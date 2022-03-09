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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.10.0
// source: orc8r/cloud/go/services/metricsd/protos/exporter.proto

package protos

import (
	context "context"
	_go "github.com/prometheus/client_model/go"
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

type SubmitMetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*ContextualizedMetric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *SubmitMetricsRequest) Reset() {
	*x = SubmitMetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitMetricsRequest) ProtoMessage() {}

func (x *SubmitMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitMetricsRequest.ProtoReflect.Descriptor instead.
func (*SubmitMetricsRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitMetricsRequest) GetMetrics() []*ContextualizedMetric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type SubmitMetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubmitMetricsResponse) Reset() {
	*x = SubmitMetricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitMetricsResponse) ProtoMessage() {}

func (x *SubmitMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitMetricsResponse.ProtoReflect.Descriptor instead.
func (*SubmitMetricsResponse) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDescGZIP(), []int{1}
}

type ContextualizedMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Family  *_go.MetricFamily `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
	Context *Context          `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *ContextualizedMetric) Reset() {
	*x = ContextualizedMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextualizedMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextualizedMetric) ProtoMessage() {}

func (x *ContextualizedMetric) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextualizedMetric.ProtoReflect.Descriptor instead.
func (*ContextualizedMetric) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDescGZIP(), []int{2}
}

func (x *ContextualizedMetric) GetFamily() *_go.MetricFamily {
	if x != nil {
		return x.Family
	}
	return nil
}

func (x *ContextualizedMetric) GetContext() *Context {
	if x != nil {
		return x.Context
	}
	return nil
}

// Context provides information to the exporter about a metric's origin.
type Context struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName,proto3" json:"metric_name,omitempty"`
	// Types that are assignable to OriginContext:
	//	*Context_CloudMetric
	//	*Context_GatewayMetric
	//	*Context_PushedMetric
	OriginContext isContext_OriginContext `protobuf_oneof:"OriginContext"`
}

func (x *Context) Reset() {
	*x = Context{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDescGZIP(), []int{3}
}

func (x *Context) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

func (m *Context) GetOriginContext() isContext_OriginContext {
	if m != nil {
		return m.OriginContext
	}
	return nil
}

func (x *Context) GetCloudMetric() *CloudContext {
	if x, ok := x.GetOriginContext().(*Context_CloudMetric); ok {
		return x.CloudMetric
	}
	return nil
}

func (x *Context) GetGatewayMetric() *GatewayContext {
	if x, ok := x.GetOriginContext().(*Context_GatewayMetric); ok {
		return x.GatewayMetric
	}
	return nil
}

func (x *Context) GetPushedMetric() *PushedContext {
	if x, ok := x.GetOriginContext().(*Context_PushedMetric); ok {
		return x.PushedMetric
	}
	return nil
}

type isContext_OriginContext interface {
	isContext_OriginContext()
}

type Context_CloudMetric struct {
	CloudMetric *CloudContext `protobuf:"bytes,2,opt,name=cloud_metric,json=cloudMetric,proto3,oneof"`
}

type Context_GatewayMetric struct {
	GatewayMetric *GatewayContext `protobuf:"bytes,3,opt,name=gateway_metric,json=gatewayMetric,proto3,oneof"`
}

type Context_PushedMetric struct {
	PushedMetric *PushedContext `protobuf:"bytes,4,opt,name=pushed_metric,json=pushedMetric,proto3,oneof"`
}

func (*Context_CloudMetric) isContext_OriginContext() {}

func (*Context_GatewayMetric) isContext_OriginContext() {}

func (*Context_PushedMetric) isContext_OriginContext() {}

// CloudContext contains context for metrics scraped from cloud services.
type CloudContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudHost string `protobuf:"bytes,1,opt,name=cloud_host,json=cloudHost,proto3" json:"cloud_host,omitempty"`
}

func (x *CloudContext) Reset() {
	*x = CloudContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudContext) ProtoMessage() {}

func (x *CloudContext) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudContext.ProtoReflect.Descriptor instead.
func (*CloudContext) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDescGZIP(), []int{4}
}

func (x *CloudContext) GetCloudHost() string {
	if x != nil {
		return x.CloudHost
	}
	return ""
}

// GatewayContext contains context for metrics submitted from gateways.
type GatewayContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	GatewayId string `protobuf:"bytes,2,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
}

func (x *GatewayContext) Reset() {
	*x = GatewayContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayContext) ProtoMessage() {}

func (x *GatewayContext) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayContext.ProtoReflect.Descriptor instead.
func (*GatewayContext) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDescGZIP(), []int{5}
}

func (x *GatewayContext) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *GatewayContext) GetGatewayId() string {
	if x != nil {
		return x.GatewayId
	}
	return ""
}

// PushedContext contains context for metrics pushed via the REST API.
type PushedContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
}

func (x *PushedContext) Reset() {
	*x = PushedContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushedContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushedContext) ProtoMessage() {}

func (x *PushedContext) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushedContext.ProtoReflect.Descriptor instead.
func (*PushedContext) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDescGZIP(), []int{6}
}

func (x *PushedContext) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

var File_orc8r_cloud_go_services_metricsd_protos_exporter_proto protoreflect.FileDescriptor

var file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDesc = []byte{
	0x0a, 0x36, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e,
	0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x64, 0x1a, 0x0d,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a,
	0x14, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f,
	0x72, 0x63, 0x38, 0x72, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x64, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x75, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x3a, 0x0a,
	0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x46, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x37, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x22, 0x9f, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x47, 0x0a, 0x0c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72,
	0x63, 0x38, 0x72, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x64, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x4d, 0x0a, 0x0e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x64, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x4a, 0x0a, 0x0d, 0x70, 0x75, 0x73, 0x68, 0x65,
	0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x64, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x75, 0x73, 0x68, 0x65, 0x64, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x42, 0x0f, 0x0a, 0x0d, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x2d, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x48,
	0x6f, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x0e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x0d, 0x50, 0x75, 0x73, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x49, 0x64, 0x32, 0x76, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x63, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x12, 0x2a, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x64, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x64, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDescOnce sync.Once
	file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDescData = file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDesc
)

func file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDescGZIP() []byte {
	file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDescOnce.Do(func() {
		file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDescData = protoimpl.X.CompressGZIP(file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDescData)
	})
	return file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDescData
}

var file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_goTypes = []interface{}{
	(*SubmitMetricsRequest)(nil),  // 0: magma.orc8r.metricsd.SubmitMetricsRequest
	(*SubmitMetricsResponse)(nil), // 1: magma.orc8r.metricsd.SubmitMetricsResponse
	(*ContextualizedMetric)(nil),  // 2: magma.orc8r.metricsd.ContextualizedMetric
	(*Context)(nil),               // 3: magma.orc8r.metricsd.Context
	(*CloudContext)(nil),          // 4: magma.orc8r.metricsd.CloudContext
	(*GatewayContext)(nil),        // 5: magma.orc8r.metricsd.GatewayContext
	(*PushedContext)(nil),         // 6: magma.orc8r.metricsd.PushedContext
	(*_go.MetricFamily)(nil),      // 7: io.prometheus.client.MetricFamily
}
var file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_depIdxs = []int32{
	2, // 0: magma.orc8r.metricsd.SubmitMetricsRequest.metrics:type_name -> magma.orc8r.metricsd.ContextualizedMetric
	7, // 1: magma.orc8r.metricsd.ContextualizedMetric.family:type_name -> io.prometheus.client.MetricFamily
	3, // 2: magma.orc8r.metricsd.ContextualizedMetric.context:type_name -> magma.orc8r.metricsd.Context
	4, // 3: magma.orc8r.metricsd.Context.cloud_metric:type_name -> magma.orc8r.metricsd.CloudContext
	5, // 4: magma.orc8r.metricsd.Context.gateway_metric:type_name -> magma.orc8r.metricsd.GatewayContext
	6, // 5: magma.orc8r.metricsd.Context.pushed_metric:type_name -> magma.orc8r.metricsd.PushedContext
	0, // 6: magma.orc8r.metricsd.MetricsExporter.Submit:input_type -> magma.orc8r.metricsd.SubmitMetricsRequest
	1, // 7: magma.orc8r.metricsd.MetricsExporter.Submit:output_type -> magma.orc8r.metricsd.SubmitMetricsResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_init() }
func file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_init() {
	if File_orc8r_cloud_go_services_metricsd_protos_exporter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitMetricsRequest); i {
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
		file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitMetricsResponse); i {
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
		file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContextualizedMetric); i {
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
		file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Context); i {
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
		file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudContext); i {
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
		file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayContext); i {
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
		file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushedContext); i {
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
	file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Context_CloudMetric)(nil),
		(*Context_GatewayMetric)(nil),
		(*Context_PushedMetric)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_goTypes,
		DependencyIndexes: file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_depIdxs,
		MessageInfos:      file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_msgTypes,
	}.Build()
	File_orc8r_cloud_go_services_metricsd_protos_exporter_proto = out.File
	file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_rawDesc = nil
	file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_goTypes = nil
	file_orc8r_cloud_go_services_metricsd_protos_exporter_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MetricsExporterClient is the client API for MetricsExporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsExporterClient interface {
	// Submit metrics to datasinks.
	Submit(ctx context.Context, in *SubmitMetricsRequest, opts ...grpc.CallOption) (*SubmitMetricsResponse, error)
}

type metricsExporterClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsExporterClient(cc grpc.ClientConnInterface) MetricsExporterClient {
	return &metricsExporterClient{cc}
}

func (c *metricsExporterClient) Submit(ctx context.Context, in *SubmitMetricsRequest, opts ...grpc.CallOption) (*SubmitMetricsResponse, error) {
	out := new(SubmitMetricsResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.metricsd.MetricsExporter/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsExporterServer is the server API for MetricsExporter service.
type MetricsExporterServer interface {
	// Submit metrics to datasinks.
	Submit(context.Context, *SubmitMetricsRequest) (*SubmitMetricsResponse, error)
}

// UnimplementedMetricsExporterServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsExporterServer struct {
}

func (*UnimplementedMetricsExporterServer) Submit(context.Context, *SubmitMetricsRequest) (*SubmitMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}

func RegisterMetricsExporterServer(s *grpc.Server, srv MetricsExporterServer) {
	s.RegisterService(&_MetricsExporter_serviceDesc, srv)
}

func _MetricsExporter_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsExporterServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.metricsd.MetricsExporter/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsExporterServer).Submit(ctx, req.(*SubmitMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricsExporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.metricsd.MetricsExporter",
	HandlerType: (*MetricsExporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _MetricsExporter_Submit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/cloud/go/services/metricsd/protos/exporter.proto",
}
