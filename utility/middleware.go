package utility

import (
	"encoding/json"
	"net/http"
)

func AddHeaderToJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		next(writer, request)
	}
}

func ValidateToken(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := ValidationJWT(request)
		if err != nil {
			writer.WriteHeader(400)
			json.NewEncoder(writer).Encode("Token isnt valid")
			return
		}
		next(writer, request)
	}
}
