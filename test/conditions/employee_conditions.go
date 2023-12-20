// Code generated by cql-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	model "github.com/FrancoLiberali/cql/model"
	models "github.com/FrancoLiberali/cql/test/models"
	"time"
)

func (employeeConditions employeeConditions) Boss(conditions ...condition.Condition[models.Employee]) condition.JoinCondition[models.Employee] {
	return condition.NewJoinCondition[models.Employee, models.Employee](conditions, "Boss", "BossID", employeeConditions.Preload(), "ID")
}
func (employeeConditions employeeConditions) PreloadBoss() condition.JoinCondition[models.Employee] {
	return employeeConditions.Boss(Employee.Preload())
}

type employeeConditions struct {
	ID        condition.Field[models.Employee, model.UUID]
	CreatedAt condition.Field[models.Employee, time.Time]
	UpdatedAt condition.Field[models.Employee, time.Time]
	DeletedAt condition.Field[models.Employee, time.Time]
	Name      condition.StringField[models.Employee]
	BossID    condition.Field[models.Employee, model.UUID]
}

var Employee = employeeConditions{
	BossID:    condition.Field[models.Employee, model.UUID]{Name: "BossID"},
	CreatedAt: condition.Field[models.Employee, time.Time]{Name: "CreatedAt"},
	DeletedAt: condition.Field[models.Employee, time.Time]{Name: "DeletedAt"},
	ID:        condition.Field[models.Employee, model.UUID]{Name: "ID"},
	Name:      condition.StringField[models.Employee]{UpdatableField: condition.UpdatableField[models.Employee, string]{Field: condition.Field[models.Employee, string]{Name: "Name"}}},
	UpdatedAt: condition.Field[models.Employee, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Employee when doing a query
func (employeeConditions employeeConditions) Preload() condition.Condition[models.Employee] {
	return condition.NewPreloadCondition[models.Employee](employeeConditions.ID, employeeConditions.CreatedAt, employeeConditions.UpdatedAt, employeeConditions.DeletedAt, employeeConditions.Name, employeeConditions.BossID)
}

// PreloadRelations allows preloading all the Employee's relation when doing a query
func (employeeConditions employeeConditions) PreloadRelations() []condition.Condition[models.Employee] {
	return []condition.Condition[models.Employee]{employeeConditions.PreloadBoss()}
}
