// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	model "github.com/FrancoLiberali/cql/model"
	overrideforeignkeyinverse "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/overrideforeignkeyinverse"
	"time"
)

type creditCardConditions struct {
	ID            condition.Field[overrideforeignkeyinverse.CreditCard, model.UUID]
	CreatedAt     condition.Field[overrideforeignkeyinverse.CreditCard, time.Time]
	UpdatedAt     condition.Field[overrideforeignkeyinverse.CreditCard, time.Time]
	DeletedAt     condition.Field[overrideforeignkeyinverse.CreditCard, time.Time]
	UserReference condition.Field[overrideforeignkeyinverse.CreditCard, model.UUID]
}

var CreditCard = creditCardConditions{
	CreatedAt:     condition.Field[overrideforeignkeyinverse.CreditCard, time.Time]{Name: "CreatedAt"},
	DeletedAt:     condition.Field[overrideforeignkeyinverse.CreditCard, time.Time]{Name: "DeletedAt"},
	ID:            condition.Field[overrideforeignkeyinverse.CreditCard, model.UUID]{Name: "ID"},
	UpdatedAt:     condition.Field[overrideforeignkeyinverse.CreditCard, time.Time]{Name: "UpdatedAt"},
	UserReference: condition.Field[overrideforeignkeyinverse.CreditCard, model.UUID]{Name: "UserReference"},
}

// Preload allows preloading the CreditCard when doing a query
func (creditCardConditions creditCardConditions) Preload() condition.Condition[overrideforeignkeyinverse.CreditCard] {
	return condition.NewPreloadCondition[overrideforeignkeyinverse.CreditCard](creditCardConditions.ID, creditCardConditions.CreatedAt, creditCardConditions.UpdatedAt, creditCardConditions.DeletedAt, creditCardConditions.UserReference)
}
