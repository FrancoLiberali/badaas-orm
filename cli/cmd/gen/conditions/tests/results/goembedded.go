// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	goembedded "github.com/ditrit/badaas-orm/cli/cmd/gen/conditions/tests/goembedded"
	orm "github.com/ditrit/badaas/orm"
	"time"
)

func GoEmbeddedId(operator orm.Operator[uint]) orm.WhereCondition[goembedded.GoEmbedded] {
	return orm.FieldCondition[goembedded.GoEmbedded, uint]{
		Field:    "ID",
		Operator: operator,
	}
}
func GoEmbeddedCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[goembedded.GoEmbedded] {
	return orm.FieldCondition[goembedded.GoEmbedded, time.Time]{
		Field:    "CreatedAt",
		Operator: operator,
	}
}
func GoEmbeddedUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[goembedded.GoEmbedded] {
	return orm.FieldCondition[goembedded.GoEmbedded, time.Time]{
		Field:    "UpdatedAt",
		Operator: operator,
	}
}
func GoEmbeddedDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[goembedded.GoEmbedded] {
	return orm.FieldCondition[goembedded.GoEmbedded, time.Time]{
		Field:    "DeletedAt",
		Operator: operator,
	}
}
func GoEmbeddedEmbeddedInt(operator orm.Operator[int]) orm.WhereCondition[goembedded.GoEmbedded] {
	return orm.FieldCondition[goembedded.GoEmbedded, int]{
		Field:    "EmbeddedInt",
		Operator: operator,
	}
}