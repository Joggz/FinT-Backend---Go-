package Helpers

import (
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