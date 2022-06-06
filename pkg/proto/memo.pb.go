// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: memo.proto

package proto

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

type Memosphere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    map[string]*Memos `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Modules map[string]*Thunk `protobuf:"bytes,2,rep,name=modules,proto3" json:"modules,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Memosphere) Reset() {
	*x = Memosphere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memosphere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memosphere) ProtoMessage() {}

func (x *Memosphere) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memosphere.ProtoReflect.Descriptor instead.
func (*Memosphere) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{0}
}

func (x *Memosphere) GetData() map[string]*Memos {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Memosphere) GetModules() map[string]*Thunk {
	if x != nil {
		return x.Modules
	}
	return nil
}

type Memos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memos []*Memory `protobuf:"bytes,1,rep,name=memos,proto3" json:"memos,omitempty"`
}

func (x *Memos) Reset() {
	*x = Memos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memos) ProtoMessage() {}

func (x *Memos) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memos.ProtoReflect.Descriptor instead.
func (*Memos) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{1}
}

func (x *Memos) GetMemos() []*Memory {
	if x != nil {
		return x.Memos
	}
	return nil
}

type Memory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input  *Value `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Output *Value `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *Memory) Reset() {
	*x = Memory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memory) ProtoMessage() {}

func (x *Memory) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memory.ProtoReflect.Descriptor instead.
func (*Memory) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{2}
}

func (x *Memory) GetInput() *Value {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *Memory) GetOutput() *Value {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_memo_proto protoreflect.FileDescriptor

var file_memo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61,
	0x73, 0x73, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84,
	0x02, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x12, 0x2e, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x61,
	0x73, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a,
	0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65,
	0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0x44, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x6f,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x47, 0x0a, 0x0c,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x62, 0x61, 0x73, 0x73, 0x2e, 0x54, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2b, 0x0a, 0x05, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x12, 0x22,
	0x0a, 0x05, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x62, 0x61, 0x73, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x05, 0x6d, 0x65, 0x6d,
	0x6f, 0x73, 0x22, 0x50, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x61,
	0x73, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x23, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x62, 0x61, 0x73, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_memo_proto_rawDescOnce sync.Once
	file_memo_proto_rawDescData = file_memo_proto_rawDesc
)

func file_memo_proto_rawDescGZIP() []byte {
	file_memo_proto_rawDescOnce.Do(func() {
		file_memo_proto_rawDescData = protoimpl.X.CompressGZIP(file_memo_proto_rawDescData)
	})
	return file_memo_proto_rawDescData
}

var file_memo_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_memo_proto_goTypes = []interface{}{
	(*Memosphere)(nil), // 0: bass.Memosphere
	(*Memos)(nil),      // 1: bass.Memos
	(*Memory)(nil),     // 2: bass.Memory
	nil,                // 3: bass.Memosphere.DataEntry
	nil,                // 4: bass.Memosphere.ModulesEntry
	(*Value)(nil),      // 5: bass.Value
	(*Thunk)(nil),      // 6: bass.Thunk
}
var file_memo_proto_depIdxs = []int32{
	3, // 0: bass.Memosphere.data:type_name -> bass.Memosphere.DataEntry
	4, // 1: bass.Memosphere.modules:type_name -> bass.Memosphere.ModulesEntry
	2, // 2: bass.Memos.memos:type_name -> bass.Memory
	5, // 3: bass.Memory.input:type_name -> bass.Value
	5, // 4: bass.Memory.output:type_name -> bass.Value
	1, // 5: bass.Memosphere.DataEntry.value:type_name -> bass.Memos
	6, // 6: bass.Memosphere.ModulesEntry.value:type_name -> bass.Thunk
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_memo_proto_init() }
func file_memo_proto_init() {
	if File_memo_proto != nil {
		return
	}
	file_bass_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_memo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memosphere); i {
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
		file_memo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memos); i {
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
		file_memo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memory); i {
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
			RawDescriptor: file_memo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_memo_proto_goTypes,
		DependencyIndexes: file_memo_proto_depIdxs,
		MessageInfos:      file_memo_proto_msgTypes,
	}.Build()
	File_memo_proto = out.File
	file_memo_proto_rawDesc = nil
	file_memo_proto_goTypes = nil
	file_memo_proto_depIdxs = nil
}
