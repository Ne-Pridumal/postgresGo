package handlers

import (
	"encoding/json"
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"net/http"
)

type aircraftActions struct {
	store.Store
}

func (a *aircraftActions) Create() http.HandlerFunc {
	type request struct {
		Key   int    `json:"key"`
		Model string `json:"model"`
		Range int    `json:"range"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}

		ac := &models.Aircraft{
			Key:   req.Key,
			Model: req.Model,
			Range: req.Range,
		}

		if _, err := a.Aircrafts().Create(ac); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		respond(w, r, http.StatusCreated, ac)
	}
}

func (a *aircraftActions) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		aircrafts, err := a.Aircrafts().GetAll()
		if err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		respond(w, r, http.StatusOK, aircrafts)
	}
}

func (a *aircraftActions) Delete() http.HandlerFunc {
	type request struct {
		Key int `json:"key"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}

		if err := a.Aircrafts().Delete(req.Key); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		// придумать нормальный ответ
		respond(w, r, http.StatusOK, "Успешный успех")
	}
}
