// Code generated by cql-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	hasmanywithpointers "github.com/FrancoLiberali/cql/cql-cli/cmd/gen/conditions/tests/hasmanywithpointers"
	model "github.com/FrancoLiberali/cql/model"
	"time"
)

func (companyWithPointersConditions companyWithPointersConditions) PreloadSellers(nestedPreloads ...condition.JoinCondition[hasmanywithpointers.SellerInPointers]) condition.Condition[hasmanywithpointers.CompanyWithPointers] {
	return condition.NewCollectionPreloadCondition[hasmanywithpointers.CompanyWithPointers, hasmanywithpointers.SellerInPointers]("Sellers", nestedPreloads)
}

type companyWithPointersConditions struct {
	ID        condition.Field[hasmanywithpointers.CompanyWithPointers, model.UUID]
	CreatedAt condition.Field[hasmanywithpointers.CompanyWithPointers, time.Time]
	UpdatedAt condition.Field[hasmanywithpointers.CompanyWithPointers, time.Time]
	DeletedAt condition.Field[hasmanywithpointers.CompanyWithPointers, time.Time]
}

var CompanyWithPointers = companyWithPointersConditions{
	CreatedAt: condition.Field[hasmanywithpointers.CompanyWithPointers, time.Time]{Name: "CreatedAt"},
	DeletedAt: condition.Field[hasmanywithpointers.CompanyWithPointers, time.Time]{Name: "DeletedAt"},
	ID:        condition.Field[hasmanywithpointers.CompanyWithPointers, model.UUID]{Name: "ID"},
	UpdatedAt: condition.Field[hasmanywithpointers.CompanyWithPointers, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the CompanyWithPointers when doing a query
func (companyWithPointersConditions companyWithPointersConditions) Preload() condition.Condition[hasmanywithpointers.CompanyWithPointers] {
	return condition.NewPreloadCondition[hasmanywithpointers.CompanyWithPointers](companyWithPointersConditions.ID, companyWithPointersConditions.CreatedAt, companyWithPointersConditions.UpdatedAt, companyWithPointersConditions.DeletedAt)
}

// PreloadRelations allows preloading all the CompanyWithPointers's relation when doing a query
func (companyWithPointersConditions companyWithPointersConditions) PreloadRelations() []condition.Condition[hasmanywithpointers.CompanyWithPointers] {
	return []condition.Condition[hasmanywithpointers.CompanyWithPointers]{companyWithPointersConditions.PreloadSellers()}
}
