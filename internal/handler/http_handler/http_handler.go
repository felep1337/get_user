package http_handler

import (
	"encoding/json"
	"http_auth/internal/domain"
	"net/http"
	"strconv"
)

type userUseCase interface {
	FindByUUID(ID int) (*domain.User, error)
}

type Handler struct {
	Data userUseCase
}

func NewHandler(data userUseCase) *Handler {
	return &Handler{Data: data}
}

func (h Handler) FindUserByID(w http.ResponseWriter, r *http.Request) {
	UserUUID := r.URL.Query().Get("id")
	ID, err := strconv.Atoi(UserUUID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	user, err := h.Data.FindByUUID(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	// Тоже полезная штука, чтобы клиент знал, что он получает JSON, а не просто текст
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
