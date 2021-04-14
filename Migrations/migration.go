package Migrations

import (
	"github.com/Joggz/FintT-Backend---Go-.git/Helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=FinT-Backend password=password sslmode=disable")
	Helpers.HandleError(err)
	return db
}

func creatAccount() {
	db := connectDB()

	users := [2]User{
		{Username: "Joggz", Email: "joggz@mailinator.com"},
		{Username: "Prince", Email: "prince@mailinator.com"},
	}

	for i := 0; i < len(users); i++ {
		generatePassword := Helpers.HashAndSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatePassword}
		db.Create(&user)

		account := Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate(){
	db := connectDB()

	db.AutoMigrate(&User{}, & Account{})
	defer db.Close()

	creatAccount()
}