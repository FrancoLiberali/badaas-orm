// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	cql "github.com/ditrit/badaas/orm/cql"
	model "github.com/ditrit/badaas/orm/model"
	models "github.com/ditrit/badaas/test/models"
	"time"
)

func (sellerConditions sellerConditions) Company(conditions ...cql.Condition[models.Company]) cql.JoinCondition[models.Seller] {
	return cql.NewJoinCondition[models.Seller, models.Company](conditions, "Company", "CompanyID", sellerConditions.Preload(), "ID")
}
func (sellerConditions sellerConditions) PreloadCompany() cql.JoinCondition[models.Seller] {
	return sellerConditions.Company(Company.Preload())
}
func (sellerConditions sellerConditions) University(conditions ...cql.Condition[models.University]) cql.JoinCondition[models.Seller] {
	return cql.NewJoinCondition[models.Seller, models.University](conditions, "University", "UniversityID", sellerConditions.Preload(), "ID")
}
func (sellerConditions sellerConditions) PreloadUniversity() cql.JoinCondition[models.Seller] {
	return sellerConditions.University(University.Preload())
}

type sellerConditions struct {
	ID           cql.Field[models.Seller, model.UUID]
	CreatedAt    cql.Field[models.Seller, time.Time]
	UpdatedAt    cql.Field[models.Seller, time.Time]
	DeletedAt    cql.Field[models.Seller, time.Time]
	Name         cql.StringField[models.Seller]
	CompanyID    cql.Field[models.Seller, model.UUID]
	UniversityID cql.Field[models.Seller, model.UUID]
}

var Seller = sellerConditions{
	CompanyID:    cql.Field[models.Seller, model.UUID]{Name: "CompanyID"},
	CreatedAt:    cql.Field[models.Seller, time.Time]{Name: "CreatedAt"},
	DeletedAt:    cql.Field[models.Seller, time.Time]{Name: "DeletedAt"},
	ID:           cql.Field[models.Seller, model.UUID]{Name: "ID"},
	Name:         cql.StringField[models.Seller]{Field: cql.Field[models.Seller, string]{Name: "Name"}},
	UniversityID: cql.Field[models.Seller, model.UUID]{Name: "UniversityID"},
	UpdatedAt:    cql.Field[models.Seller, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Seller when doing a query
func (sellerConditions sellerConditions) Preload() cql.Condition[models.Seller] {
	return cql.NewPreloadCondition[models.Seller](sellerConditions.ID, sellerConditions.CreatedAt, sellerConditions.UpdatedAt, sellerConditions.DeletedAt, sellerConditions.Name, sellerConditions.CompanyID, sellerConditions.UniversityID)
}

// PreloadRelations allows preloading all the Seller's relation when doing a query
func (sellerConditions sellerConditions) PreloadRelations() []cql.Condition[models.Seller] {
	return []cql.Condition[models.Seller]{sellerConditions.PreloadCompany(), sellerConditions.PreloadUniversity()}
}