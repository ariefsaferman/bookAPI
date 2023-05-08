package db

import (
	"bookAPI/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "book-api"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
		return nil, err
	}

	db.Debug().AutoMigrate(entity.Book{})

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
