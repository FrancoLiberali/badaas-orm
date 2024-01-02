// Code generated by cql-gen v0.0.10, DO NOT EDIT.
package hasmany

import preload "github.com/FrancoLiberali/cql/preload"

func (m Company) GetSellers() ([]Seller, error) {
	return preload.VerifyCollectionLoaded[Seller](m.Sellers)
}
func (m Seller) GetCompany() (*Company, error) {
	return preload.VerifyPointerLoaded[Company](m.CompanyID, m.Company)
}
