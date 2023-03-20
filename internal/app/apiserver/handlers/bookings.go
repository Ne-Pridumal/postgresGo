package handlers

import (
	"encoding/json"
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"net/http"
	"time"
)

type bookingActions struct {
	store.Store
}

func (a *bookingActions) Create() http.HandlerFunc {
	type request struct {
		BookRef     int    `json:"ref"`
		BookDate    string `json:"date"`
		TotalAmount int    `json:"amount"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}

		date, err := time.Parse(time.RFC3339, req.BookDate)
		if err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}

		b := &models.Booking{
			BookRef:     req.BookRef,
			BookDate:    date,
			TotalAmount: req.TotalAmount,
		}

		if _, err := a.Bookings().Create(b); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		respond(w, r, http.StatusCreated, b)
	}
}

func (a *bookingActions) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookings, err := a.Bookings().GetAll()
		if err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		respond(w, r, http.StatusOK, bookings)
	}
}

func (a *bookingActions) Delete() http.HandlerFunc {
	type request struct {
		BookRef int `json:"ref"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResp(w, r, http.StatusBadRequest, err)
			return
		}

		if err := a.Bookings().Delete(req.BookRef); err != nil {
			errorResp(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		// придумать нормальный ответ
		respond(w, r, http.StatusOK, "Успешный успех")
	}
}
