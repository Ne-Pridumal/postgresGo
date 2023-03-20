package handlers

import (
	"encoding/json"
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"net/http"
	"time"
)

type boardingPassesActions struct {
	store.Store
}

func (a *boardingPassesActions) Create() http.HandlerFunc {
	type request struct {
		FlightKey          int       `json:"flightKey"`
		ScheduledDeparture time.Time `json:"scheduledDeparture"`
		TicketKey          string    `json:"ticketKey"`
		BoardingNo         string    `json:"boardingKey"`
		SeatNo             string    `json:"seatKey"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}

		bp := &models.BoardingPass{
			FlightKey:          req.FlightKey,
			ScheduledDeparture: req.ScheduledDeparture,
			TicketKey:          req.TicketKey,
			SeatNo:             req.SeatNo,
			BoardingNo:         req.BoardingNo,
		}

		if _, err := a.BoardingPasses().Create(bp); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		respond(w, r, http.StatusCreated, bp)
	}
}

func (a *boardingPassesActions) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		boardingPasses, err := a.BoardingPasses().GetAll()
		if err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		respond(w, r, http.StatusOK, boardingPasses)
	}
}

func (a *boardingPassesActions) Delete() http.HandlerFunc {
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
		if err := a.BoardingPasses().Delete(req.FlightKey, req.TicketKey, req.ScheduledDeparture); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		// придумать нормальный ответ
		respond(w, r, http.StatusOK, "Успешный успех")
	}
}
