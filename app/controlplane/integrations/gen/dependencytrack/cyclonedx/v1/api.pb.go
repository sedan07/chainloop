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
// source: dependencytrack/cyclonedx/v1/api.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// RegistrationRequest is what the integration expect as input of the pre-registration step
type RegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *RegistrationConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	ApiKey string              `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
}

func (x *RegistrationRequest) Reset() {
	*x = RegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dependencytrack_cyclonedx_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationRequest) ProtoMessage() {}

func (x *RegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dependencytrack_cyclonedx_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationRequest.ProtoReflect.Descriptor instead.
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return file_dependencytrack_cyclonedx_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *RegistrationRequest) GetConfig() *RegistrationConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *RegistrationRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

// AttachmentRequest is what the integration expect as input of the pre-attachment step
type AttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *AttachmentConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *AttachmentRequest) Reset() {
	*x = AttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dependencytrack_cyclonedx_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentRequest) ProtoMessage() {}

func (x *AttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dependencytrack_cyclonedx_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentRequest.ProtoReflect.Descriptor instead.
func (*AttachmentRequest) Descriptor() ([]byte, []int) {
	return file_dependencytrack_cyclonedx_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *AttachmentRequest) GetConfig() *AttachmentConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

// Structs containing the information to be stored in the integration step
type RegistrationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Support the option to automatically create projects if requested
	AllowAutoCreate bool `protobuf:"varint,2,opt,name=allow_auto_create,json=allowAutoCreate,proto3" json:"allow_auto_create,omitempty"`
}

func (x *RegistrationConfig) Reset() {
	*x = RegistrationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dependencytrack_cyclonedx_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationConfig) ProtoMessage() {}

func (x *RegistrationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_dependencytrack_cyclonedx_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationConfig.ProtoReflect.Descriptor instead.
func (*RegistrationConfig) Descriptor() ([]byte, []int) {
	return file_dependencytrack_cyclonedx_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *RegistrationConfig) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *RegistrationConfig) GetAllowAutoCreate() bool {
	if x != nil {
		return x.AllowAutoCreate
	}
	return false
}

type AttachmentConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Project:
	//
	//	*AttachmentConfig_ProjectId
	//	*AttachmentConfig_ProjectName
	Project isAttachmentConfig_Project `protobuf_oneof:"project"`
}

func (x *AttachmentConfig) Reset() {
	*x = AttachmentConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dependencytrack_cyclonedx_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentConfig) ProtoMessage() {}

func (x *AttachmentConfig) ProtoReflect() protoreflect.Message {
	mi := &file_dependencytrack_cyclonedx_v1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentConfig.ProtoReflect.Descriptor instead.
func (*AttachmentConfig) Descriptor() ([]byte, []int) {
	return file_dependencytrack_cyclonedx_v1_api_proto_rawDescGZIP(), []int{3}
}

func (m *AttachmentConfig) GetProject() isAttachmentConfig_Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (x *AttachmentConfig) GetProjectId() string {
	if x, ok := x.GetProject().(*AttachmentConfig_ProjectId); ok {
		return x.ProjectId
	}
	return ""
}

func (x *AttachmentConfig) GetProjectName() string {
	if x, ok := x.GetProject().(*AttachmentConfig_ProjectName); ok {
		return x.ProjectName
	}
	return ""
}

type isAttachmentConfig_Project interface {
	isAttachmentConfig_Project()
}

type AttachmentConfig_ProjectId struct {
	// The integration might either use a pre-configured projectID
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3,oneof"`
}

type AttachmentConfig_ProjectName struct {
	// name of the project ot be auto created
	ProjectName string `protobuf:"bytes,2,opt,name=project_name,json=projectName,proto3,oneof"`
}

func (*AttachmentConfig_ProjectId) isAttachmentConfig_Project() {}

func (*AttachmentConfig_ProjectName) isAttachmentConfig_Project() {}

var File_dependencytrack_cyclonedx_v1_api_proto protoreflect.FileDescriptor

var file_dependencytrack_cyclonedx_v1_api_proto_rawDesc = []byte{
	0x0a, 0x26, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x2f, 0x63, 0x79, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x64, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63,
	0x79, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x79, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x64, 0x78,
	0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a,
	0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x5f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x2e, 0x63, 0x79, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x64, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x20, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x22, 0x72, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5d, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x64, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x63, 0x79, 0x63, 0x6c,
	0x6f, 0x6e, 0x65, 0x64, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x61, 0x0a, 0x12, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x1f, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x75, 0x74, 0x6f,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x75, 0x74, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x22, 0x68,
	0x0a, 0x10, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x66, 0x5a, 0x64, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70,
	0x2d, 0x64, 0x65, 0x76, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x2f, 0x63, 0x79, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x64, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dependencytrack_cyclonedx_v1_api_proto_rawDescOnce sync.Once
	file_dependencytrack_cyclonedx_v1_api_proto_rawDescData = file_dependencytrack_cyclonedx_v1_api_proto_rawDesc
)

func file_dependencytrack_cyclonedx_v1_api_proto_rawDescGZIP() []byte {
	file_dependencytrack_cyclonedx_v1_api_proto_rawDescOnce.Do(func() {
		file_dependencytrack_cyclonedx_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_dependencytrack_cyclonedx_v1_api_proto_rawDescData)
	})
	return file_dependencytrack_cyclonedx_v1_api_proto_rawDescData
}

var file_dependencytrack_cyclonedx_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dependencytrack_cyclonedx_v1_api_proto_goTypes = []interface{}{
	(*RegistrationRequest)(nil), // 0: integrations.dependencytrack.cyclonedx.v1.RegistrationRequest
	(*AttachmentRequest)(nil),   // 1: integrations.dependencytrack.cyclonedx.v1.AttachmentRequest
	(*RegistrationConfig)(nil),  // 2: integrations.dependencytrack.cyclonedx.v1.RegistrationConfig
	(*AttachmentConfig)(nil),    // 3: integrations.dependencytrack.cyclonedx.v1.AttachmentConfig
}
var file_dependencytrack_cyclonedx_v1_api_proto_depIdxs = []int32{
	2, // 0: integrations.dependencytrack.cyclonedx.v1.RegistrationRequest.config:type_name -> integrations.dependencytrack.cyclonedx.v1.RegistrationConfig
	3, // 1: integrations.dependencytrack.cyclonedx.v1.AttachmentRequest.config:type_name -> integrations.dependencytrack.cyclonedx.v1.AttachmentConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dependencytrack_cyclonedx_v1_api_proto_init() }
func file_dependencytrack_cyclonedx_v1_api_proto_init() {
	if File_dependencytrack_cyclonedx_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dependencytrack_cyclonedx_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationRequest); i {
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
		file_dependencytrack_cyclonedx_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentRequest); i {
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
		file_dependencytrack_cyclonedx_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationConfig); i {
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
		file_dependencytrack_cyclonedx_v1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentConfig); i {
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
	file_dependencytrack_cyclonedx_v1_api_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*AttachmentConfig_ProjectId)(nil),
		(*AttachmentConfig_ProjectName)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dependencytrack_cyclonedx_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dependencytrack_cyclonedx_v1_api_proto_goTypes,
		DependencyIndexes: file_dependencytrack_cyclonedx_v1_api_proto_depIdxs,
		MessageInfos:      file_dependencytrack_cyclonedx_v1_api_proto_msgTypes,
	}.Build()
	File_dependencytrack_cyclonedx_v1_api_proto = out.File
	file_dependencytrack_cyclonedx_v1_api_proto_rawDesc = nil
	file_dependencytrack_cyclonedx_v1_api_proto_goTypes = nil
	file_dependencytrack_cyclonedx_v1_api_proto_depIdxs = nil
}
