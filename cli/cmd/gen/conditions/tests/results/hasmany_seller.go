// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	hasmany "github.com/ditrit/badaas-orm/cli/cmd/gen/conditions/tests/hasmany"
	orm "github.com/ditrit/badaas/orm"
	"time"
)

func SellerId(operator orm.Operator[orm.UUID]) orm.WhereCondition[hasmany.Seller] {
	return orm.FieldCondition[hasmany.Seller, orm.UUID]{
		Field:    "ID",
		Operator: operator,
	}
}
func SellerCreatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[hasmany.Seller] {
	return orm.FieldCondition[hasmany.Seller, time.Time]{
		Field:    "CreatedAt",
		Operator: operator,
	}
}
func SellerUpdatedAt(operator orm.Operator[time.Time]) orm.WhereCondition[hasmany.Seller] {
	return orm.FieldCondition[hasmany.Seller, time.Time]{
		Field:    "UpdatedAt",
		Operator: operator,
	}
}
func SellerDeletedAt(operator orm.Operator[time.Time]) orm.WhereCondition[hasmany.Seller] {
	return orm.FieldCondition[hasmany.Seller, time.Time]{
		Field:    "DeletedAt",
		Operator: operator,
	}
}
func SellerCompanyId(operator orm.Operator[orm.UUID]) orm.WhereCondition[hasmany.Seller] {
	return orm.FieldCondition[hasmany.Seller, orm.UUID]{
		Field:    "CompanyID",
		Operator: operator,
	}
}