package main

import (
	"github.com/knappjf/quickquestion/internal/config"
	"github.com/knappjf/quickquestion/internal/database"
	"github.com/knappjf/quickquestion/internal/decoders"
	"github.com/knappjf/quickquestion/internal/handler"
	"github.com/knappjf/quickquestion/internal/repository"
	"github.com/knappjf/quickquestion/internal/routes"
	"go.uber.org/fx"
)

func main() {
	fx.New(opts()).Run()
}

func opts() fx.Option {
	return fx.Options(
		config.Module,
		decoders.Module,
		handler.Module,
		database.Module,
		repository.Module,
		fx.Invoke(routes.Register),
	)
}
