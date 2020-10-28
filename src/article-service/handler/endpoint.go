package handler

import (
	"article-service/article"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func createOrUpdateArticle(h *Handler, w http.ResponseWriter, r *http.Request) {

	response := response{}
	var request article.ArticleDto
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Invalid create request: ", reqBody)
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		respondJSON(w, http.StatusBadRequest, response)
		return
	}

	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		log.Println("Error encountered while unmarshalling request body", err.Error())
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		respondJSON(w, http.StatusBadRequest, response)
		return
	}

	article, err := h.articles.Create(&request)
	if err != nil {
		log.Println("Error while creating article ", err)
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		respondJSON(w, http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusCreated
	response.Message = "Success"
	response.Data = article

	respondJSON(w, http.StatusCreated, response)
}

func getByID(h *Handler, w http.ResponseWriter, r *http.Request) {
	response := response{}
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("Invalid Id: ", id)
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		respondJSON(w, http.StatusBadRequest, response)
		return
	}
	article, err := h.articles.GetByID(idInt)
	if err != nil {
		log.Println("Error while getting article by id: ", err)
		response.Status = http.StatusNotFound
		response.Message = err.Error()
		respondJSON(w, http.StatusNotFound, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = article
	respondJSON(w, http.StatusOK, response)
}

func getAll(h *Handler, w http.ResponseWriter, r *http.Request) {
	response := response{}
	articles, err := h.articles.GetAll()
	if err != nil {
		log.Println("Error while getting articles: ", err)
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		respondJSON(w, http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = articles
	respondJSON(w, http.StatusOK, response)
}

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {

	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

type response struct {
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
