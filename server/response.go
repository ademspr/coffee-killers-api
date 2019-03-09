package server

import (
	"encoding/json"
	"net/http"
)

// Error http error response
func Error(w http.ResponseWriter, code int, message string) {
	JSON(w, code, map[string]string{"error": message})
}

// JSON http ok response with json
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// JSONWithCookie http ok response with json and cookies
func JSONWithCookie(w http.ResponseWriter, code int, payload interface{}, cookie http.Cookie) {
	response, _ := json.Marshal(payload)
	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
