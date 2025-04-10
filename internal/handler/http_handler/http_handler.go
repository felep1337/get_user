package http_handler

import (
	"encoding/json"
	"http_auth/internal/use_case"
	"net/http"
	"strconv"
)

type Handler struct {
	Data *use_case.UserRepo
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	UserUUID := r.URL.Query().Get("id")
	ID, err := strconv.Atoi(UserUUID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	user, err := h.Data.FindByUUID(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}
