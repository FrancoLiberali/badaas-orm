// Code generated by cql-gen v0.0.10, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	model "github.com/FrancoLiberali/cql/model"
	models "github.com/FrancoLiberali/cql/test/models"
	"time"
)

type productConditions struct {
	ID                      condition.Field[models.Product, model.UUID]
	CreatedAt               condition.Field[models.Product, time.Time]
	UpdatedAt               condition.Field[models.Product, time.Time]
	DeletedAt               condition.Field[models.Product, time.Time]
	String                  condition.StringField[models.Product]
	Int                     condition.NumericField[models.Product, int]
	IntPointer              condition.NullableNumericField[models.Product, int]
	Float                   condition.NumericField[models.Product, float64]
	NullFloat               condition.NullableNumericField[models.Product, float64]
	Bool                    condition.BoolField[models.Product]
	NullBool                condition.NullableBoolField[models.Product]
	ByteArray               condition.UpdatableField[models.Product, []uint8]
	MultiString             condition.UpdatableField[models.Product, models.MultiString]
	ToBeEmbeddedEmbeddedInt condition.NumericField[models.Product, int]
	GormEmbeddedInt         condition.NumericField[models.Product, int]
	String2                 condition.StringField[models.Product]
}

var Product = productConditions{
	Bool:                    condition.NewBoolField[models.Product]("Bool", "", ""),
	ByteArray:               condition.NewUpdatableField[models.Product, []uint8]("ByteArray", "", ""),
	CreatedAt:               condition.NewField[models.Product, time.Time]("CreatedAt", "", ""),
	DeletedAt:               condition.NewField[models.Product, time.Time]("DeletedAt", "", ""),
	Float:                   condition.NewNumericField[models.Product, float64]("Float", "", ""),
	GormEmbeddedInt:         condition.NewNumericField[models.Product, int]("Int", "", "gorm_embedded_"),
	ID:                      condition.NewField[models.Product, model.UUID]("ID", "", ""),
	Int:                     condition.NewNumericField[models.Product, int]("Int", "", ""),
	IntPointer:              condition.NewNullableNumericField[models.Product, int]("IntPointer", "", ""),
	MultiString:             condition.NewUpdatableField[models.Product, models.MultiString]("MultiString", "", ""),
	NullBool:                condition.NewNullableBoolField[models.Product]("NullBool", "", ""),
	NullFloat:               condition.NewNullableNumericField[models.Product, float64]("NullFloat", "", ""),
	String:                  condition.NewStringField[models.Product]("String", "string_something_else", ""),
	String2:                 condition.NewStringField[models.Product]("String2", "", ""),
	ToBeEmbeddedEmbeddedInt: condition.NewNumericField[models.Product, int]("EmbeddedInt", "", ""),
	UpdatedAt:               condition.NewField[models.Product, time.Time]("UpdatedAt", "", ""),
}

// Preload allows preloading the Product when doing a query
func (productConditions productConditions) preload() condition.Condition[models.Product] {
	return condition.NewPreloadCondition[models.Product](productConditions.ID, productConditions.CreatedAt, productConditions.UpdatedAt, productConditions.DeletedAt, productConditions.String, productConditions.Int, productConditions.IntPointer, productConditions.Float, productConditions.NullFloat, productConditions.Bool, productConditions.NullBool, productConditions.ByteArray, productConditions.MultiString, productConditions.ToBeEmbeddedEmbeddedInt, productConditions.GormEmbeddedInt, productConditions.String2)
}
