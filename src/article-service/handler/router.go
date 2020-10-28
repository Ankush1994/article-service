package handler

import (
	"log"
	"net/http"
)

// RequestHandlerFunction is a type func that handles all the HTTP requests
type RequestHandlerFunction func(h *Handler, w http.ResponseWriter, r *http.Request)

func (h *Handler) setRouters() {

	log.Println("Handler setting routes..")

	h.router.Handle("/articles", h.handleRequest(createOrUpdateArticle)).Methods("POST")
	h.router.Handle("/articles", h.handleRequest(getAll)).Methods("GET")
	h.router.Handle("/articles/{id}", h.handleRequest(getByID)).Methods("GET")
}

func (h *Handler) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(h, w, r)
	}
}
