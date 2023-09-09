// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	orm "github.com/ditrit/badaas/orm"
	model "github.com/ditrit/badaas/orm/model"
	models "github.com/ditrit/badaas/testintegration/models"
	"time"
)

func (employeeConditions employeeConditions) Boss(conditions ...orm.Condition[models.Employee]) orm.JoinCondition[models.Employee] {
	return orm.NewJoinCondition[models.Employee, models.Employee](conditions, "Boss", "BossID", employeeConditions.Preload(), "ID")
}
func (employeeConditions employeeConditions) PreloadBoss() orm.JoinCondition[models.Employee] {
	return employeeConditions.Boss(Employee.Preload())
}

type employeeConditions struct {
	ID        orm.Field[models.Employee, model.UUID]
	CreatedAt orm.Field[models.Employee, time.Time]
	UpdatedAt orm.Field[models.Employee, time.Time]
	DeletedAt orm.Field[models.Employee, time.Time]
	Name      orm.StringField[models.Employee]
	BossID    orm.Field[models.Employee, model.UUID]
}

var Employee = employeeConditions{
	BossID:    orm.Field[models.Employee, model.UUID]{Name: "BossID"},
	CreatedAt: orm.Field[models.Employee, time.Time]{Name: "CreatedAt"},
	DeletedAt: orm.Field[models.Employee, time.Time]{Name: "DeletedAt"},
	ID:        orm.Field[models.Employee, model.UUID]{Name: "ID"},
	Name:      orm.StringField[models.Employee]{Field: orm.Field[models.Employee, string]{Name: "Name"}},
	UpdatedAt: orm.Field[models.Employee, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Employee when doing a query
func (employeeConditions employeeConditions) Preload() orm.Condition[models.Employee] {
	return orm.NewPreloadCondition[models.Employee](employeeConditions.ID, employeeConditions.CreatedAt, employeeConditions.UpdatedAt, employeeConditions.DeletedAt, employeeConditions.Name, employeeConditions.BossID)
}

// PreloadRelations allows preloading all the Employee's relation when doing a query
func (employeeConditions employeeConditions) PreloadRelations() []orm.Condition[models.Employee] {
	return []orm.Condition[models.Employee]{employeeConditions.PreloadBoss()}
}
