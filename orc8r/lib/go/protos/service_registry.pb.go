//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.10.0
// source: orc8r/protos/service_registry.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ListAllServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []string `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ListAllServicesResponse) Reset() {
	*x = ListAllServicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_service_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllServicesResponse) ProtoMessage() {}

func (x *ListAllServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_service_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllServicesResponse.ProtoReflect.Descriptor instead.
func (*ListAllServicesResponse) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_service_registry_proto_rawDescGZIP(), []int{0}
}

func (x *ListAllServicesResponse) GetServices() []string {
	if x != nil {
		return x.Services
	}
	return nil
}

type FindServicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *FindServicesRequest) Reset() {
	*x = FindServicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_service_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindServicesRequest) ProtoMessage() {}

func (x *FindServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_service_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindServicesRequest.ProtoReflect.Descriptor instead.
func (*FindServicesRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_service_registry_proto_rawDescGZIP(), []int{1}
}

func (x *FindServicesRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type FindServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []string `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *FindServicesResponse) Reset() {
	*x = FindServicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_service_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindServicesResponse) ProtoMessage() {}

func (x *FindServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_service_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindServicesResponse.ProtoReflect.Descriptor instead.
func (*FindServicesResponse) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_service_registry_proto_rawDescGZIP(), []int{2}
}

func (x *FindServicesResponse) GetServices() []string {
	if x != nil {
		return x.Services
	}
	return nil
}

type GetServiceAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *GetServiceAddressRequest) Reset() {
	*x = GetServiceAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_service_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceAddressRequest) ProtoMessage() {}

func (x *GetServiceAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_service_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceAddressRequest.ProtoReflect.Descriptor instead.
func (*GetServiceAddressRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_service_registry_proto_rawDescGZIP(), []int{3}
}

func (x *GetServiceAddressRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type GetServiceAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetServiceAddressResponse) Reset() {
	*x = GetServiceAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_service_registry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceAddressResponse) ProtoMessage() {}

func (x *GetServiceAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_service_registry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceAddressResponse.ProtoReflect.Descriptor instead.
func (*GetServiceAddressResponse) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_service_registry_proto_rawDescGZIP(), []int{4}
}

func (x *GetServiceAddressResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetHttpServerAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *GetHttpServerAddressRequest) Reset() {
	*x = GetHttpServerAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_service_registry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHttpServerAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHttpServerAddressRequest) ProtoMessage() {}

func (x *GetHttpServerAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_service_registry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHttpServerAddressRequest.ProtoReflect.Descriptor instead.
func (*GetHttpServerAddressRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_service_registry_proto_rawDescGZIP(), []int{5}
}

func (x *GetHttpServerAddressRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type GetHttpServerAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetHttpServerAddressResponse) Reset() {
	*x = GetHttpServerAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_service_registry_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHttpServerAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHttpServerAddressResponse) ProtoMessage() {}

func (x *GetHttpServerAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_service_registry_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHttpServerAddressResponse.ProtoReflect.Descriptor instead.
func (*GetHttpServerAddressResponse) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_service_registry_proto_rawDescGZIP(), []int{6}
}

func (x *GetHttpServerAddressResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetAnnotationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service    string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Annotation string `protobuf:"bytes,2,opt,name=annotation,proto3" json:"annotation,omitempty"`
}

func (x *GetAnnotationRequest) Reset() {
	*x = GetAnnotationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_service_registry_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnnotationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnnotationRequest) ProtoMessage() {}

func (x *GetAnnotationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_service_registry_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnnotationRequest.ProtoReflect.Descriptor instead.
func (*GetAnnotationRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_service_registry_proto_rawDescGZIP(), []int{7}
}

func (x *GetAnnotationRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *GetAnnotationRequest) GetAnnotation() string {
	if x != nil {
		return x.Annotation
	}
	return ""
}

type GetAnnotationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnnotationValue string `protobuf:"bytes,1,opt,name=annotation_value,json=annotationValue,proto3" json:"annotation_value,omitempty"`
}

func (x *GetAnnotationResponse) Reset() {
	*x = GetAnnotationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_service_registry_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnnotationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnnotationResponse) ProtoMessage() {}

func (x *GetAnnotationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_service_registry_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnnotationResponse.ProtoReflect.Descriptor instead.
func (*GetAnnotationResponse) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_service_registry_proto_rawDescGZIP(), []int{8}
}

func (x *GetAnnotationResponse) GetAnnotationValue() string {
	if x != nil {
		return x.AnnotationValue
	}
	return ""
}

var File_orc8r_protos_service_registry_proto protoreflect.FileDescriptor

var file_orc8r_protos_service_registry_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63,
	0x38, 0x72, 0x1a, 0x19, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a,
	0x17, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x22, 0x32, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x35, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x37, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x48, 0x74, 0x74, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x38, 0x0a, 0x1c, 0x47,
	0x65, 0x74, 0x48, 0x74, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x50, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x10, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xe5, 0x03, 0x0a, 0x0f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12,
	0x4c, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x11, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x24, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72,
	0x63, 0x38, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a,
	0x0c, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x20, 0x2e,
	0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x2e, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x48, 0x74, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x28, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x48, 0x74, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x74,
	0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6f, 0x72, 0x63,
	0x38, 0x72, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orc8r_protos_service_registry_proto_rawDescOnce sync.Once
	file_orc8r_protos_service_registry_proto_rawDescData = file_orc8r_protos_service_registry_proto_rawDesc
)

func file_orc8r_protos_service_registry_proto_rawDescGZIP() []byte {
	file_orc8r_protos_service_registry_proto_rawDescOnce.Do(func() {
		file_orc8r_protos_service_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_orc8r_protos_service_registry_proto_rawDescData)
	})
	return file_orc8r_protos_service_registry_proto_rawDescData
}

var file_orc8r_protos_service_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_orc8r_protos_service_registry_proto_goTypes = []interface{}{
	(*ListAllServicesResponse)(nil),      // 0: magma.orc8r.ListAllServicesResponse
	(*FindServicesRequest)(nil),          // 1: magma.orc8r.FindServicesRequest
	(*FindServicesResponse)(nil),         // 2: magma.orc8r.FindServicesResponse
	(*GetServiceAddressRequest)(nil),     // 3: magma.orc8r.GetServiceAddressRequest
	(*GetServiceAddressResponse)(nil),    // 4: magma.orc8r.GetServiceAddressResponse
	(*GetHttpServerAddressRequest)(nil),  // 5: magma.orc8r.GetHttpServerAddressRequest
	(*GetHttpServerAddressResponse)(nil), // 6: magma.orc8r.GetHttpServerAddressResponse
	(*GetAnnotationRequest)(nil),         // 7: magma.orc8r.GetAnnotationRequest
	(*GetAnnotationResponse)(nil),        // 8: magma.orc8r.GetAnnotationResponse
	(*Void)(nil),                         // 9: magma.orc8r.Void
}
var file_orc8r_protos_service_registry_proto_depIdxs = []int32{
	9, // 0: magma.orc8r.ServiceRegistry.ListAllServices:input_type -> magma.orc8r.Void
	1, // 1: magma.orc8r.ServiceRegistry.FindServices:input_type -> magma.orc8r.FindServicesRequest
	3, // 2: magma.orc8r.ServiceRegistry.GetServiceAddress:input_type -> magma.orc8r.GetServiceAddressRequest
	5, // 3: magma.orc8r.ServiceRegistry.GetHttpServerAddress:input_type -> magma.orc8r.GetHttpServerAddressRequest
	7, // 4: magma.orc8r.ServiceRegistry.GetAnnotation:input_type -> magma.orc8r.GetAnnotationRequest
	0, // 5: magma.orc8r.ServiceRegistry.ListAllServices:output_type -> magma.orc8r.ListAllServicesResponse
	2, // 6: magma.orc8r.ServiceRegistry.FindServices:output_type -> magma.orc8r.FindServicesResponse
	4, // 7: magma.orc8r.ServiceRegistry.GetServiceAddress:output_type -> magma.orc8r.GetServiceAddressResponse
	6, // 8: magma.orc8r.ServiceRegistry.GetHttpServerAddress:output_type -> magma.orc8r.GetHttpServerAddressResponse
	8, // 9: magma.orc8r.ServiceRegistry.GetAnnotation:output_type -> magma.orc8r.GetAnnotationResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orc8r_protos_service_registry_proto_init() }
func file_orc8r_protos_service_registry_proto_init() {
	if File_orc8r_protos_service_registry_proto != nil {
		return
	}
	file_orc8r_protos_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_orc8r_protos_service_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllServicesResponse); i {
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
		file_orc8r_protos_service_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindServicesRequest); i {
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
		file_orc8r_protos_service_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindServicesResponse); i {
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
		file_orc8r_protos_service_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceAddressRequest); i {
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
		file_orc8r_protos_service_registry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceAddressResponse); i {
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
		file_orc8r_protos_service_registry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHttpServerAddressRequest); i {
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
		file_orc8r_protos_service_registry_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHttpServerAddressResponse); i {
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
		file_orc8r_protos_service_registry_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnnotationRequest); i {
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
		file_orc8r_protos_service_registry_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnnotationResponse); i {
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
			RawDescriptor: file_orc8r_protos_service_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orc8r_protos_service_registry_proto_goTypes,
		DependencyIndexes: file_orc8r_protos_service_registry_proto_depIdxs,
		MessageInfos:      file_orc8r_protos_service_registry_proto_msgTypes,
	}.Build()
	File_orc8r_protos_service_registry_proto = out.File
	file_orc8r_protos_service_registry_proto_rawDesc = nil
	file_orc8r_protos_service_registry_proto_goTypes = nil
	file_orc8r_protos_service_registry_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceRegistryClient is the client API for ServiceRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceRegistryClient interface {
	// ListAllServices returns the service name of all services in the registry.
	ListAllServices(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ListAllServicesResponse, error)
	// FindServices returns the names of all services that have the passed label.
	FindServices(ctx context.Context, in *FindServicesRequest, opts ...grpc.CallOption) (*FindServicesResponse, error)
	// GetServiceAddress returns the service address of a given service.
	GetServiceAddress(ctx context.Context, in *GetServiceAddressRequest, opts ...grpc.CallOption) (*GetServiceAddressResponse, error)
	// GetHttpServerAddress returns the address for a service's HTTP server.
	GetHttpServerAddress(ctx context.Context, in *GetHttpServerAddressRequest, opts ...grpc.CallOption) (*GetHttpServerAddressResponse, error)
	// GetAnnotation returns the annotation value for the passed annotation
	// name.
	GetAnnotation(ctx context.Context, in *GetAnnotationRequest, opts ...grpc.CallOption) (*GetAnnotationResponse, error)
}

type serviceRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceRegistryClient(cc grpc.ClientConnInterface) ServiceRegistryClient {
	return &serviceRegistryClient{cc}
}

func (c *serviceRegistryClient) ListAllServices(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ListAllServicesResponse, error) {
	out := new(ListAllServicesResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.ServiceRegistry/ListAllServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceRegistryClient) FindServices(ctx context.Context, in *FindServicesRequest, opts ...grpc.CallOption) (*FindServicesResponse, error) {
	out := new(FindServicesResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.ServiceRegistry/FindServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceRegistryClient) GetServiceAddress(ctx context.Context, in *GetServiceAddressRequest, opts ...grpc.CallOption) (*GetServiceAddressResponse, error) {
	out := new(GetServiceAddressResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.ServiceRegistry/GetServiceAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceRegistryClient) GetHttpServerAddress(ctx context.Context, in *GetHttpServerAddressRequest, opts ...grpc.CallOption) (*GetHttpServerAddressResponse, error) {
	out := new(GetHttpServerAddressResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.ServiceRegistry/GetHttpServerAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceRegistryClient) GetAnnotation(ctx context.Context, in *GetAnnotationRequest, opts ...grpc.CallOption) (*GetAnnotationResponse, error) {
	out := new(GetAnnotationResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.ServiceRegistry/GetAnnotation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceRegistryServer is the server API for ServiceRegistry service.
type ServiceRegistryServer interface {
	// ListAllServices returns the service name of all services in the registry.
	ListAllServices(context.Context, *Void) (*ListAllServicesResponse, error)
	// FindServices returns the names of all services that have the passed label.
	FindServices(context.Context, *FindServicesRequest) (*FindServicesResponse, error)
	// GetServiceAddress returns the service address of a given service.
	GetServiceAddress(context.Context, *GetServiceAddressRequest) (*GetServiceAddressResponse, error)
	// GetHttpServerAddress returns the address for a service's HTTP server.
	GetHttpServerAddress(context.Context, *GetHttpServerAddressRequest) (*GetHttpServerAddressResponse, error)
	// GetAnnotation returns the annotation value for the passed annotation
	// name.
	GetAnnotation(context.Context, *GetAnnotationRequest) (*GetAnnotationResponse, error)
}

// UnimplementedServiceRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedServiceRegistryServer struct {
}

func (*UnimplementedServiceRegistryServer) ListAllServices(context.Context, *Void) (*ListAllServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllServices not implemented")
}
func (*UnimplementedServiceRegistryServer) FindServices(context.Context, *FindServicesRequest) (*FindServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindServices not implemented")
}
func (*UnimplementedServiceRegistryServer) GetServiceAddress(context.Context, *GetServiceAddressRequest) (*GetServiceAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceAddress not implemented")
}
func (*UnimplementedServiceRegistryServer) GetHttpServerAddress(context.Context, *GetHttpServerAddressRequest) (*GetHttpServerAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHttpServerAddress not implemented")
}
func (*UnimplementedServiceRegistryServer) GetAnnotation(context.Context, *GetAnnotationRequest) (*GetAnnotationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnnotation not implemented")
}

func RegisterServiceRegistryServer(s *grpc.Server, srv ServiceRegistryServer) {
	s.RegisterService(&_ServiceRegistry_serviceDesc, srv)
}

func _ServiceRegistry_ListAllServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceRegistryServer).ListAllServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.ServiceRegistry/ListAllServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceRegistryServer).ListAllServices(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceRegistry_FindServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceRegistryServer).FindServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.ServiceRegistry/FindServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceRegistryServer).FindServices(ctx, req.(*FindServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceRegistry_GetServiceAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceRegistryServer).GetServiceAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.ServiceRegistry/GetServiceAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceRegistryServer).GetServiceAddress(ctx, req.(*GetServiceAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceRegistry_GetHttpServerAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHttpServerAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceRegistryServer).GetHttpServerAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.ServiceRegistry/GetHttpServerAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceRegistryServer).GetHttpServerAddress(ctx, req.(*GetHttpServerAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceRegistry_GetAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnnotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceRegistryServer).GetAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.ServiceRegistry/GetAnnotation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceRegistryServer).GetAnnotation(ctx, req.(*GetAnnotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.ServiceRegistry",
	HandlerType: (*ServiceRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAllServices",
			Handler:    _ServiceRegistry_ListAllServices_Handler,
		},
		{
			MethodName: "FindServices",
			Handler:    _ServiceRegistry_FindServices_Handler,
		},
		{
			MethodName: "GetServiceAddress",
			Handler:    _ServiceRegistry_GetServiceAddress_Handler,
		},
		{
			MethodName: "GetHttpServerAddress",
			Handler:    _ServiceRegistry_GetHttpServerAddress_Handler,
		},
		{
			MethodName: "GetAnnotation",
			Handler:    _ServiceRegistry_GetAnnotation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/protos/service_registry.proto",
}
