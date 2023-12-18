// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	overrideforeignkey "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/overrideforeignkey"
	condition "github.com/ditrit/badaas/orm/condition"
	model "github.com/ditrit/badaas/orm/model"
	operator "github.com/ditrit/badaas/orm/operator"
	query "github.com/ditrit/badaas/orm/query"
	"reflect"
	"time"
)

var bicycleType = reflect.TypeOf(*new(overrideforeignkey.Bicycle))
var BicycleIdField = query.FieldIdentifier[model.UUID]{
	Field:     "ID",
	ModelType: bicycleType,
}

func BicycleId(operator operator.Operator[model.UUID]) condition.WhereCondition[overrideforeignkey.Bicycle] {
	return condition.NewFieldCondition[overrideforeignkey.Bicycle, model.UUID](BicycleIdField, operator)
}

var BicycleCreatedAtField = query.FieldIdentifier[time.Time]{
	Field:     "CreatedAt",
	ModelType: bicycleType,
}

func BicycleCreatedAt(operator operator.Operator[time.Time]) condition.WhereCondition[overrideforeignkey.Bicycle] {
	return condition.NewFieldCondition[overrideforeignkey.Bicycle, time.Time](BicycleCreatedAtField, operator)
}

var BicycleUpdatedAtField = query.FieldIdentifier[time.Time]{
	Field:     "UpdatedAt",
	ModelType: bicycleType,
}

func BicycleUpdatedAt(operator operator.Operator[time.Time]) condition.WhereCondition[overrideforeignkey.Bicycle] {
	return condition.NewFieldCondition[overrideforeignkey.Bicycle, time.Time](BicycleUpdatedAtField, operator)
}

var BicycleDeletedAtField = query.FieldIdentifier[time.Time]{
	Field:     "DeletedAt",
	ModelType: bicycleType,
}

func BicycleDeletedAt(operator operator.Operator[time.Time]) condition.WhereCondition[overrideforeignkey.Bicycle] {
	return condition.NewFieldCondition[overrideforeignkey.Bicycle, time.Time](BicycleDeletedAtField, operator)
}
func BicycleOwner(conditions ...condition.Condition[overrideforeignkey.Person]) condition.JoinCondition[overrideforeignkey.Bicycle] {
	return condition.NewJoinCondition[overrideforeignkey.Bicycle, overrideforeignkey.Person](conditions, "Owner", "OwnerSomethingID", BicyclePreloadAttributes, "ID")
}

var BicyclePreloadOwner = BicycleOwner(PersonPreloadAttributes)
var BicycleOwnerSomethingIdField = query.FieldIdentifier[string]{
	Field:     "OwnerSomethingID",
	ModelType: bicycleType,
}

func BicycleOwnerSomethingId(operator operator.Operator[string]) condition.WhereCondition[overrideforeignkey.Bicycle] {
	return condition.NewFieldCondition[overrideforeignkey.Bicycle, string](BicycleOwnerSomethingIdField, operator)
}

var BicyclePreloadAttributes = condition.NewPreloadCondition[overrideforeignkey.Bicycle](BicycleIdField, BicycleCreatedAtField, BicycleUpdatedAtField, BicycleDeletedAtField, BicycleOwnerSomethingIdField)
var BicyclePreloadRelations = []condition.Condition[overrideforeignkey.Bicycle]{BicyclePreloadOwner}