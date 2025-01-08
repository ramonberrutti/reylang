// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: reylang/bytecode.proto

package reylang

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bytecode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instructions []*Instruction `protobuf:"bytes,1,rep,name=instructions,proto3" json:"instructions,omitempty"`
}

func (x *Bytecode) Reset() {
	*x = Bytecode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reylang_bytecode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bytecode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bytecode) ProtoMessage() {}

func (x *Bytecode) ProtoReflect() protoreflect.Message {
	mi := &file_reylang_bytecode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bytecode.ProtoReflect.Descriptor instead.
func (*Bytecode) Descriptor() ([]byte, []int) {
	return file_reylang_bytecode_proto_rawDescGZIP(), []int{0}
}

func (x *Bytecode) GetInstructions() []*Instruction {
	if x != nil {
		return x.Instructions
	}
	return nil
}

type Instruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Opcode OpCode `protobuf:"varint,1,opt,name=opcode,proto3,enum=ramonberrutti.reylang.OpCode" json:"opcode,omitempty"`
	// Types that are assignable to Operand:
	//
	//	*Instruction_IntegerOperand
	//	*Instruction_StringOperand
	//	*Instruction_StructOperand
	Operand isInstruction_Operand `protobuf_oneof:"operand"`
}

func (x *Instruction) Reset() {
	*x = Instruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reylang_bytecode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instruction) ProtoMessage() {}

func (x *Instruction) ProtoReflect() protoreflect.Message {
	mi := &file_reylang_bytecode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instruction.ProtoReflect.Descriptor instead.
func (*Instruction) Descriptor() ([]byte, []int) {
	return file_reylang_bytecode_proto_rawDescGZIP(), []int{1}
}

func (x *Instruction) GetOpcode() OpCode {
	if x != nil {
		return x.Opcode
	}
	return OpCode_NONE
}

func (m *Instruction) GetOperand() isInstruction_Operand {
	if m != nil {
		return m.Operand
	}
	return nil
}

func (x *Instruction) GetIntegerOperand() int32 {
	if x, ok := x.GetOperand().(*Instruction_IntegerOperand); ok {
		return x.IntegerOperand
	}
	return 0
}

func (x *Instruction) GetStringOperand() string {
	if x, ok := x.GetOperand().(*Instruction_StringOperand); ok {
		return x.StringOperand
	}
	return ""
}

func (x *Instruction) GetStructOperand() *structpb.ListValue {
	if x, ok := x.GetOperand().(*Instruction_StructOperand); ok {
		return x.StructOperand
	}
	return nil
}

type isInstruction_Operand interface {
	isInstruction_Operand()
}

type Instruction_IntegerOperand struct {
	IntegerOperand int32 `protobuf:"varint,3,opt,name=integer_operand,json=integerOperand,proto3,oneof"`
}

type Instruction_StringOperand struct {
	StringOperand string `protobuf:"bytes,4,opt,name=string_operand,json=stringOperand,proto3,oneof"`
}

type Instruction_StructOperand struct {
	StructOperand *structpb.ListValue `protobuf:"bytes,5,opt,name=struct_operand,json=structOperand,proto3,oneof"`
}

func (*Instruction_IntegerOperand) isInstruction_Operand() {}

func (*Instruction_StringOperand) isInstruction_Operand() {}

func (*Instruction_StructOperand) isInstruction_Operand() {}

var File_reylang_bytecode_proto protoreflect.FileDescriptor

var file_reylang_bytecode_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x62, 0x79, 0x74, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62,
	0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x1a,
	0x14, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x08, 0x42, 0x79, 0x74, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x46, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72,
	0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xe8, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x06, 0x6f, 0x70, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62,
	0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x4f, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x29,
	0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x27, 0x0a, 0x0e, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x6e, 0x64, 0x12, 0x43, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x5f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x6e, 0x64, 0x42, 0xd2, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e,
	0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67,
	0x42, 0x0d, 0x42, 0x79, 0x74, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61,
	0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2f, 0x72, 0x65, 0x79, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x79,
	0x6c, 0x61, 0x6e, 0x67, 0xa2, 0x02, 0x03, 0x52, 0x52, 0x58, 0xaa, 0x02, 0x15, 0x52, 0x61, 0x6d,
	0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x79, 0x6c, 0x61,
	0x6e, 0x67, 0xca, 0x02, 0x15, 0x52, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74,
	0x74, 0x69, 0x5c, 0x52, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0xe2, 0x02, 0x21, 0x52, 0x61, 0x6d,
	0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x5c, 0x52, 0x65, 0x79, 0x6c, 0x61,
	0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x16, 0x52, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x3a, 0x3a,
	0x52, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reylang_bytecode_proto_rawDescOnce sync.Once
	file_reylang_bytecode_proto_rawDescData = file_reylang_bytecode_proto_rawDesc
)

func file_reylang_bytecode_proto_rawDescGZIP() []byte {
	file_reylang_bytecode_proto_rawDescOnce.Do(func() {
		file_reylang_bytecode_proto_rawDescData = protoimpl.X.CompressGZIP(file_reylang_bytecode_proto_rawDescData)
	})
	return file_reylang_bytecode_proto_rawDescData
}

var file_reylang_bytecode_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_reylang_bytecode_proto_goTypes = []interface{}{
	(*Bytecode)(nil),           // 0: ramonberrutti.reylang.Bytecode
	(*Instruction)(nil),        // 1: ramonberrutti.reylang.Instruction
	(OpCode)(0),                // 2: ramonberrutti.reylang.OpCode
	(*structpb.ListValue)(nil), // 3: google.protobuf.ListValue
}
var file_reylang_bytecode_proto_depIdxs = []int32{
	1, // 0: ramonberrutti.reylang.Bytecode.instructions:type_name -> ramonberrutti.reylang.Instruction
	2, // 1: ramonberrutti.reylang.Instruction.opcode:type_name -> ramonberrutti.reylang.OpCode
	3, // 2: ramonberrutti.reylang.Instruction.struct_operand:type_name -> google.protobuf.ListValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_reylang_bytecode_proto_init() }
func file_reylang_bytecode_proto_init() {
	if File_reylang_bytecode_proto != nil {
		return
	}
	file_reylang_opcode_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_reylang_bytecode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bytecode); i {
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
		file_reylang_bytecode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instruction); i {
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
	file_reylang_bytecode_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Instruction_IntegerOperand)(nil),
		(*Instruction_StringOperand)(nil),
		(*Instruction_StructOperand)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reylang_bytecode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reylang_bytecode_proto_goTypes,
		DependencyIndexes: file_reylang_bytecode_proto_depIdxs,
		MessageInfos:      file_reylang_bytecode_proto_msgTypes,
	}.Build()
	File_reylang_bytecode_proto = out.File
	file_reylang_bytecode_proto_rawDesc = nil
	file_reylang_bytecode_proto_goTypes = nil
	file_reylang_bytecode_proto_depIdxs = nil
}