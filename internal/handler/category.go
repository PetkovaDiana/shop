package handler

import (
	"encoding/json"
	domainModels "github.com/PetkovaDiana/shop/internal/service/models"
	"net/http"
)

func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data domainModels.GetCategoriesFilter
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "bla bla bla", http.StatusBadRequest)
		return
	}

	// Вызываем метод GetCategory с полученными ID категорий
	categories, err := h.services.Category.GetCategory(r.Context(), data)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Кодируем результат в JSON и отправляем обратно клиенту
	w.Header().Set("Content-Type", "text/json")
	json.NewEncoder(w).Encode(categories)
}
