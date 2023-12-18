package unsafe

import (
	"fmt"

	"github.com/FrancoLiberali/cql/orm/cql"
	"github.com/FrancoLiberali/cql/orm/model"
)

// Condition that can be used to express conditions that are not supported (yet?) by badaas-orm
// Example: table1.columnX = table2.columnY
type unsafeCondition[T model.Model] struct {
	SQLCondition string
	Values       []any
}

func (unsafeCondition unsafeCondition[T]) InterfaceVerificationMethod(_ T) {
	// This method is necessary to get the compiler to verify
	// that an object is of type Condition[T]
}

func (unsafeCondition unsafeCondition[T]) ApplyTo(query *cql.GormQuery, table cql.Table) error {
	return cql.ApplyWhereCondition[T](unsafeCondition, query, table)
}

func (unsafeCondition unsafeCondition[T]) GetSQL(_ *cql.GormQuery, table cql.Table) (string, []any, error) {
	return fmt.Sprintf(
		unsafeCondition.SQLCondition,
		table.Alias,
	), unsafeCondition.Values, nil
}

func (unsafeCondition unsafeCondition[T]) AffectsDeletedAt() bool {
	return false
}

// Condition that can be used to express conditions that are not supported (yet?) by badaas-orm
// Example: table1.columnX = table2.columnY
func NewCondition[T model.Model](sqlCondition string, values ...any) cql.Condition[T] {
	return unsafeCondition[T]{
		SQLCondition: sqlCondition,
		Values:       values,
	}
}
