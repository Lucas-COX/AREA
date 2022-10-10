package lib

import (
	"encoding/json"
	"log"
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
	http.Error(w, string(bytes), code)
}

func LogError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
