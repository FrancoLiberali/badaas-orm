// Code generated by cql-gen v0.0.8, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	nullabletypes "github.com/FrancoLiberali/cql-gen/cmd/gen/conditions/tests/nullabletypes"
	model "github.com/FrancoLiberali/cql/model"
	"time"
)

type nullableTypesConditions struct {
	ID        condition.Field[nullabletypes.NullableTypes, model.UUID]
	CreatedAt condition.Field[nullabletypes.NullableTypes, time.Time]
	UpdatedAt condition.Field[nullabletypes.NullableTypes, time.Time]
	DeletedAt condition.Field[nullabletypes.NullableTypes, time.Time]
	String    condition.NullableStringField[nullabletypes.NullableTypes]
	Int64     condition.NullableField[nullabletypes.NullableTypes, int64]
	Int32     condition.NullableField[nullabletypes.NullableTypes, int32]
	Int16     condition.NullableField[nullabletypes.NullableTypes, int16]
	Byte      condition.NullableField[nullabletypes.NullableTypes, int8]
	Float64   condition.NullableField[nullabletypes.NullableTypes, float64]
	Bool      condition.NullableBoolField[nullabletypes.NullableTypes]
	Time      condition.NullableField[nullabletypes.NullableTypes, time.Time]
}

var NullableTypes = nullableTypesConditions{
	Bool:      condition.NewNullableBoolField[nullabletypes.NullableTypes]("Bool", "", ""),
	Byte:      condition.NewNullableField[nullabletypes.NullableTypes, int8]("Byte", "", ""),
	CreatedAt: condition.NewField[nullabletypes.NullableTypes, time.Time]("CreatedAt", "", ""),
	DeletedAt: condition.NewField[nullabletypes.NullableTypes, time.Time]("DeletedAt", "", ""),
	Float64:   condition.NewNullableField[nullabletypes.NullableTypes, float64]("Float64", "", ""),
	ID:        condition.NewField[nullabletypes.NullableTypes, model.UUID]("ID", "", ""),
	Int16:     condition.NewNullableField[nullabletypes.NullableTypes, int16]("Int16", "", ""),
	Int32:     condition.NewNullableField[nullabletypes.NullableTypes, int32]("Int32", "", ""),
	Int64:     condition.NewNullableField[nullabletypes.NullableTypes, int64]("Int64", "", ""),
	String:    condition.NewNullableStringField[nullabletypes.NullableTypes]("String", "", ""),
	Time:      condition.NewNullableField[nullabletypes.NullableTypes, time.Time]("Time", "", ""),
	UpdatedAt: condition.NewField[nullabletypes.NullableTypes, time.Time]("UpdatedAt", "", ""),
}

// Preload allows preloading the NullableTypes when doing a query
func (nullableTypesConditions nullableTypesConditions) preload() condition.Condition[nullabletypes.NullableTypes] {
	return condition.NewPreloadCondition[nullabletypes.NullableTypes](nullableTypesConditions.ID, nullableTypesConditions.CreatedAt, nullableTypesConditions.UpdatedAt, nullableTypesConditions.DeletedAt, nullableTypesConditions.String, nullableTypesConditions.Int64, nullableTypesConditions.Int32, nullableTypesConditions.Int16, nullableTypesConditions.Byte, nullableTypesConditions.Float64, nullableTypesConditions.Bool, nullableTypesConditions.Time)
}
