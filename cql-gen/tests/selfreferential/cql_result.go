// Code generated by cql-gen v0.0.3, DO NOT EDIT.
package selfreferential

import preload "github.com/FrancoLiberali/cql/preload"

func (m Employee) GetBoss() (*Employee, error) {
	return preload.VerifyPointerLoaded[Employee](m.BossID, m.Boss)
}