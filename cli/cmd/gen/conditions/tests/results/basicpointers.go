// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	basicpointers "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/basicpointers"
	orm "github.com/ditrit/badaas/orm"
	condition "github.com/ditrit/badaas/orm/condition"
	model "github.com/ditrit/badaas/orm/model"
	query "github.com/ditrit/badaas/orm/query"
	"reflect"
	"time"
)

var basicPointersType = reflect.TypeOf(*new(basicpointers.BasicPointers))

func (basicPointersConditions basicPointersConditions) IdIs() orm.FieldIs[basicpointers.BasicPointers, model.UUID] {
	return orm.FieldIs[basicpointers.BasicPointers, model.UUID]{FieldID: basicPointersConditions.ID}
}
func (basicPointersConditions basicPointersConditions) CreatedAtIs() orm.FieldIs[basicpointers.BasicPointers, time.Time] {
	return orm.FieldIs[basicpointers.BasicPointers, time.Time]{FieldID: basicPointersConditions.CreatedAt}
}
func (basicPointersConditions basicPointersConditions) UpdatedAtIs() orm.FieldIs[basicpointers.BasicPointers, time.Time] {
	return orm.FieldIs[basicpointers.BasicPointers, time.Time]{FieldID: basicPointersConditions.UpdatedAt}
}
func (basicPointersConditions basicPointersConditions) DeletedAtIs() orm.FieldIs[basicpointers.BasicPointers, time.Time] {
	return orm.FieldIs[basicpointers.BasicPointers, time.Time]{FieldID: basicPointersConditions.DeletedAt}
}
func (basicPointersConditions basicPointersConditions) BoolIs() orm.BoolFieldIs[basicpointers.BasicPointers] {
	return orm.BoolFieldIs[basicpointers.BasicPointers]{FieldIs: orm.FieldIs[basicpointers.BasicPointers, bool]{FieldID: basicPointersConditions.Bool}}
}
func (basicPointersConditions basicPointersConditions) IntIs() orm.FieldIs[basicpointers.BasicPointers, int] {
	return orm.FieldIs[basicpointers.BasicPointers, int]{FieldID: basicPointersConditions.Int}
}
func (basicPointersConditions basicPointersConditions) Int8Is() orm.FieldIs[basicpointers.BasicPointers, int8] {
	return orm.FieldIs[basicpointers.BasicPointers, int8]{FieldID: basicPointersConditions.Int8}
}
func (basicPointersConditions basicPointersConditions) Int16Is() orm.FieldIs[basicpointers.BasicPointers, int16] {
	return orm.FieldIs[basicpointers.BasicPointers, int16]{FieldID: basicPointersConditions.Int16}
}
func (basicPointersConditions basicPointersConditions) Int32Is() orm.FieldIs[basicpointers.BasicPointers, int32] {
	return orm.FieldIs[basicpointers.BasicPointers, int32]{FieldID: basicPointersConditions.Int32}
}
func (basicPointersConditions basicPointersConditions) Int64Is() orm.FieldIs[basicpointers.BasicPointers, int64] {
	return orm.FieldIs[basicpointers.BasicPointers, int64]{FieldID: basicPointersConditions.Int64}
}
func (basicPointersConditions basicPointersConditions) UIntIs() orm.FieldIs[basicpointers.BasicPointers, uint] {
	return orm.FieldIs[basicpointers.BasicPointers, uint]{FieldID: basicPointersConditions.UInt}
}
func (basicPointersConditions basicPointersConditions) UInt8Is() orm.FieldIs[basicpointers.BasicPointers, uint8] {
	return orm.FieldIs[basicpointers.BasicPointers, uint8]{FieldID: basicPointersConditions.UInt8}
}
func (basicPointersConditions basicPointersConditions) UInt16Is() orm.FieldIs[basicpointers.BasicPointers, uint16] {
	return orm.FieldIs[basicpointers.BasicPointers, uint16]{FieldID: basicPointersConditions.UInt16}
}
func (basicPointersConditions basicPointersConditions) UInt32Is() orm.FieldIs[basicpointers.BasicPointers, uint32] {
	return orm.FieldIs[basicpointers.BasicPointers, uint32]{FieldID: basicPointersConditions.UInt32}
}
func (basicPointersConditions basicPointersConditions) UInt64Is() orm.FieldIs[basicpointers.BasicPointers, uint64] {
	return orm.FieldIs[basicpointers.BasicPointers, uint64]{FieldID: basicPointersConditions.UInt64}
}
func (basicPointersConditions basicPointersConditions) UIntptrIs() orm.FieldIs[basicpointers.BasicPointers, uintptr] {
	return orm.FieldIs[basicpointers.BasicPointers, uintptr]{FieldID: basicPointersConditions.UIntptr}
}
func (basicPointersConditions basicPointersConditions) Float32Is() orm.FieldIs[basicpointers.BasicPointers, float32] {
	return orm.FieldIs[basicpointers.BasicPointers, float32]{FieldID: basicPointersConditions.Float32}
}
func (basicPointersConditions basicPointersConditions) Float64Is() orm.FieldIs[basicpointers.BasicPointers, float64] {
	return orm.FieldIs[basicpointers.BasicPointers, float64]{FieldID: basicPointersConditions.Float64}
}
func (basicPointersConditions basicPointersConditions) Complex64Is() orm.FieldIs[basicpointers.BasicPointers, complex64] {
	return orm.FieldIs[basicpointers.BasicPointers, complex64]{FieldID: basicPointersConditions.Complex64}
}
func (basicPointersConditions basicPointersConditions) Complex128Is() orm.FieldIs[basicpointers.BasicPointers, complex128] {
	return orm.FieldIs[basicpointers.BasicPointers, complex128]{FieldID: basicPointersConditions.Complex128}
}
func (basicPointersConditions basicPointersConditions) StringIs() orm.StringFieldIs[basicpointers.BasicPointers] {
	return orm.StringFieldIs[basicpointers.BasicPointers]{FieldIs: orm.FieldIs[basicpointers.BasicPointers, string]{FieldID: basicPointersConditions.String}}
}
func (basicPointersConditions basicPointersConditions) ByteIs() orm.FieldIs[basicpointers.BasicPointers, uint8] {
	return orm.FieldIs[basicpointers.BasicPointers, uint8]{FieldID: basicPointersConditions.Byte}
}

type basicPointersConditions struct {
	ID         query.FieldIdentifier[model.UUID]
	CreatedAt  query.FieldIdentifier[time.Time]
	UpdatedAt  query.FieldIdentifier[time.Time]
	DeletedAt  query.FieldIdentifier[time.Time]
	Bool       query.FieldIdentifier[bool]
	Int        query.FieldIdentifier[int]
	Int8       query.FieldIdentifier[int8]
	Int16      query.FieldIdentifier[int16]
	Int32      query.FieldIdentifier[int32]
	Int64      query.FieldIdentifier[int64]
	UInt       query.FieldIdentifier[uint]
	UInt8      query.FieldIdentifier[uint8]
	UInt16     query.FieldIdentifier[uint16]
	UInt32     query.FieldIdentifier[uint32]
	UInt64     query.FieldIdentifier[uint64]
	UIntptr    query.FieldIdentifier[uintptr]
	Float32    query.FieldIdentifier[float32]
	Float64    query.FieldIdentifier[float64]
	Complex64  query.FieldIdentifier[complex64]
	Complex128 query.FieldIdentifier[complex128]
	String     query.FieldIdentifier[string]
	Byte       query.FieldIdentifier[uint8]
}

var BasicPointers = basicPointersConditions{
	Bool: query.FieldIdentifier[bool]{
		Field:     "Bool",
		ModelType: basicPointersType,
	},
	Byte: query.FieldIdentifier[uint8]{
		Field:     "Byte",
		ModelType: basicPointersType,
	},
	Complex128: query.FieldIdentifier[complex128]{
		Field:     "Complex128",
		ModelType: basicPointersType,
	},
	Complex64: query.FieldIdentifier[complex64]{
		Field:     "Complex64",
		ModelType: basicPointersType,
	},
	CreatedAt: query.FieldIdentifier[time.Time]{
		Field:     "CreatedAt",
		ModelType: basicPointersType,
	},
	DeletedAt: query.FieldIdentifier[time.Time]{
		Field:     "DeletedAt",
		ModelType: basicPointersType,
	},
	Float32: query.FieldIdentifier[float32]{
		Field:     "Float32",
		ModelType: basicPointersType,
	},
	Float64: query.FieldIdentifier[float64]{
		Field:     "Float64",
		ModelType: basicPointersType,
	},
	ID: query.FieldIdentifier[model.UUID]{
		Field:     "ID",
		ModelType: basicPointersType,
	},
	Int: query.FieldIdentifier[int]{
		Field:     "Int",
		ModelType: basicPointersType,
	},
	Int16: query.FieldIdentifier[int16]{
		Field:     "Int16",
		ModelType: basicPointersType,
	},
	Int32: query.FieldIdentifier[int32]{
		Field:     "Int32",
		ModelType: basicPointersType,
	},
	Int64: query.FieldIdentifier[int64]{
		Field:     "Int64",
		ModelType: basicPointersType,
	},
	Int8: query.FieldIdentifier[int8]{
		Field:     "Int8",
		ModelType: basicPointersType,
	},
	String: query.FieldIdentifier[string]{
		Field:     "String",
		ModelType: basicPointersType,
	},
	UInt: query.FieldIdentifier[uint]{
		Field:     "UInt",
		ModelType: basicPointersType,
	},
	UInt16: query.FieldIdentifier[uint16]{
		Field:     "UInt16",
		ModelType: basicPointersType,
	},
	UInt32: query.FieldIdentifier[uint32]{
		Field:     "UInt32",
		ModelType: basicPointersType,
	},
	UInt64: query.FieldIdentifier[uint64]{
		Field:     "UInt64",
		ModelType: basicPointersType,
	},
	UInt8: query.FieldIdentifier[uint8]{
		Field:     "UInt8",
		ModelType: basicPointersType,
	},
	UIntptr: query.FieldIdentifier[uintptr]{
		Field:     "UIntptr",
		ModelType: basicPointersType,
	},
	UpdatedAt: query.FieldIdentifier[time.Time]{
		Field:     "UpdatedAt",
		ModelType: basicPointersType,
	},
}

// Preload allows preloading the BasicPointers when doing a query
func (basicPointersConditions basicPointersConditions) Preload() condition.Condition[basicpointers.BasicPointers] {
	return condition.NewPreloadCondition[basicpointers.BasicPointers](basicPointersConditions.ID, basicPointersConditions.CreatedAt, basicPointersConditions.UpdatedAt, basicPointersConditions.DeletedAt, basicPointersConditions.Bool, basicPointersConditions.Int, basicPointersConditions.Int8, basicPointersConditions.Int16, basicPointersConditions.Int32, basicPointersConditions.Int64, basicPointersConditions.UInt, basicPointersConditions.UInt8, basicPointersConditions.UInt16, basicPointersConditions.UInt32, basicPointersConditions.UInt64, basicPointersConditions.UIntptr, basicPointersConditions.Float32, basicPointersConditions.Float64, basicPointersConditions.Complex64, basicPointersConditions.Complex128, basicPointersConditions.String, basicPointersConditions.Byte)
}
