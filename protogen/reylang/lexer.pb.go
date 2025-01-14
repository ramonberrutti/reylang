// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: reylang/lexer.proto

package reylang

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

type TokenType int32

const (
	TokenType_UNKNOWN     TokenType = 0 // unknown token
	TokenType_KEYWORD     TokenType = 1 // if, else, for, while, return, etc.
	TokenType_IDENTIFIER  TokenType = 2 // variable names, function names, etc.
	TokenType_STRING      TokenType = 3 // "string"
	TokenType_INTEGER     TokenType = 4 // 1
	TokenType_FLOAT       TokenType = 5 // 1.0
	TokenType_BOOLEAN     TokenType = 6 // true or false
	TokenType_OPERATOR    TokenType = 7 // + - * / % == != < <= > >= && || ! & | ^ << >> ~
	TokenType_PUNCTUATION TokenType = 8
	TokenType_LPAREN      TokenType = 9  // (
	TokenType_RPAREN      TokenType = 10 // )
	TokenType_LBRACE      TokenType = 11 // {
	TokenType_RBRACE      TokenType = 12 // }
	TokenType_LBRACKET    TokenType = 13 // [
	TokenType_RBRACKET    TokenType = 14 // ]
	TokenType_COMMA       TokenType = 15 // ,
	TokenType_DECLARATION TokenType = 16 // :=
	TokenType_ASSIGNMENT  TokenType = 17 // =
	TokenType_SEMICOLON   TokenType = 18 // ;
)

// Enum value maps for TokenType.
var (
	TokenType_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "KEYWORD",
		2:  "IDENTIFIER",
		3:  "STRING",
		4:  "INTEGER",
		5:  "FLOAT",
		6:  "BOOLEAN",
		7:  "OPERATOR",
		8:  "PUNCTUATION",
		9:  "LPAREN",
		10: "RPAREN",
		11: "LBRACE",
		12: "RBRACE",
		13: "LBRACKET",
		14: "RBRACKET",
		15: "COMMA",
		16: "DECLARATION",
		17: "ASSIGNMENT",
		18: "SEMICOLON",
	}
	TokenType_value = map[string]int32{
		"UNKNOWN":     0,
		"KEYWORD":     1,
		"IDENTIFIER":  2,
		"STRING":      3,
		"INTEGER":     4,
		"FLOAT":       5,
		"BOOLEAN":     6,
		"OPERATOR":    7,
		"PUNCTUATION": 8,
		"LPAREN":      9,
		"RPAREN":      10,
		"LBRACE":      11,
		"RBRACE":      12,
		"LBRACKET":    13,
		"RBRACKET":    14,
		"COMMA":       15,
		"DECLARATION": 16,
		"ASSIGNMENT":  17,
		"SEMICOLON":   18,
	}
)

func (x TokenType) Enum() *TokenType {
	p := new(TokenType)
	*p = x
	return p
}

func (x TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_reylang_lexer_proto_enumTypes[0].Descriptor()
}

func (TokenType) Type() protoreflect.EnumType {
	return &file_reylang_lexer_proto_enumTypes[0]
}

func (x TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenType.Descriptor instead.
func (TokenType) EnumDescriptor() ([]byte, []int) {
	return file_reylang_lexer_proto_rawDescGZIP(), []int{0}
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    TokenType `protobuf:"varint,1,opt,name=token,proto3,enum=ramonberrutti.reylang.TokenType" json:"token,omitempty"`
	Value    string    `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Position uint32    `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reylang_lexer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_reylang_lexer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_reylang_lexer_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetToken() TokenType {
	if x != nil {
		return x.Token
	}
	return TokenType_UNKNOWN
}

func (x *Token) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Token) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

var File_reylang_lexer_proto protoreflect.FileDescriptor

var file_reylang_lexer_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6c, 0x65, 0x78, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72,
	0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x71, 0x0a, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72,
	0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2a,
	0x8c, 0x02, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4b, 0x45,
	0x59, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x44, 0x45, 0x4e, 0x54,
	0x49, 0x46, 0x49, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e,
	0x47, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x10, 0x04,
	0x12, 0x09, 0x0a, 0x05, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x42,
	0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x50, 0x45, 0x52,
	0x41, 0x54, 0x4f, 0x52, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x55, 0x4e, 0x43, 0x54, 0x55,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x50, 0x41, 0x52, 0x45,
	0x4e, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x10, 0x0a, 0x12,
	0x0a, 0x0a, 0x06, 0x4c, 0x42, 0x52, 0x41, 0x43, 0x45, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x52,
	0x42, 0x52, 0x41, 0x43, 0x45, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x42, 0x52, 0x41, 0x43,
	0x4b, 0x45, 0x54, 0x10, 0x0d, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x42, 0x52, 0x41, 0x43, 0x4b, 0x45,
	0x54, 0x10, 0x0e, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x10, 0x0f, 0x12, 0x0f,
	0x0a, 0x0b, 0x44, 0x45, 0x43, 0x4c, 0x41, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x10, 0x12,
	0x0e, 0x0a, 0x0a, 0x41, 0x53, 0x53, 0x49, 0x47, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x11, 0x12,
	0x0d, 0x0a, 0x09, 0x53, 0x45, 0x4d, 0x49, 0x43, 0x4f, 0x4c, 0x4f, 0x4e, 0x10, 0x12, 0x42, 0xcf,
	0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72,
	0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x42, 0x0a, 0x4c, 0x65,
	0x78, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72,
	0x75, 0x74, 0x74, 0x69, 0x2f, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0xa2, 0x02, 0x03,
	0x52, 0x52, 0x58, 0xaa, 0x02, 0x15, 0x52, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75,
	0x74, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0xca, 0x02, 0x15, 0x52, 0x61,
	0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x5c, 0x52, 0x65, 0x79, 0x6c,
	0x61, 0x6e, 0x67, 0xe2, 0x02, 0x21, 0x52, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75,
	0x74, 0x74, 0x69, 0x5c, 0x52, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x52, 0x61, 0x6d, 0x6f, 0x6e, 0x62,
	0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x3a, 0x3a, 0x52, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reylang_lexer_proto_rawDescOnce sync.Once
	file_reylang_lexer_proto_rawDescData = file_reylang_lexer_proto_rawDesc
)

func file_reylang_lexer_proto_rawDescGZIP() []byte {
	file_reylang_lexer_proto_rawDescOnce.Do(func() {
		file_reylang_lexer_proto_rawDescData = protoimpl.X.CompressGZIP(file_reylang_lexer_proto_rawDescData)
	})
	return file_reylang_lexer_proto_rawDescData
}

var file_reylang_lexer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_reylang_lexer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_reylang_lexer_proto_goTypes = []interface{}{
	(TokenType)(0), // 0: ramonberrutti.reylang.TokenType
	(*Token)(nil),  // 1: ramonberrutti.reylang.Token
}
var file_reylang_lexer_proto_depIdxs = []int32{
	0, // 0: ramonberrutti.reylang.Token.token:type_name -> ramonberrutti.reylang.TokenType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_reylang_lexer_proto_init() }
func file_reylang_lexer_proto_init() {
	if File_reylang_lexer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reylang_lexer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_reylang_lexer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reylang_lexer_proto_goTypes,
		DependencyIndexes: file_reylang_lexer_proto_depIdxs,
		EnumInfos:         file_reylang_lexer_proto_enumTypes,
		MessageInfos:      file_reylang_lexer_proto_msgTypes,
	}.Build()
	File_reylang_lexer_proto = out.File
	file_reylang_lexer_proto_rawDesc = nil
	file_reylang_lexer_proto_goTypes = nil
	file_reylang_lexer_proto_depIdxs = nil
}
