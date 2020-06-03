package utility

import (
	"fmt"
	token "github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
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

func ValidationJWT(req *http.Request) error {
	result := ParseTokenFromRequest(req)
	_, err := token.Parse(result, func(tokenData *token.Token) (interface{}, error) {
		if _, ok := tokenData.Method.(*token.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(ApiSecret), nil
	})
	if err != nil {
		return err
	}
	return err
}

func ParseTokenFromRequest(req *http.Request) string {
	token := req.URL.Query().Get("jwt_token")
	if token != "" {
		return token
	}
	//control header inside of bearer token e.q (bearer aaaaaa like that)
	bearerToken := req.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
