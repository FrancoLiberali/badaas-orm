// Code generated by cql-gen v0.0.8, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	gormembedded "github.com/FrancoLiberali/cql-gen/cmd/gen/conditions/tests/gormembedded"
	model "github.com/FrancoLiberali/cql/model"
	"time"
)

type gormEmbeddedConditions struct {
	ID                      condition.Field[gormembedded.GormEmbedded, model.UIntID]
	CreatedAt               condition.Field[gormembedded.GormEmbedded, time.Time]
	UpdatedAt               condition.Field[gormembedded.GormEmbedded, time.Time]
	DeletedAt               condition.Field[gormembedded.GormEmbedded, time.Time]
	Int                     condition.UpdatableField[gormembedded.GormEmbedded, int]
	GormEmbeddedInt         condition.UpdatableField[gormembedded.GormEmbedded, int]
	GormEmbeddedNoPrefixInt condition.UpdatableField[gormembedded.GormEmbedded, int]
}

var GormEmbedded = gormEmbeddedConditions{
	CreatedAt:               condition.NewField[gormembedded.GormEmbedded, time.Time]("CreatedAt", "", ""),
	DeletedAt:               condition.NewField[gormembedded.GormEmbedded, time.Time]("DeletedAt", "", ""),
	GormEmbeddedInt:         condition.NewUpdatableField[gormembedded.GormEmbedded, int]("Int", "", "gorm_embedded_"),
	GormEmbeddedNoPrefixInt: condition.NewUpdatableField[gormembedded.GormEmbedded, int]("Int", "", ""),
	ID:                      condition.NewField[gormembedded.GormEmbedded, model.UIntID]("ID", "", ""),
	Int:                     condition.NewUpdatableField[gormembedded.GormEmbedded, int]("Int", "", ""),
	UpdatedAt:               condition.NewField[gormembedded.GormEmbedded, time.Time]("UpdatedAt", "", ""),
}

// Preload allows preloading the GormEmbedded when doing a query
func (gormEmbeddedConditions gormEmbeddedConditions) preload() condition.Condition[gormembedded.GormEmbedded] {
	return condition.NewPreloadCondition[gormembedded.GormEmbedded](gormEmbeddedConditions.ID, gormEmbeddedConditions.CreatedAt, gormEmbeddedConditions.UpdatedAt, gormEmbeddedConditions.DeletedAt, gormEmbeddedConditions.Int, gormEmbeddedConditions.GormEmbeddedInt, gormEmbeddedConditions.GormEmbeddedNoPrefixInt)
}
