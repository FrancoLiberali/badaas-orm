// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	orm "github.com/ditrit/badaas/orm"
	models "github.com/ditrit/badaas/testintegration/models"
	"reflect"
	"time"
)

var parent1Type = reflect.TypeOf(*new(models.Parent1))
var Parent1IdField = orm.FieldIdentifier[orm.UUID]{
	Field:     "ID",
	ModelType: parent1Type,
}

func Parent1Id(operator orm.Operator[orm.UUID]) orm.WhereCondition[models.Parent1] {
	return orm.FieldCondition[models.Parent1, orm.UUID]{
		FieldIdentifier: Parent1IdField,
		Operator:        operator,
	}
}

var Parent1CreatedAtField = orm.FieldIdentifier[time.Time]{
	Field:     "CreatedAt",
	ModelType: parent1Type,
}

func Parent1CreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[models.Parent1] {
	return orm.FieldCondition[models.Parent1, time.Time]{
		FieldIdentifier: Parent1CreatedAtField,
		Operator:        operator,
	}
}

var Parent1UpdatedAtField = orm.FieldIdentifier[time.Time]{
	Field:     "UpdatedAt",
	ModelType: parent1Type,
}

func Parent1UpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[models.Parent1] {
	return orm.FieldCondition[models.Parent1, time.Time]{
		FieldIdentifier: Parent1UpdatedAtField,
		Operator:        operator,
	}
}

var Parent1DeletedAtField = orm.FieldIdentifier[time.Time]{
	Field:     "DeletedAt",
	ModelType: parent1Type,
}

func Parent1DeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[models.Parent1] {
	return orm.FieldCondition[models.Parent1, time.Time]{
		FieldIdentifier: Parent1DeletedAtField,
		Operator:        operator,
	}
}
func Parent1ParentParent(conditions ...orm.Condition[models.ParentParent]) orm.IJoinCondition[models.Parent1] {
	return orm.JoinCondition[models.Parent1, models.ParentParent]{
		Conditions:         conditions,
		RelationField:      "ParentParent",
		T1Field:            "ParentParentID",
		T1PreloadCondition: Parent1PreloadAttributes,
		T2Field:            "ID",
	}
}

var Parent1PreloadParentParent = Parent1ParentParent(ParentParentPreloadAttributes)
var Parent1ParentParentIdField = orm.FieldIdentifier[orm.UUID]{
	Field:     "ParentParentID",
	ModelType: parent1Type,
}

func Parent1ParentParentId(operator orm.Operator[orm.UUID]) orm.WhereCondition[models.Parent1] {
	return orm.FieldCondition[models.Parent1, orm.UUID]{
		FieldIdentifier: Parent1ParentParentIdField,
		Operator:        operator,
	}
}

var Parent1PreloadAttributes = orm.NewPreloadCondition[models.Parent1](Parent1IdField, Parent1CreatedAtField, Parent1UpdatedAtField, Parent1DeletedAtField, Parent1ParentParentIdField)
var Parent1PreloadRelations = []orm.Condition[models.Parent1]{Parent1PreloadParentParent}