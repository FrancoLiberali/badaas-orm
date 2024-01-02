// Code generated by cql-gen v0.0.10, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	model "github.com/FrancoLiberali/cql/model"
	models "github.com/FrancoLiberali/cql/test/models"
	"time"
)

type universityConditions struct {
	ID        condition.Field[models.University, model.UUID]
	CreatedAt condition.Field[models.University, time.Time]
	UpdatedAt condition.Field[models.University, time.Time]
	DeletedAt condition.Field[models.University, time.Time]
	Name      condition.StringField[models.University]
}

var University = universityConditions{
	CreatedAt: condition.NewField[models.University, time.Time]("CreatedAt", "", ""),
	DeletedAt: condition.NewField[models.University, time.Time]("DeletedAt", "", ""),
	ID:        condition.NewField[models.University, model.UUID]("ID", "", ""),
	Name:      condition.NewStringField[models.University]("Name", "", ""),
	UpdatedAt: condition.NewField[models.University, time.Time]("UpdatedAt", "", ""),
}

// Preload allows preloading the University when doing a query
func (universityConditions universityConditions) preload() condition.Condition[models.University] {
	return condition.NewPreloadCondition[models.University](universityConditions.ID, universityConditions.CreatedAt, universityConditions.UpdatedAt, universityConditions.DeletedAt, universityConditions.Name)
}
