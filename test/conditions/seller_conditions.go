// Code generated by cql-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	model "github.com/FrancoLiberali/cql/model"
	models "github.com/FrancoLiberali/cql/test/models"
	"time"
)

func (sellerConditions sellerConditions) Company(conditions ...condition.Condition[models.Company]) condition.JoinCondition[models.Seller] {
	return condition.NewJoinCondition[models.Seller, models.Company](conditions, "Company", "CompanyID", sellerConditions.Preload(), "ID")
}
func (sellerConditions sellerConditions) PreloadCompany() condition.JoinCondition[models.Seller] {
	return sellerConditions.Company(Company.Preload())
}
func (sellerConditions sellerConditions) University(conditions ...condition.Condition[models.University]) condition.JoinCondition[models.Seller] {
	return condition.NewJoinCondition[models.Seller, models.University](conditions, "University", "UniversityID", sellerConditions.Preload(), "ID")
}
func (sellerConditions sellerConditions) PreloadUniversity() condition.JoinCondition[models.Seller] {
	return sellerConditions.University(University.Preload())
}

type sellerConditions struct {
	ID           condition.Field[models.Seller, model.UUID]
	CreatedAt    condition.Field[models.Seller, time.Time]
	UpdatedAt    condition.Field[models.Seller, time.Time]
	DeletedAt    condition.Field[models.Seller, time.Time]
	Name         condition.StringField[models.Seller]
	CompanyID    condition.Field[models.Seller, model.UUID]
	UniversityID condition.Field[models.Seller, model.UUID]
}

var Seller = sellerConditions{
	CompanyID:    condition.Field[models.Seller, model.UUID]{Name: "CompanyID"},
	CreatedAt:    condition.Field[models.Seller, time.Time]{Name: "CreatedAt"},
	DeletedAt:    condition.Field[models.Seller, time.Time]{Name: "DeletedAt"},
	ID:           condition.Field[models.Seller, model.UUID]{Name: "ID"},
	Name:         condition.StringField[models.Seller]{Field: condition.Field[models.Seller, string]{Name: "Name"}},
	UniversityID: condition.Field[models.Seller, model.UUID]{Name: "UniversityID"},
	UpdatedAt:    condition.Field[models.Seller, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Seller when doing a query
func (sellerConditions sellerConditions) Preload() condition.Condition[models.Seller] {
	return condition.NewPreloadCondition[models.Seller](sellerConditions.ID, sellerConditions.CreatedAt, sellerConditions.UpdatedAt, sellerConditions.DeletedAt, sellerConditions.Name, sellerConditions.CompanyID, sellerConditions.UniversityID)
}

// PreloadRelations allows preloading all the Seller's relation when doing a query
func (sellerConditions sellerConditions) PreloadRelations() []condition.Condition[models.Seller] {
	return []condition.Condition[models.Seller]{sellerConditions.PreloadCompany(), sellerConditions.PreloadUniversity()}
}
