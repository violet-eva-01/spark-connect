package utils

import proto "github.com/violet-eva-01/spark-connect/internal/generatedCustom"

type ExplainMode int

const (
	ExplainModeSimple    ExplainMode = iota
	ExplainModeExtended  ExplainMode = iota
	ExplainModeCodegen   ExplainMode = iota
	ExplainModeCost      ExplainMode = iota
	ExplainModeFormatted ExplainMode = iota
)

type StorageLevel int

const (
	StorageLevelDiskOnly          StorageLevel = iota
	StorageLevelDiskOnly2         StorageLevel = iota
	StorageLevelDiskOnly3         StorageLevel = iota
	StorageLevelMemoryAndDisk     StorageLevel = iota
	StorageLevelMemoryAndDisk2    StorageLevel = iota
	StorageLevelMemoryOnly        StorageLevel = iota
	StorageLevelMemoryOnly2       StorageLevel = iota
	StorageLevelMemoyAndDiskDeser StorageLevel = iota
	StorageLevelNone              StorageLevel = iota
	StorageLevelOffHeap           StorageLevel = iota
)

func ToProtoStorageLevel(level StorageLevel) *proto.StorageLevel {
	switch level {
	case StorageLevelDiskOnly:
		return &proto.StorageLevel{UseDisk: true, UseMemory: false, Replication: 1}
	case StorageLevelDiskOnly2:
		return &proto.StorageLevel{UseDisk: true, UseMemory: false, Replication: 2}
	case StorageLevelDiskOnly3:
		return &proto.StorageLevel{UseDisk: true, UseMemory: false, Replication: 3}
	case StorageLevelMemoryAndDisk:
		return &proto.StorageLevel{UseDisk: true, UseMemory: true, Replication: 1}
	case StorageLevelMemoryAndDisk2:
		return &proto.StorageLevel{UseDisk: true, UseMemory: true, Replication: 2}
	case StorageLevelMemoryOnly:
		return &proto.StorageLevel{UseDisk: false, UseMemory: true, Replication: 1}
	case StorageLevelMemoryOnly2:
		return &proto.StorageLevel{UseDisk: false, UseMemory: true, Replication: 2}
	case StorageLevelMemoyAndDiskDeser:
		return &proto.StorageLevel{UseDisk: true, UseMemory: true, Replication: 1, Deserialized: true}
	case StorageLevelOffHeap:
		return &proto.StorageLevel{UseDisk: true, UseMemory: true, UseOffHeap: true, Replication: 1}
	default:
		return &proto.StorageLevel{UseDisk: false, UseMemory: false, UseOffHeap: false, Replication: 1}
	}
}

func FromProtoStorageLevel(level *proto.StorageLevel) StorageLevel {
	if level.UseDisk && level.UseMemory && level.Replication <= 1 && !level.Deserialized && !level.UseOffHeap {
		return StorageLevelMemoryAndDisk
	} else if level.UseDisk && level.UseMemory && level.Replication == 2 && !level.Deserialized && !level.UseOffHeap {
		return StorageLevelMemoryAndDisk2
	} else if level.UseDisk && !level.UseMemory && level.Replication == 3 &&
		!level.Deserialized && !level.UseOffHeap {
		return StorageLevelDiskOnly3
	} else if level.UseDisk && !level.UseMemory && level.Replication == 2 &&
		!level.Deserialized && !level.UseOffHeap {
		return StorageLevelDiskOnly2
	} else if level.UseDisk && !level.UseMemory && level.Replication <= 1 &&
		!level.Deserialized && !level.UseOffHeap {
		return StorageLevelDiskOnly
	} else if !level.UseDisk && level.UseMemory && level.Replication <= 1 &&
		!level.Deserialized && !level.UseOffHeap {
		return StorageLevelMemoryOnly
	} else if !level.UseDisk && level.UseMemory && level.Replication == 2 &&
		!level.Deserialized && !level.UseOffHeap {
		return StorageLevelMemoryOnly2
	} else if level.UseDisk && level.UseMemory && level.Replication <= 1 && level.Deserialized && !level.UseOffHeap {
		return StorageLevelMemoyAndDiskDeser
	} else if !level.UseDisk && !level.UseMemory && !level.Deserialized && !level.UseOffHeap {
		return StorageLevelNone
	} else if level.UseOffHeap && !level.Deserialized {
		return StorageLevelOffHeap
	}
	return StorageLevelNone
}

type JoinType int

const (
	JoinTypeInner      JoinType = iota
	JoinTypeLeftOuter  JoinType = iota
	JoinTypeRightOuter JoinType = iota
	JoinTypeFullOuter  JoinType = iota
	JoinTypeLeftSemi   JoinType = iota
	JoinTypeLeftAnti   JoinType = iota
	JoinTypeCross      JoinType = iota
)

func ToProtoJoinType(joinType JoinType) proto.Join_JoinType {
	switch joinType {
	case JoinTypeInner:
		return proto.Join_JOIN_TYPE_INNER
	case JoinTypeLeftOuter:
		return proto.Join_JOIN_TYPE_LEFT_OUTER
	case JoinTypeRightOuter:
		return proto.Join_JOIN_TYPE_RIGHT_OUTER
	case JoinTypeFullOuter:
		return proto.Join_JOIN_TYPE_FULL_OUTER
	case JoinTypeLeftSemi:
		return proto.Join_JOIN_TYPE_LEFT_SEMI
	case JoinTypeLeftAnti:
		return proto.Join_JOIN_TYPE_LEFT_ANTI
	case JoinTypeCross:
		return proto.Join_JOIN_TYPE_CROSS
	default:
		return proto.Join_JOIN_TYPE_INNER
	}
}
