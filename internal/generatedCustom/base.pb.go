package generatedCustom

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generatedCustom code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Plan explanation mode.
type AnalyzePlanRequest_Explain_ExplainMode int32

const (
	AnalyzePlanRequest_Explain_EXPLAIN_MODE_UNSPECIFIED AnalyzePlanRequest_Explain_ExplainMode = 0
	// Generates only physical plan.
	AnalyzePlanRequest_Explain_EXPLAIN_MODE_SIMPLE AnalyzePlanRequest_Explain_ExplainMode = 1
	// Generates parsed logical plan, analyzed logical plan, optimized logical plan and physical plan.
	// Parsed Logical plan is a unresolved plan that extracted from the query. Analyzed logical plans
	// transforms which translates unresolvedAttribute and unresolvedRelation into fully typed objects.
	// The optimized logical plan transforms through a set of optimization rules, resulting in the
	// physical plan.
	AnalyzePlanRequest_Explain_EXPLAIN_MODE_EXTENDED AnalyzePlanRequest_Explain_ExplainMode = 2
	// Generates code for the statement, if any and a physical plan.
	AnalyzePlanRequest_Explain_EXPLAIN_MODE_CODEGEN AnalyzePlanRequest_Explain_ExplainMode = 3
	// If plan node statistics are available, generates a logical plan and also the statistics.
	AnalyzePlanRequest_Explain_EXPLAIN_MODE_COST AnalyzePlanRequest_Explain_ExplainMode = 4
	// Generates a physical plan outline and also node details.
	AnalyzePlanRequest_Explain_EXPLAIN_MODE_FORMATTED AnalyzePlanRequest_Explain_ExplainMode = 5
)

// Enum value maps for AnalyzePlanRequest_Explain_ExplainMode.
var (
	AnalyzePlanRequest_Explain_ExplainMode_name = map[int32]string{
		0: "EXPLAIN_MODE_UNSPECIFIED",
		1: "EXPLAIN_MODE_SIMPLE",
		2: "EXPLAIN_MODE_EXTENDED",
		3: "EXPLAIN_MODE_CODEGEN",
		4: "EXPLAIN_MODE_COST",
		5: "EXPLAIN_MODE_FORMATTED",
	}
	AnalyzePlanRequest_Explain_ExplainMode_value = map[string]int32{
		"EXPLAIN_MODE_UNSPECIFIED": 0,
		"EXPLAIN_MODE_SIMPLE":      1,
		"EXPLAIN_MODE_EXTENDED":    2,
		"EXPLAIN_MODE_CODEGEN":     3,
		"EXPLAIN_MODE_COST":        4,
		"EXPLAIN_MODE_FORMATTED":   5,
	}
)

func (x AnalyzePlanRequest_Explain_ExplainMode) Enum() *AnalyzePlanRequest_Explain_ExplainMode {
	p := new(AnalyzePlanRequest_Explain_ExplainMode)
	*p = x
	return p
}

func (x AnalyzePlanRequest_Explain_ExplainMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AnalyzePlanRequest_Explain_ExplainMode) Descriptor() protoreflect.EnumDescriptor {
	return file_spark_connect_base_proto_enumTypes[0].Descriptor()
}

func (AnalyzePlanRequest_Explain_ExplainMode) Type() protoreflect.EnumType {
	return &file_spark_connect_base_proto_enumTypes[0]
}

func (x AnalyzePlanRequest_Explain_ExplainMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AnalyzePlanRequest_Explain_ExplainMode.Descriptor instead.
func (AnalyzePlanRequest_Explain_ExplainMode) EnumDescriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{2, 1, 0}
}

type InterruptRequest_InterruptType int32

const (
	InterruptRequest_INTERRUPT_TYPE_UNSPECIFIED InterruptRequest_InterruptType = 0
	// Interrupt all running executions within the session with the provided session_id.
	InterruptRequest_INTERRUPT_TYPE_ALL InterruptRequest_InterruptType = 1
	// Interrupt all running executions within the session with the provided operation_tag.
	InterruptRequest_INTERRUPT_TYPE_TAG InterruptRequest_InterruptType = 2
	// Interrupt the running execution within the session with the provided operation_id.
	InterruptRequest_INTERRUPT_TYPE_OPERATION_ID InterruptRequest_InterruptType = 3
)

// Enum value maps for InterruptRequest_InterruptType.
var (
	InterruptRequest_InterruptType_name = map[int32]string{
		0: "INTERRUPT_TYPE_UNSPECIFIED",
		1: "INTERRUPT_TYPE_ALL",
		2: "INTERRUPT_TYPE_TAG",
		3: "INTERRUPT_TYPE_OPERATION_ID",
	}
	InterruptRequest_InterruptType_value = map[string]int32{
		"INTERRUPT_TYPE_UNSPECIFIED":  0,
		"INTERRUPT_TYPE_ALL":          1,
		"INTERRUPT_TYPE_TAG":          2,
		"INTERRUPT_TYPE_OPERATION_ID": 3,
	}
)

func (x InterruptRequest_InterruptType) Enum() *InterruptRequest_InterruptType {
	p := new(InterruptRequest_InterruptType)
	*p = x
	return p
}

func (x InterruptRequest_InterruptType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InterruptRequest_InterruptType) Descriptor() protoreflect.EnumDescriptor {
	return file_spark_connect_base_proto_enumTypes[1].Descriptor()
}

func (InterruptRequest_InterruptType) Type() protoreflect.EnumType {
	return &file_spark_connect_base_proto_enumTypes[1]
}

func (x InterruptRequest_InterruptType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InterruptRequest_InterruptType.Descriptor instead.
func (InterruptRequest_InterruptType) EnumDescriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{13, 0}
}

// A [[Plan]] is the structure that carries the runtime information for the execution from the
// client to the server. A [[Plan]] can either be of the type [[Relation]] which is a reference
// to the underlying logical plan or it can be of the [[Command]] type that is used to execute
// commands on the server.
type Plan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to OpType:
	//
	//	*Plan_Root
	//	*Plan_Command
	OpType isPlan_OpType `protobuf_oneof:"op_type"`
}

func (x *Plan) Reset() {
	*x = Plan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan) ProtoMessage() {}

func (x *Plan) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plan.ProtoReflect.Descriptor instead.
func (*Plan) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{0}
}

func (m *Plan) GetOpType() isPlan_OpType {
	if m != nil {
		return m.OpType
	}
	return nil
}

func (x *Plan) GetRoot() *Relation {
	if x, ok := x.GetOpType().(*Plan_Root); ok {
		return x.Root
	}
	return nil
}

func (x *Plan) GetCommand() *Command {
	if x, ok := x.GetOpType().(*Plan_Command); ok {
		return x.Command
	}
	return nil
}

type isPlan_OpType interface {
	isPlan_OpType()
}

type Plan_Root struct {
	Root *Relation `protobuf:"bytes,1,opt,name=root,proto3,oneof"`
}

type Plan_Command struct {
	Command *Command `protobuf:"bytes,2,opt,name=command,proto3,oneof"`
}

func (*Plan_Root) isPlan_OpType() {}

func (*Plan_Command) isPlan_OpType() {}

// User Context is used to refer to one particular user session that is executing
// queries in the backend.
type UserContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// To extend the existing user context message that is used to identify incoming requests,
	// Spark Connect leverages the Any protobuf type that can be used to inject arbitrary other
	// messages into this message. Extensions are stored as a `repeated` type to be able to
	// handle multiple active extensions.
	Extensions []*anypb.Any `protobuf:"bytes,999,rep,name=extensions,proto3" json:"extensions,omitempty"`
}

func (x *UserContext) Reset() {
	*x = UserContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserContext) ProtoMessage() {}

func (x *UserContext) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserContext.ProtoReflect.Descriptor instead.
func (*UserContext) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{1}
}

func (x *UserContext) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserContext) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserContext) GetExtensions() []*anypb.Any {
	if x != nil {
		return x.Extensions
	}
	return nil
}

// Request to perform plan analyze, optionally to explain the plan.
type AnalyzePlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	//
	// The session_id specifies a spark session for a user id (which is specified
	// by user_context.user_id). The session_id is set by the client to be able to
	// collate streaming responses from different queries within the dedicated session.
	// The id should be an UUID string of the format `00112233-4455-6677-8899-aabbccddeeff`
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// (Required) User context
	UserContext *UserContext `protobuf:"bytes,2,opt,name=user_context,json=userContext,proto3" json:"user_context,omitempty"`
	// Provides optional information about the client sending the request. This field
	// can be used for language or version specific information and is only intended for
	// logging purposes and will not be interpreted by the server.
	ClientType *string `protobuf:"bytes,3,opt,name=client_type,json=clientType,proto3,oneof" json:"client_type,omitempty"`
	// Types that are assignable to Analyze:
	//
	//	*AnalyzePlanRequest_Schema_
	//	*AnalyzePlanRequest_Explain_
	//	*AnalyzePlanRequest_TreeString_
	//	*AnalyzePlanRequest_IsLocal_
	//	*AnalyzePlanRequest_IsStreaming_
	//	*AnalyzePlanRequest_InputFiles_
	//	*AnalyzePlanRequest_SparkVersion_
	//	*AnalyzePlanRequest_DdlParse
	//	*AnalyzePlanRequest_SameSemantics_
	//	*AnalyzePlanRequest_SemanticHash_
	//	*AnalyzePlanRequest_Persist_
	//	*AnalyzePlanRequest_Unpersist_
	//	*AnalyzePlanRequest_GetStorageLevel_
	Analyze isAnalyzePlanRequest_Analyze `protobuf_oneof:"analyze"`
}

func (x *AnalyzePlanRequest) Reset() {
	*x = AnalyzePlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanRequest) ProtoMessage() {}

func (x *AnalyzePlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanRequest.ProtoReflect.Descriptor instead.
func (*AnalyzePlanRequest) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{2}
}

func (x *AnalyzePlanRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *AnalyzePlanRequest) GetUserContext() *UserContext {
	if x != nil {
		return x.UserContext
	}
	return nil
}

func (x *AnalyzePlanRequest) GetClientType() string {
	if x != nil && x.ClientType != nil {
		return *x.ClientType
	}
	return ""
}

func (m *AnalyzePlanRequest) GetAnalyze() isAnalyzePlanRequest_Analyze {
	if m != nil {
		return m.Analyze
	}
	return nil
}

func (x *AnalyzePlanRequest) GetSchema() *AnalyzePlanRequest_Schema {
	if x, ok := x.GetAnalyze().(*AnalyzePlanRequest_Schema_); ok {
		return x.Schema
	}
	return nil
}

func (x *AnalyzePlanRequest) GetExplain() *AnalyzePlanRequest_Explain {
	if x, ok := x.GetAnalyze().(*AnalyzePlanRequest_Explain_); ok {
		return x.Explain
	}
	return nil
}

func (x *AnalyzePlanRequest) GetTreeString() *AnalyzePlanRequest_TreeString {
	if x, ok := x.GetAnalyze().(*AnalyzePlanRequest_TreeString_); ok {
		return x.TreeString
	}
	return nil
}

func (x *AnalyzePlanRequest) GetIsLocal() *AnalyzePlanRequest_IsLocal {
	if x, ok := x.GetAnalyze().(*AnalyzePlanRequest_IsLocal_); ok {
		return x.IsLocal
	}
	return nil
}

func (x *AnalyzePlanRequest) GetIsStreaming() *AnalyzePlanRequest_IsStreaming {
	if x, ok := x.GetAnalyze().(*AnalyzePlanRequest_IsStreaming_); ok {
		return x.IsStreaming
	}
	return nil
}

func (x *AnalyzePlanRequest) GetInputFiles() *AnalyzePlanRequest_InputFiles {
	if x, ok := x.GetAnalyze().(*AnalyzePlanRequest_InputFiles_); ok {
		return x.InputFiles
	}
	return nil
}

func (x *AnalyzePlanRequest) GetSparkVersion() *AnalyzePlanRequest_SparkVersion {
	if x, ok := x.GetAnalyze().(*AnalyzePlanRequest_SparkVersion_); ok {
		return x.SparkVersion
	}
	return nil
}

func (x *AnalyzePlanRequest) GetDdlParse() *AnalyzePlanRequest_DDLParse {
	if x, ok := x.GetAnalyze().(*AnalyzePlanRequest_DdlParse); ok {
		return x.DdlParse
	}
	return nil
}

func (x *AnalyzePlanRequest) GetSameSemantics() *AnalyzePlanRequest_SameSemantics {
	if x, ok := x.GetAnalyze().(*AnalyzePlanRequest_SameSemantics_); ok {
		return x.SameSemantics
	}
	return nil
}

func (x *AnalyzePlanRequest) GetSemanticHash() *AnalyzePlanRequest_SemanticHash {
	if x, ok := x.GetAnalyze().(*AnalyzePlanRequest_SemanticHash_); ok {
		return x.SemanticHash
	}
	return nil
}

func (x *AnalyzePlanRequest) GetPersist() *AnalyzePlanRequest_Persist {
	if x, ok := x.GetAnalyze().(*AnalyzePlanRequest_Persist_); ok {
		return x.Persist
	}
	return nil
}

func (x *AnalyzePlanRequest) GetUnpersist() *AnalyzePlanRequest_Unpersist {
	if x, ok := x.GetAnalyze().(*AnalyzePlanRequest_Unpersist_); ok {
		return x.Unpersist
	}
	return nil
}

func (x *AnalyzePlanRequest) GetGetStorageLevel() *AnalyzePlanRequest_GetStorageLevel {
	if x, ok := x.GetAnalyze().(*AnalyzePlanRequest_GetStorageLevel_); ok {
		return x.GetStorageLevel
	}
	return nil
}

type isAnalyzePlanRequest_Analyze interface {
	isAnalyzePlanRequest_Analyze()
}

type AnalyzePlanRequest_Schema_ struct {
	Schema *AnalyzePlanRequest_Schema `protobuf:"bytes,4,opt,name=schema,proto3,oneof"`
}

type AnalyzePlanRequest_Explain_ struct {
	Explain *AnalyzePlanRequest_Explain `protobuf:"bytes,5,opt,name=explain,proto3,oneof"`
}

type AnalyzePlanRequest_TreeString_ struct {
	TreeString *AnalyzePlanRequest_TreeString `protobuf:"bytes,6,opt,name=tree_string,json=treeString,proto3,oneof"`
}

type AnalyzePlanRequest_IsLocal_ struct {
	IsLocal *AnalyzePlanRequest_IsLocal `protobuf:"bytes,7,opt,name=is_local,json=isLocal,proto3,oneof"`
}

type AnalyzePlanRequest_IsStreaming_ struct {
	IsStreaming *AnalyzePlanRequest_IsStreaming `protobuf:"bytes,8,opt,name=is_streaming,json=isStreaming,proto3,oneof"`
}

type AnalyzePlanRequest_InputFiles_ struct {
	InputFiles *AnalyzePlanRequest_InputFiles `protobuf:"bytes,9,opt,name=input_files,json=inputFiles,proto3,oneof"`
}

type AnalyzePlanRequest_SparkVersion_ struct {
	SparkVersion *AnalyzePlanRequest_SparkVersion `protobuf:"bytes,10,opt,name=spark_version,json=sparkVersion,proto3,oneof"`
}

type AnalyzePlanRequest_DdlParse struct {
	DdlParse *AnalyzePlanRequest_DDLParse `protobuf:"bytes,11,opt,name=ddl_parse,json=ddlParse,proto3,oneof"`
}

type AnalyzePlanRequest_SameSemantics_ struct {
	SameSemantics *AnalyzePlanRequest_SameSemantics `protobuf:"bytes,12,opt,name=same_semantics,json=sameSemantics,proto3,oneof"`
}

type AnalyzePlanRequest_SemanticHash_ struct {
	SemanticHash *AnalyzePlanRequest_SemanticHash `protobuf:"bytes,13,opt,name=semantic_hash,json=semanticHash,proto3,oneof"`
}

type AnalyzePlanRequest_Persist_ struct {
	Persist *AnalyzePlanRequest_Persist `protobuf:"bytes,14,opt,name=persist,proto3,oneof"`
}

type AnalyzePlanRequest_Unpersist_ struct {
	Unpersist *AnalyzePlanRequest_Unpersist `protobuf:"bytes,15,opt,name=unpersist,proto3,oneof"`
}

type AnalyzePlanRequest_GetStorageLevel_ struct {
	GetStorageLevel *AnalyzePlanRequest_GetStorageLevel `protobuf:"bytes,16,opt,name=get_storage_level,json=getStorageLevel,proto3,oneof"`
}

func (*AnalyzePlanRequest_Schema_) isAnalyzePlanRequest_Analyze() {}

func (*AnalyzePlanRequest_Explain_) isAnalyzePlanRequest_Analyze() {}

func (*AnalyzePlanRequest_TreeString_) isAnalyzePlanRequest_Analyze() {}

func (*AnalyzePlanRequest_IsLocal_) isAnalyzePlanRequest_Analyze() {}

func (*AnalyzePlanRequest_IsStreaming_) isAnalyzePlanRequest_Analyze() {}

func (*AnalyzePlanRequest_InputFiles_) isAnalyzePlanRequest_Analyze() {}

func (*AnalyzePlanRequest_SparkVersion_) isAnalyzePlanRequest_Analyze() {}

func (*AnalyzePlanRequest_DdlParse) isAnalyzePlanRequest_Analyze() {}

func (*AnalyzePlanRequest_SameSemantics_) isAnalyzePlanRequest_Analyze() {}

func (*AnalyzePlanRequest_SemanticHash_) isAnalyzePlanRequest_Analyze() {}

func (*AnalyzePlanRequest_Persist_) isAnalyzePlanRequest_Analyze() {}

func (*AnalyzePlanRequest_Unpersist_) isAnalyzePlanRequest_Analyze() {}

func (*AnalyzePlanRequest_GetStorageLevel_) isAnalyzePlanRequest_Analyze() {}

// Response to performing analysis of the query. Contains relevant metadata to be able to
// reason about the performance.
type AnalyzePlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// Types that are assignable to Result:
	//
	//	*AnalyzePlanResponse_Schema_
	//	*AnalyzePlanResponse_Explain_
	//	*AnalyzePlanResponse_TreeString_
	//	*AnalyzePlanResponse_IsLocal_
	//	*AnalyzePlanResponse_IsStreaming_
	//	*AnalyzePlanResponse_InputFiles_
	//	*AnalyzePlanResponse_SparkVersion_
	//	*AnalyzePlanResponse_DdlParse
	//	*AnalyzePlanResponse_SameSemantics_
	//	*AnalyzePlanResponse_SemanticHash_
	//	*AnalyzePlanResponse_Persist_
	//	*AnalyzePlanResponse_Unpersist_
	//	*AnalyzePlanResponse_GetStorageLevel_
	Result isAnalyzePlanResponse_Result `protobuf_oneof:"result"`
}

func (x *AnalyzePlanResponse) Reset() {
	*x = AnalyzePlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanResponse) ProtoMessage() {}

func (x *AnalyzePlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanResponse.ProtoReflect.Descriptor instead.
func (*AnalyzePlanResponse) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{3}
}

func (x *AnalyzePlanResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (m *AnalyzePlanResponse) GetResult() isAnalyzePlanResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *AnalyzePlanResponse) GetSchema() *AnalyzePlanResponse_Schema {
	if x, ok := x.GetResult().(*AnalyzePlanResponse_Schema_); ok {
		return x.Schema
	}
	return nil
}

func (x *AnalyzePlanResponse) GetExplain() *AnalyzePlanResponse_Explain {
	if x, ok := x.GetResult().(*AnalyzePlanResponse_Explain_); ok {
		return x.Explain
	}
	return nil
}

func (x *AnalyzePlanResponse) GetTreeString() *AnalyzePlanResponse_TreeString {
	if x, ok := x.GetResult().(*AnalyzePlanResponse_TreeString_); ok {
		return x.TreeString
	}
	return nil
}

func (x *AnalyzePlanResponse) GetIsLocal() *AnalyzePlanResponse_IsLocal {
	if x, ok := x.GetResult().(*AnalyzePlanResponse_IsLocal_); ok {
		return x.IsLocal
	}
	return nil
}

func (x *AnalyzePlanResponse) GetIsStreaming() *AnalyzePlanResponse_IsStreaming {
	if x, ok := x.GetResult().(*AnalyzePlanResponse_IsStreaming_); ok {
		return x.IsStreaming
	}
	return nil
}

func (x *AnalyzePlanResponse) GetInputFiles() *AnalyzePlanResponse_InputFiles {
	if x, ok := x.GetResult().(*AnalyzePlanResponse_InputFiles_); ok {
		return x.InputFiles
	}
	return nil
}

func (x *AnalyzePlanResponse) GetSparkVersion() *AnalyzePlanResponse_SparkVersion {
	if x, ok := x.GetResult().(*AnalyzePlanResponse_SparkVersion_); ok {
		return x.SparkVersion
	}
	return nil
}

func (x *AnalyzePlanResponse) GetDdlParse() *AnalyzePlanResponse_DDLParse {
	if x, ok := x.GetResult().(*AnalyzePlanResponse_DdlParse); ok {
		return x.DdlParse
	}
	return nil
}

func (x *AnalyzePlanResponse) GetSameSemantics() *AnalyzePlanResponse_SameSemantics {
	if x, ok := x.GetResult().(*AnalyzePlanResponse_SameSemantics_); ok {
		return x.SameSemantics
	}
	return nil
}

func (x *AnalyzePlanResponse) GetSemanticHash() *AnalyzePlanResponse_SemanticHash {
	if x, ok := x.GetResult().(*AnalyzePlanResponse_SemanticHash_); ok {
		return x.SemanticHash
	}
	return nil
}

func (x *AnalyzePlanResponse) GetPersist() *AnalyzePlanResponse_Persist {
	if x, ok := x.GetResult().(*AnalyzePlanResponse_Persist_); ok {
		return x.Persist
	}
	return nil
}

func (x *AnalyzePlanResponse) GetUnpersist() *AnalyzePlanResponse_Unpersist {
	if x, ok := x.GetResult().(*AnalyzePlanResponse_Unpersist_); ok {
		return x.Unpersist
	}
	return nil
}

func (x *AnalyzePlanResponse) GetGetStorageLevel() *AnalyzePlanResponse_GetStorageLevel {
	if x, ok := x.GetResult().(*AnalyzePlanResponse_GetStorageLevel_); ok {
		return x.GetStorageLevel
	}
	return nil
}

type isAnalyzePlanResponse_Result interface {
	isAnalyzePlanResponse_Result()
}

type AnalyzePlanResponse_Schema_ struct {
	Schema *AnalyzePlanResponse_Schema `protobuf:"bytes,2,opt,name=schema,proto3,oneof"`
}

type AnalyzePlanResponse_Explain_ struct {
	Explain *AnalyzePlanResponse_Explain `protobuf:"bytes,3,opt,name=explain,proto3,oneof"`
}

type AnalyzePlanResponse_TreeString_ struct {
	TreeString *AnalyzePlanResponse_TreeString `protobuf:"bytes,4,opt,name=tree_string,json=treeString,proto3,oneof"`
}

type AnalyzePlanResponse_IsLocal_ struct {
	IsLocal *AnalyzePlanResponse_IsLocal `protobuf:"bytes,5,opt,name=is_local,json=isLocal,proto3,oneof"`
}

type AnalyzePlanResponse_IsStreaming_ struct {
	IsStreaming *AnalyzePlanResponse_IsStreaming `protobuf:"bytes,6,opt,name=is_streaming,json=isStreaming,proto3,oneof"`
}

type AnalyzePlanResponse_InputFiles_ struct {
	InputFiles *AnalyzePlanResponse_InputFiles `protobuf:"bytes,7,opt,name=input_files,json=inputFiles,proto3,oneof"`
}

type AnalyzePlanResponse_SparkVersion_ struct {
	SparkVersion *AnalyzePlanResponse_SparkVersion `protobuf:"bytes,8,opt,name=spark_version,json=sparkVersion,proto3,oneof"`
}

type AnalyzePlanResponse_DdlParse struct {
	DdlParse *AnalyzePlanResponse_DDLParse `protobuf:"bytes,9,opt,name=ddl_parse,json=ddlParse,proto3,oneof"`
}

type AnalyzePlanResponse_SameSemantics_ struct {
	SameSemantics *AnalyzePlanResponse_SameSemantics `protobuf:"bytes,10,opt,name=same_semantics,json=sameSemantics,proto3,oneof"`
}

type AnalyzePlanResponse_SemanticHash_ struct {
	SemanticHash *AnalyzePlanResponse_SemanticHash `protobuf:"bytes,11,opt,name=semantic_hash,json=semanticHash,proto3,oneof"`
}

type AnalyzePlanResponse_Persist_ struct {
	Persist *AnalyzePlanResponse_Persist `protobuf:"bytes,12,opt,name=persist,proto3,oneof"`
}

type AnalyzePlanResponse_Unpersist_ struct {
	Unpersist *AnalyzePlanResponse_Unpersist `protobuf:"bytes,13,opt,name=unpersist,proto3,oneof"`
}

type AnalyzePlanResponse_GetStorageLevel_ struct {
	GetStorageLevel *AnalyzePlanResponse_GetStorageLevel `protobuf:"bytes,14,opt,name=get_storage_level,json=getStorageLevel,proto3,oneof"`
}

func (*AnalyzePlanResponse_Schema_) isAnalyzePlanResponse_Result() {}

func (*AnalyzePlanResponse_Explain_) isAnalyzePlanResponse_Result() {}

func (*AnalyzePlanResponse_TreeString_) isAnalyzePlanResponse_Result() {}

func (*AnalyzePlanResponse_IsLocal_) isAnalyzePlanResponse_Result() {}

func (*AnalyzePlanResponse_IsStreaming_) isAnalyzePlanResponse_Result() {}

func (*AnalyzePlanResponse_InputFiles_) isAnalyzePlanResponse_Result() {}

func (*AnalyzePlanResponse_SparkVersion_) isAnalyzePlanResponse_Result() {}

func (*AnalyzePlanResponse_DdlParse) isAnalyzePlanResponse_Result() {}

func (*AnalyzePlanResponse_SameSemantics_) isAnalyzePlanResponse_Result() {}

func (*AnalyzePlanResponse_SemanticHash_) isAnalyzePlanResponse_Result() {}

func (*AnalyzePlanResponse_Persist_) isAnalyzePlanResponse_Result() {}

func (*AnalyzePlanResponse_Unpersist_) isAnalyzePlanResponse_Result() {}

func (*AnalyzePlanResponse_GetStorageLevel_) isAnalyzePlanResponse_Result() {}

// A request to be executed by the service.
type ExecutePlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	//
	// The session_id specifies a spark session for a user id (which is specified
	// by user_context.user_id). The session_id is set by the client to be able to
	// collate streaming responses from different queries within the dedicated session.
	// The id should be an UUID string of the format `00112233-4455-6677-8899-aabbccddeeff`
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// (Required) User context
	//
	// user_context.user_id and session+id both identify a unique remote spark session on the
	// server side.
	UserContext *UserContext `protobuf:"bytes,2,opt,name=user_context,json=userContext,proto3" json:"user_context,omitempty"`
	// (Optional)
	// Provide an id for this request. If not provided, it will be generatedCustom by the server.
	// It is returned in every ExecutePlanResponse.operation_id of the ExecutePlan response stream.
	// The id must be an UUID string of the format `00112233-4455-6677-8899-aabbccddeeff`
	OperationId *string `protobuf:"bytes,6,opt,name=operation_id,json=operationId,proto3,oneof" json:"operation_id,omitempty"`
	// (Required) The logical plan to be executed / analyzed.
	Plan *Plan `protobuf:"bytes,3,opt,name=plan,proto3" json:"plan,omitempty"`
	// Provides optional information about the client sending the request. This field
	// can be used for language or version specific information and is only intended for
	// logging purposes and will not be interpreted by the server.
	ClientType *string `protobuf:"bytes,4,opt,name=client_type,json=clientType,proto3,oneof" json:"client_type,omitempty"`
	// Repeated element for options that can be passed to the request. This element is currently
	// unused but allows to pass in an extension value used for arbitrary options.
	RequestOptions []*ExecutePlanRequest_RequestOption `protobuf:"bytes,5,rep,name=request_options,json=requestOptions,proto3" json:"request_options,omitempty"`
	// Tags to tag the given execution with.
	// Tags cannot contain ',' character and cannot be empty strings.
	// Used by Interrupt with interrupt.tag.
	Tags []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ExecutePlanRequest) Reset() {
	*x = ExecutePlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutePlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePlanRequest) ProtoMessage() {}

func (x *ExecutePlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePlanRequest.ProtoReflect.Descriptor instead.
func (*ExecutePlanRequest) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{4}
}

func (x *ExecutePlanRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ExecutePlanRequest) GetUserContext() *UserContext {
	if x != nil {
		return x.UserContext
	}
	return nil
}

func (x *ExecutePlanRequest) GetOperationId() string {
	if x != nil && x.OperationId != nil {
		return *x.OperationId
	}
	return ""
}

func (x *ExecutePlanRequest) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

func (x *ExecutePlanRequest) GetClientType() string {
	if x != nil && x.ClientType != nil {
		return *x.ClientType
	}
	return ""
}

func (x *ExecutePlanRequest) GetRequestOptions() []*ExecutePlanRequest_RequestOption {
	if x != nil {
		return x.RequestOptions
	}
	return nil
}

func (x *ExecutePlanRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// The response of a query, can be one or more for each request. Responses belonging to the
// same input query, carry the same `session_id`.
type ExecutePlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// Identifies the ExecutePlan execution.
	// If set by the client in ExecutePlanRequest.operationId, that value is returned.
	// Otherwise generatedCustom by the server.
	// It is an UUID string of the format `00112233-4455-6677-8899-aabbccddeeff`
	OperationId string `protobuf:"bytes,12,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// Identified the response in the stream.
	// The id is an UUID string of the format `00112233-4455-6677-8899-aabbccddeeff`
	ResponseId string `protobuf:"bytes,13,opt,name=response_id,json=responseId,proto3" json:"response_id,omitempty"`
	// Union type for the different response messages.
	//
	// Types that are assignable to ResponseType:
	//
	//	*ExecutePlanResponse_ArrowBatch_
	//	*ExecutePlanResponse_SqlCommandResult_
	//	*ExecutePlanResponse_WriteStreamOperationStartResult
	//	*ExecutePlanResponse_StreamingQueryCommandResult
	//	*ExecutePlanResponse_GetResourcesCommandResult
	//	*ExecutePlanResponse_StreamingQueryManagerCommandResult
	//	*ExecutePlanResponse_ResultComplete_
	//	*ExecutePlanResponse_Extension
	ResponseType isExecutePlanResponse_ResponseType `protobuf_oneof:"response_type"`
	// Metrics for the query execution. Typically, this field is only present in the last
	// batch of results and then represent the overall state of the query execution.
	Metrics *ExecutePlanResponse_Metrics `protobuf:"bytes,4,opt,name=metrics,proto3" json:"metrics,omitempty"`
	// The metrics observed during the execution of the query plan.
	ObservedMetrics []*ExecutePlanResponse_ObservedMetrics `protobuf:"bytes,6,rep,name=observed_metrics,json=observedMetrics,proto3" json:"observed_metrics,omitempty"`
	// (Optional) The Spark schema. This field is available when `collect` is called.
	Schema *DataType `protobuf:"bytes,7,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *ExecutePlanResponse) Reset() {
	*x = ExecutePlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutePlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePlanResponse) ProtoMessage() {}

func (x *ExecutePlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePlanResponse.ProtoReflect.Descriptor instead.
func (*ExecutePlanResponse) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{5}
}

func (x *ExecutePlanResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ExecutePlanResponse) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

func (x *ExecutePlanResponse) GetResponseId() string {
	if x != nil {
		return x.ResponseId
	}
	return ""
}

func (m *ExecutePlanResponse) GetResponseType() isExecutePlanResponse_ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return nil
}

func (x *ExecutePlanResponse) GetArrowBatch() *ExecutePlanResponse_ArrowBatch {
	if x, ok := x.GetResponseType().(*ExecutePlanResponse_ArrowBatch_); ok {
		return x.ArrowBatch
	}
	return nil
}

func (x *ExecutePlanResponse) GetSqlCommandResult() *ExecutePlanResponse_SqlCommandResult {
	if x, ok := x.GetResponseType().(*ExecutePlanResponse_SqlCommandResult_); ok {
		return x.SqlCommandResult
	}
	return nil
}

func (x *ExecutePlanResponse) GetWriteStreamOperationStartResult() *WriteStreamOperationStartResult {
	if x, ok := x.GetResponseType().(*ExecutePlanResponse_WriteStreamOperationStartResult); ok {
		return x.WriteStreamOperationStartResult
	}
	return nil
}

func (x *ExecutePlanResponse) GetStreamingQueryCommandResult() *StreamingQueryCommandResult {
	if x, ok := x.GetResponseType().(*ExecutePlanResponse_StreamingQueryCommandResult); ok {
		return x.StreamingQueryCommandResult
	}
	return nil
}

func (x *ExecutePlanResponse) GetGetResourcesCommandResult() *GetResourcesCommandResult {
	if x, ok := x.GetResponseType().(*ExecutePlanResponse_GetResourcesCommandResult); ok {
		return x.GetResourcesCommandResult
	}
	return nil
}

func (x *ExecutePlanResponse) GetStreamingQueryManagerCommandResult() *StreamingQueryManagerCommandResult {
	if x, ok := x.GetResponseType().(*ExecutePlanResponse_StreamingQueryManagerCommandResult); ok {
		return x.StreamingQueryManagerCommandResult
	}
	return nil
}

func (x *ExecutePlanResponse) GetResultComplete() *ExecutePlanResponse_ResultComplete {
	if x, ok := x.GetResponseType().(*ExecutePlanResponse_ResultComplete_); ok {
		return x.ResultComplete
	}
	return nil
}

func (x *ExecutePlanResponse) GetExtension() *anypb.Any {
	if x, ok := x.GetResponseType().(*ExecutePlanResponse_Extension); ok {
		return x.Extension
	}
	return nil
}

func (x *ExecutePlanResponse) GetMetrics() *ExecutePlanResponse_Metrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *ExecutePlanResponse) GetObservedMetrics() []*ExecutePlanResponse_ObservedMetrics {
	if x != nil {
		return x.ObservedMetrics
	}
	return nil
}

func (x *ExecutePlanResponse) GetSchema() *DataType {
	if x != nil {
		return x.Schema
	}
	return nil
}

type isExecutePlanResponse_ResponseType interface {
	isExecutePlanResponse_ResponseType()
}

type ExecutePlanResponse_ArrowBatch_ struct {
	ArrowBatch *ExecutePlanResponse_ArrowBatch `protobuf:"bytes,2,opt,name=arrow_batch,json=arrowBatch,proto3,oneof"`
}

type ExecutePlanResponse_SqlCommandResult_ struct {
	// Special case for executing SQL commands.
	SqlCommandResult *ExecutePlanResponse_SqlCommandResult `protobuf:"bytes,5,opt,name=sql_command_result,json=sqlCommandResult,proto3,oneof"`
}

type ExecutePlanResponse_WriteStreamOperationStartResult struct {
	// Response for a streaming query.
	WriteStreamOperationStartResult *WriteStreamOperationStartResult `protobuf:"bytes,8,opt,name=write_stream_operation_start_result,json=writeStreamOperationStartResult,proto3,oneof"`
}

type ExecutePlanResponse_StreamingQueryCommandResult struct {
	// Response for commands on a streaming query.
	StreamingQueryCommandResult *StreamingQueryCommandResult `protobuf:"bytes,9,opt,name=streaming_query_command_result,json=streamingQueryCommandResult,proto3,oneof"`
}

type ExecutePlanResponse_GetResourcesCommandResult struct {
	// Response for 'SparkContext.resources'.
	GetResourcesCommandResult *GetResourcesCommandResult `protobuf:"bytes,10,opt,name=get_resources_command_result,json=getResourcesCommandResult,proto3,oneof"`
}

type ExecutePlanResponse_StreamingQueryManagerCommandResult struct {
	// Response for commands on the streaming query manager.
	StreamingQueryManagerCommandResult *StreamingQueryManagerCommandResult `protobuf:"bytes,11,opt,name=streaming_query_manager_command_result,json=streamingQueryManagerCommandResult,proto3,oneof"`
}

type ExecutePlanResponse_ResultComplete_ struct {
	// Response type informing if the stream is complete in reattachable execution.
	ResultComplete *ExecutePlanResponse_ResultComplete `protobuf:"bytes,14,opt,name=result_complete,json=resultComplete,proto3,oneof"`
}

type ExecutePlanResponse_Extension struct {
	// Support arbitrary result objects.
	Extension *anypb.Any `protobuf:"bytes,999,opt,name=extension,proto3,oneof"`
}

func (*ExecutePlanResponse_ArrowBatch_) isExecutePlanResponse_ResponseType() {}

func (*ExecutePlanResponse_SqlCommandResult_) isExecutePlanResponse_ResponseType() {}

func (*ExecutePlanResponse_WriteStreamOperationStartResult) isExecutePlanResponse_ResponseType() {}

func (*ExecutePlanResponse_StreamingQueryCommandResult) isExecutePlanResponse_ResponseType() {}

func (*ExecutePlanResponse_GetResourcesCommandResult) isExecutePlanResponse_ResponseType() {}

func (*ExecutePlanResponse_StreamingQueryManagerCommandResult) isExecutePlanResponse_ResponseType() {}

func (*ExecutePlanResponse_ResultComplete_) isExecutePlanResponse_ResponseType() {}

func (*ExecutePlanResponse_Extension) isExecutePlanResponse_ResponseType() {}

// The key-value pair for the config request and response.
type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// (Optional) The value.
	Value *string `protobuf:"bytes,2,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{6}
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

// Request to update or fetch the configurations.
type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	//
	// The session_id specifies a spark session for a user id (which is specified
	// by user_context.user_id). The session_id is set by the client to be able to
	// collate streaming responses from different queries within the dedicated session.
	// The id should be an UUID string of the format `00112233-4455-6677-8899-aabbccddeeff`
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// (Required) User context
	UserContext *UserContext `protobuf:"bytes,2,opt,name=user_context,json=userContext,proto3" json:"user_context,omitempty"`
	// (Required) The operation for the config.
	Operation *ConfigRequest_Operation `protobuf:"bytes,3,opt,name=operation,proto3" json:"operation,omitempty"`
	// Provides optional information about the client sending the request. This field
	// can be used for language or version specific information and is only intended for
	// logging purposes and will not be interpreted by the server.
	ClientType *string `protobuf:"bytes,4,opt,name=client_type,json=clientType,proto3,oneof" json:"client_type,omitempty"`
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{7}
}

func (x *ConfigRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ConfigRequest) GetUserContext() *UserContext {
	if x != nil {
		return x.UserContext
	}
	return nil
}

func (x *ConfigRequest) GetOperation() *ConfigRequest_Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *ConfigRequest) GetClientType() string {
	if x != nil && x.ClientType != nil {
		return *x.ClientType
	}
	return ""
}

// Response to the config request.
type ConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// (Optional) The result key-value pairs.
	//
	// Available when the operation is 'Get', 'GetWithDefault', 'GetOption', 'GetAll'.
	// Also available for the operation 'IsModifiable' with boolean string "true" and "false".
	Pairs []*KeyValue `protobuf:"bytes,2,rep,name=pairs,proto3" json:"pairs,omitempty"`
	// (Optional)
	//
	// Warning messages for deprecated or unsupported configurations.
	Warnings []string `protobuf:"bytes,3,rep,name=warnings,proto3" json:"warnings,omitempty"`
}

func (x *ConfigResponse) Reset() {
	*x = ConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigResponse) ProtoMessage() {}

func (x *ConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigResponse.ProtoReflect.Descriptor instead.
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{8}
}

func (x *ConfigResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ConfigResponse) GetPairs() []*KeyValue {
	if x != nil {
		return x.Pairs
	}
	return nil
}

func (x *ConfigResponse) GetWarnings() []string {
	if x != nil {
		return x.Warnings
	}
	return nil
}

// Request to transfer client-local artifacts.
type AddArtifactsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	//
	// The session_id specifies a spark session for a user id (which is specified
	// by user_context.user_id). The session_id is set by the client to be able to
	// collate streaming responses from different queries within the dedicated session.
	// The id should be an UUID string of the format `00112233-4455-6677-8899-aabbccddeeff`
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// User context
	UserContext *UserContext `protobuf:"bytes,2,opt,name=user_context,json=userContext,proto3" json:"user_context,omitempty"`
	// Provides optional information about the client sending the request. This field
	// can be used for language or version specific information and is only intended for
	// logging purposes and will not be interpreted by the server.
	ClientType *string `protobuf:"bytes,6,opt,name=client_type,json=clientType,proto3,oneof" json:"client_type,omitempty"`
	// The payload is either a batch of artifacts or a partial chunk of a large artifact.
	//
	// Types that are assignable to Payload:
	//
	//	*AddArtifactsRequest_Batch_
	//	*AddArtifactsRequest_BeginChunk
	//	*AddArtifactsRequest_Chunk
	Payload isAddArtifactsRequest_Payload `protobuf_oneof:"payload"`
}

func (x *AddArtifactsRequest) Reset() {
	*x = AddArtifactsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddArtifactsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddArtifactsRequest) ProtoMessage() {}

func (x *AddArtifactsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddArtifactsRequest.ProtoReflect.Descriptor instead.
func (*AddArtifactsRequest) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{9}
}

func (x *AddArtifactsRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *AddArtifactsRequest) GetUserContext() *UserContext {
	if x != nil {
		return x.UserContext
	}
	return nil
}

func (x *AddArtifactsRequest) GetClientType() string {
	if x != nil && x.ClientType != nil {
		return *x.ClientType
	}
	return ""
}

func (m *AddArtifactsRequest) GetPayload() isAddArtifactsRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *AddArtifactsRequest) GetBatch() *AddArtifactsRequest_Batch {
	if x, ok := x.GetPayload().(*AddArtifactsRequest_Batch_); ok {
		return x.Batch
	}
	return nil
}

func (x *AddArtifactsRequest) GetBeginChunk() *AddArtifactsRequest_BeginChunkedArtifact {
	if x, ok := x.GetPayload().(*AddArtifactsRequest_BeginChunk); ok {
		return x.BeginChunk
	}
	return nil
}

func (x *AddArtifactsRequest) GetChunk() *AddArtifactsRequest_ArtifactChunk {
	if x, ok := x.GetPayload().(*AddArtifactsRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isAddArtifactsRequest_Payload interface {
	isAddArtifactsRequest_Payload()
}

type AddArtifactsRequest_Batch_ struct {
	Batch *AddArtifactsRequest_Batch `protobuf:"bytes,3,opt,name=batch,proto3,oneof"`
}

type AddArtifactsRequest_BeginChunk struct {
	// The metadata and the initial chunk of a large artifact chunked into multiple requests.
	// The server side is notified about the total size of the large artifact as well as the
	// number of chunks to expect.
	BeginChunk *AddArtifactsRequest_BeginChunkedArtifact `protobuf:"bytes,4,opt,name=begin_chunk,json=beginChunk,proto3,oneof"`
}

type AddArtifactsRequest_Chunk struct {
	// A chunk of an artifact excluding metadata. This can be any chunk of a large artifact
	// excluding the first chunk (which is included in `BeginChunkedArtifact`).
	Chunk *AddArtifactsRequest_ArtifactChunk `protobuf:"bytes,5,opt,name=chunk,proto3,oneof"`
}

func (*AddArtifactsRequest_Batch_) isAddArtifactsRequest_Payload() {}

func (*AddArtifactsRequest_BeginChunk) isAddArtifactsRequest_Payload() {}

func (*AddArtifactsRequest_Chunk) isAddArtifactsRequest_Payload() {}

// Response to adding an artifact. Contains relevant metadata to verify successful transfer of
// artifact(s).
type AddArtifactsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of artifact(s) seen by the server.
	Artifacts []*AddArtifactsResponse_ArtifactSummary `protobuf:"bytes,1,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
}

func (x *AddArtifactsResponse) Reset() {
	*x = AddArtifactsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddArtifactsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddArtifactsResponse) ProtoMessage() {}

func (x *AddArtifactsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddArtifactsResponse.ProtoReflect.Descriptor instead.
func (*AddArtifactsResponse) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{10}
}

func (x *AddArtifactsResponse) GetArtifacts() []*AddArtifactsResponse_ArtifactSummary {
	if x != nil {
		return x.Artifacts
	}
	return nil
}

// Request to get current statuses of artifacts at the server side.
type ArtifactStatusesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	//
	// The session_id specifies a spark session for a user id (which is specified
	// by user_context.user_id). The session_id is set by the client to be able to
	// collate streaming responses from different queries within the dedicated session.
	// The id should be an UUID string of the format `00112233-4455-6677-8899-aabbccddeeff`
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// User context
	UserContext *UserContext `protobuf:"bytes,2,opt,name=user_context,json=userContext,proto3" json:"user_context,omitempty"`
	// Provides optional information about the client sending the request. This field
	// can be used for language or version specific information and is only intended for
	// logging purposes and will not be interpreted by the server.
	ClientType *string `protobuf:"bytes,3,opt,name=client_type,json=clientType,proto3,oneof" json:"client_type,omitempty"`
	// The name of the artifact is expected in the form of a "Relative Path" that is made up of a
	// sequence of directories and the final file element.
	// Examples of "Relative Path"s: "jars/test.jar", "classes/xyz.class", "abc.xyz", "a/b/X.jar".
	// The server is expected to maintain the hierarchy of files as defined by their name. (i.e
	// The relative path of the file on the server's filesystem will be the same as the name of
	// the provided artifact)
	Names []string `protobuf:"bytes,4,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *ArtifactStatusesRequest) Reset() {
	*x = ArtifactStatusesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactStatusesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactStatusesRequest) ProtoMessage() {}

func (x *ArtifactStatusesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactStatusesRequest.ProtoReflect.Descriptor instead.
func (*ArtifactStatusesRequest) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{11}
}

func (x *ArtifactStatusesRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ArtifactStatusesRequest) GetUserContext() *UserContext {
	if x != nil {
		return x.UserContext
	}
	return nil
}

func (x *ArtifactStatusesRequest) GetClientType() string {
	if x != nil && x.ClientType != nil {
		return *x.ClientType
	}
	return ""
}

func (x *ArtifactStatusesRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

// Response to checking artifact statuses.
type ArtifactStatusesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A map of artifact names to their statuses.
	Statuses map[string]*ArtifactStatusesResponse_ArtifactStatus `protobuf:"bytes,1,rep,name=statuses,proto3" json:"statuses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ArtifactStatusesResponse) Reset() {
	*x = ArtifactStatusesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactStatusesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactStatusesResponse) ProtoMessage() {}

func (x *ArtifactStatusesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactStatusesResponse.ProtoReflect.Descriptor instead.
func (*ArtifactStatusesResponse) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{12}
}

func (x *ArtifactStatusesResponse) GetStatuses() map[string]*ArtifactStatusesResponse_ArtifactStatus {
	if x != nil {
		return x.Statuses
	}
	return nil
}

type InterruptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	//
	// The session_id specifies a spark session for a user id (which is specified
	// by user_context.user_id). The session_id is set by the client to be able to
	// collate streaming responses from different queries within the dedicated session.
	// The id should be an UUID string of the format `00112233-4455-6677-8899-aabbccddeeff`
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// (Required) User context
	UserContext *UserContext `protobuf:"bytes,2,opt,name=user_context,json=userContext,proto3" json:"user_context,omitempty"`
	// Provides optional information about the client sending the request. This field
	// can be used for language or version specific information and is only intended for
	// logging purposes and will not be interpreted by the server.
	ClientType *string `protobuf:"bytes,3,opt,name=client_type,json=clientType,proto3,oneof" json:"client_type,omitempty"`
	// (Required) The type of interrupt to execute.
	InterruptType InterruptRequest_InterruptType `protobuf:"varint,4,opt,name=interrupt_type,json=interruptType,proto3,enum=spark.connect.InterruptRequest_InterruptType" json:"interrupt_type,omitempty"`
	// Types that are assignable to Interrupt:
	//
	//	*InterruptRequest_OperationTag
	//	*InterruptRequest_OperationId
	Interrupt isInterruptRequest_Interrupt `protobuf_oneof:"interrupt"`
}

func (x *InterruptRequest) Reset() {
	*x = InterruptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterruptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterruptRequest) ProtoMessage() {}

func (x *InterruptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterruptRequest.ProtoReflect.Descriptor instead.
func (*InterruptRequest) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{13}
}

func (x *InterruptRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *InterruptRequest) GetUserContext() *UserContext {
	if x != nil {
		return x.UserContext
	}
	return nil
}

func (x *InterruptRequest) GetClientType() string {
	if x != nil && x.ClientType != nil {
		return *x.ClientType
	}
	return ""
}

func (x *InterruptRequest) GetInterruptType() InterruptRequest_InterruptType {
	if x != nil {
		return x.InterruptType
	}
	return InterruptRequest_INTERRUPT_TYPE_UNSPECIFIED
}

func (m *InterruptRequest) GetInterrupt() isInterruptRequest_Interrupt {
	if m != nil {
		return m.Interrupt
	}
	return nil
}

func (x *InterruptRequest) GetOperationTag() string {
	if x, ok := x.GetInterrupt().(*InterruptRequest_OperationTag); ok {
		return x.OperationTag
	}
	return ""
}

func (x *InterruptRequest) GetOperationId() string {
	if x, ok := x.GetInterrupt().(*InterruptRequest_OperationId); ok {
		return x.OperationId
	}
	return ""
}

type isInterruptRequest_Interrupt interface {
	isInterruptRequest_Interrupt()
}

type InterruptRequest_OperationTag struct {
	// if interrupt_tag == INTERRUPT_TYPE_TAG, interrupt operation with this tag.
	OperationTag string `protobuf:"bytes,5,opt,name=operation_tag,json=operationTag,proto3,oneof"`
}

type InterruptRequest_OperationId struct {
	// if interrupt_tag == INTERRUPT_TYPE_OPERATION_ID, interrupt operation with this operation_id.
	OperationId string `protobuf:"bytes,6,opt,name=operation_id,json=operationId,proto3,oneof"`
}

func (*InterruptRequest_OperationTag) isInterruptRequest_Interrupt() {}

func (*InterruptRequest_OperationId) isInterruptRequest_Interrupt() {}

type InterruptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Session id in which the interrupt was running.
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// Operation ids of the executions which were interrupted.
	InterruptedIds []string `protobuf:"bytes,2,rep,name=interrupted_ids,json=interruptedIds,proto3" json:"interrupted_ids,omitempty"`
}

func (x *InterruptResponse) Reset() {
	*x = InterruptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterruptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterruptResponse) ProtoMessage() {}

func (x *InterruptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterruptResponse.ProtoReflect.Descriptor instead.
func (*InterruptResponse) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{14}
}

func (x *InterruptResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *InterruptResponse) GetInterruptedIds() []string {
	if x != nil {
		return x.InterruptedIds
	}
	return nil
}

type ReattachOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, the request can be reattached to using ReattachExecute.
	// ReattachExecute can be used either if the stream broke with a GRPC network error,
	// or if the server closed the stream without sending a response with StreamStatus.complete=true.
	// The server will keep a buffer of responses in case a response is lost, and
	// ReattachExecute needs to back-track.
	//
	// If false, the execution response stream will will not be reattachable, and all responses are
	// immediately released by the server after being sent.
	Reattachable bool `protobuf:"varint,1,opt,name=reattachable,proto3" json:"reattachable,omitempty"`
}

func (x *ReattachOptions) Reset() {
	*x = ReattachOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReattachOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReattachOptions) ProtoMessage() {}

func (x *ReattachOptions) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReattachOptions.ProtoReflect.Descriptor instead.
func (*ReattachOptions) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{15}
}

func (x *ReattachOptions) GetReattachable() bool {
	if x != nil {
		return x.Reattachable
	}
	return false
}

type ReattachExecuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	//
	// The session_id of the request to reattach to.
	// This must be an id of existing session.
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// (Required) User context
	//
	// user_context.user_id and session+id both identify a unique remote spark session on the
	// server side.
	UserContext *UserContext `protobuf:"bytes,2,opt,name=user_context,json=userContext,proto3" json:"user_context,omitempty"`
	// (Required)
	// Provide an id of the request to reattach to.
	// This must be an id of existing operation.
	OperationId string `protobuf:"bytes,3,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// Provides optional information about the client sending the request. This field
	// can be used for language or version specific information and is only intended for
	// logging purposes and will not be interpreted by the server.
	ClientType *string `protobuf:"bytes,4,opt,name=client_type,json=clientType,proto3,oneof" json:"client_type,omitempty"`
	// (Optional)
	// Last already processed response id from the response stream.
	// After reattach, server will resume the response stream after that response.
	// If not specified, server will restart the stream from the start.
	//
	// Note: server controls the amount of responses that it buffers and it may drop responses,
	// that are far behind the latest returned response, so this can't be used to arbitrarily
	// scroll back the cursor. If the response is no longer available, this will result in an error.
	LastResponseId *string `protobuf:"bytes,5,opt,name=last_response_id,json=lastResponseId,proto3,oneof" json:"last_response_id,omitempty"`
}

func (x *ReattachExecuteRequest) Reset() {
	*x = ReattachExecuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReattachExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReattachExecuteRequest) ProtoMessage() {}

func (x *ReattachExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReattachExecuteRequest.ProtoReflect.Descriptor instead.
func (*ReattachExecuteRequest) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{16}
}

func (x *ReattachExecuteRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ReattachExecuteRequest) GetUserContext() *UserContext {
	if x != nil {
		return x.UserContext
	}
	return nil
}

func (x *ReattachExecuteRequest) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

func (x *ReattachExecuteRequest) GetClientType() string {
	if x != nil && x.ClientType != nil {
		return *x.ClientType
	}
	return ""
}

func (x *ReattachExecuteRequest) GetLastResponseId() string {
	if x != nil && x.LastResponseId != nil {
		return *x.LastResponseId
	}
	return ""
}

type ReleaseExecuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	//
	// The session_id of the request to reattach to.
	// This must be an id of existing session.
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// (Required) User context
	//
	// user_context.user_id and session+id both identify a unique remote spark session on the
	// server side.
	UserContext *UserContext `protobuf:"bytes,2,opt,name=user_context,json=userContext,proto3" json:"user_context,omitempty"`
	// (Required)
	// Provide an id of the request to reattach to.
	// This must be an id of existing operation.
	OperationId string `protobuf:"bytes,3,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// Provides optional information about the client sending the request. This field
	// can be used for language or version specific information and is only intended for
	// logging purposes and will not be interpreted by the server.
	ClientType *string `protobuf:"bytes,4,opt,name=client_type,json=clientType,proto3,oneof" json:"client_type,omitempty"`
	// Types that are assignable to Release:
	//
	//	*ReleaseExecuteRequest_ReleaseAll_
	//	*ReleaseExecuteRequest_ReleaseUntil_
	Release isReleaseExecuteRequest_Release `protobuf_oneof:"release"`
}

func (x *ReleaseExecuteRequest) Reset() {
	*x = ReleaseExecuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseExecuteRequest) ProtoMessage() {}

func (x *ReleaseExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseExecuteRequest.ProtoReflect.Descriptor instead.
func (*ReleaseExecuteRequest) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{17}
}

func (x *ReleaseExecuteRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ReleaseExecuteRequest) GetUserContext() *UserContext {
	if x != nil {
		return x.UserContext
	}
	return nil
}

func (x *ReleaseExecuteRequest) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

func (x *ReleaseExecuteRequest) GetClientType() string {
	if x != nil && x.ClientType != nil {
		return *x.ClientType
	}
	return ""
}

func (m *ReleaseExecuteRequest) GetRelease() isReleaseExecuteRequest_Release {
	if m != nil {
		return m.Release
	}
	return nil
}

func (x *ReleaseExecuteRequest) GetReleaseAll() *ReleaseExecuteRequest_ReleaseAll {
	if x, ok := x.GetRelease().(*ReleaseExecuteRequest_ReleaseAll_); ok {
		return x.ReleaseAll
	}
	return nil
}

func (x *ReleaseExecuteRequest) GetReleaseUntil() *ReleaseExecuteRequest_ReleaseUntil {
	if x, ok := x.GetRelease().(*ReleaseExecuteRequest_ReleaseUntil_); ok {
		return x.ReleaseUntil
	}
	return nil
}

type isReleaseExecuteRequest_Release interface {
	isReleaseExecuteRequest_Release()
}

type ReleaseExecuteRequest_ReleaseAll_ struct {
	ReleaseAll *ReleaseExecuteRequest_ReleaseAll `protobuf:"bytes,5,opt,name=release_all,json=releaseAll,proto3,oneof"`
}

type ReleaseExecuteRequest_ReleaseUntil_ struct {
	ReleaseUntil *ReleaseExecuteRequest_ReleaseUntil `protobuf:"bytes,6,opt,name=release_until,json=releaseUntil,proto3,oneof"`
}

func (*ReleaseExecuteRequest_ReleaseAll_) isReleaseExecuteRequest_Release() {}

func (*ReleaseExecuteRequest_ReleaseUntil_) isReleaseExecuteRequest_Release() {}

type ReleaseExecuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Session id in which the release was running.
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// Operation id of the operation on which the release executed.
	// If the operation couldn't be found (because e.g. it was concurrently released), will be unset.
	// Otherwise, it will be equal to the operation_id from request.
	OperationId *string `protobuf:"bytes,2,opt,name=operation_id,json=operationId,proto3,oneof" json:"operation_id,omitempty"`
}

func (x *ReleaseExecuteResponse) Reset() {
	*x = ReleaseExecuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseExecuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseExecuteResponse) ProtoMessage() {}

func (x *ReleaseExecuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseExecuteResponse.ProtoReflect.Descriptor instead.
func (*ReleaseExecuteResponse) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{18}
}

func (x *ReleaseExecuteResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ReleaseExecuteResponse) GetOperationId() string {
	if x != nil && x.OperationId != nil {
		return *x.OperationId
	}
	return ""
}

type AnalyzePlanRequest_Schema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The logical plan to be analyzed.
	Plan *Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (x *AnalyzePlanRequest_Schema) Reset() {
	*x = AnalyzePlanRequest_Schema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanRequest_Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanRequest_Schema) ProtoMessage() {}

func (x *AnalyzePlanRequest_Schema) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanRequest_Schema.ProtoReflect.Descriptor instead.
func (*AnalyzePlanRequest_Schema) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{2, 0}
}

func (x *AnalyzePlanRequest_Schema) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

// Explains the input plan based on a configurable mode.
type AnalyzePlanRequest_Explain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The logical plan to be analyzed.
	Plan *Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	// (Required) For analyzePlan rpc calls, configure the mode to explain plan in strings.
	ExplainMode AnalyzePlanRequest_Explain_ExplainMode `protobuf:"varint,2,opt,name=explain_mode,json=explainMode,proto3,enum=spark.connect.AnalyzePlanRequest_Explain_ExplainMode" json:"explain_mode,omitempty"`
}

func (x *AnalyzePlanRequest_Explain) Reset() {
	*x = AnalyzePlanRequest_Explain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[20]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanRequest_Explain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanRequest_Explain) ProtoMessage() {}

func (x *AnalyzePlanRequest_Explain) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[20]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanRequest_Explain.ProtoReflect.Descriptor instead.
func (*AnalyzePlanRequest_Explain) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{2, 1}
}

func (x *AnalyzePlanRequest_Explain) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

func (x *AnalyzePlanRequest_Explain) GetExplainMode() AnalyzePlanRequest_Explain_ExplainMode {
	if x != nil {
		return x.ExplainMode
	}
	return AnalyzePlanRequest_Explain_EXPLAIN_MODE_UNSPECIFIED
}

type AnalyzePlanRequest_TreeString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The logical plan to be analyzed.
	Plan *Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	// (Optional) Max level of the schema.
	Level *int32 `protobuf:"varint,2,opt,name=level,proto3,oneof" json:"level,omitempty"`
}

func (x *AnalyzePlanRequest_TreeString) Reset() {
	*x = AnalyzePlanRequest_TreeString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[21]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanRequest_TreeString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanRequest_TreeString) ProtoMessage() {}

func (x *AnalyzePlanRequest_TreeString) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[21]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanRequest_TreeString.ProtoReflect.Descriptor instead.
func (*AnalyzePlanRequest_TreeString) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{2, 2}
}

func (x *AnalyzePlanRequest_TreeString) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

func (x *AnalyzePlanRequest_TreeString) GetLevel() int32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

type AnalyzePlanRequest_IsLocal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The logical plan to be analyzed.
	Plan *Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (x *AnalyzePlanRequest_IsLocal) Reset() {
	*x = AnalyzePlanRequest_IsLocal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[22]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanRequest_IsLocal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanRequest_IsLocal) ProtoMessage() {}

func (x *AnalyzePlanRequest_IsLocal) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[22]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanRequest_IsLocal.ProtoReflect.Descriptor instead.
func (*AnalyzePlanRequest_IsLocal) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{2, 3}
}

func (x *AnalyzePlanRequest_IsLocal) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

type AnalyzePlanRequest_IsStreaming struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The logical plan to be analyzed.
	Plan *Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (x *AnalyzePlanRequest_IsStreaming) Reset() {
	*x = AnalyzePlanRequest_IsStreaming{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[23]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanRequest_IsStreaming) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanRequest_IsStreaming) ProtoMessage() {}

func (x *AnalyzePlanRequest_IsStreaming) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[23]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanRequest_IsStreaming.ProtoReflect.Descriptor instead.
func (*AnalyzePlanRequest_IsStreaming) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{2, 4}
}

func (x *AnalyzePlanRequest_IsStreaming) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

type AnalyzePlanRequest_InputFiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The logical plan to be analyzed.
	Plan *Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (x *AnalyzePlanRequest_InputFiles) Reset() {
	*x = AnalyzePlanRequest_InputFiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[24]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanRequest_InputFiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanRequest_InputFiles) ProtoMessage() {}

func (x *AnalyzePlanRequest_InputFiles) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[24]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanRequest_InputFiles.ProtoReflect.Descriptor instead.
func (*AnalyzePlanRequest_InputFiles) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{2, 5}
}

func (x *AnalyzePlanRequest_InputFiles) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

type AnalyzePlanRequest_SparkVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnalyzePlanRequest_SparkVersion) Reset() {
	*x = AnalyzePlanRequest_SparkVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[25]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanRequest_SparkVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanRequest_SparkVersion) ProtoMessage() {}

func (x *AnalyzePlanRequest_SparkVersion) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[25]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanRequest_SparkVersion.ProtoReflect.Descriptor instead.
func (*AnalyzePlanRequest_SparkVersion) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{2, 6}
}

type AnalyzePlanRequest_DDLParse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The DDL formatted string to be parsed.
	DdlString string `protobuf:"bytes,1,opt,name=ddl_string,json=ddlString,proto3" json:"ddl_string,omitempty"`
}

func (x *AnalyzePlanRequest_DDLParse) Reset() {
	*x = AnalyzePlanRequest_DDLParse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[26]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanRequest_DDLParse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanRequest_DDLParse) ProtoMessage() {}

func (x *AnalyzePlanRequest_DDLParse) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[26]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanRequest_DDLParse.ProtoReflect.Descriptor instead.
func (*AnalyzePlanRequest_DDLParse) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{2, 7}
}

func (x *AnalyzePlanRequest_DDLParse) GetDdlString() string {
	if x != nil {
		return x.DdlString
	}
	return ""
}

// Returns `true` when the logical query plans  are equal and therefore return same results.
type AnalyzePlanRequest_SameSemantics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The plan to be compared.
	TargetPlan *Plan `protobuf:"bytes,1,opt,name=target_plan,json=targetPlan,proto3" json:"target_plan,omitempty"`
	// (Required) The other plan to be compared.
	OtherPlan *Plan `protobuf:"bytes,2,opt,name=other_plan,json=otherPlan,proto3" json:"other_plan,omitempty"`
}

func (x *AnalyzePlanRequest_SameSemantics) Reset() {
	*x = AnalyzePlanRequest_SameSemantics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[27]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanRequest_SameSemantics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanRequest_SameSemantics) ProtoMessage() {}

func (x *AnalyzePlanRequest_SameSemantics) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[27]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanRequest_SameSemantics.ProtoReflect.Descriptor instead.
func (*AnalyzePlanRequest_SameSemantics) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{2, 8}
}

func (x *AnalyzePlanRequest_SameSemantics) GetTargetPlan() *Plan {
	if x != nil {
		return x.TargetPlan
	}
	return nil
}

func (x *AnalyzePlanRequest_SameSemantics) GetOtherPlan() *Plan {
	if x != nil {
		return x.OtherPlan
	}
	return nil
}

type AnalyzePlanRequest_SemanticHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The logical plan to get a hashCode.
	Plan *Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (x *AnalyzePlanRequest_SemanticHash) Reset() {
	*x = AnalyzePlanRequest_SemanticHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[28]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanRequest_SemanticHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanRequest_SemanticHash) ProtoMessage() {}

func (x *AnalyzePlanRequest_SemanticHash) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[28]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanRequest_SemanticHash.ProtoReflect.Descriptor instead.
func (*AnalyzePlanRequest_SemanticHash) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{2, 9}
}

func (x *AnalyzePlanRequest_SemanticHash) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

type AnalyzePlanRequest_Persist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The logical plan to persist.
	Relation *Relation `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"`
	// (Optional) The storage level.
	StorageLevel *StorageLevel `protobuf:"bytes,2,opt,name=storage_level,json=storageLevel,proto3,oneof" json:"storage_level,omitempty"`
}

func (x *AnalyzePlanRequest_Persist) Reset() {
	*x = AnalyzePlanRequest_Persist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[29]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanRequest_Persist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanRequest_Persist) ProtoMessage() {}

func (x *AnalyzePlanRequest_Persist) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[29]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanRequest_Persist.ProtoReflect.Descriptor instead.
func (*AnalyzePlanRequest_Persist) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{2, 10}
}

func (x *AnalyzePlanRequest_Persist) GetRelation() *Relation {
	if x != nil {
		return x.Relation
	}
	return nil
}

func (x *AnalyzePlanRequest_Persist) GetStorageLevel() *StorageLevel {
	if x != nil {
		return x.StorageLevel
	}
	return nil
}

type AnalyzePlanRequest_Unpersist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The logical plan to unpersist.
	Relation *Relation `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"`
	// (Optional) Whether to block until all blocks are deleted.
	Blocking *bool `protobuf:"varint,2,opt,name=blocking,proto3,oneof" json:"blocking,omitempty"`
}

func (x *AnalyzePlanRequest_Unpersist) Reset() {
	*x = AnalyzePlanRequest_Unpersist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[30]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanRequest_Unpersist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanRequest_Unpersist) ProtoMessage() {}

func (x *AnalyzePlanRequest_Unpersist) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[30]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanRequest_Unpersist.ProtoReflect.Descriptor instead.
func (*AnalyzePlanRequest_Unpersist) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{2, 11}
}

func (x *AnalyzePlanRequest_Unpersist) GetRelation() *Relation {
	if x != nil {
		return x.Relation
	}
	return nil
}

func (x *AnalyzePlanRequest_Unpersist) GetBlocking() bool {
	if x != nil && x.Blocking != nil {
		return *x.Blocking
	}
	return false
}

type AnalyzePlanRequest_GetStorageLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The logical plan to get the storage level.
	Relation *Relation `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *AnalyzePlanRequest_GetStorageLevel) Reset() {
	*x = AnalyzePlanRequest_GetStorageLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[31]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanRequest_GetStorageLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanRequest_GetStorageLevel) ProtoMessage() {}

func (x *AnalyzePlanRequest_GetStorageLevel) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[31]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanRequest_GetStorageLevel.ProtoReflect.Descriptor instead.
func (*AnalyzePlanRequest_GetStorageLevel) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{2, 12}
}

func (x *AnalyzePlanRequest_GetStorageLevel) GetRelation() *Relation {
	if x != nil {
		return x.Relation
	}
	return nil
}

type AnalyzePlanResponse_Schema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema *DataType `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *AnalyzePlanResponse_Schema) Reset() {
	*x = AnalyzePlanResponse_Schema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[32]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanResponse_Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanResponse_Schema) ProtoMessage() {}

func (x *AnalyzePlanResponse_Schema) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[32]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanResponse_Schema.ProtoReflect.Descriptor instead.
func (*AnalyzePlanResponse_Schema) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{3, 0}
}

func (x *AnalyzePlanResponse_Schema) GetSchema() *DataType {
	if x != nil {
		return x.Schema
	}
	return nil
}

type AnalyzePlanResponse_Explain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExplainString string `protobuf:"bytes,1,opt,name=explain_string,json=explainString,proto3" json:"explain_string,omitempty"`
}

func (x *AnalyzePlanResponse_Explain) Reset() {
	*x = AnalyzePlanResponse_Explain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[33]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanResponse_Explain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanResponse_Explain) ProtoMessage() {}

func (x *AnalyzePlanResponse_Explain) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[33]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanResponse_Explain.ProtoReflect.Descriptor instead.
func (*AnalyzePlanResponse_Explain) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{3, 1}
}

func (x *AnalyzePlanResponse_Explain) GetExplainString() string {
	if x != nil {
		return x.ExplainString
	}
	return ""
}

type AnalyzePlanResponse_TreeString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TreeString string `protobuf:"bytes,1,opt,name=tree_string,json=treeString,proto3" json:"tree_string,omitempty"`
}

func (x *AnalyzePlanResponse_TreeString) Reset() {
	*x = AnalyzePlanResponse_TreeString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[34]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanResponse_TreeString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanResponse_TreeString) ProtoMessage() {}

func (x *AnalyzePlanResponse_TreeString) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[34]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanResponse_TreeString.ProtoReflect.Descriptor instead.
func (*AnalyzePlanResponse_TreeString) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{3, 2}
}

func (x *AnalyzePlanResponse_TreeString) GetTreeString() string {
	if x != nil {
		return x.TreeString
	}
	return ""
}

type AnalyzePlanResponse_IsLocal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsLocal bool `protobuf:"varint,1,opt,name=is_local,json=isLocal,proto3" json:"is_local,omitempty"`
}

func (x *AnalyzePlanResponse_IsLocal) Reset() {
	*x = AnalyzePlanResponse_IsLocal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[35]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanResponse_IsLocal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanResponse_IsLocal) ProtoMessage() {}

func (x *AnalyzePlanResponse_IsLocal) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[35]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanResponse_IsLocal.ProtoReflect.Descriptor instead.
func (*AnalyzePlanResponse_IsLocal) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{3, 3}
}

func (x *AnalyzePlanResponse_IsLocal) GetIsLocal() bool {
	if x != nil {
		return x.IsLocal
	}
	return false
}

type AnalyzePlanResponse_IsStreaming struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsStreaming bool `protobuf:"varint,1,opt,name=is_streaming,json=isStreaming,proto3" json:"is_streaming,omitempty"`
}

func (x *AnalyzePlanResponse_IsStreaming) Reset() {
	*x = AnalyzePlanResponse_IsStreaming{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[36]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanResponse_IsStreaming) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanResponse_IsStreaming) ProtoMessage() {}

func (x *AnalyzePlanResponse_IsStreaming) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[36]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanResponse_IsStreaming.ProtoReflect.Descriptor instead.
func (*AnalyzePlanResponse_IsStreaming) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{3, 4}
}

func (x *AnalyzePlanResponse_IsStreaming) GetIsStreaming() bool {
	if x != nil {
		return x.IsStreaming
	}
	return false
}

type AnalyzePlanResponse_InputFiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A best-effort snapshot of the files that compose this Dataset
	Files []string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *AnalyzePlanResponse_InputFiles) Reset() {
	*x = AnalyzePlanResponse_InputFiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[37]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanResponse_InputFiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanResponse_InputFiles) ProtoMessage() {}

func (x *AnalyzePlanResponse_InputFiles) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[37]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanResponse_InputFiles.ProtoReflect.Descriptor instead.
func (*AnalyzePlanResponse_InputFiles) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{3, 5}
}

func (x *AnalyzePlanResponse_InputFiles) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

type AnalyzePlanResponse_SparkVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *AnalyzePlanResponse_SparkVersion) Reset() {
	*x = AnalyzePlanResponse_SparkVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[38]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanResponse_SparkVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanResponse_SparkVersion) ProtoMessage() {}

func (x *AnalyzePlanResponse_SparkVersion) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[38]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanResponse_SparkVersion.ProtoReflect.Descriptor instead.
func (*AnalyzePlanResponse_SparkVersion) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{3, 6}
}

func (x *AnalyzePlanResponse_SparkVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type AnalyzePlanResponse_DDLParse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parsed *DataType `protobuf:"bytes,1,opt,name=parsed,proto3" json:"parsed,omitempty"`
}

func (x *AnalyzePlanResponse_DDLParse) Reset() {
	*x = AnalyzePlanResponse_DDLParse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[39]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanResponse_DDLParse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanResponse_DDLParse) ProtoMessage() {}

func (x *AnalyzePlanResponse_DDLParse) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[39]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanResponse_DDLParse.ProtoReflect.Descriptor instead.
func (*AnalyzePlanResponse_DDLParse) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{3, 7}
}

func (x *AnalyzePlanResponse_DDLParse) GetParsed() *DataType {
	if x != nil {
		return x.Parsed
	}
	return nil
}

type AnalyzePlanResponse_SameSemantics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AnalyzePlanResponse_SameSemantics) Reset() {
	*x = AnalyzePlanResponse_SameSemantics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[40]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanResponse_SameSemantics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanResponse_SameSemantics) ProtoMessage() {}

func (x *AnalyzePlanResponse_SameSemantics) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[40]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanResponse_SameSemantics.ProtoReflect.Descriptor instead.
func (*AnalyzePlanResponse_SameSemantics) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{3, 8}
}

func (x *AnalyzePlanResponse_SameSemantics) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type AnalyzePlanResponse_SemanticHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AnalyzePlanResponse_SemanticHash) Reset() {
	*x = AnalyzePlanResponse_SemanticHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[41]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanResponse_SemanticHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanResponse_SemanticHash) ProtoMessage() {}

func (x *AnalyzePlanResponse_SemanticHash) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[41]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanResponse_SemanticHash.ProtoReflect.Descriptor instead.
func (*AnalyzePlanResponse_SemanticHash) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{3, 9}
}

func (x *AnalyzePlanResponse_SemanticHash) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type AnalyzePlanResponse_Persist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnalyzePlanResponse_Persist) Reset() {
	*x = AnalyzePlanResponse_Persist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[42]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanResponse_Persist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanResponse_Persist) ProtoMessage() {}

func (x *AnalyzePlanResponse_Persist) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[42]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanResponse_Persist.ProtoReflect.Descriptor instead.
func (*AnalyzePlanResponse_Persist) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{3, 10}
}

type AnalyzePlanResponse_Unpersist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnalyzePlanResponse_Unpersist) Reset() {
	*x = AnalyzePlanResponse_Unpersist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[43]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanResponse_Unpersist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanResponse_Unpersist) ProtoMessage() {}

func (x *AnalyzePlanResponse_Unpersist) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[43]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanResponse_Unpersist.ProtoReflect.Descriptor instead.
func (*AnalyzePlanResponse_Unpersist) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{3, 11}
}

type AnalyzePlanResponse_GetStorageLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The StorageLevel as a result of get_storage_level request.
	StorageLevel *StorageLevel `protobuf:"bytes,1,opt,name=storage_level,json=storageLevel,proto3" json:"storage_level,omitempty"`
}

func (x *AnalyzePlanResponse_GetStorageLevel) Reset() {
	*x = AnalyzePlanResponse_GetStorageLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[44]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzePlanResponse_GetStorageLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzePlanResponse_GetStorageLevel) ProtoMessage() {}

func (x *AnalyzePlanResponse_GetStorageLevel) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[44]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzePlanResponse_GetStorageLevel.ProtoReflect.Descriptor instead.
func (*AnalyzePlanResponse_GetStorageLevel) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{3, 12}
}

func (x *AnalyzePlanResponse_GetStorageLevel) GetStorageLevel() *StorageLevel {
	if x != nil {
		return x.StorageLevel
	}
	return nil
}

type ExecutePlanRequest_RequestOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RequestOption:
	//
	//	*ExecutePlanRequest_RequestOption_ReattachOptions
	//	*ExecutePlanRequest_RequestOption_Extension
	RequestOption isExecutePlanRequest_RequestOption_RequestOption `protobuf_oneof:"request_option"`
}

func (x *ExecutePlanRequest_RequestOption) Reset() {
	*x = ExecutePlanRequest_RequestOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[45]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutePlanRequest_RequestOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePlanRequest_RequestOption) ProtoMessage() {}

func (x *ExecutePlanRequest_RequestOption) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[45]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePlanRequest_RequestOption.ProtoReflect.Descriptor instead.
func (*ExecutePlanRequest_RequestOption) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{4, 0}
}

func (m *ExecutePlanRequest_RequestOption) GetRequestOption() isExecutePlanRequest_RequestOption_RequestOption {
	if m != nil {
		return m.RequestOption
	}
	return nil
}

func (x *ExecutePlanRequest_RequestOption) GetReattachOptions() *ReattachOptions {
	if x, ok := x.GetRequestOption().(*ExecutePlanRequest_RequestOption_ReattachOptions); ok {
		return x.ReattachOptions
	}
	return nil
}

func (x *ExecutePlanRequest_RequestOption) GetExtension() *anypb.Any {
	if x, ok := x.GetRequestOption().(*ExecutePlanRequest_RequestOption_Extension); ok {
		return x.Extension
	}
	return nil
}

type isExecutePlanRequest_RequestOption_RequestOption interface {
	isExecutePlanRequest_RequestOption_RequestOption()
}

type ExecutePlanRequest_RequestOption_ReattachOptions struct {
	ReattachOptions *ReattachOptions `protobuf:"bytes,1,opt,name=reattach_options,json=reattachOptions,proto3,oneof"`
}

type ExecutePlanRequest_RequestOption_Extension struct {
	// Extension type for request options
	Extension *anypb.Any `protobuf:"bytes,999,opt,name=extension,proto3,oneof"`
}

func (*ExecutePlanRequest_RequestOption_ReattachOptions) isExecutePlanRequest_RequestOption_RequestOption() {
}

func (*ExecutePlanRequest_RequestOption_Extension) isExecutePlanRequest_RequestOption_RequestOption() {
}

// A SQL command returns an opaque Relation that can be directly used as input for the next
// call.
type ExecutePlanResponse_SqlCommandResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relation *Relation `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *ExecutePlanResponse_SqlCommandResult) Reset() {
	*x = ExecutePlanResponse_SqlCommandResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[46]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutePlanResponse_SqlCommandResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePlanResponse_SqlCommandResult) ProtoMessage() {}

func (x *ExecutePlanResponse_SqlCommandResult) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[46]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePlanResponse_SqlCommandResult.ProtoReflect.Descriptor instead.
func (*ExecutePlanResponse_SqlCommandResult) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ExecutePlanResponse_SqlCommandResult) GetRelation() *Relation {
	if x != nil {
		return x.Relation
	}
	return nil
}

// Batch results of metrics.
type ExecutePlanResponse_ArrowBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowCount int64  `protobuf:"varint,1,opt,name=row_count,json=rowCount,proto3" json:"row_count,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ExecutePlanResponse_ArrowBatch) Reset() {
	*x = ExecutePlanResponse_ArrowBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[47]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutePlanResponse_ArrowBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePlanResponse_ArrowBatch) ProtoMessage() {}

func (x *ExecutePlanResponse_ArrowBatch) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[47]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePlanResponse_ArrowBatch.ProtoReflect.Descriptor instead.
func (*ExecutePlanResponse_ArrowBatch) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{5, 1}
}

func (x *ExecutePlanResponse_ArrowBatch) GetRowCount() int64 {
	if x != nil {
		return x.RowCount
	}
	return 0
}

func (x *ExecutePlanResponse_ArrowBatch) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ExecutePlanResponse_Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*ExecutePlanResponse_Metrics_MetricObject `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *ExecutePlanResponse_Metrics) Reset() {
	*x = ExecutePlanResponse_Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[48]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutePlanResponse_Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePlanResponse_Metrics) ProtoMessage() {}

func (x *ExecutePlanResponse_Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[48]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePlanResponse_Metrics.ProtoReflect.Descriptor instead.
func (*ExecutePlanResponse_Metrics) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{5, 2}
}

func (x *ExecutePlanResponse_Metrics) GetMetrics() []*ExecutePlanResponse_Metrics_MetricObject {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type ExecutePlanResponse_ObservedMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Values []*Expression_Literal `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ExecutePlanResponse_ObservedMetrics) Reset() {
	*x = ExecutePlanResponse_ObservedMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[49]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutePlanResponse_ObservedMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePlanResponse_ObservedMetrics) ProtoMessage() {}

func (x *ExecutePlanResponse_ObservedMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[49]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePlanResponse_ObservedMetrics.ProtoReflect.Descriptor instead.
func (*ExecutePlanResponse_ObservedMetrics) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{5, 3}
}

func (x *ExecutePlanResponse_ObservedMetrics) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExecutePlanResponse_ObservedMetrics) GetValues() []*Expression_Literal {
	if x != nil {
		return x.Values
	}
	return nil
}

type ExecutePlanResponse_ResultComplete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExecutePlanResponse_ResultComplete) Reset() {
	*x = ExecutePlanResponse_ResultComplete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[50]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutePlanResponse_ResultComplete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePlanResponse_ResultComplete) ProtoMessage() {}

func (x *ExecutePlanResponse_ResultComplete) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[50]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePlanResponse_ResultComplete.ProtoReflect.Descriptor instead.
func (*ExecutePlanResponse_ResultComplete) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{5, 4}
}

type ExecutePlanResponse_Metrics_MetricObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string                                              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PlanId           int64                                               `protobuf:"varint,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Parent           int64                                               `protobuf:"varint,3,opt,name=parent,proto3" json:"parent,omitempty"`
	ExecutionMetrics map[string]*ExecutePlanResponse_Metrics_MetricValue `protobuf:"bytes,4,rep,name=execution_metrics,json=executionMetrics,proto3" json:"execution_metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExecutePlanResponse_Metrics_MetricObject) Reset() {
	*x = ExecutePlanResponse_Metrics_MetricObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[51]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutePlanResponse_Metrics_MetricObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePlanResponse_Metrics_MetricObject) ProtoMessage() {}

func (x *ExecutePlanResponse_Metrics_MetricObject) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[51]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePlanResponse_Metrics_MetricObject.ProtoReflect.Descriptor instead.
func (*ExecutePlanResponse_Metrics_MetricObject) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{5, 2, 0}
}

func (x *ExecutePlanResponse_Metrics_MetricObject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExecutePlanResponse_Metrics_MetricObject) GetPlanId() int64 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

func (x *ExecutePlanResponse_Metrics_MetricObject) GetParent() int64 {
	if x != nil {
		return x.Parent
	}
	return 0
}

func (x *ExecutePlanResponse_Metrics_MetricObject) GetExecutionMetrics() map[string]*ExecutePlanResponse_Metrics_MetricValue {
	if x != nil {
		return x.ExecutionMetrics
	}
	return nil
}

type ExecutePlanResponse_Metrics_MetricValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value      int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	MetricType string `protobuf:"bytes,3,opt,name=metric_type,json=metricType,proto3" json:"metric_type,omitempty"`
}

func (x *ExecutePlanResponse_Metrics_MetricValue) Reset() {
	*x = ExecutePlanResponse_Metrics_MetricValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[52]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutePlanResponse_Metrics_MetricValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePlanResponse_Metrics_MetricValue) ProtoMessage() {}

func (x *ExecutePlanResponse_Metrics_MetricValue) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[52]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePlanResponse_Metrics_MetricValue.ProtoReflect.Descriptor instead.
func (*ExecutePlanResponse_Metrics_MetricValue) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{5, 2, 1}
}

func (x *ExecutePlanResponse_Metrics_MetricValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExecutePlanResponse_Metrics_MetricValue) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ExecutePlanResponse_Metrics_MetricValue) GetMetricType() string {
	if x != nil {
		return x.MetricType
	}
	return ""
}

type ConfigRequest_Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to OpType:
	//
	//	*ConfigRequest_Operation_Set
	//	*ConfigRequest_Operation_Get
	//	*ConfigRequest_Operation_GetWithDefault
	//	*ConfigRequest_Operation_GetOption
	//	*ConfigRequest_Operation_GetAll
	//	*ConfigRequest_Operation_Unset
	//	*ConfigRequest_Operation_IsModifiable
	OpType isConfigRequest_Operation_OpType `protobuf_oneof:"op_type"`
}

func (x *ConfigRequest_Operation) Reset() {
	*x = ConfigRequest_Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[54]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_Operation) ProtoMessage() {}

func (x *ConfigRequest_Operation) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[54]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_Operation.ProtoReflect.Descriptor instead.
func (*ConfigRequest_Operation) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{7, 0}
}

func (m *ConfigRequest_Operation) GetOpType() isConfigRequest_Operation_OpType {
	if m != nil {
		return m.OpType
	}
	return nil
}

func (x *ConfigRequest_Operation) GetSet() *ConfigRequest_Set {
	if x, ok := x.GetOpType().(*ConfigRequest_Operation_Set); ok {
		return x.Set
	}
	return nil
}

func (x *ConfigRequest_Operation) GetGet() *ConfigRequest_Get {
	if x, ok := x.GetOpType().(*ConfigRequest_Operation_Get); ok {
		return x.Get
	}
	return nil
}

func (x *ConfigRequest_Operation) GetGetWithDefault() *ConfigRequest_GetWithDefault {
	if x, ok := x.GetOpType().(*ConfigRequest_Operation_GetWithDefault); ok {
		return x.GetWithDefault
	}
	return nil
}

func (x *ConfigRequest_Operation) GetGetOption() *ConfigRequest_GetOption {
	if x, ok := x.GetOpType().(*ConfigRequest_Operation_GetOption); ok {
		return x.GetOption
	}
	return nil
}

func (x *ConfigRequest_Operation) GetGetAll() *ConfigRequest_GetAll {
	if x, ok := x.GetOpType().(*ConfigRequest_Operation_GetAll); ok {
		return x.GetAll
	}
	return nil
}

func (x *ConfigRequest_Operation) GetUnset() *ConfigRequest_Unset {
	if x, ok := x.GetOpType().(*ConfigRequest_Operation_Unset); ok {
		return x.Unset
	}
	return nil
}

func (x *ConfigRequest_Operation) GetIsModifiable() *ConfigRequest_IsModifiable {
	if x, ok := x.GetOpType().(*ConfigRequest_Operation_IsModifiable); ok {
		return x.IsModifiable
	}
	return nil
}

type isConfigRequest_Operation_OpType interface {
	isConfigRequest_Operation_OpType()
}

type ConfigRequest_Operation_Set struct {
	Set *ConfigRequest_Set `protobuf:"bytes,1,opt,name=set,proto3,oneof"`
}

type ConfigRequest_Operation_Get struct {
	Get *ConfigRequest_Get `protobuf:"bytes,2,opt,name=get,proto3,oneof"`
}

type ConfigRequest_Operation_GetWithDefault struct {
	GetWithDefault *ConfigRequest_GetWithDefault `protobuf:"bytes,3,opt,name=get_with_default,json=getWithDefault,proto3,oneof"`
}

type ConfigRequest_Operation_GetOption struct {
	GetOption *ConfigRequest_GetOption `protobuf:"bytes,4,opt,name=get_option,json=getOption,proto3,oneof"`
}

type ConfigRequest_Operation_GetAll struct {
	GetAll *ConfigRequest_GetAll `protobuf:"bytes,5,opt,name=get_all,json=getAll,proto3,oneof"`
}

type ConfigRequest_Operation_Unset struct {
	Unset *ConfigRequest_Unset `protobuf:"bytes,6,opt,name=unset,proto3,oneof"`
}

type ConfigRequest_Operation_IsModifiable struct {
	IsModifiable *ConfigRequest_IsModifiable `protobuf:"bytes,7,opt,name=is_modifiable,json=isModifiable,proto3,oneof"`
}

func (*ConfigRequest_Operation_Set) isConfigRequest_Operation_OpType() {}

func (*ConfigRequest_Operation_Get) isConfigRequest_Operation_OpType() {}

func (*ConfigRequest_Operation_GetWithDefault) isConfigRequest_Operation_OpType() {}

func (*ConfigRequest_Operation_GetOption) isConfigRequest_Operation_OpType() {}

func (*ConfigRequest_Operation_GetAll) isConfigRequest_Operation_OpType() {}

func (*ConfigRequest_Operation_Unset) isConfigRequest_Operation_OpType() {}

func (*ConfigRequest_Operation_IsModifiable) isConfigRequest_Operation_OpType() {}

type ConfigRequest_Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The config key-value pairs to set.
	Pairs []*KeyValue `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *ConfigRequest_Set) Reset() {
	*x = ConfigRequest_Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[55]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_Set) ProtoMessage() {}

func (x *ConfigRequest_Set) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[55]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_Set.ProtoReflect.Descriptor instead.
func (*ConfigRequest_Set) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{7, 1}
}

func (x *ConfigRequest_Set) GetPairs() []*KeyValue {
	if x != nil {
		return x.Pairs
	}
	return nil
}

type ConfigRequest_Get struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The config keys to get.
	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *ConfigRequest_Get) Reset() {
	*x = ConfigRequest_Get{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[56]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_Get) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_Get) ProtoMessage() {}

func (x *ConfigRequest_Get) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[56]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_Get.ProtoReflect.Descriptor instead.
func (*ConfigRequest_Get) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{7, 2}
}

func (x *ConfigRequest_Get) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type ConfigRequest_GetWithDefault struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The config key-value paris to get. The value will be used as the default value.
	Pairs []*KeyValue `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *ConfigRequest_GetWithDefault) Reset() {
	*x = ConfigRequest_GetWithDefault{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[57]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_GetWithDefault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_GetWithDefault) ProtoMessage() {}

func (x *ConfigRequest_GetWithDefault) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[57]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_GetWithDefault.ProtoReflect.Descriptor instead.
func (*ConfigRequest_GetWithDefault) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{7, 3}
}

func (x *ConfigRequest_GetWithDefault) GetPairs() []*KeyValue {
	if x != nil {
		return x.Pairs
	}
	return nil
}

type ConfigRequest_GetOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The config keys to get optionally.
	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *ConfigRequest_GetOption) Reset() {
	*x = ConfigRequest_GetOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[58]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_GetOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_GetOption) ProtoMessage() {}

func (x *ConfigRequest_GetOption) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[58]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_GetOption.ProtoReflect.Descriptor instead.
func (*ConfigRequest_GetOption) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{7, 4}
}

func (x *ConfigRequest_GetOption) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type ConfigRequest_GetAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Optional) The prefix of the config key to get.
	Prefix *string `protobuf:"bytes,1,opt,name=prefix,proto3,oneof" json:"prefix,omitempty"`
}

func (x *ConfigRequest_GetAll) Reset() {
	*x = ConfigRequest_GetAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[59]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_GetAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_GetAll) ProtoMessage() {}

func (x *ConfigRequest_GetAll) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[59]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_GetAll.ProtoReflect.Descriptor instead.
func (*ConfigRequest_GetAll) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{7, 5}
}

func (x *ConfigRequest_GetAll) GetPrefix() string {
	if x != nil && x.Prefix != nil {
		return *x.Prefix
	}
	return ""
}

type ConfigRequest_Unset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The config keys to unset.
	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *ConfigRequest_Unset) Reset() {
	*x = ConfigRequest_Unset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[60]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_Unset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_Unset) ProtoMessage() {}

func (x *ConfigRequest_Unset) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[60]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_Unset.ProtoReflect.Descriptor instead.
func (*ConfigRequest_Unset) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{7, 6}
}

func (x *ConfigRequest_Unset) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type ConfigRequest_IsModifiable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The config keys to check the config is modifiable.
	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *ConfigRequest_IsModifiable) Reset() {
	*x = ConfigRequest_IsModifiable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[61]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_IsModifiable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_IsModifiable) ProtoMessage() {}

func (x *ConfigRequest_IsModifiable) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[61]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_IsModifiable.ProtoReflect.Descriptor instead.
func (*ConfigRequest_IsModifiable) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{7, 7}
}

func (x *ConfigRequest_IsModifiable) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

// A chunk of an Artifact.
type AddArtifactsRequest_ArtifactChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data chunk.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// CRC to allow server to verify integrity of the chunk.
	Crc int64 `protobuf:"varint,2,opt,name=crc,proto3" json:"crc,omitempty"`
}

func (x *AddArtifactsRequest_ArtifactChunk) Reset() {
	*x = AddArtifactsRequest_ArtifactChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[62]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddArtifactsRequest_ArtifactChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddArtifactsRequest_ArtifactChunk) ProtoMessage() {}

func (x *AddArtifactsRequest_ArtifactChunk) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[62]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddArtifactsRequest_ArtifactChunk.ProtoReflect.Descriptor instead.
func (*AddArtifactsRequest_ArtifactChunk) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{9, 0}
}

func (x *AddArtifactsRequest_ArtifactChunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AddArtifactsRequest_ArtifactChunk) GetCrc() int64 {
	if x != nil {
		return x.Crc
	}
	return 0
}

// An artifact that is contained in a single `ArtifactChunk`.
// Generally, this message represents tiny artifacts such as REPL-generatedCustom class files.
type AddArtifactsRequest_SingleChunkArtifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the artifact is expected in the form of a "Relative Path" that is made up of a
	// sequence of directories and the final file element.
	// Examples of "Relative Path"s: "jars/test.jar", "classes/xyz.class", "abc.xyz", "a/b/X.jar".
	// The server is expected to maintain the hierarchy of files as defined by their name. (i.e
	// The relative path of the file on the server's filesystem will be the same as the name of
	// the provided artifact)
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A single data chunk.
	Data *AddArtifactsRequest_ArtifactChunk `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AddArtifactsRequest_SingleChunkArtifact) Reset() {
	*x = AddArtifactsRequest_SingleChunkArtifact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[63]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddArtifactsRequest_SingleChunkArtifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddArtifactsRequest_SingleChunkArtifact) ProtoMessage() {}

func (x *AddArtifactsRequest_SingleChunkArtifact) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[63]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddArtifactsRequest_SingleChunkArtifact.ProtoReflect.Descriptor instead.
func (*AddArtifactsRequest_SingleChunkArtifact) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{9, 1}
}

func (x *AddArtifactsRequest_SingleChunkArtifact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddArtifactsRequest_SingleChunkArtifact) GetData() *AddArtifactsRequest_ArtifactChunk {
	if x != nil {
		return x.Data
	}
	return nil
}

// A number of `SingleChunkArtifact` batched into a single RPC.
type AddArtifactsRequest_Batch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Artifacts []*AddArtifactsRequest_SingleChunkArtifact `protobuf:"bytes,1,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
}

func (x *AddArtifactsRequest_Batch) Reset() {
	*x = AddArtifactsRequest_Batch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[64]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddArtifactsRequest_Batch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddArtifactsRequest_Batch) ProtoMessage() {}

func (x *AddArtifactsRequest_Batch) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[64]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddArtifactsRequest_Batch.ProtoReflect.Descriptor instead.
func (*AddArtifactsRequest_Batch) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{9, 2}
}

func (x *AddArtifactsRequest_Batch) GetArtifacts() []*AddArtifactsRequest_SingleChunkArtifact {
	if x != nil {
		return x.Artifacts
	}
	return nil
}

// Signals the beginning/start of a chunked artifact.
// A large artifact is transferred through a payload of `BeginChunkedArtifact` followed by a
// sequence of `ArtifactChunk`s.
type AddArtifactsRequest_BeginChunkedArtifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the artifact undergoing chunking. Follows the same conventions as the `name` in
	// the `Artifact` message.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Total size of the artifact in bytes.
	TotalBytes int64 `protobuf:"varint,2,opt,name=total_bytes,json=totalBytes,proto3" json:"total_bytes,omitempty"`
	// Number of chunks the artifact is split into.
	// This includes the `initial_chunk`.
	NumChunks int64 `protobuf:"varint,3,opt,name=num_chunks,json=numChunks,proto3" json:"num_chunks,omitempty"`
	// The first/initial chunk.
	InitialChunk *AddArtifactsRequest_ArtifactChunk `protobuf:"bytes,4,opt,name=initial_chunk,json=initialChunk,proto3" json:"initial_chunk,omitempty"`
}

func (x *AddArtifactsRequest_BeginChunkedArtifact) Reset() {
	*x = AddArtifactsRequest_BeginChunkedArtifact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[65]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddArtifactsRequest_BeginChunkedArtifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddArtifactsRequest_BeginChunkedArtifact) ProtoMessage() {}

func (x *AddArtifactsRequest_BeginChunkedArtifact) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[65]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddArtifactsRequest_BeginChunkedArtifact.ProtoReflect.Descriptor instead.
func (*AddArtifactsRequest_BeginChunkedArtifact) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{9, 3}
}

func (x *AddArtifactsRequest_BeginChunkedArtifact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddArtifactsRequest_BeginChunkedArtifact) GetTotalBytes() int64 {
	if x != nil {
		return x.TotalBytes
	}
	return 0
}

func (x *AddArtifactsRequest_BeginChunkedArtifact) GetNumChunks() int64 {
	if x != nil {
		return x.NumChunks
	}
	return 0
}

func (x *AddArtifactsRequest_BeginChunkedArtifact) GetInitialChunk() *AddArtifactsRequest_ArtifactChunk {
	if x != nil {
		return x.InitialChunk
	}
	return nil
}

// Metadata of an artifact.
type AddArtifactsResponse_ArtifactSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Whether the CRC (Cyclic Redundancy Check) is successful on server verification.
	// The server discards any artifact that fails the CRC.
	// If false, the client may choose to resend the artifact specified by `name`.
	IsCrcSuccessful bool `protobuf:"varint,2,opt,name=is_crc_successful,json=isCrcSuccessful,proto3" json:"is_crc_successful,omitempty"`
}

func (x *AddArtifactsResponse_ArtifactSummary) Reset() {
	*x = AddArtifactsResponse_ArtifactSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[66]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddArtifactsResponse_ArtifactSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddArtifactsResponse_ArtifactSummary) ProtoMessage() {}

func (x *AddArtifactsResponse_ArtifactSummary) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[66]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddArtifactsResponse_ArtifactSummary.ProtoReflect.Descriptor instead.
func (*AddArtifactsResponse_ArtifactSummary) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{10, 0}
}

func (x *AddArtifactsResponse_ArtifactSummary) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddArtifactsResponse_ArtifactSummary) GetIsCrcSuccessful() bool {
	if x != nil {
		return x.IsCrcSuccessful
	}
	return false
}

type ArtifactStatusesResponse_ArtifactStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Exists or not particular artifact at the server.
	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *ArtifactStatusesResponse_ArtifactStatus) Reset() {
	*x = ArtifactStatusesResponse_ArtifactStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[67]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactStatusesResponse_ArtifactStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactStatusesResponse_ArtifactStatus) ProtoMessage() {}

func (x *ArtifactStatusesResponse_ArtifactStatus) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[67]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactStatusesResponse_ArtifactStatus.ProtoReflect.Descriptor instead.
func (*ArtifactStatusesResponse_ArtifactStatus) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{12, 0}
}

func (x *ArtifactStatusesResponse_ArtifactStatus) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

// Release and close operation completely.
// This will also interrupt the query if it is running execution, and wait for it to be torn down.
type ReleaseExecuteRequest_ReleaseAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReleaseExecuteRequest_ReleaseAll) Reset() {
	*x = ReleaseExecuteRequest_ReleaseAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[69]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseExecuteRequest_ReleaseAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseExecuteRequest_ReleaseAll) ProtoMessage() {}

func (x *ReleaseExecuteRequest_ReleaseAll) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[69]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseExecuteRequest_ReleaseAll.ProtoReflect.Descriptor instead.
func (*ReleaseExecuteRequest_ReleaseAll) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{17, 0}
}

// Release all responses from the operation response stream up to and including
// the response with the given by response_id.
// While server determines by itself how much of a buffer of responses to keep, client providing
// explicit release calls will help reduce resource consumption.
// Noop if response_id not found in cached responses.
type ReleaseExecuteRequest_ReleaseUntil struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseId string `protobuf:"bytes,1,opt,name=response_id,json=responseId,proto3" json:"response_id,omitempty"`
}

func (x *ReleaseExecuteRequest_ReleaseUntil) Reset() {
	*x = ReleaseExecuteRequest_ReleaseUntil{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_base_proto_msgTypes[70]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseExecuteRequest_ReleaseUntil) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseExecuteRequest_ReleaseUntil) ProtoMessage() {}

func (x *ReleaseExecuteRequest_ReleaseUntil) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_base_proto_msgTypes[70]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseExecuteRequest_ReleaseUntil.ProtoReflect.Descriptor instead.
func (*ReleaseExecuteRequest_ReleaseUntil) Descriptor() ([]byte, []int) {
	return file_spark_connect_base_proto_rawDescGZIP(), []int{17, 1}
}

func (x *ReleaseExecuteRequest_ReleaseUntil) GetResponseId() string {
	if x != nil {
		return x.ResponseId
	}
	return ""
}

var File_spark_connect_base_proto protoreflect.FileDescriptor

var file_spark_connect_base_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x04, 0x50, 0x6c, 0x61,
	0x6e, 0x12, 0x2d, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74,
	0x12, 0x32, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x7a, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xe7, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xf5, 0x12, 0x0a, 0x12,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x24, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x45, 0x0a, 0x07, 0x65, 0x78,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45,
	0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x07, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x12, 0x4f, 0x0a, 0x0b, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x72, 0x65, 0x65, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x46, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x48,
	0x00, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x52, 0x0a, 0x0c, 0x69, 0x73,
	0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x49, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x48,
	0x00, 0x52, 0x0b, 0x69, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x4f,
	0x0a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12,
	0x55, 0x0a, 0x0d, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x09, 0x64, 0x64, 0x6c, 0x5f, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x44, 0x4c,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x48, 0x00, 0x52, 0x08, 0x64, 0x64, 0x6c, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x12, 0x58, 0x0a, 0x0e, 0x73, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74,
	0x69, 0x63, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x61, 0x6d,
	0x65, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x61,
	0x6d, 0x65, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x73, 0x12, 0x55, 0x0a, 0x0d, 0x73,
	0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x48, 0x61,
	0x73, 0x68, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x45, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x09, 0x75, 0x6e, 0x70,
	0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x55, 0x6e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x75, 0x6e, 0x70,
	0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x12, 0x5f, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x0f, 0x67, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0x31, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x12, 0x27, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x1a, 0xbb, 0x02, 0x0a, 0x07, 0x45,
	0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x12, 0x27, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12,
	0x58, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e,
	0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x65, 0x78,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x0b, 0x45, 0x78,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x58, 0x50,
	0x4c, 0x41, 0x49, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x58, 0x50, 0x4c, 0x41,
	0x49, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x19, 0x0a, 0x15, 0x45, 0x58, 0x50, 0x4c, 0x41, 0x49, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45,
	0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x45,
	0x58, 0x50, 0x4c, 0x41, 0x49, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x47, 0x45, 0x4e, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x58, 0x50, 0x4c, 0x41, 0x49, 0x4e,
	0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x4f, 0x53, 0x54, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16,
	0x45, 0x58, 0x50, 0x4c, 0x41, 0x49, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x4f, 0x52,
	0x4d, 0x41, 0x54, 0x54, 0x45, 0x44, 0x10, 0x05, 0x1a, 0x5a, 0x0a, 0x0a, 0x54, 0x72, 0x65, 0x65,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12,
	0x19, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x1a, 0x32, 0x0a, 0x07, 0x49, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12,
	0x27, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x1a, 0x36, 0x0a, 0x0b, 0x49, 0x73, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e,
	0x1a, 0x35, 0x0a, 0x0a, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x27,
	0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x1a, 0x0e, 0x0a, 0x0c, 0x53, 0x70, 0x61, 0x72, 0x6b,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x29, 0x0a, 0x08, 0x44, 0x44, 0x4c, 0x50, 0x61,
	0x72, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x64, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x64, 0x6c, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x1a, 0x79, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74,
	0x69, 0x63, 0x73, 0x12, 0x34, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x0a, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x32, 0x0a, 0x0a, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x09, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x1a, 0x37, 0x0a,
	0x0c, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x48, 0x61, 0x73, 0x68, 0x12, 0x27, 0x0a,
	0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x1a, 0x97, 0x01, 0x0a, 0x07, 0x50, 0x65, 0x72, 0x73, 0x69,
	0x73, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x0c, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x10,
	0x0a, 0x0e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x1a, 0x6e, 0x0a, 0x09, 0x55, 0x6e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a,
	0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x1a, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x7a, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x99, 0x0d, 0x0a, 0x13, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x48, 0x00, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12,
	0x46, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x07,
	0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x12, 0x50, 0x0a, 0x0b, 0x74, 0x72, 0x65, 0x65, 0x5f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x54, 0x72, 0x65, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x0a, 0x74,
	0x72, 0x65, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x47, 0x0a, 0x08, 0x69, 0x73, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x49, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x12, 0x53, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65,
	0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x73, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x0b, 0x69, 0x73, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x50, 0x0a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x0d, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x4a, 0x0a, 0x09, 0x64, 0x64, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x44, 0x4c, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x48, 0x00, 0x52, 0x08, 0x64, 0x64, 0x6c, 0x50, 0x61, 0x72, 0x73, 0x65, 0x12, 0x59, 0x0a,
	0x0e, 0x73, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x61, 0x6d, 0x65, 0x53, 0x65,
	0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x61, 0x6d, 0x65, 0x53,
	0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x73, 0x12, 0x56, 0x0a, 0x0d, 0x73, 0x65, 0x6d, 0x61,
	0x6e, 0x74, 0x69, 0x63, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x48, 0x61, 0x73, 0x68,
	0x48, 0x00, 0x52, 0x0c, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x46, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x07, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x09, 0x75, 0x6e, 0x70, 0x65,
	0x72, 0x73, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x55, 0x6e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x75, 0x6e, 0x70,
	0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x12, 0x60, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x0f, 0x67, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0x39, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x1a, 0x30, 0x0a, 0x07, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x12, 0x25,
	0x0a, 0x0e, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x2d, 0x0a, 0x0a, 0x54, 0x72, 0x65, 0x65, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x65, 0x65, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x1a, 0x24, 0x0a, 0x07, 0x49, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x1a, 0x30, 0x0a, 0x0b, 0x49, 0x73,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x69, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x1a, 0x22, 0x0a, 0x0a,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x1a, 0x28, 0x0a, 0x0c, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x3b, 0x0a, 0x08, 0x44, 0x44,
	0x4c, 0x50, 0x61, 0x72, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x1a, 0x27, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x65, 0x53,
	0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x1a, 0x26, 0x0a, 0x0c, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x09, 0x0a, 0x07, 0x50, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x1a, 0x0b, 0x0a, 0x09, 0x55, 0x6e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74,
	0x1a, 0x53, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x40, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0xa0, 0x04, 0x0a, 0x12, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x26, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x04,
	0x70, 0x6c, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x58, 0x0a, 0x0f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0xa5, 0x01, 0x0a, 0x0d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x10, 0x72,
	0x65, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x0f, 0x72, 0x65, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xe7, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x48, 0x00, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x10, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x99, 0x0f, 0x0a, 0x13, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x12, 0x50, 0x0a,
	0x0b, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x63, 0x0a, 0x12, 0x73, 0x71, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x53, 0x71, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x48, 0x00, 0x52, 0x10, 0x73, 0x71, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x7e, 0x0a, 0x23, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x48, 0x00, 0x52, 0x1f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x71, 0x0a, 0x1e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x1b, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x6b, 0x0a, 0x1c, 0x67, 0x65, 0x74, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x19, 0x67, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x87, 0x01, 0x0a, 0x26, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x22, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x5c,
	0x0a, 0x0f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x09,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xe7, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x5d, 0x0a, 0x10, 0x6f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x0f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x47, 0x0a, 0x10, 0x53, 0x71, 0x6c,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x33, 0x0a,
	0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x3d, 0x0a, 0x0a, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x85, 0x04, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x51, 0x0a,
	0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37,
	0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x1a, 0xcc, 0x02, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x7a, 0x0a, 0x11, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x4d, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x10, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x1a, 0x7b, 0x0a, 0x15, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x4c, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x58, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x60, 0x0a, 0x0f, 0x4f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x39, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x74, 0x65,
	0x72, 0x61, 0x6c, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x10, 0x0a, 0x0e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x0f, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x41,
	0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x84, 0x08, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x44, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x1a, 0xf2, 0x03,
	0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x03, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x03, 0x73, 0x65,
	0x74, 0x12, 0x34, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x48, 0x00, 0x52, 0x03, 0x67, 0x65, 0x74, 0x12, 0x57, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x77,
	0x69, 0x74, 0x68, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x48, 0x00,
	0x52, 0x0e, 0x67, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x12, 0x47, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x09,
	0x67, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x07, 0x67, 0x65, 0x74,
	0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48,
	0x00, 0x52, 0x06, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x3a, 0x0a, 0x05, 0x75, 0x6e, 0x73,
	0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x05,
	0x75, 0x6e, 0x73, 0x65, 0x74, 0x12, 0x50, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x73, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x73, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x1a, 0x34, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x70, 0x61, 0x69,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x1a, 0x19, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x1a, 0x3f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x70,
	0x61, 0x69, 0x72, 0x73, 0x1a, 0x1f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x65, 0x79, 0x73, 0x1a, 0x30, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x1b, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x1a, 0x1b, 0x0a, 0x05, 0x55, 0x6e, 0x73, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x1a, 0x22, 0x0a, 0x0c, 0x49, 0x73, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x7a, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x05, 0x70, 0x61, 0x69,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x73, 0x22, 0xe7, 0x06, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x40, 0x0a, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x41, 0x64, 0x64, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x48, 0x00, 0x52, 0x05, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x5a, 0x0a, 0x0b, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x65, 0x67, 0x69,
	0x6e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x48, 0x00, 0x52, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x48,
	0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x64,
	0x64, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48,
	0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x35, 0x0a, 0x0d, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x72, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x72, 0x63, 0x1a,
	0x6f, 0x0a, 0x13, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x5d, 0x0a, 0x05, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x54, 0x0a, 0x09, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x1a,
	0xc1, 0x01, 0x0a, 0x14, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x65, 0x64,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x55, 0x0a, 0x0d,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xbc,
	0x01, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52,
	0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x1a, 0x51, 0x0a, 0x0f, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x63, 0x72, 0x63, 0x5f, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73,
	0x43, 0x72, 0x63, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x22, 0xc3, 0x01,
	0x0a, 0x17, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x8c, 0x02, 0x0a, 0x18, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x51, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x65, 0x73, 0x1a, 0x28, 0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x1a, 0x73, 0x0a,
	0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x4c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xd8, 0x03, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x54, 0x0a, 0x0e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x25, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x67, 0x12, 0x23, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x80, 0x01,
	0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x52, 0x55, 0x50, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x16, 0x0a, 0x12, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x52, 0x55, 0x50, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x52, 0x55, 0x50, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x41, 0x47, 0x10, 0x02, 0x12,
	0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x52, 0x55, 0x50, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x03,
	0x42, 0x0b, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5b, 0x0a,
	0x11, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x65, 0x64,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x72, 0x75, 0x70, 0x74, 0x65, 0x64, 0x49, 0x64, 0x73, 0x22, 0x35, 0x0a, 0x0f, 0x52, 0x65,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x65, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c,
	0x65, 0x22, 0x93, 0x02, 0x0a, 0x16, 0x52, 0x65, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x0e, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x22, 0xc6, 0x03, 0x0a, 0x15, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x52, 0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x6c, 0x6c, 0x48, 0x00,
	0x52, 0x0a, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x58, 0x0a, 0x0d,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x1a, 0x0c, 0x0a, 0x0a, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x41, 0x6c, 0x6c, 0x1a, 0x2f, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x55,
	0x6e, 0x74, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x70, 0x0a, 0x16, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0c, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x32, 0xe7, 0x05, 0x0a, 0x13, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x56, 0x0a, 0x0b, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x12, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x12, 0x63, 0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x72, 0x75, 0x70, 0x74, 0x12, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0f, 0x52, 0x65, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5f, 0x0a, 0x0e, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x24, 0x2e,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x0a, 0x1e,
	0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spark_connect_base_proto_rawDescOnce sync.Once
	file_spark_connect_base_proto_rawDescData = file_spark_connect_base_proto_rawDesc
)

func file_spark_connect_base_proto_rawDescGZIP() []byte {
	file_spark_connect_base_proto_rawDescOnce.Do(func() {
		file_spark_connect_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_spark_connect_base_proto_rawDescData)
	})
	return file_spark_connect_base_proto_rawDescData
}

var file_spark_connect_base_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_spark_connect_base_proto_msgTypes = make([]protoimpl.MessageInfo, 71)
var file_spark_connect_base_proto_goTypes = []interface{}{
	(AnalyzePlanRequest_Explain_ExplainMode)(0),      // 0: spark.connect.AnalyzePlanRequest.Explain.ExplainMode
	(InterruptRequest_InterruptType)(0),              // 1: spark.connect.InterruptRequest.InterruptType
	(*Plan)(nil),                                     // 2: spark.connect.Plan
	(*UserContext)(nil),                              // 3: spark.connect.UserContext
	(*AnalyzePlanRequest)(nil),                       // 4: spark.connect.AnalyzePlanRequest
	(*AnalyzePlanResponse)(nil),                      // 5: spark.connect.AnalyzePlanResponse
	(*ExecutePlanRequest)(nil),                       // 6: spark.connect.ExecutePlanRequest
	(*ExecutePlanResponse)(nil),                      // 7: spark.connect.ExecutePlanResponse
	(*KeyValue)(nil),                                 // 8: spark.connect.KeyValue
	(*ConfigRequest)(nil),                            // 9: spark.connect.ConfigRequest
	(*ConfigResponse)(nil),                           // 10: spark.connect.ConfigResponse
	(*AddArtifactsRequest)(nil),                      // 11: spark.connect.AddArtifactsRequest
	(*AddArtifactsResponse)(nil),                     // 12: spark.connect.AddArtifactsResponse
	(*ArtifactStatusesRequest)(nil),                  // 13: spark.connect.ArtifactStatusesRequest
	(*ArtifactStatusesResponse)(nil),                 // 14: spark.connect.ArtifactStatusesResponse
	(*InterruptRequest)(nil),                         // 15: spark.connect.InterruptRequest
	(*InterruptResponse)(nil),                        // 16: spark.connect.InterruptResponse
	(*ReattachOptions)(nil),                          // 17: spark.connect.ReattachOptions
	(*ReattachExecuteRequest)(nil),                   // 18: spark.connect.ReattachExecuteRequest
	(*ReleaseExecuteRequest)(nil),                    // 19: spark.connect.ReleaseExecuteRequest
	(*ReleaseExecuteResponse)(nil),                   // 20: spark.connect.ReleaseExecuteResponse
	(*AnalyzePlanRequest_Schema)(nil),                // 21: spark.connect.AnalyzePlanRequest.Schema
	(*AnalyzePlanRequest_Explain)(nil),               // 22: spark.connect.AnalyzePlanRequest.Explain
	(*AnalyzePlanRequest_TreeString)(nil),            // 23: spark.connect.AnalyzePlanRequest.TreeString
	(*AnalyzePlanRequest_IsLocal)(nil),               // 24: spark.connect.AnalyzePlanRequest.IsLocal
	(*AnalyzePlanRequest_IsStreaming)(nil),           // 25: spark.connect.AnalyzePlanRequest.IsStreaming
	(*AnalyzePlanRequest_InputFiles)(nil),            // 26: spark.connect.AnalyzePlanRequest.InputFiles
	(*AnalyzePlanRequest_SparkVersion)(nil),          // 27: spark.connect.AnalyzePlanRequest.SparkVersion
	(*AnalyzePlanRequest_DDLParse)(nil),              // 28: spark.connect.AnalyzePlanRequest.DDLParse
	(*AnalyzePlanRequest_SameSemantics)(nil),         // 29: spark.connect.AnalyzePlanRequest.SameSemantics
	(*AnalyzePlanRequest_SemanticHash)(nil),          // 30: spark.connect.AnalyzePlanRequest.SemanticHash
	(*AnalyzePlanRequest_Persist)(nil),               // 31: spark.connect.AnalyzePlanRequest.Persist
	(*AnalyzePlanRequest_Unpersist)(nil),             // 32: spark.connect.AnalyzePlanRequest.Unpersist
	(*AnalyzePlanRequest_GetStorageLevel)(nil),       // 33: spark.connect.AnalyzePlanRequest.GetStorageLevel
	(*AnalyzePlanResponse_Schema)(nil),               // 34: spark.connect.AnalyzePlanResponse.Schema
	(*AnalyzePlanResponse_Explain)(nil),              // 35: spark.connect.AnalyzePlanResponse.Explain
	(*AnalyzePlanResponse_TreeString)(nil),           // 36: spark.connect.AnalyzePlanResponse.TreeString
	(*AnalyzePlanResponse_IsLocal)(nil),              // 37: spark.connect.AnalyzePlanResponse.IsLocal
	(*AnalyzePlanResponse_IsStreaming)(nil),          // 38: spark.connect.AnalyzePlanResponse.IsStreaming
	(*AnalyzePlanResponse_InputFiles)(nil),           // 39: spark.connect.AnalyzePlanResponse.InputFiles
	(*AnalyzePlanResponse_SparkVersion)(nil),         // 40: spark.connect.AnalyzePlanResponse.SparkVersion
	(*AnalyzePlanResponse_DDLParse)(nil),             // 41: spark.connect.AnalyzePlanResponse.DDLParse
	(*AnalyzePlanResponse_SameSemantics)(nil),        // 42: spark.connect.AnalyzePlanResponse.SameSemantics
	(*AnalyzePlanResponse_SemanticHash)(nil),         // 43: spark.connect.AnalyzePlanResponse.SemanticHash
	(*AnalyzePlanResponse_Persist)(nil),              // 44: spark.connect.AnalyzePlanResponse.Persist
	(*AnalyzePlanResponse_Unpersist)(nil),            // 45: spark.connect.AnalyzePlanResponse.Unpersist
	(*AnalyzePlanResponse_GetStorageLevel)(nil),      // 46: spark.connect.AnalyzePlanResponse.GetStorageLevel
	(*ExecutePlanRequest_RequestOption)(nil),         // 47: spark.connect.ExecutePlanRequest.RequestOption
	(*ExecutePlanResponse_SqlCommandResult)(nil),     // 48: spark.connect.ExecutePlanResponse.SqlCommandResult
	(*ExecutePlanResponse_ArrowBatch)(nil),           // 49: spark.connect.ExecutePlanResponse.ArrowBatch
	(*ExecutePlanResponse_Metrics)(nil),              // 50: spark.connect.ExecutePlanResponse.Metrics
	(*ExecutePlanResponse_ObservedMetrics)(nil),      // 51: spark.connect.ExecutePlanResponse.ObservedMetrics
	(*ExecutePlanResponse_ResultComplete)(nil),       // 52: spark.connect.ExecutePlanResponse.ResultComplete
	(*ExecutePlanResponse_Metrics_MetricObject)(nil), // 53: spark.connect.ExecutePlanResponse.Metrics.MetricObject
	(*ExecutePlanResponse_Metrics_MetricValue)(nil),  // 54: spark.connect.ExecutePlanResponse.Metrics.MetricValue
	nil,                                              // 55: spark.connect.ExecutePlanResponse.Metrics.MetricObject.ExecutionMetricsEntry
	(*ConfigRequest_Operation)(nil),                  // 56: spark.connect.ConfigRequest.Operation
	(*ConfigRequest_Set)(nil),                        // 57: spark.connect.ConfigRequest.Set
	(*ConfigRequest_Get)(nil),                        // 58: spark.connect.ConfigRequest.Get
	(*ConfigRequest_GetWithDefault)(nil),             // 59: spark.connect.ConfigRequest.GetWithDefault
	(*ConfigRequest_GetOption)(nil),                  // 60: spark.connect.ConfigRequest.GetOption
	(*ConfigRequest_GetAll)(nil),                     // 61: spark.connect.ConfigRequest.GetAll
	(*ConfigRequest_Unset)(nil),                      // 62: spark.connect.ConfigRequest.Unset
	(*ConfigRequest_IsModifiable)(nil),               // 63: spark.connect.ConfigRequest.IsModifiable
	(*AddArtifactsRequest_ArtifactChunk)(nil),        // 64: spark.connect.AddArtifactsRequest.ArtifactChunk
	(*AddArtifactsRequest_SingleChunkArtifact)(nil),  // 65: spark.connect.AddArtifactsRequest.SingleChunkArtifact
	(*AddArtifactsRequest_Batch)(nil),                // 66: spark.connect.AddArtifactsRequest.Batch
	(*AddArtifactsRequest_BeginChunkedArtifact)(nil), // 67: spark.connect.AddArtifactsRequest.BeginChunkedArtifact
	(*AddArtifactsResponse_ArtifactSummary)(nil),     // 68: spark.connect.AddArtifactsResponse.ArtifactSummary
	(*ArtifactStatusesResponse_ArtifactStatus)(nil),  // 69: spark.connect.ArtifactStatusesResponse.ArtifactStatus
	nil,                                      // 70: spark.connect.ArtifactStatusesResponse.StatusesEntry
	(*ReleaseExecuteRequest_ReleaseAll)(nil), // 71: spark.connect.ReleaseExecuteRequest.ReleaseAll
	(*ReleaseExecuteRequest_ReleaseUntil)(nil), // 72: spark.connect.ReleaseExecuteRequest.ReleaseUntil
	(*Relation)(nil),                           // 73: spark.connect.Relation
	(*Command)(nil),                            // 74: spark.connect.Command
	(*anypb.Any)(nil),                          // 75: google.protobuf.Any
	(*WriteStreamOperationStartResult)(nil),    // 76: spark.connect.WriteStreamOperationStartResult
	(*StreamingQueryCommandResult)(nil),        // 77: spark.connect.StreamingQueryCommandResult
	(*GetResourcesCommandResult)(nil),          // 78: spark.connect.GetResourcesCommandResult
	(*StreamingQueryManagerCommandResult)(nil), // 79: spark.connect.StreamingQueryManagerCommandResult
	(*DataType)(nil),                           // 80: spark.connect.DataType
	(*StorageLevel)(nil),                       // 81: spark.connect.StorageLevel
	(*Expression_Literal)(nil),                 // 82: spark.connect.Expression.Literal
}
var file_spark_connect_base_proto_depIdxs = []int32{
	73,  // 0: spark.connect.Plan.root:type_name -> spark.connect.Relation
	74,  // 1: spark.connect.Plan.command:type_name -> spark.connect.Command
	75,  // 2: spark.connect.UserContext.extensions:type_name -> google.protobuf.Any
	3,   // 3: spark.connect.AnalyzePlanRequest.user_context:type_name -> spark.connect.UserContext
	21,  // 4: spark.connect.AnalyzePlanRequest.schema:type_name -> spark.connect.AnalyzePlanRequest.Schema
	22,  // 5: spark.connect.AnalyzePlanRequest.explain:type_name -> spark.connect.AnalyzePlanRequest.Explain
	23,  // 6: spark.connect.AnalyzePlanRequest.tree_string:type_name -> spark.connect.AnalyzePlanRequest.TreeString
	24,  // 7: spark.connect.AnalyzePlanRequest.is_local:type_name -> spark.connect.AnalyzePlanRequest.IsLocal
	25,  // 8: spark.connect.AnalyzePlanRequest.is_streaming:type_name -> spark.connect.AnalyzePlanRequest.IsStreaming
	26,  // 9: spark.connect.AnalyzePlanRequest.input_files:type_name -> spark.connect.AnalyzePlanRequest.InputFiles
	27,  // 10: spark.connect.AnalyzePlanRequest.spark_version:type_name -> spark.connect.AnalyzePlanRequest.SparkVersion
	28,  // 11: spark.connect.AnalyzePlanRequest.ddl_parse:type_name -> spark.connect.AnalyzePlanRequest.DDLParse
	29,  // 12: spark.connect.AnalyzePlanRequest.same_semantics:type_name -> spark.connect.AnalyzePlanRequest.SameSemantics
	30,  // 13: spark.connect.AnalyzePlanRequest.semantic_hash:type_name -> spark.connect.AnalyzePlanRequest.SemanticHash
	31,  // 14: spark.connect.AnalyzePlanRequest.persist:type_name -> spark.connect.AnalyzePlanRequest.Persist
	32,  // 15: spark.connect.AnalyzePlanRequest.unpersist:type_name -> spark.connect.AnalyzePlanRequest.Unpersist
	33,  // 16: spark.connect.AnalyzePlanRequest.get_storage_level:type_name -> spark.connect.AnalyzePlanRequest.GetStorageLevel
	34,  // 17: spark.connect.AnalyzePlanResponse.schema:type_name -> spark.connect.AnalyzePlanResponse.Schema
	35,  // 18: spark.connect.AnalyzePlanResponse.explain:type_name -> spark.connect.AnalyzePlanResponse.Explain
	36,  // 19: spark.connect.AnalyzePlanResponse.tree_string:type_name -> spark.connect.AnalyzePlanResponse.TreeString
	37,  // 20: spark.connect.AnalyzePlanResponse.is_local:type_name -> spark.connect.AnalyzePlanResponse.IsLocal
	38,  // 21: spark.connect.AnalyzePlanResponse.is_streaming:type_name -> spark.connect.AnalyzePlanResponse.IsStreaming
	39,  // 22: spark.connect.AnalyzePlanResponse.input_files:type_name -> spark.connect.AnalyzePlanResponse.InputFiles
	40,  // 23: spark.connect.AnalyzePlanResponse.spark_version:type_name -> spark.connect.AnalyzePlanResponse.SparkVersion
	41,  // 24: spark.connect.AnalyzePlanResponse.ddl_parse:type_name -> spark.connect.AnalyzePlanResponse.DDLParse
	42,  // 25: spark.connect.AnalyzePlanResponse.same_semantics:type_name -> spark.connect.AnalyzePlanResponse.SameSemantics
	43,  // 26: spark.connect.AnalyzePlanResponse.semantic_hash:type_name -> spark.connect.AnalyzePlanResponse.SemanticHash
	44,  // 27: spark.connect.AnalyzePlanResponse.persist:type_name -> spark.connect.AnalyzePlanResponse.Persist
	45,  // 28: spark.connect.AnalyzePlanResponse.unpersist:type_name -> spark.connect.AnalyzePlanResponse.Unpersist
	46,  // 29: spark.connect.AnalyzePlanResponse.get_storage_level:type_name -> spark.connect.AnalyzePlanResponse.GetStorageLevel
	3,   // 30: spark.connect.ExecutePlanRequest.user_context:type_name -> spark.connect.UserContext
	2,   // 31: spark.connect.ExecutePlanRequest.plan:type_name -> spark.connect.Plan
	47,  // 32: spark.connect.ExecutePlanRequest.request_options:type_name -> spark.connect.ExecutePlanRequest.RequestOption
	49,  // 33: spark.connect.ExecutePlanResponse.arrow_batch:type_name -> spark.connect.ExecutePlanResponse.ArrowBatch
	48,  // 34: spark.connect.ExecutePlanResponse.sql_command_result:type_name -> spark.connect.ExecutePlanResponse.SqlCommandResult
	76,  // 35: spark.connect.ExecutePlanResponse.write_stream_operation_start_result:type_name -> spark.connect.WriteStreamOperationStartResult
	77,  // 36: spark.connect.ExecutePlanResponse.streaming_query_command_result:type_name -> spark.connect.StreamingQueryCommandResult
	78,  // 37: spark.connect.ExecutePlanResponse.get_resources_command_result:type_name -> spark.connect.GetResourcesCommandResult
	79,  // 38: spark.connect.ExecutePlanResponse.streaming_query_manager_command_result:type_name -> spark.connect.StreamingQueryManagerCommandResult
	52,  // 39: spark.connect.ExecutePlanResponse.result_complete:type_name -> spark.connect.ExecutePlanResponse.ResultComplete
	75,  // 40: spark.connect.ExecutePlanResponse.extension:type_name -> google.protobuf.Any
	50,  // 41: spark.connect.ExecutePlanResponse.metrics:type_name -> spark.connect.ExecutePlanResponse.Metrics
	51,  // 42: spark.connect.ExecutePlanResponse.observed_metrics:type_name -> spark.connect.ExecutePlanResponse.ObservedMetrics
	80,  // 43: spark.connect.ExecutePlanResponse.schema:type_name -> spark.connect.DataType
	3,   // 44: spark.connect.ConfigRequest.user_context:type_name -> spark.connect.UserContext
	56,  // 45: spark.connect.ConfigRequest.operation:type_name -> spark.connect.ConfigRequest.Operation
	8,   // 46: spark.connect.ConfigResponse.pairs:type_name -> spark.connect.KeyValue
	3,   // 47: spark.connect.AddArtifactsRequest.user_context:type_name -> spark.connect.UserContext
	66,  // 48: spark.connect.AddArtifactsRequest.batch:type_name -> spark.connect.AddArtifactsRequest.Batch
	67,  // 49: spark.connect.AddArtifactsRequest.begin_chunk:type_name -> spark.connect.AddArtifactsRequest.BeginChunkedArtifact
	64,  // 50: spark.connect.AddArtifactsRequest.chunk:type_name -> spark.connect.AddArtifactsRequest.ArtifactChunk
	68,  // 51: spark.connect.AddArtifactsResponse.artifacts:type_name -> spark.connect.AddArtifactsResponse.ArtifactSummary
	3,   // 52: spark.connect.ArtifactStatusesRequest.user_context:type_name -> spark.connect.UserContext
	70,  // 53: spark.connect.ArtifactStatusesResponse.statuses:type_name -> spark.connect.ArtifactStatusesResponse.StatusesEntry
	3,   // 54: spark.connect.InterruptRequest.user_context:type_name -> spark.connect.UserContext
	1,   // 55: spark.connect.InterruptRequest.interrupt_type:type_name -> spark.connect.InterruptRequest.InterruptType
	3,   // 56: spark.connect.ReattachExecuteRequest.user_context:type_name -> spark.connect.UserContext
	3,   // 57: spark.connect.ReleaseExecuteRequest.user_context:type_name -> spark.connect.UserContext
	71,  // 58: spark.connect.ReleaseExecuteRequest.release_all:type_name -> spark.connect.ReleaseExecuteRequest.ReleaseAll
	72,  // 59: spark.connect.ReleaseExecuteRequest.release_until:type_name -> spark.connect.ReleaseExecuteRequest.ReleaseUntil
	2,   // 60: spark.connect.AnalyzePlanRequest.Schema.plan:type_name -> spark.connect.Plan
	2,   // 61: spark.connect.AnalyzePlanRequest.Explain.plan:type_name -> spark.connect.Plan
	0,   // 62: spark.connect.AnalyzePlanRequest.Explain.explain_mode:type_name -> spark.connect.AnalyzePlanRequest.Explain.ExplainMode
	2,   // 63: spark.connect.AnalyzePlanRequest.TreeString.plan:type_name -> spark.connect.Plan
	2,   // 64: spark.connect.AnalyzePlanRequest.IsLocal.plan:type_name -> spark.connect.Plan
	2,   // 65: spark.connect.AnalyzePlanRequest.IsStreaming.plan:type_name -> spark.connect.Plan
	2,   // 66: spark.connect.AnalyzePlanRequest.InputFiles.plan:type_name -> spark.connect.Plan
	2,   // 67: spark.connect.AnalyzePlanRequest.SameSemantics.target_plan:type_name -> spark.connect.Plan
	2,   // 68: spark.connect.AnalyzePlanRequest.SameSemantics.other_plan:type_name -> spark.connect.Plan
	2,   // 69: spark.connect.AnalyzePlanRequest.SemanticHash.plan:type_name -> spark.connect.Plan
	73,  // 70: spark.connect.AnalyzePlanRequest.Persist.relation:type_name -> spark.connect.Relation
	81,  // 71: spark.connect.AnalyzePlanRequest.Persist.storage_level:type_name -> spark.connect.StorageLevel
	73,  // 72: spark.connect.AnalyzePlanRequest.Unpersist.relation:type_name -> spark.connect.Relation
	73,  // 73: spark.connect.AnalyzePlanRequest.GetStorageLevel.relation:type_name -> spark.connect.Relation
	80,  // 74: spark.connect.AnalyzePlanResponse.Schema.schema:type_name -> spark.connect.DataType
	80,  // 75: spark.connect.AnalyzePlanResponse.DDLParse.parsed:type_name -> spark.connect.DataType
	81,  // 76: spark.connect.AnalyzePlanResponse.GetStorageLevel.storage_level:type_name -> spark.connect.StorageLevel
	17,  // 77: spark.connect.ExecutePlanRequest.RequestOption.reattach_options:type_name -> spark.connect.ReattachOptions
	75,  // 78: spark.connect.ExecutePlanRequest.RequestOption.extension:type_name -> google.protobuf.Any
	73,  // 79: spark.connect.ExecutePlanResponse.SqlCommandResult.relation:type_name -> spark.connect.Relation
	53,  // 80: spark.connect.ExecutePlanResponse.Metrics.metrics:type_name -> spark.connect.ExecutePlanResponse.Metrics.MetricObject
	82,  // 81: spark.connect.ExecutePlanResponse.ObservedMetrics.values:type_name -> spark.connect.Expression.Literal
	55,  // 82: spark.connect.ExecutePlanResponse.Metrics.MetricObject.execution_metrics:type_name -> spark.connect.ExecutePlanResponse.Metrics.MetricObject.ExecutionMetricsEntry
	54,  // 83: spark.connect.ExecutePlanResponse.Metrics.MetricObject.ExecutionMetricsEntry.value:type_name -> spark.connect.ExecutePlanResponse.Metrics.MetricValue
	57,  // 84: spark.connect.ConfigRequest.Operation.set:type_name -> spark.connect.ConfigRequest.Set
	58,  // 85: spark.connect.ConfigRequest.Operation.get:type_name -> spark.connect.ConfigRequest.Get
	59,  // 86: spark.connect.ConfigRequest.Operation.get_with_default:type_name -> spark.connect.ConfigRequest.GetWithDefault
	60,  // 87: spark.connect.ConfigRequest.Operation.get_option:type_name -> spark.connect.ConfigRequest.GetOption
	61,  // 88: spark.connect.ConfigRequest.Operation.get_all:type_name -> spark.connect.ConfigRequest.GetAll
	62,  // 89: spark.connect.ConfigRequest.Operation.unset:type_name -> spark.connect.ConfigRequest.Unset
	63,  // 90: spark.connect.ConfigRequest.Operation.is_modifiable:type_name -> spark.connect.ConfigRequest.IsModifiable
	8,   // 91: spark.connect.ConfigRequest.Set.pairs:type_name -> spark.connect.KeyValue
	8,   // 92: spark.connect.ConfigRequest.GetWithDefault.pairs:type_name -> spark.connect.KeyValue
	64,  // 93: spark.connect.AddArtifactsRequest.SingleChunkArtifact.data:type_name -> spark.connect.AddArtifactsRequest.ArtifactChunk
	65,  // 94: spark.connect.AddArtifactsRequest.Batch.artifacts:type_name -> spark.connect.AddArtifactsRequest.SingleChunkArtifact
	64,  // 95: spark.connect.AddArtifactsRequest.BeginChunkedArtifact.initial_chunk:type_name -> spark.connect.AddArtifactsRequest.ArtifactChunk
	69,  // 96: spark.connect.ArtifactStatusesResponse.StatusesEntry.value:type_name -> spark.connect.ArtifactStatusesResponse.ArtifactStatus
	6,   // 97: spark.connect.SparkConnectService.ExecutePlan:input_type -> spark.connect.ExecutePlanRequest
	4,   // 98: spark.connect.SparkConnectService.AnalyzePlan:input_type -> spark.connect.AnalyzePlanRequest
	9,   // 99: spark.connect.SparkConnectService.Config:input_type -> spark.connect.ConfigRequest
	11,  // 100: spark.connect.SparkConnectService.AddArtifacts:input_type -> spark.connect.AddArtifactsRequest
	13,  // 101: spark.connect.SparkConnectService.ArtifactStatus:input_type -> spark.connect.ArtifactStatusesRequest
	15,  // 102: spark.connect.SparkConnectService.Interrupt:input_type -> spark.connect.InterruptRequest
	18,  // 103: spark.connect.SparkConnectService.ReattachExecute:input_type -> spark.connect.ReattachExecuteRequest
	19,  // 104: spark.connect.SparkConnectService.ReleaseExecute:input_type -> spark.connect.ReleaseExecuteRequest
	7,   // 105: spark.connect.SparkConnectService.ExecutePlan:output_type -> spark.connect.ExecutePlanResponse
	5,   // 106: spark.connect.SparkConnectService.AnalyzePlan:output_type -> spark.connect.AnalyzePlanResponse
	10,  // 107: spark.connect.SparkConnectService.Config:output_type -> spark.connect.ConfigResponse
	12,  // 108: spark.connect.SparkConnectService.AddArtifacts:output_type -> spark.connect.AddArtifactsResponse
	14,  // 109: spark.connect.SparkConnectService.ArtifactStatus:output_type -> spark.connect.ArtifactStatusesResponse
	16,  // 110: spark.connect.SparkConnectService.Interrupt:output_type -> spark.connect.InterruptResponse
	7,   // 111: spark.connect.SparkConnectService.ReattachExecute:output_type -> spark.connect.ExecutePlanResponse
	20,  // 112: spark.connect.SparkConnectService.ReleaseExecute:output_type -> spark.connect.ReleaseExecuteResponse
	105, // [105:113] is the sub-list for method output_type
	97,  // [97:105] is the sub-list for method input_type
	97,  // [97:97] is the sub-list for extension type_name
	97,  // [97:97] is the sub-list for extension extendee
	0,   // [0:97] is the sub-list for field type_name
}

func init() { file_spark_connect_base_proto_init() }
func file_spark_connect_base_proto_init() {
	if File_spark_connect_base_proto != nil {
		return
	}
	file_spark_connect_commands_proto_init()
	file_spark_connect_common_proto_init()
	file_spark_connect_expressions_proto_init()
	file_spark_connect_relations_proto_init()
	file_spark_connect_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spark_connect_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plan); i {
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
		file_spark_connect_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserContext); i {
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
		file_spark_connect_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanRequest); i {
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
		file_spark_connect_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanResponse); i {
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
		file_spark_connect_base_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutePlanRequest); i {
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
		file_spark_connect_base_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutePlanResponse); i {
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
		file_spark_connect_base_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
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
		file_spark_connect_base_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
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
		file_spark_connect_base_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigResponse); i {
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
		file_spark_connect_base_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddArtifactsRequest); i {
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
		file_spark_connect_base_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddArtifactsResponse); i {
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
		file_spark_connect_base_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactStatusesRequest); i {
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
		file_spark_connect_base_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactStatusesResponse); i {
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
		file_spark_connect_base_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterruptRequest); i {
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
		file_spark_connect_base_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterruptResponse); i {
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
		file_spark_connect_base_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReattachOptions); i {
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
		file_spark_connect_base_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReattachExecuteRequest); i {
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
		file_spark_connect_base_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseExecuteRequest); i {
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
		file_spark_connect_base_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseExecuteResponse); i {
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
		file_spark_connect_base_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanRequest_Schema); i {
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
		file_spark_connect_base_proto_msgTypes[20].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanRequest_Explain); i {
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
		file_spark_connect_base_proto_msgTypes[21].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanRequest_TreeString); i {
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
		file_spark_connect_base_proto_msgTypes[22].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanRequest_IsLocal); i {
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
		file_spark_connect_base_proto_msgTypes[23].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanRequest_IsStreaming); i {
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
		file_spark_connect_base_proto_msgTypes[24].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanRequest_InputFiles); i {
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
		file_spark_connect_base_proto_msgTypes[25].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanRequest_SparkVersion); i {
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
		file_spark_connect_base_proto_msgTypes[26].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanRequest_DDLParse); i {
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
		file_spark_connect_base_proto_msgTypes[27].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanRequest_SameSemantics); i {
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
		file_spark_connect_base_proto_msgTypes[28].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanRequest_SemanticHash); i {
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
		file_spark_connect_base_proto_msgTypes[29].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanRequest_Persist); i {
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
		file_spark_connect_base_proto_msgTypes[30].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanRequest_Unpersist); i {
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
		file_spark_connect_base_proto_msgTypes[31].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanRequest_GetStorageLevel); i {
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
		file_spark_connect_base_proto_msgTypes[32].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanResponse_Schema); i {
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
		file_spark_connect_base_proto_msgTypes[33].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanResponse_Explain); i {
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
		file_spark_connect_base_proto_msgTypes[34].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanResponse_TreeString); i {
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
		file_spark_connect_base_proto_msgTypes[35].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanResponse_IsLocal); i {
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
		file_spark_connect_base_proto_msgTypes[36].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanResponse_IsStreaming); i {
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
		file_spark_connect_base_proto_msgTypes[37].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanResponse_InputFiles); i {
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
		file_spark_connect_base_proto_msgTypes[38].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanResponse_SparkVersion); i {
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
		file_spark_connect_base_proto_msgTypes[39].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanResponse_DDLParse); i {
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
		file_spark_connect_base_proto_msgTypes[40].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanResponse_SameSemantics); i {
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
		file_spark_connect_base_proto_msgTypes[41].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanResponse_SemanticHash); i {
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
		file_spark_connect_base_proto_msgTypes[42].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanResponse_Persist); i {
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
		file_spark_connect_base_proto_msgTypes[43].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanResponse_Unpersist); i {
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
		file_spark_connect_base_proto_msgTypes[44].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzePlanResponse_GetStorageLevel); i {
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
		file_spark_connect_base_proto_msgTypes[45].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutePlanRequest_RequestOption); i {
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
		file_spark_connect_base_proto_msgTypes[46].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutePlanResponse_SqlCommandResult); i {
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
		file_spark_connect_base_proto_msgTypes[47].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutePlanResponse_ArrowBatch); i {
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
		file_spark_connect_base_proto_msgTypes[48].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutePlanResponse_Metrics); i {
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
		file_spark_connect_base_proto_msgTypes[49].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutePlanResponse_ObservedMetrics); i {
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
		file_spark_connect_base_proto_msgTypes[50].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutePlanResponse_ResultComplete); i {
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
		file_spark_connect_base_proto_msgTypes[51].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutePlanResponse_Metrics_MetricObject); i {
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
		file_spark_connect_base_proto_msgTypes[52].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutePlanResponse_Metrics_MetricValue); i {
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
		file_spark_connect_base_proto_msgTypes[54].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_Operation); i {
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
		file_spark_connect_base_proto_msgTypes[55].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_Set); i {
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
		file_spark_connect_base_proto_msgTypes[56].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_Get); i {
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
		file_spark_connect_base_proto_msgTypes[57].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_GetWithDefault); i {
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
		file_spark_connect_base_proto_msgTypes[58].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_GetOption); i {
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
		file_spark_connect_base_proto_msgTypes[59].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_GetAll); i {
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
		file_spark_connect_base_proto_msgTypes[60].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_Unset); i {
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
		file_spark_connect_base_proto_msgTypes[61].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_IsModifiable); i {
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
		file_spark_connect_base_proto_msgTypes[62].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddArtifactsRequest_ArtifactChunk); i {
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
		file_spark_connect_base_proto_msgTypes[63].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddArtifactsRequest_SingleChunkArtifact); i {
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
		file_spark_connect_base_proto_msgTypes[64].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddArtifactsRequest_Batch); i {
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
		file_spark_connect_base_proto_msgTypes[65].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddArtifactsRequest_BeginChunkedArtifact); i {
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
		file_spark_connect_base_proto_msgTypes[66].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddArtifactsResponse_ArtifactSummary); i {
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
		file_spark_connect_base_proto_msgTypes[67].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactStatusesResponse_ArtifactStatus); i {
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
		file_spark_connect_base_proto_msgTypes[69].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseExecuteRequest_ReleaseAll); i {
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
		file_spark_connect_base_proto_msgTypes[70].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseExecuteRequest_ReleaseUntil); i {
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
	file_spark_connect_base_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Plan_Root)(nil),
		(*Plan_Command)(nil),
	}
	file_spark_connect_base_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*AnalyzePlanRequest_Schema_)(nil),
		(*AnalyzePlanRequest_Explain_)(nil),
		(*AnalyzePlanRequest_TreeString_)(nil),
		(*AnalyzePlanRequest_IsLocal_)(nil),
		(*AnalyzePlanRequest_IsStreaming_)(nil),
		(*AnalyzePlanRequest_InputFiles_)(nil),
		(*AnalyzePlanRequest_SparkVersion_)(nil),
		(*AnalyzePlanRequest_DdlParse)(nil),
		(*AnalyzePlanRequest_SameSemantics_)(nil),
		(*AnalyzePlanRequest_SemanticHash_)(nil),
		(*AnalyzePlanRequest_Persist_)(nil),
		(*AnalyzePlanRequest_Unpersist_)(nil),
		(*AnalyzePlanRequest_GetStorageLevel_)(nil),
	}
	file_spark_connect_base_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*AnalyzePlanResponse_Schema_)(nil),
		(*AnalyzePlanResponse_Explain_)(nil),
		(*AnalyzePlanResponse_TreeString_)(nil),
		(*AnalyzePlanResponse_IsLocal_)(nil),
		(*AnalyzePlanResponse_IsStreaming_)(nil),
		(*AnalyzePlanResponse_InputFiles_)(nil),
		(*AnalyzePlanResponse_SparkVersion_)(nil),
		(*AnalyzePlanResponse_DdlParse)(nil),
		(*AnalyzePlanResponse_SameSemantics_)(nil),
		(*AnalyzePlanResponse_SemanticHash_)(nil),
		(*AnalyzePlanResponse_Persist_)(nil),
		(*AnalyzePlanResponse_Unpersist_)(nil),
		(*AnalyzePlanResponse_GetStorageLevel_)(nil),
	}
	file_spark_connect_base_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_spark_connect_base_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ExecutePlanResponse_ArrowBatch_)(nil),
		(*ExecutePlanResponse_SqlCommandResult_)(nil),
		(*ExecutePlanResponse_WriteStreamOperationStartResult)(nil),
		(*ExecutePlanResponse_StreamingQueryCommandResult)(nil),
		(*ExecutePlanResponse_GetResourcesCommandResult)(nil),
		(*ExecutePlanResponse_StreamingQueryManagerCommandResult)(nil),
		(*ExecutePlanResponse_ResultComplete_)(nil),
		(*ExecutePlanResponse_Extension)(nil),
	}
	file_spark_connect_base_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_spark_connect_base_proto_msgTypes[7].OneofWrappers = []interface{}{}
	file_spark_connect_base_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*AddArtifactsRequest_Batch_)(nil),
		(*AddArtifactsRequest_BeginChunk)(nil),
		(*AddArtifactsRequest_Chunk)(nil),
	}
	file_spark_connect_base_proto_msgTypes[11].OneofWrappers = []interface{}{}
	file_spark_connect_base_proto_msgTypes[13].OneofWrappers = []interface{}{
		(*InterruptRequest_OperationTag)(nil),
		(*InterruptRequest_OperationId)(nil),
	}
	file_spark_connect_base_proto_msgTypes[16].OneofWrappers = []interface{}{}
	file_spark_connect_base_proto_msgTypes[17].OneofWrappers = []interface{}{
		(*ReleaseExecuteRequest_ReleaseAll_)(nil),
		(*ReleaseExecuteRequest_ReleaseUntil_)(nil),
	}
	file_spark_connect_base_proto_msgTypes[18].OneofWrappers = []interface{}{}
	file_spark_connect_base_proto_msgTypes[21].OneofWrappers = []interface{}{}
	file_spark_connect_base_proto_msgTypes[29].OneofWrappers = []interface{}{}
	file_spark_connect_base_proto_msgTypes[30].OneofWrappers = []interface{}{}
	file_spark_connect_base_proto_msgTypes[45].OneofWrappers = []interface{}{
		(*ExecutePlanRequest_RequestOption_ReattachOptions)(nil),
		(*ExecutePlanRequest_RequestOption_Extension)(nil),
	}
	file_spark_connect_base_proto_msgTypes[54].OneofWrappers = []interface{}{
		(*ConfigRequest_Operation_Set)(nil),
		(*ConfigRequest_Operation_Get)(nil),
		(*ConfigRequest_Operation_GetWithDefault)(nil),
		(*ConfigRequest_Operation_GetOption)(nil),
		(*ConfigRequest_Operation_GetAll)(nil),
		(*ConfigRequest_Operation_Unset)(nil),
		(*ConfigRequest_Operation_IsModifiable)(nil),
	}
	file_spark_connect_base_proto_msgTypes[59].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spark_connect_base_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   71,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spark_connect_base_proto_goTypes,
		DependencyIndexes: file_spark_connect_base_proto_depIdxs,
		EnumInfos:         file_spark_connect_base_proto_enumTypes,
		MessageInfos:      file_spark_connect_base_proto_msgTypes,
	}.Build()
	File_spark_connect_base_proto = out.File
	file_spark_connect_base_proto_rawDesc = nil
	file_spark_connect_base_proto_goTypes = nil
	file_spark_connect_base_proto_depIdxs = nil
}
