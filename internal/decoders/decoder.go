package decoders

import (
	"github.com/knappjf/quickquestion/internal/decoders/interfaces"
	"net/http"
)

type decoder struct {
	decoders []interfaces.Decoder
}

func (d *decoder) DecodeRequest(request *http.Request, destination interface{}) (bool, error) {
	var decoded bool
	var err error

	for _, dec := range d.decoders {
		decoded, err = dec.DecodeRequest(request, destination)
		if err != nil {
			return false, err
		}

		if decoded {
			return decoded, nil
		}
	}

	return false, ErrUnsupportedMediaType
}
