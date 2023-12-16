// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	package1 "github.com/ditrit/badaas-orm/cli/cmd/gen/conditions/tests/multiplepackage/package1"
	package2 "github.com/ditrit/badaas-orm/cli/cmd/gen/conditions/tests/multiplepackage/package2"
	orm "github.com/ditrit/badaas/orm"
	"time"
)

func Package1Id(operator orm.Operator[orm.UUID]) orm.WhereCondition[package1.Package1] {
	return orm.FieldCondition[package1.Package1, orm.UUID]{
		Field:    "ID",
		Operator: operator,
	}
}
func Package1CreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[package1.Package1] {
	return orm.FieldCondition[package1.Package1, time.Time]{
		Field:    "CreatedAt",
		Operator: operator,
	}
}
func Package1UpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[package1.Package1] {
	return orm.FieldCondition[package1.Package1, time.Time]{
		Field:    "UpdatedAt",
		Operator: operator,
	}
}
func Package1DeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[package1.Package1] {
	return orm.FieldCondition[package1.Package1, time.Time]{
		Field:    "DeletedAt",
		Operator: operator,
	}
}
func Package1Package2(conditions ...orm.Condition[package2.Package2]) orm.Condition[package1.Package1] {
	return orm.JoinCondition[package1.Package1, package2.Package2]{
		Conditions: conditions,
		T1Field:    "ID",
		T2Field:    "Package1ID",
	}
}
func Package2Package1(conditions ...orm.Condition[package1.Package1]) orm.Condition[package2.Package2] {
	return orm.JoinCondition[package2.Package2, package1.Package1]{
		Conditions: conditions,
		T1Field:    "Package1ID",
		T2Field:    "ID",
	}
}
