package handler

import (
	"article-service/article"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var articlDto = article.ArticleDto{
	ID:      1,
	Author:  "Test Author",
	Content: "Test Content",
	Title:   "Test Title",
}

func Test_createArticle_success(t *testing.T) {
	articlesMock, handler := InjectHandlerDependenciesMock()
	articlesMock.On("Create").Return(&articlDto, nil)

	requestBody, _ := json.Marshal(map[string]string{
		"author":  "Test Author",
		"content": "Test Content",
		"title":   "Test Title",
	})
	request, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(requestBody))
	responseRecorder := httptest.NewRecorder()
	createOrUpdateArticle(handler, responseRecorder, request)
	if http.StatusCreated != responseRecorder.Code {
		t.Error("Invalid response code")
	}
}

func Test_GetById_failure(t *testing.T) {
	articlesMock, handler := InjectHandlerDependenciesMock()
	articlesMock.On("GetByID").Return(&articlDto, nil)
	request, _ := http.NewRequest("GET", "/articles/asd", nil)
	responseRecorder := httptest.NewRecorder()
	getByID(handler, responseRecorder, request)
	if http.StatusBadRequest != responseRecorder.Code {
		t.Error("Invalid response code")
	}
}

func Test_GetAll_Success(t *testing.T) {
	articlesMock, handler := InjectHandlerDependenciesMock()
	articlesMock.On("GetAll").Return([]*article.ArticleDto{&articlDto}, nil)

	request, _ := http.NewRequest("GET", "/articles", nil)
	responseRecorder := httptest.NewRecorder()
	getAll(handler, responseRecorder, request)
	if http.StatusOK != responseRecorder.Code {
		t.Error("Invalid response code")
	}
}
