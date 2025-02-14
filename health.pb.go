// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: github.com/na4ma4/go-healthcheck/health.proto

package healthcheck

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type Status int32

const (
	Status_UNKNOWN  Status = 0
	Status_STARTING Status = 1
	Status_RUNNING  Status = 2
	Status_FINISHED Status = 3
	Status_ERRORED  Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "STARTING",
		2: "RUNNING",
		3: "FINISHED",
		4: "ERRORED",
	}
	Status_value = map[string]int32{
		"UNKNOWN":  0,
		"STARTING": 1,
		"RUNNING":  2,
		"FINISHED": 3,
		"ERRORED":  4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_na4ma4_go_healthcheck_health_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_github_com_na4ma4_go_healthcheck_health_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_github_com_na4ma4_go_healthcheck_health_proto_rawDescGZIP(), []int{0}
}

type ReportStatus int32

const (
	ReportStatus_NOTSET ReportStatus = 0
	ReportStatus_RED    ReportStatus = 1
	ReportStatus_YELLOW ReportStatus = 2
	ReportStatus_GREEN  ReportStatus = 3
)

// Enum value maps for ReportStatus.
var (
	ReportStatus_name = map[int32]string{
		0: "NOTSET",
		1: "RED",
		2: "YELLOW",
		3: "GREEN",
	}
	ReportStatus_value = map[string]int32{
		"NOTSET": 0,
		"RED":    1,
		"YELLOW": 2,
		"GREEN":  3,
	}
)

func (x ReportStatus) Enum() *ReportStatus {
	p := new(ReportStatus)
	*p = x
	return p
}

func (x ReportStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_na4ma4_go_healthcheck_health_proto_enumTypes[1].Descriptor()
}

func (ReportStatus) Type() protoreflect.EnumType {
	return &file_github_com_na4ma4_go_healthcheck_health_proto_enumTypes[1]
}

func (x ReportStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportStatus.Descriptor instead.
func (ReportStatus) EnumDescriptor() ([]byte, []int) {
	return file_github_com_na4ma4_go_healthcheck_health_proto_rawDescGZIP(), []int{1}
}

type CoreProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items map[string]*ItemProto `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CoreProto) Reset() {
	*x = CoreProto{}
	mi := &file_github_com_na4ma4_go_healthcheck_health_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CoreProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreProto) ProtoMessage() {}

func (x *CoreProto) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_na4ma4_go_healthcheck_health_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreProto.ProtoReflect.Descriptor instead.
func (*CoreProto) Descriptor() ([]byte, []int) {
	return file_github_com_na4ma4_go_healthcheck_health_proto_rawDescGZIP(), []int{0}
}

func (x *CoreProto) GetItems() map[string]*ItemProto {
	if x != nil {
		return x.Items
	}
	return nil
}

type ItemProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Lifecycle []*EventProto                    `protobuf:"bytes,2,rep,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	Times     map[int32]*timestamppb.Timestamp `protobuf:"bytes,3,rep,name=times,proto3" json:"times,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Duration  *durationpb.Duration             `protobuf:"bytes,10,opt,name=duration,proto3" json:"duration,omitempty"`
	StartTime *timestamppb.Timestamp           `protobuf:"bytes,11,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Status    Status                           `protobuf:"varint,30,opt,name=status,proto3,enum=healthcheck.Status" json:"status,omitempty"`
}

func (x *ItemProto) Reset() {
	*x = ItemProto{}
	mi := &file_github_com_na4ma4_go_healthcheck_health_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ItemProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemProto) ProtoMessage() {}

func (x *ItemProto) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_na4ma4_go_healthcheck_health_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemProto.ProtoReflect.Descriptor instead.
func (*ItemProto) Descriptor() ([]byte, []int) {
	return file_github_com_na4ma4_go_healthcheck_health_proto_rawDescGZIP(), []int{1}
}

func (x *ItemProto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ItemProto) GetLifecycle() []*EventProto {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *ItemProto) GetTimes() map[int32]*timestamppb.Timestamp {
	if x != nil {
		return x.Times
	}
	return nil
}

func (x *ItemProto) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *ItemProto) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *ItemProto) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN
}

type EventProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Status    Status                 `protobuf:"varint,2,opt,name=status,proto3,enum=healthcheck.Status" json:"status,omitempty"`
	Message   string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EventProto) Reset() {
	*x = EventProto{}
	mi := &file_github_com_na4ma4_go_healthcheck_health_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventProto) ProtoMessage() {}

func (x *EventProto) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_na4ma4_go_healthcheck_health_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventProto.ProtoReflect.Descriptor instead.
func (*EventProto) Descriptor() ([]byte, []int) {
	return file_github_com_na4ma4_go_healthcheck_health_proto_rawDescGZIP(), []int{2}
}

func (x *EventProto) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *EventProto) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN
}

func (x *EventProto) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var file_github_com_na4ma4_go_healthcheck_health_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         33000,
		Name:          "healthcheck.display_name",
		Tag:           "bytes,33000,opt,name=display_name",
		Filename:      "github.com/na4ma4/go-healthcheck/health.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string display_name = 33000;
	E_DisplayName = &file_github_com_na4ma4_go_healthcheck_health_proto_extTypes[0]
)

var File_github_com_na4ma4_go_healthcheck_health_proto protoreflect.FileDescriptor

var file_github_com_na4ma4_go_healthcheck_health_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x34,
	0x6d, 0x61, 0x34, 0x2f, 0x67, 0x6f, 0x2d, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x96, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x50, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x84, 0x03, 0x0a, 0x09, 0x49, 0x74, 0x65,
	0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x6c, 0x69,
	0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x12, 0x37, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x54, 0x0a, 0x0a, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x8d, 0x01, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x38,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a,
	0x8e, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x1a, 0x0b, 0xc2, 0x8e, 0x10, 0x07, 0x75, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x1a, 0x0c, 0xc2, 0x8e, 0x10, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x1a, 0x0b, 0xc2,
	0x8e, 0x10, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x49,
	0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x03, 0x1a, 0x0c, 0xc2, 0x8e, 0x10, 0x08, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x45,
	0x44, 0x10, 0x04, 0x1a, 0x0b, 0xc2, 0x8e, 0x10, 0x07, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x64,
	0x2a, 0x3a, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x54, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x59, 0x45, 0x4c, 0x4c, 0x4f, 0x57, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x03, 0x3a, 0x46, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xe8, 0x81, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x34, 0x6d, 0x61, 0x34, 0x2f, 0x67, 0x6f, 0x2d, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x3b, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_na4ma4_go_healthcheck_health_proto_rawDescOnce sync.Once
	file_github_com_na4ma4_go_healthcheck_health_proto_rawDescData = file_github_com_na4ma4_go_healthcheck_health_proto_rawDesc
)

func file_github_com_na4ma4_go_healthcheck_health_proto_rawDescGZIP() []byte {
	file_github_com_na4ma4_go_healthcheck_health_proto_rawDescOnce.Do(func() {
		file_github_com_na4ma4_go_healthcheck_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_na4ma4_go_healthcheck_health_proto_rawDescData)
	})
	return file_github_com_na4ma4_go_healthcheck_health_proto_rawDescData
}

var file_github_com_na4ma4_go_healthcheck_health_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_na4ma4_go_healthcheck_health_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_na4ma4_go_healthcheck_health_proto_goTypes = []any{
	(Status)(0),                           // 0: healthcheck.Status
	(ReportStatus)(0),                     // 1: healthcheck.ReportStatus
	(*CoreProto)(nil),                     // 2: healthcheck.CoreProto
	(*ItemProto)(nil),                     // 3: healthcheck.ItemProto
	(*EventProto)(nil),                    // 4: healthcheck.EventProto
	nil,                                   // 5: healthcheck.CoreProto.ItemsEntry
	nil,                                   // 6: healthcheck.ItemProto.TimesEntry
	(*durationpb.Duration)(nil),           // 7: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil),         // 8: google.protobuf.Timestamp
	(*descriptorpb.EnumValueOptions)(nil), // 9: google.protobuf.EnumValueOptions
}
var file_github_com_na4ma4_go_healthcheck_health_proto_depIdxs = []int32{
	5,  // 0: healthcheck.CoreProto.items:type_name -> healthcheck.CoreProto.ItemsEntry
	4,  // 1: healthcheck.ItemProto.lifecycle:type_name -> healthcheck.EventProto
	6,  // 2: healthcheck.ItemProto.times:type_name -> healthcheck.ItemProto.TimesEntry
	7,  // 3: healthcheck.ItemProto.duration:type_name -> google.protobuf.Duration
	8,  // 4: healthcheck.ItemProto.start_time:type_name -> google.protobuf.Timestamp
	0,  // 5: healthcheck.ItemProto.status:type_name -> healthcheck.Status
	8,  // 6: healthcheck.EventProto.timestamp:type_name -> google.protobuf.Timestamp
	0,  // 7: healthcheck.EventProto.status:type_name -> healthcheck.Status
	3,  // 8: healthcheck.CoreProto.ItemsEntry.value:type_name -> healthcheck.ItemProto
	8,  // 9: healthcheck.ItemProto.TimesEntry.value:type_name -> google.protobuf.Timestamp
	9,  // 10: healthcheck.display_name:extendee -> google.protobuf.EnumValueOptions
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	10, // [10:11] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_github_com_na4ma4_go_healthcheck_health_proto_init() }
func file_github_com_na4ma4_go_healthcheck_health_proto_init() {
	if File_github_com_na4ma4_go_healthcheck_health_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_na4ma4_go_healthcheck_health_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_github_com_na4ma4_go_healthcheck_health_proto_goTypes,
		DependencyIndexes: file_github_com_na4ma4_go_healthcheck_health_proto_depIdxs,
		EnumInfos:         file_github_com_na4ma4_go_healthcheck_health_proto_enumTypes,
		MessageInfos:      file_github_com_na4ma4_go_healthcheck_health_proto_msgTypes,
		ExtensionInfos:    file_github_com_na4ma4_go_healthcheck_health_proto_extTypes,
	}.Build()
	File_github_com_na4ma4_go_healthcheck_health_proto = out.File
	file_github_com_na4ma4_go_healthcheck_health_proto_rawDesc = nil
	file_github_com_na4ma4_go_healthcheck_health_proto_goTypes = nil
	file_github_com_na4ma4_go_healthcheck_health_proto_depIdxs = nil
}
