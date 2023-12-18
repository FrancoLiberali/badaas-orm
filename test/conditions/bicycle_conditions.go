// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	model "github.com/FrancoLiberali/cql/model"
	models "github.com/FrancoLiberali/cql/test/models"
	"time"
)

func (bicycleConditions bicycleConditions) Owner(conditions ...condition.Condition[models.Person]) condition.JoinCondition[models.Bicycle] {
	return condition.NewJoinCondition[models.Bicycle, models.Person](conditions, "Owner", "OwnerName", bicycleConditions.Preload(), "Name")
}
func (bicycleConditions bicycleConditions) PreloadOwner() condition.JoinCondition[models.Bicycle] {
	return bicycleConditions.Owner(Person.Preload())
}

type bicycleConditions struct {
	ID        condition.Field[models.Bicycle, model.UUID]
	CreatedAt condition.Field[models.Bicycle, time.Time]
	UpdatedAt condition.Field[models.Bicycle, time.Time]
	DeletedAt condition.Field[models.Bicycle, time.Time]
	Name      condition.StringField[models.Bicycle]
	OwnerName condition.StringField[models.Bicycle]
}

var Bicycle = bicycleConditions{
	CreatedAt: condition.Field[models.Bicycle, time.Time]{Name: "CreatedAt"},
	DeletedAt: condition.Field[models.Bicycle, time.Time]{Name: "DeletedAt"},
	ID:        condition.Field[models.Bicycle, model.UUID]{Name: "ID"},
	Name:      condition.StringField[models.Bicycle]{Field: condition.Field[models.Bicycle, string]{Name: "Name"}},
	OwnerName: condition.StringField[models.Bicycle]{Field: condition.Field[models.Bicycle, string]{Name: "OwnerName"}},
	UpdatedAt: condition.Field[models.Bicycle, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Bicycle when doing a query
func (bicycleConditions bicycleConditions) Preload() condition.Condition[models.Bicycle] {
	return condition.NewPreloadCondition[models.Bicycle](bicycleConditions.ID, bicycleConditions.CreatedAt, bicycleConditions.UpdatedAt, bicycleConditions.DeletedAt, bicycleConditions.Name, bicycleConditions.OwnerName)
}

// PreloadRelations allows preloading all the Bicycle's relation when doing a query
func (bicycleConditions bicycleConditions) PreloadRelations() []condition.Condition[models.Bicycle] {
	return []condition.Condition[models.Bicycle]{bicycleConditions.PreloadOwner()}
}
