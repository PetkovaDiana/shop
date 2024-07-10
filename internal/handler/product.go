package handler

import (
	"encoding/json"
	domainModels "github.com/PetkovaDiana/shop/internal/service/models"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data domainModels.GetProductsFilter
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "ops ", http.StatusBadRequest)
		return
	}

	products, err := h.services.Product.GetProduct(r.Context(), data)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/json")
	json.NewEncoder(w).Encode(products)
}
