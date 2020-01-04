package main

import (
	"github.com/knappjf/quickquestion/internal/database"
	jsonDecoder "github.com/knappjf/quickquestion/internal/decoders/json"
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
		handler.Module,
		database.Module,
		repository.Module,
		jsonDecoder.Module,
		fx.Invoke(routes.Register),
	)
}
