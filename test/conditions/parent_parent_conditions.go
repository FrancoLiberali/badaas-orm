// Code generated by cql-gen v0.0.10, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	model "github.com/FrancoLiberali/cql/model"
	models "github.com/FrancoLiberali/cql/test/models"
	"time"
)

type parentParentConditions struct {
	ID        condition.Field[models.ParentParent, model.UUID]
	CreatedAt condition.Field[models.ParentParent, time.Time]
	UpdatedAt condition.Field[models.ParentParent, time.Time]
	DeletedAt condition.Field[models.ParentParent, time.Time]
	Name      condition.StringField[models.ParentParent]
	Number    condition.NumericField[models.ParentParent, int]
}

var ParentParent = parentParentConditions{
	CreatedAt: condition.NewField[models.ParentParent, time.Time]("CreatedAt", "", ""),
	DeletedAt: condition.NewField[models.ParentParent, time.Time]("DeletedAt", "", ""),
	ID:        condition.NewField[models.ParentParent, model.UUID]("ID", "", ""),
	Name:      condition.NewStringField[models.ParentParent]("Name", "", ""),
	Number:    condition.NewNumericField[models.ParentParent, int]("Number", "", ""),
	UpdatedAt: condition.NewField[models.ParentParent, time.Time]("UpdatedAt", "", ""),
}

// Preload allows preloading the ParentParent when doing a query
func (parentParentConditions parentParentConditions) preload() condition.Condition[models.ParentParent] {
	return condition.NewPreloadCondition[models.ParentParent](parentParentConditions.ID, parentParentConditions.CreatedAt, parentParentConditions.UpdatedAt, parentParentConditions.DeletedAt, parentParentConditions.Name, parentParentConditions.Number)
}
