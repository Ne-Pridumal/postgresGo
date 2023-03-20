package handlers

import (
	"encoding/json"
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"net/http"
)

type airportActions struct {
	store.Store
}

func (a *airportActions) Create() http.HandlerFunc {
	type request struct {
		Key   int    `json:"key"`
		Name  string `json:"name"`
		City  string `json:"city"`
		Coord string `json:"coordinates"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}

		ap := &models.Airport{
			Key:         req.Key,
			Name:        req.Name,
			City:        req.City,
			Coordinates: req.Coord,
		}

		if _, err := a.Airports().Create(ap); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		respond(w, r, http.StatusCreated, ap)
	}
}

func (a *airportActions) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		aircrafts, err := a.Airports().GetAll()
		if err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		respond(w, r, http.StatusOK, aircrafts)
	}
}

func (a *airportActions) Delete() http.HandlerFunc {
	type request struct {
		Key int `json:"key"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}

		if err := a.Airports().Delete(req.Key); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		// придумать нормальный ответ
		respond(w, r, http.StatusOK, "Успешный успех")
	}
}
