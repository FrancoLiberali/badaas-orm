// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	orm "github.com/ditrit/badaas/orm"
	model "github.com/ditrit/badaas/orm/model"
	models "github.com/ditrit/badaas/testintegration/models"
	"time"
)

func (phoneConditions phoneConditions) Brand(conditions ...orm.Condition[models.Brand]) orm.JoinCondition[models.Phone] {
	return orm.NewJoinCondition[models.Phone, models.Brand](conditions, "Brand", "BrandID", phoneConditions.Preload(), "ID")
}
func (phoneConditions phoneConditions) PreloadBrand() orm.JoinCondition[models.Phone] {
	return phoneConditions.Brand(Brand.Preload())
}

type phoneConditions struct {
	ID        orm.Field[models.Phone, model.UIntID]
	CreatedAt orm.Field[models.Phone, time.Time]
	UpdatedAt orm.Field[models.Phone, time.Time]
	DeletedAt orm.Field[models.Phone, time.Time]
	Name      orm.StringField[models.Phone]
	BrandID   orm.Field[models.Phone, uint]
}

var Phone = phoneConditions{
	BrandID:   orm.Field[models.Phone, uint]{Name: "BrandID"},
	CreatedAt: orm.Field[models.Phone, time.Time]{Name: "CreatedAt"},
	DeletedAt: orm.Field[models.Phone, time.Time]{Name: "DeletedAt"},
	ID:        orm.Field[models.Phone, model.UIntID]{Name: "ID"},
	Name:      orm.StringField[models.Phone]{Field: orm.Field[models.Phone, string]{Name: "Name"}},
	UpdatedAt: orm.Field[models.Phone, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Phone when doing a query
func (phoneConditions phoneConditions) Preload() orm.Condition[models.Phone] {
	return orm.NewPreloadCondition[models.Phone](phoneConditions.ID, phoneConditions.CreatedAt, phoneConditions.UpdatedAt, phoneConditions.DeletedAt, phoneConditions.Name, phoneConditions.BrandID)
}

// PreloadRelations allows preloading all the Phone's relation when doing a query
func (phoneConditions phoneConditions) PreloadRelations() []orm.Condition[models.Phone] {
	return []orm.Condition[models.Phone]{phoneConditions.PreloadBrand()}
}
