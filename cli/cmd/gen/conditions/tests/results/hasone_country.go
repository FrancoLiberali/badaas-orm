// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	hasone "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/hasone"
	cql "github.com/ditrit/badaas/orm/cql"
	model "github.com/ditrit/badaas/orm/model"
	"time"
)

func (countryConditions countryConditions) Capital(conditions ...cql.Condition[hasone.City]) cql.JoinCondition[hasone.Country] {
	return cql.NewJoinCondition[hasone.Country, hasone.City](conditions, "Capital", "ID", countryConditions.Preload(), "CountryID")
}
func (countryConditions countryConditions) PreloadCapital() cql.JoinCondition[hasone.Country] {
	return countryConditions.Capital(City.Preload())
}

type countryConditions struct {
	ID        cql.Field[hasone.Country, model.UUID]
	CreatedAt cql.Field[hasone.Country, time.Time]
	UpdatedAt cql.Field[hasone.Country, time.Time]
	DeletedAt cql.Field[hasone.Country, time.Time]
}

var Country = countryConditions{
	CreatedAt: cql.Field[hasone.Country, time.Time]{Name: "CreatedAt"},
	DeletedAt: cql.Field[hasone.Country, time.Time]{Name: "DeletedAt"},
	ID:        cql.Field[hasone.Country, model.UUID]{Name: "ID"},
	UpdatedAt: cql.Field[hasone.Country, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Country when doing a query
func (countryConditions countryConditions) Preload() cql.Condition[hasone.Country] {
	return cql.NewPreloadCondition[hasone.Country](countryConditions.ID, countryConditions.CreatedAt, countryConditions.UpdatedAt, countryConditions.DeletedAt)
}

// PreloadRelations allows preloading all the Country's relation when doing a query
func (countryConditions countryConditions) PreloadRelations() []cql.Condition[hasone.Country] {
	return []cql.Condition[hasone.Country]{countryConditions.PreloadCapital()}
}
