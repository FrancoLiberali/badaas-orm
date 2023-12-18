// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	hasmany "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/hasmany"
	cql "github.com/FrancoLiberali/cql/orm/cql"
	model "github.com/FrancoLiberali/cql/orm/model"
	"time"
)

func (companyConditions companyConditions) PreloadSellers(nestedPreloads ...cql.JoinCondition[hasmany.Seller]) cql.Condition[hasmany.Company] {
	return cql.NewCollectionPreloadCondition[hasmany.Company, hasmany.Seller]("Sellers", nestedPreloads)
}

type companyConditions struct {
	ID        cql.Field[hasmany.Company, model.UUID]
	CreatedAt cql.Field[hasmany.Company, time.Time]
	UpdatedAt cql.Field[hasmany.Company, time.Time]
	DeletedAt cql.Field[hasmany.Company, time.Time]
}

var Company = companyConditions{
	CreatedAt: cql.Field[hasmany.Company, time.Time]{Name: "CreatedAt"},
	DeletedAt: cql.Field[hasmany.Company, time.Time]{Name: "DeletedAt"},
	ID:        cql.Field[hasmany.Company, model.UUID]{Name: "ID"},
	UpdatedAt: cql.Field[hasmany.Company, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Company when doing a query
func (companyConditions companyConditions) Preload() cql.Condition[hasmany.Company] {
	return cql.NewPreloadCondition[hasmany.Company](companyConditions.ID, companyConditions.CreatedAt, companyConditions.UpdatedAt, companyConditions.DeletedAt)
}

// PreloadRelations allows preloading all the Company's relation when doing a query
func (companyConditions companyConditions) PreloadRelations() []cql.Condition[hasmany.Company] {
	return []cql.Condition[hasmany.Company]{companyConditions.PreloadSellers()}
}
