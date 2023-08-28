// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	orm "github.com/ditrit/badaas/orm"
	condition "github.com/ditrit/badaas/orm/condition"
	model "github.com/ditrit/badaas/orm/model"
	query "github.com/ditrit/badaas/orm/query"
	models "github.com/ditrit/badaas/testintegration/models"
	"reflect"
	"time"
)

var sellerType = reflect.TypeOf(*new(models.Seller))

func (sellerConditions sellerConditions) IdIs() orm.FieldIs[models.Seller, model.UUID] {
	return orm.FieldIs[models.Seller, model.UUID]{FieldID: sellerConditions.ID}
}
func (sellerConditions sellerConditions) CreatedAtIs() orm.FieldIs[models.Seller, time.Time] {
	return orm.FieldIs[models.Seller, time.Time]{FieldID: sellerConditions.CreatedAt}
}
func (sellerConditions sellerConditions) UpdatedAtIs() orm.FieldIs[models.Seller, time.Time] {
	return orm.FieldIs[models.Seller, time.Time]{FieldID: sellerConditions.UpdatedAt}
}
func (sellerConditions sellerConditions) DeletedAtIs() orm.FieldIs[models.Seller, time.Time] {
	return orm.FieldIs[models.Seller, time.Time]{FieldID: sellerConditions.DeletedAt}
}
func (sellerConditions sellerConditions) NameIs() orm.StringFieldIs[models.Seller] {
	return orm.StringFieldIs[models.Seller]{FieldIs: orm.FieldIs[models.Seller, string]{FieldID: sellerConditions.Name}}
}
func (sellerConditions sellerConditions) Company(conditions ...condition.Condition[models.Company]) condition.JoinCondition[models.Seller] {
	return condition.NewJoinCondition[models.Seller, models.Company](conditions, "Company", "CompanyID", sellerConditions.Preload(), "ID")
}
func (sellerConditions sellerConditions) PreloadCompany() condition.JoinCondition[models.Seller] {
	return sellerConditions.Company(Company.Preload())
}
func (sellerConditions sellerConditions) CompanyIdIs() orm.FieldIs[models.Seller, model.UUID] {
	return orm.FieldIs[models.Seller, model.UUID]{FieldID: sellerConditions.CompanyID}
}
func (sellerConditions sellerConditions) University(conditions ...condition.Condition[models.University]) condition.JoinCondition[models.Seller] {
	return condition.NewJoinCondition[models.Seller, models.University](conditions, "University", "UniversityID", sellerConditions.Preload(), "ID")
}
func (sellerConditions sellerConditions) PreloadUniversity() condition.JoinCondition[models.Seller] {
	return sellerConditions.University(University.Preload())
}
func (sellerConditions sellerConditions) UniversityIdIs() orm.FieldIs[models.Seller, model.UUID] {
	return orm.FieldIs[models.Seller, model.UUID]{FieldID: sellerConditions.UniversityID}
}

type sellerConditions struct {
	ID           query.FieldIdentifier[model.UUID]
	CreatedAt    query.FieldIdentifier[time.Time]
	UpdatedAt    query.FieldIdentifier[time.Time]
	DeletedAt    query.FieldIdentifier[time.Time]
	Name         query.FieldIdentifier[string]
	CompanyID    query.FieldIdentifier[model.UUID]
	UniversityID query.FieldIdentifier[model.UUID]
}

var Seller = sellerConditions{
	CompanyID: query.FieldIdentifier[model.UUID]{
		Field:     "CompanyID",
		ModelType: sellerType,
	},
	CreatedAt: query.FieldIdentifier[time.Time]{
		Field:     "CreatedAt",
		ModelType: sellerType,
	},
	DeletedAt: query.FieldIdentifier[time.Time]{
		Field:     "DeletedAt",
		ModelType: sellerType,
	},
	ID: query.FieldIdentifier[model.UUID]{
		Field:     "ID",
		ModelType: sellerType,
	},
	Name: query.FieldIdentifier[string]{
		Field:     "Name",
		ModelType: sellerType,
	},
	UniversityID: query.FieldIdentifier[model.UUID]{
		Field:     "UniversityID",
		ModelType: sellerType,
	},
	UpdatedAt: query.FieldIdentifier[time.Time]{
		Field:     "UpdatedAt",
		ModelType: sellerType,
	},
}

// Preload allows preloading the Seller when doing a query
func (sellerConditions sellerConditions) Preload() condition.Condition[models.Seller] {
	return condition.NewPreloadCondition[models.Seller](sellerConditions.ID, sellerConditions.CreatedAt, sellerConditions.UpdatedAt, sellerConditions.DeletedAt, sellerConditions.Name, sellerConditions.CompanyID, sellerConditions.UniversityID)
}

// PreloadRelations allows preloading all the Seller's relation when doing a query
func (sellerConditions sellerConditions) PreloadRelations() []condition.Condition[models.Seller] {
	return []condition.Condition[models.Seller]{sellerConditions.PreloadCompany(), sellerConditions.PreloadUniversity()}
}
