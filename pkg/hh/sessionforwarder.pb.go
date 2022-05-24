// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pkg/hh/sessionforwarder.proto

package hh

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

type Bug int32

const (
	Bug_open   Bug = 0
	Bug_closed Bug = 1
)

// Enum value maps for Bug.
var (
	Bug_name = map[int32]string{
		0: "open",
		1: "closed",
	}
	Bug_value = map[string]int32{
		"open":   0,
		"closed": 1,
	}
)

func (x Bug) Enum() *Bug {
	p := new(Bug)
	*p = x
	return p
}

func (x Bug) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Bug) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_hh_sessionforwarder_proto_enumTypes[0].Descriptor()
}

func (Bug) Type() protoreflect.EnumType {
	return &file_pkg_hh_sessionforwarder_proto_enumTypes[0]
}

func (x Bug) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Bug.Descriptor instead.
func (Bug) EnumDescriptor() ([]byte, []int) {
	return file_pkg_hh_sessionforwarder_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_hh_sessionforwarder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_hh_sessionforwarder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_pkg_hh_sessionforwarder_proto_rawDescGZIP(), []int{0}
}

type ReportSessionsStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionStats []*SessionStats `protobuf:"bytes,1,rep,name=session_stats,json=sessionStats,proto3" json:"session_stats,omitempty"`
}

func (x *ReportSessionsStatsRequest) Reset() {
	*x = ReportSessionsStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_hh_sessionforwarder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportSessionsStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportSessionsStatsRequest) ProtoMessage() {}

func (x *ReportSessionsStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_hh_sessionforwarder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportSessionsStatsRequest.ProtoReflect.Descriptor instead.
func (*ReportSessionsStatsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_hh_sessionforwarder_proto_rawDescGZIP(), []int{1}
}

func (x *ReportSessionsStatsRequest) GetSessionStats() []*SessionStats {
	if x != nil {
		return x.SessionStats
	}
	return nil
}

type SessionStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid                   string               `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Usage                 []*SessionStatsUsage `protobuf:"bytes,2,rep,name=usage,proto3" json:"usage,omitempty"`
	Imsi                  string               `protobuf:"bytes,3,opt,name=imsi,proto3" json:"imsi,omitempty"`
	Imei                  string               `protobuf:"bytes,4,opt,name=imei,proto3" json:"imei,omitempty"`
	Msisdn                string               `protobuf:"bytes,5,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	PlmnId                string               `protobuf:"bytes,6,opt,name=plmn_id,json=plmnId,proto3" json:"plmn_id,omitempty"`
	PgwPlmnId             string               `protobuf:"bytes,7,opt,name=pgw_plmn_id,json=pgwPlmnId,proto3" json:"pgw_plmn_id,omitempty"`
	Apnni                 string               `protobuf:"bytes,8,opt,name=apnni,proto3" json:"apnni,omitempty"`
	PdpType               string               `protobuf:"bytes,9,opt,name=pdp_type,json=pdpType,proto3" json:"pdp_type,omitempty"`
	UeIp                  string               `protobuf:"bytes,10,opt,name=ue_ip,json=ueIp,proto3" json:"ue_ip,omitempty"`
	SgwIp                 string               `protobuf:"bytes,11,opt,name=sgw_ip,json=sgwIp,proto3" json:"sgw_ip,omitempty"`
	PgwIp                 string               `protobuf:"bytes,12,opt,name=pgw_ip,json=pgwIp,proto3" json:"pgw_ip,omitempty"`
	UeIpv6                string               `protobuf:"bytes,13,opt,name=ue_ipv6,json=ueIpv6,proto3" json:"ue_ipv6,omitempty"`
	SgwIpv6               string               `protobuf:"bytes,14,opt,name=sgw_ipv6,json=sgwIpv6,proto3" json:"sgw_ipv6,omitempty"`
	PgwIpv6               string               `protobuf:"bytes,15,opt,name=pgw_ipv6,json=pgwIpv6,proto3" json:"pgw_ipv6,omitempty"`
	RatType               uint32               `protobuf:"varint,16,opt,name=rat_type,json=ratType,proto3" json:"rat_type,omitempty"`
	SessionStartTime      int64                `protobuf:"varint,17,opt,name=session_start_time,json=sessionStartTime,proto3" json:"session_start_time,omitempty"`
	FinalRecord           bool                 `protobuf:"varint,18,opt,name=final_record,json=finalRecord,proto3" json:"final_record,omitempty"`
	LocalSequenceNumber   uint64               `protobuf:"varint,19,opt,name=local_sequence_number,json=localSequenceNumber,proto3" json:"local_sequence_number,omitempty"`
	RecordOpeningTime     uint64               `protobuf:"varint,20,opt,name=record_opening_time,json=recordOpeningTime,proto3" json:"record_opening_time,omitempty"`
	RecordDuration        int64                `protobuf:"varint,21,opt,name=record_duration,json=recordDuration,proto3" json:"record_duration,omitempty"`
	CauseForRecordClosing uint32               `protobuf:"varint,22,opt,name=cause_for_record_closing,json=causeForRecordClosing,proto3" json:"cause_for_record_closing,omitempty"`
	RecordType            uint32               `protobuf:"varint,23,opt,name=record_type,json=recordType,proto3" json:"record_type,omitempty"`
	ChargingId            uint32               `protobuf:"varint,24,opt,name=charging_id,json=chargingId,proto3" json:"charging_id,omitempty"`
	RawUserLocationInfo   []byte               `protobuf:"bytes,25,opt,name=raw_user_location_info,json=rawUserLocationInfo,proto3" json:"raw_user_location_info,omitempty"`
	UeTimezone            string               `protobuf:"bytes,26,opt,name=ue_timezone,json=ueTimezone,proto3" json:"ue_timezone,omitempty"`
}

func (x *SessionStats) Reset() {
	*x = SessionStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_hh_sessionforwarder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionStats) ProtoMessage() {}

func (x *SessionStats) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_hh_sessionforwarder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionStats.ProtoReflect.Descriptor instead.
func (*SessionStats) Descriptor() ([]byte, []int) {
	return file_pkg_hh_sessionforwarder_proto_rawDescGZIP(), []int{2}
}

func (x *SessionStats) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *SessionStats) GetUsage() []*SessionStatsUsage {
	if x != nil {
		return x.Usage
	}
	return nil
}

func (x *SessionStats) GetImsi() string {
	if x != nil {
		return x.Imsi
	}
	return ""
}

func (x *SessionStats) GetImei() string {
	if x != nil {
		return x.Imei
	}
	return ""
}

func (x *SessionStats) GetMsisdn() string {
	if x != nil {
		return x.Msisdn
	}
	return ""
}

func (x *SessionStats) GetPlmnId() string {
	if x != nil {
		return x.PlmnId
	}
	return ""
}

func (x *SessionStats) GetPgwPlmnId() string {
	if x != nil {
		return x.PgwPlmnId
	}
	return ""
}

func (x *SessionStats) GetApnni() string {
	if x != nil {
		return x.Apnni
	}
	return ""
}

func (x *SessionStats) GetPdpType() string {
	if x != nil {
		return x.PdpType
	}
	return ""
}

func (x *SessionStats) GetUeIp() string {
	if x != nil {
		return x.UeIp
	}
	return ""
}

func (x *SessionStats) GetSgwIp() string {
	if x != nil {
		return x.SgwIp
	}
	return ""
}

func (x *SessionStats) GetPgwIp() string {
	if x != nil {
		return x.PgwIp
	}
	return ""
}

func (x *SessionStats) GetUeIpv6() string {
	if x != nil {
		return x.UeIpv6
	}
	return ""
}

func (x *SessionStats) GetSgwIpv6() string {
	if x != nil {
		return x.SgwIpv6
	}
	return ""
}

func (x *SessionStats) GetPgwIpv6() string {
	if x != nil {
		return x.PgwIpv6
	}
	return ""
}

func (x *SessionStats) GetRatType() uint32 {
	if x != nil {
		return x.RatType
	}
	return 0
}

func (x *SessionStats) GetSessionStartTime() int64 {
	if x != nil {
		return x.SessionStartTime
	}
	return 0
}

func (x *SessionStats) GetFinalRecord() bool {
	if x != nil {
		return x.FinalRecord
	}
	return false
}

func (x *SessionStats) GetLocalSequenceNumber() uint64 {
	if x != nil {
		return x.LocalSequenceNumber
	}
	return 0
}

func (x *SessionStats) GetRecordOpeningTime() uint64 {
	if x != nil {
		return x.RecordOpeningTime
	}
	return 0
}

func (x *SessionStats) GetRecordDuration() int64 {
	if x != nil {
		return x.RecordDuration
	}
	return 0
}

func (x *SessionStats) GetCauseForRecordClosing() uint32 {
	if x != nil {
		return x.CauseForRecordClosing
	}
	return 0
}

func (x *SessionStats) GetRecordType() uint32 {
	if x != nil {
		return x.RecordType
	}
	return 0
}

func (x *SessionStats) GetChargingId() uint32 {
	if x != nil {
		return x.ChargingId
	}
	return 0
}

func (x *SessionStats) GetRawUserLocationInfo() []byte {
	if x != nil {
		return x.RawUserLocationInfo
	}
	return nil
}

func (x *SessionStats) GetUeTimezone() string {
	if x != nil {
		return x.UeTimezone
	}
	return ""
}

type SessionStatsUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuleId    string `protobuf:"bytes,1,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
	BytesTx   uint64 `protobuf:"varint,2,opt,name=bytes_tx,json=bytesTx,proto3" json:"bytes_tx,omitempty"`
	BytesRx   uint64 `protobuf:"varint,3,opt,name=bytes_rx,json=bytesRx,proto3" json:"bytes_rx,omitempty"`
	DroppedTx uint64 `protobuf:"varint,4,opt,name=dropped_tx,json=droppedTx,proto3" json:"dropped_tx,omitempty"`
	DroppedRx uint64 `protobuf:"varint,5,opt,name=dropped_rx,json=droppedRx,proto3" json:"dropped_rx,omitempty"`
}

func (x *SessionStatsUsage) Reset() {
	*x = SessionStatsUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_hh_sessionforwarder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionStatsUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionStatsUsage) ProtoMessage() {}

func (x *SessionStatsUsage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_hh_sessionforwarder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionStatsUsage.ProtoReflect.Descriptor instead.
func (*SessionStatsUsage) Descriptor() ([]byte, []int) {
	return file_pkg_hh_sessionforwarder_proto_rawDescGZIP(), []int{3}
}

func (x *SessionStatsUsage) GetRuleId() string {
	if x != nil {
		return x.RuleId
	}
	return ""
}

func (x *SessionStatsUsage) GetBytesTx() uint64 {
	if x != nil {
		return x.BytesTx
	}
	return 0
}

func (x *SessionStatsUsage) GetBytesRx() uint64 {
	if x != nil {
		return x.BytesRx
	}
	return 0
}

func (x *SessionStatsUsage) GetDroppedTx() uint64 {
	if x != nil {
		return x.DroppedTx
	}
	return 0
}

func (x *SessionStatsUsage) GetDroppedRx() uint64 {
	if x != nil {
		return x.DroppedRx
	}
	return 0
}

var File_pkg_hh_sessionforwarder_proto protoreflect.FileDescriptor

var file_pkg_hh_sessionforwarder_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x68, 0x68, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x68, 0x68, 0x22, 0x07, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x53, 0x0a, 0x1a,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0d, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x68, 0x68, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x22, 0xd3, 0x06, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x73, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x68, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x6d, 0x73, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x65, 0x69, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6d, 0x65, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x73, 0x69,
	0x73, 0x64, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x73, 0x69, 0x73, 0x64,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x6d, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x6d, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x70, 0x67,
	0x77, 0x5f, 0x70, 0x6c, 0x6d, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x67, 0x77, 0x50, 0x6c, 0x6d, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70,
	0x6e, 0x6e, 0x69, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x6e, 0x6e, 0x69,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x64, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x64, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x75,
	0x65, 0x5f, 0x69, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x65, 0x49, 0x70,
	0x12, 0x15, 0x0a, 0x06, 0x73, 0x67, 0x77, 0x5f, 0x69, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x67, 0x77, 0x49, 0x70, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x67, 0x77, 0x5f, 0x69,
	0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x67, 0x77, 0x49, 0x70, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x65, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x65, 0x49, 0x70, 0x76, 0x36, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x67, 0x77, 0x5f, 0x69,
	0x70, 0x76, 0x36, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x67, 0x77, 0x49, 0x70,
	0x76, 0x36, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x67, 0x77, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x67, 0x77, 0x49, 0x70, 0x76, 0x36, 0x12, 0x19, 0x0a,
	0x08, 0x72, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x72, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x13, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a,
	0x13, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x4f, 0x70, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x18, 0x63, 0x61, 0x75, 0x73, 0x65, 0x5f,
	0x66, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x69,
	0x6e, 0x67, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x63, 0x61, 0x75, 0x73, 0x65, 0x46,
	0x6f, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x49,
	0x64, 0x12, 0x33, 0x0a, 0x16, 0x72, 0x61, 0x77, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x13, 0x72, 0x61, 0x77, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x11, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f,
	0x74, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54,
	0x78, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x62, 0x79, 0x74, 0x65, 0x73, 0x52, 0x78, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x74, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x54, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x72, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x52, 0x78, 0x2a, 0x1b, 0x0a, 0x03, 0x62, 0x75,
	0x67, 0x12, 0x08, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x64, 0x10, 0x01, 0x32, 0x56, 0x0a, 0x10, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x13, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x1e, 0x2e, 0x68, 0x68, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x09, 0x2e, 0x68, 0x68, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61,
	0x61, 0x72, 0x61, 0x6f, 0x72, 0x61, 0x2f, 0x68, 0x68, 0x2d, 0x73, 0x74, 0x75, 0x62, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x68, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_hh_sessionforwarder_proto_rawDescOnce sync.Once
	file_pkg_hh_sessionforwarder_proto_rawDescData = file_pkg_hh_sessionforwarder_proto_rawDesc
)

func file_pkg_hh_sessionforwarder_proto_rawDescGZIP() []byte {
	file_pkg_hh_sessionforwarder_proto_rawDescOnce.Do(func() {
		file_pkg_hh_sessionforwarder_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_hh_sessionforwarder_proto_rawDescData)
	})
	return file_pkg_hh_sessionforwarder_proto_rawDescData
}

var file_pkg_hh_sessionforwarder_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_hh_sessionforwarder_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_hh_sessionforwarder_proto_goTypes = []interface{}{
	(Bug)(0),                           // 0: hh.bug
	(*Empty)(nil),                      // 1: hh.empty
	(*ReportSessionsStatsRequest)(nil), // 2: hh.reportSessionsStatsRequest
	(*SessionStats)(nil),               // 3: hh.sessionStats
	(*SessionStatsUsage)(nil),          // 4: hh.sessionStatsUsage
}
var file_pkg_hh_sessionforwarder_proto_depIdxs = []int32{
	3, // 0: hh.reportSessionsStatsRequest.session_stats:type_name -> hh.sessionStats
	4, // 1: hh.sessionStats.usage:type_name -> hh.sessionStatsUsage
	2, // 2: hh.sessionForwarder.reportSessionsStats:input_type -> hh.reportSessionsStatsRequest
	1, // 3: hh.sessionForwarder.reportSessionsStats:output_type -> hh.empty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_hh_sessionforwarder_proto_init() }
func file_pkg_hh_sessionforwarder_proto_init() {
	if File_pkg_hh_sessionforwarder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_hh_sessionforwarder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_pkg_hh_sessionforwarder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportSessionsStatsRequest); i {
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
		file_pkg_hh_sessionforwarder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionStats); i {
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
		file_pkg_hh_sessionforwarder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionStatsUsage); i {
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
			RawDescriptor: file_pkg_hh_sessionforwarder_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_hh_sessionforwarder_proto_goTypes,
		DependencyIndexes: file_pkg_hh_sessionforwarder_proto_depIdxs,
		EnumInfos:         file_pkg_hh_sessionforwarder_proto_enumTypes,
		MessageInfos:      file_pkg_hh_sessionforwarder_proto_msgTypes,
	}.Build()
	File_pkg_hh_sessionforwarder_proto = out.File
	file_pkg_hh_sessionforwarder_proto_rawDesc = nil
	file_pkg_hh_sessionforwarder_proto_goTypes = nil
	file_pkg_hh_sessionforwarder_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SessionForwarderClient is the client API for SessionForwarder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionForwarderClient interface {
	ReportSessionsStats(ctx context.Context, in *ReportSessionsStatsRequest, opts ...grpc.CallOption) (*Empty, error)
}

type sessionForwarderClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionForwarderClient(cc grpc.ClientConnInterface) SessionForwarderClient {
	return &sessionForwarderClient{cc}
}

func (c *sessionForwarderClient) ReportSessionsStats(ctx context.Context, in *ReportSessionsStatsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/hh.sessionForwarder/reportSessionsStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionForwarderServer is the server API for SessionForwarder service.
type SessionForwarderServer interface {
	ReportSessionsStats(context.Context, *ReportSessionsStatsRequest) (*Empty, error)
}

// UnimplementedSessionForwarderServer can be embedded to have forward compatible implementations.
type UnimplementedSessionForwarderServer struct {
}

func (*UnimplementedSessionForwarderServer) ReportSessionsStats(context.Context, *ReportSessionsStatsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportSessionsStats not implemented")
}

func RegisterSessionForwarderServer(s *grpc.Server, srv SessionForwarderServer) {
	s.RegisterService(&_SessionForwarder_serviceDesc, srv)
}

func _SessionForwarder_ReportSessionsStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportSessionsStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionForwarderServer).ReportSessionsStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hh.sessionForwarder/ReportSessionsStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionForwarderServer).ReportSessionsStats(ctx, req.(*ReportSessionsStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionForwarder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hh.sessionForwarder",
	HandlerType: (*SessionForwarderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "reportSessionsStats",
			Handler:    _SessionForwarder_ReportSessionsStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/hh/sessionforwarder.proto",
}