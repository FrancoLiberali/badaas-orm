// Code generated by cql-gen v0.0.8, DO NOT EDIT.
package conditions

import (
	hasmany "github.com/FrancoLiberali/cql-gen/cmd/gen/conditions/tests/hasmany"
	condition "github.com/FrancoLiberali/cql/condition"
	model "github.com/FrancoLiberali/cql/model"
	"time"
)

type companyConditions struct {
	ID        condition.Field[hasmany.Company, model.UUID]
	CreatedAt condition.Field[hasmany.Company, time.Time]
	UpdatedAt condition.Field[hasmany.Company, time.Time]
	DeletedAt condition.Field[hasmany.Company, time.Time]
	Sellers   condition.Collection[hasmany.Company, hasmany.Seller]
}

var Company = companyConditions{
	CreatedAt: condition.NewField[hasmany.Company, time.Time]("CreatedAt", "", ""),
	DeletedAt: condition.NewField[hasmany.Company, time.Time]("DeletedAt", "", ""),
	ID:        condition.NewField[hasmany.Company, model.UUID]("ID", "", ""),
	Sellers:   condition.NewCollection[hasmany.Company, hasmany.Seller]("Sellers", "ID", "CompanyID"),
	UpdatedAt: condition.NewField[hasmany.Company, time.Time]("UpdatedAt", "", ""),
}

// Preload allows preloading the Company when doing a query
func (companyConditions companyConditions) preload() condition.Condition[hasmany.Company] {
	return condition.NewPreloadCondition[hasmany.Company](companyConditions.ID, companyConditions.CreatedAt, companyConditions.UpdatedAt, companyConditions.DeletedAt)
}
