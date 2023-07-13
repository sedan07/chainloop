//
// Copyright 2023 The Chainloop Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: controlplane/v1/cas_backends.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CASBackendServiceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CASBackendServiceListRequest) Reset() {
	*x = CASBackendServiceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_cas_backends_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CASBackendServiceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CASBackendServiceListRequest) ProtoMessage() {}

func (x *CASBackendServiceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_cas_backends_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CASBackendServiceListRequest.ProtoReflect.Descriptor instead.
func (*CASBackendServiceListRequest) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_cas_backends_proto_rawDescGZIP(), []int{0}
}

type CASBackendServiceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*CASBackendItem `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *CASBackendServiceListResponse) Reset() {
	*x = CASBackendServiceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_cas_backends_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CASBackendServiceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CASBackendServiceListResponse) ProtoMessage() {}

func (x *CASBackendServiceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_cas_backends_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CASBackendServiceListResponse.ProtoReflect.Descriptor instead.
func (*CASBackendServiceListResponse) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_cas_backends_proto_rawDescGZIP(), []int{1}
}

func (x *CASBackendServiceListResponse) GetResult() []*CASBackendItem {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_controlplane_v1_cas_backends_proto protoreflect.FileDescriptor

var file_controlplane_v1_cas_backends_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x73, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e,
	0x0a, 0x1c, 0x43, 0x41, 0x53, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x58,
	0x0a, 0x1d, 0x43, 0x41, 0x53, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x37, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x41, 0x53, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x7a, 0x0a, 0x11, 0x43, 0x41, 0x53, 0x42,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x41, 0x53, 0x42, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x41, 0x53, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70, 0x2d, 0x64, 0x65, 0x76,
	0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controlplane_v1_cas_backends_proto_rawDescOnce sync.Once
	file_controlplane_v1_cas_backends_proto_rawDescData = file_controlplane_v1_cas_backends_proto_rawDesc
)

func file_controlplane_v1_cas_backends_proto_rawDescGZIP() []byte {
	file_controlplane_v1_cas_backends_proto_rawDescOnce.Do(func() {
		file_controlplane_v1_cas_backends_proto_rawDescData = protoimpl.X.CompressGZIP(file_controlplane_v1_cas_backends_proto_rawDescData)
	})
	return file_controlplane_v1_cas_backends_proto_rawDescData
}

var file_controlplane_v1_cas_backends_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controlplane_v1_cas_backends_proto_goTypes = []interface{}{
	(*CASBackendServiceListRequest)(nil),  // 0: controlplane.v1.CASBackendServiceListRequest
	(*CASBackendServiceListResponse)(nil), // 1: controlplane.v1.CASBackendServiceListResponse
	(*CASBackendItem)(nil),                // 2: controlplane.v1.CASBackendItem
}
var file_controlplane_v1_cas_backends_proto_depIdxs = []int32{
	2, // 0: controlplane.v1.CASBackendServiceListResponse.result:type_name -> controlplane.v1.CASBackendItem
	0, // 1: controlplane.v1.CASBackendService.List:input_type -> controlplane.v1.CASBackendServiceListRequest
	1, // 2: controlplane.v1.CASBackendService.List:output_type -> controlplane.v1.CASBackendServiceListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_controlplane_v1_cas_backends_proto_init() }
func file_controlplane_v1_cas_backends_proto_init() {
	if File_controlplane_v1_cas_backends_proto != nil {
		return
	}
	file_controlplane_v1_response_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_controlplane_v1_cas_backends_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CASBackendServiceListRequest); i {
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
		file_controlplane_v1_cas_backends_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CASBackendServiceListResponse); i {
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
			RawDescriptor: file_controlplane_v1_cas_backends_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controlplane_v1_cas_backends_proto_goTypes,
		DependencyIndexes: file_controlplane_v1_cas_backends_proto_depIdxs,
		MessageInfos:      file_controlplane_v1_cas_backends_proto_msgTypes,
	}.Build()
	File_controlplane_v1_cas_backends_proto = out.File
	file_controlplane_v1_cas_backends_proto_rawDesc = nil
	file_controlplane_v1_cas_backends_proto_goTypes = nil
	file_controlplane_v1_cas_backends_proto_depIdxs = nil
}