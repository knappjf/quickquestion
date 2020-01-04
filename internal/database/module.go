package database

import (
	"github.com/gocraft/dbr/v2"
	"go.uber.org/fx"
)

var Module = fx.Provide(New)

func New() (dbr.SessionRunner, error) {
	conn, err := dbr.Open("sqlite3", "things.db", nil) // TODO: extract database name to config
	if err != nil {
		return nil, err
	}

	return conn.NewSession(nil), nil
}
