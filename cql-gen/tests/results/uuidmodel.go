// Code generated by cql-gen v0.0.8, DO NOT EDIT.
package conditions

import (
	uuidmodel "github.com/FrancoLiberali/cql-gen/cmd/gen/conditions/tests/uuidmodel"
	condition "github.com/FrancoLiberali/cql/condition"
	model "github.com/FrancoLiberali/cql/model"
	"time"
)

type uuidModelConditions struct {
	ID        condition.Field[uuidmodel.UUIDModel, model.UUID]
	CreatedAt condition.Field[uuidmodel.UUIDModel, time.Time]
	UpdatedAt condition.Field[uuidmodel.UUIDModel, time.Time]
	DeletedAt condition.Field[uuidmodel.UUIDModel, time.Time]
}

var UUIDModel = uuidModelConditions{
	CreatedAt: condition.NewField[uuidmodel.UUIDModel, time.Time]("CreatedAt", "", ""),
	DeletedAt: condition.NewField[uuidmodel.UUIDModel, time.Time]("DeletedAt", "", ""),
	ID:        condition.NewField[uuidmodel.UUIDModel, model.UUID]("ID", "", ""),
	UpdatedAt: condition.NewField[uuidmodel.UUIDModel, time.Time]("UpdatedAt", "", ""),
}

// Preload allows preloading the UUIDModel when doing a query
func (uuidModelConditions uuidModelConditions) preload() condition.Condition[uuidmodel.UUIDModel] {
	return condition.NewPreloadCondition[uuidmodel.UUIDModel](uuidModelConditions.ID, uuidModelConditions.CreatedAt, uuidModelConditions.UpdatedAt, uuidModelConditions.DeletedAt)
}
