package Migrations

import (
	"github.com/Joggz/FintT-Backend---Go-.git/Helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/Joggz/FintT-Backend---Go-.git/Interface"
	"github.com/Joggz/FintT-Backend---Go-.git/Helpers"
)

type User struct {
	gorm.Model
	Username string
	Email string
	Password string
}

type Account struct {
	gorm.Model
	Type string
	Name string
	Balance uint
	UserID uint
}


func creatAccount() {
	db := Helpers.ConnectDB()

	users := [2]Interface.User{
		{Username: "Joggz", Email: "joggz@mailinator.com"},
		{Username: "Prince", Email: "prince@mailinator.com"},
	}

	for i := 0; i < len(users); i++ {
		generatePassword := Helpers.HashAndSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatePassword}
		db.Create(&user)

		account := Interface.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate(){
	db := Helpers.ConnectDB()
	User := Interface.User{}
	Account := Interface.Account{}

	db.AutoMigrate(&User, & Account)
	defer db.Close()

	creatAccount()
}