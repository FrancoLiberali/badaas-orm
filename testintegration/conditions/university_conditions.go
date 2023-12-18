// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	cql "github.com/ditrit/badaas/orm/cql"
	model "github.com/ditrit/badaas/orm/model"
	models "github.com/ditrit/badaas/testintegration/models"
	"time"
)

type universityConditions struct {
	ID        cql.Field[models.University, model.UUID]
	CreatedAt cql.Field[models.University, time.Time]
	UpdatedAt cql.Field[models.University, time.Time]
	DeletedAt cql.Field[models.University, time.Time]
	Name      cql.StringField[models.University]
}

var University = universityConditions{
	CreatedAt: cql.Field[models.University, time.Time]{Name: "CreatedAt"},
	DeletedAt: cql.Field[models.University, time.Time]{Name: "DeletedAt"},
	ID:        cql.Field[models.University, model.UUID]{Name: "ID"},
	Name:      cql.StringField[models.University]{Field: cql.Field[models.University, string]{Name: "Name"}},
	UpdatedAt: cql.Field[models.University, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the University when doing a query
func (universityConditions universityConditions) Preload() cql.Condition[models.University] {
	return cql.NewPreloadCondition[models.University](universityConditions.ID, universityConditions.CreatedAt, universityConditions.UpdatedAt, universityConditions.DeletedAt, universityConditions.Name)
}
