package goembedded

import "github.com/FrancoLiberali/cql/orm/model"

type ToBeEmbedded struct {
	Int int
}

type GoEmbedded struct {
	model.UIntModel

	Int int
	ToBeEmbedded
}
