// Code generated by cql-gen v0.0.8, DO NOT EDIT.
package conditions

import (
	basicpointers "github.com/FrancoLiberali/cql-gen/cmd/gen/conditions/tests/basicpointers"
	condition "github.com/FrancoLiberali/cql/condition"
	model "github.com/FrancoLiberali/cql/model"
	"time"
)

type basicPointersConditions struct {
	ID         condition.Field[basicpointers.BasicPointers, model.UUID]
	CreatedAt  condition.Field[basicpointers.BasicPointers, time.Time]
	UpdatedAt  condition.Field[basicpointers.BasicPointers, time.Time]
	DeletedAt  condition.Field[basicpointers.BasicPointers, time.Time]
	Bool       condition.NullableBoolField[basicpointers.BasicPointers]
	Int        condition.NullableField[basicpointers.BasicPointers, int]
	Int8       condition.NullableField[basicpointers.BasicPointers, int8]
	Int16      condition.NullableField[basicpointers.BasicPointers, int16]
	Int32      condition.NullableField[basicpointers.BasicPointers, int32]
	Int64      condition.NullableField[basicpointers.BasicPointers, int64]
	UInt       condition.NullableField[basicpointers.BasicPointers, uint]
	UInt8      condition.NullableField[basicpointers.BasicPointers, uint8]
	UInt16     condition.NullableField[basicpointers.BasicPointers, uint16]
	UInt32     condition.NullableField[basicpointers.BasicPointers, uint32]
	UInt64     condition.NullableField[basicpointers.BasicPointers, uint64]
	UIntptr    condition.NullableField[basicpointers.BasicPointers, uintptr]
	Float32    condition.NullableField[basicpointers.BasicPointers, float32]
	Float64    condition.NullableField[basicpointers.BasicPointers, float64]
	Complex64  condition.NullableField[basicpointers.BasicPointers, complex64]
	Complex128 condition.NullableField[basicpointers.BasicPointers, complex128]
	String     condition.NullableStringField[basicpointers.BasicPointers]
	Byte       condition.NullableField[basicpointers.BasicPointers, uint8]
}

var BasicPointers = basicPointersConditions{
	Bool:       condition.NewNullableBoolField[basicpointers.BasicPointers]("Bool", "", ""),
	Byte:       condition.NewNullableField[basicpointers.BasicPointers, uint8]("Byte", "", ""),
	Complex128: condition.NewNullableField[basicpointers.BasicPointers, complex128]("Complex128", "", ""),
	Complex64:  condition.NewNullableField[basicpointers.BasicPointers, complex64]("Complex64", "", ""),
	CreatedAt:  condition.NewField[basicpointers.BasicPointers, time.Time]("CreatedAt", "", ""),
	DeletedAt:  condition.NewField[basicpointers.BasicPointers, time.Time]("DeletedAt", "", ""),
	Float32:    condition.NewNullableField[basicpointers.BasicPointers, float32]("Float32", "", ""),
	Float64:    condition.NewNullableField[basicpointers.BasicPointers, float64]("Float64", "", ""),
	ID:         condition.NewField[basicpointers.BasicPointers, model.UUID]("ID", "", ""),
	Int:        condition.NewNullableField[basicpointers.BasicPointers, int]("Int", "", ""),
	Int16:      condition.NewNullableField[basicpointers.BasicPointers, int16]("Int16", "", ""),
	Int32:      condition.NewNullableField[basicpointers.BasicPointers, int32]("Int32", "", ""),
	Int64:      condition.NewNullableField[basicpointers.BasicPointers, int64]("Int64", "", ""),
	Int8:       condition.NewNullableField[basicpointers.BasicPointers, int8]("Int8", "", ""),
	String:     condition.NewNullableStringField[basicpointers.BasicPointers]("String", "", ""),
	UInt:       condition.NewNullableField[basicpointers.BasicPointers, uint]("UInt", "", ""),
	UInt16:     condition.NewNullableField[basicpointers.BasicPointers, uint16]("UInt16", "", ""),
	UInt32:     condition.NewNullableField[basicpointers.BasicPointers, uint32]("UInt32", "", ""),
	UInt64:     condition.NewNullableField[basicpointers.BasicPointers, uint64]("UInt64", "", ""),
	UInt8:      condition.NewNullableField[basicpointers.BasicPointers, uint8]("UInt8", "", ""),
	UIntptr:    condition.NewNullableField[basicpointers.BasicPointers, uintptr]("UIntptr", "", ""),
	UpdatedAt:  condition.NewField[basicpointers.BasicPointers, time.Time]("UpdatedAt", "", ""),
}

// Preload allows preloading the BasicPointers when doing a query
func (basicPointersConditions basicPointersConditions) preload() condition.Condition[basicpointers.BasicPointers] {
	return condition.NewPreloadCondition[basicpointers.BasicPointers](basicPointersConditions.ID, basicPointersConditions.CreatedAt, basicPointersConditions.UpdatedAt, basicPointersConditions.DeletedAt, basicPointersConditions.Bool, basicPointersConditions.Int, basicPointersConditions.Int8, basicPointersConditions.Int16, basicPointersConditions.Int32, basicPointersConditions.Int64, basicPointersConditions.UInt, basicPointersConditions.UInt8, basicPointersConditions.UInt16, basicPointersConditions.UInt32, basicPointersConditions.UInt64, basicPointersConditions.UIntptr, basicPointersConditions.Float32, basicPointersConditions.Float64, basicPointersConditions.Complex64, basicPointersConditions.Complex128, basicPointersConditions.String, basicPointersConditions.Byte)
}
