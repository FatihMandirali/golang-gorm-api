package middleware

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}

// set error message in Error struct
func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			var err Error
			err = SetError(err, "Token not found")
			data, _ := json.Marshal(err)
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		}
	}
}
