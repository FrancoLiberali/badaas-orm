// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	selfreferential "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/selfreferential"
	cql "github.com/ditrit/badaas/orm/cql"
	model "github.com/ditrit/badaas/orm/model"
	"time"
)

func (employeeConditions employeeConditions) Boss(conditions ...cql.Condition[selfreferential.Employee]) cql.JoinCondition[selfreferential.Employee] {
	return cql.NewJoinCondition[selfreferential.Employee, selfreferential.Employee](conditions, "Boss", "BossID", employeeConditions.Preload(), "ID")
}
func (employeeConditions employeeConditions) PreloadBoss() cql.JoinCondition[selfreferential.Employee] {
	return employeeConditions.Boss(Employee.Preload())
}

type employeeConditions struct {
	ID        cql.Field[selfreferential.Employee, model.UUID]
	CreatedAt cql.Field[selfreferential.Employee, time.Time]
	UpdatedAt cql.Field[selfreferential.Employee, time.Time]
	DeletedAt cql.Field[selfreferential.Employee, time.Time]
	BossID    cql.Field[selfreferential.Employee, model.UUID]
}

var Employee = employeeConditions{
	BossID:    cql.Field[selfreferential.Employee, model.UUID]{Name: "BossID"},
	CreatedAt: cql.Field[selfreferential.Employee, time.Time]{Name: "CreatedAt"},
	DeletedAt: cql.Field[selfreferential.Employee, time.Time]{Name: "DeletedAt"},
	ID:        cql.Field[selfreferential.Employee, model.UUID]{Name: "ID"},
	UpdatedAt: cql.Field[selfreferential.Employee, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Employee when doing a query
func (employeeConditions employeeConditions) Preload() cql.Condition[selfreferential.Employee] {
	return cql.NewPreloadCondition[selfreferential.Employee](employeeConditions.ID, employeeConditions.CreatedAt, employeeConditions.UpdatedAt, employeeConditions.DeletedAt, employeeConditions.BossID)
}

// PreloadRelations allows preloading all the Employee's relation when doing a query
func (employeeConditions employeeConditions) PreloadRelations() []cql.Condition[selfreferential.Employee] {
	return []cql.Condition[selfreferential.Employee]{employeeConditions.PreloadBoss()}
}
