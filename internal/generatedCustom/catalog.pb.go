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

// Catalog messages are marked as unstable.
type Catalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to CatType:
	//
	//	*Catalog_CurrentDatabase
	//	*Catalog_SetCurrentDatabase
	//	*Catalog_ListDatabases
	//	*Catalog_ListTables
	//	*Catalog_ListFunctions
	//	*Catalog_ListColumns
	//	*Catalog_GetDatabase
	//	*Catalog_GetTable
	//	*Catalog_GetFunction
	//	*Catalog_DatabaseExists
	//	*Catalog_TableExists
	//	*Catalog_FunctionExists
	//	*Catalog_CreateExternalTable
	//	*Catalog_CreateTable
	//	*Catalog_DropTempView
	//	*Catalog_DropGlobalTempView
	//	*Catalog_RecoverPartitions
	//	*Catalog_IsCached
	//	*Catalog_CacheTable
	//	*Catalog_UncacheTable
	//	*Catalog_ClearCache
	//	*Catalog_RefreshTable
	//	*Catalog_RefreshByPath
	//	*Catalog_CurrentCatalog
	//	*Catalog_SetCurrentCatalog
	//	*Catalog_ListCatalogs
	CatType isCatalog_CatType `protobuf_oneof:"cat_type"`
}

func (x *Catalog) Reset() {
	*x = Catalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Catalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Catalog) ProtoMessage() {}

func (x *Catalog) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Catalog.ProtoReflect.Descriptor instead.
func (*Catalog) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{0}
}

func (m *Catalog) GetCatType() isCatalog_CatType {
	if m != nil {
		return m.CatType
	}
	return nil
}

func (x *Catalog) GetCurrentDatabase() *CurrentDatabase {
	if x, ok := x.GetCatType().(*Catalog_CurrentDatabase); ok {
		return x.CurrentDatabase
	}
	return nil
}

func (x *Catalog) GetSetCurrentDatabase() *SetCurrentDatabase {
	if x, ok := x.GetCatType().(*Catalog_SetCurrentDatabase); ok {
		return x.SetCurrentDatabase
	}
	return nil
}

func (x *Catalog) GetListDatabases() *ListDatabases {
	if x, ok := x.GetCatType().(*Catalog_ListDatabases); ok {
		return x.ListDatabases
	}
	return nil
}

func (x *Catalog) GetListTables() *ListTables {
	if x, ok := x.GetCatType().(*Catalog_ListTables); ok {
		return x.ListTables
	}
	return nil
}

func (x *Catalog) GetListFunctions() *ListFunctions {
	if x, ok := x.GetCatType().(*Catalog_ListFunctions); ok {
		return x.ListFunctions
	}
	return nil
}

func (x *Catalog) GetListColumns() *ListColumns {
	if x, ok := x.GetCatType().(*Catalog_ListColumns); ok {
		return x.ListColumns
	}
	return nil
}

func (x *Catalog) GetGetDatabase() *GetDatabase {
	if x, ok := x.GetCatType().(*Catalog_GetDatabase); ok {
		return x.GetDatabase
	}
	return nil
}

func (x *Catalog) GetGetTable() *GetTable {
	if x, ok := x.GetCatType().(*Catalog_GetTable); ok {
		return x.GetTable
	}
	return nil
}

func (x *Catalog) GetGetFunction() *GetFunction {
	if x, ok := x.GetCatType().(*Catalog_GetFunction); ok {
		return x.GetFunction
	}
	return nil
}

func (x *Catalog) GetDatabaseExists() *DatabaseExists {
	if x, ok := x.GetCatType().(*Catalog_DatabaseExists); ok {
		return x.DatabaseExists
	}
	return nil
}

func (x *Catalog) GetTableExists() *TableExists {
	if x, ok := x.GetCatType().(*Catalog_TableExists); ok {
		return x.TableExists
	}
	return nil
}

func (x *Catalog) GetFunctionExists() *FunctionExists {
	if x, ok := x.GetCatType().(*Catalog_FunctionExists); ok {
		return x.FunctionExists
	}
	return nil
}

func (x *Catalog) GetCreateExternalTable() *CreateExternalTable {
	if x, ok := x.GetCatType().(*Catalog_CreateExternalTable); ok {
		return x.CreateExternalTable
	}
	return nil
}

func (x *Catalog) GetCreateTable() *CreateTable {
	if x, ok := x.GetCatType().(*Catalog_CreateTable); ok {
		return x.CreateTable
	}
	return nil
}

func (x *Catalog) GetDropTempView() *DropTempView {
	if x, ok := x.GetCatType().(*Catalog_DropTempView); ok {
		return x.DropTempView
	}
	return nil
}

func (x *Catalog) GetDropGlobalTempView() *DropGlobalTempView {
	if x, ok := x.GetCatType().(*Catalog_DropGlobalTempView); ok {
		return x.DropGlobalTempView
	}
	return nil
}

func (x *Catalog) GetRecoverPartitions() *RecoverPartitions {
	if x, ok := x.GetCatType().(*Catalog_RecoverPartitions); ok {
		return x.RecoverPartitions
	}
	return nil
}

func (x *Catalog) GetIsCached() *IsCached {
	if x, ok := x.GetCatType().(*Catalog_IsCached); ok {
		return x.IsCached
	}
	return nil
}

func (x *Catalog) GetCacheTable() *CacheTable {
	if x, ok := x.GetCatType().(*Catalog_CacheTable); ok {
		return x.CacheTable
	}
	return nil
}

func (x *Catalog) GetUncacheTable() *UncacheTable {
	if x, ok := x.GetCatType().(*Catalog_UncacheTable); ok {
		return x.UncacheTable
	}
	return nil
}

func (x *Catalog) GetClearCache() *ClearCache {
	if x, ok := x.GetCatType().(*Catalog_ClearCache); ok {
		return x.ClearCache
	}
	return nil
}

func (x *Catalog) GetRefreshTable() *RefreshTable {
	if x, ok := x.GetCatType().(*Catalog_RefreshTable); ok {
		return x.RefreshTable
	}
	return nil
}

func (x *Catalog) GetRefreshByPath() *RefreshByPath {
	if x, ok := x.GetCatType().(*Catalog_RefreshByPath); ok {
		return x.RefreshByPath
	}
	return nil
}

func (x *Catalog) GetCurrentCatalog() *CurrentCatalog {
	if x, ok := x.GetCatType().(*Catalog_CurrentCatalog); ok {
		return x.CurrentCatalog
	}
	return nil
}

func (x *Catalog) GetSetCurrentCatalog() *SetCurrentCatalog {
	if x, ok := x.GetCatType().(*Catalog_SetCurrentCatalog); ok {
		return x.SetCurrentCatalog
	}
	return nil
}

func (x *Catalog) GetListCatalogs() *ListCatalogs {
	if x, ok := x.GetCatType().(*Catalog_ListCatalogs); ok {
		return x.ListCatalogs
	}
	return nil
}

type isCatalog_CatType interface {
	isCatalog_CatType()
}

type Catalog_CurrentDatabase struct {
	CurrentDatabase *CurrentDatabase `protobuf:"bytes,1,opt,name=current_database,json=currentDatabase,proto3,oneof"`
}

type Catalog_SetCurrentDatabase struct {
	SetCurrentDatabase *SetCurrentDatabase `protobuf:"bytes,2,opt,name=set_current_database,json=setCurrentDatabase,proto3,oneof"`
}

type Catalog_ListDatabases struct {
	ListDatabases *ListDatabases `protobuf:"bytes,3,opt,name=list_databases,json=listDatabases,proto3,oneof"`
}

type Catalog_ListTables struct {
	ListTables *ListTables `protobuf:"bytes,4,opt,name=list_tables,json=listTables,proto3,oneof"`
}

type Catalog_ListFunctions struct {
	ListFunctions *ListFunctions `protobuf:"bytes,5,opt,name=list_functions,json=listFunctions,proto3,oneof"`
}

type Catalog_ListColumns struct {
	ListColumns *ListColumns `protobuf:"bytes,6,opt,name=list_columns,json=listColumns,proto3,oneof"`
}

type Catalog_GetDatabase struct {
	GetDatabase *GetDatabase `protobuf:"bytes,7,opt,name=get_database,json=getDatabase,proto3,oneof"`
}

type Catalog_GetTable struct {
	GetTable *GetTable `protobuf:"bytes,8,opt,name=get_table,json=getTable,proto3,oneof"`
}

type Catalog_GetFunction struct {
	GetFunction *GetFunction `protobuf:"bytes,9,opt,name=get_function,json=getFunction,proto3,oneof"`
}

type Catalog_DatabaseExists struct {
	DatabaseExists *DatabaseExists `protobuf:"bytes,10,opt,name=database_exists,json=databaseExists,proto3,oneof"`
}

type Catalog_TableExists struct {
	TableExists *TableExists `protobuf:"bytes,11,opt,name=table_exists,json=tableExists,proto3,oneof"`
}

type Catalog_FunctionExists struct {
	FunctionExists *FunctionExists `protobuf:"bytes,12,opt,name=function_exists,json=functionExists,proto3,oneof"`
}

type Catalog_CreateExternalTable struct {
	CreateExternalTable *CreateExternalTable `protobuf:"bytes,13,opt,name=create_external_table,json=createExternalTable,proto3,oneof"`
}

type Catalog_CreateTable struct {
	CreateTable *CreateTable `protobuf:"bytes,14,opt,name=create_table,json=createTable,proto3,oneof"`
}

type Catalog_DropTempView struct {
	DropTempView *DropTempView `protobuf:"bytes,15,opt,name=drop_temp_view,json=dropTempView,proto3,oneof"`
}

type Catalog_DropGlobalTempView struct {
	DropGlobalTempView *DropGlobalTempView `protobuf:"bytes,16,opt,name=drop_global_temp_view,json=dropGlobalTempView,proto3,oneof"`
}

type Catalog_RecoverPartitions struct {
	RecoverPartitions *RecoverPartitions `protobuf:"bytes,17,opt,name=recover_partitions,json=recoverPartitions,proto3,oneof"`
}

type Catalog_IsCached struct {
	IsCached *IsCached `protobuf:"bytes,18,opt,name=is_cached,json=isCached,proto3,oneof"`
}

type Catalog_CacheTable struct {
	CacheTable *CacheTable `protobuf:"bytes,19,opt,name=cache_table,json=cacheTable,proto3,oneof"`
}

type Catalog_UncacheTable struct {
	UncacheTable *UncacheTable `protobuf:"bytes,20,opt,name=uncache_table,json=uncacheTable,proto3,oneof"`
}

type Catalog_ClearCache struct {
	ClearCache *ClearCache `protobuf:"bytes,21,opt,name=clear_cache,json=clearCache,proto3,oneof"`
}

type Catalog_RefreshTable struct {
	RefreshTable *RefreshTable `protobuf:"bytes,22,opt,name=refresh_table,json=refreshTable,proto3,oneof"`
}

type Catalog_RefreshByPath struct {
	RefreshByPath *RefreshByPath `protobuf:"bytes,23,opt,name=refresh_by_path,json=refreshByPath,proto3,oneof"`
}

type Catalog_CurrentCatalog struct {
	CurrentCatalog *CurrentCatalog `protobuf:"bytes,24,opt,name=current_catalog,json=currentCatalog,proto3,oneof"`
}

type Catalog_SetCurrentCatalog struct {
	SetCurrentCatalog *SetCurrentCatalog `protobuf:"bytes,25,opt,name=set_current_catalog,json=setCurrentCatalog,proto3,oneof"`
}

type Catalog_ListCatalogs struct {
	ListCatalogs *ListCatalogs `protobuf:"bytes,26,opt,name=list_catalogs,json=listCatalogs,proto3,oneof"`
}

func (*Catalog_CurrentDatabase) isCatalog_CatType() {}

func (*Catalog_SetCurrentDatabase) isCatalog_CatType() {}

func (*Catalog_ListDatabases) isCatalog_CatType() {}

func (*Catalog_ListTables) isCatalog_CatType() {}

func (*Catalog_ListFunctions) isCatalog_CatType() {}

func (*Catalog_ListColumns) isCatalog_CatType() {}

func (*Catalog_GetDatabase) isCatalog_CatType() {}

func (*Catalog_GetTable) isCatalog_CatType() {}

func (*Catalog_GetFunction) isCatalog_CatType() {}

func (*Catalog_DatabaseExists) isCatalog_CatType() {}

func (*Catalog_TableExists) isCatalog_CatType() {}

func (*Catalog_FunctionExists) isCatalog_CatType() {}

func (*Catalog_CreateExternalTable) isCatalog_CatType() {}

func (*Catalog_CreateTable) isCatalog_CatType() {}

func (*Catalog_DropTempView) isCatalog_CatType() {}

func (*Catalog_DropGlobalTempView) isCatalog_CatType() {}

func (*Catalog_RecoverPartitions) isCatalog_CatType() {}

func (*Catalog_IsCached) isCatalog_CatType() {}

func (*Catalog_CacheTable) isCatalog_CatType() {}

func (*Catalog_UncacheTable) isCatalog_CatType() {}

func (*Catalog_ClearCache) isCatalog_CatType() {}

func (*Catalog_RefreshTable) isCatalog_CatType() {}

func (*Catalog_RefreshByPath) isCatalog_CatType() {}

func (*Catalog_CurrentCatalog) isCatalog_CatType() {}

func (*Catalog_SetCurrentCatalog) isCatalog_CatType() {}

func (*Catalog_ListCatalogs) isCatalog_CatType() {}

// See `spark.catalog.currentDatabase`
type CurrentDatabase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CurrentDatabase) Reset() {
	*x = CurrentDatabase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentDatabase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentDatabase) ProtoMessage() {}

func (x *CurrentDatabase) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentDatabase.ProtoReflect.Descriptor instead.
func (*CurrentDatabase) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{1}
}

// See `spark.catalog.setCurrentDatabase`
type SetCurrentDatabase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}

func (x *SetCurrentDatabase) Reset() {
	*x = SetCurrentDatabase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetCurrentDatabase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCurrentDatabase) ProtoMessage() {}

func (x *SetCurrentDatabase) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCurrentDatabase.ProtoReflect.Descriptor instead.
func (*SetCurrentDatabase) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{2}
}

func (x *SetCurrentDatabase) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

// See `spark.catalog.listDatabases`
type ListDatabases struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Optional) The pattern that the database name needs to match
	Pattern *string `protobuf:"bytes,1,opt,name=pattern,proto3,oneof" json:"pattern,omitempty"`
}

func (x *ListDatabases) Reset() {
	*x = ListDatabases{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatabases) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatabases) ProtoMessage() {}

func (x *ListDatabases) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatabases.ProtoReflect.Descriptor instead.
func (*ListDatabases) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{3}
}

func (x *ListDatabases) GetPattern() string {
	if x != nil && x.Pattern != nil {
		return *x.Pattern
	}
	return ""
}

// See `spark.catalog.listTables`
type ListTables struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Optional)
	DbName *string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3,oneof" json:"db_name,omitempty"`
	// (Optional) The pattern that the table name needs to match
	Pattern *string `protobuf:"bytes,2,opt,name=pattern,proto3,oneof" json:"pattern,omitempty"`
}

func (x *ListTables) Reset() {
	*x = ListTables{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTables) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTables) ProtoMessage() {}

func (x *ListTables) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTables.ProtoReflect.Descriptor instead.
func (*ListTables) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{4}
}

func (x *ListTables) GetDbName() string {
	if x != nil && x.DbName != nil {
		return *x.DbName
	}
	return ""
}

func (x *ListTables) GetPattern() string {
	if x != nil && x.Pattern != nil {
		return *x.Pattern
	}
	return ""
}

// See `spark.catalog.listFunctions`
type ListFunctions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Optional)
	DbName *string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3,oneof" json:"db_name,omitempty"`
	// (Optional) The pattern that the function name needs to match
	Pattern *string `protobuf:"bytes,2,opt,name=pattern,proto3,oneof" json:"pattern,omitempty"`
}

func (x *ListFunctions) Reset() {
	*x = ListFunctions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFunctions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFunctions) ProtoMessage() {}

func (x *ListFunctions) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFunctions.ProtoReflect.Descriptor instead.
func (*ListFunctions) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{5}
}

func (x *ListFunctions) GetDbName() string {
	if x != nil && x.DbName != nil {
		return *x.DbName
	}
	return ""
}

func (x *ListFunctions) GetPattern() string {
	if x != nil && x.Pattern != nil {
		return *x.Pattern
	}
	return ""
}

// See `spark.catalog.listColumns`
type ListColumns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// (Optional)
	DbName *string `protobuf:"bytes,2,opt,name=db_name,json=dbName,proto3,oneof" json:"db_name,omitempty"`
}

func (x *ListColumns) Reset() {
	*x = ListColumns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListColumns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListColumns) ProtoMessage() {}

func (x *ListColumns) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListColumns.ProtoReflect.Descriptor instead.
func (*ListColumns) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{6}
}

func (x *ListColumns) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *ListColumns) GetDbName() string {
	if x != nil && x.DbName != nil {
		return *x.DbName
	}
	return ""
}

// See `spark.catalog.getDatabase`
type GetDatabase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}

func (x *GetDatabase) Reset() {
	*x = GetDatabase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDatabase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatabase) ProtoMessage() {}

func (x *GetDatabase) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatabase.ProtoReflect.Descriptor instead.
func (*GetDatabase) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{7}
}

func (x *GetDatabase) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

// See `spark.catalog.getTable`
type GetTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// (Optional)
	DbName *string `protobuf:"bytes,2,opt,name=db_name,json=dbName,proto3,oneof" json:"db_name,omitempty"`
}

func (x *GetTable) Reset() {
	*x = GetTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTable) ProtoMessage() {}

func (x *GetTable) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTable.ProtoReflect.Descriptor instead.
func (*GetTable) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{8}
}

func (x *GetTable) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *GetTable) GetDbName() string {
	if x != nil && x.DbName != nil {
		return *x.DbName
	}
	return ""
}

// See `spark.catalog.getFunction`
type GetFunction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	// (Optional)
	DbName *string `protobuf:"bytes,2,opt,name=db_name,json=dbName,proto3,oneof" json:"db_name,omitempty"`
}

func (x *GetFunction) Reset() {
	*x = GetFunction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFunction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFunction) ProtoMessage() {}

func (x *GetFunction) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFunction.ProtoReflect.Descriptor instead.
func (*GetFunction) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{9}
}

func (x *GetFunction) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *GetFunction) GetDbName() string {
	if x != nil && x.DbName != nil {
		return *x.DbName
	}
	return ""
}

// See `spark.catalog.databaseExists`
type DatabaseExists struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}

func (x *DatabaseExists) Reset() {
	*x = DatabaseExists{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseExists) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseExists) ProtoMessage() {}

func (x *DatabaseExists) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseExists.ProtoReflect.Descriptor instead.
func (*DatabaseExists) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{10}
}

func (x *DatabaseExists) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

// See `spark.catalog.tableExists`
type TableExists struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// (Optional)
	DbName *string `protobuf:"bytes,2,opt,name=db_name,json=dbName,proto3,oneof" json:"db_name,omitempty"`
}

func (x *TableExists) Reset() {
	*x = TableExists{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableExists) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableExists) ProtoMessage() {}

func (x *TableExists) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableExists.ProtoReflect.Descriptor instead.
func (*TableExists) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{11}
}

func (x *TableExists) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *TableExists) GetDbName() string {
	if x != nil && x.DbName != nil {
		return *x.DbName
	}
	return ""
}

// See `spark.catalog.functionExists`
type FunctionExists struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	// (Optional)
	DbName *string `protobuf:"bytes,2,opt,name=db_name,json=dbName,proto3,oneof" json:"db_name,omitempty"`
}

func (x *FunctionExists) Reset() {
	*x = FunctionExists{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionExists) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionExists) ProtoMessage() {}

func (x *FunctionExists) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionExists.ProtoReflect.Descriptor instead.
func (*FunctionExists) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{12}
}

func (x *FunctionExists) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *FunctionExists) GetDbName() string {
	if x != nil && x.DbName != nil {
		return *x.DbName
	}
	return ""
}

// See `spark.catalog.createExternalTable`
type CreateExternalTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// (Optional)
	Path *string `protobuf:"bytes,2,opt,name=path,proto3,oneof" json:"path,omitempty"`
	// (Optional)
	Source *string `protobuf:"bytes,3,opt,name=source,proto3,oneof" json:"source,omitempty"`
	// (Optional)
	Schema *DataType `protobuf:"bytes,4,opt,name=schema,proto3,oneof" json:"schema,omitempty"`
	// Options could be empty for valid data source format.
	// The map key is case insensitive.
	Options map[string]string `protobuf:"bytes,5,rep,name=options,proto3" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CreateExternalTable) Reset() {
	*x = CreateExternalTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExternalTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExternalTable) ProtoMessage() {}

func (x *CreateExternalTable) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExternalTable.ProtoReflect.Descriptor instead.
func (*CreateExternalTable) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{13}
}

func (x *CreateExternalTable) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *CreateExternalTable) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *CreateExternalTable) GetSource() string {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return ""
}

func (x *CreateExternalTable) GetSchema() *DataType {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *CreateExternalTable) GetOptions() map[string]string {
	if x != nil {
		return x.Options
	}
	return nil
}

// See `spark.catalog.createTable`
type CreateTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// (Optional)
	Path *string `protobuf:"bytes,2,opt,name=path,proto3,oneof" json:"path,omitempty"`
	// (Optional)
	Source *string `protobuf:"bytes,3,opt,name=source,proto3,oneof" json:"source,omitempty"`
	// (Optional)
	Description *string `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// (Optional)
	Schema *DataType `protobuf:"bytes,5,opt,name=schema,proto3,oneof" json:"schema,omitempty"`
	// Options could be empty for valid data source format.
	// The map key is case insensitive.
	Options map[string]string `protobuf:"bytes,6,rep,name=options,proto3" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CreateTable) Reset() {
	*x = CreateTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTable) ProtoMessage() {}

func (x *CreateTable) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTable.ProtoReflect.Descriptor instead.
func (*CreateTable) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{14}
}

func (x *CreateTable) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *CreateTable) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *CreateTable) GetSource() string {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return ""
}

func (x *CreateTable) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *CreateTable) GetSchema() *DataType {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *CreateTable) GetOptions() map[string]string {
	if x != nil {
		return x.Options
	}
	return nil
}

// See `spark.catalog.dropTempView`
type DropTempView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	ViewName string `protobuf:"bytes,1,opt,name=view_name,json=viewName,proto3" json:"view_name,omitempty"`
}

func (x *DropTempView) Reset() {
	*x = DropTempView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropTempView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropTempView) ProtoMessage() {}

func (x *DropTempView) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropTempView.ProtoReflect.Descriptor instead.
func (*DropTempView) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{15}
}

func (x *DropTempView) GetViewName() string {
	if x != nil {
		return x.ViewName
	}
	return ""
}

// See `spark.catalog.dropGlobalTempView`
type DropGlobalTempView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	ViewName string `protobuf:"bytes,1,opt,name=view_name,json=viewName,proto3" json:"view_name,omitempty"`
}

func (x *DropGlobalTempView) Reset() {
	*x = DropGlobalTempView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropGlobalTempView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropGlobalTempView) ProtoMessage() {}

func (x *DropGlobalTempView) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropGlobalTempView.ProtoReflect.Descriptor instead.
func (*DropGlobalTempView) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{16}
}

func (x *DropGlobalTempView) GetViewName() string {
	if x != nil {
		return x.ViewName
	}
	return ""
}

// See `spark.catalog.recoverPartitions`
type RecoverPartitions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
}

func (x *RecoverPartitions) Reset() {
	*x = RecoverPartitions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecoverPartitions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecoverPartitions) ProtoMessage() {}

func (x *RecoverPartitions) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecoverPartitions.ProtoReflect.Descriptor instead.
func (*RecoverPartitions) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{17}
}

func (x *RecoverPartitions) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

// See `spark.catalog.isCached`
type IsCached struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
}

func (x *IsCached) Reset() {
	*x = IsCached{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsCached) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsCached) ProtoMessage() {}

func (x *IsCached) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsCached.ProtoReflect.Descriptor instead.
func (*IsCached) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{18}
}

func (x *IsCached) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

// See `spark.catalog.cacheTable`
type CacheTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// (Optional)
	StorageLevel *StorageLevel `protobuf:"bytes,2,opt,name=storage_level,json=storageLevel,proto3,oneof" json:"storage_level,omitempty"`
}

func (x *CacheTable) Reset() {
	*x = CacheTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheTable) ProtoMessage() {}

func (x *CacheTable) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheTable.ProtoReflect.Descriptor instead.
func (*CacheTable) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{19}
}

func (x *CacheTable) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *CacheTable) GetStorageLevel() *StorageLevel {
	if x != nil {
		return x.StorageLevel
	}
	return nil
}

// See `spark.catalog.uncacheTable`
type UncacheTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
}

func (x *UncacheTable) Reset() {
	*x = UncacheTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[20]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UncacheTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UncacheTable) ProtoMessage() {}

func (x *UncacheTable) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[20]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UncacheTable.ProtoReflect.Descriptor instead.
func (*UncacheTable) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{20}
}

func (x *UncacheTable) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

// See `spark.catalog.clearCache`
type ClearCache struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearCache) Reset() {
	*x = ClearCache{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[21]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearCache) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearCache) ProtoMessage() {}

func (x *ClearCache) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[21]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearCache.ProtoReflect.Descriptor instead.
func (*ClearCache) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{21}
}

// See `spark.catalog.refreshTable`
type RefreshTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
}

func (x *RefreshTable) Reset() {
	*x = RefreshTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[22]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTable) ProtoMessage() {}

func (x *RefreshTable) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[22]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTable.ProtoReflect.Descriptor instead.
func (*RefreshTable) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{22}
}

func (x *RefreshTable) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

// See `spark.catalog.refreshByPath`
type RefreshByPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *RefreshByPath) Reset() {
	*x = RefreshByPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[23]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshByPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshByPath) ProtoMessage() {}

func (x *RefreshByPath) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[23]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshByPath.ProtoReflect.Descriptor instead.
func (*RefreshByPath) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{23}
}

func (x *RefreshByPath) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// See `spark.catalog.currentCatalog`
type CurrentCatalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CurrentCatalog) Reset() {
	*x = CurrentCatalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[24]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentCatalog) ProtoMessage() {}

func (x *CurrentCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[24]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentCatalog.ProtoReflect.Descriptor instead.
func (*CurrentCatalog) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{24}
}

// See `spark.catalog.setCurrentCatalog`
type SetCurrentCatalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required)
	CatalogName string `protobuf:"bytes,1,opt,name=catalog_name,json=catalogName,proto3" json:"catalog_name,omitempty"`
}

func (x *SetCurrentCatalog) Reset() {
	*x = SetCurrentCatalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[25]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetCurrentCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCurrentCatalog) ProtoMessage() {}

func (x *SetCurrentCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[25]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCurrentCatalog.ProtoReflect.Descriptor instead.
func (*SetCurrentCatalog) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{25}
}

func (x *SetCurrentCatalog) GetCatalogName() string {
	if x != nil {
		return x.CatalogName
	}
	return ""
}

// See `spark.catalog.listCatalogs`
type ListCatalogs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Optional) The pattern that the catalog name needs to match
	Pattern *string `protobuf:"bytes,1,opt,name=pattern,proto3,oneof" json:"pattern,omitempty"`
}

func (x *ListCatalogs) Reset() {
	*x = ListCatalogs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spark_connect_catalog_proto_msgTypes[26]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCatalogs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatalogs) ProtoMessage() {}

func (x *ListCatalogs) ProtoReflect() protoreflect.Message {
	mi := &file_spark_connect_catalog_proto_msgTypes[26]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatalogs.ProtoReflect.Descriptor instead.
func (*ListCatalogs) Descriptor() ([]byte, []int) {
	return file_spark_connect_catalog_proto_rawDescGZIP(), []int{26}
}

func (x *ListCatalogs) GetPattern() string {
	if x != nil && x.Pattern != nil {
		return *x.Pattern
	}
	return ""
}

var File_spark_connect_catalog_proto protoreflect.FileDescriptor

var file_spark_connect_catalog_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x1a, 0x1a, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x0e, 0x0a, 0x07, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12,
	0x4b, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x14,
	0x73, 0x65, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x12, 0x73, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x6c, 0x69, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x6c, 0x69,
	0x73, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x0e, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00,
	0x52, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x3f, 0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x73, 0x48, 0x00, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73,
	0x12, 0x3f, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52,
	0x08, 0x67, 0x65, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x67, 0x65, 0x74,
	0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x67,
	0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x0f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x48, 0x00, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x48, 0x0a, 0x0f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x48, 0x00, 0x52,
	0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12,
	0x58, 0x0a, 0x15, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x48, 0x00, 0x52, 0x13, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x64, 0x72,
	0x6f, 0x70, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x56, 0x69, 0x65, 0x77, 0x48,
	0x00, 0x52, 0x0c, 0x64, 0x72, 0x6f, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x56, 0x69, 0x65, 0x77, 0x12,
	0x56, 0x0a, 0x15, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x74,
	0x65, 0x6d, 0x70, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44,
	0x72, 0x6f, 0x70, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x56, 0x69, 0x65,
	0x77, 0x48, 0x00, 0x52, 0x12, 0x64, 0x72, 0x6f, 0x70, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x54,
	0x65, 0x6d, 0x70, 0x56, 0x69, 0x65, 0x77, 0x12, 0x51, 0x0a, 0x12, 0x72, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x11, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x49, 0x73,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x69, 0x73, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x64, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x42, 0x0a, 0x0d, 0x75, 0x6e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x55, 0x6e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x75, 0x6e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x62, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x42, 0x79, 0x50, 0x61, 0x74, 0x68, 0x48, 0x00, 0x52,
	0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x42, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x48,
	0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x52, 0x0a, 0x13, 0x73, 0x65, 0x74, 0x5f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x18,
	0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x48, 0x00, 0x52, 0x11, 0x73, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x42, 0x0a, 0x0d,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73,
	0x48, 0x00, 0x52, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73,
	0x42, 0x0a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x11, 0x0a, 0x0f,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x22,
	0x2d, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3a,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x12,
	0x1d, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x22, 0x61, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x64, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x22, 0x64, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c,
	0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07,
	0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x6e, 0x22, 0x56, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x53, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x07,
	0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64,
	0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x56, 0x0a, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5f, 0x0a, 0x0e, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc6, 0x02, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x48, 0x02, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x88, 0x01, 0x01, 0x12, 0x49, 0x0a,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x22, 0xed, 0x02, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x34, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x48, 0x03, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x22, 0x2b, 0x0a, 0x0c, 0x44, 0x72, 0x6f, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x56,
	0x69, 0x65, 0x77, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x31, 0x0a, 0x12, 0x44, 0x72, 0x6f, 0x70, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x54, 0x65,
	0x6d, 0x70, 0x56, 0x69, 0x65, 0x77, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x65, 0x77, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x08, 0x49, 0x73, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x0a, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x45, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x2d, 0x0a, 0x0c, 0x55, 0x6e, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x43, 0x6c, 0x65, 0x61,
	0x72, 0x43, 0x61, 0x63, 0x68, 0x65, 0x22, 0x2d, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x42, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x22, 0x36, 0x0a, 0x11,
	0x53, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x42,
	0x36, 0x0a, 0x1e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spark_connect_catalog_proto_rawDescOnce sync.Once
	file_spark_connect_catalog_proto_rawDescData = file_spark_connect_catalog_proto_rawDesc
)

func file_spark_connect_catalog_proto_rawDescGZIP() []byte {
	file_spark_connect_catalog_proto_rawDescOnce.Do(func() {
		file_spark_connect_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_spark_connect_catalog_proto_rawDescData)
	})
	return file_spark_connect_catalog_proto_rawDescData
}

var file_spark_connect_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 29)
var file_spark_connect_catalog_proto_goTypes = []interface{}{
	(*Catalog)(nil),             // 0: spark.connect.Catalog
	(*CurrentDatabase)(nil),     // 1: spark.connect.CurrentDatabase
	(*SetCurrentDatabase)(nil),  // 2: spark.connect.SetCurrentDatabase
	(*ListDatabases)(nil),       // 3: spark.connect.ListDatabases
	(*ListTables)(nil),          // 4: spark.connect.ListTables
	(*ListFunctions)(nil),       // 5: spark.connect.ListFunctions
	(*ListColumns)(nil),         // 6: spark.connect.ListColumns
	(*GetDatabase)(nil),         // 7: spark.connect.GetDatabase
	(*GetTable)(nil),            // 8: spark.connect.GetTable
	(*GetFunction)(nil),         // 9: spark.connect.GetFunction
	(*DatabaseExists)(nil),      // 10: spark.connect.DatabaseExists
	(*TableExists)(nil),         // 11: spark.connect.TableExists
	(*FunctionExists)(nil),      // 12: spark.connect.FunctionExists
	(*CreateExternalTable)(nil), // 13: spark.connect.CreateExternalTable
	(*CreateTable)(nil),         // 14: spark.connect.CreateTable
	(*DropTempView)(nil),        // 15: spark.connect.DropTempView
	(*DropGlobalTempView)(nil),  // 16: spark.connect.DropGlobalTempView
	(*RecoverPartitions)(nil),   // 17: spark.connect.RecoverPartitions
	(*IsCached)(nil),            // 18: spark.connect.IsCached
	(*CacheTable)(nil),          // 19: spark.connect.CacheTable
	(*UncacheTable)(nil),        // 20: spark.connect.UncacheTable
	(*ClearCache)(nil),          // 21: spark.connect.ClearCache
	(*RefreshTable)(nil),        // 22: spark.connect.RefreshTable
	(*RefreshByPath)(nil),       // 23: spark.connect.RefreshByPath
	(*CurrentCatalog)(nil),      // 24: spark.connect.CurrentCatalog
	(*SetCurrentCatalog)(nil),   // 25: spark.connect.SetCurrentCatalog
	(*ListCatalogs)(nil),        // 26: spark.connect.ListCatalogs
	nil,                         // 27: spark.connect.CreateExternalTable.OptionsEntry
	nil,                         // 28: spark.connect.CreateTable.OptionsEntry
	(*DataType)(nil),            // 29: spark.connect.DataType
	(*StorageLevel)(nil),        // 30: spark.connect.StorageLevel
}
var file_spark_connect_catalog_proto_depIdxs = []int32{
	1,  // 0: spark.connect.Catalog.current_database:type_name -> spark.connect.CurrentDatabase
	2,  // 1: spark.connect.Catalog.set_current_database:type_name -> spark.connect.SetCurrentDatabase
	3,  // 2: spark.connect.Catalog.list_databases:type_name -> spark.connect.ListDatabases
	4,  // 3: spark.connect.Catalog.list_tables:type_name -> spark.connect.ListTables
	5,  // 4: spark.connect.Catalog.list_functions:type_name -> spark.connect.ListFunctions
	6,  // 5: spark.connect.Catalog.list_columns:type_name -> spark.connect.ListColumns
	7,  // 6: spark.connect.Catalog.get_database:type_name -> spark.connect.GetDatabase
	8,  // 7: spark.connect.Catalog.get_table:type_name -> spark.connect.GetTable
	9,  // 8: spark.connect.Catalog.get_function:type_name -> spark.connect.GetFunction
	10, // 9: spark.connect.Catalog.database_exists:type_name -> spark.connect.DatabaseExists
	11, // 10: spark.connect.Catalog.table_exists:type_name -> spark.connect.TableExists
	12, // 11: spark.connect.Catalog.function_exists:type_name -> spark.connect.FunctionExists
	13, // 12: spark.connect.Catalog.create_external_table:type_name -> spark.connect.CreateExternalTable
	14, // 13: spark.connect.Catalog.create_table:type_name -> spark.connect.CreateTable
	15, // 14: spark.connect.Catalog.drop_temp_view:type_name -> spark.connect.DropTempView
	16, // 15: spark.connect.Catalog.drop_global_temp_view:type_name -> spark.connect.DropGlobalTempView
	17, // 16: spark.connect.Catalog.recover_partitions:type_name -> spark.connect.RecoverPartitions
	18, // 17: spark.connect.Catalog.is_cached:type_name -> spark.connect.IsCached
	19, // 18: spark.connect.Catalog.cache_table:type_name -> spark.connect.CacheTable
	20, // 19: spark.connect.Catalog.uncache_table:type_name -> spark.connect.UncacheTable
	21, // 20: spark.connect.Catalog.clear_cache:type_name -> spark.connect.ClearCache
	22, // 21: spark.connect.Catalog.refresh_table:type_name -> spark.connect.RefreshTable
	23, // 22: spark.connect.Catalog.refresh_by_path:type_name -> spark.connect.RefreshByPath
	24, // 23: spark.connect.Catalog.current_catalog:type_name -> spark.connect.CurrentCatalog
	25, // 24: spark.connect.Catalog.set_current_catalog:type_name -> spark.connect.SetCurrentCatalog
	26, // 25: spark.connect.Catalog.list_catalogs:type_name -> spark.connect.ListCatalogs
	29, // 26: spark.connect.CreateExternalTable.schema:type_name -> spark.connect.DataType
	27, // 27: spark.connect.CreateExternalTable.options:type_name -> spark.connect.CreateExternalTable.OptionsEntry
	29, // 28: spark.connect.CreateTable.schema:type_name -> spark.connect.DataType
	28, // 29: spark.connect.CreateTable.options:type_name -> spark.connect.CreateTable.OptionsEntry
	30, // 30: spark.connect.CacheTable.storage_level:type_name -> spark.connect.StorageLevel
	31, // [31:31] is the sub-list for method output_type
	31, // [31:31] is the sub-list for method input_type
	31, // [31:31] is the sub-list for extension type_name
	31, // [31:31] is the sub-list for extension extendee
	0,  // [0:31] is the sub-list for field type_name
}

func init() { file_spark_connect_catalog_proto_init() }
func file_spark_connect_catalog_proto_init() {
	if File_spark_connect_catalog_proto != nil {
		return
	}
	file_spark_connect_common_proto_init()
	file_spark_connect_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spark_connect_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Catalog); i {
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
		file_spark_connect_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentDatabase); i {
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
		file_spark_connect_catalog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetCurrentDatabase); i {
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
		file_spark_connect_catalog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDatabases); i {
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
		file_spark_connect_catalog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTables); i {
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
		file_spark_connect_catalog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFunctions); i {
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
		file_spark_connect_catalog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListColumns); i {
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
		file_spark_connect_catalog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDatabase); i {
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
		file_spark_connect_catalog_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTable); i {
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
		file_spark_connect_catalog_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFunction); i {
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
		file_spark_connect_catalog_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseExists); i {
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
		file_spark_connect_catalog_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableExists); i {
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
		file_spark_connect_catalog_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionExists); i {
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
		file_spark_connect_catalog_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExternalTable); i {
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
		file_spark_connect_catalog_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTable); i {
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
		file_spark_connect_catalog_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropTempView); i {
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
		file_spark_connect_catalog_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropGlobalTempView); i {
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
		file_spark_connect_catalog_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecoverPartitions); i {
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
		file_spark_connect_catalog_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsCached); i {
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
		file_spark_connect_catalog_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheTable); i {
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
		file_spark_connect_catalog_proto_msgTypes[20].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UncacheTable); i {
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
		file_spark_connect_catalog_proto_msgTypes[21].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearCache); i {
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
		file_spark_connect_catalog_proto_msgTypes[22].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTable); i {
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
		file_spark_connect_catalog_proto_msgTypes[23].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshByPath); i {
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
		file_spark_connect_catalog_proto_msgTypes[24].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentCatalog); i {
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
		file_spark_connect_catalog_proto_msgTypes[25].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetCurrentCatalog); i {
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
		file_spark_connect_catalog_proto_msgTypes[26].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCatalogs); i {
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
	file_spark_connect_catalog_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Catalog_CurrentDatabase)(nil),
		(*Catalog_SetCurrentDatabase)(nil),
		(*Catalog_ListDatabases)(nil),
		(*Catalog_ListTables)(nil),
		(*Catalog_ListFunctions)(nil),
		(*Catalog_ListColumns)(nil),
		(*Catalog_GetDatabase)(nil),
		(*Catalog_GetTable)(nil),
		(*Catalog_GetFunction)(nil),
		(*Catalog_DatabaseExists)(nil),
		(*Catalog_TableExists)(nil),
		(*Catalog_FunctionExists)(nil),
		(*Catalog_CreateExternalTable)(nil),
		(*Catalog_CreateTable)(nil),
		(*Catalog_DropTempView)(nil),
		(*Catalog_DropGlobalTempView)(nil),
		(*Catalog_RecoverPartitions)(nil),
		(*Catalog_IsCached)(nil),
		(*Catalog_CacheTable)(nil),
		(*Catalog_UncacheTable)(nil),
		(*Catalog_ClearCache)(nil),
		(*Catalog_RefreshTable)(nil),
		(*Catalog_RefreshByPath)(nil),
		(*Catalog_CurrentCatalog)(nil),
		(*Catalog_SetCurrentCatalog)(nil),
		(*Catalog_ListCatalogs)(nil),
	}
	file_spark_connect_catalog_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_spark_connect_catalog_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_spark_connect_catalog_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_spark_connect_catalog_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_spark_connect_catalog_proto_msgTypes[8].OneofWrappers = []interface{}{}
	file_spark_connect_catalog_proto_msgTypes[9].OneofWrappers = []interface{}{}
	file_spark_connect_catalog_proto_msgTypes[11].OneofWrappers = []interface{}{}
	file_spark_connect_catalog_proto_msgTypes[12].OneofWrappers = []interface{}{}
	file_spark_connect_catalog_proto_msgTypes[13].OneofWrappers = []interface{}{}
	file_spark_connect_catalog_proto_msgTypes[14].OneofWrappers = []interface{}{}
	file_spark_connect_catalog_proto_msgTypes[19].OneofWrappers = []interface{}{}
	file_spark_connect_catalog_proto_msgTypes[26].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spark_connect_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   29,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spark_connect_catalog_proto_goTypes,
		DependencyIndexes: file_spark_connect_catalog_proto_depIdxs,
		MessageInfos:      file_spark_connect_catalog_proto_msgTypes,
	}.Build()
	File_spark_connect_catalog_proto = out.File
	file_spark_connect_catalog_proto_rawDesc = nil
	file_spark_connect_catalog_proto_goTypes = nil
	file_spark_connect_catalog_proto_depIdxs = nil
}
