package generated

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

// This message describes the logical [[DataType]] of something. It does not carry the value
// itself but only describes it.
type DataType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*DataType_Null
	//	*DataType_Binary_
	//	*DataType_Boolean_
	//	*DataType_Byte_
	//	*DataType_Short_
	//	*DataType_Integer_
	//	*DataType_Long_
	//	*DataType_Float_
	//	*DataType_Double_
	//	*DataType_Decimal_
	//	*DataType_String_
	//	*DataType_Char_
	//	*DataType_VarChar_
	//	*DataType_Date_
	//	*DataType_Timestamp_
	//	*DataType_TimestampNtz
	//	*DataType_CalendarInterval_
	//	*DataType_YearMonthInterval_
	//	*DataType_DayTimeInterval_
	//	*DataType_Array_
	//	*DataType_Struct_
	//	*DataType_Map_
	//	*DataType_Udt
	//	*DataType_Unparsed_
	Kind isDataType_Kind `protobuf_oneof:"kind"`
}

func (x *DataType) Reset() {
	*x = DataType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType) ProtoMessage() {}

func (x *DataType) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType.ProtoReflect.Descriptor instead.
func (*DataType) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0}
}

func (m *DataType) GetKind() isDataType_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *DataType) GetNull() *DataType_NULL {
	if x, ok := x.GetKind().(*DataType_Null); ok {
		return x.Null
	}
	return nil
}

func (x *DataType) GetBinary() *DataType_Binary {
	if x, ok := x.GetKind().(*DataType_Binary_); ok {
		return x.Binary
	}
	return nil
}

func (x *DataType) GetBoolean() *DataType_Boolean {
	if x, ok := x.GetKind().(*DataType_Boolean_); ok {
		return x.Boolean
	}
	return nil
}

func (x *DataType) GetByte() *DataType_Byte {
	if x, ok := x.GetKind().(*DataType_Byte_); ok {
		return x.Byte
	}
	return nil
}

func (x *DataType) GetShort() *DataType_Short {
	if x, ok := x.GetKind().(*DataType_Short_); ok {
		return x.Short
	}
	return nil
}

func (x *DataType) GetInteger() *DataType_Integer {
	if x, ok := x.GetKind().(*DataType_Integer_); ok {
		return x.Integer
	}
	return nil
}

func (x *DataType) GetLong() *DataType_Long {
	if x, ok := x.GetKind().(*DataType_Long_); ok {
		return x.Long
	}
	return nil
}

func (x *DataType) GetFloat() *DataType_Float {
	if x, ok := x.GetKind().(*DataType_Float_); ok {
		return x.Float
	}
	return nil
}

func (x *DataType) GetDouble() *DataType_Double {
	if x, ok := x.GetKind().(*DataType_Double_); ok {
		return x.Double
	}
	return nil
}

func (x *DataType) GetDecimal() *DataType_Decimal {
	if x, ok := x.GetKind().(*DataType_Decimal_); ok {
		return x.Decimal
	}
	return nil
}

func (x *DataType) GetString_() *DataType_String {
	if x, ok := x.GetKind().(*DataType_String_); ok {
		return x.String_
	}
	return nil
}

func (x *DataType) GetChar() *DataType_Char {
	if x, ok := x.GetKind().(*DataType_Char_); ok {
		return x.Char
	}
	return nil
}

func (x *DataType) GetVarChar() *DataType_VarChar {
	if x, ok := x.GetKind().(*DataType_VarChar_); ok {
		return x.VarChar
	}
	return nil
}

func (x *DataType) GetDate() *DataType_Date {
	if x, ok := x.GetKind().(*DataType_Date_); ok {
		return x.Date
	}
	return nil
}

func (x *DataType) GetTimestamp() *DataType_Timestamp {
	if x, ok := x.GetKind().(*DataType_Timestamp_); ok {
		return x.Timestamp
	}
	return nil
}

func (x *DataType) GetTimestampNtz() *DataType_TimestampNTZ {
	if x, ok := x.GetKind().(*DataType_TimestampNtz); ok {
		return x.TimestampNtz
	}
	return nil
}

func (x *DataType) GetCalendarInterval() *DataType_CalendarInterval {
	if x, ok := x.GetKind().(*DataType_CalendarInterval_); ok {
		return x.CalendarInterval
	}
	return nil
}

func (x *DataType) GetYearMonthInterval() *DataType_YearMonthInterval {
	if x, ok := x.GetKind().(*DataType_YearMonthInterval_); ok {
		return x.YearMonthInterval
	}
	return nil
}

func (x *DataType) GetDayTimeInterval() *DataType_DayTimeInterval {
	if x, ok := x.GetKind().(*DataType_DayTimeInterval_); ok {
		return x.DayTimeInterval
	}
	return nil
}

func (x *DataType) GetArray() *DataType_Array {
	if x, ok := x.GetKind().(*DataType_Array_); ok {
		return x.Array
	}
	return nil
}

func (x *DataType) GetStruct() *DataType_Struct {
	if x, ok := x.GetKind().(*DataType_Struct_); ok {
		return x.Struct
	}
	return nil
}

func (x *DataType) GetMap() *DataType_Map {
	if x, ok := x.GetKind().(*DataType_Map_); ok {
		return x.Map
	}
	return nil
}

func (x *DataType) GetUdt() *DataType_UDT {
	if x, ok := x.GetKind().(*DataType_Udt); ok {
		return x.Udt
	}
	return nil
}

func (x *DataType) GetUnparsed() *DataType_Unparsed {
	if x, ok := x.GetKind().(*DataType_Unparsed_); ok {
		return x.Unparsed
	}
	return nil
}

type isDataType_Kind interface {
	isDataType_Kind()
}

type DataType_Null struct {
	Null *DataType_NULL `protobuf:"bytes,1,opt,name=null,proto3,oneof"`
}

type DataType_Binary_ struct {
	Binary *DataType_Binary `protobuf:"bytes,2,opt,name=binary,proto3,oneof"`
}

type DataType_Boolean_ struct {
	Boolean *DataType_Boolean `protobuf:"bytes,3,opt,name=boolean,proto3,oneof"`
}

type DataType_Byte_ struct {
	// Numeric types
	Byte *DataType_Byte `protobuf:"bytes,4,opt,name=byte,proto3,oneof"`
}

type DataType_Short_ struct {
	Short *DataType_Short `protobuf:"bytes,5,opt,name=short,proto3,oneof"`
}

type DataType_Integer_ struct {
	Integer *DataType_Integer `protobuf:"bytes,6,opt,name=integer,proto3,oneof"`
}

type DataType_Long_ struct {
	Long *DataType_Long `protobuf:"bytes,7,opt,name=long,proto3,oneof"`
}

type DataType_Float_ struct {
	Float *DataType_Float `protobuf:"bytes,8,opt,name=float,proto3,oneof"`
}

type DataType_Double_ struct {
	Double *DataType_Double `protobuf:"bytes,9,opt,name=double,proto3,oneof"`
}

type DataType_Decimal_ struct {
	Decimal *DataType_Decimal `protobuf:"bytes,10,opt,name=decimal,proto3,oneof"`
}

type DataType_String_ struct {
	// String types
	String_ *DataType_String `protobuf:"bytes,11,opt,name=string,proto3,oneof"`
}

type DataType_Char_ struct {
	Char *DataType_Char `protobuf:"bytes,12,opt,name=char,proto3,oneof"`
}

type DataType_VarChar_ struct {
	VarChar *DataType_VarChar `protobuf:"bytes,13,opt,name=var_char,json=varChar,proto3,oneof"`
}

type DataType_Date_ struct {
	// Datatime types
	Date *DataType_Date `protobuf:"bytes,14,opt,name=date,proto3,oneof"`
}

type DataType_Timestamp_ struct {
	Timestamp *DataType_Timestamp `protobuf:"bytes,15,opt,name=timestamp,proto3,oneof"`
}

type DataType_TimestampNtz struct {
	TimestampNtz *DataType_TimestampNTZ `protobuf:"bytes,16,opt,name=timestamp_ntz,json=timestampNtz,proto3,oneof"`
}

type DataType_CalendarInterval_ struct {
	// Interval types
	CalendarInterval *DataType_CalendarInterval `protobuf:"bytes,17,opt,name=calendar_interval,json=calendarInterval,proto3,oneof"`
}

type DataType_YearMonthInterval_ struct {
	YearMonthInterval *DataType_YearMonthInterval `protobuf:"bytes,18,opt,name=year_month_interval,json=yearMonthInterval,proto3,oneof"`
}

type DataType_DayTimeInterval_ struct {
	DayTimeInterval *DataType_DayTimeInterval `protobuf:"bytes,19,opt,name=day_time_interval,json=dayTimeInterval,proto3,oneof"`
}

type DataType_Array_ struct {
	// Complex types
	Array *DataType_Array `protobuf:"bytes,20,opt,name=array,proto3,oneof"`
}

type DataType_Struct_ struct {
	Struct *DataType_Struct `protobuf:"bytes,21,opt,name=struct,proto3,oneof"`
}

type DataType_Map_ struct {
	Map *DataType_Map `protobuf:"bytes,22,opt,name=map,proto3,oneof"`
}

type DataType_Udt struct {
	// UserDefinedType
	Udt *DataType_UDT `protobuf:"bytes,23,opt,name=udt,proto3,oneof"`
}

type DataType_Unparsed_ struct {
	// UnparsedDataType
	Unparsed *DataType_Unparsed `protobuf:"bytes,24,opt,name=unparsed,proto3,oneof"`
}

func (*DataType_Null) isDataType_Kind() {}

func (*DataType_Binary_) isDataType_Kind() {}

func (*DataType_Boolean_) isDataType_Kind() {}

func (*DataType_Byte_) isDataType_Kind() {}

func (*DataType_Short_) isDataType_Kind() {}

func (*DataType_Integer_) isDataType_Kind() {}

func (*DataType_Long_) isDataType_Kind() {}

func (*DataType_Float_) isDataType_Kind() {}

func (*DataType_Double_) isDataType_Kind() {}

func (*DataType_Decimal_) isDataType_Kind() {}

func (*DataType_String_) isDataType_Kind() {}

func (*DataType_Char_) isDataType_Kind() {}

func (*DataType_VarChar_) isDataType_Kind() {}

func (*DataType_Date_) isDataType_Kind() {}

func (*DataType_Timestamp_) isDataType_Kind() {}

func (*DataType_TimestampNtz) isDataType_Kind() {}

func (*DataType_CalendarInterval_) isDataType_Kind() {}

func (*DataType_YearMonthInterval_) isDataType_Kind() {}

func (*DataType_DayTimeInterval_) isDataType_Kind() {}

func (*DataType_Array_) isDataType_Kind() {}

func (*DataType_Struct_) isDataType_Kind() {}

func (*DataType_Map_) isDataType_Kind() {}

func (*DataType_Udt) isDataType_Kind() {}

func (*DataType_Unparsed_) isDataType_Kind() {}

type DataType_Boolean struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeVariationReference uint32 `protobuf:"varint,1,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_Boolean) Reset() {
	*x = DataType_Boolean{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Boolean) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Boolean) ProtoMessage() {}

func (x *DataType_Boolean) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Boolean.ProtoReflect.Descriptor instead.
func (*DataType_Boolean) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DataType_Boolean) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_Byte struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeVariationReference uint32 `protobuf:"varint,1,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_Byte) Reset() {
	*x = DataType_Byte{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Byte) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Byte) ProtoMessage() {}

func (x *DataType_Byte) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Byte.ProtoReflect.Descriptor instead.
func (*DataType_Byte) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 1}
}

func (x *DataType_Byte) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_Short struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeVariationReference uint32 `protobuf:"varint,1,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_Short) Reset() {
	*x = DataType_Short{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Short) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Short) ProtoMessage() {}

func (x *DataType_Short) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Short.ProtoReflect.Descriptor instead.
func (*DataType_Short) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 2}
}

func (x *DataType_Short) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_Integer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeVariationReference uint32 `protobuf:"varint,1,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_Integer) Reset() {
	*x = DataType_Integer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Integer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Integer) ProtoMessage() {}

func (x *DataType_Integer) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Integer.ProtoReflect.Descriptor instead.
func (*DataType_Integer) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 3}
}

func (x *DataType_Integer) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_Long struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeVariationReference uint32 `protobuf:"varint,1,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_Long) Reset() {
	*x = DataType_Long{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Long) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Long) ProtoMessage() {}

func (x *DataType_Long) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Long.ProtoReflect.Descriptor instead.
func (*DataType_Long) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 4}
}

func (x *DataType_Long) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_Float struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeVariationReference uint32 `protobuf:"varint,1,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_Float) Reset() {
	*x = DataType_Float{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Float) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Float) ProtoMessage() {}

func (x *DataType_Float) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Float.ProtoReflect.Descriptor instead.
func (*DataType_Float) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 5}
}

func (x *DataType_Float) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_Double struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeVariationReference uint32 `protobuf:"varint,1,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_Double) Reset() {
	*x = DataType_Double{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Double) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Double) ProtoMessage() {}

func (x *DataType_Double) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Double.ProtoReflect.Descriptor instead.
func (*DataType_Double) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 6}
}

func (x *DataType_Double) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_String struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeVariationReference uint32 `protobuf:"varint,1,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_String) Reset() {
	*x = DataType_String{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_String) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_String) ProtoMessage() {}

func (x *DataType_String) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_String.ProtoReflect.Descriptor instead.
func (*DataType_String) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 7}
}

func (x *DataType_String) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_Binary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeVariationReference uint32 `protobuf:"varint,1,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_Binary) Reset() {
	*x = DataType_Binary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Binary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Binary) ProtoMessage() {}

func (x *DataType_Binary) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Binary.ProtoReflect.Descriptor instead.
func (*DataType_Binary) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 8}
}

func (x *DataType_Binary) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_NULL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeVariationReference uint32 `protobuf:"varint,1,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_NULL) Reset() {
	*x = DataType_NULL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_NULL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_NULL) ProtoMessage() {}

func (x *DataType_NULL) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_NULL.ProtoReflect.Descriptor instead.
func (*DataType_NULL) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 9}
}

func (x *DataType_NULL) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_Timestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeVariationReference uint32 `protobuf:"varint,1,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_Timestamp) Reset() {
	*x = DataType_Timestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Timestamp) ProtoMessage() {}

func (x *DataType_Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Timestamp.ProtoReflect.Descriptor instead.
func (*DataType_Timestamp) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 10}
}

func (x *DataType_Timestamp) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeVariationReference uint32 `protobuf:"varint,1,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_Date) Reset() {
	*x = DataType_Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Date) ProtoMessage() {}

func (x *DataType_Date) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Date.ProtoReflect.Descriptor instead.
func (*DataType_Date) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 11}
}

func (x *DataType_Date) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_TimestampNTZ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeVariationReference uint32 `protobuf:"varint,1,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_TimestampNTZ) Reset() {
	*x = DataType_TimestampNTZ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_TimestampNTZ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_TimestampNTZ) ProtoMessage() {}

func (x *DataType_TimestampNTZ) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_TimestampNTZ.ProtoReflect.Descriptor instead.
func (*DataType_TimestampNTZ) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 12}
}

func (x *DataType_TimestampNTZ) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_CalendarInterval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeVariationReference uint32 `protobuf:"varint,1,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_CalendarInterval) Reset() {
	*x = DataType_CalendarInterval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_CalendarInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_CalendarInterval) ProtoMessage() {}

func (x *DataType_CalendarInterval) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_CalendarInterval.ProtoReflect.Descriptor instead.
func (*DataType_CalendarInterval) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 13}
}

func (x *DataType_CalendarInterval) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_YearMonthInterval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartField             *int32 `protobuf:"varint,1,opt,name=start_field,json=startField,proto3,oneof" json:"start_field,omitempty"`
	EndField               *int32 `protobuf:"varint,2,opt,name=end_field,json=endField,proto3,oneof" json:"end_field,omitempty"`
	TypeVariationReference uint32 `protobuf:"varint,3,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_YearMonthInterval) Reset() {
	*x = DataType_YearMonthInterval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_YearMonthInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_YearMonthInterval) ProtoMessage() {}

func (x *DataType_YearMonthInterval) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_YearMonthInterval.ProtoReflect.Descriptor instead.
func (*DataType_YearMonthInterval) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 14}
}

func (x *DataType_YearMonthInterval) GetStartField() int32 {
	if x != nil && x.StartField != nil {
		return *x.StartField
	}
	return 0
}

func (x *DataType_YearMonthInterval) GetEndField() int32 {
	if x != nil && x.EndField != nil {
		return *x.EndField
	}
	return 0
}

func (x *DataType_YearMonthInterval) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_DayTimeInterval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartField             *int32 `protobuf:"varint,1,opt,name=start_field,json=startField,proto3,oneof" json:"start_field,omitempty"`
	EndField               *int32 `protobuf:"varint,2,opt,name=end_field,json=endField,proto3,oneof" json:"end_field,omitempty"`
	TypeVariationReference uint32 `protobuf:"varint,3,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_DayTimeInterval) Reset() {
	*x = DataType_DayTimeInterval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_DayTimeInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_DayTimeInterval) ProtoMessage() {}

func (x *DataType_DayTimeInterval) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_DayTimeInterval.ProtoReflect.Descriptor instead.
func (*DataType_DayTimeInterval) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 15}
}

func (x *DataType_DayTimeInterval) GetStartField() int32 {
	if x != nil && x.StartField != nil {
		return *x.StartField
	}
	return 0
}

func (x *DataType_DayTimeInterval) GetEndField() int32 {
	if x != nil && x.EndField != nil {
		return *x.EndField
	}
	return 0
}

func (x *DataType_DayTimeInterval) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

// Start compound types.
type DataType_Char struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length                 int32  `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	TypeVariationReference uint32 `protobuf:"varint,2,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_Char) Reset() {
	*x = DataType_Char{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Char) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Char) ProtoMessage() {}

func (x *DataType_Char) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Char.ProtoReflect.Descriptor instead.
func (*DataType_Char) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 16}
}

func (x *DataType_Char) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *DataType_Char) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_VarChar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length                 int32  `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	TypeVariationReference uint32 `protobuf:"varint,2,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_VarChar) Reset() {
	*x = DataType_VarChar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_VarChar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_VarChar) ProtoMessage() {}

func (x *DataType_VarChar) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_VarChar.ProtoReflect.Descriptor instead.
func (*DataType_VarChar) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 17}
}

func (x *DataType_VarChar) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *DataType_VarChar) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_Decimal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scale                  *int32 `protobuf:"varint,1,opt,name=scale,proto3,oneof" json:"scale,omitempty"`
	Precision              *int32 `protobuf:"varint,2,opt,name=precision,proto3,oneof" json:"precision,omitempty"`
	TypeVariationReference uint32 `protobuf:"varint,3,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_Decimal) Reset() {
	*x = DataType_Decimal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Decimal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Decimal) ProtoMessage() {}

func (x *DataType_Decimal) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Decimal.ProtoReflect.Descriptor instead.
func (*DataType_Decimal) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 18}
}

func (x *DataType_Decimal) GetScale() int32 {
	if x != nil && x.Scale != nil {
		return *x.Scale
	}
	return 0
}

func (x *DataType_Decimal) GetPrecision() int32 {
	if x != nil && x.Precision != nil {
		return *x.Precision
	}
	return 0
}

func (x *DataType_Decimal) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_StructField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DataType *DataType `protobuf:"bytes,2,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	Nullable bool      `protobuf:"varint,3,opt,name=nullable,proto3" json:"nullable,omitempty"`
	Metadata *string   `protobuf:"bytes,4,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
}

func (x *DataType_StructField) Reset() {
	*x = DataType_StructField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[20]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_StructField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_StructField) ProtoMessage() {}

func (x *DataType_StructField) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[20]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_StructField.ProtoReflect.Descriptor instead.
func (*DataType_StructField) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 19}
}

func (x *DataType_StructField) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataType_StructField) GetDataType() *DataType {
	if x != nil {
		return x.DataType
	}
	return nil
}

func (x *DataType_StructField) GetNullable() bool {
	if x != nil {
		return x.Nullable
	}
	return false
}

func (x *DataType_StructField) GetMetadata() string {
	if x != nil && x.Metadata != nil {
		return *x.Metadata
	}
	return ""
}

type DataType_Struct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields                 []*DataType_StructField `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	TypeVariationReference uint32                  `protobuf:"varint,2,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_Struct) Reset() {
	*x = DataType_Struct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[21]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Struct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Struct) ProtoMessage() {}

func (x *DataType_Struct) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[21]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Struct.ProtoReflect.Descriptor instead.
func (*DataType_Struct) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 20}
}

func (x *DataType_Struct) GetFields() []*DataType_StructField {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *DataType_Struct) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_Array struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ElementType            *DataType `protobuf:"bytes,1,opt,name=element_type,json=elementType,proto3" json:"element_type,omitempty"`
	ContainsNull           bool      `protobuf:"varint,2,opt,name=contains_null,json=containsNull,proto3" json:"contains_null,omitempty"`
	TypeVariationReference uint32    `protobuf:"varint,3,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_Array) Reset() {
	*x = DataType_Array{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[22]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Array) ProtoMessage() {}

func (x *DataType_Array) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[22]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Array.ProtoReflect.Descriptor instead.
func (*DataType_Array) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 21}
}

func (x *DataType_Array) GetElementType() *DataType {
	if x != nil {
		return x.ElementType
	}
	return nil
}

func (x *DataType_Array) GetContainsNull() bool {
	if x != nil {
		return x.ContainsNull
	}
	return false
}

func (x *DataType_Array) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_Map struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyType                *DataType `protobuf:"bytes,1,opt,name=key_type,json=keyType,proto3" json:"key_type,omitempty"`
	ValueType              *DataType `protobuf:"bytes,2,opt,name=value_type,json=valueType,proto3" json:"value_type,omitempty"`
	ValueContainsNull      bool      `protobuf:"varint,3,opt,name=value_contains_null,json=valueContainsNull,proto3" json:"value_contains_null,omitempty"`
	TypeVariationReference uint32    `protobuf:"varint,4,opt,name=type_variation_reference,json=typeVariationReference,proto3" json:"type_variation_reference,omitempty"`
}

func (x *DataType_Map) Reset() {
	*x = DataType_Map{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[23]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Map) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Map) ProtoMessage() {}

func (x *DataType_Map) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[23]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Map.ProtoReflect.Descriptor instead.
func (*DataType_Map) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 22}
}

func (x *DataType_Map) GetKeyType() *DataType {
	if x != nil {
		return x.KeyType
	}
	return nil
}

func (x *DataType_Map) GetValueType() *DataType {
	if x != nil {
		return x.ValueType
	}
	return nil
}

func (x *DataType_Map) GetValueContainsNull() bool {
	if x != nil {
		return x.ValueContainsNull
	}
	return false
}

func (x *DataType_Map) GetTypeVariationReference() uint32 {
	if x != nil {
		return x.TypeVariationReference
	}
	return 0
}

type DataType_UDT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                  string    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	JvmClass              *string   `protobuf:"bytes,2,opt,name=jvm_class,json=jvmClass,proto3,oneof" json:"jvm_class,omitempty"`
	PythonClass           *string   `protobuf:"bytes,3,opt,name=python_class,json=pythonClass,proto3,oneof" json:"python_class,omitempty"`
	SerializedPythonClass *string   `protobuf:"bytes,4,opt,name=serialized_python_class,json=serializedPythonClass,proto3,oneof" json:"serialized_python_class,omitempty"`
	SqlType               *DataType `protobuf:"bytes,5,opt,name=sql_type,json=sqlType,proto3" json:"sql_type,omitempty"`
}

func (x *DataType_UDT) Reset() {
	*x = DataType_UDT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[24]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_UDT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_UDT) ProtoMessage() {}

func (x *DataType_UDT) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[24]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_UDT.ProtoReflect.Descriptor instead.
func (*DataType_UDT) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 23}
}

func (x *DataType_UDT) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DataType_UDT) GetJvmClass() string {
	if x != nil && x.JvmClass != nil {
		return *x.JvmClass
	}
	return ""
}

func (x *DataType_UDT) GetPythonClass() string {
	if x != nil && x.PythonClass != nil {
		return *x.PythonClass
	}
	return ""
}

func (x *DataType_UDT) GetSerializedPythonClass() string {
	if x != nil && x.SerializedPythonClass != nil {
		return *x.SerializedPythonClass
	}
	return ""
}

func (x *DataType_UDT) GetSqlType() *DataType {
	if x != nil {
		return x.SqlType
	}
	return nil
}

type DataType_Unparsed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The unparsed data type string
	DataTypeString string `protobuf:"bytes,1,opt,name=data_type_string,json=dataTypeString,proto3" json:"data_type_string,omitempty"`
}

func (x *DataType_Unparsed) Reset() {
	*x = DataType_Unparsed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_types_proto_msgTypes[25]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataType_Unparsed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataType_Unparsed) ProtoMessage() {}

func (x *DataType_Unparsed) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_types_proto_msgTypes[25]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataType_Unparsed.ProtoReflect.Descriptor instead.
func (*DataType_Unparsed) Descriptor() ([]byte, []int) {
	return file_spark_connect_types_proto_rawDescGZIP(), []int{0, 24}
}

func (x *DataType_Unparsed) GetDataTypeString() string {
	if x != nil {
		return x.DataTypeString
	}
	return ""
}

var File_spark_connect_types_proto protoreflect.FileDescriptor

var file_spark_connect_types_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x22, 0xc7, 0x20, 0x0a, 0x08, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x6e, 0x75, 0x6c, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x4e,
	0x55, 0x4c, 0x4c, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x75, 0x6c, 0x6c, 0x12, 0x38, 0x0a, 0x06, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x48, 0x00, 0x52, 0x06, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x48, 0x00, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x65,
	0x61, 0x6e, 0x12, 0x32, 0x0a, 0x04, 0x62, 0x79, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x48, 0x00,
	0x52, 0x04, 0x62, 0x79, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x48, 0x00, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x3b, 0x0a,
	0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x48,
	0x00, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x04, 0x6c, 0x6f,
	0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x35,
	0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x48, 0x00, 0x52, 0x05,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12,
	0x3b, 0x0a, 0x07, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61,
	0x6c, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x38, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x32, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x72, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x68,
	0x61, 0x72, 0x48, 0x00, 0x52, 0x04, 0x63, 0x68, 0x61, 0x72, 0x12, 0x3c, 0x0a, 0x08, 0x76, 0x61,
	0x72, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x56, 0x61, 0x72, 0x43, 0x68, 0x61, 0x72, 0x48, 0x00, 0x52,
	0x07, 0x76, 0x61, 0x72, 0x43, 0x68, 0x61, 0x72, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x41, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x4b, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6e, 0x74, 0x7a,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4e, 0x54, 0x5a, 0x48, 0x00, 0x52, 0x0c,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4e, 0x74, 0x7a, 0x12, 0x57, 0x0a, 0x11,
	0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x48, 0x00, 0x52, 0x10, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x5b, 0x0a, 0x13, 0x79, 0x65, 0x61, 0x72, 0x5f, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x59, 0x65, 0x61, 0x72,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x48, 0x00, 0x52,
	0x11, 0x79, 0x65, 0x61, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x12, 0x55, 0x0a, 0x11, 0x64, 0x61, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x0f, 0x64, 0x61, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x35, 0x0a, 0x05, 0x61, 0x72, 0x72,
	0x61, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x00, 0x52, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79,
	0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x2f, 0x0a, 0x03, 0x6d, 0x61,
	0x70, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x4d, 0x61, 0x70, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x12, 0x2f, 0x0a, 0x03, 0x75,
	0x64, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x55, 0x44, 0x54, 0x48, 0x00, 0x52, 0x03, 0x75, 0x64, 0x74, 0x12, 0x3e, 0x0a, 0x08,
	0x75, 0x6e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x55, 0x6e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64,
	0x48, 0x00, 0x52, 0x08, 0x75, 0x6e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x1a, 0x43, 0x0a, 0x07,
	0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79, 0x70, 0x65, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x1a, 0x40, 0x0a, 0x04, 0x42, 0x79, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79, 0x70,
	0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x1a, 0x41, 0x0a, 0x05, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x38, 0x0a, 0x18,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16,
	0x74, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x43, 0x0a, 0x07, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65,
	0x72, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x40, 0x0a, 0x04, 0x4c,
	0x6f, 0x6e, 0x67, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x41, 0x0a,
	0x05, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79, 0x70, 0x65, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x1a, 0x42, 0x0a, 0x06, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79,
	0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x1a, 0x42, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x38,
	0x0a, 0x18, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x16, 0x74, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x42, 0x0a, 0x06, 0x42, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x40, 0x0a, 0x04,
	0x4e, 0x55, 0x4c, 0x4c, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x45,
	0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x38, 0x0a, 0x18, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x74,
	0x79, 0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x40, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a,
	0x18, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x16, 0x74, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x48, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x4e, 0x54, 0x5a, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79, 0x70, 0x65, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x1a, 0x4c, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a,
	0xb3, 0x01, 0x0a, 0x11, 0x59, 0x65, 0x61, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x65,
	0x6e, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01,
	0x52, 0x08, 0x65, 0x6e, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a,
	0x18, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x16, 0x74, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x65, 0x6e, 0x64, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0xb1, 0x01, 0x0a, 0x0f, 0x44, 0x61, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x20, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x01, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x65, 0x6e, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0x58, 0x0a, 0x04, 0x43, 0x68, 0x61,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79, 0x70,
	0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x1a, 0x5b, 0x0a, 0x07, 0x56, 0x61, 0x72, 0x43, 0x68, 0x61, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79, 0x70, 0x65, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x1a, 0x99, 0x01, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x05,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x09, 0x70, 0x72,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79,
	0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0xa1, 0x01, 0x0a,
	0x0b, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x34, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x7f, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79, 0x70, 0x65, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x1a, 0xa2, 0x01, 0x0a, 0x05, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x3a, 0x0a, 0x0c, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x73, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x4e, 0x75, 0x6c, 0x6c, 0x12, 0x38, 0x0a, 0x18,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16,
	0x74, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0xdb, 0x01, 0x0a, 0x03, 0x4d, 0x61, 0x70, 0x12, 0x32,
	0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x5f, 0x6e, 0x75, 0x6c,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x4e, 0x75, 0x6c, 0x6c, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x74, 0x79,
	0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x1a, 0x8f, 0x02, 0x0a, 0x03, 0x55, 0x44, 0x54, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x20, 0x0a, 0x09, 0x6a, 0x76, 0x6d, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6a, 0x76, 0x6d, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x70, 0x79, 0x74, 0x68,
	0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x17, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x15, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x50, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x71, 0x6c, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x07, 0x73, 0x71, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x6a, 0x76, 0x6d, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x70, 0x79,
	0x74, 0x68, 0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e,
	0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x1a, 0x34, 0x0a, 0x08, 0x55, 0x6e, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x06, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x42, 0x36, 0x0a, 0x1e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spark_connect_types_proto_rawDescOnce sync.Once
	file_spark_connect_types_proto_rawDescData = file_spark_connect_types_proto_rawDesc
)

func file_spark_connect_types_proto_rawDescGZIP() []byte {
	file_spark_connect_types_proto_rawDescOnce.Do(func() {
		file_spark_connect_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_spark_connect_types_proto_rawDescData)
	})
	return file_spark_connect_types_proto_rawDescData
}

var file_spark_connect_types_proto_msgTypes = make([]protoimpl.MessageInfo, 26)
var file_spark_connect_types_proto_goTypes = []interface{}{
	(*DataType)(nil),                   // 0: spark.connect.DataType
	(*DataType_Boolean)(nil),           // 1: spark.connect.DataType.Boolean
	(*DataType_Byte)(nil),              // 2: spark.connect.DataType.Byte
	(*DataType_Short)(nil),             // 3: spark.connect.DataType.Short
	(*DataType_Integer)(nil),           // 4: spark.connect.DataType.Integer
	(*DataType_Long)(nil),              // 5: spark.connect.DataType.Long
	(*DataType_Float)(nil),             // 6: spark.connect.DataType.Float
	(*DataType_Double)(nil),            // 7: spark.connect.DataType.Double
	(*DataType_String)(nil),            // 8: spark.connect.DataType.String
	(*DataType_Binary)(nil),            // 9: spark.connect.DataType.Binary
	(*DataType_NULL)(nil),              // 10: spark.connect.DataType.NULL
	(*DataType_Timestamp)(nil),         // 11: spark.connect.DataType.Timestamp
	(*DataType_Date)(nil),              // 12: spark.connect.DataType.Date
	(*DataType_TimestampNTZ)(nil),      // 13: spark.connect.DataType.TimestampNTZ
	(*DataType_CalendarInterval)(nil),  // 14: spark.connect.DataType.CalendarInterval
	(*DataType_YearMonthInterval)(nil), // 15: spark.connect.DataType.YearMonthInterval
	(*DataType_DayTimeInterval)(nil),   // 16: spark.connect.DataType.DayTimeInterval
	(*DataType_Char)(nil),              // 17: spark.connect.DataType.Char
	(*DataType_VarChar)(nil),           // 18: spark.connect.DataType.VarChar
	(*DataType_Decimal)(nil),           // 19: spark.connect.DataType.Decimal
	(*DataType_StructField)(nil),       // 20: spark.connect.DataType.StructField
	(*DataType_Struct)(nil),            // 21: spark.connect.DataType.Struct
	(*DataType_Array)(nil),             // 22: spark.connect.DataType.Array
	(*DataType_Map)(nil),               // 23: spark.connect.DataType.Map
	(*DataType_UDT)(nil),               // 24: spark.connect.DataType.UDT
	(*DataType_Unparsed)(nil),          // 25: spark.connect.DataType.Unparsed
}
var file_spark_connect_types_proto_depIdxs = []int32{
	10, // 0: spark.connect.DataType.null:type_name -> spark.connect.DataType.NULL
	9,  // 1: spark.connect.DataType.binary:type_name -> spark.connect.DataType.Binary
	1,  // 2: spark.connect.DataType.boolean:type_name -> spark.connect.DataType.Boolean
	2,  // 3: spark.connect.DataType.byte:type_name -> spark.connect.DataType.Byte
	3,  // 4: spark.connect.DataType.short:type_name -> spark.connect.DataType.Short
	4,  // 5: spark.connect.DataType.integer:type_name -> spark.connect.DataType.Integer
	5,  // 6: spark.connect.DataType.long:type_name -> spark.connect.DataType.Long
	6,  // 7: spark.connect.DataType.float:type_name -> spark.connect.DataType.Float
	7,  // 8: spark.connect.DataType.double:type_name -> spark.connect.DataType.Double
	19, // 9: spark.connect.DataType.decimal:type_name -> spark.connect.DataType.Decimal
	8,  // 10: spark.connect.DataType.string:type_name -> spark.connect.DataType.String
	17, // 11: spark.connect.DataType.char:type_name -> spark.connect.DataType.Char
	18, // 12: spark.connect.DataType.var_char:type_name -> spark.connect.DataType.VarChar
	12, // 13: spark.connect.DataType.date:type_name -> spark.connect.DataType.Date
	11, // 14: spark.connect.DataType.timestamp:type_name -> spark.connect.DataType.Timestamp
	13, // 15: spark.connect.DataType.timestamp_ntz:type_name -> spark.connect.DataType.TimestampNTZ
	14, // 16: spark.connect.DataType.calendar_interval:type_name -> spark.connect.DataType.CalendarInterval
	15, // 17: spark.connect.DataType.year_month_interval:type_name -> spark.connect.DataType.YearMonthInterval
	16, // 18: spark.connect.DataType.day_time_interval:type_name -> spark.connect.DataType.DayTimeInterval
	22, // 19: spark.connect.DataType.array:type_name -> spark.connect.DataType.Array
	21, // 20: spark.connect.DataType.struct:type_name -> spark.connect.DataType.Struct
	23, // 21: spark.connect.DataType.map:type_name -> spark.connect.DataType.Map
	24, // 22: spark.connect.DataType.udt:type_name -> spark.connect.DataType.UDT
	25, // 23: spark.connect.DataType.unparsed:type_name -> spark.connect.DataType.Unparsed
	0,  // 24: spark.connect.DataType.StructField.data_type:type_name -> spark.connect.DataType
	20, // 25: spark.connect.DataType.Struct.fields:type_name -> spark.connect.DataType.StructField
	0,  // 26: spark.connect.DataType.Array.element_type:type_name -> spark.connect.DataType
	0,  // 27: spark.connect.DataType.Map.key_type:type_name -> spark.connect.DataType
	0,  // 28: spark.connect.DataType.Map.value_type:type_name -> spark.connect.DataType
	0,  // 29: spark.connect.DataType.UDT.sql_type:type_name -> spark.connect.DataType
	30, // [30:30] is the sub-list for method output_type
	30, // [30:30] is the sub-list for method input_type
	30, // [30:30] is the sub-list for extension type_name
	30, // [30:30] is the sub-list for extension extendee
	0,  // [0:30] is the sub-list for field type_name
}

func init() { file_spark_connect_types_proto_init() }
func file_spark_connect_types_proto_init() {
	if File_spark_connect_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spark_connect_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType); i {
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
		file_spark_connect_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Boolean); i {
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
		file_spark_connect_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Byte); i {
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
		file_spark_connect_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Short); i {
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
		file_spark_connect_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Integer); i {
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
		file_spark_connect_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Long); i {
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
		file_spark_connect_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Float); i {
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
		file_spark_connect_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Double); i {
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
		file_spark_connect_types_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_String); i {
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
		file_spark_connect_types_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Binary); i {
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
		file_spark_connect_types_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_NULL); i {
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
		file_spark_connect_types_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Timestamp); i {
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
		file_spark_connect_types_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Date); i {
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
		file_spark_connect_types_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_TimestampNTZ); i {
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
		file_spark_connect_types_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_CalendarInterval); i {
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
		file_spark_connect_types_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_YearMonthInterval); i {
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
		file_spark_connect_types_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_DayTimeInterval); i {
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
		file_spark_connect_types_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Char); i {
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
		file_spark_connect_types_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_VarChar); i {
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
		file_spark_connect_types_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Decimal); i {
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
		file_spark_connect_types_proto_msgTypes[20].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_StructField); i {
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
		file_spark_connect_types_proto_msgTypes[21].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Struct); i {
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
		file_spark_connect_types_proto_msgTypes[22].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Array); i {
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
		file_spark_connect_types_proto_msgTypes[23].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Map); i {
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
		file_spark_connect_types_proto_msgTypes[24].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_UDT); i {
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
		file_spark_connect_types_proto_msgTypes[25].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataType_Unparsed); i {
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
	file_spark_connect_types_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DataType_Null)(nil),
		(*DataType_Binary_)(nil),
		(*DataType_Boolean_)(nil),
		(*DataType_Byte_)(nil),
		(*DataType_Short_)(nil),
		(*DataType_Integer_)(nil),
		(*DataType_Long_)(nil),
		(*DataType_Float_)(nil),
		(*DataType_Double_)(nil),
		(*DataType_Decimal_)(nil),
		(*DataType_String_)(nil),
		(*DataType_Char_)(nil),
		(*DataType_VarChar_)(nil),
		(*DataType_Date_)(nil),
		(*DataType_Timestamp_)(nil),
		(*DataType_TimestampNtz)(nil),
		(*DataType_CalendarInterval_)(nil),
		(*DataType_YearMonthInterval_)(nil),
		(*DataType_DayTimeInterval_)(nil),
		(*DataType_Array_)(nil),
		(*DataType_Struct_)(nil),
		(*DataType_Map_)(nil),
		(*DataType_Udt)(nil),
		(*DataType_Unparsed_)(nil),
	}
	file_spark_connect_types_proto_msgTypes[15].OneofWrappers = []interface{}{}
	file_spark_connect_types_proto_msgTypes[16].OneofWrappers = []interface{}{}
	file_spark_connect_types_proto_msgTypes[19].OneofWrappers = []interface{}{}
	file_spark_connect_types_proto_msgTypes[20].OneofWrappers = []interface{}{}
	file_spark_connect_types_proto_msgTypes[24].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spark_connect_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   26,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spark_connect_types_proto_goTypes,
		DependencyIndexes: file_spark_connect_types_proto_depIdxs,
		MessageInfos:      file_spark_connect_types_proto_msgTypes,
	}.Build()
	File_spark_connect_types_proto = out.File
	file_spark_connect_types_proto_rawDesc = nil
	file_spark_connect_types_proto_goTypes = nil
	file_spark_connect_types_proto_depIdxs = nil
}
