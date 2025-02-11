package generated

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Expression_Window_WindowFrame_FrameType int32

const (
	Expression_Window_WindowFrame_FRAME_TYPE_UNDEFINED Expression_Window_WindowFrame_FrameType = 0
	// RowFrame treats rows in a partition individually.
	Expression_Window_WindowFrame_FRAME_TYPE_ROW Expression_Window_WindowFrame_FrameType = 1
	// RangeFrame treats rows in a partition as groups of peers.
	// All rows having the same 'ORDER BY' ordering are considered as peers.
	Expression_Window_WindowFrame_FRAME_TYPE_RANGE Expression_Window_WindowFrame_FrameType = 2
)

// Enum value maps for Expression_Window_WindowFrame_FrameType.
var (
	Expression_Window_WindowFrame_FrameType_name = map[int32]string{
		0: "FRAME_TYPE_UNDEFINED",
		1: "FRAME_TYPE_ROW",
		2: "FRAME_TYPE_RANGE",
	}
	Expression_Window_WindowFrame_FrameType_value = map[string]int32{
		"FRAME_TYPE_UNDEFINED": 0,
		"FRAME_TYPE_ROW":       1,
		"FRAME_TYPE_RANGE":     2,
	}
)

func (x Expression_Window_WindowFrame_FrameType) Enum() *Expression_Window_WindowFrame_FrameType {
	p := new(Expression_Window_WindowFrame_FrameType)
	*p = x
	return p
}

func (x Expression_Window_WindowFrame_FrameType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Expression_Window_WindowFrame_FrameType) Descriptor() protoreflect.EnumDescriptor {
	return file_spark_connect_expressions_proto_enumTypes[0].Descriptor()
}

func (Expression_Window_WindowFrame_FrameType) Type() protoreflect.EnumType {
	return &file_spark_connect_expressions_proto_enumTypes[0]
}

func (x Expression_Window_WindowFrame_FrameType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Expression_Window_WindowFrame_FrameType.Descriptor instead.
func (Expression_Window_WindowFrame_FrameType) EnumDescriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

type Expression_SortOrder_SortDirection int32

const (
	Expression_SortOrder_SORT_DIRECTION_UNSPECIFIED Expression_SortOrder_SortDirection = 0
	Expression_SortOrder_SORT_DIRECTION_ASCENDING   Expression_SortOrder_SortDirection = 1
	Expression_SortOrder_SORT_DIRECTION_DESCENDING  Expression_SortOrder_SortDirection = 2
)

// Enum value maps for Expression_SortOrder_SortDirection.
var (
	Expression_SortOrder_SortDirection_name = map[int32]string{
		0: "SORT_DIRECTION_UNSPECIFIED",
		1: "SORT_DIRECTION_ASCENDING",
		2: "SORT_DIRECTION_DESCENDING",
	}
	Expression_SortOrder_SortDirection_value = map[string]int32{
		"SORT_DIRECTION_UNSPECIFIED": 0,
		"SORT_DIRECTION_ASCENDING":   1,
		"SORT_DIRECTION_DESCENDING":  2,
	}
)

func (x Expression_SortOrder_SortDirection) Enum() *Expression_SortOrder_SortDirection {
	p := new(Expression_SortOrder_SortDirection)
	*p = x
	return p
}

func (x Expression_SortOrder_SortDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Expression_SortOrder_SortDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_spark_connect_expressions_proto_enumTypes[1].Descriptor()
}

func (Expression_SortOrder_SortDirection) Type() protoreflect.EnumType {
	return &file_spark_connect_expressions_proto_enumTypes[1]
}

func (x Expression_SortOrder_SortDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Expression_SortOrder_SortDirection.Descriptor instead.
func (Expression_SortOrder_SortDirection) EnumDescriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 1, 0}
}

type Expression_SortOrder_NullOrdering int32

const (
	Expression_SortOrder_SORT_NULLS_UNSPECIFIED Expression_SortOrder_NullOrdering = 0
	Expression_SortOrder_SORT_NULLS_FIRST       Expression_SortOrder_NullOrdering = 1
	Expression_SortOrder_SORT_NULLS_LAST        Expression_SortOrder_NullOrdering = 2
)

// Enum value maps for Expression_SortOrder_NullOrdering.
var (
	Expression_SortOrder_NullOrdering_name = map[int32]string{
		0: "SORT_NULLS_UNSPECIFIED",
		1: "SORT_NULLS_FIRST",
		2: "SORT_NULLS_LAST",
	}
	Expression_SortOrder_NullOrdering_value = map[string]int32{
		"SORT_NULLS_UNSPECIFIED": 0,
		"SORT_NULLS_FIRST":       1,
		"SORT_NULLS_LAST":        2,
	}
)

func (x Expression_SortOrder_NullOrdering) Enum() *Expression_SortOrder_NullOrdering {
	p := new(Expression_SortOrder_NullOrdering)
	*p = x
	return p
}

func (x Expression_SortOrder_NullOrdering) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Expression_SortOrder_NullOrdering) Descriptor() protoreflect.EnumDescriptor {
	return file_spark_connect_expressions_proto_enumTypes[2].Descriptor()
}

func (Expression_SortOrder_NullOrdering) Type() protoreflect.EnumType {
	return &file_spark_connect_expressions_proto_enumTypes[2]
}

func (x Expression_SortOrder_NullOrdering) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Expression_SortOrder_NullOrdering.Descriptor instead.
func (Expression_SortOrder_NullOrdering) EnumDescriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 1, 1}
}

// Expression used to refer to fields, functions and similar. This can be used everywhere
// expressions in SQL appear.
type Expression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ExprType:
	//
	//	*Expression_Literal_
	//	*Expression_UnresolvedAttribute_
	//	*Expression_UnresolvedFunction_
	//	*Expression_ExpressionString_
	//	*Expression_UnresolvedStar_
	//	*Expression_Alias_
	//	*Expression_Cast_
	//	*Expression_UnresolvedRegex_
	//	*Expression_SortOrder_
	//	*Expression_LambdaFunction_
	//	*Expression_Window_
	//	*Expression_UnresolvedExtractValue_
	//	*Expression_UpdateFields_
	//	*Expression_UnresolvedNamedLambdaVariable_
	//	*Expression_CommonInlineUserDefinedFunction
	//	*Expression_CallFunction
	//	*Expression_Extension
	ExprType isExpression_ExprType `protobuf_oneof:"expr_type"`
}

func (x *Expression) Reset() {
	*x = Expression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression) ProtoMessage() {}

func (x *Expression) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression.ProtoReflect.Descriptor instead.
func (*Expression) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0}
}

func (m *Expression) GetExprType() isExpression_ExprType {
	if m != nil {
		return m.ExprType
	}
	return nil
}

func (x *Expression) GetLiteral() *Expression_Literal {
	if x, ok := x.GetExprType().(*Expression_Literal_); ok {
		return x.Literal
	}
	return nil
}

func (x *Expression) GetUnresolvedAttribute() *Expression_UnresolvedAttribute {
	if x, ok := x.GetExprType().(*Expression_UnresolvedAttribute_); ok {
		return x.UnresolvedAttribute
	}
	return nil
}

func (x *Expression) GetUnresolvedFunction() *Expression_UnresolvedFunction {
	if x, ok := x.GetExprType().(*Expression_UnresolvedFunction_); ok {
		return x.UnresolvedFunction
	}
	return nil
}

func (x *Expression) GetExpressionString() *Expression_ExpressionString {
	if x, ok := x.GetExprType().(*Expression_ExpressionString_); ok {
		return x.ExpressionString
	}
	return nil
}

func (x *Expression) GetUnresolvedStar() *Expression_UnresolvedStar {
	if x, ok := x.GetExprType().(*Expression_UnresolvedStar_); ok {
		return x.UnresolvedStar
	}
	return nil
}

func (x *Expression) GetAlias() *Expression_Alias {
	if x, ok := x.GetExprType().(*Expression_Alias_); ok {
		return x.Alias
	}
	return nil
}

func (x *Expression) GetCast() *Expression_Cast {
	if x, ok := x.GetExprType().(*Expression_Cast_); ok {
		return x.Cast
	}
	return nil
}

func (x *Expression) GetUnresolvedRegex() *Expression_UnresolvedRegex {
	if x, ok := x.GetExprType().(*Expression_UnresolvedRegex_); ok {
		return x.UnresolvedRegex
	}
	return nil
}

func (x *Expression) GetSortOrder() *Expression_SortOrder {
	if x, ok := x.GetExprType().(*Expression_SortOrder_); ok {
		return x.SortOrder
	}
	return nil
}

func (x *Expression) GetLambdaFunction() *Expression_LambdaFunction {
	if x, ok := x.GetExprType().(*Expression_LambdaFunction_); ok {
		return x.LambdaFunction
	}
	return nil
}

func (x *Expression) GetWindow() *Expression_Window {
	if x, ok := x.GetExprType().(*Expression_Window_); ok {
		return x.Window
	}
	return nil
}

func (x *Expression) GetUnresolvedExtractValue() *Expression_UnresolvedExtractValue {
	if x, ok := x.GetExprType().(*Expression_UnresolvedExtractValue_); ok {
		return x.UnresolvedExtractValue
	}
	return nil
}

func (x *Expression) GetUpdateFields() *Expression_UpdateFields {
	if x, ok := x.GetExprType().(*Expression_UpdateFields_); ok {
		return x.UpdateFields
	}
	return nil
}

func (x *Expression) GetUnresolvedNamedLambdaVariable() *Expression_UnresolvedNamedLambdaVariable {
	if x, ok := x.GetExprType().(*Expression_UnresolvedNamedLambdaVariable_); ok {
		return x.UnresolvedNamedLambdaVariable
	}
	return nil
}

func (x *Expression) GetCommonInlineUserDefinedFunction() *CommonInlineUserDefinedFunction {
	if x, ok := x.GetExprType().(*Expression_CommonInlineUserDefinedFunction); ok {
		return x.CommonInlineUserDefinedFunction
	}
	return nil
}

func (x *Expression) GetCallFunction() *CallFunction {
	if x, ok := x.GetExprType().(*Expression_CallFunction); ok {
		return x.CallFunction
	}
	return nil
}

func (x *Expression) GetExtension() *anypb.Any {
	if x, ok := x.GetExprType().(*Expression_Extension); ok {
		return x.Extension
	}
	return nil
}

type isExpression_ExprType interface {
	isExpression_ExprType()
}

type Expression_Literal_ struct {
	Literal *Expression_Literal `protobuf:"bytes,1,opt,name=literal,proto3,oneof"`
}

type Expression_UnresolvedAttribute_ struct {
	UnresolvedAttribute *Expression_UnresolvedAttribute `protobuf:"bytes,2,opt,name=unresolved_attribute,json=unresolvedAttribute,proto3,oneof"`
}

type Expression_UnresolvedFunction_ struct {
	UnresolvedFunction *Expression_UnresolvedFunction `protobuf:"bytes,3,opt,name=unresolved_function,json=unresolvedFunction,proto3,oneof"`
}

type Expression_ExpressionString_ struct {
	ExpressionString *Expression_ExpressionString `protobuf:"bytes,4,opt,name=expression_string,json=expressionString,proto3,oneof"`
}

type Expression_UnresolvedStar_ struct {
	UnresolvedStar *Expression_UnresolvedStar `protobuf:"bytes,5,opt,name=unresolved_star,json=unresolvedStar,proto3,oneof"`
}

type Expression_Alias_ struct {
	Alias *Expression_Alias `protobuf:"bytes,6,opt,name=alias,proto3,oneof"`
}

type Expression_Cast_ struct {
	Cast *Expression_Cast `protobuf:"bytes,7,opt,name=cast,proto3,oneof"`
}

type Expression_UnresolvedRegex_ struct {
	UnresolvedRegex *Expression_UnresolvedRegex `protobuf:"bytes,8,opt,name=unresolved_regex,json=unresolvedRegex,proto3,oneof"`
}

type Expression_SortOrder_ struct {
	SortOrder *Expression_SortOrder `protobuf:"bytes,9,opt,name=sort_order,json=sortOrder,proto3,oneof"`
}

type Expression_LambdaFunction_ struct {
	LambdaFunction *Expression_LambdaFunction `protobuf:"bytes,10,opt,name=lambda_function,json=lambdaFunction,proto3,oneof"`
}

type Expression_Window_ struct {
	Window *Expression_Window `protobuf:"bytes,11,opt,name=window,proto3,oneof"`
}

type Expression_UnresolvedExtractValue_ struct {
	UnresolvedExtractValue *Expression_UnresolvedExtractValue `protobuf:"bytes,12,opt,name=unresolved_extract_value,json=unresolvedExtractValue,proto3,oneof"`
}

type Expression_UpdateFields_ struct {
	UpdateFields *Expression_UpdateFields `protobuf:"bytes,13,opt,name=update_fields,json=updateFields,proto3,oneof"`
}

type Expression_UnresolvedNamedLambdaVariable_ struct {
	UnresolvedNamedLambdaVariable *Expression_UnresolvedNamedLambdaVariable `protobuf:"bytes,14,opt,name=unresolved_named_lambda_variable,json=unresolvedNamedLambdaVariable,proto3,oneof"`
}

type Expression_CommonInlineUserDefinedFunction struct {
	CommonInlineUserDefinedFunction *CommonInlineUserDefinedFunction `protobuf:"bytes,15,opt,name=common_inline_user_defined_function,json=commonInlineUserDefinedFunction,proto3,oneof"`
}

type Expression_CallFunction struct {
	CallFunction *CallFunction `protobuf:"bytes,16,opt,name=call_function,json=callFunction,proto3,oneof"`
}

type Expression_Extension struct {
	// This field is used to mark extensions to the protocol. When plugins generate arbitrary
	// relations they can add them here. During the planning the correct resolution is done.
	Extension *anypb.Any `protobuf:"bytes,999,opt,name=extension,proto3,oneof"`
}

func (*Expression_Literal_) isExpression_ExprType() {}

func (*Expression_UnresolvedAttribute_) isExpression_ExprType() {}

func (*Expression_UnresolvedFunction_) isExpression_ExprType() {}

func (*Expression_ExpressionString_) isExpression_ExprType() {}

func (*Expression_UnresolvedStar_) isExpression_ExprType() {}

func (*Expression_Alias_) isExpression_ExprType() {}

func (*Expression_Cast_) isExpression_ExprType() {}

func (*Expression_UnresolvedRegex_) isExpression_ExprType() {}

func (*Expression_SortOrder_) isExpression_ExprType() {}

func (*Expression_LambdaFunction_) isExpression_ExprType() {}

func (*Expression_Window_) isExpression_ExprType() {}

func (*Expression_UnresolvedExtractValue_) isExpression_ExprType() {}

func (*Expression_UpdateFields_) isExpression_ExprType() {}

func (*Expression_UnresolvedNamedLambdaVariable_) isExpression_ExprType() {}

func (*Expression_CommonInlineUserDefinedFunction) isExpression_ExprType() {}

func (*Expression_CallFunction) isExpression_ExprType() {}

func (*Expression_Extension) isExpression_ExprType() {}

type CommonInlineUserDefinedFunction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) Name of the user-defined function.
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	// (Optional) Indicate if the user-defined function is deterministic.
	Deterministic bool `protobuf:"varint,2,opt,name=deterministic,proto3" json:"deterministic,omitempty"`
	// (Optional) Function arguments. Empty arguments are allowed.
	Arguments []*Expression `protobuf:"bytes,3,rep,name=arguments,proto3" json:"arguments,omitempty"`
	// (Required) Indicate the function type of the user-defined function.
	//
	// Types that are assignable to Function:
	//
	//	*CommonInlineUserDefinedFunction_PythonUdf
	//	*CommonInlineUserDefinedFunction_ScalarScalaUdf
	//	*CommonInlineUserDefinedFunction_JavaUdf
	Function isCommonInlineUserDefinedFunction_Function `protobuf_oneof:"function"`
}

func (x *CommonInlineUserDefinedFunction) Reset() {
	*x = CommonInlineUserDefinedFunction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonInlineUserDefinedFunction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonInlineUserDefinedFunction) ProtoMessage() {}

func (x *CommonInlineUserDefinedFunction) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonInlineUserDefinedFunction.ProtoReflect.Descriptor instead.
func (*CommonInlineUserDefinedFunction) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{1}
}

func (x *CommonInlineUserDefinedFunction) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *CommonInlineUserDefinedFunction) GetDeterministic() bool {
	if x != nil {
		return x.Deterministic
	}
	return false
}

func (x *CommonInlineUserDefinedFunction) GetArguments() []*Expression {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (m *CommonInlineUserDefinedFunction) GetFunction() isCommonInlineUserDefinedFunction_Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (x *CommonInlineUserDefinedFunction) GetPythonUdf() *PythonUDF {
	if x, ok := x.GetFunction().(*CommonInlineUserDefinedFunction_PythonUdf); ok {
		return x.PythonUdf
	}
	return nil
}

func (x *CommonInlineUserDefinedFunction) GetScalarScalaUdf() *ScalarScalaUDF {
	if x, ok := x.GetFunction().(*CommonInlineUserDefinedFunction_ScalarScalaUdf); ok {
		return x.ScalarScalaUdf
	}
	return nil
}

func (x *CommonInlineUserDefinedFunction) GetJavaUdf() *JavaUDF {
	if x, ok := x.GetFunction().(*CommonInlineUserDefinedFunction_JavaUdf); ok {
		return x.JavaUdf
	}
	return nil
}

type isCommonInlineUserDefinedFunction_Function interface {
	isCommonInlineUserDefinedFunction_Function()
}

type CommonInlineUserDefinedFunction_PythonUdf struct {
	PythonUdf *PythonUDF `protobuf:"bytes,4,opt,name=python_udf,json=pythonUdf,proto3,oneof"`
}

type CommonInlineUserDefinedFunction_ScalarScalaUdf struct {
	ScalarScalaUdf *ScalarScalaUDF `protobuf:"bytes,5,opt,name=scalar_scala_udf,json=scalarScalaUdf,proto3,oneof"`
}

type CommonInlineUserDefinedFunction_JavaUdf struct {
	JavaUdf *JavaUDF `protobuf:"bytes,6,opt,name=java_udf,json=javaUdf,proto3,oneof"`
}

func (*CommonInlineUserDefinedFunction_PythonUdf) isCommonInlineUserDefinedFunction_Function() {}

func (*CommonInlineUserDefinedFunction_ScalarScalaUdf) isCommonInlineUserDefinedFunction_Function() {}

func (*CommonInlineUserDefinedFunction_JavaUdf) isCommonInlineUserDefinedFunction_Function() {}

type PythonUDF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) Output type of the Python UDF
	OutputType *DataType `protobuf:"bytes,1,opt,name=output_type,json=outputType,proto3" json:"output_type,omitempty"`
	// (Required) EvalType of the Python UDF
	EvalType int32 `protobuf:"varint,2,opt,name=eval_type,json=evalType,proto3" json:"eval_type,omitempty"`
	// (Required) The encoded commands of the Python UDF
	Command []byte `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	// (Required) Python version being used in the client.
	PythonVer string `protobuf:"bytes,4,opt,name=python_ver,json=pythonVer,proto3" json:"python_ver,omitempty"`
}

func (x *PythonUDF) Reset() {
	*x = PythonUDF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PythonUDF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PythonUDF) ProtoMessage() {}

func (x *PythonUDF) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PythonUDF.ProtoReflect.Descriptor instead.
func (*PythonUDF) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{2}
}

func (x *PythonUDF) GetOutputType() *DataType {
	if x != nil {
		return x.OutputType
	}
	return nil
}

func (x *PythonUDF) GetEvalType() int32 {
	if x != nil {
		return x.EvalType
	}
	return 0
}

func (x *PythonUDF) GetCommand() []byte {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *PythonUDF) GetPythonVer() string {
	if x != nil {
		return x.PythonVer
	}
	return ""
}

type ScalarScalaUDF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) Serialized JVM object containing UDF definition, input encoders and output encoder
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// (Optional) Input type(s) of the UDF
	InputTypes []*DataType `protobuf:"bytes,2,rep,name=inputTypes,proto3" json:"inputTypes,omitempty"`
	// (Required) Output type of the UDF
	OutputType *DataType `protobuf:"bytes,3,opt,name=outputType,proto3" json:"outputType,omitempty"`
	// (Required) True if the UDF can return null value
	Nullable bool `protobuf:"varint,4,opt,name=nullable,proto3" json:"nullable,omitempty"`
}

func (x *ScalarScalaUDF) Reset() {
	*x = ScalarScalaUDF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScalarScalaUDF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScalarScalaUDF) ProtoMessage() {}

func (x *ScalarScalaUDF) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScalarScalaUDF.ProtoReflect.Descriptor instead.
func (*ScalarScalaUDF) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{3}
}

func (x *ScalarScalaUDF) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ScalarScalaUDF) GetInputTypes() []*DataType {
	if x != nil {
		return x.InputTypes
	}
	return nil
}

func (x *ScalarScalaUDF) GetOutputType() *DataType {
	if x != nil {
		return x.OutputType
	}
	return nil
}

func (x *ScalarScalaUDF) GetNullable() bool {
	if x != nil {
		return x.Nullable
	}
	return false
}

type JavaUDF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) Fully qualified name of Java class
	ClassName string `protobuf:"bytes,1,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	// (Optional) Output type of the Java UDF
	OutputType *DataType `protobuf:"bytes,2,opt,name=output_type,json=outputType,proto3,oneof" json:"output_type,omitempty"`
	// (Required) Indicate if the Java user-defined function is an aggregate function
	Aggregate bool `protobuf:"varint,3,opt,name=aggregate,proto3" json:"aggregate,omitempty"`
}

func (x *JavaUDF) Reset() {
	*x = JavaUDF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JavaUDF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JavaUDF) ProtoMessage() {}

func (x *JavaUDF) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JavaUDF.ProtoReflect.Descriptor instead.
func (*JavaUDF) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{4}
}

func (x *JavaUDF) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *JavaUDF) GetOutputType() *DataType {
	if x != nil {
		return x.OutputType
	}
	return nil
}

func (x *JavaUDF) GetAggregate() bool {
	if x != nil {
		return x.Aggregate
	}
	return false
}

type CallFunction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) Unparsed name of the SQL function.
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	// (Optional) Function arguments. Empty arguments are allowed.
	Arguments []*Expression `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *CallFunction) Reset() {
	*x = CallFunction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallFunction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallFunction) ProtoMessage() {}

func (x *CallFunction) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallFunction.ProtoReflect.Descriptor instead.
func (*CallFunction) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{5}
}

func (x *CallFunction) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *CallFunction) GetArguments() []*Expression {
	if x != nil {
		return x.Arguments
	}
	return nil
}

// Expression for the OVER clause or WINDOW clause.
type Expression_Window struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The window function.
	WindowFunction *Expression `protobuf:"bytes,1,opt,name=window_function,json=windowFunction,proto3" json:"window_function,omitempty"`
	// (Optional) The way that input rows are partitioned.
	PartitionSpec []*Expression `protobuf:"bytes,2,rep,name=partition_spec,json=partitionSpec,proto3" json:"partition_spec,omitempty"`
	// (Optional) Ordering of rows in a partition.
	OrderSpec []*Expression_SortOrder `protobuf:"bytes,3,rep,name=order_spec,json=orderSpec,proto3" json:"order_spec,omitempty"`
	// (Optional) Window frame in a partition.
	//
	// If not set, it will be treated as 'UnspecifiedFrame'.
	FrameSpec *Expression_Window_WindowFrame `protobuf:"bytes,4,opt,name=frame_spec,json=frameSpec,proto3" json:"frame_spec,omitempty"`
}

func (x *Expression_Window) Reset() {
	*x = Expression_Window{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_Window) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_Window) ProtoMessage() {}

func (x *Expression_Window) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_Window.ProtoReflect.Descriptor instead.
func (*Expression_Window) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Expression_Window) GetWindowFunction() *Expression {
	if x != nil {
		return x.WindowFunction
	}
	return nil
}

func (x *Expression_Window) GetPartitionSpec() []*Expression {
	if x != nil {
		return x.PartitionSpec
	}
	return nil
}

func (x *Expression_Window) GetOrderSpec() []*Expression_SortOrder {
	if x != nil {
		return x.OrderSpec
	}
	return nil
}

func (x *Expression_Window) GetFrameSpec() *Expression_Window_WindowFrame {
	if x != nil {
		return x.FrameSpec
	}
	return nil
}

// SortOrder is used to specify the  data ordering, it is normally used in Sort and Window.
// It is an unevaluable expression and cannot be evaluated, so can not be used in Projection.
type Expression_SortOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The expression to be sorted.
	Child *Expression `protobuf:"bytes,1,opt,name=child,proto3" json:"child,omitempty"`
	// (Required) The sort direction, should be ASCENDING or DESCENDING.
	Direction Expression_SortOrder_SortDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=spark.connect.Expression_SortOrder_SortDirection" json:"direction,omitempty"`
	// (Required) How to deal with NULLs, should be NULLS_FIRST or NULLS_LAST.
	NullOrdering Expression_SortOrder_NullOrdering `protobuf:"varint,3,opt,name=null_ordering,json=nullOrdering,proto3,enum=spark.connect.Expression_SortOrder_NullOrdering" json:"null_ordering,omitempty"`
}

func (x *Expression_SortOrder) Reset() {
	*x = Expression_SortOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_SortOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_SortOrder) ProtoMessage() {}

func (x *Expression_SortOrder) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_SortOrder.ProtoReflect.Descriptor instead.
func (*Expression_SortOrder) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Expression_SortOrder) GetChild() *Expression {
	if x != nil {
		return x.Child
	}
	return nil
}

func (x *Expression_SortOrder) GetDirection() Expression_SortOrder_SortDirection {
	if x != nil {
		return x.Direction
	}
	return Expression_SortOrder_SORT_DIRECTION_UNSPECIFIED
}

func (x *Expression_SortOrder) GetNullOrdering() Expression_SortOrder_NullOrdering {
	if x != nil {
		return x.NullOrdering
	}
	return Expression_SortOrder_SORT_NULLS_UNSPECIFIED
}

type Expression_Cast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) the expression to be casted.
	Expr *Expression `protobuf:"bytes,1,opt,name=expr,proto3" json:"expr,omitempty"`
	// (Required) the data type that the expr to be casted to.
	//
	// Types that are assignable to CastToType:
	//
	//	*Expression_Cast_Type
	//	*Expression_Cast_TypeStr
	CastToType isExpression_Cast_CastToType `protobuf_oneof:"cast_to_type"`
}

func (x *Expression_Cast) Reset() {
	*x = Expression_Cast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_Cast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_Cast) ProtoMessage() {}

func (x *Expression_Cast) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_Cast.ProtoReflect.Descriptor instead.
func (*Expression_Cast) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Expression_Cast) GetExpr() *Expression {
	if x != nil {
		return x.Expr
	}
	return nil
}

func (m *Expression_Cast) GetCastToType() isExpression_Cast_CastToType {
	if m != nil {
		return m.CastToType
	}
	return nil
}

func (x *Expression_Cast) GetType() *DataType {
	if x, ok := x.GetCastToType().(*Expression_Cast_Type); ok {
		return x.Type
	}
	return nil
}

func (x *Expression_Cast) GetTypeStr() string {
	if x, ok := x.GetCastToType().(*Expression_Cast_TypeStr); ok {
		return x.TypeStr
	}
	return ""
}

type isExpression_Cast_CastToType interface {
	isExpression_Cast_CastToType()
}

type Expression_Cast_Type struct {
	Type *DataType `protobuf:"bytes,2,opt,name=type,proto3,oneof"`
}

type Expression_Cast_TypeStr struct {
	// If this is set, Server will use Catalyst parser to parse this string to DataType.
	TypeStr string `protobuf:"bytes,3,opt,name=type_str,json=typeStr,proto3,oneof"`
}

func (*Expression_Cast_Type) isExpression_Cast_CastToType() {}

func (*Expression_Cast_TypeStr) isExpression_Cast_CastToType() {}

type Expression_Literal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to LiteralType:
	//
	//	*Expression_Literal_Null
	//	*Expression_Literal_Binary
	//	*Expression_Literal_Boolean
	//	*Expression_Literal_Byte
	//	*Expression_Literal_Short
	//	*Expression_Literal_Integer
	//	*Expression_Literal_Long
	//	*Expression_Literal_Float
	//	*Expression_Literal_Double
	//	*Expression_Literal_Decimal_
	//	*Expression_Literal_String_
	//	*Expression_Literal_Date
	//	*Expression_Literal_Timestamp
	//	*Expression_Literal_TimestampNtz
	//	*Expression_Literal_CalendarInterval_
	//	*Expression_Literal_YearMonthInterval
	//	*Expression_Literal_DayTimeInterval
	//	*Expression_Literal_Array_
	//	*Expression_Literal_Map_
	//	*Expression_Literal_Struct_
	LiteralType isExpression_Literal_LiteralType `protobuf_oneof:"literal_type"`
}

func (x *Expression_Literal) Reset() {
	*x = Expression_Literal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_Literal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_Literal) ProtoMessage() {}

func (x *Expression_Literal) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_Literal.ProtoReflect.Descriptor instead.
func (*Expression_Literal) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 3}
}

func (m *Expression_Literal) GetLiteralType() isExpression_Literal_LiteralType {
	if m != nil {
		return m.LiteralType
	}
	return nil
}

func (x *Expression_Literal) GetNull() *DataType {
	if x, ok := x.GetLiteralType().(*Expression_Literal_Null); ok {
		return x.Null
	}
	return nil
}

func (x *Expression_Literal) GetBinary() []byte {
	if x, ok := x.GetLiteralType().(*Expression_Literal_Binary); ok {
		return x.Binary
	}
	return nil
}

func (x *Expression_Literal) GetBoolean() bool {
	if x, ok := x.GetLiteralType().(*Expression_Literal_Boolean); ok {
		return x.Boolean
	}
	return false
}

func (x *Expression_Literal) GetByte() int32 {
	if x, ok := x.GetLiteralType().(*Expression_Literal_Byte); ok {
		return x.Byte
	}
	return 0
}

func (x *Expression_Literal) GetShort() int32 {
	if x, ok := x.GetLiteralType().(*Expression_Literal_Short); ok {
		return x.Short
	}
	return 0
}

func (x *Expression_Literal) GetInteger() int32 {
	if x, ok := x.GetLiteralType().(*Expression_Literal_Integer); ok {
		return x.Integer
	}
	return 0
}

func (x *Expression_Literal) GetLong() int64 {
	if x, ok := x.GetLiteralType().(*Expression_Literal_Long); ok {
		return x.Long
	}
	return 0
}

func (x *Expression_Literal) GetFloat() float32 {
	if x, ok := x.GetLiteralType().(*Expression_Literal_Float); ok {
		return x.Float
	}
	return 0
}

func (x *Expression_Literal) GetDouble() float64 {
	if x, ok := x.GetLiteralType().(*Expression_Literal_Double); ok {
		return x.Double
	}
	return 0
}

func (x *Expression_Literal) GetDecimal() *Expression_Literal_Decimal {
	if x, ok := x.GetLiteralType().(*Expression_Literal_Decimal_); ok {
		return x.Decimal
	}
	return nil
}

func (x *Expression_Literal) GetString_() string {
	if x, ok := x.GetLiteralType().(*Expression_Literal_String_); ok {
		return x.String_
	}
	return ""
}

func (x *Expression_Literal) GetDate() int32 {
	if x, ok := x.GetLiteralType().(*Expression_Literal_Date); ok {
		return x.Date
	}
	return 0
}

func (x *Expression_Literal) GetTimestamp() int64 {
	if x, ok := x.GetLiteralType().(*Expression_Literal_Timestamp); ok {
		return x.Timestamp
	}
	return 0
}

func (x *Expression_Literal) GetTimestampNtz() int64 {
	if x, ok := x.GetLiteralType().(*Expression_Literal_TimestampNtz); ok {
		return x.TimestampNtz
	}
	return 0
}

func (x *Expression_Literal) GetCalendarInterval() *Expression_Literal_CalendarInterval {
	if x, ok := x.GetLiteralType().(*Expression_Literal_CalendarInterval_); ok {
		return x.CalendarInterval
	}
	return nil
}

func (x *Expression_Literal) GetYearMonthInterval() int32 {
	if x, ok := x.GetLiteralType().(*Expression_Literal_YearMonthInterval); ok {
		return x.YearMonthInterval
	}
	return 0
}

func (x *Expression_Literal) GetDayTimeInterval() int64 {
	if x, ok := x.GetLiteralType().(*Expression_Literal_DayTimeInterval); ok {
		return x.DayTimeInterval
	}
	return 0
}

func (x *Expression_Literal) GetArray() *Expression_Literal_Array {
	if x, ok := x.GetLiteralType().(*Expression_Literal_Array_); ok {
		return x.Array
	}
	return nil
}

func (x *Expression_Literal) GetMap() *Expression_Literal_Map {
	if x, ok := x.GetLiteralType().(*Expression_Literal_Map_); ok {
		return x.Map
	}
	return nil
}

func (x *Expression_Literal) GetStruct() *Expression_Literal_Struct {
	if x, ok := x.GetLiteralType().(*Expression_Literal_Struct_); ok {
		return x.Struct
	}
	return nil
}

type isExpression_Literal_LiteralType interface {
	isExpression_Literal_LiteralType()
}

type Expression_Literal_Null struct {
	Null *DataType `protobuf:"bytes,1,opt,name=null,proto3,oneof"`
}

type Expression_Literal_Binary struct {
	Binary []byte `protobuf:"bytes,2,opt,name=binary,proto3,oneof"`
}

type Expression_Literal_Boolean struct {
	Boolean bool `protobuf:"varint,3,opt,name=boolean,proto3,oneof"`
}

type Expression_Literal_Byte struct {
	Byte int32 `protobuf:"varint,4,opt,name=byte,proto3,oneof"`
}

type Expression_Literal_Short struct {
	Short int32 `protobuf:"varint,5,opt,name=short,proto3,oneof"`
}

type Expression_Literal_Integer struct {
	Integer int32 `protobuf:"varint,6,opt,name=integer,proto3,oneof"`
}

type Expression_Literal_Long struct {
	Long int64 `protobuf:"varint,7,opt,name=long,proto3,oneof"`
}

type Expression_Literal_Float struct {
	Float float32 `protobuf:"fixed32,10,opt,name=float,proto3,oneof"`
}

type Expression_Literal_Double struct {
	Double float64 `protobuf:"fixed64,11,opt,name=double,proto3,oneof"`
}

type Expression_Literal_Decimal_ struct {
	Decimal *Expression_Literal_Decimal `protobuf:"bytes,12,opt,name=decimal,proto3,oneof"`
}

type Expression_Literal_String_ struct {
	String_ string `protobuf:"bytes,13,opt,name=string,proto3,oneof"`
}

type Expression_Literal_Date struct {
	// Date in units of days since the UNIX epoch.
	Date int32 `protobuf:"varint,16,opt,name=date,proto3,oneof"`
}

type Expression_Literal_Timestamp struct {
	// Timestamp in units of microseconds since the UNIX epoch.
	Timestamp int64 `protobuf:"varint,17,opt,name=timestamp,proto3,oneof"`
}

type Expression_Literal_TimestampNtz struct {
	// Timestamp in units of microseconds since the UNIX epoch (without timezone information).
	TimestampNtz int64 `protobuf:"varint,18,opt,name=timestamp_ntz,json=timestampNtz,proto3,oneof"`
}

type Expression_Literal_CalendarInterval_ struct {
	CalendarInterval *Expression_Literal_CalendarInterval `protobuf:"bytes,19,opt,name=calendar_interval,json=calendarInterval,proto3,oneof"`
}

type Expression_Literal_YearMonthInterval struct {
	YearMonthInterval int32 `protobuf:"varint,20,opt,name=year_month_interval,json=yearMonthInterval,proto3,oneof"`
}

type Expression_Literal_DayTimeInterval struct {
	DayTimeInterval int64 `protobuf:"varint,21,opt,name=day_time_interval,json=dayTimeInterval,proto3,oneof"`
}

type Expression_Literal_Array_ struct {
	Array *Expression_Literal_Array `protobuf:"bytes,22,opt,name=array,proto3,oneof"`
}

type Expression_Literal_Map_ struct {
	Map *Expression_Literal_Map `protobuf:"bytes,23,opt,name=map,proto3,oneof"`
}

type Expression_Literal_Struct_ struct {
	Struct *Expression_Literal_Struct `protobuf:"bytes,24,opt,name=struct,proto3,oneof"`
}

func (*Expression_Literal_Null) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_Binary) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_Boolean) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_Byte) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_Short) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_Integer) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_Long) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_Float) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_Double) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_Decimal_) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_String_) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_Date) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_Timestamp) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_TimestampNtz) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_CalendarInterval_) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_YearMonthInterval) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_DayTimeInterval) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_Array_) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_Map_) isExpression_Literal_LiteralType() {}

func (*Expression_Literal_Struct_) isExpression_Literal_LiteralType() {}

// An unresolved attribute that is not explicitly bound to a specific column, but the column
// is resolved during analysis by name.
type Expression_UnresolvedAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) An identifier that will be parsed by Catalyst parser. This should follow the
	// Spark SQL identifier syntax.
	UnparsedIdentifier string `protobuf:"bytes,1,opt,name=unparsed_identifier,json=unparsedIdentifier,proto3" json:"unparsed_identifier,omitempty"`
	// (Optional) The id of corresponding connect plan.
	PlanId *int64 `protobuf:"varint,2,opt,name=plan_id,json=planId,proto3,oneof" json:"plan_id,omitempty"`
}

func (x *Expression_UnresolvedAttribute) Reset() {
	*x = Expression_UnresolvedAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_UnresolvedAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_UnresolvedAttribute) ProtoMessage() {}

func (x *Expression_UnresolvedAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_UnresolvedAttribute.ProtoReflect.Descriptor instead.
func (*Expression_UnresolvedAttribute) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Expression_UnresolvedAttribute) GetUnparsedIdentifier() string {
	if x != nil {
		return x.UnparsedIdentifier
	}
	return ""
}

func (x *Expression_UnresolvedAttribute) GetPlanId() int64 {
	if x != nil && x.PlanId != nil {
		return *x.PlanId
	}
	return 0
}

// An unresolved function is not explicitly bound to one explicit function, but the function
// is resolved during analysis following Sparks name resolution rules.
type Expression_UnresolvedFunction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) name (or unparsed name for user defined function) for the unresolved function.
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	// (Optional) Function arguments. Empty arguments are allowed.
	Arguments []*Expression `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
	// (Required) Indicate if this function should be applied on distinct values.
	IsDistinct bool `protobuf:"varint,3,opt,name=is_distinct,json=isDistinct,proto3" json:"is_distinct,omitempty"`
	// (Required) Indicate if this is a user defined function.
	//
	// When it is not a user defined function, Connect will use the function name directly.
	// When it is a user defined function, Connect will parse the function name first.
	IsUserDefinedFunction bool `protobuf:"varint,4,opt,name=is_user_defined_function,json=isUserDefinedFunction,proto3" json:"is_user_defined_function,omitempty"`
}

func (x *Expression_UnresolvedFunction) Reset() {
	*x = Expression_UnresolvedFunction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_UnresolvedFunction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_UnresolvedFunction) ProtoMessage() {}

func (x *Expression_UnresolvedFunction) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_UnresolvedFunction.ProtoReflect.Descriptor instead.
func (*Expression_UnresolvedFunction) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 5}
}

func (x *Expression_UnresolvedFunction) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *Expression_UnresolvedFunction) GetArguments() []*Expression {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *Expression_UnresolvedFunction) GetIsDistinct() bool {
	if x != nil {
		return x.IsDistinct
	}
	return false
}

func (x *Expression_UnresolvedFunction) GetIsUserDefinedFunction() bool {
	if x != nil {
		return x.IsUserDefinedFunction
	}
	return false
}

// Expression as string.
type Expression_ExpressionString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) A SQL expression that will be parsed by Catalyst parser.
	Expression string `protobuf:"bytes,1,opt,name=expression,proto3" json:"expression,omitempty"`
}

func (x *Expression_ExpressionString) Reset() {
	*x = Expression_ExpressionString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_ExpressionString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_ExpressionString) ProtoMessage() {}

func (x *Expression_ExpressionString) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_ExpressionString.ProtoReflect.Descriptor instead.
func (*Expression_ExpressionString) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 6}
}

func (x *Expression_ExpressionString) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

// UnresolvedStar is used to expand all the fields of a relation or struct.
type Expression_UnresolvedStar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Optional) The target of the expansion.
	//
	// If set, it should end with '.*' and will be parsed by 'parseAttributeName'
	// in the server side.
	UnparsedTarget *string `protobuf:"bytes,1,opt,name=unparsed_target,json=unparsedTarget,proto3,oneof" json:"unparsed_target,omitempty"`
}

func (x *Expression_UnresolvedStar) Reset() {
	*x = Expression_UnresolvedStar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_UnresolvedStar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_UnresolvedStar) ProtoMessage() {}

func (x *Expression_UnresolvedStar) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_UnresolvedStar.ProtoReflect.Descriptor instead.
func (*Expression_UnresolvedStar) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 7}
}

func (x *Expression_UnresolvedStar) GetUnparsedTarget() string {
	if x != nil && x.UnparsedTarget != nil {
		return *x.UnparsedTarget
	}
	return ""
}

// Represents all of the input attributes to a given relational operator, for example in
// "SELECT `(id)?+.+` FROM ...".
type Expression_UnresolvedRegex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The column name used to extract column with regex.
	ColName string `protobuf:"bytes,1,opt,name=col_name,json=colName,proto3" json:"col_name,omitempty"`
	// (Optional) The id of corresponding connect plan.
	PlanId *int64 `protobuf:"varint,2,opt,name=plan_id,json=planId,proto3,oneof" json:"plan_id,omitempty"`
}

func (x *Expression_UnresolvedRegex) Reset() {
	*x = Expression_UnresolvedRegex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_UnresolvedRegex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_UnresolvedRegex) ProtoMessage() {}

func (x *Expression_UnresolvedRegex) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_UnresolvedRegex.ProtoReflect.Descriptor instead.
func (*Expression_UnresolvedRegex) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 8}
}

func (x *Expression_UnresolvedRegex) GetColName() string {
	if x != nil {
		return x.ColName
	}
	return ""
}

func (x *Expression_UnresolvedRegex) GetPlanId() int64 {
	if x != nil && x.PlanId != nil {
		return *x.PlanId
	}
	return 0
}

// Extracts a value or values from an Expression
type Expression_UnresolvedExtractValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The expression to extract value from, can be
	// Map, Array, Struct or array of Structs.
	Child *Expression `protobuf:"bytes,1,opt,name=child,proto3" json:"child,omitempty"`
	// (Required) The expression to describe the extraction, can be
	// key of Map, index of Array, field name of Struct.
	Extraction *Expression `protobuf:"bytes,2,opt,name=extraction,proto3" json:"extraction,omitempty"`
}

func (x *Expression_UnresolvedExtractValue) Reset() {
	*x = Expression_UnresolvedExtractValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_UnresolvedExtractValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_UnresolvedExtractValue) ProtoMessage() {}

func (x *Expression_UnresolvedExtractValue) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_UnresolvedExtractValue.ProtoReflect.Descriptor instead.
func (*Expression_UnresolvedExtractValue) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 9}
}

func (x *Expression_UnresolvedExtractValue) GetChild() *Expression {
	if x != nil {
		return x.Child
	}
	return nil
}

func (x *Expression_UnresolvedExtractValue) GetExtraction() *Expression {
	if x != nil {
		return x.Extraction
	}
	return nil
}

// Add, replace or drop a field of `StructType` expression by name.
type Expression_UpdateFields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The struct expression.
	StructExpression *Expression `protobuf:"bytes,1,opt,name=struct_expression,json=structExpression,proto3" json:"struct_expression,omitempty"`
	// (Required) The field name.
	FieldName string `protobuf:"bytes,2,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	// (Optional) The expression to add or replace.
	//
	// When not set, it means this field will be dropped.
	ValueExpression *Expression `protobuf:"bytes,3,opt,name=value_expression,json=valueExpression,proto3" json:"value_expression,omitempty"`
}

func (x *Expression_UpdateFields) Reset() {
	*x = Expression_UpdateFields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_UpdateFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_UpdateFields) ProtoMessage() {}

func (x *Expression_UpdateFields) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_UpdateFields.ProtoReflect.Descriptor instead.
func (*Expression_UpdateFields) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 10}
}

func (x *Expression_UpdateFields) GetStructExpression() *Expression {
	if x != nil {
		return x.StructExpression
	}
	return nil
}

func (x *Expression_UpdateFields) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *Expression_UpdateFields) GetValueExpression() *Expression {
	if x != nil {
		return x.ValueExpression
	}
	return nil
}

type Expression_Alias struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The expression that alias will be added on.
	Expr *Expression `protobuf:"bytes,1,opt,name=expr,proto3" json:"expr,omitempty"`
	// (Required) a list of name parts for the alias.
	//
	// Scalar columns only has one name that presents.
	Name []string `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty"`
	// (Optional) Alias metadata expressed as a JSON map.
	Metadata *string `protobuf:"bytes,3,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
}

func (x *Expression_Alias) Reset() {
	*x = Expression_Alias{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_Alias) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_Alias) ProtoMessage() {}

func (x *Expression_Alias) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_Alias.ProtoReflect.Descriptor instead.
func (*Expression_Alias) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 11}
}

func (x *Expression_Alias) GetExpr() *Expression {
	if x != nil {
		return x.Expr
	}
	return nil
}

func (x *Expression_Alias) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Expression_Alias) GetMetadata() string {
	if x != nil && x.Metadata != nil {
		return *x.Metadata
	}
	return ""
}

type Expression_LambdaFunction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The lambda function.
	//
	// The function body should use 'UnresolvedAttribute' as arguments, the sever side will
	// replace 'UnresolvedAttribute' with 'UnresolvedNamedLambdaVariable'.
	Function *Expression `protobuf:"bytes,1,opt,name=function,proto3" json:"function,omitempty"`
	// (Required) Function variables. Must contains 1 ~ 3 variables.
	Arguments []*Expression_UnresolvedNamedLambdaVariable `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *Expression_LambdaFunction) Reset() {
	*x = Expression_LambdaFunction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_LambdaFunction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_LambdaFunction) ProtoMessage() {}

func (x *Expression_LambdaFunction) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_LambdaFunction.ProtoReflect.Descriptor instead.
func (*Expression_LambdaFunction) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 12}
}

func (x *Expression_LambdaFunction) GetFunction() *Expression {
	if x != nil {
		return x.Function
	}
	return nil
}

func (x *Expression_LambdaFunction) GetArguments() []*Expression_UnresolvedNamedLambdaVariable {
	if x != nil {
		return x.Arguments
	}
	return nil
}

type Expression_UnresolvedNamedLambdaVariable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) a list of name parts for the variable. Must not be empty.
	NameParts []string `protobuf:"bytes,1,rep,name=name_parts,json=nameParts,proto3" json:"name_parts,omitempty"`
}

func (x *Expression_UnresolvedNamedLambdaVariable) Reset() {
	*x = Expression_UnresolvedNamedLambdaVariable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_UnresolvedNamedLambdaVariable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_UnresolvedNamedLambdaVariable) ProtoMessage() {}

func (x *Expression_UnresolvedNamedLambdaVariable) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_UnresolvedNamedLambdaVariable.ProtoReflect.Descriptor instead.
func (*Expression_UnresolvedNamedLambdaVariable) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 13}
}

func (x *Expression_UnresolvedNamedLambdaVariable) GetNameParts() []string {
	if x != nil {
		return x.NameParts
	}
	return nil
}

// The window frame
type Expression_Window_WindowFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The type of the frame.
	FrameType Expression_Window_WindowFrame_FrameType `protobuf:"varint,1,opt,name=frame_type,json=frameType,proto3,enum=spark.connect.Expression_Window_WindowFrame_FrameType" json:"frame_type,omitempty"`
	// (Required) The lower bound of the frame.
	Lower *Expression_Window_WindowFrame_FrameBoundary `protobuf:"bytes,2,opt,name=lower,proto3" json:"lower,omitempty"`
	// (Required) The upper bound of the frame.
	Upper *Expression_Window_WindowFrame_FrameBoundary `protobuf:"bytes,3,opt,name=upper,proto3" json:"upper,omitempty"`
}

func (x *Expression_Window_WindowFrame) Reset() {
	*x = Expression_Window_WindowFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[20]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_Window_WindowFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_Window_WindowFrame) ProtoMessage() {}

func (x *Expression_Window_WindowFrame) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[20]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_Window_WindowFrame.ProtoReflect.Descriptor instead.
func (*Expression_Window_WindowFrame) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Expression_Window_WindowFrame) GetFrameType() Expression_Window_WindowFrame_FrameType {
	if x != nil {
		return x.FrameType
	}
	return Expression_Window_WindowFrame_FRAME_TYPE_UNDEFINED
}

func (x *Expression_Window_WindowFrame) GetLower() *Expression_Window_WindowFrame_FrameBoundary {
	if x != nil {
		return x.Lower
	}
	return nil
}

func (x *Expression_Window_WindowFrame) GetUpper() *Expression_Window_WindowFrame_FrameBoundary {
	if x != nil {
		return x.Upper
	}
	return nil
}

type Expression_Window_WindowFrame_FrameBoundary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Boundary:
	//
	//	*Expression_Window_WindowFrame_FrameBoundary_CurrentRow
	//	*Expression_Window_WindowFrame_FrameBoundary_Unbounded
	//	*Expression_Window_WindowFrame_FrameBoundary_Value
	Boundary isExpression_Window_WindowFrame_FrameBoundary_Boundary `protobuf_oneof:"boundary"`
}

func (x *Expression_Window_WindowFrame_FrameBoundary) Reset() {
	*x = Expression_Window_WindowFrame_FrameBoundary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[21]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_Window_WindowFrame_FrameBoundary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_Window_WindowFrame_FrameBoundary) ProtoMessage() {}

func (x *Expression_Window_WindowFrame_FrameBoundary) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[21]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_Window_WindowFrame_FrameBoundary.ProtoReflect.Descriptor instead.
func (*Expression_Window_WindowFrame_FrameBoundary) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

func (m *Expression_Window_WindowFrame_FrameBoundary) GetBoundary() isExpression_Window_WindowFrame_FrameBoundary_Boundary {
	if m != nil {
		return m.Boundary
	}
	return nil
}

func (x *Expression_Window_WindowFrame_FrameBoundary) GetCurrentRow() bool {
	if x, ok := x.GetBoundary().(*Expression_Window_WindowFrame_FrameBoundary_CurrentRow); ok {
		return x.CurrentRow
	}
	return false
}

func (x *Expression_Window_WindowFrame_FrameBoundary) GetUnbounded() bool {
	if x, ok := x.GetBoundary().(*Expression_Window_WindowFrame_FrameBoundary_Unbounded); ok {
		return x.Unbounded
	}
	return false
}

func (x *Expression_Window_WindowFrame_FrameBoundary) GetValue() *Expression {
	if x, ok := x.GetBoundary().(*Expression_Window_WindowFrame_FrameBoundary_Value); ok {
		return x.Value
	}
	return nil
}

type isExpression_Window_WindowFrame_FrameBoundary_Boundary interface {
	isExpression_Window_WindowFrame_FrameBoundary_Boundary()
}

type Expression_Window_WindowFrame_FrameBoundary_CurrentRow struct {
	// CURRENT ROW boundary
	CurrentRow bool `protobuf:"varint,1,opt,name=current_row,json=currentRow,proto3,oneof"`
}

type Expression_Window_WindowFrame_FrameBoundary_Unbounded struct {
	// UNBOUNDED boundary.
	// For lower bound, it will be converted to 'UnboundedPreceding'.
	// for upper bound, it will be converted to 'UnboundedFollowing'.
	Unbounded bool `protobuf:"varint,2,opt,name=unbounded,proto3,oneof"`
}

type Expression_Window_WindowFrame_FrameBoundary_Value struct {
	// This is an expression for future proofing. We are expecting literals on the server side.
	Value *Expression `protobuf:"bytes,3,opt,name=value,proto3,oneof"`
}

func (*Expression_Window_WindowFrame_FrameBoundary_CurrentRow) isExpression_Window_WindowFrame_FrameBoundary_Boundary() {
}

func (*Expression_Window_WindowFrame_FrameBoundary_Unbounded) isExpression_Window_WindowFrame_FrameBoundary_Boundary() {
}

func (*Expression_Window_WindowFrame_FrameBoundary_Value) isExpression_Window_WindowFrame_FrameBoundary_Boundary() {
}

type Expression_Literal_Decimal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the string representation.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// The maximum number of digits allowed in the value.
	// the maximum precision is 38.
	Precision *int32 `protobuf:"varint,2,opt,name=precision,proto3,oneof" json:"precision,omitempty"`
	// declared scale of decimal literal
	Scale *int32 `protobuf:"varint,3,opt,name=scale,proto3,oneof" json:"scale,omitempty"`
}

func (x *Expression_Literal_Decimal) Reset() {
	*x = Expression_Literal_Decimal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[22]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_Literal_Decimal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_Literal_Decimal) ProtoMessage() {}

func (x *Expression_Literal_Decimal) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[22]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_Literal_Decimal.ProtoReflect.Descriptor instead.
func (*Expression_Literal_Decimal) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 3, 0}
}

func (x *Expression_Literal_Decimal) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Expression_Literal_Decimal) GetPrecision() int32 {
	if x != nil && x.Precision != nil {
		return *x.Precision
	}
	return 0
}

func (x *Expression_Literal_Decimal) GetScale() int32 {
	if x != nil && x.Scale != nil {
		return *x.Scale
	}
	return 0
}

type Expression_Literal_CalendarInterval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Months       int32 `protobuf:"varint,1,opt,name=months,proto3" json:"months,omitempty"`
	Days         int32 `protobuf:"varint,2,opt,name=days,proto3" json:"days,omitempty"`
	Microseconds int64 `protobuf:"varint,3,opt,name=microseconds,proto3" json:"microseconds,omitempty"`
}

func (x *Expression_Literal_CalendarInterval) Reset() {
	*x = Expression_Literal_CalendarInterval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[23]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_Literal_CalendarInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_Literal_CalendarInterval) ProtoMessage() {}

func (x *Expression_Literal_CalendarInterval) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[23]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_Literal_CalendarInterval.ProtoReflect.Descriptor instead.
func (*Expression_Literal_CalendarInterval) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 3, 1}
}

func (x *Expression_Literal_CalendarInterval) GetMonths() int32 {
	if x != nil {
		return x.Months
	}
	return 0
}

func (x *Expression_Literal_CalendarInterval) GetDays() int32 {
	if x != nil {
		return x.Days
	}
	return 0
}

func (x *Expression_Literal_CalendarInterval) GetMicroseconds() int64 {
	if x != nil {
		return x.Microseconds
	}
	return 0
}

type Expression_Literal_Array struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ElementType *DataType             `protobuf:"bytes,1,opt,name=element_type,json=elementType,proto3" json:"element_type,omitempty"`
	Elements    []*Expression_Literal `protobuf:"bytes,2,rep,name=elements,proto3" json:"elements,omitempty"`
}

func (x *Expression_Literal_Array) Reset() {
	*x = Expression_Literal_Array{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[24]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_Literal_Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_Literal_Array) ProtoMessage() {}

func (x *Expression_Literal_Array) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[24]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_Literal_Array.ProtoReflect.Descriptor instead.
func (*Expression_Literal_Array) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 3, 2}
}

func (x *Expression_Literal_Array) GetElementType() *DataType {
	if x != nil {
		return x.ElementType
	}
	return nil
}

func (x *Expression_Literal_Array) GetElements() []*Expression_Literal {
	if x != nil {
		return x.Elements
	}
	return nil
}

type Expression_Literal_Map struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyType   *DataType             `protobuf:"bytes,1,opt,name=key_type,json=keyType,proto3" json:"key_type,omitempty"`
	ValueType *DataType             `protobuf:"bytes,2,opt,name=value_type,json=valueType,proto3" json:"value_type,omitempty"`
	Keys      []*Expression_Literal `protobuf:"bytes,3,rep,name=keys,proto3" json:"keys,omitempty"`
	Values    []*Expression_Literal `protobuf:"bytes,4,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Expression_Literal_Map) Reset() {
	*x = Expression_Literal_Map{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[25]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_Literal_Map) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_Literal_Map) ProtoMessage() {}

func (x *Expression_Literal_Map) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[25]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_Literal_Map.ProtoReflect.Descriptor instead.
func (*Expression_Literal_Map) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 3, 3}
}

func (x *Expression_Literal_Map) GetKeyType() *DataType {
	if x != nil {
		return x.KeyType
	}
	return nil
}

func (x *Expression_Literal_Map) GetValueType() *DataType {
	if x != nil {
		return x.ValueType
	}
	return nil
}

func (x *Expression_Literal_Map) GetKeys() []*Expression_Literal {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *Expression_Literal_Map) GetValues() []*Expression_Literal {
	if x != nil {
		return x.Values
	}
	return nil
}

type Expression_Literal_Struct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StructType *DataType             `protobuf:"bytes,1,opt,name=struct_type,json=structType,proto3" json:"struct_type,omitempty"`
	Elements   []*Expression_Literal `protobuf:"bytes,2,rep,name=elements,proto3" json:"elements,omitempty"`
}

func (x *Expression_Literal_Struct) Reset() {
	*x = Expression_Literal_Struct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_expressions_proto_msgTypes[26]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expression_Literal_Struct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression_Literal_Struct) ProtoMessage() {}

func (x *Expression_Literal_Struct) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_expressions_proto_msgTypes[26]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression_Literal_Struct.ProtoReflect.Descriptor instead.
func (*Expression_Literal_Struct) Descriptor() ([]byte, []int) {
	return file_spark_connect_expressions_proto_rawDescGZIP(), []int{0, 3, 4}
}

func (x *Expression_Literal_Struct) GetStructType() *DataType {
	if x != nil {
		return x.StructType
	}
	return nil
}

func (x *Expression_Literal_Struct) GetElements() []*Expression_Literal {
	if x != nil {
		return x.Elements
	}
	return nil
}

var File_spark_connect_expressions_proto protoreflect.FileDescriptor

var file_spark_connect_expressions_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f,
	0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x2b, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x07, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x69, 0x74,
	0x65, 0x72, 0x61, 0x6c, 0x12, 0x62, 0x0a, 0x14, 0x75, 0x6e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x6e,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x48, 0x00, 0x52, 0x13, 0x75, 0x6e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x5f, 0x0a, 0x13, 0x75, 0x6e, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x55, 0x6e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x12, 0x75, 0x6e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a, 0x11, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x48, 0x00, 0x52, 0x10, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x12, 0x53, 0x0a, 0x0f, 0x75, 0x6e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x6e, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x48, 0x00, 0x52, 0x0e, 0x75, 0x6e, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x12, 0x37, 0x0a, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x48, 0x00, 0x52, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x63, 0x61, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x04, 0x63, 0x61, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x10, 0x75, 0x6e, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x6e,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x52, 0x65, 0x67, 0x65, 0x78, 0x48, 0x00, 0x52,
	0x0f, 0x75, 0x6e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x52, 0x65, 0x67, 0x65, 0x78,
	0x12, 0x44, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x09, 0x73, 0x6f, 0x72,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x0f, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61,
	0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x61, 0x6d, 0x62, 0x64,
	0x61, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x6c, 0x61, 0x6d,
	0x62, 0x64, 0x61, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x06, 0x77,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x48, 0x00, 0x52,
	0x06, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x6c, 0x0a, 0x18, 0x75, 0x6e, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x6e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x16, 0x75,
	0x6e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4d, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x20, 0x75, 0x6e, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61,
	0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x6e, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x1d, 0x75, 0x6e, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x4c, 0x61, 0x6d, 0x62, 0x64,
	0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x7e, 0x0a, 0x23, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x1f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65,
	0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x0d, 0x63, 0x61, 0x6c,
	0x6c, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a,
	0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xe7, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x8f, 0x06, 0x0a, 0x06, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12,
	0x42, 0x0a, 0x0f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x42, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x09,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4b, 0x0a, 0x0a, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e,
	0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x53, 0x70, 0x65, 0x63, 0x1a, 0xed, 0x03, 0x0a, 0x0b, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x55, 0x0a, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x57, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x50, 0x0a,
	0x05, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x57,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x52, 0x05, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12,
	0x50, 0x0a, 0x05, 0x75, 0x70, 0x70, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x52, 0x05, 0x75, 0x70, 0x70, 0x65,
	0x72, 0x1a, 0x91, 0x01, 0x0a, 0x0d, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72,
	0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x52, 0x6f, 0x77, 0x12, 0x1e, 0x0a, 0x09, 0x75, 0x6e, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x75, 0x6e, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x72, 0x79, 0x22, 0x4f, 0x0a, 0x09, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e,
	0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x4f, 0x57, 0x10, 0x01,
	0x12, 0x14, 0x0a, 0x10, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x10, 0x02, 0x1a, 0xa9, 0x03, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x05, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x4f, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x6f,
	0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x0d, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x0c, 0x6e, 0x75, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x6c, 0x0a,
	0x0d, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x1a, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c,
	0x0a, 0x18, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x41, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19,
	0x53, 0x4f, 0x52, 0x54, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44,
	0x45, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x22, 0x55, 0x0a, 0x0c, 0x4e,
	0x75, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x16, 0x53,
	0x4f, 0x52, 0x54, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x4f, 0x52, 0x54, 0x5f,
	0x4e, 0x55, 0x4c, 0x4c, 0x53, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x53, 0x5f, 0x4c, 0x41, 0x53, 0x54,
	0x10, 0x02, 0x1a, 0x91, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x65,
	0x78, 0x70, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x65, 0x78, 0x70, 0x72, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x48, 0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x08, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x74,
	0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x74,
	0x6f, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x9b, 0x0c, 0x0a, 0x07, 0x4c, 0x69, 0x74, 0x65, 0x72,
	0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x04, 0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x75, 0x6c,
	0x6c, 0x12, 0x18, 0x0a, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x07, 0x62,
	0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07,
	0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x14, 0x0a, 0x04, 0x62, 0x79, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x62, 0x79, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12,
	0x18, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x48,
	0x00, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x45, 0x0a, 0x07, 0x64, 0x65, 0x63,
	0x69, 0x6d, 0x61, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x44, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c,
	0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1e, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x25, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6e, 0x74,
	0x7a, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x4e, 0x74, 0x7a, 0x12, 0x61, 0x0a, 0x11, 0x63, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69,
	0x74, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x10, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x13, 0x79, 0x65,
	0x61, 0x72, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x11, 0x79, 0x65, 0x61, 0x72, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x11,
	0x64, 0x61, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0f, 0x64, 0x61, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x3f, 0x0a, 0x05, 0x61, 0x72,
	0x72, 0x61, 0x79, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x41, 0x72, 0x72,
	0x61, 0x79, 0x48, 0x00, 0x52, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x12, 0x39, 0x0a, 0x03, 0x6d,
	0x61, 0x70, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4d, 0x61, 0x70, 0x48,
	0x00, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x12, 0x42, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x75, 0x0a, 0x07, 0x44, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x70,
	0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52,
	0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x72,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x1a, 0x62, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x79,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x1a, 0x82, 0x01, 0x0a, 0x05, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12,
	0x3a, 0x0a, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c,
	0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0xe3, 0x01, 0x0a, 0x03, 0x4d,
	0x61, 0x70, 0x12, 0x32, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35,
	0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x52,
	0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x39, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x1a, 0x81, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x1a, 0x70, 0x0a, 0x13, 0x55, 0x6e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x75,
	0x6e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x75, 0x6e, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x07,
	0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70,
	0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x1a, 0xcc, 0x01, 0x0a, 0x12, 0x55, 0x6e, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a,
	0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x73, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x69, 0x73, 0x44, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x12, 0x37, 0x0a, 0x18,
	0x69, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x5f,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15,
	0x69, 0x73, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x32, 0x0a, 0x10, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x52, 0x0a, 0x0e, 0x55, 0x6e, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x12, 0x2c, 0x0a, 0x0f, 0x75,
	0x6e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x75, 0x6e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x88, 0x01, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x75, 0x6e,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x1a, 0x56, 0x0a,
	0x0f, 0x55, 0x6e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x52, 0x65, 0x67, 0x65, 0x78,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x07, 0x70,
	0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06,
	0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x5f, 0x69, 0x64, 0x1a, 0x84, 0x01, 0x0a, 0x16, 0x55, 0x6e, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x2f, 0x0a, 0x05, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x63, 0x68, 0x69, 0x6c,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0xbb, 0x01, 0x0a,
	0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x46, 0x0a,
	0x11, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x10, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x78, 0x0a, 0x05, 0x41, 0x6c,
	0x69, 0x61, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x65, 0x78, 0x70, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x65, 0x78,
	0x70, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x9e, 0x01, 0x0a, 0x0e, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55,
	0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x6e, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x4c, 0x61, 0x6d, 0x62,
	0x64, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x3e, 0x0a, 0x1d, 0x55, 0x6e, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x70,
	0x61, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x50, 0x61, 0x72, 0x74, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x22, 0xec, 0x02, 0x0a, 0x1f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x64,
	0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x12, 0x37, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x79,
	0x74, 0x68, 0x6f, 0x6e, 0x5f, 0x75, 0x64, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x50,
	0x79, 0x74, 0x68, 0x6f, 0x6e, 0x55, 0x44, 0x46, 0x48, 0x00, 0x52, 0x09, 0x70, 0x79, 0x74, 0x68,
	0x6f, 0x6e, 0x55, 0x64, 0x66, 0x12, 0x49, 0x0a, 0x10, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f,
	0x73, 0x63, 0x61, 0x6c, 0x61, 0x5f, 0x75, 0x64, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x55, 0x44, 0x46, 0x48, 0x00,
	0x52, 0x0e, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x55, 0x64, 0x66,
	0x12, 0x33, 0x0a, 0x08, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x75, 0x64, 0x66, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x4a, 0x61, 0x76, 0x61, 0x55, 0x44, 0x46, 0x48, 0x00, 0x52, 0x07, 0x6a, 0x61,
	0x76, 0x61, 0x55, 0x64, 0x66, 0x42, 0x0a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x9b, 0x01, 0x0a, 0x09, 0x50, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x55, 0x44, 0x46, 0x12,
	0x38, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x76, 0x61,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x76,
	0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x22,
	0xb8, 0x01, 0x0a, 0x0e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x55,
	0x44, 0x46, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x37, 0x0a, 0x0a,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x07, 0x4a,
	0x61, 0x76, 0x61, 0x55, 0x44, 0x46, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x6c, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x42, 0x36, 0x0a, 0x1e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spark_connect_expressions_proto_rawDescOnce sync.Once
	file_spark_connect_expressions_proto_rawDescData = file_spark_connect_expressions_proto_rawDesc
)

func file_spark_connect_expressions_proto_rawDescGZIP() []byte {
	file_spark_connect_expressions_proto_rawDescOnce.Do(func() {
		file_spark_connect_expressions_proto_rawDescData = protoimpl.X.CompressGZIP(file_spark_connect_expressions_proto_rawDescData)
	})
	return file_spark_connect_expressions_proto_rawDescData
}

var file_spark_connect_expressions_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_spark_connect_expressions_proto_msgTypes = make([]protoimpl.MessageInfo, 27)
var file_spark_connect_expressions_proto_goTypes = []interface{}{
	(Expression_Window_WindowFrame_FrameType)(0),        // 0: spark.connect.Expression.Window.WindowFrame.FrameType
	(Expression_SortOrder_SortDirection)(0),             // 1: spark.connect.Expression.SortOrder.SortDirection
	(Expression_SortOrder_NullOrdering)(0),              // 2: spark.connect.Expression.SortOrder.NullOrdering
	(*Expression)(nil),                                  // 3: spark.connect.Expression
	(*CommonInlineUserDefinedFunction)(nil),             // 4: spark.connect.CommonInlineUserDefinedFunction
	(*PythonUDF)(nil),                                   // 5: spark.connect.PythonUDF
	(*ScalarScalaUDF)(nil),                              // 6: spark.connect.ScalarScalaUDF
	(*JavaUDF)(nil),                                     // 7: spark.connect.JavaUDF
	(*CallFunction)(nil),                                // 8: spark.connect.CallFunction
	(*Expression_Window)(nil),                           // 9: spark.connect.Expression.Window
	(*Expression_SortOrder)(nil),                        // 10: spark.connect.Expression.SortOrder
	(*Expression_Cast)(nil),                             // 11: spark.connect.Expression.Cast
	(*Expression_Literal)(nil),                          // 12: spark.connect.Expression.Literal
	(*Expression_UnresolvedAttribute)(nil),              // 13: spark.connect.Expression.UnresolvedAttribute
	(*Expression_UnresolvedFunction)(nil),               // 14: spark.connect.Expression.UnresolvedFunction
	(*Expression_ExpressionString)(nil),                 // 15: spark.connect.Expression.ExpressionString
	(*Expression_UnresolvedStar)(nil),                   // 16: spark.connect.Expression.UnresolvedStar
	(*Expression_UnresolvedRegex)(nil),                  // 17: spark.connect.Expression.UnresolvedRegex
	(*Expression_UnresolvedExtractValue)(nil),           // 18: spark.connect.Expression.UnresolvedExtractValue
	(*Expression_UpdateFields)(nil),                     // 19: spark.connect.Expression.UpdateFields
	(*Expression_Alias)(nil),                            // 20: spark.connect.Expression.Alias
	(*Expression_LambdaFunction)(nil),                   // 21: spark.connect.Expression.LambdaFunction
	(*Expression_UnresolvedNamedLambdaVariable)(nil),    // 22: spark.connect.Expression.UnresolvedNamedLambdaVariable
	(*Expression_Window_WindowFrame)(nil),               // 23: spark.connect.Expression.Window.WindowFrame
	(*Expression_Window_WindowFrame_FrameBoundary)(nil), // 24: spark.connect.Expression.Window.WindowFrame.FrameBoundary
	(*Expression_Literal_Decimal)(nil),                  // 25: spark.connect.Expression.Literal.Decimal
	(*Expression_Literal_CalendarInterval)(nil),         // 26: spark.connect.Expression.Literal.CalendarInterval
	(*Expression_Literal_Array)(nil),                    // 27: spark.connect.Expression.Literal.Array
	(*Expression_Literal_Map)(nil),                      // 28: spark.connect.Expression.Literal.Map
	(*Expression_Literal_Struct)(nil),                   // 29: spark.connect.Expression.Literal.Struct
	(*anypb.Any)(nil),                                   // 30: google.protobuf.Any
	(*DataType)(nil),                                    // 31: spark.connect.DataType
}
var file_spark_connect_expressions_proto_depIdxs = []int32{
	12, // 0: spark.connect.Expression.literal:type_name -> spark.connect.Expression.Literal
	13, // 1: spark.connect.Expression.unresolved_attribute:type_name -> spark.connect.Expression.UnresolvedAttribute
	14, // 2: spark.connect.Expression.unresolved_function:type_name -> spark.connect.Expression.UnresolvedFunction
	15, // 3: spark.connect.Expression.expression_string:type_name -> spark.connect.Expression.ExpressionString
	16, // 4: spark.connect.Expression.unresolved_star:type_name -> spark.connect.Expression.UnresolvedStar
	20, // 5: spark.connect.Expression.alias:type_name -> spark.connect.Expression.Alias
	11, // 6: spark.connect.Expression.cast:type_name -> spark.connect.Expression.Cast
	17, // 7: spark.connect.Expression.unresolved_regex:type_name -> spark.connect.Expression.UnresolvedRegex
	10, // 8: spark.connect.Expression.sort_order:type_name -> spark.connect.Expression.SortOrder
	21, // 9: spark.connect.Expression.lambda_function:type_name -> spark.connect.Expression.LambdaFunction
	9,  // 10: spark.connect.Expression.window:type_name -> spark.connect.Expression.Window
	18, // 11: spark.connect.Expression.unresolved_extract_value:type_name -> spark.connect.Expression.UnresolvedExtractValue
	19, // 12: spark.connect.Expression.update_fields:type_name -> spark.connect.Expression.UpdateFields
	22, // 13: spark.connect.Expression.unresolved_named_lambda_variable:type_name -> spark.connect.Expression.UnresolvedNamedLambdaVariable
	4,  // 14: spark.connect.Expression.common_inline_user_defined_function:type_name -> spark.connect.CommonInlineUserDefinedFunction
	8,  // 15: spark.connect.Expression.call_function:type_name -> spark.connect.CallFunction
	30, // 16: spark.connect.Expression.extension:type_name -> google.protobuf.Any
	3,  // 17: spark.connect.CommonInlineUserDefinedFunction.arguments:type_name -> spark.connect.Expression
	5,  // 18: spark.connect.CommonInlineUserDefinedFunction.python_udf:type_name -> spark.connect.PythonUDF
	6,  // 19: spark.connect.CommonInlineUserDefinedFunction.scalar_scala_udf:type_name -> spark.connect.ScalarScalaUDF
	7,  // 20: spark.connect.CommonInlineUserDefinedFunction.java_udf:type_name -> spark.connect.JavaUDF
	31, // 21: spark.connect.PythonUDF.output_type:type_name -> spark.connect.DataType
	31, // 22: spark.connect.ScalarScalaUDF.inputTypes:type_name -> spark.connect.DataType
	31, // 23: spark.connect.ScalarScalaUDF.outputType:type_name -> spark.connect.DataType
	31, // 24: spark.connect.JavaUDF.output_type:type_name -> spark.connect.DataType
	3,  // 25: spark.connect.CallFunction.arguments:type_name -> spark.connect.Expression
	3,  // 26: spark.connect.Expression.Window.window_function:type_name -> spark.connect.Expression
	3,  // 27: spark.connect.Expression.Window.partition_spec:type_name -> spark.connect.Expression
	10, // 28: spark.connect.Expression.Window.order_spec:type_name -> spark.connect.Expression.SortOrder
	23, // 29: spark.connect.Expression.Window.frame_spec:type_name -> spark.connect.Expression.Window.WindowFrame
	3,  // 30: spark.connect.Expression.SortOrder.child:type_name -> spark.connect.Expression
	1,  // 31: spark.connect.Expression.SortOrder.direction:type_name -> spark.connect.Expression.SortOrder.SortDirection
	2,  // 32: spark.connect.Expression.SortOrder.null_ordering:type_name -> spark.connect.Expression.SortOrder.NullOrdering
	3,  // 33: spark.connect.Expression.Cast.expr:type_name -> spark.connect.Expression
	31, // 34: spark.connect.Expression.Cast.type:type_name -> spark.connect.DataType
	31, // 35: spark.connect.Expression.Literal.null:type_name -> spark.connect.DataType
	25, // 36: spark.connect.Expression.Literal.decimal:type_name -> spark.connect.Expression.Literal.Decimal
	26, // 37: spark.connect.Expression.Literal.calendar_interval:type_name -> spark.connect.Expression.Literal.CalendarInterval
	27, // 38: spark.connect.Expression.Literal.array:type_name -> spark.connect.Expression.Literal.Array
	28, // 39: spark.connect.Expression.Literal.map:type_name -> spark.connect.Expression.Literal.Map
	29, // 40: spark.connect.Expression.Literal.struct:type_name -> spark.connect.Expression.Literal.Struct
	3,  // 41: spark.connect.Expression.UnresolvedFunction.arguments:type_name -> spark.connect.Expression
	3,  // 42: spark.connect.Expression.UnresolvedExtractValue.child:type_name -> spark.connect.Expression
	3,  // 43: spark.connect.Expression.UnresolvedExtractValue.extraction:type_name -> spark.connect.Expression
	3,  // 44: spark.connect.Expression.UpdateFields.struct_expression:type_name -> spark.connect.Expression
	3,  // 45: spark.connect.Expression.UpdateFields.value_expression:type_name -> spark.connect.Expression
	3,  // 46: spark.connect.Expression.Alias.expr:type_name -> spark.connect.Expression
	3,  // 47: spark.connect.Expression.LambdaFunction.function:type_name -> spark.connect.Expression
	22, // 48: spark.connect.Expression.LambdaFunction.arguments:type_name -> spark.connect.Expression.UnresolvedNamedLambdaVariable
	0,  // 49: spark.connect.Expression.Window.WindowFrame.frame_type:type_name -> spark.connect.Expression.Window.WindowFrame.FrameType
	24, // 50: spark.connect.Expression.Window.WindowFrame.lower:type_name -> spark.connect.Expression.Window.WindowFrame.FrameBoundary
	24, // 51: spark.connect.Expression.Window.WindowFrame.upper:type_name -> spark.connect.Expression.Window.WindowFrame.FrameBoundary
	3,  // 52: spark.connect.Expression.Window.WindowFrame.FrameBoundary.value:type_name -> spark.connect.Expression
	31, // 53: spark.connect.Expression.Literal.Array.element_type:type_name -> spark.connect.DataType
	12, // 54: spark.connect.Expression.Literal.Array.elements:type_name -> spark.connect.Expression.Literal
	31, // 55: spark.connect.Expression.Literal.Map.key_type:type_name -> spark.connect.DataType
	31, // 56: spark.connect.Expression.Literal.Map.value_type:type_name -> spark.connect.DataType
	12, // 57: spark.connect.Expression.Literal.Map.keys:type_name -> spark.connect.Expression.Literal
	12, // 58: spark.connect.Expression.Literal.Map.values:type_name -> spark.connect.Expression.Literal
	31, // 59: spark.connect.Expression.Literal.Struct.struct_type:type_name -> spark.connect.DataType
	12, // 60: spark.connect.Expression.Literal.Struct.elements:type_name -> spark.connect.Expression.Literal
	61, // [61:61] is the sub-list for method output_type
	61, // [61:61] is the sub-list for method input_type
	61, // [61:61] is the sub-list for extension type_name
	61, // [61:61] is the sub-list for extension extendee
	0,  // [0:61] is the sub-list for field type_name
}

func init() { file_spark_connect_expressions_proto_init() }
func file_spark_connect_expressions_proto_init() {
	if File_spark_connect_expressions_proto != nil {
		return
	}
	file_spark_connect_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spark_connect_expressions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression); i {
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
		file_spark_connect_expressions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonInlineUserDefinedFunction); i {
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
		file_spark_connect_expressions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PythonUDF); i {
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
		file_spark_connect_expressions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScalarScalaUDF); i {
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
		file_spark_connect_expressions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JavaUDF); i {
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
		file_spark_connect_expressions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallFunction); i {
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
		file_spark_connect_expressions_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_Window); i {
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
		file_spark_connect_expressions_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_SortOrder); i {
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
		file_spark_connect_expressions_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_Cast); i {
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
		file_spark_connect_expressions_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_Literal); i {
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
		file_spark_connect_expressions_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_UnresolvedAttribute); i {
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
		file_spark_connect_expressions_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_UnresolvedFunction); i {
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
		file_spark_connect_expressions_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_ExpressionString); i {
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
		file_spark_connect_expressions_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_UnresolvedStar); i {
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
		file_spark_connect_expressions_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_UnresolvedRegex); i {
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
		file_spark_connect_expressions_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_UnresolvedExtractValue); i {
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
		file_spark_connect_expressions_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_UpdateFields); i {
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
		file_spark_connect_expressions_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_Alias); i {
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
		file_spark_connect_expressions_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_LambdaFunction); i {
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
		file_spark_connect_expressions_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_UnresolvedNamedLambdaVariable); i {
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
		file_spark_connect_expressions_proto_msgTypes[20].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_Window_WindowFrame); i {
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
		file_spark_connect_expressions_proto_msgTypes[21].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_Window_WindowFrame_FrameBoundary); i {
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
		file_spark_connect_expressions_proto_msgTypes[22].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_Literal_Decimal); i {
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
		file_spark_connect_expressions_proto_msgTypes[23].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_Literal_CalendarInterval); i {
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
		file_spark_connect_expressions_proto_msgTypes[24].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_Literal_Array); i {
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
		file_spark_connect_expressions_proto_msgTypes[25].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_Literal_Map); i {
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
		file_spark_connect_expressions_proto_msgTypes[26].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expression_Literal_Struct); i {
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
	file_spark_connect_expressions_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Expression_Literal_)(nil),
		(*Expression_UnresolvedAttribute_)(nil),
		(*Expression_UnresolvedFunction_)(nil),
		(*Expression_ExpressionString_)(nil),
		(*Expression_UnresolvedStar_)(nil),
		(*Expression_Alias_)(nil),
		(*Expression_Cast_)(nil),
		(*Expression_UnresolvedRegex_)(nil),
		(*Expression_SortOrder_)(nil),
		(*Expression_LambdaFunction_)(nil),
		(*Expression_Window_)(nil),
		(*Expression_UnresolvedExtractValue_)(nil),
		(*Expression_UpdateFields_)(nil),
		(*Expression_UnresolvedNamedLambdaVariable_)(nil),
		(*Expression_CommonInlineUserDefinedFunction)(nil),
		(*Expression_CallFunction)(nil),
		(*Expression_Extension)(nil),
	}
	file_spark_connect_expressions_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CommonInlineUserDefinedFunction_PythonUdf)(nil),
		(*CommonInlineUserDefinedFunction_ScalarScalaUdf)(nil),
		(*CommonInlineUserDefinedFunction_JavaUdf)(nil),
	}
	file_spark_connect_expressions_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_spark_connect_expressions_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*Expression_Cast_Type)(nil),
		(*Expression_Cast_TypeStr)(nil),
	}
	file_spark_connect_expressions_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*Expression_Literal_Null)(nil),
		(*Expression_Literal_Binary)(nil),
		(*Expression_Literal_Boolean)(nil),
		(*Expression_Literal_Byte)(nil),
		(*Expression_Literal_Short)(nil),
		(*Expression_Literal_Integer)(nil),
		(*Expression_Literal_Long)(nil),
		(*Expression_Literal_Float)(nil),
		(*Expression_Literal_Double)(nil),
		(*Expression_Literal_Decimal_)(nil),
		(*Expression_Literal_String_)(nil),
		(*Expression_Literal_Date)(nil),
		(*Expression_Literal_Timestamp)(nil),
		(*Expression_Literal_TimestampNtz)(nil),
		(*Expression_Literal_CalendarInterval_)(nil),
		(*Expression_Literal_YearMonthInterval)(nil),
		(*Expression_Literal_DayTimeInterval)(nil),
		(*Expression_Literal_Array_)(nil),
		(*Expression_Literal_Map_)(nil),
		(*Expression_Literal_Struct_)(nil),
	}
	file_spark_connect_expressions_proto_msgTypes[10].OneofWrappers = []interface{}{}
	file_spark_connect_expressions_proto_msgTypes[13].OneofWrappers = []interface{}{}
	file_spark_connect_expressions_proto_msgTypes[14].OneofWrappers = []interface{}{}
	file_spark_connect_expressions_proto_msgTypes[17].OneofWrappers = []interface{}{}
	file_spark_connect_expressions_proto_msgTypes[21].OneofWrappers = []interface{}{
		(*Expression_Window_WindowFrame_FrameBoundary_CurrentRow)(nil),
		(*Expression_Window_WindowFrame_FrameBoundary_Unbounded)(nil),
		(*Expression_Window_WindowFrame_FrameBoundary_Value)(nil),
	}
	file_spark_connect_expressions_proto_msgTypes[22].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spark_connect_expressions_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   27,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spark_connect_expressions_proto_goTypes,
		DependencyIndexes: file_spark_connect_expressions_proto_depIdxs,
		EnumInfos:         file_spark_connect_expressions_proto_enumTypes,
		MessageInfos:      file_spark_connect_expressions_proto_msgTypes,
	}.Build()
	File_spark_connect_expressions_proto = out.File
	file_spark_connect_expressions_proto_rawDesc = nil
	file_spark_connect_expressions_proto_goTypes = nil
	file_spark_connect_expressions_proto_depIdxs = nil
}
