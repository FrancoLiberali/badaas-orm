// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	orm "github.com/ditrit/badaas/orm"
	models "github.com/ditrit/badaas/testintegration/models"
	"time"
)

func CountryId(operator orm.Operator[orm.UUID]) orm.WhereCondition[models.Country] {
	return orm.FieldCondition[models.Country, orm.UUID]{
		Field:    "ID",
		Operator: operator,
	}
}
func CountryCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[models.Country] {
	return orm.FieldCondition[models.Country, time.Time]{
		Field:    "CreatedAt",
		Operator: operator,
	}
}
func CountryUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[models.Country] {
	return orm.FieldCondition[models.Country, time.Time]{
		Field:    "UpdatedAt",
		Operator: operator,
	}
}
func CountryDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[models.Country] {
	return orm.FieldCondition[models.Country, time.Time]{
		Field:    "DeletedAt",
		Operator: operator,
	}
}
func CountryName(operator orm.Operator[string]) orm.WhereCondition[models.Country] {
	return orm.FieldCondition[models.Country, string]{
		Field:    "Name",
		Operator: operator,
	}
}
func CountryCapital(conditions ...orm.Condition[models.City]) orm.Condition[models.Country] {
	return orm.JoinCondition[models.Country, models.City]{
		Conditions: conditions,
		T1Field:    "ID",
		T2Field:    "CountryID",
	}
}
func CityCountry(conditions ...orm.Condition[models.Country]) orm.Condition[models.City] {
	return orm.JoinCondition[models.City, models.Country]{
		Conditions: conditions,
		T1Field:    "CountryID",
		T2Field:    "ID",
	}
}