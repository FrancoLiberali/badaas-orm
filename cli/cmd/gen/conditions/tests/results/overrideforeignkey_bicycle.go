// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	overrideforeignkey "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/overrideforeignkey"
	cql "github.com/FrancoLiberali/cql/orm/cql"
	model "github.com/FrancoLiberali/cql/orm/model"
	"time"
)

func (bicycleConditions bicycleConditions) Owner(conditions ...cql.Condition[overrideforeignkey.Person]) cql.JoinCondition[overrideforeignkey.Bicycle] {
	return cql.NewJoinCondition[overrideforeignkey.Bicycle, overrideforeignkey.Person](conditions, "Owner", "OwnerSomethingID", bicycleConditions.Preload(), "ID")
}
func (bicycleConditions bicycleConditions) PreloadOwner() cql.JoinCondition[overrideforeignkey.Bicycle] {
	return bicycleConditions.Owner(Person.Preload())
}

type bicycleConditions struct {
	ID               cql.Field[overrideforeignkey.Bicycle, model.UUID]
	CreatedAt        cql.Field[overrideforeignkey.Bicycle, time.Time]
	UpdatedAt        cql.Field[overrideforeignkey.Bicycle, time.Time]
	DeletedAt        cql.Field[overrideforeignkey.Bicycle, time.Time]
	OwnerSomethingID cql.StringField[overrideforeignkey.Bicycle]
}

var Bicycle = bicycleConditions{
	CreatedAt:        cql.Field[overrideforeignkey.Bicycle, time.Time]{Name: "CreatedAt"},
	DeletedAt:        cql.Field[overrideforeignkey.Bicycle, time.Time]{Name: "DeletedAt"},
	ID:               cql.Field[overrideforeignkey.Bicycle, model.UUID]{Name: "ID"},
	OwnerSomethingID: cql.StringField[overrideforeignkey.Bicycle]{Field: cql.Field[overrideforeignkey.Bicycle, string]{Name: "OwnerSomethingID"}},
	UpdatedAt:        cql.Field[overrideforeignkey.Bicycle, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Bicycle when doing a query
func (bicycleConditions bicycleConditions) Preload() cql.Condition[overrideforeignkey.Bicycle] {
	return cql.NewPreloadCondition[overrideforeignkey.Bicycle](bicycleConditions.ID, bicycleConditions.CreatedAt, bicycleConditions.UpdatedAt, bicycleConditions.DeletedAt, bicycleConditions.OwnerSomethingID)
}

// PreloadRelations allows preloading all the Bicycle's relation when doing a query
func (bicycleConditions bicycleConditions) PreloadRelations() []cql.Condition[overrideforeignkey.Bicycle] {
	return []cql.Condition[overrideforeignkey.Bicycle]{bicycleConditions.PreloadOwner()}
}
