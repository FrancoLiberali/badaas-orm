// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	overridereferencesinverse "github.com/ditrit/badaas-cli/cmd/gen/conditions/tests/overridereferencesinverse"
	cql "github.com/ditrit/badaas/orm/cql"
	model "github.com/ditrit/badaas/orm/model"
	"time"
)

func (computerConditions computerConditions) Processor(conditions ...cql.Condition[overridereferencesinverse.Processor]) cql.JoinCondition[overridereferencesinverse.Computer] {
	return cql.NewJoinCondition[overridereferencesinverse.Computer, overridereferencesinverse.Processor](conditions, "Processor", "Name", computerConditions.Preload(), "ComputerName")
}
func (computerConditions computerConditions) PreloadProcessor() cql.JoinCondition[overridereferencesinverse.Computer] {
	return computerConditions.Processor(Processor.Preload())
}

type computerConditions struct {
	ID        cql.Field[overridereferencesinverse.Computer, model.UUID]
	CreatedAt cql.Field[overridereferencesinverse.Computer, time.Time]
	UpdatedAt cql.Field[overridereferencesinverse.Computer, time.Time]
	DeletedAt cql.Field[overridereferencesinverse.Computer, time.Time]
	Name      cql.StringField[overridereferencesinverse.Computer]
}

var Computer = computerConditions{
	CreatedAt: cql.Field[overridereferencesinverse.Computer, time.Time]{Name: "CreatedAt"},
	DeletedAt: cql.Field[overridereferencesinverse.Computer, time.Time]{Name: "DeletedAt"},
	ID:        cql.Field[overridereferencesinverse.Computer, model.UUID]{Name: "ID"},
	Name:      cql.StringField[overridereferencesinverse.Computer]{Field: cql.Field[overridereferencesinverse.Computer, string]{Name: "Name"}},
	UpdatedAt: cql.Field[overridereferencesinverse.Computer, time.Time]{Name: "UpdatedAt"},
}

// Preload allows preloading the Computer when doing a query
func (computerConditions computerConditions) Preload() cql.Condition[overridereferencesinverse.Computer] {
	return cql.NewPreloadCondition[overridereferencesinverse.Computer](computerConditions.ID, computerConditions.CreatedAt, computerConditions.UpdatedAt, computerConditions.DeletedAt, computerConditions.Name)
}

// PreloadRelations allows preloading all the Computer's relation when doing a query
func (computerConditions computerConditions) PreloadRelations() []cql.Condition[overridereferencesinverse.Computer] {
	return []cql.Condition[overridereferencesinverse.Computer]{computerConditions.PreloadProcessor()}
}
