// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	belongsto "github.com/ditrit/badaas-orm/cli/cmd/gen/conditions/tests/belongsto"
	orm "github.com/ditrit/badaas/orm"
	"time"
)

func OwnerId(operator orm.Operator[orm.UUID]) orm.WhereCondition[belongsto.Owner] {
	return orm.FieldCondition[belongsto.Owner, orm.UUID]{
		FieldIdentifier: orm.IDFieldID,
		Operator:        operator,
	}
}
func OwnerCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[belongsto.Owner] {
	return orm.FieldCondition[belongsto.Owner, time.Time]{
		FieldIdentifier: orm.CreatedAtFieldID,
		Operator:        operator,
	}
}
func OwnerUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[belongsto.Owner] {
	return orm.FieldCondition[belongsto.Owner, time.Time]{
		FieldIdentifier: orm.UpdatedAtFieldID,
		Operator:        operator,
	}
}
func OwnerDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[belongsto.Owner] {
	return orm.FieldCondition[belongsto.Owner, time.Time]{
		FieldIdentifier: orm.DeletedAtFieldID,
		Operator:        operator,
	}
}

var OwnerPreloadAttributes = orm.NewPreloadCondition[belongsto.Owner]()
