package json

import (
	"encoding/json"
	"github.com/knappjf/quickquestion/internal/decoders/interfaces"
	"net/http"
	"strings"
)

func New() interfaces.Decoder {
	return jsonDecoder{}
}

type jsonDecoder struct{}

func (jd jsonDecoder) DecodeRequest(r *http.Request, dest interface{}) (bool, error) {
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
