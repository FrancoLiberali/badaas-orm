// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	overrideforeignkeyinverse "github.com/ditrit/badaas-orm/cli/cmd/gen/conditions/tests/overrideforeignkeyinverse"
	orm "github.com/ditrit/badaas/orm"
	"time"
)

func UserId(operator orm.Operator[orm.UUID]) orm.WhereCondition[overrideforeignkeyinverse.User] {
	return orm.FieldCondition[overrideforeignkeyinverse.User, orm.UUID]{
		FieldIdentifier: orm.IDFieldID,
		Operator:        operator,
	}
}
func UserCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[overrideforeignkeyinverse.User] {
	return orm.FieldCondition[overrideforeignkeyinverse.User, time.Time]{
		FieldIdentifier: orm.CreatedAtFieldID,
		Operator:        operator,
	}
}
func UserUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[overrideforeignkeyinverse.User] {
	return orm.FieldCondition[overrideforeignkeyinverse.User, time.Time]{
		FieldIdentifier: orm.UpdatedAtFieldID,
		Operator:        operator,
	}
}
func UserDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[overrideforeignkeyinverse.User] {
	return orm.FieldCondition[overrideforeignkeyinverse.User, time.Time]{
		FieldIdentifier: orm.DeletedAtFieldID,
		Operator:        operator,
	}
}
func UserCreditCard(conditions ...orm.Condition[overrideforeignkeyinverse.CreditCard]) orm.IJoinCondition[overrideforeignkeyinverse.User] {
	return orm.JoinCondition[overrideforeignkeyinverse.User, overrideforeignkeyinverse.CreditCard]{
		Conditions:         conditions,
		RelationField:      "CreditCard",
		T1Field:            "ID",
		T1PreloadCondition: UserPreloadAttributes,
		T2Field:            "UserReference",
	}
}

var UserPreloadCreditCard = UserCreditCard(CreditCardPreloadAttributes)
var UserPreloadAttributes = orm.NewPreloadCondition[overrideforeignkeyinverse.User]()
var UserPreloadRelations = []orm.Condition[overrideforeignkeyinverse.User]{UserPreloadCreditCard}
