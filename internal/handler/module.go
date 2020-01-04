package handler

import (
	"github.com/knappjf/quickquestion/internal/handler/thing_handler"
	"go.uber.org/fx"
)

var Module = fx.Options(
	thing_handler.Module,
)
