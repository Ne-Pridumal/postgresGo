package handlers

import (
	"encoding/json"
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"net/http"
)

type seatsActions struct {
	store.Store
}

func (a *seatsActions) Create() http.HandlerFunc {
	type request struct {
		Number        string `json:"key"`
		AircraftKey   int    `json:"aircraftKey"`
		FareCondition string `json:"fareCondition"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}

		s := &models.Seat{
			AircraftKey:    req.AircraftKey,
			Number:         req.Number,
			FareConditions: req.FareCondition,
		}

		if _, err := a.Seats().Create(s); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		respond(w, r, http.StatusCreated, s)
	}
}

func (a *seatsActions) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookings, err := a.Seats().GetAll()
		if err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		respond(w, r, http.StatusOK, bookings)
	}
}

func (a *seatsActions) Delete() http.HandlerFunc {
	type request struct {
		Number      string `json:"key"`
		AircraftKey int    `json:"aircraftKey"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}

		if err := a.Seats().Delete(req.AircraftKey, req.Number); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		// придумать нормальный ответ
		respond(w, r, http.StatusOK, "Успешный успех")
	}
}
