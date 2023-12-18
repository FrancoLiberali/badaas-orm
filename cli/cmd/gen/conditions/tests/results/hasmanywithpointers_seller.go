// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	hasmanywithpointers "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/hasmanywithpointers"
	cql "github.com/ditrit/badaas/orm/cql"
	model "github.com/ditrit/badaas/orm/model"
	"time"
)

func (sellerInPointersConditions sellerInPointersConditions) Company(conditions ...cql.Condition[hasmanywithpointers.CompanyWithPointers]) cql.JoinCondition[hasmanywithpointers.SellerInPointers] {
	return cql.NewJoinCondition[hasmanywithpointers.SellerInPointers, hasmanywithpointers.CompanyWithPointers](conditions, "Company", "CompanyID", sellerInPointersConditions.Preload(), "ID")
}
func (sellerInPointersConditions sellerInPointersConditions) PreloadCompany() cql.JoinCondition[hasmanywithpointers.SellerInPointers] {
	return sellerInPointersConditions.Company(CompanyWithPointers.Preload())
}

type sellerInPointersConditions struct {
	ID        cql.Field[hasmanywithpointers.SellerInPointers, model.UUID]
	CreatedAt cql.Field[hasmanywithpointers.SellerInPointers, time.Time]
	UpdatedAt cql.Field[hasmanywithpointers.SellerInPointers, time.Time]
	DeletedAt cql.Field[hasmanywithpointers.SellerInPointers, time.Time]
	CompanyID cql.Field[hasmanywithpointers.SellerInPointers, model.UUID]
}

var SellerInPointers = sellerInPointersConditions{
	CompanyID: cql.Field[hasmanywithpointers.SellerInPointers, model.UUID]{Name: "CompanyID"},
	CreatedAt: cql.Field[hasmanywithpointers.SellerInPointers, time.Time]{Name: "CreatedAt"},
	DeletedAt: cql.Field[hasmanywithpointers.SellerInPointers, time.Time]{Name: "DeletedAt"},
	ID:        cql.Field[hasmanywithpointers.SellerInPointers, model.UUID]{Name: "ID"},
	UpdatedAt: cql.Field[hasmanywithpointers.SellerInPointers, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the SellerInPointers when doing a query
func (sellerInPointersConditions sellerInPointersConditions) Preload() cql.Condition[hasmanywithpointers.SellerInPointers] {
	return cql.NewPreloadCondition[hasmanywithpointers.SellerInPointers](sellerInPointersConditions.ID, sellerInPointersConditions.CreatedAt, sellerInPointersConditions.UpdatedAt, sellerInPointersConditions.DeletedAt, sellerInPointersConditions.CompanyID)
}

// PreloadRelations allows preloading all the SellerInPointers's relation when doing a query
func (sellerInPointersConditions sellerInPointersConditions) PreloadRelations() []cql.Condition[hasmanywithpointers.SellerInPointers] {
	return []cql.Condition[hasmanywithpointers.SellerInPointers]{sellerInPointersConditions.PreloadCompany()}
}
