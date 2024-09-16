package handler

import (
	"fmt"
	"github.com/PetkovaDiana/shop/internal/service"
	"net/http"
	"time"
)

const (
	basePath = "/api"
	authApi  = basePath + "/auth"
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
	mx.HandleFunc(fmt.Sprintf("%s/get-products", basePath), h.baseHandler(h.GetProducts))     // Get
	mx.HandleFunc(fmt.Sprintf("%s/sing-in", authApi), h.baseHandler(h.signIn))                //Post
	mx.HandleFunc(fmt.Sprintf("%s/sing-up", authApi), h.baseHandler(h.signUp))                // Post
	mx.HandleFunc(fmt.Sprintf("%s/basket", authApi), h.baseHandler(h.basket))                 // Post

	return mx
}

func (h *Handler) baseHandler(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now())
		fmt.Println(r.Method)

		handlerFunc(w, r)
	}
}
