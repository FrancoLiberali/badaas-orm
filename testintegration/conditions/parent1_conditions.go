// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	orm "github.com/ditrit/badaas/orm"
	model "github.com/ditrit/badaas/orm/model"
	models "github.com/ditrit/badaas/testintegration/models"
	"time"
)

func (parent1Conditions parent1Conditions) ParentParent(conditions ...orm.Condition[models.ParentParent]) orm.JoinCondition[models.Parent1] {
	return orm.NewJoinCondition[models.Parent1, models.ParentParent](conditions, "ParentParent", "ParentParentID", parent1Conditions.Preload(), "ID")
}
func (parent1Conditions parent1Conditions) PreloadParentParent() orm.JoinCondition[models.Parent1] {
	return parent1Conditions.ParentParent(ParentParent.Preload())
}

type parent1Conditions struct {
	ID             orm.Field[models.Parent1, model.UUID]
	CreatedAt      orm.Field[models.Parent1, time.Time]
	UpdatedAt      orm.Field[models.Parent1, time.Time]
	DeletedAt      orm.Field[models.Parent1, time.Time]
	ParentParentID orm.Field[models.Parent1, model.UUID]
}

var Parent1 = parent1Conditions{
	CreatedAt:      orm.Field[models.Parent1, time.Time]{Name: "CreatedAt"},
	DeletedAt:      orm.Field[models.Parent1, time.Time]{Name: "DeletedAt"},
	ID:             orm.Field[models.Parent1, model.UUID]{Name: "ID"},
	ParentParentID: orm.Field[models.Parent1, model.UUID]{Name: "ParentParentID"},
	UpdatedAt:      orm.Field[models.Parent1, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Parent1 when doing a query
func (parent1Conditions parent1Conditions) Preload() orm.Condition[models.Parent1] {
	return orm.NewPreloadCondition[models.Parent1](parent1Conditions.ID, parent1Conditions.CreatedAt, parent1Conditions.UpdatedAt, parent1Conditions.DeletedAt, parent1Conditions.ParentParentID)
}

// PreloadRelations allows preloading all the Parent1's relation when doing a query
func (parent1Conditions parent1Conditions) PreloadRelations() []orm.Condition[models.Parent1] {
	return []orm.Condition[models.Parent1]{parent1Conditions.PreloadParentParent()}
}
