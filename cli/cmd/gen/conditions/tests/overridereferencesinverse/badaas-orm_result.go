// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package overridereferencesinverse

import preload "github.com/FrancoLiberali/cql/orm/preload"

func (m Computer) GetProcessor() (*Processor, error) {
	return preload.VerifyStructLoaded[Processor](&m.Processor)
}
