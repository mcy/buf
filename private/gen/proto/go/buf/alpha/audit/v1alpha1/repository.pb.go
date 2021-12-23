// Copyright 2020-2021 Buf Technologies, Inc.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: buf/alpha/audit/v1alpha1/repository.proto

package auditv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BufAlphaRegistryV1Alpha1Visibility int32

const (
	BufAlphaRegistryV1Alpha1Visibility_BUF_ALPHA_REGISTRY_V1_ALPHA1_VISIBILITY_UNSPECIFIED BufAlphaRegistryV1Alpha1Visibility = 0
	BufAlphaRegistryV1Alpha1Visibility_BUF_ALPHA_REGISTRY_V1_ALPHA1_VISIBILITY_PUBLIC      BufAlphaRegistryV1Alpha1Visibility = 1
	BufAlphaRegistryV1Alpha1Visibility_BUF_ALPHA_REGISTRY_V1_ALPHA1_VISIBILITY_PRIVATE     BufAlphaRegistryV1Alpha1Visibility = 2
)

// Enum value maps for BufAlphaRegistryV1Alpha1Visibility.
var (
	BufAlphaRegistryV1Alpha1Visibility_name = map[int32]string{
		0: "BUF_ALPHA_REGISTRY_V1_ALPHA1_VISIBILITY_UNSPECIFIED",
		1: "BUF_ALPHA_REGISTRY_V1_ALPHA1_VISIBILITY_PUBLIC",
		2: "BUF_ALPHA_REGISTRY_V1_ALPHA1_VISIBILITY_PRIVATE",
	}
	BufAlphaRegistryV1Alpha1Visibility_value = map[string]int32{
		"BUF_ALPHA_REGISTRY_V1_ALPHA1_VISIBILITY_UNSPECIFIED": 0,
		"BUF_ALPHA_REGISTRY_V1_ALPHA1_VISIBILITY_PUBLIC":      1,
		"BUF_ALPHA_REGISTRY_V1_ALPHA1_VISIBILITY_PRIVATE":     2,
	}
)

func (x BufAlphaRegistryV1Alpha1Visibility) Enum() *BufAlphaRegistryV1Alpha1Visibility {
	p := new(BufAlphaRegistryV1Alpha1Visibility)
	*p = x
	return p
}

func (x BufAlphaRegistryV1Alpha1Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BufAlphaRegistryV1Alpha1Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_alpha_audit_v1alpha1_repository_proto_enumTypes[0].Descriptor()
}

func (BufAlphaRegistryV1Alpha1Visibility) Type() protoreflect.EnumType {
	return &file_buf_alpha_audit_v1alpha1_repository_proto_enumTypes[0]
}

func (x BufAlphaRegistryV1Alpha1Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BufAlphaRegistryV1Alpha1Visibility.Descriptor instead.
func (BufAlphaRegistryV1Alpha1Visibility) EnumDescriptor() ([]byte, []int) {
	return file_buf_alpha_audit_v1alpha1_repository_proto_rawDescGZIP(), []int{0}
}

type BufAlphaRegistryV1Alpha1RepositoryBranch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Field number '3' reserved for the update_time.
	Name         string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	RepositoryId string `protobuf:"bytes,5,opt,name=repository_id,json=repositoryId,proto3" json:"repository_id,omitempty"`
}

func (x *BufAlphaRegistryV1Alpha1RepositoryBranch) Reset() {
	*x = BufAlphaRegistryV1Alpha1RepositoryBranch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_audit_v1alpha1_repository_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BufAlphaRegistryV1Alpha1RepositoryBranch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BufAlphaRegistryV1Alpha1RepositoryBranch) ProtoMessage() {}

func (x *BufAlphaRegistryV1Alpha1RepositoryBranch) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_audit_v1alpha1_repository_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BufAlphaRegistryV1Alpha1RepositoryBranch.ProtoReflect.Descriptor instead.
func (*BufAlphaRegistryV1Alpha1RepositoryBranch) Descriptor() ([]byte, []int) {
	return file_buf_alpha_audit_v1alpha1_repository_proto_rawDescGZIP(), []int{0}
}

func (x *BufAlphaRegistryV1Alpha1RepositoryBranch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BufAlphaRegistryV1Alpha1RepositoryBranch) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *BufAlphaRegistryV1Alpha1RepositoryBranch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BufAlphaRegistryV1Alpha1RepositoryBranch) GetRepositoryId() string {
	if x != nil {
		return x.RepositoryId
	}
	return ""
}

type BufAlphaRegistryV1Alpha1RepositoryTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Field number '3' reserved for the update_time.
	Name       string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	CommitName string `protobuf:"bytes,5,opt,name=commit_name,json=commitName,proto3" json:"commit_name,omitempty"`
	Author     string `protobuf:"bytes,6,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *BufAlphaRegistryV1Alpha1RepositoryTag) Reset() {
	*x = BufAlphaRegistryV1Alpha1RepositoryTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_audit_v1alpha1_repository_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BufAlphaRegistryV1Alpha1RepositoryTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BufAlphaRegistryV1Alpha1RepositoryTag) ProtoMessage() {}

func (x *BufAlphaRegistryV1Alpha1RepositoryTag) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_audit_v1alpha1_repository_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BufAlphaRegistryV1Alpha1RepositoryTag.ProtoReflect.Descriptor instead.
func (*BufAlphaRegistryV1Alpha1RepositoryTag) Descriptor() ([]byte, []int) {
	return file_buf_alpha_audit_v1alpha1_repository_proto_rawDescGZIP(), []int{1}
}

func (x *BufAlphaRegistryV1Alpha1RepositoryTag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BufAlphaRegistryV1Alpha1RepositoryTag) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *BufAlphaRegistryV1Alpha1RepositoryTag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BufAlphaRegistryV1Alpha1RepositoryTag) GetCommitName() string {
	if x != nil {
		return x.CommitName
	}
	return ""
}

func (x *BufAlphaRegistryV1Alpha1RepositoryTag) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type BufAlphaRegistryV1Alpha1RepositoryCommit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateTime       *timestamppb.Timestamp                   `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Digest           string                                   `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	Name             string                                   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Branch           string                                   `protobuf:"bytes,5,opt,name=branch,proto3" json:"branch,omitempty"`
	CommitSequenceId int64                                    `protobuf:"varint,6,opt,name=commit_sequence_id,json=commitSequenceId,proto3" json:"commit_sequence_id,omitempty"`
	Author           string                                   `protobuf:"bytes,7,opt,name=author,proto3" json:"author,omitempty"`
	Tags             []*BufAlphaRegistryV1Alpha1RepositoryTag `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *BufAlphaRegistryV1Alpha1RepositoryCommit) Reset() {
	*x = BufAlphaRegistryV1Alpha1RepositoryCommit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_audit_v1alpha1_repository_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BufAlphaRegistryV1Alpha1RepositoryCommit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BufAlphaRegistryV1Alpha1RepositoryCommit) ProtoMessage() {}

func (x *BufAlphaRegistryV1Alpha1RepositoryCommit) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_audit_v1alpha1_repository_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BufAlphaRegistryV1Alpha1RepositoryCommit.ProtoReflect.Descriptor instead.
func (*BufAlphaRegistryV1Alpha1RepositoryCommit) Descriptor() ([]byte, []int) {
	return file_buf_alpha_audit_v1alpha1_repository_proto_rawDescGZIP(), []int{2}
}

func (x *BufAlphaRegistryV1Alpha1RepositoryCommit) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BufAlphaRegistryV1Alpha1RepositoryCommit) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *BufAlphaRegistryV1Alpha1RepositoryCommit) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *BufAlphaRegistryV1Alpha1RepositoryCommit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BufAlphaRegistryV1Alpha1RepositoryCommit) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *BufAlphaRegistryV1Alpha1RepositoryCommit) GetCommitSequenceId() int64 {
	if x != nil {
		return x.CommitSequenceId
	}
	return 0
}

func (x *BufAlphaRegistryV1Alpha1RepositoryCommit) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *BufAlphaRegistryV1Alpha1RepositoryCommit) GetTags() []*BufAlphaRegistryV1Alpha1RepositoryTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

type BufAlphaRegistryV1Alpha1RepositoryTrack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Field number '3' reserved for the update_time.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BufAlphaRegistryV1Alpha1RepositoryTrack) Reset() {
	*x = BufAlphaRegistryV1Alpha1RepositoryTrack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_audit_v1alpha1_repository_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BufAlphaRegistryV1Alpha1RepositoryTrack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BufAlphaRegistryV1Alpha1RepositoryTrack) ProtoMessage() {}

func (x *BufAlphaRegistryV1Alpha1RepositoryTrack) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_audit_v1alpha1_repository_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BufAlphaRegistryV1Alpha1RepositoryTrack.ProtoReflect.Descriptor instead.
func (*BufAlphaRegistryV1Alpha1RepositoryTrack) Descriptor() ([]byte, []int) {
	return file_buf_alpha_audit_v1alpha1_repository_proto_rawDescGZIP(), []int{3}
}

func (x *BufAlphaRegistryV1Alpha1RepositoryTrack) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BufAlphaRegistryV1Alpha1RepositoryTrack) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *BufAlphaRegistryV1Alpha1RepositoryTrack) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_buf_alpha_audit_v1alpha1_repository_proto protoreflect.FileDescriptor

var file_buf_alpha_audit_v1alpha1_repository_proto_rawDesc = []byte{
	0x0a, 0x29, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x62, 0x75, 0x66,
	0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x28, 0x42, 0x75, 0x66, 0x41, 0x6c,
	0x70, 0x68, 0x61, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x56, 0x31, 0x41, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0xc1, 0x01, 0x0a, 0x25, 0x42, 0x75,
	0x66, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x56, 0x31,
	0x41, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0xd6, 0x02,
	0x0a, 0x28, 0x42, 0x75, 0x66, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x56, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x2c, 0x0a, 0x12, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x53, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3f, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x75, 0x66, 0x41, 0x6c,
	0x70, 0x68, 0x61, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x56, 0x31, 0x41, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x61, 0x67,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x27, 0x42, 0x75, 0x66, 0x41, 0x6c,
	0x70, 0x68, 0x61, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x56, 0x31, 0x41, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x2a, 0xc6, 0x01, 0x0a, 0x22, 0x42, 0x75, 0x66, 0x41, 0x6c, 0x70, 0x68, 0x61,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x56, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x33, 0x42, 0x55,
	0x46, 0x5f, 0x41, 0x4c, 0x50, 0x48, 0x41, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x59,
	0x5f, 0x56, 0x31, 0x5f, 0x41, 0x4c, 0x50, 0x48, 0x41, 0x31, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x32, 0x0a, 0x2e, 0x42, 0x55, 0x46, 0x5f, 0x41, 0x4c, 0x50, 0x48, 0x41,
	0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x59, 0x5f, 0x56, 0x31, 0x5f, 0x41, 0x4c, 0x50,
	0x48, 0x41, 0x31, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x50,
	0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x01, 0x12, 0x33, 0x0a, 0x2f, 0x42, 0x55, 0x46, 0x5f, 0x41,
	0x4c, 0x50, 0x48, 0x41, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x59, 0x5f, 0x56, 0x31,
	0x5f, 0x41, 0x4c, 0x50, 0x48, 0x41, 0x31, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49,
	0x54, 0x59, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x02, 0x42, 0x87, 0x02, 0x0a,
	0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0f, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x41, 0x41, 0xaa, 0x02, 0x18, 0x42, 0x75,
	0x66, 0x2e, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x18, 0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c, 0x70,
	0x68, 0x61, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xe2, 0x02, 0x24, 0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x42, 0x75, 0x66, 0x3a, 0x3a,
	0x41, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x3a, 0x41, 0x75, 0x64, 0x69, 0x74, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_alpha_audit_v1alpha1_repository_proto_rawDescOnce sync.Once
	file_buf_alpha_audit_v1alpha1_repository_proto_rawDescData = file_buf_alpha_audit_v1alpha1_repository_proto_rawDesc
)

func file_buf_alpha_audit_v1alpha1_repository_proto_rawDescGZIP() []byte {
	file_buf_alpha_audit_v1alpha1_repository_proto_rawDescOnce.Do(func() {
		file_buf_alpha_audit_v1alpha1_repository_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_alpha_audit_v1alpha1_repository_proto_rawDescData)
	})
	return file_buf_alpha_audit_v1alpha1_repository_proto_rawDescData
}

var file_buf_alpha_audit_v1alpha1_repository_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buf_alpha_audit_v1alpha1_repository_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_buf_alpha_audit_v1alpha1_repository_proto_goTypes = []interface{}{
	(BufAlphaRegistryV1Alpha1Visibility)(0),          // 0: buf.alpha.audit.v1alpha1.BufAlphaRegistryV1Alpha1Visibility
	(*BufAlphaRegistryV1Alpha1RepositoryBranch)(nil), // 1: buf.alpha.audit.v1alpha1.BufAlphaRegistryV1Alpha1RepositoryBranch
	(*BufAlphaRegistryV1Alpha1RepositoryTag)(nil),    // 2: buf.alpha.audit.v1alpha1.BufAlphaRegistryV1Alpha1RepositoryTag
	(*BufAlphaRegistryV1Alpha1RepositoryCommit)(nil), // 3: buf.alpha.audit.v1alpha1.BufAlphaRegistryV1Alpha1RepositoryCommit
	(*BufAlphaRegistryV1Alpha1RepositoryTrack)(nil),  // 4: buf.alpha.audit.v1alpha1.BufAlphaRegistryV1Alpha1RepositoryTrack
	(*timestamppb.Timestamp)(nil),                    // 5: google.protobuf.Timestamp
}
var file_buf_alpha_audit_v1alpha1_repository_proto_depIdxs = []int32{
	5, // 0: buf.alpha.audit.v1alpha1.BufAlphaRegistryV1Alpha1RepositoryBranch.create_time:type_name -> google.protobuf.Timestamp
	5, // 1: buf.alpha.audit.v1alpha1.BufAlphaRegistryV1Alpha1RepositoryTag.create_time:type_name -> google.protobuf.Timestamp
	5, // 2: buf.alpha.audit.v1alpha1.BufAlphaRegistryV1Alpha1RepositoryCommit.create_time:type_name -> google.protobuf.Timestamp
	2, // 3: buf.alpha.audit.v1alpha1.BufAlphaRegistryV1Alpha1RepositoryCommit.tags:type_name -> buf.alpha.audit.v1alpha1.BufAlphaRegistryV1Alpha1RepositoryTag
	5, // 4: buf.alpha.audit.v1alpha1.BufAlphaRegistryV1Alpha1RepositoryTrack.create_time:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_buf_alpha_audit_v1alpha1_repository_proto_init() }
func file_buf_alpha_audit_v1alpha1_repository_proto_init() {
	if File_buf_alpha_audit_v1alpha1_repository_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_alpha_audit_v1alpha1_repository_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BufAlphaRegistryV1Alpha1RepositoryBranch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buf_alpha_audit_v1alpha1_repository_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BufAlphaRegistryV1Alpha1RepositoryTag); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buf_alpha_audit_v1alpha1_repository_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BufAlphaRegistryV1Alpha1RepositoryCommit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buf_alpha_audit_v1alpha1_repository_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BufAlphaRegistryV1Alpha1RepositoryTrack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_alpha_audit_v1alpha1_repository_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_alpha_audit_v1alpha1_repository_proto_goTypes,
		DependencyIndexes: file_buf_alpha_audit_v1alpha1_repository_proto_depIdxs,
		EnumInfos:         file_buf_alpha_audit_v1alpha1_repository_proto_enumTypes,
		MessageInfos:      file_buf_alpha_audit_v1alpha1_repository_proto_msgTypes,
	}.Build()
	File_buf_alpha_audit_v1alpha1_repository_proto = out.File
	file_buf_alpha_audit_v1alpha1_repository_proto_rawDesc = nil
	file_buf_alpha_audit_v1alpha1_repository_proto_goTypes = nil
	file_buf_alpha_audit_v1alpha1_repository_proto_depIdxs = nil
}
