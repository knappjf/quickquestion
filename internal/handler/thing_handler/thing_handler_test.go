package thing_handler

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	jsonDecoder "github.com/knappjf/quickquestion/internal/decoders/json"
	"github.com/knappjf/quickquestion/internal/models"
	"github.com/knappjf/quickquestion/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_CreateThing(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	thing := models.Thing{
		Name:        "foo",
		Description: "this is a foo",
		Enabled:     true,
	}

	mockThingRepository := mock_repository.NewMockThingRepository(controller)
	mockThingRepository.EXPECT().CreateThing(thing).Return(nil)

	body, err := json.Marshal(thing)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	decoder := jsonDecoder.New()

	req, err := http.NewRequest("POST", "/v1/things", bytes.NewReader(body))
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.ResponseRecorder{}

	h := thingHandler{
		repo:    mockThingRepository,
		decoder: decoder,
	}

	h.CreateThing(&resp, req, nil)
}

func TestHandler_GetThing(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	thing := models.Thing{
		Name:        "foo",
		Description: "this is a foo",
		Enabled:     true,
	}

	mockThingRepository := mock_repository.NewMockThingRepository(controller)
	mockThingRepository.EXPECT().GetThing(123).Return(thing, nil)

	// TODO: look into using httprequest.NewRequest for this
	req, err := http.NewRequest("GET", "/v1/things/123", nil)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	resp := httptest.NewRecorder()
	params := httprouter.Params{
		httprouter.Param{
			Key:   "id",
			Value: "123",
		},
	}

	h := thingHandler{
		repo: mockThingRepository,
	}
	h.GetThing(resp, req, params)

	assert.Equal(t, http.StatusOK, resp.Code)

	var actual models.Thing
	assert.Nil(t, json.Unmarshal(resp.Body.Bytes(), &actual))
	assert.Equal(t, thing, actual)
}

func TestHandler_DeleteThing(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockThingRepository := mock_repository.NewMockThingRepository(controller)
	mockThingRepository.EXPECT().DeleteThing(123).Return(nil)

	req, err := http.NewRequest("DELETE", "/v1/things/123", nil)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	resp := httptest.NewRecorder()
	params := httprouter.Params{
		httprouter.Param{
			Key:   "id",
			Value: "123",
		},
	}

	h := thingHandler{
		repo: mockThingRepository,
	}

	h.DeleteThing(resp, req, params)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestHandler_UpdateThing(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	thing := models.Thing{
		Name:        "foo",
		Description: "this is a foo",
		Enabled:     true,
	}

	mockThingRepository := mock_repository.NewMockThingRepository(controller)
	mockThingRepository.EXPECT().UpdateThing(123, thing).Return(nil)

	body, err := json.Marshal(thing)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	decoder := jsonDecoder.New()

	req, err := http.NewRequest("POST", "/v1/things/123", bytes.NewReader(body))
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	h := thingHandler{
		repo:    mockThingRepository,
		decoder: decoder,
	}

	params := httprouter.Params{
		httprouter.Param{
			Key:   "id",
			Value: "123",
		},
	}

	h.UpdateThing(resp, req, params)

	assert.Equal(t, http.StatusOK, resp.Code)
}
