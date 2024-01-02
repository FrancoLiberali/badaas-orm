// Code generated by cql-gen v0.0.8, DO NOT EDIT.
package conditions

import (
	condition "github.com/FrancoLiberali/cql/condition"
	belongsto "github.com/FrancoLiberali/cql-gen/cmd/gen/conditions/tests/belongsto"
	model "github.com/FrancoLiberali/cql/model"
	"time"
)

func (ownedConditions ownedConditions) Owner(conditions ...condition.Condition[belongsto.Owner]) condition.JoinCondition[belongsto.Owned] {
	return condition.NewJoinCondition[belongsto.Owned, belongsto.Owner](conditions, "Owner", "OwnerID", ownedConditions.preload(), "ID", Owner.preload())
}

type ownedConditions struct {
	ID        condition.Field[belongsto.Owned, model.UUID]
	CreatedAt condition.Field[belongsto.Owned, time.Time]
	UpdatedAt condition.Field[belongsto.Owned, time.Time]
	DeletedAt condition.Field[belongsto.Owned, time.Time]
	OwnerID   condition.UpdatableField[belongsto.Owned, model.UUID]
}

var Owned = ownedConditions{
	CreatedAt: condition.NewField[belongsto.Owned, time.Time]("CreatedAt", "", ""),
	DeletedAt: condition.NewField[belongsto.Owned, time.Time]("DeletedAt", "", ""),
	ID:        condition.NewField[belongsto.Owned, model.UUID]("ID", "", ""),
	OwnerID:   condition.NewUpdatableField[belongsto.Owned, model.UUID]("OwnerID", "", ""),
	UpdatedAt: condition.NewField[belongsto.Owned, time.Time]("UpdatedAt", "", ""),
}

// Preload allows preloading the Owned when doing a query
func (ownedConditions ownedConditions) preload() condition.Condition[belongsto.Owned] {
	return condition.NewPreloadCondition[belongsto.Owned](ownedConditions.ID, ownedConditions.CreatedAt, ownedConditions.UpdatedAt, ownedConditions.DeletedAt, ownedConditions.OwnerID)
}
