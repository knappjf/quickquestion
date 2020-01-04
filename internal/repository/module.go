package repository

import (
	"github.com/gocraft/dbr/v2"
	"go.uber.org/fx"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	SessionRunner dbr.SessionRunner
}

func New(p Params) (ThingRepository, error) {
	return NewThingRepository(p.SessionRunner), nil
}
