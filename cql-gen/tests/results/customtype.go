// Code generated by cql-gen v0.1.0, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	customtype "github.com/FrancoLiberali/cql/cql-gen/cmd/gen/conditions/tests/customtype"
	model "github.com/FrancoLiberali/cql/model"
	"time"
)

type customTypeConditions struct {
	ID        condition.Field[customtype.CustomType, model.UUID]
	CreatedAt condition.Field[customtype.CustomType, time.Time]
	UpdatedAt condition.Field[customtype.CustomType, time.Time]
	DeletedAt condition.Field[customtype.CustomType, time.Time]
	Custom    condition.UpdatableField[customtype.CustomType, customtype.MultiString]
}

var CustomType = customTypeConditions{
	CreatedAt: condition.NewField[customtype.CustomType, time.Time]("CreatedAt", "", ""),
	Custom:    condition.NewUpdatableField[customtype.CustomType, customtype.MultiString]("Custom", "", ""),
	DeletedAt: condition.NewField[customtype.CustomType, time.Time]("DeletedAt", "", ""),
	ID:        condition.NewField[customtype.CustomType, model.UUID]("ID", "", ""),
	UpdatedAt: condition.NewField[customtype.CustomType, time.Time]("UpdatedAt", "", ""),
}

// Preload allows preloading the CustomType when doing a query
func (customTypeConditions customTypeConditions) preload() condition.Condition[customtype.CustomType] {
	return condition.NewPreloadCondition[customtype.CustomType](customTypeConditions.ID, customTypeConditions.CreatedAt, customTypeConditions.UpdatedAt, customTypeConditions.DeletedAt, customTypeConditions.Custom)
}
