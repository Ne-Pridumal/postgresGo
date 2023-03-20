package handlers

import (
	"encoding/json"
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"net/http"
	"time"
)

type flightActions struct {
	store.Store
}

func (a *flightActions) Create() http.HandlerFunc {
	type request struct {
		FlightKey          int       `json:"key"`
		ScheduledDeparture time.Time `json:"scheduledDeparture"`
		ScheduledArrival   time.Time `json:"scheduledArrival"`
		DepartureAirport   int       `json:"departureAirport"`
		ArrivalAirport     int       `json:"arrivalAirport"`
		Status             string    `json:"status"`
		AircraftKey        int       `json:"aircraftKey"`
		ActualDeparture    time.Time `json:"actualDeparture"`
		ActualArrival      time.Time `json:"actualArrival"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}

		f := &models.Flight{
			FlightKey:          req.FlightKey,
			ScheduledDeparture: req.ScheduledDeparture,
			ScheduledArrival:   req.ScheduledArrival,
			DepartureAirport:   req.DepartureAirport,
			ArrivalAirport:     req.ArrivalAirport,
			Status:             req.Status,
			AircraftKey:        req.AircraftKey,
			ActualDeparture:    req.ActualDeparture,
			ActualArrival:      req.ActualArrival,
		}

		if _, err := a.Flights().Create(f); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		respond(w, r, http.StatusCreated, f)
	}
}

func (a *flightActions) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flights, err := a.Flights().GetAll()
		if err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		respond(w, r, http.StatusOK, flights)
	}
}

func (a *flightActions) Delete() http.HandlerFunc {
	type request struct {
		FlightKey          int       `json:"key"`
		ScheduledDeparture time.Time `json:"scheduledDeparture"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}
		if err := a.Flights().Delete(req.FlightKey, req.ScheduledDeparture); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		// придумать нормальный ответ
		respond(w, r, http.StatusOK, "Успешный успех")
	}
}
