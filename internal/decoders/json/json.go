//go:generate mockgen -source json.go -destination mocks/json.go
package json

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type JsonDecoder interface {
	DecodeRequest(*http.Request, interface{}) error //TODO: Make this a generic interface and implement msgpack decoder
}

func New() JsonDecoder {
	return &jsonDecoder{}
}

type jsonDecoder struct{}

// TODO: Rewrite this to support multiple types of decoders
func (jd *jsonDecoder) DecodeRequest(r *http.Request, dest interface{}) error {
	contentTypeParts := strings.Split(r.Header.Get("Content-Type"), ";")
	if contentTypeParts[0] == "application/json" {
		decoder := json.NewDecoder(r.Body)
		return decoder.Decode(dest)
	}

	return errors.New("Unsupported media type")
}
