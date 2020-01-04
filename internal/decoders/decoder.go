package decoders

import "net/http"

type Decoder interface {
	DecodeRequest(*http.Request, interface{}) (bool, error)
}

type decoder struct {
	decoders []Decoder
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
