// Code generated by cql-gen v0.0.10, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	model "github.com/FrancoLiberali/cql/model"
	models "github.com/FrancoLiberali/cql/test/models"
	"time"
)

func (countryConditions countryConditions) Capital(conditions ...condition.Condition[models.City]) condition.JoinCondition[models.Country] {
	return condition.NewJoinCondition[models.Country, models.City](conditions, "Capital", "ID", countryConditions.preload(), "CountryID", City.preload())
}

type countryConditions struct {
	ID        condition.Field[models.Country, model.UUID]
	CreatedAt condition.Field[models.Country, time.Time]
	UpdatedAt condition.Field[models.Country, time.Time]
	DeletedAt condition.Field[models.Country, time.Time]
	Name      condition.StringField[models.Country]
}

var Country = countryConditions{
	CreatedAt: condition.NewField[models.Country, time.Time]("CreatedAt", "", ""),
	DeletedAt: condition.NewField[models.Country, time.Time]("DeletedAt", "", ""),
	ID:        condition.NewField[models.Country, model.UUID]("ID", "", ""),
	Name:      condition.NewStringField[models.Country]("Name", "", ""),
	UpdatedAt: condition.NewField[models.Country, time.Time]("UpdatedAt", "", ""),
}

// Preload allows preloading the Country when doing a query
func (countryConditions countryConditions) preload() condition.Condition[models.Country] {
	return condition.NewPreloadCondition[models.Country](countryConditions.ID, countryConditions.CreatedAt, countryConditions.UpdatedAt, countryConditions.DeletedAt, countryConditions.Name)
}
