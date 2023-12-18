// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	basicslicespointer "github.com/FrancoLiberali/cql/cql-cli/cmd/gen/conditions/tests/basicslicespointer"
	model "github.com/FrancoLiberali/cql/model"
	"time"
)

type basicSlicesPointerConditions struct {
	ID         condition.Field[basicslicespointer.BasicSlicesPointer, model.UUID]
	CreatedAt  condition.Field[basicslicespointer.BasicSlicesPointer, time.Time]
	UpdatedAt  condition.Field[basicslicespointer.BasicSlicesPointer, time.Time]
	DeletedAt  condition.Field[basicslicespointer.BasicSlicesPointer, time.Time]
	Bool       condition.Field[basicslicespointer.BasicSlicesPointer, []bool]
	Int        condition.Field[basicslicespointer.BasicSlicesPointer, []int]
	Int8       condition.Field[basicslicespointer.BasicSlicesPointer, []int8]
	Int16      condition.Field[basicslicespointer.BasicSlicesPointer, []int16]
	Int32      condition.Field[basicslicespointer.BasicSlicesPointer, []int32]
	Int64      condition.Field[basicslicespointer.BasicSlicesPointer, []int64]
	UInt       condition.Field[basicslicespointer.BasicSlicesPointer, []uint]
	UInt8      condition.Field[basicslicespointer.BasicSlicesPointer, []uint8]
	UInt16     condition.Field[basicslicespointer.BasicSlicesPointer, []uint16]
	UInt32     condition.Field[basicslicespointer.BasicSlicesPointer, []uint32]
	UInt64     condition.Field[basicslicespointer.BasicSlicesPointer, []uint64]
	UIntptr    condition.Field[basicslicespointer.BasicSlicesPointer, []uintptr]
	Float32    condition.Field[basicslicespointer.BasicSlicesPointer, []float32]
	Float64    condition.Field[basicslicespointer.BasicSlicesPointer, []float64]
	Complex64  condition.Field[basicslicespointer.BasicSlicesPointer, []complex64]
	Complex128 condition.Field[basicslicespointer.BasicSlicesPointer, []complex128]
	String     condition.Field[basicslicespointer.BasicSlicesPointer, []string]
	Byte       condition.Field[basicslicespointer.BasicSlicesPointer, []uint8]
}

var BasicSlicesPointer = basicSlicesPointerConditions{
	Bool:       condition.Field[basicslicespointer.BasicSlicesPointer, []bool]{Name: "Bool"},
	Byte:       condition.Field[basicslicespointer.BasicSlicesPointer, []uint8]{Name: "Byte"},
	Complex128: condition.Field[basicslicespointer.BasicSlicesPointer, []complex128]{Name: "Complex128"},
	Complex64:  condition.Field[basicslicespointer.BasicSlicesPointer, []complex64]{Name: "Complex64"},
	CreatedAt:  condition.Field[basicslicespointer.BasicSlicesPointer, time.Time]{Name: "CreatedAt"},
	DeletedAt:  condition.Field[basicslicespointer.BasicSlicesPointer, time.Time]{Name: "DeletedAt"},
	Float32:    condition.Field[basicslicespointer.BasicSlicesPointer, []float32]{Name: "Float32"},
	Float64:    condition.Field[basicslicespointer.BasicSlicesPointer, []float64]{Name: "Float64"},
	ID:         condition.Field[basicslicespointer.BasicSlicesPointer, model.UUID]{Name: "ID"},
	Int:        condition.Field[basicslicespointer.BasicSlicesPointer, []int]{Name: "Int"},
	Int16:      condition.Field[basicslicespointer.BasicSlicesPointer, []int16]{Name: "Int16"},
	Int32:      condition.Field[basicslicespointer.BasicSlicesPointer, []int32]{Name: "Int32"},
	Int64:      condition.Field[basicslicespointer.BasicSlicesPointer, []int64]{Name: "Int64"},
	Int8:       condition.Field[basicslicespointer.BasicSlicesPointer, []int8]{Name: "Int8"},
	String:     condition.Field[basicslicespointer.BasicSlicesPointer, []string]{Name: "String"},
	UInt:       condition.Field[basicslicespointer.BasicSlicesPointer, []uint]{Name: "UInt"},
	UInt16:     condition.Field[basicslicespointer.BasicSlicesPointer, []uint16]{Name: "UInt16"},
	UInt32:     condition.Field[basicslicespointer.BasicSlicesPointer, []uint32]{Name: "UInt32"},
	UInt64:     condition.Field[basicslicespointer.BasicSlicesPointer, []uint64]{Name: "UInt64"},
	UInt8:      condition.Field[basicslicespointer.BasicSlicesPointer, []uint8]{Name: "UInt8"},
	UIntptr:    condition.Field[basicslicespointer.BasicSlicesPointer, []uintptr]{Name: "UIntptr"},
	UpdatedAt:  condition.Field[basicslicespointer.BasicSlicesPointer, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the BasicSlicesPointer when doing a query
func (basicSlicesPointerConditions basicSlicesPointerConditions) Preload() condition.Condition[basicslicespointer.BasicSlicesPointer] {
	return condition.NewPreloadCondition[basicslicespointer.BasicSlicesPointer](basicSlicesPointerConditions.ID, basicSlicesPointerConditions.CreatedAt, basicSlicesPointerConditions.UpdatedAt, basicSlicesPointerConditions.DeletedAt, basicSlicesPointerConditions.Bool, basicSlicesPointerConditions.Int, basicSlicesPointerConditions.Int8, basicSlicesPointerConditions.Int16, basicSlicesPointerConditions.Int32, basicSlicesPointerConditions.Int64, basicSlicesPointerConditions.UInt, basicSlicesPointerConditions.UInt8, basicSlicesPointerConditions.UInt16, basicSlicesPointerConditions.UInt32, basicSlicesPointerConditions.UInt64, basicSlicesPointerConditions.UIntptr, basicSlicesPointerConditions.Float32, basicSlicesPointerConditions.Float64, basicSlicesPointerConditions.Complex64, basicSlicesPointerConditions.Complex128, basicSlicesPointerConditions.String, basicSlicesPointerConditions.Byte)
}
