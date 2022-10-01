package lib

import (
	"encoding/json"
	"net/http"
)

type ErrorBody struct {
	Message string `json:"message"`
}

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func SendError(w http.ResponseWriter, code int, message string) {
	var body ErrorBody = ErrorBody{
		Message: message,
	}
	bytes, err := json.Marshal(body)
	CheckError(err)
	w.WriteHeader(code)
	w.Write(bytes)
}
