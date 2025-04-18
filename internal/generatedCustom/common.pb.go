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

// StorageLevel for persisting Datasets/Tables.
type StorageLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) Whether the cache should use disk or not.
	UseDisk bool `protobuf:"varint,1,opt,name=use_disk,json=useDisk,proto3" json:"use_disk,omitempty"`
	// (Required) Whether the cache should use memory or not.
	UseMemory bool `protobuf:"varint,2,opt,name=use_memory,json=useMemory,proto3" json:"use_memory,omitempty"`
	// (Required) Whether the cache should use off-heap or not.
	UseOffHeap bool `protobuf:"varint,3,opt,name=use_off_heap,json=useOffHeap,proto3" json:"use_off_heap,omitempty"`
	// (Required) Whether the cached data is deserialized or not.
	Deserialized bool `protobuf:"varint,4,opt,name=deserialized,proto3" json:"deserialized,omitempty"`
	// (Required) The number of replicas.
	Replication int32 `protobuf:"varint,5,opt,name=replication,proto3" json:"replication,omitempty"`
}

func (x *StorageLevel) Reset() {
	*x = StorageLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageLevel) ProtoMessage() {}

func (x *StorageLevel) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageLevel.ProtoReflect.Descriptor instead.
func (*StorageLevel) Descriptor() ([]byte, []int) {
	return file_spark_connect_common_proto_rawDescGZIP(), []int{0}
}

func (x *StorageLevel) GetUseDisk() bool {
	if x != nil {
		return x.UseDisk
	}
	return false
}

func (x *StorageLevel) GetUseMemory() bool {
	if x != nil {
		return x.UseMemory
	}
	return false
}

func (x *StorageLevel) GetUseOffHeap() bool {
	if x != nil {
		return x.UseOffHeap
	}
	return false
}

func (x *StorageLevel) GetDeserialized() bool {
	if x != nil {
		return x.Deserialized
	}
	return false
}

func (x *StorageLevel) GetReplication() int32 {
	if x != nil {
		return x.Replication
	}
	return 0
}

// ResourceInformation to hold information about a type of Resource.
// The corresponding class is 'org.apache.spark.resource.ResourceInformation'
type ResourceInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The name of the resource
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Required) An array of strings describing the addresses of the resource.
	Addresses []string `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *ResourceInformation) Reset() {
	*x = ResourceInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceInformation) ProtoMessage() {}

func (x *ResourceInformation) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceInformation.ProtoReflect.Descriptor instead.
func (*ResourceInformation) Descriptor() ([]byte, []int) {
	return file_spark_connect_common_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceInformation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceInformation) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

var File_spark_connect_common_proto protoreflect.FileDescriptor

var file_spark_connect_common_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x22, 0xb0, 0x01, 0x0a, 0x0c,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x75, 0x73, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x5f, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x5f, 0x6f, 0x66,
	0x66, 0x5f, 0x68, 0x65, 0x61, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x73,
	0x65, 0x4f, 0x66, 0x66, 0x48, 0x65, 0x61, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x64, 0x65, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x47,
	0x0a, 0x13, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x42, 0x36, 0x0a, 0x1e, 0x6f, 0x72, 0x67, 0x2e, 0x61,
	0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x12, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spark_connect_common_proto_rawDescOnce sync.Once
	file_spark_connect_common_proto_rawDescData = file_spark_connect_common_proto_rawDesc
)

func file_spark_connect_common_proto_rawDescGZIP() []byte {
	file_spark_connect_common_proto_rawDescOnce.Do(func() {
		file_spark_connect_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_spark_connect_common_proto_rawDescData)
	})
	return file_spark_connect_common_proto_rawDescData
}

var file_spark_connect_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spark_connect_common_proto_goTypes = []interface{}{
	(*StorageLevel)(nil),        // 0: spark.connect.StorageLevel
	(*ResourceInformation)(nil), // 1: spark.connect.ResourceInformation
}
var file_spark_connect_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spark_connect_common_proto_init() }
func file_spark_connect_common_proto_init() {
	if File_spark_connect_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spark_connect_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageLevel); i {
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
		file_spark_connect_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceInformation); i {
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
			RawDescriptor: file_spark_connect_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spark_connect_common_proto_goTypes,
		DependencyIndexes: file_spark_connect_common_proto_depIdxs,
		MessageInfos:      file_spark_connect_common_proto_msgTypes,
	}.Build()
	File_spark_connect_common_proto = out.File
	file_spark_connect_common_proto_rawDesc = nil
	file_spark_connect_common_proto_goTypes = nil
	file_spark_connect_common_proto_depIdxs = nil
}
