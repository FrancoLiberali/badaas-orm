// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	orm "github.com/ditrit/badaas/orm"
	model "github.com/ditrit/badaas/orm/model"
	models "github.com/ditrit/badaas/testintegration/models"
	"time"
)

func (cityConditions cityConditions) Country(conditions ...orm.Condition[models.Country]) orm.JoinCondition[models.City] {
	return orm.NewJoinCondition[models.City, models.Country](conditions, "Country", "CountryID", cityConditions.Preload(), "ID")
}
func (cityConditions cityConditions) PreloadCountry() orm.JoinCondition[models.City] {
	return cityConditions.Country(Country.Preload())
}

type cityConditions struct {
	ID        orm.Field[models.City, model.UUID]
	CreatedAt orm.Field[models.City, time.Time]
	UpdatedAt orm.Field[models.City, time.Time]
	DeletedAt orm.Field[models.City, time.Time]
	Name      orm.StringField[models.City]
	CountryID orm.Field[models.City, model.UUID]
}

var City = cityConditions{
	CountryID: orm.Field[models.City, model.UUID]{Name: "CountryID"},
	CreatedAt: orm.Field[models.City, time.Time]{Name: "CreatedAt"},
	DeletedAt: orm.Field[models.City, time.Time]{Name: "DeletedAt"},
	ID:        orm.Field[models.City, model.UUID]{Name: "ID"},
	Name:      orm.StringField[models.City]{Field: orm.Field[models.City, string]{Name: "Name"}},
	UpdatedAt: orm.Field[models.City, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the City when doing a query
func (cityConditions cityConditions) Preload() orm.Condition[models.City] {
	return orm.NewPreloadCondition[models.City](cityConditions.ID, cityConditions.CreatedAt, cityConditions.UpdatedAt, cityConditions.DeletedAt, cityConditions.Name, cityConditions.CountryID)
}

// PreloadRelations allows preloading all the City's relation when doing a query
func (cityConditions cityConditions) PreloadRelations() []orm.Condition[models.City] {
	return []orm.Condition[models.City]{cityConditions.PreloadCountry()}
}
