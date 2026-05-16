package main

import (
	"encoding/json"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello API",
	})
}

func main() {
	http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)
}
