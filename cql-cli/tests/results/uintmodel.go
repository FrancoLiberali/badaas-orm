// Code generated by cql-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	uintmodel "github.com/FrancoLiberali/cql/cql-cli/cmd/gen/conditions/tests/uintmodel"
	model "github.com/FrancoLiberali/cql/model"
	"time"
)

type uintModelConditions struct {
	ID        condition.Field[uintmodel.UintModel, model.UIntID]
	CreatedAt condition.Field[uintmodel.UintModel, time.Time]
	UpdatedAt condition.Field[uintmodel.UintModel, time.Time]
	DeletedAt condition.Field[uintmodel.UintModel, time.Time]
}

var UintModel = uintModelConditions{
	CreatedAt: condition.Field[uintmodel.UintModel, time.Time]{Name: "CreatedAt"},
	DeletedAt: condition.Field[uintmodel.UintModel, time.Time]{Name: "DeletedAt"},
	ID:        condition.Field[uintmodel.UintModel, model.UIntID]{Name: "ID"},
	UpdatedAt: condition.Field[uintmodel.UintModel, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the UintModel when doing a query
func (uintModelConditions uintModelConditions) Preload() condition.Condition[uintmodel.UintModel] {
	return condition.NewPreloadCondition[uintmodel.UintModel](uintModelConditions.ID, uintModelConditions.CreatedAt, uintModelConditions.UpdatedAt, uintModelConditions.DeletedAt)
}