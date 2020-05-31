package utility

import (
	token "github.com/dgrijalva/jwt-go"
	"time"
)

const ApiSecret = "MY_SCRT"

func ObtainTokenWithUser(user int) (string, error) {
	data := setTokenData(user)
	return data.SignedString(ApiSecret)
}

func setTokenData(user int) *token.Token {
	claim := token.MapClaims{}
	claim["authorized"] = true
	claim["user_id"] = user
	claim["exp"] = time.Now().Add(time.Hour * 1).Unix()
	return token.NewWithClaims(token.SigningMethodES256, claim)
}
