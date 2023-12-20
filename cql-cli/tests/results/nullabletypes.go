// Code generated by cql-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	nullabletypes "github.com/FrancoLiberali/cql/cql-cli/cmd/gen/conditions/tests/nullabletypes"
	model "github.com/FrancoLiberali/cql/model"
	"time"
)

type nullableTypesConditions struct {
	ID        condition.Field[nullabletypes.NullableTypes, model.UUID]
	CreatedAt condition.Field[nullabletypes.NullableTypes, time.Time]
	UpdatedAt condition.Field[nullabletypes.NullableTypes, time.Time]
	DeletedAt condition.Field[nullabletypes.NullableTypes, time.Time]
	String    condition.StringField[nullabletypes.NullableTypes]
	Int64     condition.UpdatableField[nullabletypes.NullableTypes, int64]
	Int32     condition.UpdatableField[nullabletypes.NullableTypes, int32]
	Int16     condition.UpdatableField[nullabletypes.NullableTypes, int16]
	Byte      condition.UpdatableField[nullabletypes.NullableTypes, int8]
	Float64   condition.UpdatableField[nullabletypes.NullableTypes, float64]
	Bool      condition.BoolField[nullabletypes.NullableTypes]
	Time      condition.UpdatableField[nullabletypes.NullableTypes, time.Time]
}

var NullableTypes = nullableTypesConditions{
	Bool:      condition.BoolField[nullabletypes.NullableTypes]{UpdatableField: condition.UpdatableField[nullabletypes.NullableTypes, bool]{Field: condition.Field[nullabletypes.NullableTypes, bool]{Name: "Bool"}}},
	Byte:      condition.UpdatableField[nullabletypes.NullableTypes, int8]{Field: condition.Field[nullabletypes.NullableTypes, int8]{Name: "Byte"}},
	CreatedAt: condition.Field[nullabletypes.NullableTypes, time.Time]{Name: "CreatedAt"},
	DeletedAt: condition.Field[nullabletypes.NullableTypes, time.Time]{Name: "DeletedAt"},
	Float64:   condition.UpdatableField[nullabletypes.NullableTypes, float64]{Field: condition.Field[nullabletypes.NullableTypes, float64]{Name: "Float64"}},
	ID:        condition.Field[nullabletypes.NullableTypes, model.UUID]{Name: "ID"},
	Int16:     condition.UpdatableField[nullabletypes.NullableTypes, int16]{Field: condition.Field[nullabletypes.NullableTypes, int16]{Name: "Int16"}},
	Int32:     condition.UpdatableField[nullabletypes.NullableTypes, int32]{Field: condition.Field[nullabletypes.NullableTypes, int32]{Name: "Int32"}},
	Int64:     condition.UpdatableField[nullabletypes.NullableTypes, int64]{Field: condition.Field[nullabletypes.NullableTypes, int64]{Name: "Int64"}},
	String:    condition.StringField[nullabletypes.NullableTypes]{UpdatableField: condition.UpdatableField[nullabletypes.NullableTypes, string]{Field: condition.Field[nullabletypes.NullableTypes, string]{Name: "String"}}},
	Time:      condition.UpdatableField[nullabletypes.NullableTypes, time.Time]{Field: condition.Field[nullabletypes.NullableTypes, time.Time]{Name: "Time"}},
	UpdatedAt: condition.Field[nullabletypes.NullableTypes, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the NullableTypes when doing a query
func (nullableTypesConditions nullableTypesConditions) Preload() condition.Condition[nullabletypes.NullableTypes] {
	return condition.NewPreloadCondition[nullabletypes.NullableTypes](nullableTypesConditions.ID, nullableTypesConditions.CreatedAt, nullableTypesConditions.UpdatedAt, nullableTypesConditions.DeletedAt, nullableTypesConditions.String, nullableTypesConditions.Int64, nullableTypesConditions.Int32, nullableTypesConditions.Int16, nullableTypesConditions.Byte, nullableTypesConditions.Float64, nullableTypesConditions.Bool, nullableTypesConditions.Time)
}
