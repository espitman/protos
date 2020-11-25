// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: acm/acm.proto

package acm

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RequestCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"title",validate:"required"
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title" validate:"required"`
	// @inject_tag: json:"host",validate:"required"
	Host int32 `protobuf:"varint,2,opt,name=host,proto3" json:"host" validate:"required"`
	// @inject_tag: json:"description",validate:"required"
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description" validate:"required"`
	// @inject_tag: json:"type",validate:"required"
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type" validate:"required"`
	// @inject_tag: json:"reservationType",validate:"required|oneof=instant preApprove"
	ReservationType string `protobuf:"bytes,5,opt,name=reservationType,proto3" json:"reservationType" validate:"required|oneof=instant preApprove"`
	// @inject_tag: json:"placeType",validate:"required|oneof=personalRoom sharedRoom"
	PlaceType string `protobuf:"bytes,6,opt,name=placeType,proto3" json:"placeType" validate:"required|oneof=personalRoom sharedRoom"`
	// @inject_tag: json:"pricingType",validate:"required"
	PricingType string `protobuf:"bytes,7,opt,name=pricingType,proto3" json:"pricingType" validate:"required"`
	// @inject_tag: json:"minNight",validate:"required|number|min=1|max=99"
	MinNight int32 `protobuf:"varint,8,opt,name=minNight,proto3" json:"minNight" validate:"required|number|min=1|max=99"`
	// @inject_tag: json:"capacity",validate:"required"
	Capacity *RequestCreate_Capacity `protobuf:"bytes,9,opt,name=capacity,proto3" json:"capacity" validate:"required"`
}

func (x *RequestCreate) Reset() {
	*x = RequestCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acm_acm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCreate) ProtoMessage() {}

func (x *RequestCreate) ProtoReflect() protoreflect.Message {
	mi := &file_acm_acm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCreate.ProtoReflect.Descriptor instead.
func (*RequestCreate) Descriptor() ([]byte, []int) {
	return file_acm_acm_proto_rawDescGZIP(), []int{0}
}

func (x *RequestCreate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RequestCreate) GetHost() int32 {
	if x != nil {
		return x.Host
	}
	return 0
}

func (x *RequestCreate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RequestCreate) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RequestCreate) GetReservationType() string {
	if x != nil {
		return x.ReservationType
	}
	return ""
}

func (x *RequestCreate) GetPlaceType() string {
	if x != nil {
		return x.PlaceType
	}
	return ""
}

func (x *RequestCreate) GetPricingType() string {
	if x != nil {
		return x.PricingType
	}
	return ""
}

func (x *RequestCreate) GetMinNight() int32 {
	if x != nil {
		return x.MinNight
	}
	return 0
}

func (x *RequestCreate) GetCapacity() *RequestCreate_Capacity {
	if x != nil {
		return x.Capacity
	}
	return nil
}

type RequestDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestDetails) Reset() {
	*x = RequestDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acm_acm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDetails) ProtoMessage() {}

func (x *RequestDetails) ProtoReflect() protoreflect.Message {
	mi := &file_acm_acm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDetails.ProtoReflect.Descriptor instead.
func (*RequestDetails) Descriptor() ([]byte, []int) {
	return file_acm_acm_proto_rawDescGZIP(), []int{1}
}

func (x *RequestDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ResponseDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title           string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Host            int32  `protobuf:"varint,3,opt,name=host,proto3" json:"host,omitempty"`
	Description     string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Type            string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	ReservationType string `protobuf:"bytes,6,opt,name=reservationType,proto3" json:"reservationType,omitempty"`
	PlaceType       string `protobuf:"bytes,7,opt,name=placeType,proto3" json:"placeType,omitempty"`
	PricingType     string `protobuf:"bytes,8,opt,name=pricingType,proto3" json:"pricingType,omitempty"`
	MinNight        int32  `protobuf:"varint,9,opt,name=minNight,proto3" json:"minNight,omitempty"`
}

func (x *ResponseDetails) Reset() {
	*x = ResponseDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acm_acm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDetails) ProtoMessage() {}

func (x *ResponseDetails) ProtoReflect() protoreflect.Message {
	mi := &file_acm_acm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseDetails.ProtoReflect.Descriptor instead.
func (*ResponseDetails) Descriptor() ([]byte, []int) {
	return file_acm_acm_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResponseDetails) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ResponseDetails) GetHost() int32 {
	if x != nil {
		return x.Host
	}
	return 0
}

func (x *ResponseDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ResponseDetails) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ResponseDetails) GetReservationType() string {
	if x != nil {
		return x.ReservationType
	}
	return ""
}

func (x *ResponseDetails) GetPlaceType() string {
	if x != nil {
		return x.PlaceType
	}
	return ""
}

func (x *ResponseDetails) GetPricingType() string {
	if x != nil {
		return x.PricingType
	}
	return ""
}

func (x *ResponseDetails) GetMinNight() int32 {
	if x != nil {
		return x.MinNight
	}
	return 0
}

type RequestCreate_Capacity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"beds",validate:"required"
	Beds *RequestCreate_Capacity_Beds `protobuf:"bytes,1,opt,name=beds,proto3" json:"beds" validate:"required"`
	// @inject_tag: json:"guests",validate:"required"
	Guests *RequestCreate_Capacity_Guests `protobuf:"bytes,2,opt,name=guests,proto3" json:"guests" validate:"required"`
}

func (x *RequestCreate_Capacity) Reset() {
	*x = RequestCreate_Capacity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acm_acm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCreate_Capacity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCreate_Capacity) ProtoMessage() {}

func (x *RequestCreate_Capacity) ProtoReflect() protoreflect.Message {
	mi := &file_acm_acm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCreate_Capacity.ProtoReflect.Descriptor instead.
func (*RequestCreate_Capacity) Descriptor() ([]byte, []int) {
	return file_acm_acm_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RequestCreate_Capacity) GetBeds() *RequestCreate_Capacity_Beds {
	if x != nil {
		return x.Beds
	}
	return nil
}

func (x *RequestCreate_Capacity) GetGuests() *RequestCreate_Capacity_Guests {
	if x != nil {
		return x.Guests
	}
	return nil
}

type RequestCreate_Capacity_Beds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"twin",validate:"number"
	Twin int32 `protobuf:"varint,1,opt,name=twin,proto3" json:"twin" validate:"number"`
	// @inject_tag: json:"single",validate:"number"
	Single int32 `protobuf:"varint,2,opt,name=single,proto3" json:"single" validate:"number"`
	// @inject_tag: json:"double",validate:"number"
	Double int32 `protobuf:"varint,3,opt,name=double,proto3" json:"double" validate:"number"`
	// @inject_tag: json:"mattress",validate:"number"
	Mattress int32 `protobuf:"varint,4,opt,name=mattress,proto3" json:"mattress" validate:"number"`
}

func (x *RequestCreate_Capacity_Beds) Reset() {
	*x = RequestCreate_Capacity_Beds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acm_acm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCreate_Capacity_Beds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCreate_Capacity_Beds) ProtoMessage() {}

func (x *RequestCreate_Capacity_Beds) ProtoReflect() protoreflect.Message {
	mi := &file_acm_acm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCreate_Capacity_Beds.ProtoReflect.Descriptor instead.
func (*RequestCreate_Capacity_Beds) Descriptor() ([]byte, []int) {
	return file_acm_acm_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *RequestCreate_Capacity_Beds) GetTwin() int32 {
	if x != nil {
		return x.Twin
	}
	return 0
}

func (x *RequestCreate_Capacity_Beds) GetSingle() int32 {
	if x != nil {
		return x.Single
	}
	return 0
}

func (x *RequestCreate_Capacity_Beds) GetDouble() int32 {
	if x != nil {
		return x.Double
	}
	return 0
}

func (x *RequestCreate_Capacity_Beds) GetMattress() int32 {
	if x != nil {
		return x.Mattress
	}
	return 0
}

type RequestCreate_Capacity_Guests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"base",validate:"required|number|min=1"
	Base int32 `protobuf:"varint,1,opt,name=base,proto3" json:"base" validate:"required|number|min=1"`
	// @inject_tag: json:"extra",validate:"number"
	Extra int32 `protobuf:"varint,2,opt,name=extra,proto3" json:"extra" validate:"number"`
}

func (x *RequestCreate_Capacity_Guests) Reset() {
	*x = RequestCreate_Capacity_Guests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acm_acm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCreate_Capacity_Guests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCreate_Capacity_Guests) ProtoMessage() {}

func (x *RequestCreate_Capacity_Guests) ProtoReflect() protoreflect.Message {
	mi := &file_acm_acm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCreate_Capacity_Guests.ProtoReflect.Descriptor instead.
func (*RequestCreate_Capacity_Guests) Descriptor() ([]byte, []int) {
	return file_acm_acm_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *RequestCreate_Capacity_Guests) GetBase() int32 {
	if x != nil {
		return x.Base
	}
	return 0
}

func (x *RequestCreate_Capacity_Guests) GetExtra() int32 {
	if x != nil {
		return x.Extra
	}
	return 0
}

var File_acm_acm_proto protoreflect.FileDescriptor

var file_acm_acm_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x63, 0x6d, 0x2f, 0x61, 0x63, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x63, 0x6d, 0x22, 0xc9, 0x04, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x4e, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x4e, 0x69, 0x67, 0x68, 0x74, 0x12, 0x37, 0x0a, 0x08,
	0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x08, 0x63, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x1a, 0x98, 0x02, 0x0a, 0x08, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x34, 0x0a, 0x04, 0x62, 0x65, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x42, 0x65,
	0x64, 0x73, 0x52, 0x04, 0x62, 0x65, 0x64, 0x73, 0x12, 0x3a, 0x0a, 0x06, 0x67, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x06, 0x67, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x1a, 0x66, 0x0a, 0x04, 0x42, 0x65, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x77, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x77, 0x69, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x74, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x74, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x32, 0x0a, 0x06,
	0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x22, 0x20, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x87, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x4e, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x4e, 0x69, 0x67, 0x68, 0x74, 0x32, 0x7a, 0x0a, 0x0a,
	0x41, 0x63, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x14, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00,
	0x12, 0x36, 0x0a, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x13, 0x2e, 0x61, 0x63,
	0x6d, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x1a, 0x14, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_acm_acm_proto_rawDescOnce sync.Once
	file_acm_acm_proto_rawDescData = file_acm_acm_proto_rawDesc
)

func file_acm_acm_proto_rawDescGZIP() []byte {
	file_acm_acm_proto_rawDescOnce.Do(func() {
		file_acm_acm_proto_rawDescData = protoimpl.X.CompressGZIP(file_acm_acm_proto_rawDescData)
	})
	return file_acm_acm_proto_rawDescData
}

var file_acm_acm_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_acm_acm_proto_goTypes = []interface{}{
	(*RequestCreate)(nil),                 // 0: acm.RequestCreate
	(*RequestDetails)(nil),                // 1: acm.RequestDetails
	(*ResponseDetails)(nil),               // 2: acm.ResponseDetails
	(*RequestCreate_Capacity)(nil),        // 3: acm.RequestCreate.Capacity
	(*RequestCreate_Capacity_Beds)(nil),   // 4: acm.RequestCreate.Capacity.Beds
	(*RequestCreate_Capacity_Guests)(nil), // 5: acm.RequestCreate.Capacity.Guests
}
var file_acm_acm_proto_depIdxs = []int32{
	3, // 0: acm.RequestCreate.capacity:type_name -> acm.RequestCreate.Capacity
	4, // 1: acm.RequestCreate.Capacity.beds:type_name -> acm.RequestCreate.Capacity.Beds
	5, // 2: acm.RequestCreate.Capacity.guests:type_name -> acm.RequestCreate.Capacity.Guests
	0, // 3: acm.AcmService.Create:input_type -> acm.RequestCreate
	1, // 4: acm.AcmService.Details:input_type -> acm.RequestDetails
	2, // 5: acm.AcmService.Create:output_type -> acm.ResponseDetails
	2, // 6: acm.AcmService.Details:output_type -> acm.ResponseDetails
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_acm_acm_proto_init() }
func file_acm_acm_proto_init() {
	if File_acm_acm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_acm_acm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCreate); i {
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
		file_acm_acm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDetails); i {
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
		file_acm_acm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseDetails); i {
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
		file_acm_acm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCreate_Capacity); i {
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
		file_acm_acm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCreate_Capacity_Beds); i {
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
		file_acm_acm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCreate_Capacity_Guests); i {
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
			RawDescriptor: file_acm_acm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_acm_acm_proto_goTypes,
		DependencyIndexes: file_acm_acm_proto_depIdxs,
		MessageInfos:      file_acm_acm_proto_msgTypes,
	}.Build()
	File_acm_acm_proto = out.File
	file_acm_acm_proto_rawDesc = nil
	file_acm_acm_proto_goTypes = nil
	file_acm_acm_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AcmServiceClient is the client API for AcmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AcmServiceClient interface {
	Create(ctx context.Context, in *RequestCreate, opts ...grpc.CallOption) (*ResponseDetails, error)
	Details(ctx context.Context, in *RequestDetails, opts ...grpc.CallOption) (*ResponseDetails, error)
}

type acmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAcmServiceClient(cc grpc.ClientConnInterface) AcmServiceClient {
	return &acmServiceClient{cc}
}

func (c *acmServiceClient) Create(ctx context.Context, in *RequestCreate, opts ...grpc.CallOption) (*ResponseDetails, error) {
	out := new(ResponseDetails)
	err := c.cc.Invoke(ctx, "/acm.AcmService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *acmServiceClient) Details(ctx context.Context, in *RequestDetails, opts ...grpc.CallOption) (*ResponseDetails, error) {
	out := new(ResponseDetails)
	err := c.cc.Invoke(ctx, "/acm.AcmService/Details", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AcmServiceServer is the server API for AcmService service.
type AcmServiceServer interface {
	Create(context.Context, *RequestCreate) (*ResponseDetails, error)
	Details(context.Context, *RequestDetails) (*ResponseDetails, error)
}

// UnimplementedAcmServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAcmServiceServer struct {
}

func (*UnimplementedAcmServiceServer) Create(context.Context, *RequestCreate) (*ResponseDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAcmServiceServer) Details(context.Context, *RequestDetails) (*ResponseDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Details not implemented")
}

func RegisterAcmServiceServer(s *grpc.Server, srv AcmServiceServer) {
	s.RegisterService(&_AcmService_serviceDesc, srv)
}

func _AcmService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcmServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/acm.AcmService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcmServiceServer).Create(ctx, req.(*RequestCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _AcmService_Details_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcmServiceServer).Details(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/acm.AcmService/Details",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcmServiceServer).Details(ctx, req.(*RequestDetails))
	}
	return interceptor(ctx, in, info, handler)
}

var _AcmService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "acm.AcmService",
	HandlerType: (*AcmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AcmService_Create_Handler,
		},
		{
			MethodName: "Details",
			Handler:    _AcmService_Details_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "acm/acm.proto",
}
