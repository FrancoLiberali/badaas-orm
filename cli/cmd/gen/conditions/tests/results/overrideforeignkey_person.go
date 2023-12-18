// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	model "github.com/FrancoLiberali/cql/model"
	overrideforeignkey "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/overrideforeignkey"
	"time"
)

type personConditions struct {
	ID        condition.Field[overrideforeignkey.Person, model.UUID]
	CreatedAt condition.Field[overrideforeignkey.Person, time.Time]
	UpdatedAt condition.Field[overrideforeignkey.Person, time.Time]
	DeletedAt condition.Field[overrideforeignkey.Person, time.Time]
}

var Person = personConditions{
	CreatedAt: condition.Field[overrideforeignkey.Person, time.Time]{Name: "CreatedAt"},
	DeletedAt: condition.Field[overrideforeignkey.Person, time.Time]{Name: "DeletedAt"},
	ID:        condition.Field[overrideforeignkey.Person, model.UUID]{Name: "ID"},
	UpdatedAt: condition.Field[overrideforeignkey.Person, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Person when doing a query
func (personConditions personConditions) Preload() condition.Condition[overrideforeignkey.Person] {
	return condition.NewPreloadCondition[overrideforeignkey.Person](personConditions.ID, personConditions.CreatedAt, personConditions.UpdatedAt, personConditions.DeletedAt)
}
