package handlers

import (
	"ne-pridumal/go-postgress/internal/app/store"
	"net/http"
)

type Actions interface {
	GetAll() http.HandlerFunc
	Create() http.HandlerFunc
	DeleteById(int) http.HandlerFunc
}

type Handlers struct {
	store    store.Store
	bookings Actions
}

func NewHandler(store store.Store) *Handlers {
	return &Handlers{
		store: store,
	}
}

func (h *Handlers) Bookings() Actions {
	if h.bookings != nil {
		return h.bookings
	}
	h.bookings = &BookingActions{
		Store: h.store,
	}
	return h.bookings
}
