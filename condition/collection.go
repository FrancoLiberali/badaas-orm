package condition

import (
	"github.com/FrancoLiberali/cql/model"
)

type Collection[TObject model.Model, TAttribute model.Model] struct {
	name    string
	t1Field string
	t2Field string
}

// Preload collection of models
//
// nestedPreloads can be used to preload relations of the models inside the collection
func (collection Collection[TObject, TAttribute]) Preload(nestedPreloads ...JoinCondition[TAttribute]) Condition[TObject] {
	return NewCollectionPreloadCondition[TObject, TAttribute](collection.name, nestedPreloads)
}

// Any generates a condition that is true if at least one model in the collection fulfills the conditions
func (collection Collection[TObject, TAttribute]) Any(conditions ...WhereCondition[TAttribute]) WhereCondition[TObject] {
	return existsCondition[TObject, TAttribute]{
		Conditions:    conditions,
		RelationField: collection.name,
		T1Field:       collection.t1Field,
		T2Field:       collection.t2Field,
	}
}

func NewCollection[TObject model.Model, TAttribute model.Model](name, t1Field, t2Field string) Collection[TObject, TAttribute] {
	return Collection[TObject, TAttribute]{
		name:    name,
		t1Field: t1Field,
		t2Field: t2Field,
	}
}
