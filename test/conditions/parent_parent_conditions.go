// Code generated by cql-cli v0.0.0, DO NOT EDIT.
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
	Number    condition.Field[models.ParentParent, int]
}

var ParentParent = parentParentConditions{
	CreatedAt: condition.Field[models.ParentParent, time.Time]{Name: "CreatedAt"},
	DeletedAt: condition.Field[models.ParentParent, time.Time]{Name: "DeletedAt"},
	ID:        condition.Field[models.ParentParent, model.UUID]{Name: "ID"},
	Name:      condition.StringField[models.ParentParent]{Field: condition.Field[models.ParentParent, string]{Name: "Name"}},
	Number:    condition.Field[models.ParentParent, int]{Name: "Number"},
	UpdatedAt: condition.Field[models.ParentParent, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the ParentParent when doing a query
func (parentParentConditions parentParentConditions) Preload() condition.Condition[models.ParentParent] {
	return condition.NewPreloadCondition[models.ParentParent](parentParentConditions.ID, parentParentConditions.CreatedAt, parentParentConditions.UpdatedAt, parentParentConditions.DeletedAt, parentParentConditions.Name, parentParentConditions.Number)
}
