package Helpers

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func HandleError(err error) {
	if err != nil {
		panic(err.Error())
		log.Fatal("Error:", err.Error)
	}
}

func HashAndSalt (password []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	HandleError(err)
	return  string(hashed)
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=FinT-Backend password=password sslmode=disable")
	HandleError(err)
	return db
}