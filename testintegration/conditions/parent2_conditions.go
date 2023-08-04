// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	orm "github.com/ditrit/badaas/orm"
	models "github.com/ditrit/badaas/testintegration/models"
	"reflect"
	"time"
)

var parent2Type = reflect.TypeOf(*new(models.Parent2))
var Parent2IdField = orm.FieldIdentifier[orm.UUID]{
	Field:     "ID",
	ModelType: parent2Type,
}

func Parent2Id(operator orm.Operator[orm.UUID]) orm.WhereCondition[models.Parent2] {
	return orm.FieldCondition[models.Parent2, orm.UUID]{
		FieldIdentifier: Parent2IdField,
		Operator:        operator,
	}
}

var Parent2CreatedAtField = orm.FieldIdentifier[time.Time]{
	Field:     "CreatedAt",
	ModelType: parent2Type,
}

func Parent2CreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[models.Parent2] {
	return orm.FieldCondition[models.Parent2, time.Time]{
		FieldIdentifier: Parent2CreatedAtField,
		Operator:        operator,
	}
}

var Parent2UpdatedAtField = orm.FieldIdentifier[time.Time]{
	Field:     "UpdatedAt",
	ModelType: parent2Type,
}

func Parent2UpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[models.Parent2] {
	return orm.FieldCondition[models.Parent2, time.Time]{
		FieldIdentifier: Parent2UpdatedAtField,
		Operator:        operator,
	}
}

var Parent2DeletedAtField = orm.FieldIdentifier[time.Time]{
	Field:     "DeletedAt",
	ModelType: parent2Type,
}

func Parent2DeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[models.Parent2] {
	return orm.FieldCondition[models.Parent2, time.Time]{
		FieldIdentifier: Parent2DeletedAtField,
		Operator:        operator,
	}
}
func Parent2ParentParent(conditions ...orm.Condition[models.ParentParent]) orm.IJoinCondition[models.Parent2] {
	return orm.JoinCondition[models.Parent2, models.ParentParent]{
		Conditions:         conditions,
		RelationField:      "ParentParent",
		T1Field:            "ParentParentID",
		T1PreloadCondition: Parent2PreloadAttributes,
		T2Field:            "ID",
	}
}

var Parent2PreloadParentParent = Parent2ParentParent(ParentParentPreloadAttributes)
var Parent2ParentParentIdField = orm.FieldIdentifier[orm.UUID]{
	Field:     "ParentParentID",
	ModelType: parent2Type,
}

func Parent2ParentParentId(operator orm.Operator[orm.UUID]) orm.WhereCondition[models.Parent2] {
	return orm.FieldCondition[models.Parent2, orm.UUID]{
		FieldIdentifier: Parent2ParentParentIdField,
		Operator:        operator,
	}
}

var Parent2PreloadAttributes = orm.NewPreloadCondition[models.Parent2](Parent2IdField, Parent2CreatedAtField, Parent2UpdatedAtField, Parent2DeletedAtField, Parent2ParentParentIdField)
var Parent2PreloadRelations = []orm.Condition[models.Parent2]{Parent2PreloadParentParent}