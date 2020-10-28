package main

import (
	articleservice "article-service"
	"article-service/config"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Println("Starting article service")

	config, err := config.GetServiceConfig()
	if err != nil {
		log.Fatalf("Error getting config : %v Exiting..", err.Error())
	}

	handler, err := articleservice.New(config)
	if err != nil {
		log.Fatal("Unable to Instantiate article service instance ", err.Error())
	}

	handler.Start()

}
