// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package overrideforeignkey

import preload "github.com/FrancoLiberali/cql/orm/preload"

func (m Bicycle) GetOwner() (*Person, error) {
	return preload.VerifyStructLoaded[Person](&m.Owner)
}
