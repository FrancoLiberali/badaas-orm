// Code generated by cql-cli v0.0.0, DO NOT EDIT.
package overrideforeignkeyinverse

import preload "github.com/FrancoLiberali/cql/preload"

func (m User) GetCreditCard() (*CreditCard, error) {
	return preload.VerifyStructLoaded[CreditCard](&m.CreditCard)
}