package decoders

import (
	"github.com/knappjf/quickquestion/internal/decoders/json"
	"go.uber.org/fx"
)

var Module = fx.Options(
	json.Module,
)
