package interfaces

import "net/http"

type Decoder interface {
	DecodeRequest(*http.Request, interface{}) (bool, error)
}
