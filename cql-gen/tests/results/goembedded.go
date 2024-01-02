// Code generated by cql-gen v0.0.8, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	goembedded "github.com/FrancoLiberali/cql-gen/cmd/gen/conditions/tests/goembedded"
	model "github.com/FrancoLiberali/cql/model"
	"time"
)

type goEmbeddedConditions struct {
	ID              condition.Field[goembedded.GoEmbedded, model.UIntID]
	CreatedAt       condition.Field[goembedded.GoEmbedded, time.Time]
	UpdatedAt       condition.Field[goembedded.GoEmbedded, time.Time]
	DeletedAt       condition.Field[goembedded.GoEmbedded, time.Time]
	Int             condition.UpdatableField[goembedded.GoEmbedded, int]
	ToBeEmbeddedInt condition.UpdatableField[goembedded.GoEmbedded, int]
}

var GoEmbedded = goEmbeddedConditions{
	CreatedAt:       condition.NewField[goembedded.GoEmbedded, time.Time]("CreatedAt", "", ""),
	DeletedAt:       condition.NewField[goembedded.GoEmbedded, time.Time]("DeletedAt", "", ""),
	ID:              condition.NewField[goembedded.GoEmbedded, model.UIntID]("ID", "", ""),
	Int:             condition.NewUpdatableField[goembedded.GoEmbedded, int]("Int", "", ""),
	ToBeEmbeddedInt: condition.NewUpdatableField[goembedded.GoEmbedded, int]("Int", "", ""),
	UpdatedAt:       condition.NewField[goembedded.GoEmbedded, time.Time]("UpdatedAt", "", ""),
}

// Preload allows preloading the GoEmbedded when doing a query
func (goEmbeddedConditions goEmbeddedConditions) preload() condition.Condition[goembedded.GoEmbedded] {
	return condition.NewPreloadCondition[goembedded.GoEmbedded](goEmbeddedConditions.ID, goEmbeddedConditions.CreatedAt, goEmbeddedConditions.UpdatedAt, goEmbeddedConditions.DeletedAt, goEmbeddedConditions.Int, goEmbeddedConditions.ToBeEmbeddedInt)
}
