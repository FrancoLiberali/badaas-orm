// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	orm "github.com/ditrit/badaas/orm"
	model "github.com/ditrit/badaas/orm/model"
	models "github.com/ditrit/badaas/testintegration/models"
	"time"
)

func (parent2Conditions parent2Conditions) ParentParent(conditions ...orm.Condition[models.ParentParent]) orm.JoinCondition[models.Parent2] {
	return orm.NewJoinCondition[models.Parent2, models.ParentParent](conditions, "ParentParent", "ParentParentID", parent2Conditions.Preload(), "ID")
}
func (parent2Conditions parent2Conditions) PreloadParentParent() orm.JoinCondition[models.Parent2] {
	return parent2Conditions.ParentParent(ParentParent.Preload())
}

type parent2Conditions struct {
	ID             orm.Field[models.Parent2, model.UUID]
	CreatedAt      orm.Field[models.Parent2, time.Time]
	UpdatedAt      orm.Field[models.Parent2, time.Time]
	DeletedAt      orm.Field[models.Parent2, time.Time]
	ParentParentID orm.Field[models.Parent2, model.UUID]
}

var Parent2 = parent2Conditions{
	CreatedAt:      orm.Field[models.Parent2, time.Time]{Name: "CreatedAt"},
	DeletedAt:      orm.Field[models.Parent2, time.Time]{Name: "DeletedAt"},
	ID:             orm.Field[models.Parent2, model.UUID]{Name: "ID"},
	ParentParentID: orm.Field[models.Parent2, model.UUID]{Name: "ParentParentID"},
	UpdatedAt:      orm.Field[models.Parent2, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Parent2 when doing a query
func (parent2Conditions parent2Conditions) Preload() orm.Condition[models.Parent2] {
	return orm.NewPreloadCondition[models.Parent2](parent2Conditions.ID, parent2Conditions.CreatedAt, parent2Conditions.UpdatedAt, parent2Conditions.DeletedAt, parent2Conditions.ParentParentID)
}

// PreloadRelations allows preloading all the Parent2's relation when doing a query
func (parent2Conditions parent2Conditions) PreloadRelations() []orm.Condition[models.Parent2] {
	return []orm.Condition[models.Parent2]{parent2Conditions.PreloadParentParent()}
}
