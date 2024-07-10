package handler

import (
	"encoding/json"
	"errors"
	repoErrors "github.com/PetkovaDiana/shop/internal/repository/errors"
	domainModels "github.com/PetkovaDiana/shop/internal/service/models"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var data domainModels.CreateClient
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "bla", http.StatusBadRequest)
		return
	}

	err := h.services.Authorization.CreateClient(r.Context(), data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var data domainModels.AuthClient

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	token, err := h.services.Authorization.AuthClient(r.Context(), data)
	if err != nil {
		if errors.As(err, &repoErrors.ErrClientNotFound{}) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	ck := http.Cookie{
		Name:     "token",
		Domain:   "",
		Path:     basePath,
		HttpOnly: true,
		Secure:   true,
	}

	ck.Value = token

	http.SetCookie(w, &ck)
	w.WriteHeader(http.StatusOK)

}
