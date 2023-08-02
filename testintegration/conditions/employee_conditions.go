// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	orm "github.com/ditrit/badaas/orm"
	models "github.com/ditrit/badaas/testintegration/models"
	"time"
)

func EmployeeId(operator orm.Operator[orm.UUID]) orm.WhereCondition[models.Employee] {
	return orm.FieldCondition[models.Employee, orm.UUID]{
		Field:    "ID",
		Operator: operator,
	}
}
func EmployeeCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[models.Employee] {
	return orm.FieldCondition[models.Employee, time.Time]{
		Field:    "CreatedAt",
		Operator: operator,
	}
}
func EmployeeUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[models.Employee] {
	return orm.FieldCondition[models.Employee, time.Time]{
		Field:    "UpdatedAt",
		Operator: operator,
	}
}
func EmployeeDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[models.Employee] {
	return orm.FieldCondition[models.Employee, time.Time]{
		Field:    "DeletedAt",
		Operator: operator,
	}
}
func EmployeeName(operator orm.Operator[string]) orm.WhereCondition[models.Employee] {
	return orm.FieldCondition[models.Employee, string]{
		Field:    "Name",
		Operator: operator,
	}
}
func EmployeeBoss(conditions ...orm.Condition[models.Employee]) orm.Condition[models.Employee] {
	return orm.JoinCondition[models.Employee, models.Employee]{
		Conditions: conditions,
		T1Field:    "BossID",
		T2Field:    "ID",
	}
}
func EmployeeBossId(operator orm.Operator[orm.UUID]) orm.WhereCondition[models.Employee] {
	return orm.FieldCondition[models.Employee, orm.UUID]{
		Field:    "BossID",
		Operator: operator,
	}
}
