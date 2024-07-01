package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	category, err := h.services.Category.GetAllCategory()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(category)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to encode category", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
