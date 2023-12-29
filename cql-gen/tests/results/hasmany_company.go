// Code generated by cql-gen v0.0.6, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	hasmany "github.com/FrancoLiberali/cql/cql-gen/cmd/gen/conditions/tests/hasmany"
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
	CreatedAt: condition.Field[hasmany.Company, time.Time]{Name: "CreatedAt"},
	DeletedAt: condition.Field[hasmany.Company, time.Time]{Name: "DeletedAt"},
	ID:        condition.Field[hasmany.Company, model.UUID]{Name: "ID"},
	Sellers:   condition.NewCollection[hasmany.Company, hasmany.Seller]("Sellers"),
	UpdatedAt: condition.Field[hasmany.Company, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Company when doing a query
func (companyConditions companyConditions) preload() condition.Condition[hasmany.Company] {
	return condition.NewPreloadCondition[hasmany.Company](companyConditions.ID, companyConditions.CreatedAt, companyConditions.UpdatedAt, companyConditions.DeletedAt)
}
