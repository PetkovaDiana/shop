package handler

import (
	"fmt"
	"github.com/PetkovaDiana/shop/internal/service"
	"net/http"
	"time"
)

const (
	basePath = "/api"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mx := http.NewServeMux()

	mx.HandleFunc(fmt.Sprintf("%s/get-categories", basePath), h.baseHandler(h.GetCategories)) // Get

	return mx
}

func (h *Handler) baseHandler(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now())
		fmt.Println(r.Method)

		handlerFunc(w, r)
	}
}
