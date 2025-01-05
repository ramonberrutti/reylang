// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/reylang/node.proto

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

type ComparisonNode_Operator int32

const (
	ComparisonNode_OPERATOR_NONE ComparisonNode_Operator = 0
	ComparisonNode_OPERATOR_EQ   ComparisonNode_Operator = 1
	ComparisonNode_OPERATOR_NE   ComparisonNode_Operator = 2
	ComparisonNode_OPERATOR_LT   ComparisonNode_Operator = 3
	ComparisonNode_OPERATOR_LE   ComparisonNode_Operator = 4
	ComparisonNode_OPERATOR_GT   ComparisonNode_Operator = 5
	ComparisonNode_OPERATOR_GE   ComparisonNode_Operator = 6
)

// Enum value maps for ComparisonNode_Operator.
var (
	ComparisonNode_Operator_name = map[int32]string{
		0: "OPERATOR_NONE",
		1: "OPERATOR_EQ",
		2: "OPERATOR_NE",
		3: "OPERATOR_LT",
		4: "OPERATOR_LE",
		5: "OPERATOR_GT",
		6: "OPERATOR_GE",
	}
	ComparisonNode_Operator_value = map[string]int32{
		"OPERATOR_NONE": 0,
		"OPERATOR_EQ":   1,
		"OPERATOR_NE":   2,
		"OPERATOR_LT":   3,
		"OPERATOR_LE":   4,
		"OPERATOR_GT":   5,
		"OPERATOR_GE":   6,
	}
)

func (x ComparisonNode_Operator) Enum() *ComparisonNode_Operator {
	p := new(ComparisonNode_Operator)
	*p = x
	return p
}

func (x ComparisonNode_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ComparisonNode_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_reylang_node_proto_enumTypes[0].Descriptor()
}

func (ComparisonNode_Operator) Type() protoreflect.EnumType {
	return &file_proto_reylang_node_proto_enumTypes[0]
}

func (x ComparisonNode_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ComparisonNode_Operator.Descriptor instead.
func (ComparisonNode_Operator) EnumDescriptor() ([]byte, []int) {
	return file_proto_reylang_node_proto_rawDescGZIP(), []int{6, 0}
}

// Node is the base type for all nodes in the AST.
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to NodeType:
	//
	//	*Node_LiteralNode
	//	*Node_IdentifierNode
	//	*Node_IfNode
	//	*Node_ForNode
	//	*Node_ForRangeNode
	//	*Node_ComparisonNode
	//	*Node_FunctionCallNode
	NodeType isNode_NodeType `protobuf_oneof:"node_type"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reylang_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reylang_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_proto_reylang_node_proto_rawDescGZIP(), []int{0}
}

func (m *Node) GetNodeType() isNode_NodeType {
	if m != nil {
		return m.NodeType
	}
	return nil
}

func (x *Node) GetLiteralNode() *LiteralNode {
	if x, ok := x.GetNodeType().(*Node_LiteralNode); ok {
		return x.LiteralNode
	}
	return nil
}

func (x *Node) GetIdentifierNode() *IdentifierNode {
	if x, ok := x.GetNodeType().(*Node_IdentifierNode); ok {
		return x.IdentifierNode
	}
	return nil
}

func (x *Node) GetIfNode() *IfNode {
	if x, ok := x.GetNodeType().(*Node_IfNode); ok {
		return x.IfNode
	}
	return nil
}

func (x *Node) GetForNode() *ForNode {
	if x, ok := x.GetNodeType().(*Node_ForNode); ok {
		return x.ForNode
	}
	return nil
}

func (x *Node) GetForRangeNode() *ForRangeNode {
	if x, ok := x.GetNodeType().(*Node_ForRangeNode); ok {
		return x.ForRangeNode
	}
	return nil
}

func (x *Node) GetComparisonNode() *ComparisonNode {
	if x, ok := x.GetNodeType().(*Node_ComparisonNode); ok {
		return x.ComparisonNode
	}
	return nil
}

func (x *Node) GetFunctionCallNode() *FunctionCallNode {
	if x, ok := x.GetNodeType().(*Node_FunctionCallNode); ok {
		return x.FunctionCallNode
	}
	return nil
}

type isNode_NodeType interface {
	isNode_NodeType()
}

type Node_LiteralNode struct {
	LiteralNode *LiteralNode `protobuf:"bytes,1,opt,name=literal_node,json=literalNode,proto3,oneof"`
}

type Node_IdentifierNode struct {
	IdentifierNode *IdentifierNode `protobuf:"bytes,2,opt,name=identifier_node,json=identifierNode,proto3,oneof"`
}

type Node_IfNode struct {
	IfNode *IfNode `protobuf:"bytes,3,opt,name=if_node,json=ifNode,proto3,oneof"`
}

type Node_ForNode struct {
	ForNode *ForNode `protobuf:"bytes,4,opt,name=for_node,json=forNode,proto3,oneof"`
}

type Node_ForRangeNode struct {
	ForRangeNode *ForRangeNode `protobuf:"bytes,5,opt,name=for_range_node,json=forRangeNode,proto3,oneof"`
}

type Node_ComparisonNode struct {
	ComparisonNode *ComparisonNode `protobuf:"bytes,6,opt,name=comparison_node,json=comparisonNode,proto3,oneof"`
}

type Node_FunctionCallNode struct {
	FunctionCallNode *FunctionCallNode `protobuf:"bytes,7,opt,name=function_call_node,json=functionCallNode,proto3,oneof"`
}

func (*Node_LiteralNode) isNode_NodeType() {}

func (*Node_IdentifierNode) isNode_NodeType() {}

func (*Node_IfNode) isNode_NodeType() {}

func (*Node_ForNode) isNode_NodeType() {}

func (*Node_ForRangeNode) isNode_NodeType() {}

func (*Node_ComparisonNode) isNode_NodeType() {}

func (*Node_FunctionCallNode) isNode_NodeType() {}

type LiteralNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to LiteralType:
	//
	//	*LiteralNode_String_
	//	*LiteralNode_Integer
	//	*LiteralNode_Float
	//	*LiteralNode_Boolean
	LiteralType isLiteralNode_LiteralType `protobuf_oneof:"literal_type"`
}

func (x *LiteralNode) Reset() {
	*x = LiteralNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reylang_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiteralNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiteralNode) ProtoMessage() {}

func (x *LiteralNode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reylang_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiteralNode.ProtoReflect.Descriptor instead.
func (*LiteralNode) Descriptor() ([]byte, []int) {
	return file_proto_reylang_node_proto_rawDescGZIP(), []int{1}
}

func (m *LiteralNode) GetLiteralType() isLiteralNode_LiteralType {
	if m != nil {
		return m.LiteralType
	}
	return nil
}

func (x *LiteralNode) GetString_() string {
	if x, ok := x.GetLiteralType().(*LiteralNode_String_); ok {
		return x.String_
	}
	return ""
}

func (x *LiteralNode) GetInteger() int32 {
	if x, ok := x.GetLiteralType().(*LiteralNode_Integer); ok {
		return x.Integer
	}
	return 0
}

func (x *LiteralNode) GetFloat() float32 {
	if x, ok := x.GetLiteralType().(*LiteralNode_Float); ok {
		return x.Float
	}
	return 0
}

func (x *LiteralNode) GetBoolean() bool {
	if x, ok := x.GetLiteralType().(*LiteralNode_Boolean); ok {
		return x.Boolean
	}
	return false
}

type isLiteralNode_LiteralType interface {
	isLiteralNode_LiteralType()
}

type LiteralNode_String_ struct {
	String_ string `protobuf:"bytes,1,opt,name=string,proto3,oneof"`
}

type LiteralNode_Integer struct {
	Integer int32 `protobuf:"varint,2,opt,name=integer,proto3,oneof"`
}

type LiteralNode_Float struct {
	Float float32 `protobuf:"fixed32,3,opt,name=float,proto3,oneof"`
}

type LiteralNode_Boolean struct {
	Boolean bool `protobuf:"varint,4,opt,name=boolean,proto3,oneof"`
}

func (*LiteralNode_String_) isLiteralNode_LiteralType() {}

func (*LiteralNode_Integer) isLiteralNode_LiteralType() {}

func (*LiteralNode_Float) isLiteralNode_LiteralType() {}

func (*LiteralNode_Boolean) isLiteralNode_LiteralType() {}

type IdentifierNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *IdentifierNode) Reset() {
	*x = IdentifierNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reylang_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentifierNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentifierNode) ProtoMessage() {}

func (x *IdentifierNode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reylang_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentifierNode.ProtoReflect.Descriptor instead.
func (*IdentifierNode) Descriptor() ([]byte, []int) {
	return file_proto_reylang_node_proto_rawDescGZIP(), []int{2}
}

func (x *IdentifierNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type IfNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Condition *ComparisonNode `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	Body      []*Node         `protobuf:"bytes,2,rep,name=body,proto3" json:"body,omitempty"`
	ElseBody  []*Node         `protobuf:"bytes,3,rep,name=else_body,json=elseBody,proto3" json:"else_body,omitempty"`
}

func (x *IfNode) Reset() {
	*x = IfNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reylang_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IfNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IfNode) ProtoMessage() {}

func (x *IfNode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reylang_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IfNode.ProtoReflect.Descriptor instead.
func (*IfNode) Descriptor() ([]byte, []int) {
	return file_proto_reylang_node_proto_rawDescGZIP(), []int{3}
}

func (x *IfNode) GetCondition() *ComparisonNode {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *IfNode) GetBody() []*Node {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *IfNode) GetElseBody() []*Node {
	if x != nil {
		return x.ElseBody
	}
	return nil
}

// Going to generate the ForRangeNode from the ForNode.
type ForNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Initializer *Node   `protobuf:"bytes,1,opt,name=initializer,proto3" json:"initializer,omitempty"`
	Condition   *Node   `protobuf:"bytes,2,opt,name=condition,proto3" json:"condition,omitempty"`
	Increment   *Node   `protobuf:"bytes,3,opt,name=increment,proto3" json:"increment,omitempty"`
	Body        []*Node `protobuf:"bytes,4,rep,name=body,proto3" json:"body,omitempty"`
}

func (x *ForNode) Reset() {
	*x = ForNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reylang_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForNode) ProtoMessage() {}

func (x *ForNode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reylang_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForNode.ProtoReflect.Descriptor instead.
func (*ForNode) Descriptor() ([]byte, []int) {
	return file_proto_reylang_node_proto_rawDescGZIP(), []int{4}
}

func (x *ForNode) GetInitializer() *Node {
	if x != nil {
		return x.Initializer
	}
	return nil
}

func (x *ForNode) GetCondition() *Node {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *ForNode) GetIncrement() *Node {
	if x != nil {
		return x.Increment
	}
	return nil
}

func (x *ForNode) GetBody() []*Node {
	if x != nil {
		return x.Body
	}
	return nil
}

type ForRangeNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        *IdentifierNode `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value      *IdentifierNode `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Collection *IdentifierNode `protobuf:"bytes,3,opt,name=collection,proto3" json:"collection,omitempty"`
	Body       []*Node         `protobuf:"bytes,4,rep,name=body,proto3" json:"body,omitempty"`
}

func (x *ForRangeNode) Reset() {
	*x = ForRangeNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reylang_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForRangeNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForRangeNode) ProtoMessage() {}

func (x *ForRangeNode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reylang_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForRangeNode.ProtoReflect.Descriptor instead.
func (*ForRangeNode) Descriptor() ([]byte, []int) {
	return file_proto_reylang_node_proto_rawDescGZIP(), []int{5}
}

func (x *ForRangeNode) GetKey() *IdentifierNode {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *ForRangeNode) GetValue() *IdentifierNode {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ForRangeNode) GetCollection() *IdentifierNode {
	if x != nil {
		return x.Collection
	}
	return nil
}

func (x *ForRangeNode) GetBody() []*Node {
	if x != nil {
		return x.Body
	}
	return nil
}

type ComparisonNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Left     *Node                   `protobuf:"bytes,1,opt,name=left,proto3" json:"left,omitempty"`
	Right    *Node                   `protobuf:"bytes,2,opt,name=right,proto3" json:"right,omitempty"`
	Operator ComparisonNode_Operator `protobuf:"varint,3,opt,name=operator,proto3,enum=ramonberrutti.reylang.ComparisonNode_Operator" json:"operator,omitempty"`
}

func (x *ComparisonNode) Reset() {
	*x = ComparisonNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reylang_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComparisonNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComparisonNode) ProtoMessage() {}

func (x *ComparisonNode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reylang_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComparisonNode.ProtoReflect.Descriptor instead.
func (*ComparisonNode) Descriptor() ([]byte, []int) {
	return file_proto_reylang_node_proto_rawDescGZIP(), []int{6}
}

func (x *ComparisonNode) GetLeft() *Node {
	if x != nil {
		return x.Left
	}
	return nil
}

func (x *ComparisonNode) GetRight() *Node {
	if x != nil {
		return x.Right
	}
	return nil
}

func (x *ComparisonNode) GetOperator() ComparisonNode_Operator {
	if x != nil {
		return x.Operator
	}
	return ComparisonNode_OPERATOR_NONE
}

type FunctionCallNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      *IdentifierNode `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Arguments []*Node         `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *FunctionCallNode) Reset() {
	*x = FunctionCallNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reylang_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionCallNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionCallNode) ProtoMessage() {}

func (x *FunctionCallNode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reylang_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionCallNode.ProtoReflect.Descriptor instead.
func (*FunctionCallNode) Descriptor() ([]byte, []int) {
	return file_proto_reylang_node_proto_rawDescGZIP(), []int{7}
}

func (x *FunctionCallNode) GetName() *IdentifierNode {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *FunctionCallNode) GetArguments() []*Node {
	if x != nil {
		return x.Arguments
	}
	return nil
}

var File_proto_reylang_node_proto protoreflect.FileDescriptor

var file_proto_reylang_node_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x72, 0x61, 0x6d, 0x6f,
	0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e,
	0x67, 0x22, 0x9d, 0x04, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x6c, 0x69,
	0x74, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69,
	0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c,
	0x4e, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72,
	0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4e,
	0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x69, 0x66, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65,
	0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x49,
	0x66, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x06, 0x69, 0x66, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x3b, 0x0a, 0x08, 0x66, 0x6f, 0x72, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74,
	0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x72, 0x4e, 0x6f, 0x64,
	0x65, 0x48, 0x00, 0x52, 0x07, 0x66, 0x6f, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x4b, 0x0a, 0x0e,
	0x66, 0x6f, 0x72, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72,
	0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x72,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x66, 0x6f, 0x72,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74,
	0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x69, 0x73, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x57, 0x0a, 0x12, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62,
	0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x4e, 0x6f, 0x64, 0x65,
	0x48, 0x00, 0x52, 0x10, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c,
	0x4e, 0x6f, 0x64, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x07, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x07,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12,
	0x1a, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x6c,
	0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0xb8, 0x01, 0x0a, 0x06, 0x49, 0x66, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x43, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e,
	0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73,
	0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2f, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e,
	0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x65, 0x6c, 0x73, 0x65, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72,
	0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x08, 0x65, 0x6c, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x22, 0xef, 0x01, 0x0a,
	0x07, 0x46, 0x6f, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65,
	0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x61, 0x6d,
	0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x09, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72,
	0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x09, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x61,
	0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0xfc,
	0x01, 0x0a, 0x0c, 0x46, 0x6f, 0x72, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x37, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72,
	0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62,
	0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x61, 0x6d, 0x6f,
	0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x61, 0x6d,
	0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0xc6, 0x02,
	0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x2f, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72,
	0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6c, 0x65, 0x66,
	0x74, 0x12, 0x31, 0x0a, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69,
	0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x72,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x4a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65,
	0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x22, 0x83, 0x01, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x11, 0x0a,
	0x0d, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x45, 0x51, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x4e, 0x45,
	0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x4c,
	0x54, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f,
	0x4c, 0x45, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52,
	0x5f, 0x47, 0x54, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f,
	0x52, 0x5f, 0x47, 0x45, 0x10, 0x06, 0x22, 0x88, 0x01, 0x0a, 0x10, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x61, 0x6d, 0x6f,
	0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x61, 0x6d, 0x6f,
	0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x42, 0xd4, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62,
	0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x42,
	0x09, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65,
	0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2f, 0x72, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65,
	0x79, 0x6c, 0x61, 0x6e, 0x67, 0xa2, 0x02, 0x03, 0x52, 0x52, 0x58, 0xaa, 0x02, 0x15, 0x52, 0x61,
	0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x79, 0x6c,
	0x61, 0x6e, 0x67, 0xca, 0x02, 0x15, 0x52, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75,
	0x74, 0x74, 0x69, 0x5c, 0x52, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0xe2, 0x02, 0x21, 0x52, 0x61,
	0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x5c, 0x52, 0x65, 0x79, 0x6c,
	0x61, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x16, 0x52, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x3a,
	0x3a, 0x52, 0x65, 0x79, 0x6c, 0x61, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_reylang_node_proto_rawDescOnce sync.Once
	file_proto_reylang_node_proto_rawDescData = file_proto_reylang_node_proto_rawDesc
)

func file_proto_reylang_node_proto_rawDescGZIP() []byte {
	file_proto_reylang_node_proto_rawDescOnce.Do(func() {
		file_proto_reylang_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_reylang_node_proto_rawDescData)
	})
	return file_proto_reylang_node_proto_rawDescData
}

var file_proto_reylang_node_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_reylang_node_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_reylang_node_proto_goTypes = []interface{}{
	(ComparisonNode_Operator)(0), // 0: ramonberrutti.reylang.ComparisonNode.Operator
	(*Node)(nil),                 // 1: ramonberrutti.reylang.Node
	(*LiteralNode)(nil),          // 2: ramonberrutti.reylang.LiteralNode
	(*IdentifierNode)(nil),       // 3: ramonberrutti.reylang.IdentifierNode
	(*IfNode)(nil),               // 4: ramonberrutti.reylang.IfNode
	(*ForNode)(nil),              // 5: ramonberrutti.reylang.ForNode
	(*ForRangeNode)(nil),         // 6: ramonberrutti.reylang.ForRangeNode
	(*ComparisonNode)(nil),       // 7: ramonberrutti.reylang.ComparisonNode
	(*FunctionCallNode)(nil),     // 8: ramonberrutti.reylang.FunctionCallNode
}
var file_proto_reylang_node_proto_depIdxs = []int32{
	2,  // 0: ramonberrutti.reylang.Node.literal_node:type_name -> ramonberrutti.reylang.LiteralNode
	3,  // 1: ramonberrutti.reylang.Node.identifier_node:type_name -> ramonberrutti.reylang.IdentifierNode
	4,  // 2: ramonberrutti.reylang.Node.if_node:type_name -> ramonberrutti.reylang.IfNode
	5,  // 3: ramonberrutti.reylang.Node.for_node:type_name -> ramonberrutti.reylang.ForNode
	6,  // 4: ramonberrutti.reylang.Node.for_range_node:type_name -> ramonberrutti.reylang.ForRangeNode
	7,  // 5: ramonberrutti.reylang.Node.comparison_node:type_name -> ramonberrutti.reylang.ComparisonNode
	8,  // 6: ramonberrutti.reylang.Node.function_call_node:type_name -> ramonberrutti.reylang.FunctionCallNode
	7,  // 7: ramonberrutti.reylang.IfNode.condition:type_name -> ramonberrutti.reylang.ComparisonNode
	1,  // 8: ramonberrutti.reylang.IfNode.body:type_name -> ramonberrutti.reylang.Node
	1,  // 9: ramonberrutti.reylang.IfNode.else_body:type_name -> ramonberrutti.reylang.Node
	1,  // 10: ramonberrutti.reylang.ForNode.initializer:type_name -> ramonberrutti.reylang.Node
	1,  // 11: ramonberrutti.reylang.ForNode.condition:type_name -> ramonberrutti.reylang.Node
	1,  // 12: ramonberrutti.reylang.ForNode.increment:type_name -> ramonberrutti.reylang.Node
	1,  // 13: ramonberrutti.reylang.ForNode.body:type_name -> ramonberrutti.reylang.Node
	3,  // 14: ramonberrutti.reylang.ForRangeNode.key:type_name -> ramonberrutti.reylang.IdentifierNode
	3,  // 15: ramonberrutti.reylang.ForRangeNode.value:type_name -> ramonberrutti.reylang.IdentifierNode
	3,  // 16: ramonberrutti.reylang.ForRangeNode.collection:type_name -> ramonberrutti.reylang.IdentifierNode
	1,  // 17: ramonberrutti.reylang.ForRangeNode.body:type_name -> ramonberrutti.reylang.Node
	1,  // 18: ramonberrutti.reylang.ComparisonNode.left:type_name -> ramonberrutti.reylang.Node
	1,  // 19: ramonberrutti.reylang.ComparisonNode.right:type_name -> ramonberrutti.reylang.Node
	0,  // 20: ramonberrutti.reylang.ComparisonNode.operator:type_name -> ramonberrutti.reylang.ComparisonNode.Operator
	3,  // 21: ramonberrutti.reylang.FunctionCallNode.name:type_name -> ramonberrutti.reylang.IdentifierNode
	1,  // 22: ramonberrutti.reylang.FunctionCallNode.arguments:type_name -> ramonberrutti.reylang.Node
	23, // [23:23] is the sub-list for method output_type
	23, // [23:23] is the sub-list for method input_type
	23, // [23:23] is the sub-list for extension type_name
	23, // [23:23] is the sub-list for extension extendee
	0,  // [0:23] is the sub-list for field type_name
}

func init() { file_proto_reylang_node_proto_init() }
func file_proto_reylang_node_proto_init() {
	if File_proto_reylang_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_reylang_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_proto_reylang_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiteralNode); i {
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
		file_proto_reylang_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentifierNode); i {
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
		file_proto_reylang_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IfNode); i {
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
		file_proto_reylang_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForNode); i {
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
		file_proto_reylang_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForRangeNode); i {
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
		file_proto_reylang_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComparisonNode); i {
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
		file_proto_reylang_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionCallNode); i {
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
	file_proto_reylang_node_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Node_LiteralNode)(nil),
		(*Node_IdentifierNode)(nil),
		(*Node_IfNode)(nil),
		(*Node_ForNode)(nil),
		(*Node_ForRangeNode)(nil),
		(*Node_ComparisonNode)(nil),
		(*Node_FunctionCallNode)(nil),
	}
	file_proto_reylang_node_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*LiteralNode_String_)(nil),
		(*LiteralNode_Integer)(nil),
		(*LiteralNode_Float)(nil),
		(*LiteralNode_Boolean)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_reylang_node_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_reylang_node_proto_goTypes,
		DependencyIndexes: file_proto_reylang_node_proto_depIdxs,
		EnumInfos:         file_proto_reylang_node_proto_enumTypes,
		MessageInfos:      file_proto_reylang_node_proto_msgTypes,
	}.Build()
	File_proto_reylang_node_proto = out.File
	file_proto_reylang_node_proto_rawDesc = nil
	file_proto_reylang_node_proto_goTypes = nil
	file_proto_reylang_node_proto_depIdxs = nil
}
