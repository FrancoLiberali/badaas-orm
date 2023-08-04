// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	orm "github.com/ditrit/badaas/orm"
	models "github.com/ditrit/badaas/testintegration/models"
	"reflect"
	"time"
)

var employeeType = reflect.TypeOf(*new(models.Employee))
var EmployeeIdField = orm.FieldIdentifier[orm.UUID]{
	Field:     "ID",
	ModelType: employeeType,
}

func EmployeeId(operator orm.Operator[orm.UUID]) orm.WhereCondition[models.Employee] {
	return orm.FieldCondition[models.Employee, orm.UUID]{
		FieldIdentifier: EmployeeIdField,
		Operator:        operator,
	}
}

var EmployeeCreatedAtField = orm.FieldIdentifier[time.Time]{
	Field:     "CreatedAt",
	ModelType: employeeType,
}

func EmployeeCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[models.Employee] {
	return orm.FieldCondition[models.Employee, time.Time]{
		FieldIdentifier: EmployeeCreatedAtField,
		Operator:        operator,
	}
}

var EmployeeUpdatedAtField = orm.FieldIdentifier[time.Time]{
	Field:     "UpdatedAt",
	ModelType: employeeType,
}

func EmployeeUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[models.Employee] {
	return orm.FieldCondition[models.Employee, time.Time]{
		FieldIdentifier: EmployeeUpdatedAtField,
		Operator:        operator,
	}
}

var EmployeeDeletedAtField = orm.FieldIdentifier[time.Time]{
	Field:     "DeletedAt",
	ModelType: employeeType,
}

func EmployeeDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[models.Employee] {
	return orm.FieldCondition[models.Employee, time.Time]{
		FieldIdentifier: EmployeeDeletedAtField,
		Operator:        operator,
	}
}

var EmployeeNameField = orm.FieldIdentifier[string]{
	Field:     "Name",
	ModelType: employeeType,
}

func EmployeeName(operator orm.Operator[string]) orm.WhereCondition[models.Employee] {
	return orm.FieldCondition[models.Employee, string]{
		FieldIdentifier: EmployeeNameField,
		Operator:        operator,
	}
}
func EmployeeBoss(conditions ...orm.Condition[models.Employee]) orm.IJoinCondition[models.Employee] {
	return orm.JoinCondition[models.Employee, models.Employee]{
		Conditions:         conditions,
		RelationField:      "Boss",
		T1Field:            "BossID",
		T1PreloadCondition: EmployeePreloadAttributes,
		T2Field:            "ID",
	}
}

var EmployeePreloadBoss = EmployeeBoss(EmployeePreloadAttributes)
var EmployeeBossIdField = orm.FieldIdentifier[orm.UUID]{
	Field:     "BossID",
	ModelType: employeeType,
}

func EmployeeBossId(operator orm.Operator[orm.UUID]) orm.WhereCondition[models.Employee] {
	return orm.FieldCondition[models.Employee, orm.UUID]{
		FieldIdentifier: EmployeeBossIdField,
		Operator:        operator,
	}
}

var EmployeePreloadAttributes = orm.NewPreloadCondition[models.Employee](EmployeeIdField, EmployeeCreatedAtField, EmployeeUpdatedAtField, EmployeeDeletedAtField, EmployeeNameField, EmployeeBossIdField)
var EmployeePreloadRelations = []orm.Condition[models.Employee]{EmployeePreloadBoss}
