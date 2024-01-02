// Code generated by cql-gen v0.0.10, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	overridereferences "github.com/FrancoLiberali/cql/cql-gen/cmd/gen/conditions/tests/overridereferences"
	model "github.com/FrancoLiberali/cql/model"
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
	CreatedAt: condition.NewField[overridereferences.Brand, time.Time]("CreatedAt", "", ""),
	DeletedAt: condition.NewField[overridereferences.Brand, time.Time]("DeletedAt", "", ""),
	ID:        condition.NewField[overridereferences.Brand, model.UUID]("ID", "", ""),
	Name:      condition.NewStringField[overridereferences.Brand]("Name", "", ""),
	UpdatedAt: condition.NewField[overridereferences.Brand, time.Time]("UpdatedAt", "", ""),
}

// Preload allows preloading the Brand when doing a query
func (brandConditions brandConditions) preload() condition.Condition[overridereferences.Brand] {
	return condition.NewPreloadCondition[overridereferences.Brand](brandConditions.ID, brandConditions.CreatedAt, brandConditions.UpdatedAt, brandConditions.DeletedAt, brandConditions.Name)
}
