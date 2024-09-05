// Copyright 2020-2024 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package buflsp implements a language server for Protobuf.
//
// The main entry-point of this package is the New() function, which creates a new LSP server.
package buflsp

import (
	"context"
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/bufbuild/buf/private/buf/bufcli"
	"github.com/bufbuild/buf/private/buf/bufctl"
	"github.com/bufbuild/buf/private/bufpkg/bufmodule"
	"github.com/bufbuild/buf/private/gen/data/datawkt"
	"github.com/bufbuild/buf/private/pkg/app/appext"
	"github.com/bufbuild/buf/private/pkg/storage"
	"github.com/bufbuild/buf/private/pkg/storage/storageos"
	"github.com/bufbuild/buf/private/pkg/tracing"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// ErrNotInit is returned by LSP server methods that are called without first
// initializing the server.
var ErrNotInit = fmt.Errorf("the first call to the server must be the %q method", protocol.MethodInitialize)

// New spawns a new LSP server, listening on the given stream.
//
// Returns a context for managing the server.
func Serve(
	ctx context.Context,
	container appext.Container,
	stream jsonrpc2.Stream,
) (jsonrpc2.Conn, error) {
	controller, err := bufcli.NewController(container)
	if err != nil {
		return nil, err
	}

	bucketProvider := storageos.NewProvider(storageos.ProviderWithSymlinks())
	bucket, err := bucketProvider.NewReadWriteBucket("/",
		storageos.ReadWriteBucketWithSymlinksIfSupported())
	if err != nil {
		return nil, err
	}

	conn := jsonrpc2.NewConn(stream)
	server := &server{
		conn:   conn,
		client: protocol.ClientDispatcher(conn, container.Logger()),

		logger:     container.Logger(),
		tracer:     tracing.NewTracer(container.Tracer()),
		controller: controller,
		rootBucket: bucket,
	}
	server.files = newFiles(server)
	off := protocol.TraceOff
	server.traceValue.Store(&off)

	conn.Go(ctx, server.makeHandler())
	return conn, nil
}

// server is an LSP server.
//
// This type contains all the necessary book-keeping for keeping the server running.
// Its handler methods are not defined in buflsp.go; they are defined in other files, grouped
// according to the groupings in https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification
type server struct {
	nyiServer

	conn   jsonrpc2.Conn
	client protocol.Client

	logger     *zap.Logger
	tracer     tracing.Tracer
	controller bufctl.Controller
	rootBucket storage.ReadBucket
	files      *files

	wktModuleSet bufmodule.ModuleSet

	// These are atomics, because they are read often add written to
	// almost never, but potentially concurrently. Having them side-by-side
	// is fine; they are almost never written to so false sharing is not a
	// concern.
	initParams atomic.Pointer[protocol.InitializeParams]
	traceValue atomic.Pointer[protocol.TraceValue]
}

// checkInit is a helper that checks if initialization has occurred and, if not,
// returns an appropriate error.
func (server *server) checkInit() error {
	if server.initParams.Load() != nil {
		return nil
	}
	return ErrNotInit
}

// init performs *actual* initialization of the server. This is called by Initialize().
//
// It may only be called once for a given server.
func (server *server) init(ctx context.Context, params *protocol.InitializeParams) error {
	if server.initParams.Load() != nil {
		return fmt.Errorf("called the %q method more than once", protocol.MethodInitialize)
	}
	server.initParams.Store(params)

	// TODO: set up logging. We need to forward everything from server.logger through to
	// the client, if tracing is turned on. The right way to do this is with an extra
	// goroutine and some channels.

	// Load the WKTs asap. They're always needed and don't need to hit the filesystem.
	if err := server.loadWTKs(ctx); err != nil {
		return err
	}

	return nil
}

// loadWKTs loads a ModuleSet for the well-known types.
func (server *server) loadWTKs(ctx context.Context) (err error) {
	builder := bufmodule.NewModuleSetBuilder(
		ctx,
		server.tracer,
		bufmodule.NopModuleDataProvider,
		bufmodule.NopCommitProvider,
	)
	// DO NOT MERGE: is isTarget necessary?
	builder.AddLocalModule(datawkt.ReadBucket, "." /*isTarget=*/, true)
	server.wktModuleSet, err = builder.Build()
	return
}

// makeHandler constructs an RPC handler that wraps the default one from jsonrpc2. This allows us
// to inject debug logging, tracing, and timeouts to requests.
func (server *server) makeHandler() jsonrpc2.Handler {
	actual := protocol.ServerHandler(server, nil)
	return func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) (err error) {
		ctx, span := server.tracer.Start(ctx,
			tracing.WithErr(&err),
			tracing.WithAttributes(attribute.String("method", req.Method())),
		)
		defer span.End()

		server.logger.Debug(
			"processing request",
			zap.String("method", req.Method()),
			zap.ByteString("params", req.Params()),
		)

		// Each request is handled in a separate goroutine, and has a fixed timeout.
		// This is to enforce responsiveness on the LSP side: if something is going to take
		// a long time, it should be offloaded.
		ctx, done := context.WithTimeout(ctx, 3*time.Second)
		ctx = withReentrancy(ctx)

		go func() {
			defer done()
			err = actual(ctx, server.adaptReplier(reply, req), req)
		}()

		<-ctx.Done()
		if ctx.Err() == context.DeadlineExceeded {
			// Don't return this error; that will kill the whole server!
			server.logger.Sugar().Errorf(
				"timed out while handling %s; this is likely a bug", req.Method())
		}
		return
	}
}

// adaptReplier wraps a jsonrpc2.Replier, allowing us to inject logging and tracing and so on.
func (server *server) adaptReplier(reply jsonrpc2.Replier, req jsonrpc2.Request) jsonrpc2.Replier {
	return func(ctx context.Context, result any, err error) error {
		if err != nil {
			server.logger.Debug(
				"responding",
				zap.String("method", req.Method()),
				zap.Error(err),
			)
		} else {
			json, _ := json.Marshal(result)
			server.logger.Debug(
				"responding",
				zap.String("method", req.Method()),
				zap.ByteString("params", json),
			)
		}

		return reply(ctx, result, err)
	}
}
