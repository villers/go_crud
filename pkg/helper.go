package pkg

import (
	"encoding/json"
	"net/http"
)

type JsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func RespondWithNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, input interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(input); err != nil {
		panic(err)
	}
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, JsonErr{Code: code, Text: message})
}
