// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/msg.proto

package go_libuiohook

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

type KeyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyCode uint32 `protobuf:"varint,2,opt,name=key_code,json=keyCode,proto3" json:"key_code,omitempty"`
	RawCode uint32 `protobuf:"varint,3,opt,name=raw_code,json=rawCode,proto3" json:"raw_code,omitempty"`
	KeyChar uint32 `protobuf:"varint,4,opt,name=key_char,json=keyChar,proto3" json:"key_char,omitempty"`
}

func (x *KeyData) Reset() {
	*x = KeyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyData) ProtoMessage() {}

func (x *KeyData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyData.ProtoReflect.Descriptor instead.
func (*KeyData) Descriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{0}
}

func (x *KeyData) GetKeyCode() uint32 {
	if x != nil {
		return x.KeyCode
	}
	return 0
}

func (x *KeyData) GetRawCode() uint32 {
	if x != nil {
		return x.RawCode
	}
	return 0
}

func (x *KeyData) GetKeyChar() uint32 {
	if x != nil {
		return x.KeyChar
	}
	return 0
}

type MouseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Button uint32 `protobuf:"varint,2,opt,name=button,proto3" json:"button,omitempty"`
	Clicks uint32 `protobuf:"varint,3,opt,name=clicks,proto3" json:"clicks,omitempty"`
	X      int32  `protobuf:"varint,4,opt,name=x,proto3" json:"x,omitempty"`
	Y      int32  `protobuf:"varint,5,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *MouseData) Reset() {
	*x = MouseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MouseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MouseData) ProtoMessage() {}

func (x *MouseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MouseData.ProtoReflect.Descriptor instead.
func (*MouseData) Descriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{1}
}

func (x *MouseData) GetButton() uint32 {
	if x != nil {
		return x.Button
	}
	return 0
}

func (x *MouseData) GetClicks() uint32 {
	if x != nil {
		return x.Clicks
	}
	return 0
}

func (x *MouseData) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *MouseData) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type WheelData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      uint32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Clicks    uint32 `protobuf:"varint,2,opt,name=clicks,proto3" json:"clicks,omitempty"`
	X         int32  `protobuf:"varint,3,opt,name=x,proto3" json:"x,omitempty"`
	Y         int32  `protobuf:"varint,4,opt,name=y,proto3" json:"y,omitempty"`
	Amount    uint32 `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Rotation  int32  `protobuf:"varint,6,opt,name=rotation,proto3" json:"rotation,omitempty"`
	Direction uint32 `protobuf:"varint,7,opt,name=direction,proto3" json:"direction,omitempty"`
}

func (x *WheelData) Reset() {
	*x = WheelData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WheelData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WheelData) ProtoMessage() {}

func (x *WheelData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WheelData.ProtoReflect.Descriptor instead.
func (*WheelData) Descriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{2}
}

func (x *WheelData) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *WheelData) GetClicks() uint32 {
	if x != nil {
		return x.Clicks
	}
	return 0
}

func (x *WheelData) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *WheelData) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *WheelData) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *WheelData) GetRotation() int32 {
	if x != nil {
		return x.Rotation
	}
	return 0
}

func (x *WheelData) GetDirection() uint32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     uint32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Time     uint64 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Mask     uint32 `protobuf:"varint,3,opt,name=mask,proto3" json:"mask,omitempty"`
	Reserved uint32 `protobuf:"varint,4,opt,name=reserved,proto3" json:"reserved,omitempty"`
	// Types that are assignable to Data:
	//
	//	*Msg_Key
	//	*Msg_Mouse
	//	*Msg_Wheel
	Data isMsg_Data `protobuf_oneof:"data"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{3}
}

func (x *Msg) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Msg) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Msg) GetMask() uint32 {
	if x != nil {
		return x.Mask
	}
	return 0
}

func (x *Msg) GetReserved() uint32 {
	if x != nil {
		return x.Reserved
	}
	return 0
}

func (m *Msg) GetData() isMsg_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Msg) GetKey() *KeyData {
	if x, ok := x.GetData().(*Msg_Key); ok {
		return x.Key
	}
	return nil
}

func (x *Msg) GetMouse() *MouseData {
	if x, ok := x.GetData().(*Msg_Mouse); ok {
		return x.Mouse
	}
	return nil
}

func (x *Msg) GetWheel() *WheelData {
	if x, ok := x.GetData().(*Msg_Wheel); ok {
		return x.Wheel
	}
	return nil
}

type isMsg_Data interface {
	isMsg_Data()
}

type Msg_Key struct {
	Key *KeyData `protobuf:"bytes,5,opt,name=key,proto3,oneof"`
}

type Msg_Mouse struct {
	Mouse *MouseData `protobuf:"bytes,6,opt,name=mouse,proto3,oneof"`
}

type Msg_Wheel struct {
	Wheel *WheelData `protobuf:"bytes,7,opt,name=wheel,proto3,oneof"`
}

func (*Msg_Key) isMsg_Data() {}

func (*Msg_Mouse) isMsg_Data() {}

func (*Msg_Wheel) isMsg_Data() {}

var File_proto_msg_proto protoreflect.FileDescriptor

var file_proto_msg_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x75, 0x69, 0x6f, 0x68, 0x6f, 0x6f, 0x6b, 0x22, 0x5a, 0x0a, 0x07, 0x4b, 0x65,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6b,
	0x65, 0x79, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x43, 0x68, 0x61, 0x72, 0x22, 0x57, 0x0a, 0x09, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6c, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x6c, 0x69,
	0x63, 0x6b, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22,
	0xa5, 0x01, 0x0a, 0x09, 0x57, 0x68, 0x65, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe3, 0x01, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x69, 0x6f, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x4b,
	0x65, 0x79, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a,
	0x05, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x75,
	0x69, 0x6f, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x48, 0x00, 0x52, 0x05, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x77, 0x68, 0x65,
	0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x75, 0x69, 0x6f, 0x68, 0x6f,
	0x6f, 0x6b, 0x2e, 0x57, 0x68, 0x65, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x05,
	0x77, 0x68, 0x65, 0x65, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x72, 0x6f, 0x6e,
	0x70, 0x61, 0x72, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x69, 0x62, 0x75, 0x69, 0x6f, 0x68, 0x6f,
	0x6f, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_msg_proto_rawDescOnce sync.Once
	file_proto_msg_proto_rawDescData = file_proto_msg_proto_rawDesc
)

func file_proto_msg_proto_rawDescGZIP() []byte {
	file_proto_msg_proto_rawDescOnce.Do(func() {
		file_proto_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_msg_proto_rawDescData)
	})
	return file_proto_msg_proto_rawDescData
}

var file_proto_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_msg_proto_goTypes = []interface{}{
	(*KeyData)(nil),   // 0: uiohook.KeyData
	(*MouseData)(nil), // 1: uiohook.MouseData
	(*WheelData)(nil), // 2: uiohook.WheelData
	(*Msg)(nil),       // 3: uiohook.Msg
}
var file_proto_msg_proto_depIdxs = []int32{
	0, // 0: uiohook.Msg.key:type_name -> uiohook.KeyData
	1, // 1: uiohook.Msg.mouse:type_name -> uiohook.MouseData
	2, // 2: uiohook.Msg.wheel:type_name -> uiohook.WheelData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_msg_proto_init() }
func file_proto_msg_proto_init() {
	if File_proto_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyData); i {
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
		file_proto_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MouseData); i {
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
		file_proto_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WheelData); i {
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
		file_proto_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
	file_proto_msg_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Msg_Key)(nil),
		(*Msg_Mouse)(nil),
		(*Msg_Wheel)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_msg_proto_goTypes,
		DependencyIndexes: file_proto_msg_proto_depIdxs,
		MessageInfos:      file_proto_msg_proto_msgTypes,
	}.Build()
	File_proto_msg_proto = out.File
	file_proto_msg_proto_rawDesc = nil
	file_proto_msg_proto_goTypes = nil
	file_proto_msg_proto_depIdxs = nil
}
