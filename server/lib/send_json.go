package lib

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, body any) {
	bytes, err := json.Marshal(body)
	CheckError(err)
	w.Write(bytes)
}
