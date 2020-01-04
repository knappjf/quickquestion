package database

import (
	"github.com/gocraft/dbr/v2"
	"github.com/knappjf/quickquestion/internal/config"
	"go.uber.org/fx"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	Config config.Config
}

func New(p Params) (dbr.SessionRunner, error) {
	conn, err := dbr.Open("sqlite3", p.Config.Database, nil)
	if err != nil {
		return nil, err
	}

	return conn.NewSession(nil), nil
}
