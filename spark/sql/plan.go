package sql

import (
	"sync/atomic"

	proto "github.com/violet-eva-01/spark-connect/internal/generatedCustom"
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

func newReadTableRelationAndOptions(table string, options map[string]string) *proto.Relation {
	return &proto.Relation{
		Common: &proto.RelationCommon{
			PlanId: newPlanId(),
		},
		RelType: &proto.Relation_Read{
			Read: &proto.Read{
				ReadType: &proto.Read_NamedTable_{
					NamedTable: &proto.Read_NamedTable{
						UnparsedIdentifier: table,
						Options:            options,
					},
				},
			},
		},
	}
}

func newReadStreamTableRelation(table string) *proto.Relation {
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
				IsStreaming: true,
			},
		},
	}
}

func newReadStreamTableRelationAndOptions(table string, options map[string]string) *proto.Relation {
	return &proto.Relation{
		Common: &proto.RelationCommon{
			PlanId: newPlanId(),
		},
		RelType: &proto.Relation_Read{
			Read: &proto.Read{
				ReadType: &proto.Read_NamedTable_{
					NamedTable: &proto.Read_NamedTable{
						UnparsedIdentifier: table,
						Options:            options,
					},
				},
				IsStreaming: true,
			},
		},
	}
}

func newReadWithFormat(format string) *proto.Relation {
	return &proto.Relation{
		RelType: &proto.Relation_Read{
			Read: &proto.Read{
				ReadType: &proto.Read_DataSource_{
					DataSource: &proto.Read_DataSource{
						Format: &format,
					},
				},
			},
		},
	}
}

func newReadWithFormatAndOptions(format string, options map[string]string) *proto.Relation {
	return &proto.Relation{
		RelType: &proto.Relation_Read{
			Read: &proto.Read{
				ReadType: &proto.Read_DataSource_{
					DataSource: &proto.Read_DataSource{
						Format:  &format,
						Options: options,
					},
				},
			},
		},
	}
}

func newReadWithFormatAndPath(path []string, format string) *proto.Relation {
	return &proto.Relation{
		RelType: &proto.Relation_Read{
			Read: &proto.Read{
				ReadType: &proto.Read_DataSource_{
					DataSource: &proto.Read_DataSource{
						Format: &format,
						Paths:  path,
					},
				},
			},
		},
	}
}

func newReadWithFormatAndPathAndOptions(path []string, format string, options map[string]string) *proto.Relation {
	return &proto.Relation{
		RelType: &proto.Relation_Read{
			Read: &proto.Read{
				ReadType: &proto.Read_DataSource_{
					DataSource: &proto.Read_DataSource{
						Format:  &format,
						Paths:   path,
						Options: options,
					},
				},
			},
		},
	}
}

func newReadStreamWithFormatAndPath(path []string, format string) *proto.Relation {
	return &proto.Relation{
		RelType: &proto.Relation_Read{
			Read: &proto.Read{
				ReadType: &proto.Read_DataSource_{
					DataSource: &proto.Read_DataSource{
						Format: &format,
						Paths:  path,
					},
				},
				IsStreaming: true,
			},
		},
	}
}

func newReadStreamWithFormatAndPathAndOptions(path []string, format string, options map[string]string) *proto.Relation {
	return &proto.Relation{
		RelType: &proto.Relation_Read{
			Read: &proto.Read{
				ReadType: &proto.Read_DataSource_{
					DataSource: &proto.Read_DataSource{
						Format:  &format,
						Paths:   path,
						Options: options,
					},
				},
				IsStreaming: true,
			},
		},
	}
}

func newReadStreamWithFormat(format string) *proto.Relation {
	return &proto.Relation{
		RelType: &proto.Relation_Read{
			Read: &proto.Read{
				ReadType: &proto.Read_DataSource_{
					DataSource: &proto.Read_DataSource{
						Format: &format,
					},
				},
				IsStreaming: true,
			},
		},
	}
}

func newReadStreamWithFormatAndOptions(format string, options map[string]string) *proto.Relation {
	return &proto.Relation{
		RelType: &proto.Relation_Read{
			Read: &proto.Read{
				ReadType: &proto.Read_DataSource_{
					DataSource: &proto.Read_DataSource{
						Format:  &format,
						Options: options,
					},
				},
				IsStreaming: true,
			},
		},
	}
}
