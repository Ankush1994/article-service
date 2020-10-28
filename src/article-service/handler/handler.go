// Package handler provides a HTTP handler implementation for the API endpoints
package handler

import (
	"article-service/article"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Handler defines the initalized instance of the Billing service
type Handler struct {
	router   *mux.Router
	address  string
	port     int64
	articles article.IArticle
}

// New instantiates the handler
func New(address string, port int64, articles *article.Articles) (*Handler, error) {

	addr := fmt.Sprintf("%s:%d", address, port)
	h := &Handler{
		articles: articles,
	}
	h.address = address
	h.port = port
	h.router = mux.NewRouter()
	h.setRouters()
	log.Printf("Starting HTTP Handler at Address - %s", addr)
	return h, nil
}

// Start uses the handler instance to listen and serve the incoming requests
func (h *Handler) Start() {

	host := fmt.Sprintf("%s:%d", h.address, h.port)

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"x-access-token", "Content-Type", "Access-Control-Allow-Credentials", "Access-Control-Allow-Origin"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})

	err := http.ListenAndServe(host, handlers.CORS(originsOk,
		methodsOk, headersOk)(h.router))
	if err != nil {
		log.Fatalf("Error running ListenAndServe %v ", err)
	}
}
