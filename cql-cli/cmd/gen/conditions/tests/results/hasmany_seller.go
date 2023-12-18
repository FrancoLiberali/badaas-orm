// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	hasmany "github.com/FrancoLiberali/cql/cql-cli/cmd/gen/conditions/tests/hasmany"
	model "github.com/FrancoLiberali/cql/model"
	"time"
)

func (sellerConditions sellerConditions) Company(conditions ...condition.Condition[hasmany.Company]) condition.JoinCondition[hasmany.Seller] {
	return condition.NewJoinCondition[hasmany.Seller, hasmany.Company](conditions, "Company", "CompanyID", sellerConditions.Preload(), "ID")
}
func (sellerConditions sellerConditions) PreloadCompany() condition.JoinCondition[hasmany.Seller] {
	return sellerConditions.Company(Company.Preload())
}

type sellerConditions struct {
	ID        condition.Field[hasmany.Seller, model.UUID]
	CreatedAt condition.Field[hasmany.Seller, time.Time]
	UpdatedAt condition.Field[hasmany.Seller, time.Time]
	DeletedAt condition.Field[hasmany.Seller, time.Time]
	CompanyID condition.Field[hasmany.Seller, model.UUID]
}

var Seller = sellerConditions{
	CompanyID: condition.Field[hasmany.Seller, model.UUID]{Name: "CompanyID"},
	CreatedAt: condition.Field[hasmany.Seller, time.Time]{Name: "CreatedAt"},
	DeletedAt: condition.Field[hasmany.Seller, time.Time]{Name: "DeletedAt"},
	ID:        condition.Field[hasmany.Seller, model.UUID]{Name: "ID"},
	UpdatedAt: condition.Field[hasmany.Seller, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Seller when doing a query
func (sellerConditions sellerConditions) Preload() condition.Condition[hasmany.Seller] {
	return condition.NewPreloadCondition[hasmany.Seller](sellerConditions.ID, sellerConditions.CreatedAt, sellerConditions.UpdatedAt, sellerConditions.DeletedAt, sellerConditions.CompanyID)
}

// PreloadRelations allows preloading all the Seller's relation when doing a query
func (sellerConditions sellerConditions) PreloadRelations() []condition.Condition[hasmany.Seller] {
	return []condition.Condition[hasmany.Seller]{sellerConditions.PreloadCompany()}
}
