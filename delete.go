package cql

import (
	"gorm.io/gorm"

	"github.com/FrancoLiberali/cql/condition"
	"github.com/FrancoLiberali/cql/model"
)

// Create a Delete to which the conditions are applied inside transaction tx
func Delete[T model.Model](tx *gorm.DB, conditions ...condition.Condition[T]) *condition.Delete[T] {
	return condition.NewDelete(tx, conditions...)
}