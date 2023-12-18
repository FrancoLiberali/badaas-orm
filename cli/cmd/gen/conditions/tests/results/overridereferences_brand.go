// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	model "github.com/FrancoLiberali/cql/model"
	overridereferences "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/overridereferences"
	"time"
)

type brandConditions struct {
	ID        condition.Field[overridereferences.Brand, model.UUID]
	CreatedAt condition.Field[overridereferences.Brand, time.Time]
	UpdatedAt condition.Field[overridereferences.Brand, time.Time]
	DeletedAt condition.Field[overridereferences.Brand, time.Time]
	Name      condition.StringField[overridereferences.Brand]
}

var Brand = brandConditions{
	CreatedAt: condition.Field[overridereferences.Brand, time.Time]{Name: "CreatedAt"},
	DeletedAt: condition.Field[overridereferences.Brand, time.Time]{Name: "DeletedAt"},
	ID:        condition.Field[overridereferences.Brand, model.UUID]{Name: "ID"},
	Name:      condition.StringField[overridereferences.Brand]{Field: condition.Field[overridereferences.Brand, string]{Name: "Name"}},
	UpdatedAt: condition.Field[overridereferences.Brand, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Brand when doing a query
func (brandConditions brandConditions) Preload() condition.Condition[overridereferences.Brand] {
	return condition.NewPreloadCondition[overridereferences.Brand](brandConditions.ID, brandConditions.CreatedAt, brandConditions.UpdatedAt, brandConditions.DeletedAt, brandConditions.Name)
}
