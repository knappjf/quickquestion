//go:generate mockgen -source json.go -destination mocks/json.go
package json

import (
	"encoding/json"
	"net/http"
	"strings"
)

func New() *jsonDecoder {
	return &jsonDecoder{}
}

type jsonDecoder struct{}

func (jd *jsonDecoder) DecodeRequest(r *http.Request, dest interface{}) (bool, error) {
	contentTypeParts := strings.Split(r.Header.Get("Content-Type"), ";")

	if contentTypeParts[0] == "application/json" {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(dest)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}
