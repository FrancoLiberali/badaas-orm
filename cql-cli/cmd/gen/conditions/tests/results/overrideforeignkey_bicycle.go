// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	overrideforeignkey "github.com/FrancoLiberali/cql/cql-cli/cmd/gen/conditions/tests/overrideforeignkey"
	model "github.com/FrancoLiberali/cql/model"
	"time"
)

func (bicycleConditions bicycleConditions) Owner(conditions ...condition.Condition[overrideforeignkey.Person]) condition.JoinCondition[overrideforeignkey.Bicycle] {
	return condition.NewJoinCondition[overrideforeignkey.Bicycle, overrideforeignkey.Person](conditions, "Owner", "OwnerSomethingID", bicycleConditions.Preload(), "ID")
}
func (bicycleConditions bicycleConditions) PreloadOwner() condition.JoinCondition[overrideforeignkey.Bicycle] {
	return bicycleConditions.Owner(Person.Preload())
}

type bicycleConditions struct {
	ID               condition.Field[overrideforeignkey.Bicycle, model.UUID]
	CreatedAt        condition.Field[overrideforeignkey.Bicycle, time.Time]
	UpdatedAt        condition.Field[overrideforeignkey.Bicycle, time.Time]
	DeletedAt        condition.Field[overrideforeignkey.Bicycle, time.Time]
	OwnerSomethingID condition.StringField[overrideforeignkey.Bicycle]
}

var Bicycle = bicycleConditions{
	CreatedAt:        condition.Field[overrideforeignkey.Bicycle, time.Time]{Name: "CreatedAt"},
	DeletedAt:        condition.Field[overrideforeignkey.Bicycle, time.Time]{Name: "DeletedAt"},
	ID:               condition.Field[overrideforeignkey.Bicycle, model.UUID]{Name: "ID"},
	OwnerSomethingID: condition.StringField[overrideforeignkey.Bicycle]{Field: condition.Field[overrideforeignkey.Bicycle, string]{Name: "OwnerSomethingID"}},
	UpdatedAt:        condition.Field[overrideforeignkey.Bicycle, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Bicycle when doing a query
func (bicycleConditions bicycleConditions) Preload() condition.Condition[overrideforeignkey.Bicycle] {
	return condition.NewPreloadCondition[overrideforeignkey.Bicycle](bicycleConditions.ID, bicycleConditions.CreatedAt, bicycleConditions.UpdatedAt, bicycleConditions.DeletedAt, bicycleConditions.OwnerSomethingID)
}

// PreloadRelations allows preloading all the Bicycle's relation when doing a query
func (bicycleConditions bicycleConditions) PreloadRelations() []condition.Condition[overrideforeignkey.Bicycle] {
	return []condition.Condition[overrideforeignkey.Bicycle]{bicycleConditions.PreloadOwner()}
}
