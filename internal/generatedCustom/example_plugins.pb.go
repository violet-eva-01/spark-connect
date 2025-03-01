package generatedCustom

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generatedCustom code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExamplePluginRelation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input       *Relation `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	CustomField string    `protobuf:"bytes,2,opt,name=custom_field,json=customField,proto3" json:"custom_field,omitempty"`
}

func (x *ExamplePluginRelation) Reset() {
	*x = ExamplePluginRelation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_example_plugins_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExamplePluginRelation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExamplePluginRelation) ProtoMessage() {}

func (x *ExamplePluginRelation) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_example_plugins_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExamplePluginRelation.ProtoReflect.Descriptor instead.
func (*ExamplePluginRelation) Descriptor() ([]byte, []int) {
	return file_spark_connect_example_plugins_proto_rawDescGZIP(), []int{0}
}

func (x *ExamplePluginRelation) GetInput() *Relation {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *ExamplePluginRelation) GetCustomField() string {
	if x != nil {
		return x.CustomField
	}
	return ""
}

type ExamplePluginExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Child       *Expression `protobuf:"bytes,1,opt,name=child,proto3" json:"child,omitempty"`
	CustomField string      `protobuf:"bytes,2,opt,name=custom_field,json=customField,proto3" json:"custom_field,omitempty"`
}

func (x *ExamplePluginExpression) Reset() {
	*x = ExamplePluginExpression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_example_plugins_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExamplePluginExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExamplePluginExpression) ProtoMessage() {}

func (x *ExamplePluginExpression) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_example_plugins_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExamplePluginExpression.ProtoReflect.Descriptor instead.
func (*ExamplePluginExpression) Descriptor() ([]byte, []int) {
	return file_spark_connect_example_plugins_proto_rawDescGZIP(), []int{1}
}

func (x *ExamplePluginExpression) GetChild() *Expression {
	if x != nil {
		return x.Child
	}
	return nil
}

func (x *ExamplePluginExpression) GetCustomField() string {
	if x != nil {
		return x.CustomField
	}
	return ""
}

type ExamplePluginCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomField string `protobuf:"bytes,1,opt,name=custom_field,json=customField,proto3" json:"custom_field,omitempty"`
}

func (x *ExamplePluginCommand) Reset() {
	*x = ExamplePluginCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_example_plugins_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExamplePluginCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExamplePluginCommand) ProtoMessage() {}

func (x *ExamplePluginCommand) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_example_plugins_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExamplePluginCommand.ProtoReflect.Descriptor instead.
func (*ExamplePluginCommand) Descriptor() ([]byte, []int) {
	return file_spark_connect_example_plugins_proto_rawDescGZIP(), []int{2}
}

func (x *ExamplePluginCommand) GetCustomField() string {
	if x != nil {
		return x.CustomField
	}
	return ""
}

var File_spark_connect_example_plugins_proto protoreflect.FileDescriptor

var file_spark_connect_example_plugins_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x1a, 0x1d, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x15, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22,
	0x6d, 0x0a, 0x17, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x05, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x39,
	0x0a, 0x14, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x36, 0x0a, 0x1e, 0x6f, 0x72, 0x67,
	0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x12, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spark_connect_example_plugins_proto_rawDescOnce sync.Once
	file_spark_connect_example_plugins_proto_rawDescData = file_spark_connect_example_plugins_proto_rawDesc
)

func file_spark_connect_example_plugins_proto_rawDescGZIP() []byte {
	file_spark_connect_example_plugins_proto_rawDescOnce.Do(func() {
		file_spark_connect_example_plugins_proto_rawDescData = protoimpl.X.CompressGZIP(file_spark_connect_example_plugins_proto_rawDescData)
	})
	return file_spark_connect_example_plugins_proto_rawDescData
}

var file_spark_connect_example_plugins_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spark_connect_example_plugins_proto_goTypes = []interface{}{
	(*ExamplePluginRelation)(nil),   // 0: spark.connect.ExamplePluginRelation
	(*ExamplePluginExpression)(nil), // 1: spark.connect.ExamplePluginExpression
	(*ExamplePluginCommand)(nil),    // 2: spark.connect.ExamplePluginCommand
	(*Relation)(nil),                // 3: spark.connect.Relation
	(*Expression)(nil),              // 4: spark.connect.Expression
}
var file_spark_connect_example_plugins_proto_depIdxs = []int32{
	3, // 0: spark.connect.ExamplePluginRelation.input:type_name -> spark.connect.Relation
	4, // 1: spark.connect.ExamplePluginExpression.child:type_name -> spark.connect.Expression
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spark_connect_example_plugins_proto_init() }
func file_spark_connect_example_plugins_proto_init() {
	if File_spark_connect_example_plugins_proto != nil {
		return
	}
	file_spark_connect_relations_proto_init()
	file_spark_connect_expressions_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spark_connect_example_plugins_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExamplePluginRelation); i {
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
		file_spark_connect_example_plugins_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExamplePluginExpression); i {
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
		file_spark_connect_example_plugins_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExamplePluginCommand); i {
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
			RawDescriptor: file_spark_connect_example_plugins_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spark_connect_example_plugins_proto_goTypes,
		DependencyIndexes: file_spark_connect_example_plugins_proto_depIdxs,
		MessageInfos:      file_spark_connect_example_plugins_proto_msgTypes,
	}.Build()
	File_spark_connect_example_plugins_proto = out.File
	file_spark_connect_example_plugins_proto_rawDesc = nil
	file_spark_connect_example_plugins_proto_goTypes = nil
	file_spark_connect_example_plugins_proto_depIdxs = nil
}
