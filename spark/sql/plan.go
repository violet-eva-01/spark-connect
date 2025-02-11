package sql

import (
	"sync/atomic"

	proto "github.com/violet-eva-01/spark-connect/internal/generated"
)

var atomicInt64 atomic.Int64

func newPlanId() *int64 {
	v := atomicInt64.Add(1)
	return &v
}

func resetPlanIdForTesting() {
	atomicInt64.Swap(0)
}

func newReadTableRelation(table string) *proto.Relation {
	return &proto.Relation{
		Common: &proto.RelationCommon{
			PlanId: newPlanId(),
		},
		RelType: &proto.Relation_Read{
			Read: &proto.Read{
				ReadType: &proto.Read_NamedTable_{
					NamedTable: &proto.Read_NamedTable{
						UnparsedIdentifier: table,
					},
				},
			},
		},
	}
}

func newReadWithFormatAndPath(path, format string) *proto.Relation {
	return &proto.Relation{
		RelType: &proto.Relation_Read{
			Read: &proto.Read{
				ReadType: &proto.Read_DataSource_{
					DataSource: &proto.Read_DataSource{
						Format: &format,
						Paths:  []string{path},
					},
				},
			},
		},
	}
}

func newReadWithFormatAndPathAndOptions(path, format string, options map[string]string) *proto.Relation {
	return &proto.Relation{
		RelType: &proto.Relation_Read{
			Read: &proto.Read{
				ReadType: &proto.Read_DataSource_{
					DataSource: &proto.Read_DataSource{
						Format:  &format,
						Paths:   []string{path},
						Options: options,
					},
				},
			},
		},
	}
}
