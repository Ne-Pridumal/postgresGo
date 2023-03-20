package handlers

import (
	"encoding/json"
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"net/http"
	"time"
)

type ticketsActions struct {
	store.Store
}

func (a *ticketsActions) Create() http.HandlerFunc {
	type request struct {
		Key           string `json:"key"`
		BookRef       int    `json:"bookRef"`
		PassengerId   int    `json:"passengerId"`
		PassengerName string `json:"passengerName"`
		ContactDate   string `json:"contactDate"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}

		date, err := time.Parse(time.RFC3339, req.ContactDate)
		if err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}

		t := &models.Ticket{
			Key:           req.Key,
			BookRef:       req.BookRef,
			PassengerId:   req.PassengerId,
			PassengerName: req.PassengerName,
			ContactDate:   date,
		}

		if _, err := a.Tickets().Create(t); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		respond(w, r, http.StatusCreated, t)
	}
}

func (a *ticketsActions) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tickets, err := a.Tickets().GetAll()
		if err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		respond(w, r, http.StatusOK, tickets)
	}
}

func (a *ticketsActions) Delete() http.HandlerFunc {
	type request struct {
		Key     string `json:"key"`
		BookRef int    `json:"bookRef"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}
		if err := a.Tickets().Delete(req.Key, req.BookRef); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		respond(w, r, http.StatusOK, "Успешный успех")
	}
}
