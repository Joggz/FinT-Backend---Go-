package User

import (
	"time"

	"github.com/Joggz/FintT-Backend---Go-.git/Helpers"
	"github.com/Joggz/FintT-Backend---Go-.git/Interface"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)


func Login( username string, pass string) map[string] interface{}{
	user := Interface.User{}
	db := Helpers.ConnectDB()

	if db.Where("username = ? ", username).First(&user).RecordNotFound(){
		return map[string]interface{}{"message": "Record not found"}
	}

	pErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if pErr == bcrypt.ErrMismatchedHashAndPassword && pErr != nil {
		return map[string]interface{}{"message": "Wrong Password"}
	}
	
	accounts := []Interface.ResponseAccount{}
	db.Table("accounts").Select("id, name, balance").Where("user_id = ? ", user.ID).Scan(&accounts)

	//var accounts  []Interface.ResponseAccount
	//db.Table("accounts").Select("id, balance, name").Where("users_id = ?", user.ID).Scan(&account)

	responseUser := &Interface.ResponseUser{
		ID: user.ID,
		UserName: user.Username,
		Email: user.Email,
		Account: accounts,
	}

	defer db.Close()

	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry": time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	Helpers.HandleError(err)

	response := map[string]interface{}{"message": "succesfu	l"}
	response["jwt"] = token
	response["user"] = responseUser

	return response
}