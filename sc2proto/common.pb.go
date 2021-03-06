// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: common.proto

package sc2proto

import (
	"reflect"
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Race int32

const (
	Race_NoRace  Race = 0
	Race_Terran  Race = 1
	Race_Zerg    Race = 2
	Race_Protoss Race = 3
	Race_Random  Race = 4
)

// Enum value maps for Race.
var (
	Race_name = map[int32]string{
		0: "NoRace",
		1: "Terran",
		2: "Zerg",
		3: "Protoss",
		4: "Random",
	}
	Race_value = map[string]int32{
		"NoRace":  0,
		"Terran":  1,
		"Zerg":    2,
		"Protoss": 3,
		"Random":  4,
	}
)

func (x Race) Enum() *Race {
	p := new(Race)
	*p = x
	return p
}

func (x Race) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Race) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (Race) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x Race) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Race) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Race(num)
	return nil
}

// Deprecated: Use Race.Descriptor instead.
func (Race) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type AvailableAbility struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbilityId     *int32 `protobuf:"varint,1,opt,name=ability_id,json=abilityId" json:"ability_id,omitempty"`
	RequiresPoint *bool  `protobuf:"varint,2,opt,name=requires_point,json=requiresPoint" json:"requires_point,omitempty"`
}

func (x *AvailableAbility) Reset() {
	*x = AvailableAbility{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableAbility) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableAbility) ProtoMessage() {}

func (x *AvailableAbility) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableAbility.ProtoReflect.Descriptor instead.
func (*AvailableAbility) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *AvailableAbility) GetAbilityId() int32 {
	if x != nil && x.AbilityId != nil {
		return *x.AbilityId
	}
	return 0
}

func (x *AvailableAbility) GetRequiresPoint() bool {
	if x != nil && x.RequiresPoint != nil {
		return *x.RequiresPoint
	}
	return false
}

type ImageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BitsPerPixel *int32   `protobuf:"varint,1,opt,name=bits_per_pixel,json=bitsPerPixel" json:"bits_per_pixel,omitempty"` // Number of bits per pixel; 8 bits for a byte etc.
	Size         *Size2DI `protobuf:"bytes,2,opt,name=size" json:"size,omitempty"`                                        // Dimension in pixels.
	Data         []byte   `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`                                        // Binary data; the size of this buffer in bytes is width * height * bits_per_pixel / 8.
}

func (x *ImageData) Reset() {
	*x = ImageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageData) ProtoMessage() {}

func (x *ImageData) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageData.ProtoReflect.Descriptor instead.
func (*ImageData) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *ImageData) GetBitsPerPixel() int32 {
	if x != nil && x.BitsPerPixel != nil {
		return *x.BitsPerPixel
	}
	return 0
}

func (x *ImageData) GetSize() *Size2DI {
	if x != nil {
		return x.Size
	}
	return nil
}

func (x *ImageData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Point on the screen/minimap (e.g., 0..64).
// Note: bottom left of the screen is 0, 0.
type PointI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X *int32 `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y *int32 `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
}

func (x *PointI) Reset() {
	*x = PointI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PointI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointI) ProtoMessage() {}

func (x *PointI) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointI.ProtoReflect.Descriptor instead.
func (*PointI) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *PointI) GetX() int32 {
	if x != nil && x.X != nil {
		return *x.X
	}
	return 0
}

func (x *PointI) GetY() int32 {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return 0
}

// Screen space rectangular area.
type RectangleI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P0 *PointI `protobuf:"bytes,1,opt,name=p0" json:"p0,omitempty"`
	P1 *PointI `protobuf:"bytes,2,opt,name=p1" json:"p1,omitempty"`
}

func (x *RectangleI) Reset() {
	*x = RectangleI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RectangleI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RectangleI) ProtoMessage() {}

func (x *RectangleI) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RectangleI.ProtoReflect.Descriptor instead.
func (*RectangleI) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *RectangleI) GetP0() *PointI {
	if x != nil {
		return x.P0
	}
	return nil
}

func (x *RectangleI) GetP1() *PointI {
	if x != nil {
		return x.P1
	}
	return nil
}

// Point on the game board, 0..255.
// Note: bottom left of the screen is 0, 0.
type Point2D struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X *float32 `protobuf:"fixed32,1,opt,name=x" json:"x,omitempty"`
	Y *float32 `protobuf:"fixed32,2,opt,name=y" json:"y,omitempty"`
}

func (x *Point2D) Reset() {
	*x = Point2D{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point2D) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point2D) ProtoMessage() {}

func (x *Point2D) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point2D.ProtoReflect.Descriptor instead.
func (*Point2D) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *Point2D) GetX() float32 {
	if x != nil && x.X != nil {
		return *x.X
	}
	return 0
}

func (x *Point2D) GetY() float32 {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return 0
}

// Point on the game board, 0..255.
// Note: bottom left of the screen is 0, 0.
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X *float32 `protobuf:"fixed32,1,opt,name=x" json:"x,omitempty"`
	Y *float32 `protobuf:"fixed32,2,opt,name=y" json:"y,omitempty"`
	Z *float32 `protobuf:"fixed32,3,opt,name=z" json:"z,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *Point) GetX() float32 {
	if x != nil && x.X != nil {
		return *x.X
	}
	return 0
}

func (x *Point) GetY() float32 {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return 0
}

func (x *Point) GetZ() float32 {
	if x != nil && x.Z != nil {
		return *x.Z
	}
	return 0
}

// Screen dimensions.
type Size2DI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X *int32 `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y *int32 `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
}

func (x *Size2DI) Reset() {
	*x = Size2DI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Size2DI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Size2DI) ProtoMessage() {}

func (x *Size2DI) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Size2DI.ProtoReflect.Descriptor instead.
func (*Size2DI) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{6}
}

func (x *Size2DI) GetX() int32 {
	if x != nil && x.X != nil {
		return *x.X
	}
	return 0
}

func (x *Size2DI) GetY() int32 {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return 0
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x63, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x10, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x22, 0x6c, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x24, 0x0a, 0x0e, 0x62, 0x69, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x69, 0x78, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x62, 0x69, 0x74, 0x73, 0x50, 0x65, 0x72,
	0x50, 0x69, 0x78, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x63, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x69, 0x7a, 0x65, 0x32, 0x44, 0x49, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x24, 0x0a, 0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x50, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x74, 0x61, 0x6e,
	0x67, 0x6c, 0x65, 0x49, 0x12, 0x20, 0x0a, 0x02, 0x70, 0x30, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x73, 0x63, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x49, 0x52, 0x02, 0x70, 0x30, 0x12, 0x20, 0x0a, 0x02, 0x70, 0x31, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x63, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x49, 0x52, 0x02, 0x70, 0x31, 0x22, 0x25, 0x0a, 0x07, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x32, 0x44, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x22,
	0x31, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x01, 0x7a, 0x22, 0x25, 0x0a, 0x07, 0x53, 0x69, 0x7a, 0x65, 0x32, 0x44, 0x49, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x2a, 0x41, 0x0a, 0x04, 0x52, 0x61, 0x63,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x6f, 0x52, 0x61, 0x63, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x54, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x65, 0x72,
	0x67, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x73, 0x10, 0x03,
	0x12, 0x0a, 0x0a, 0x06, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x10, 0x04, 0x42, 0x2c, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x69, 0x6e, 0x77, 0x75,
	0x7a, 0x68, 0x61, 0x6f, 0x2f, 0x73, 0x63, 0x32, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x67,
	0x6f, 0x2f, 0x73, 0x63, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_common_proto_goTypes = []interface{}{
	(Race)(0),                // 0: sc2proto.Race
	(*AvailableAbility)(nil), // 1: sc2proto.AvailableAbility
	(*ImageData)(nil),        // 2: sc2proto.ImageData
	(*PointI)(nil),           // 3: sc2proto.PointI
	(*RectangleI)(nil),       // 4: sc2proto.RectangleI
	(*Point2D)(nil),          // 5: sc2proto.Point2D
	(*Point)(nil),            // 6: sc2proto.Point
	(*Size2DI)(nil),          // 7: sc2proto.Size2DI
}
var file_common_proto_depIdxs = []int32{
	7, // 0: sc2proto.ImageData.size:type_name -> sc2proto.Size2DI
	3, // 1: sc2proto.RectangleI.p0:type_name -> sc2proto.PointI
	3, // 2: sc2proto.RectangleI.p1:type_name -> sc2proto.PointI
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableAbility); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageData); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PointI); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RectangleI); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point2D); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Size2DI); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
