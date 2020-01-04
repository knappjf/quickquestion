package thing_handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	jsonDecoder "github.com/knappjf/quickquestion/internal/decoders/json"
	"github.com/knappjf/quickquestion/internal/models"
	"github.com/knappjf/quickquestion/internal/repository"
	"go.uber.org/fx"
	"net/http"
	"strconv"
)

var Module = fx.Provide(New)

type ThingHandler interface {
	CreateThing(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetThing(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	DeleteThing(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateThing(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type Params struct {
	fx.In

	Repository repository.ThingRepository
	Decoder    jsonDecoder.JsonDecoder
}

func New(p Params) (ThingHandler, error) {
	return &thingHandler{
		repo:    p.Repository,
		decoder: p.Decoder,
	}, nil
}

type thingHandler struct {
	repo    repository.ThingRepository
	decoder jsonDecoder.JsonDecoder
}

func (h *thingHandler) CreateThing(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var thing models.Thing

	if err := h.decoder.DecodeRequest(r, &thing); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.repo.CreateThing(thing)

	if err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *thingHandler) GetThing(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idParam := params.ByName("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	thing, err := h.repo.GetThing(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	thingJson, err := json.Marshal(thing)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = w.Write(thingJson)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *thingHandler) DeleteThing(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idParam := params.ByName("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.repo.DeleteThing(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *thingHandler) UpdateThing(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idParam := params.ByName("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var thing models.Thing

	if err := h.decoder.DecodeRequest(r, &thing); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.repo.UpdateThing(id, thing); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
