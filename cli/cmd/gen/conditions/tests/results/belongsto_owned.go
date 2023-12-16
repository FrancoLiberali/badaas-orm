// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	belongsto "github.com/ditrit/badaas-orm/cli/cmd/gen/conditions/tests/belongsto"
	orm "github.com/ditrit/badaas/orm"
	"time"
)

func OwnedId(operator orm.Operator[orm.UUID]) orm.WhereCondition[belongsto.Owned] {
	return orm.FieldCondition[belongsto.Owned, orm.UUID]{
		Field:    "ID",
		Operator: operator,
	}
}
func OwnedCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[belongsto.Owned] {
	return orm.FieldCondition[belongsto.Owned, time.Time]{
		Field:    "CreatedAt",
		Operator: operator,
	}
}
func OwnedUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[belongsto.Owned] {
	return orm.FieldCondition[belongsto.Owned, time.Time]{
		Field:    "UpdatedAt",
		Operator: operator,
	}
}
func OwnedDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[belongsto.Owned] {
	return orm.FieldCondition[belongsto.Owned, time.Time]{
		Field:    "DeletedAt",
		Operator: operator,
	}
}
func OwnedOwner(conditions ...orm.Condition[belongsto.Owner]) orm.Condition[belongsto.Owned] {
	return orm.JoinCondition[belongsto.Owned, belongsto.Owner]{
		Conditions: conditions,
		T1Field:    "OwnerID",
		T2Field:    "ID",
	}
}
func OwnedOwnerId(operator orm.Operator[orm.UUID]) orm.WhereCondition[belongsto.Owned] {
	return orm.FieldCondition[belongsto.Owned, orm.UUID]{
		Field:    "OwnerID",
		Operator: operator,
	}
}