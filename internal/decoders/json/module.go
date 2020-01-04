package json

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(New)
