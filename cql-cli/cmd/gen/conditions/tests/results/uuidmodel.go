// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	uuidmodel "github.com/FrancoLiberali/cql/cql-cli/cmd/gen/conditions/tests/uuidmodel"
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
	CreatedAt: condition.Field[uuidmodel.UUIDModel, time.Time]{Name: "CreatedAt"},
	DeletedAt: condition.Field[uuidmodel.UUIDModel, time.Time]{Name: "DeletedAt"},
	ID:        condition.Field[uuidmodel.UUIDModel, model.UUID]{Name: "ID"},
	UpdatedAt: condition.Field[uuidmodel.UUIDModel, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the UUIDModel when doing a query
func (uuidModelConditions uuidModelConditions) Preload() condition.Condition[uuidmodel.UUIDModel] {
	return condition.NewPreloadCondition[uuidmodel.UUIDModel](uuidModelConditions.ID, uuidModelConditions.CreatedAt, uuidModelConditions.UpdatedAt, uuidModelConditions.DeletedAt)
}
