package utility

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseData(writer http.ResponseWriter, data interface{}, statusCode int) {
	writer.WriteHeader(statusCode)
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		fmt.Fprintf(writer, "%s", err.Error())
	}
}
