// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: proto/bodyjson/bodyjson.proto

package bodyjson

import (
	examplepb "github.com/RyoJerryYu/protoc-gen-pluginx/tests/protoc-gen-ts-grpc-gateway-cli/integration-everything-simple/server/proto/examplepb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WellKnownTypesHolder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayloadCheck string                  `protobuf:"bytes,100,opt,name=payload_check,json=payloadCheck,proto3" json:"payload_check,omitempty"`
	Any          *anypb.Any              `protobuf:"bytes,1,opt,name=any,proto3" json:"any,omitempty"`
	Empty        *emptypb.Empty          `protobuf:"bytes,2,opt,name=empty,proto3" json:"empty,omitempty"`
	Struct       *structpb.Struct        `protobuf:"bytes,3,opt,name=struct,proto3" json:"struct,omitempty"`
	Value        *structpb.Value         `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	ListValue    *structpb.ListValue     `protobuf:"bytes,5,opt,name=list_value,json=listValue,proto3" json:"list_value,omitempty"`
	Timestamp    *timestamppb.Timestamp  `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Duration     *durationpb.Duration    `protobuf:"bytes,7,opt,name=duration,proto3" json:"duration,omitempty"`
	FieldMask    *fieldmaskpb.FieldMask  `protobuf:"bytes,8,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	BoolValue    *wrapperspb.BoolValue   `protobuf:"bytes,9,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	Int32Value   *wrapperspb.Int32Value  `protobuf:"bytes,10,opt,name=int32_value,json=int32Value,proto3" json:"int32_value,omitempty"`
	Uint32Value  *wrapperspb.UInt32Value `protobuf:"bytes,11,opt,name=uint32_value,json=uint32Value,proto3" json:"uint32_value,omitempty"`
	Int64Value   *wrapperspb.Int64Value  `protobuf:"bytes,12,opt,name=int64_value,json=int64Value,proto3" json:"int64_value,omitempty"`
	Uint64Value  *wrapperspb.UInt64Value `protobuf:"bytes,13,opt,name=uint64_value,json=uint64Value,proto3" json:"uint64_value,omitempty"`
	StringValue  *wrapperspb.StringValue `protobuf:"bytes,14,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	BytesValue   *wrapperspb.BytesValue  `protobuf:"bytes,15,opt,name=bytes_value,json=bytesValue,proto3" json:"bytes_value,omitempty"`
	DoubleValue  *wrapperspb.DoubleValue `protobuf:"bytes,16,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	FloatValue   *wrapperspb.FloatValue  `protobuf:"bytes,17,opt,name=float_value,json=floatValue,proto3" json:"float_value,omitempty"`
}

func (x *WellKnownTypesHolder) Reset() {
	*x = WellKnownTypesHolder{}
	mi := &file_proto_bodyjson_bodyjson_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WellKnownTypesHolder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WellKnownTypesHolder) ProtoMessage() {}

func (x *WellKnownTypesHolder) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bodyjson_bodyjson_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WellKnownTypesHolder.ProtoReflect.Descriptor instead.
func (*WellKnownTypesHolder) Descriptor() ([]byte, []int) {
	return file_proto_bodyjson_bodyjson_proto_rawDescGZIP(), []int{0}
}

func (x *WellKnownTypesHolder) GetPayloadCheck() string {
	if x != nil {
		return x.PayloadCheck
	}
	return ""
}

func (x *WellKnownTypesHolder) GetAny() *anypb.Any {
	if x != nil {
		return x.Any
	}
	return nil
}

func (x *WellKnownTypesHolder) GetEmpty() *emptypb.Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}

func (x *WellKnownTypesHolder) GetStruct() *structpb.Struct {
	if x != nil {
		return x.Struct
	}
	return nil
}

func (x *WellKnownTypesHolder) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *WellKnownTypesHolder) GetListValue() *structpb.ListValue {
	if x != nil {
		return x.ListValue
	}
	return nil
}

func (x *WellKnownTypesHolder) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *WellKnownTypesHolder) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *WellKnownTypesHolder) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

func (x *WellKnownTypesHolder) GetBoolValue() *wrapperspb.BoolValue {
	if x != nil {
		return x.BoolValue
	}
	return nil
}

func (x *WellKnownTypesHolder) GetInt32Value() *wrapperspb.Int32Value {
	if x != nil {
		return x.Int32Value
	}
	return nil
}

func (x *WellKnownTypesHolder) GetUint32Value() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Uint32Value
	}
	return nil
}

func (x *WellKnownTypesHolder) GetInt64Value() *wrapperspb.Int64Value {
	if x != nil {
		return x.Int64Value
	}
	return nil
}

func (x *WellKnownTypesHolder) GetUint64Value() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Uint64Value
	}
	return nil
}

func (x *WellKnownTypesHolder) GetStringValue() *wrapperspb.StringValue {
	if x != nil {
		return x.StringValue
	}
	return nil
}

func (x *WellKnownTypesHolder) GetBytesValue() *wrapperspb.BytesValue {
	if x != nil {
		return x.BytesValue
	}
	return nil
}

func (x *WellKnownTypesHolder) GetDoubleValue() *wrapperspb.DoubleValue {
	if x != nil {
		return x.DoubleValue
	}
	return nil
}

func (x *WellKnownTypesHolder) GetFloatValue() *wrapperspb.FloatValue {
	if x != nil {
		return x.FloatValue
	}
	return nil
}

var File_proto_bodyjson_bodyjson_proto protoreflect.FileDescriptor

var file_proto_bodyjson_bodyjson_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e,
	0x2f, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2f, 0x61, 0x5f, 0x62, 0x69, 0x74, 0x5f, 0x6f, 0x66,
	0x5f, 0x65, 0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8e, 0x08, 0x0a, 0x14, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x26, 0x0a, 0x03, 0x61, 0x6e, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x03, 0x61, 0x6e, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x05,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x62,
	0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x62, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x32, 0xd1, 0x0b, 0x0a, 0x0f, 0x42, 0x6f, 0x64, 0x79, 0x4a, 0x53, 0x4f, 0x4e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x42, 0x69, 0x74,
	0x4f, 0x66, 0x45, 0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x1a, 0x21, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x41,
	0x42, 0x69, 0x74, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22,
	0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x6a,
	0x73, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x91,
	0x01, 0x0a, 0x17, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x42, 0x69,
	0x74, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x1a, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e,
	0x41, 0x42, 0x69, 0x74, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x9e, 0x01, 0x0a, 0x16, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e,
	0x41, 0x42, 0x69, 0x74, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x70, 0x62, 0x2e, 0x41, 0x42, 0x69, 0x74, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x72, 0x79, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x22, 0x3e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x38, 0x3a, 0x15, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e,
	0x2f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x62,
	0x6f, 0x64, 0x79, 0x12, 0x8e, 0x01, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b,
	0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x1a,
	0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e,
	0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x48,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f,
	0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x8f, 0x01, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x6c, 0x6c,
	0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f,
	0x6e, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x0a,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61,
	0x73, 0x6b, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x85, 0x01, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b,
	0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x1a,
	0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e,
	0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x48,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x06, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x6a,
	0x73, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x82,
	0x01, 0x0a, 0x0d, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f,
	0x6e, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62,
	0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x25, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x16, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62,
	0x6f, 0x64, 0x79, 0x12, 0x8f, 0x01, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b,
	0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x1a,
	0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e,
	0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x48,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x0a, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x8c, 0x01, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x57, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b,
	0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x1a,
	0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e,
	0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x48,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x0b, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0xaf, 0x01, 0x0a, 0x16, 0x50, 0x61, 0x74, 0x63, 0x68, 0x42, 0x6f,
	0x64, 0x79, 0x57, 0x69, 0x74, 0x68, 0x50, 0x61, 0x74, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12,
	0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70,
	0x62, 0x2e, 0x41, 0x42, 0x69, 0x74, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x42, 0x69, 0x74, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x72, 0x79,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x4f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x49, 0x3a, 0x0d, 0x73,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x32, 0x38, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x62, 0x6f, 0x64, 0x79, 0x77, 0x69, 0x74, 0x68, 0x70, 0x61, 0x74, 0x68, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x2f, 0x7b, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x80, 0x02, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x42, 0x0d, 0x42,
	0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x81,
	0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x79, 0x6f, 0x4a,
	0x65, 0x72, 0x72, 0x79, 0x59, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x78, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x74, 0x73, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x63, 0x6c, 0x69, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x65, 0x76, 0x65, 0x72, 0x79,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x6a, 0x73,
	0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x50, 0x42, 0x58, 0xaa, 0x02, 0x0e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x42, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0xca, 0x02, 0x0e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5c, 0x42, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0xe2, 0x02, 0x1a, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5c, 0x42, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3a,
	0x3a, 0x42, 0x6f, 0x64, 0x79, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_bodyjson_bodyjson_proto_rawDescOnce sync.Once
	file_proto_bodyjson_bodyjson_proto_rawDescData = file_proto_bodyjson_bodyjson_proto_rawDesc
)

func file_proto_bodyjson_bodyjson_proto_rawDescGZIP() []byte {
	file_proto_bodyjson_bodyjson_proto_rawDescOnce.Do(func() {
		file_proto_bodyjson_bodyjson_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bodyjson_bodyjson_proto_rawDescData)
	})
	return file_proto_bodyjson_bodyjson_proto_rawDescData
}

var file_proto_bodyjson_bodyjson_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_bodyjson_bodyjson_proto_goTypes = []any{
	(*WellKnownTypesHolder)(nil),       // 0: proto.bodyjson.WellKnownTypesHolder
	(*anypb.Any)(nil),                  // 1: google.protobuf.Any
	(*emptypb.Empty)(nil),              // 2: google.protobuf.Empty
	(*structpb.Struct)(nil),            // 3: google.protobuf.Struct
	(*structpb.Value)(nil),             // 4: google.protobuf.Value
	(*structpb.ListValue)(nil),         // 5: google.protobuf.ListValue
	(*timestamppb.Timestamp)(nil),      // 6: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),        // 7: google.protobuf.Duration
	(*fieldmaskpb.FieldMask)(nil),      // 8: google.protobuf.FieldMask
	(*wrapperspb.BoolValue)(nil),       // 9: google.protobuf.BoolValue
	(*wrapperspb.Int32Value)(nil),      // 10: google.protobuf.Int32Value
	(*wrapperspb.UInt32Value)(nil),     // 11: google.protobuf.UInt32Value
	(*wrapperspb.Int64Value)(nil),      // 12: google.protobuf.Int64Value
	(*wrapperspb.UInt64Value)(nil),     // 13: google.protobuf.UInt64Value
	(*wrapperspb.StringValue)(nil),     // 14: google.protobuf.StringValue
	(*wrapperspb.BytesValue)(nil),      // 15: google.protobuf.BytesValue
	(*wrapperspb.DoubleValue)(nil),     // 16: google.protobuf.DoubleValue
	(*wrapperspb.FloatValue)(nil),      // 17: google.protobuf.FloatValue
	(*examplepb.ABitOfEverything)(nil), // 18: proto.examplepb.ABitOfEverything
}
var file_proto_bodyjson_bodyjson_proto_depIdxs = []int32{
	1,  // 0: proto.bodyjson.WellKnownTypesHolder.any:type_name -> google.protobuf.Any
	2,  // 1: proto.bodyjson.WellKnownTypesHolder.empty:type_name -> google.protobuf.Empty
	3,  // 2: proto.bodyjson.WellKnownTypesHolder.struct:type_name -> google.protobuf.Struct
	4,  // 3: proto.bodyjson.WellKnownTypesHolder.value:type_name -> google.protobuf.Value
	5,  // 4: proto.bodyjson.WellKnownTypesHolder.list_value:type_name -> google.protobuf.ListValue
	6,  // 5: proto.bodyjson.WellKnownTypesHolder.timestamp:type_name -> google.protobuf.Timestamp
	7,  // 6: proto.bodyjson.WellKnownTypesHolder.duration:type_name -> google.protobuf.Duration
	8,  // 7: proto.bodyjson.WellKnownTypesHolder.field_mask:type_name -> google.protobuf.FieldMask
	9,  // 8: proto.bodyjson.WellKnownTypesHolder.bool_value:type_name -> google.protobuf.BoolValue
	10, // 9: proto.bodyjson.WellKnownTypesHolder.int32_value:type_name -> google.protobuf.Int32Value
	11, // 10: proto.bodyjson.WellKnownTypesHolder.uint32_value:type_name -> google.protobuf.UInt32Value
	12, // 11: proto.bodyjson.WellKnownTypesHolder.int64_value:type_name -> google.protobuf.Int64Value
	13, // 12: proto.bodyjson.WellKnownTypesHolder.uint64_value:type_name -> google.protobuf.UInt64Value
	14, // 13: proto.bodyjson.WellKnownTypesHolder.string_value:type_name -> google.protobuf.StringValue
	15, // 14: proto.bodyjson.WellKnownTypesHolder.bytes_value:type_name -> google.protobuf.BytesValue
	16, // 15: proto.bodyjson.WellKnownTypesHolder.double_value:type_name -> google.protobuf.DoubleValue
	17, // 16: proto.bodyjson.WellKnownTypesHolder.float_value:type_name -> google.protobuf.FloatValue
	18, // 17: proto.bodyjson.BodyJSONService.PostStringBody:input_type -> proto.examplepb.ABitOfEverything
	18, // 18: proto.bodyjson.BodyJSONService.PostRepeatedMessageBody:input_type -> proto.examplepb.ABitOfEverything
	18, // 19: proto.bodyjson.BodyJSONService.PostRepeatedStringBody:input_type -> proto.examplepb.ABitOfEverything
	0,  // 20: proto.bodyjson.BodyJSONService.PostTimestampBody:input_type -> proto.bodyjson.WellKnownTypesHolder
	0,  // 21: proto.bodyjson.BodyJSONService.PostFieldMaskBody:input_type -> proto.bodyjson.WellKnownTypesHolder
	0,  // 22: proto.bodyjson.BodyJSONService.PostStructBody:input_type -> proto.bodyjson.WellKnownTypesHolder
	0,  // 23: proto.bodyjson.BodyJSONService.PostValueBody:input_type -> proto.bodyjson.WellKnownTypesHolder
	0,  // 24: proto.bodyjson.BodyJSONService.PostListValueBody:input_type -> proto.bodyjson.WellKnownTypesHolder
	0,  // 25: proto.bodyjson.BodyJSONService.PostWrapperBody:input_type -> proto.bodyjson.WellKnownTypesHolder
	18, // 26: proto.bodyjson.BodyJSONService.PatchBodyWithPathParam:input_type -> proto.examplepb.ABitOfEverything
	18, // 27: proto.bodyjson.BodyJSONService.PostStringBody:output_type -> proto.examplepb.ABitOfEverything
	18, // 28: proto.bodyjson.BodyJSONService.PostRepeatedMessageBody:output_type -> proto.examplepb.ABitOfEverything
	18, // 29: proto.bodyjson.BodyJSONService.PostRepeatedStringBody:output_type -> proto.examplepb.ABitOfEverything
	0,  // 30: proto.bodyjson.BodyJSONService.PostTimestampBody:output_type -> proto.bodyjson.WellKnownTypesHolder
	0,  // 31: proto.bodyjson.BodyJSONService.PostFieldMaskBody:output_type -> proto.bodyjson.WellKnownTypesHolder
	0,  // 32: proto.bodyjson.BodyJSONService.PostStructBody:output_type -> proto.bodyjson.WellKnownTypesHolder
	0,  // 33: proto.bodyjson.BodyJSONService.PostValueBody:output_type -> proto.bodyjson.WellKnownTypesHolder
	0,  // 34: proto.bodyjson.BodyJSONService.PostListValueBody:output_type -> proto.bodyjson.WellKnownTypesHolder
	0,  // 35: proto.bodyjson.BodyJSONService.PostWrapperBody:output_type -> proto.bodyjson.WellKnownTypesHolder
	18, // 36: proto.bodyjson.BodyJSONService.PatchBodyWithPathParam:output_type -> proto.examplepb.ABitOfEverything
	27, // [27:37] is the sub-list for method output_type
	17, // [17:27] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_proto_bodyjson_bodyjson_proto_init() }
func file_proto_bodyjson_bodyjson_proto_init() {
	if File_proto_bodyjson_bodyjson_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_bodyjson_bodyjson_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_bodyjson_bodyjson_proto_goTypes,
		DependencyIndexes: file_proto_bodyjson_bodyjson_proto_depIdxs,
		MessageInfos:      file_proto_bodyjson_bodyjson_proto_msgTypes,
	}.Build()
	File_proto_bodyjson_bodyjson_proto = out.File
	file_proto_bodyjson_bodyjson_proto_rawDesc = nil
	file_proto_bodyjson_bodyjson_proto_goTypes = nil
	file_proto_bodyjson_bodyjson_proto_depIdxs = nil
}
