package utils

import (
	"CAKE-STORE/entity"
	"encoding/json"
	"net/http"
)

func HandleErrorResponse(w http.ResponseWriter, code int, status string, msg string) {
	err := entity.Error{
		Code:    code,
		Status:  status,
		Message: msg,
	}
	errorByte, error := json.Marshal(err)
	if error != nil {
		panic(error)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(errorByte)
}
