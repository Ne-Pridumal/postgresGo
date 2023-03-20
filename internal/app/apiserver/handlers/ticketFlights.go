package handlers

import (
	"encoding/json"
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"net/http"
	"time"
)

type ticketFlightsActions struct {
	store.Store
}

func (a *ticketFlightsActions) Create() http.HandlerFunc {
	type request struct {
		FlightKey          int       `json:"flightKey"`
		ScheduledDeparture time.Time `json:"scheduledDeparture"`
		TicketKey          string    `json:"ticketKey"`
		FareCondition      string    `json:"fareCondition"`
		Amount             int       `json:"amount"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}

		tf := &models.TicketFlights{
			FlightKey:          req.FlightKey,
			ScheduledDeparture: req.ScheduledDeparture,
			Amount:             req.Amount,
			TicketKey:          req.TicketKey,
			FareCondition:      req.FareCondition,
		}

		if _, err := a.TicketFlights().Create(tf); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		respond(w, r, http.StatusCreated, tf)
	}
}

func (a *ticketFlightsActions) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flights, err := a.TicketFlights().GetAll()
		if err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		respond(w, r, http.StatusOK, flights)
	}
}

func (a *ticketFlightsActions) Delete() http.HandlerFunc {
	type request struct {
		FlightKey          int       `json:"flightKey"`
		ScheduledDeparture time.Time `json:"scheduledDeparture"`
		TicketKey          string    `json:"ticketKey"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}
		if err := a.TicketFlights().Delete(req.FlightKey, req.TicketKey, req.ScheduledDeparture); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		// придумать нормальный ответ
		respond(w, r, http.StatusOK, "Успешный успех")
	}
}
