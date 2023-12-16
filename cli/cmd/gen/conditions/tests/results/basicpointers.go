// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	basicpointers "github.com/ditrit/badaas-orm/cli/cmd/gen/conditions/tests/basicpointers"
	orm "github.com/ditrit/badaas/orm"
	"time"
)

func BasicPointersId(operator orm.Operator[orm.UUID]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, orm.UUID]{
		Field:    "ID",
		Operator: operator,
	}
}
func BasicPointersCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, time.Time]{
		Field:    "CreatedAt",
		Operator: operator,
	}
}
func BasicPointersUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, time.Time]{
		Field:    "UpdatedAt",
		Operator: operator,
	}
}
func BasicPointersDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, time.Time]{
		Field:    "DeletedAt",
		Operator: operator,
	}
}
func BasicPointersBool(operator orm.Operator[bool]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, bool]{
		Field:    "Bool",
		Operator: operator,
	}
}
func BasicPointersInt(operator orm.Operator[int]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, int]{
		Field:    "Int",
		Operator: operator,
	}
}
func BasicPointersInt8(operator orm.Operator[int8]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, int8]{
		Field:    "Int8",
		Operator: operator,
	}
}
func BasicPointersInt16(operator orm.Operator[int16]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, int16]{
		Field:    "Int16",
		Operator: operator,
	}
}
func BasicPointersInt32(operator orm.Operator[int32]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, int32]{
		Field:    "Int32",
		Operator: operator,
	}
}
func BasicPointersInt64(operator orm.Operator[int64]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, int64]{
		Field:    "Int64",
		Operator: operator,
	}
}
func BasicPointersUInt(operator orm.Operator[uint]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, uint]{
		Field:    "UInt",
		Operator: operator,
	}
}
func BasicPointersUInt8(operator orm.Operator[uint8]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, uint8]{
		Field:    "UInt8",
		Operator: operator,
	}
}
func BasicPointersUInt16(operator orm.Operator[uint16]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, uint16]{
		Field:    "UInt16",
		Operator: operator,
	}
}
func BasicPointersUInt32(operator orm.Operator[uint32]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, uint32]{
		Field:    "UInt32",
		Operator: operator,
	}
}
func BasicPointersUInt64(operator orm.Operator[uint64]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, uint64]{
		Field:    "UInt64",
		Operator: operator,
	}
}
func BasicPointersUIntptr(operator orm.Operator[uintptr]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, uintptr]{
		Field:    "UIntptr",
		Operator: operator,
	}
}
func BasicPointersFloat32(operator orm.Operator[float32]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, float32]{
		Field:    "Float32",
		Operator: operator,
	}
}
func BasicPointersFloat64(operator orm.Operator[float64]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, float64]{
		Field:    "Float64",
		Operator: operator,
	}
}
func BasicPointersComplex64(operator orm.Operator[complex64]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, complex64]{
		Field:    "Complex64",
		Operator: operator,
	}
}
func BasicPointersComplex128(operator orm.Operator[complex128]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, complex128]{
		Field:    "Complex128",
		Operator: operator,
	}
}
func BasicPointersString(operator orm.Operator[string]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, string]{
		Field:    "String",
		Operator: operator,
	}
}
func BasicPointersByte(operator orm.Operator[uint8]) orm.WhereCondition[basicpointers.BasicPointers] {
	return orm.FieldCondition[basicpointers.BasicPointers, uint8]{
		Field:    "Byte",
		Operator: operator,
	}
}
