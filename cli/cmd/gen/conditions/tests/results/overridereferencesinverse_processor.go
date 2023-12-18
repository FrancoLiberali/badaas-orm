// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	overridereferencesinverse "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/overridereferencesinverse"
	cql "github.com/ditrit/badaas/orm/cql"
	model "github.com/ditrit/badaas/orm/model"
	"time"
)

type processorConditions struct {
	ID           cql.Field[overridereferencesinverse.Processor, model.UUID]
	CreatedAt    cql.Field[overridereferencesinverse.Processor, time.Time]
	UpdatedAt    cql.Field[overridereferencesinverse.Processor, time.Time]
	DeletedAt    cql.Field[overridereferencesinverse.Processor, time.Time]
	ComputerName cql.StringField[overridereferencesinverse.Processor]
}

var Processor = processorConditions{
	ComputerName: cql.StringField[overridereferencesinverse.Processor]{Field: cql.Field[overridereferencesinverse.Processor, string]{Name: "ComputerName"}},
	CreatedAt:    cql.Field[overridereferencesinverse.Processor, time.Time]{Name: "CreatedAt"},
	DeletedAt:    cql.Field[overridereferencesinverse.Processor, time.Time]{Name: "DeletedAt"},
	ID:           cql.Field[overridereferencesinverse.Processor, model.UUID]{Name: "ID"},
	UpdatedAt:    cql.Field[overridereferencesinverse.Processor, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Processor when doing a query
func (processorConditions processorConditions) Preload() cql.Condition[overridereferencesinverse.Processor] {
	return cql.NewPreloadCondition[overridereferencesinverse.Processor](processorConditions.ID, processorConditions.CreatedAt, processorConditions.UpdatedAt, processorConditions.DeletedAt, processorConditions.ComputerName)
}
