// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.10.0
// source: lte/protos/oai/ngap_state.proto

package oai

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

type NgapTimer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // long
	// Deprecated: Do not use.
	Sec  int32 `protobuf:"varint,2,opt,name=sec,proto3" json:"sec,omitempty"`   // long
	Msec int32 `protobuf:"varint,3,opt,name=msec,proto3" json:"msec,omitempty"` // long
}

func (x *NgapTimer) Reset() {
	*x = NgapTimer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_ngap_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NgapTimer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NgapTimer) ProtoMessage() {}

func (x *NgapTimer) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_ngap_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NgapTimer.ProtoReflect.Descriptor instead.
func (*NgapTimer) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_ngap_state_proto_rawDescGZIP(), []int{0}
}

func (x *NgapTimer) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Deprecated: Do not use.
func (x *NgapTimer) GetSec() int32 {
	if x != nil {
		return x.Sec
	}
	return 0
}

func (x *NgapTimer) GetMsec() int32 {
	if x != nil {
		return x.Msec
	}
	return 0
}

type Ngap_SupportedTaiItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tac            uint32   `protobuf:"varint,1,opt,name=tac,proto3" json:"tac,omitempty"`
	BplmnlistCount uint32   `protobuf:"varint,2,opt,name=bplmnlist_count,json=bplmnlistCount,proto3" json:"bplmnlist_count,omitempty"`
	Bplmns         [][]byte `protobuf:"bytes,3,rep,name=bplmns,proto3" json:"bplmns,omitempty"`
}

func (x *Ngap_SupportedTaiItems) Reset() {
	*x = Ngap_SupportedTaiItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_ngap_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ngap_SupportedTaiItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ngap_SupportedTaiItems) ProtoMessage() {}

func (x *Ngap_SupportedTaiItems) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_ngap_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ngap_SupportedTaiItems.ProtoReflect.Descriptor instead.
func (*Ngap_SupportedTaiItems) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_ngap_state_proto_rawDescGZIP(), []int{1}
}

func (x *Ngap_SupportedTaiItems) GetTac() uint32 {
	if x != nil {
		return x.Tac
	}
	return 0
}

func (x *Ngap_SupportedTaiItems) GetBplmnlistCount() uint32 {
	if x != nil {
		return x.BplmnlistCount
	}
	return 0
}

func (x *Ngap_SupportedTaiItems) GetBplmns() [][]byte {
	if x != nil {
		return x.Bplmns
	}
	return nil
}

// supported_ta_list_t
type Ngap_SupportedTaList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListCount         uint32                    `protobuf:"varint,1,opt,name=list_count,json=listCount,proto3" json:"list_count,omitempty"`
	SupportedTaiItems []*Ngap_SupportedTaiItems `protobuf:"bytes,2,rep,name=supported_tai_items,json=supportedTaiItems,proto3" json:"supported_tai_items,omitempty"`
}

func (x *Ngap_SupportedTaList) Reset() {
	*x = Ngap_SupportedTaList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_ngap_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ngap_SupportedTaList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ngap_SupportedTaList) ProtoMessage() {}

func (x *Ngap_SupportedTaList) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_ngap_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ngap_SupportedTaList.ProtoReflect.Descriptor instead.
func (*Ngap_SupportedTaList) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_ngap_state_proto_rawDescGZIP(), []int{2}
}

func (x *Ngap_SupportedTaList) GetListCount() uint32 {
	if x != nil {
		return x.ListCount
	}
	return 0
}

func (x *Ngap_SupportedTaList) GetSupportedTaiItems() []*Ngap_SupportedTaiItems {
	if x != nil {
		return x.SupportedTaiItems
	}
	return nil
}

type GnbDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GnbId            uint32                `protobuf:"varint,1,opt,name=gnb_id,json=gnbId,proto3" json:"gnb_id,omitempty"`
	NgState          int32                 `protobuf:"varint,2,opt,name=ng_state,json=ngState,proto3" json:"ng_state,omitempty"`                              // enum amf_ng_gnb_state_s
	GnbName          []byte                `protobuf:"bytes,3,opt,name=gnb_name,json=gnbName,proto3" json:"gnb_name,omitempty"`                               // char[]
	DefaultPagingDrx uint32                `protobuf:"varint,4,opt,name=default_paging_drx,json=defaultPagingDrx,proto3" json:"default_paging_drx,omitempty"` // uint8_t
	NbUeAssociated   uint32                `protobuf:"varint,5,opt,name=nb_ue_associated,json=nbUeAssociated,proto3" json:"nb_ue_associated,omitempty"`
	SctpAssocId      uint32                `protobuf:"varint,6,opt,name=sctp_assoc_id,json=sctpAssocId,proto3" json:"sctp_assoc_id,omitempty"`                                                                       // sctp_assoc_id_t
	NextSctpStream   uint32                `protobuf:"varint,7,opt,name=next_sctp_stream,json=nextSctpStream,proto3" json:"next_sctp_stream,omitempty"`                                                              // sctp_stream_id_t
	Instreams        uint32                `protobuf:"varint,8,opt,name=instreams,proto3" json:"instreams,omitempty"`                                                                                                // sctp_stream_id_t
	Outstreams       uint32                `protobuf:"varint,9,opt,name=outstreams,proto3" json:"outstreams,omitempty"`                                                                                              // sctp_stream_id_t
	UeIds            map[uint64]uint64     `protobuf:"bytes,10,rep,name=ue_ids,json=ueIds,proto3" json:"ue_ids,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // gnb_ue_ngap_id -> comp_ngap_id
	SupportedTaList  *Ngap_SupportedTaList `protobuf:"bytes,11,opt,name=supported_ta_list,json=supportedTaList,proto3" json:"supported_ta_list,omitempty"`                                                           // TAs supported by gNB
}

func (x *GnbDescription) Reset() {
	*x = GnbDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_ngap_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GnbDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GnbDescription) ProtoMessage() {}

func (x *GnbDescription) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_ngap_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GnbDescription.ProtoReflect.Descriptor instead.
func (*GnbDescription) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_ngap_state_proto_rawDescGZIP(), []int{3}
}

func (x *GnbDescription) GetGnbId() uint32 {
	if x != nil {
		return x.GnbId
	}
	return 0
}

func (x *GnbDescription) GetNgState() int32 {
	if x != nil {
		return x.NgState
	}
	return 0
}

func (x *GnbDescription) GetGnbName() []byte {
	if x != nil {
		return x.GnbName
	}
	return nil
}

func (x *GnbDescription) GetDefaultPagingDrx() uint32 {
	if x != nil {
		return x.DefaultPagingDrx
	}
	return 0
}

func (x *GnbDescription) GetNbUeAssociated() uint32 {
	if x != nil {
		return x.NbUeAssociated
	}
	return 0
}

func (x *GnbDescription) GetSctpAssocId() uint32 {
	if x != nil {
		return x.SctpAssocId
	}
	return 0
}

func (x *GnbDescription) GetNextSctpStream() uint32 {
	if x != nil {
		return x.NextSctpStream
	}
	return 0
}

func (x *GnbDescription) GetInstreams() uint32 {
	if x != nil {
		return x.Instreams
	}
	return 0
}

func (x *GnbDescription) GetOutstreams() uint32 {
	if x != nil {
		return x.Outstreams
	}
	return 0
}

func (x *GnbDescription) GetUeIds() map[uint64]uint64 {
	if x != nil {
		return x.UeIds
	}
	return nil
}

func (x *GnbDescription) GetSupportedTaList() *Ngap_SupportedTaList {
	if x != nil {
		return x.SupportedTaList
	}
	return nil
}

type Ngap_UeDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NgUeState             int32      `protobuf:"varint,2,opt,name=ng_ue_state,json=ngUeState,proto3" json:"ng_ue_state,omitempty"`                                        // enum ng_ue_state_s
	GnbUeNgapId           uint32     `protobuf:"varint,3,opt,name=gnb_ue_ngap_id,json=gnbUeNgapId,proto3" json:"gnb_ue_ngap_id,omitempty"`                                // gnb_ue_ngap_id_t
	AmfUeNgapId           uint32     `protobuf:"varint,4,opt,name=amf_ue_ngap_id,json=amfUeNgapId,proto3" json:"amf_ue_ngap_id,omitempty"`                                // amf_ue_ngap_id_t
	SctpStreamRecv        uint32     `protobuf:"varint,5,opt,name=sctp_stream_recv,json=sctpStreamRecv,proto3" json:"sctp_stream_recv,omitempty"`                         // sctp_stream_id_t
	SctpStreamSend        uint32     `protobuf:"varint,6,opt,name=sctp_stream_send,json=sctpStreamSend,proto3" json:"sctp_stream_send,omitempty"`                         // sctp_stream_id_t
	NgapUeContextRelTimer *NgapTimer `protobuf:"bytes,7,opt,name=ngap_ue_context_rel_timer,json=ngapUeContextRelTimer,proto3" json:"ngap_ue_context_rel_timer,omitempty"` // struct ngap_timer_t
	SctpAssocId           uint32     `protobuf:"varint,8,opt,name=sctp_assoc_id,json=sctpAssocId,proto3" json:"sctp_assoc_id,omitempty"`                                  // sctp_assoc_id_t
}

func (x *Ngap_UeDescription) Reset() {
	*x = Ngap_UeDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_ngap_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ngap_UeDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ngap_UeDescription) ProtoMessage() {}

func (x *Ngap_UeDescription) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_ngap_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ngap_UeDescription.ProtoReflect.Descriptor instead.
func (*Ngap_UeDescription) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_ngap_state_proto_rawDescGZIP(), []int{4}
}

func (x *Ngap_UeDescription) GetNgUeState() int32 {
	if x != nil {
		return x.NgUeState
	}
	return 0
}

func (x *Ngap_UeDescription) GetGnbUeNgapId() uint32 {
	if x != nil {
		return x.GnbUeNgapId
	}
	return 0
}

func (x *Ngap_UeDescription) GetAmfUeNgapId() uint32 {
	if x != nil {
		return x.AmfUeNgapId
	}
	return 0
}

func (x *Ngap_UeDescription) GetSctpStreamRecv() uint32 {
	if x != nil {
		return x.SctpStreamRecv
	}
	return 0
}

func (x *Ngap_UeDescription) GetSctpStreamSend() uint32 {
	if x != nil {
		return x.SctpStreamSend
	}
	return 0
}

func (x *Ngap_UeDescription) GetNgapUeContextRelTimer() *NgapTimer {
	if x != nil {
		return x.NgapUeContextRelTimer
	}
	return nil
}

func (x *Ngap_UeDescription) GetSctpAssocId() uint32 {
	if x != nil {
		return x.SctpAssocId
	}
	return 0
}

type NgapState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gnbs          map[uint32]*GnbDescription `protobuf:"bytes,1,rep,name=gnbs,proto3" json:"gnbs,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`                    // gnbid -> GnbDescription
	Amfid2Associd map[uint32]uint32          `protobuf:"bytes,2,rep,name=amfid2associd,proto3" json:"amfid2associd,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // amfueid -> ue associd
	NumGnbs       uint32                     `protobuf:"varint,3,opt,name=num_gnbs,json=numGnbs,proto3" json:"num_gnbs,omitempty"`
}

func (x *NgapState) Reset() {
	*x = NgapState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_ngap_state_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NgapState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NgapState) ProtoMessage() {}

func (x *NgapState) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_ngap_state_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NgapState.ProtoReflect.Descriptor instead.
func (*NgapState) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_ngap_state_proto_rawDescGZIP(), []int{5}
}

func (x *NgapState) GetGnbs() map[uint32]*GnbDescription {
	if x != nil {
		return x.Gnbs
	}
	return nil
}

func (x *NgapState) GetAmfid2Associd() map[uint32]uint32 {
	if x != nil {
		return x.Amfid2Associd
	}
	return nil
}

func (x *NgapState) GetNumGnbs() uint32 {
	if x != nil {
		return x.NumGnbs
	}
	return 0
}

type NgapImsiMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AmfUeIdImsiMap map[uint64]uint64 `protobuf:"bytes,1,rep,name=amf_ue_id_imsi_map,json=amfUeIdImsiMap,proto3" json:"amf_ue_id_imsi_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // amf_ngap_ue_id => IMSI64
}

func (x *NgapImsiMap) Reset() {
	*x = NgapImsiMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_ngap_state_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NgapImsiMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NgapImsiMap) ProtoMessage() {}

func (x *NgapImsiMap) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_ngap_state_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NgapImsiMap.ProtoReflect.Descriptor instead.
func (*NgapImsiMap) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_ngap_state_proto_rawDescGZIP(), []int{6}
}

func (x *NgapImsiMap) GetAmfUeIdImsiMap() map[uint64]uint64 {
	if x != nil {
		return x.AmfUeIdImsiMap
	}
	return nil
}

var File_lte_protos_oai_ngap_state_proto protoreflect.FileDescriptor

var file_lte_protos_oai_ngap_state_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6c, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x61, 0x69,
	0x2f, 0x6e, 0x67, 0x61, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69,
	0x22, 0x45, 0x0a, 0x09, 0x4e, 0x67, 0x61, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x03, 0x73, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x02, 0x18, 0x01, 0x52, 0x03,
	0x73, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x73, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x6d, 0x73, 0x65, 0x63, 0x22, 0x6b, 0x0a, 0x16, 0x4e, 0x67, 0x61, 0x70, 0x5f,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x61, 0x69, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x74, 0x61, 0x63, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x70, 0x6c, 0x6d, 0x6e, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x62, 0x70,
	0x6c, 0x6d, 0x6e, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x70, 0x6c, 0x6d, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x70,
	0x6c, 0x6d, 0x6e, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x14, 0x4e, 0x67, 0x61, 0x70, 0x5f, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x55, 0x0a, 0x13,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x69, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x4e, 0x67, 0x61, 0x70, 0x5f, 0x53,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x61, 0x69, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x11, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x61, 0x69, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x8d, 0x04, 0x0a, 0x0e, 0x47, 0x6e, 0x62, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x67, 0x6e, 0x62, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x67, 0x6e, 0x62, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x6e, 0x62, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x67, 0x6e, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x72, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x10, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x44, 0x72,
	0x78, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x62, 0x5f, 0x75, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x6f, 0x63,
	0x69, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6e, 0x62, 0x55,
	0x65, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73,
	0x63, 0x74, 0x70, 0x5f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x73, 0x63, 0x74, 0x70, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x49, 0x64, 0x12,
	0x28, 0x0a, 0x10, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x63, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6e, 0x65, 0x78, 0x74, 0x53,
	0x63, 0x74, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x69, 0x6e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6f, 0x75, 0x74,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x3f, 0x0a, 0x06, 0x75, 0x65, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e,
	0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x47, 0x6e, 0x62, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x65, 0x49, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x75, 0x65, 0x49, 0x64, 0x73, 0x12, 0x4f, 0x0a, 0x11, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e,
	0x6f, 0x61, 0x69, 0x2e, 0x4e, 0x67, 0x61, 0x70, 0x5f, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x54, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x54, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x38, 0x0a, 0x0a, 0x55, 0x65, 0x49,
	0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xca, 0x02, 0x0a, 0x12, 0x4e, 0x67, 0x61, 0x70, 0x5f, 0x55, 0x65, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0b, 0x6e, 0x67,
	0x5f, 0x75, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x6e, 0x67, 0x55, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0e, 0x67, 0x6e,
	0x62, 0x5f, 0x75, 0x65, 0x5f, 0x6e, 0x67, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x67, 0x6e, 0x62, 0x55, 0x65, 0x4e, 0x67, 0x61, 0x70, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0e, 0x61, 0x6d, 0x66, 0x5f, 0x75, 0x65, 0x5f, 0x6e, 0x67, 0x61, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x6d, 0x66, 0x55, 0x65, 0x4e, 0x67,
	0x61, 0x70, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x63, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x63, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e,
	0x73, 0x63, 0x74, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x63, 0x76, 0x12, 0x28,
	0x0a, 0x10, 0x73, 0x63, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x73, 0x65,
	0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x63, 0x74, 0x70, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x52, 0x0a, 0x19, 0x6e, 0x67, 0x61, 0x70,
	0x5f, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x61,
	0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x4e, 0x67, 0x61, 0x70,
	0x54, 0x69, 0x6d, 0x65, 0x72, 0x52, 0x15, 0x6e, 0x67, 0x61, 0x70, 0x55, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d,
	0x73, 0x63, 0x74, 0x70, 0x5f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x63, 0x74, 0x70, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x49, 0x64,
	0x22, 0xcb, 0x02, 0x0a, 0x09, 0x4e, 0x67, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x36,
	0x0a, 0x04, 0x67, 0x6e, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x4e, 0x67, 0x61,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x6e, 0x62, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x67, 0x6e, 0x62, 0x73, 0x12, 0x51, 0x0a, 0x0d, 0x61, 0x6d, 0x66, 0x69, 0x64, 0x32,
	0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x4e, 0x67,
	0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x6d, 0x66, 0x69, 0x64, 0x32, 0x61, 0x73,
	0x73, 0x6f, 0x63, 0x69, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x61, 0x6d, 0x66, 0x69,
	0x64, 0x32, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x75, 0x6d,
	0x5f, 0x67, 0x6e, 0x62, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x75, 0x6d,
	0x47, 0x6e, 0x62, 0x73, 0x1a, 0x56, 0x0a, 0x09, 0x47, 0x6e, 0x62, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f,
	0x61, 0x69, 0x2e, 0x47, 0x6e, 0x62, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x40, 0x0a, 0x12,
	0x41, 0x6d, 0x66, 0x69, 0x64, 0x32, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xac,
	0x01, 0x0a, 0x0b, 0x4e, 0x67, 0x61, 0x70, 0x49, 0x6d, 0x73, 0x69, 0x4d, 0x61, 0x70, 0x12, 0x5a,
	0x0a, 0x12, 0x61, 0x6d, 0x66, 0x5f, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x69, 0x6d, 0x73, 0x69,
	0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x4e, 0x67, 0x61, 0x70, 0x49,
	0x6d, 0x73, 0x69, 0x4d, 0x61, 0x70, 0x2e, 0x41, 0x6d, 0x66, 0x55, 0x65, 0x49, 0x64, 0x49, 0x6d,
	0x73, 0x69, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x61, 0x6d, 0x66, 0x55,
	0x65, 0x49, 0x64, 0x49, 0x6d, 0x73, 0x69, 0x4d, 0x61, 0x70, 0x1a, 0x41, 0x0a, 0x13, 0x41, 0x6d,
	0x66, 0x55, 0x65, 0x49, 0x64, 0x49, 0x6d, 0x73, 0x69, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x1f, 0x5a,
	0x1d, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6c, 0x74, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x61, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lte_protos_oai_ngap_state_proto_rawDescOnce sync.Once
	file_lte_protos_oai_ngap_state_proto_rawDescData = file_lte_protos_oai_ngap_state_proto_rawDesc
)

func file_lte_protos_oai_ngap_state_proto_rawDescGZIP() []byte {
	file_lte_protos_oai_ngap_state_proto_rawDescOnce.Do(func() {
		file_lte_protos_oai_ngap_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_lte_protos_oai_ngap_state_proto_rawDescData)
	})
	return file_lte_protos_oai_ngap_state_proto_rawDescData
}

var file_lte_protos_oai_ngap_state_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_lte_protos_oai_ngap_state_proto_goTypes = []interface{}{
	(*NgapTimer)(nil),              // 0: magma.lte.oai.NgapTimer
	(*Ngap_SupportedTaiItems)(nil), // 1: magma.lte.oai.Ngap_SupportedTaiItems
	(*Ngap_SupportedTaList)(nil),   // 2: magma.lte.oai.Ngap_SupportedTaList
	(*GnbDescription)(nil),         // 3: magma.lte.oai.GnbDescription
	(*Ngap_UeDescription)(nil),     // 4: magma.lte.oai.Ngap_UeDescription
	(*NgapState)(nil),              // 5: magma.lte.oai.NgapState
	(*NgapImsiMap)(nil),            // 6: magma.lte.oai.NgapImsiMap
	nil,                            // 7: magma.lte.oai.GnbDescription.UeIdsEntry
	nil,                            // 8: magma.lte.oai.NgapState.GnbsEntry
	nil,                            // 9: magma.lte.oai.NgapState.Amfid2associdEntry
	nil,                            // 10: magma.lte.oai.NgapImsiMap.AmfUeIdImsiMapEntry
}
var file_lte_protos_oai_ngap_state_proto_depIdxs = []int32{
	1,  // 0: magma.lte.oai.Ngap_SupportedTaList.supported_tai_items:type_name -> magma.lte.oai.Ngap_SupportedTaiItems
	7,  // 1: magma.lte.oai.GnbDescription.ue_ids:type_name -> magma.lte.oai.GnbDescription.UeIdsEntry
	2,  // 2: magma.lte.oai.GnbDescription.supported_ta_list:type_name -> magma.lte.oai.Ngap_SupportedTaList
	0,  // 3: magma.lte.oai.Ngap_UeDescription.ngap_ue_context_rel_timer:type_name -> magma.lte.oai.NgapTimer
	8,  // 4: magma.lte.oai.NgapState.gnbs:type_name -> magma.lte.oai.NgapState.GnbsEntry
	9,  // 5: magma.lte.oai.NgapState.amfid2associd:type_name -> magma.lte.oai.NgapState.Amfid2associdEntry
	10, // 6: magma.lte.oai.NgapImsiMap.amf_ue_id_imsi_map:type_name -> magma.lte.oai.NgapImsiMap.AmfUeIdImsiMapEntry
	3,  // 7: magma.lte.oai.NgapState.GnbsEntry.value:type_name -> magma.lte.oai.GnbDescription
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_lte_protos_oai_ngap_state_proto_init() }
func file_lte_protos_oai_ngap_state_proto_init() {
	if File_lte_protos_oai_ngap_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lte_protos_oai_ngap_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NgapTimer); i {
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
		file_lte_protos_oai_ngap_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ngap_SupportedTaiItems); i {
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
		file_lte_protos_oai_ngap_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ngap_SupportedTaList); i {
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
		file_lte_protos_oai_ngap_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GnbDescription); i {
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
		file_lte_protos_oai_ngap_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ngap_UeDescription); i {
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
		file_lte_protos_oai_ngap_state_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NgapState); i {
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
		file_lte_protos_oai_ngap_state_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NgapImsiMap); i {
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
			RawDescriptor: file_lte_protos_oai_ngap_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lte_protos_oai_ngap_state_proto_goTypes,
		DependencyIndexes: file_lte_protos_oai_ngap_state_proto_depIdxs,
		MessageInfos:      file_lte_protos_oai_ngap_state_proto_msgTypes,
	}.Build()
	File_lte_protos_oai_ngap_state_proto = out.File
	file_lte_protos_oai_ngap_state_proto_rawDesc = nil
	file_lte_protos_oai_ngap_state_proto_goTypes = nil
	file_lte_protos_oai_ngap_state_proto_depIdxs = nil
}
