package dbclient

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	// _ import needed for GORM
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// New - creates a new instance of Db client
func New(host string, port int, dbName string, user string, password string) *gorm.DB {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbName, password)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln("Error while connecting to db:", err)
		return nil
	}
	db.DB().SetMaxOpenConns(25)
	db.DB().SetMaxIdleConns(25)
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	log.Println("Connected to DB!!")
	return db
}
