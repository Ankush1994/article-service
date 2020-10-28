// Package articleservice provides API for article service
package articleservice

import (
	"article-service/article"
	"article-service/config"
	"article-service/dbclient"
	"article-service/handler"
	"log"
)

// New creates a instance of the article service and intializes all the
// dependencies like Handler, DB etc
func New(config *config.ServiceConfig) (*handler.Handler, error) {
	dbClient := dbclient.New(config.DBHost, config.DBPort, config.DBName, config.DBUser, config.DBPassword)
	articleInstance := article.New(dbClient)
	handler, err := handler.New(config.Address, config.Port, articleInstance)
	if err != nil {
		log.Fatalf("Error initializing handler : %v Exiting... ", err)
	}
	return handler, nil
}
