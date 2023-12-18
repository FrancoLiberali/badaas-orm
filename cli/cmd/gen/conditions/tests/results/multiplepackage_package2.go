// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	package2 "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/multiplepackage/package2"
	cql "github.com/ditrit/badaas/orm/cql"
	model "github.com/ditrit/badaas/orm/model"
	"time"
)

type package2Conditions struct {
	ID         cql.Field[package2.Package2, model.UUID]
	CreatedAt  cql.Field[package2.Package2, time.Time]
	UpdatedAt  cql.Field[package2.Package2, time.Time]
	DeletedAt  cql.Field[package2.Package2, time.Time]
	Package1ID cql.Field[package2.Package2, model.UUID]
}

var Package2 = package2Conditions{
	CreatedAt:  cql.Field[package2.Package2, time.Time]{Name: "CreatedAt"},
	DeletedAt:  cql.Field[package2.Package2, time.Time]{Name: "DeletedAt"},
	ID:         cql.Field[package2.Package2, model.UUID]{Name: "ID"},
	Package1ID: cql.Field[package2.Package2, model.UUID]{Name: "Package1ID"},
	UpdatedAt:  cql.Field[package2.Package2, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Package2 when doing a query
func (package2Conditions package2Conditions) Preload() cql.Condition[package2.Package2] {
	return cql.NewPreloadCondition[package2.Package2](package2Conditions.ID, package2Conditions.CreatedAt, package2Conditions.UpdatedAt, package2Conditions.DeletedAt, package2Conditions.Package1ID)
}
