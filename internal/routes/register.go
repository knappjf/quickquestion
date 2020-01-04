package routes

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/knappjf/quickquestion/internal/config"
	"github.com/knappjf/quickquestion/internal/handler/thing_handler"
	"go.uber.org/fx"
	"net/http"
)

type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
	Handler   thing_handler.ThingHandler
	Config    config.Config
}

func Register(p Params) {
	router := httprouter.New()
	router.Handle("GET", "/v1/things/:id", p.Handler.GetThing)
	router.Handle("POST", "/v1/things", p.Handler.CreateThing)
	router.Handle("POST", "/v1/things/:id", p.Handler.UpdateThing)
	router.Handle("DELETE", "/v1/things/:id", p.Handler.DeleteThing)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", p.Config.Address, p.Config.Port),
		Handler: router,
	}

	p.Lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return server.Shutdown(ctx)
			},
		})
}
