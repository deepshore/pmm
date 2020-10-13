// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: qanpb/profile.proto

package qanpb

import (
	context "context"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// ReportRequest defines filtering of metrics report for db server or other dimentions.
type ReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeriodStartFrom *timestamp.Timestamp   `protobuf:"bytes,1,opt,name=period_start_from,json=periodStartFrom,proto3" json:"period_start_from,omitempty"`
	PeriodStartTo   *timestamp.Timestamp   `protobuf:"bytes,2,opt,name=period_start_to,json=periodStartTo,proto3" json:"period_start_to,omitempty"`
	GroupBy         string                 `protobuf:"bytes,3,opt,name=group_by,json=groupBy,proto3" json:"group_by,omitempty"`
	Labels          []*ReportMapFieldEntry `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	Columns         []string               `protobuf:"bytes,5,rep,name=columns,proto3" json:"columns,omitempty"`
	OrderBy         string                 `protobuf:"bytes,6,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	Offset          uint32                 `protobuf:"varint,7,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit           uint32                 `protobuf:"varint,8,opt,name=limit,proto3" json:"limit,omitempty"`
	MainMetric      string                 `protobuf:"bytes,9,opt,name=main_metric,json=mainMetric,proto3" json:"main_metric,omitempty"`
	Search          string                 `protobuf:"bytes,10,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qanpb_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qanpb_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRequest.ProtoReflect.Descriptor instead.
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return file_qanpb_profile_proto_rawDescGZIP(), []int{0}
}

func (x *ReportRequest) GetPeriodStartFrom() *timestamp.Timestamp {
	if x != nil {
		return x.PeriodStartFrom
	}
	return nil
}

func (x *ReportRequest) GetPeriodStartTo() *timestamp.Timestamp {
	if x != nil {
		return x.PeriodStartTo
	}
	return nil
}

func (x *ReportRequest) GetGroupBy() string {
	if x != nil {
		return x.GroupBy
	}
	return ""
}

func (x *ReportRequest) GetLabels() []*ReportMapFieldEntry {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ReportRequest) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *ReportRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ReportRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ReportRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ReportRequest) GetMainMetric() string {
	if x != nil {
		return x.MainMetric
	}
	return ""
}

func (x *ReportRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

// ReportMapFieldEntry allows to pass labels/dimentions in form like {"server": ["db1", "db2"...]}.
type ReportMapFieldEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []string `protobuf:"bytes,2,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *ReportMapFieldEntry) Reset() {
	*x = ReportMapFieldEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qanpb_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportMapFieldEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportMapFieldEntry) ProtoMessage() {}

func (x *ReportMapFieldEntry) ProtoReflect() protoreflect.Message {
	mi := &file_qanpb_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportMapFieldEntry.ProtoReflect.Descriptor instead.
func (*ReportMapFieldEntry) Descriptor() ([]byte, []int) {
	return file_qanpb_profile_proto_rawDescGZIP(), []int{1}
}

func (x *ReportMapFieldEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ReportMapFieldEntry) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

// ReportReply is list of reports per quieryids, hosts etc.
type ReportReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalRows uint32 `protobuf:"varint,1,opt,name=total_rows,json=totalRows,proto3" json:"total_rows,omitempty"`
	Offset    uint32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit     uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Rows      []*Row `protobuf:"bytes,4,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *ReportReply) Reset() {
	*x = ReportReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qanpb_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportReply) ProtoMessage() {}

func (x *ReportReply) ProtoReflect() protoreflect.Message {
	mi := &file_qanpb_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportReply.ProtoReflect.Descriptor instead.
func (*ReportReply) Descriptor() ([]byte, []int) {
	return file_qanpb_profile_proto_rawDescGZIP(), []int{2}
}

func (x *ReportReply) GetTotalRows() uint32 {
	if x != nil {
		return x.TotalRows
	}
	return 0
}

func (x *ReportReply) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ReportReply) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ReportReply) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

// Row define metrics for selected dimention.
type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank        uint32             `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Dimension   string             `protobuf:"bytes,2,opt,name=dimension,proto3" json:"dimension,omitempty"`
	Metrics     map[string]*Metric `protobuf:"bytes,3,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Sparkline   []*Point           `protobuf:"bytes,4,rep,name=sparkline,proto3" json:"sparkline,omitempty"`
	Fingerprint string             `protobuf:"bytes,5,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	NumQueries  uint32             `protobuf:"varint,6,opt,name=num_queries,json=numQueries,proto3" json:"num_queries,omitempty"`
	Qps         float32            `protobuf:"fixed32,7,opt,name=qps,proto3" json:"qps,omitempty"`
	Load        float32            `protobuf:"fixed32,8,opt,name=load,proto3" json:"load,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qanpb_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_qanpb_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_qanpb_profile_proto_rawDescGZIP(), []int{3}
}

func (x *Row) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Row) GetDimension() string {
	if x != nil {
		return x.Dimension
	}
	return ""
}

func (x *Row) GetMetrics() map[string]*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *Row) GetSparkline() []*Point {
	if x != nil {
		return x.Sparkline
	}
	return nil
}

func (x *Row) GetFingerprint() string {
	if x != nil {
		return x.Fingerprint
	}
	return ""
}

func (x *Row) GetNumQueries() uint32 {
	if x != nil {
		return x.NumQueries
	}
	return 0
}

func (x *Row) GetQps() float32 {
	if x != nil {
		return x.Qps
	}
	return 0
}

func (x *Row) GetLoad() float32 {
	if x != nil {
		return x.Load
	}
	return 0
}

// Metric cell.
type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats *Stat `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qanpb_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_qanpb_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_qanpb_profile_proto_rawDescGZIP(), []int{4}
}

func (x *Metric) GetStats() *Stat {
	if x != nil {
		return x.Stats
	}
	return nil
}

// Stat is statistics of specific metric.
type Stat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rate      float32 `protobuf:"fixed32,1,opt,name=rate,proto3" json:"rate,omitempty"`
	Cnt       float32 `protobuf:"fixed32,2,opt,name=cnt,proto3" json:"cnt,omitempty"`
	Sum       float32 `protobuf:"fixed32,3,opt,name=sum,proto3" json:"sum,omitempty"`
	Min       float32 `protobuf:"fixed32,4,opt,name=min,proto3" json:"min,omitempty"`
	Max       float32 `protobuf:"fixed32,5,opt,name=max,proto3" json:"max,omitempty"`
	P99       float32 `protobuf:"fixed32,6,opt,name=p99,proto3" json:"p99,omitempty"`
	Avg       float32 `protobuf:"fixed32,7,opt,name=avg,proto3" json:"avg,omitempty"`
	SumPerSec float32 `protobuf:"fixed32,8,opt,name=sum_per_sec,json=sumPerSec,proto3" json:"sum_per_sec,omitempty"`
}

func (x *Stat) Reset() {
	*x = Stat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qanpb_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stat) ProtoMessage() {}

func (x *Stat) ProtoReflect() protoreflect.Message {
	mi := &file_qanpb_profile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stat.ProtoReflect.Descriptor instead.
func (*Stat) Descriptor() ([]byte, []int) {
	return file_qanpb_profile_proto_rawDescGZIP(), []int{5}
}

func (x *Stat) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *Stat) GetCnt() float32 {
	if x != nil {
		return x.Cnt
	}
	return 0
}

func (x *Stat) GetSum() float32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

func (x *Stat) GetMin() float32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Stat) GetMax() float32 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *Stat) GetP99() float32 {
	if x != nil {
		return x.P99
	}
	return 0
}

func (x *Stat) GetAvg() float32 {
	if x != nil {
		return x.Avg
	}
	return 0
}

func (x *Stat) GetSumPerSec() float32 {
	if x != nil {
		return x.SumPerSec
	}
	return 0
}

var File_qanpb_profile_proto protoreflect.FileDescriptor

var file_qanpb_profile_proto_rawDesc = []byte{
	0x0a, 0x13, 0x71, 0x61, 0x6e, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x71, 0x61, 0x6e, 0x70, 0x62, 0x2f, 0x71, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8c, 0x03, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x42, 0x0a, 0x0f,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0d, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x6f,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x71, 0x61,
	0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x61, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x22, 0x3d, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x80, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x6f, 0x77, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x24, 0x0a,
	0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x71, 0x61,
	0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72,
	0x6f, 0x77, 0x73, 0x22, 0xdc, 0x02, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x61, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a,
	0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x6f, 0x77,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x30, 0x0a, 0x09, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x71, 0x61, 0x6e, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x09, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75,
	0x6d, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x6e, 0x75, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x71,
	0x70, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x71, 0x70, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6c, 0x6f, 0x61,
	0x64, 0x1a, 0x4f, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x31, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x27, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x71, 0x61,
	0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x61,
	0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x03, 0x63, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x39,
	0x39, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x70, 0x39, 0x39, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x76, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x61, 0x76, 0x67, 0x12, 0x1e,
	0x0a, 0x0b, 0x73, 0x75, 0x6d, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x73, 0x75, 0x6d, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x32, 0x6a,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x5f, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x76, 0x30, 0x2f, 0x71, 0x61, 0x6e, 0x2f, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x11, 0x5a, 0x0f, 0x61, 0x70,
	0x69, 0x2f, 0x71, 0x61, 0x6e, 0x70, 0x62, 0x3b, 0x71, 0x61, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qanpb_profile_proto_rawDescOnce sync.Once
	file_qanpb_profile_proto_rawDescData = file_qanpb_profile_proto_rawDesc
)

func file_qanpb_profile_proto_rawDescGZIP() []byte {
	file_qanpb_profile_proto_rawDescOnce.Do(func() {
		file_qanpb_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_qanpb_profile_proto_rawDescData)
	})
	return file_qanpb_profile_proto_rawDescData
}

var file_qanpb_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_qanpb_profile_proto_goTypes = []interface{}{
	(*ReportRequest)(nil),       // 0: qan.v1beta1.ReportRequest
	(*ReportMapFieldEntry)(nil), // 1: qan.v1beta1.ReportMapFieldEntry
	(*ReportReply)(nil),         // 2: qan.v1beta1.ReportReply
	(*Row)(nil),                 // 3: qan.v1beta1.Row
	(*Metric)(nil),              // 4: qan.v1beta1.Metric
	(*Stat)(nil),                // 5: qan.v1beta1.Stat
	nil,                         // 6: qan.v1beta1.Row.MetricsEntry
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*Point)(nil),               // 8: qan.v1beta1.Point
}
var file_qanpb_profile_proto_depIdxs = []int32{
	7, // 0: qan.v1beta1.ReportRequest.period_start_from:type_name -> google.protobuf.Timestamp
	7, // 1: qan.v1beta1.ReportRequest.period_start_to:type_name -> google.protobuf.Timestamp
	1, // 2: qan.v1beta1.ReportRequest.labels:type_name -> qan.v1beta1.ReportMapFieldEntry
	3, // 3: qan.v1beta1.ReportReply.rows:type_name -> qan.v1beta1.Row
	6, // 4: qan.v1beta1.Row.metrics:type_name -> qan.v1beta1.Row.MetricsEntry
	8, // 5: qan.v1beta1.Row.sparkline:type_name -> qan.v1beta1.Point
	5, // 6: qan.v1beta1.Metric.stats:type_name -> qan.v1beta1.Stat
	4, // 7: qan.v1beta1.Row.MetricsEntry.value:type_name -> qan.v1beta1.Metric
	0, // 8: qan.v1beta1.Profile.GetReport:input_type -> qan.v1beta1.ReportRequest
	2, // 9: qan.v1beta1.Profile.GetReport:output_type -> qan.v1beta1.ReportReply
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_qanpb_profile_proto_init() }
func file_qanpb_profile_proto_init() {
	if File_qanpb_profile_proto != nil {
		return
	}
	file_qanpb_qan_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_qanpb_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRequest); i {
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
		file_qanpb_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportMapFieldEntry); i {
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
		file_qanpb_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportReply); i {
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
		file_qanpb_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_qanpb_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
		file_qanpb_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stat); i {
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
			RawDescriptor: file_qanpb_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qanpb_profile_proto_goTypes,
		DependencyIndexes: file_qanpb_profile_proto_depIdxs,
		MessageInfos:      file_qanpb_profile_proto_msgTypes,
	}.Build()
	File_qanpb_profile_proto = out.File
	file_qanpb_profile_proto_rawDesc = nil
	file_qanpb_profile_proto_goTypes = nil
	file_qanpb_profile_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileClient interface {
	// GetReport returns list of metrics group by queryid or other dimentions.
	GetReport(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error)
}

type profileClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileClient(cc grpc.ClientConnInterface) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) GetReport(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error) {
	out := new(ReportReply)
	err := c.cc.Invoke(ctx, "/qan.v1beta1.Profile/GetReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
type ProfileServer interface {
	// GetReport returns list of metrics group by queryid or other dimentions.
	GetReport(context.Context, *ReportRequest) (*ReportReply, error)
}

// UnimplementedProfileServer can be embedded to have forward compatible implementations.
type UnimplementedProfileServer struct {
}

func (*UnimplementedProfileServer) GetReport(context.Context, *ReportRequest) (*ReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}

func RegisterProfileServer(s *grpc.Server, srv ProfileServer) {
	s.RegisterService(&_Profile_serviceDesc, srv)
}

func _Profile_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.v1beta1.Profile/GetReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetReport(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Profile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qan.v1beta1.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReport",
			Handler:    _Profile_GetReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qanpb/profile.proto",
}
