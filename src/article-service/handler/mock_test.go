package handler

import (
	"article-service/article"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

type ArticlesMock struct {
	mock.Mock
}

func (a *ArticlesMock) Create(request *article.ArticleDto) (*article.ArticleDto, error) {
	args := a.Called()
	return args.Get(0).(*article.ArticleDto), args.Error(1)
}

func (a *ArticlesMock) GetByID(id int64) (*article.ArticleDto, error) {
	args := a.Called()
	return args.Get(0).(*article.ArticleDto), args.Error(1)
}

func (a *ArticlesMock) GetAll() ([]*article.ArticleDto, error) {
	args := a.Called()
	return args.Get(0).([]*article.ArticleDto), args.Error(1)
}

func InjectHandlerDependenciesMock() (*ArticlesMock, *Handler) {
	articlesMock := new(ArticlesMock)
	handler := Handler{
		articles: articlesMock,
		address:  "localhost",
		port:     8080,
		router:   mux.NewRouter(),
	}
	return articlesMock, &handler
}
