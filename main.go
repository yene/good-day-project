package main

import (
	"encoding/json"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// TODO: serve frontend

	http.HandleFunc("/", frontend)
	// TODO: handle answers
	http.HandleFunc("/api/answer", frontend)
	http.ListenAndServe(":3000", nil)
}

func frontend(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// HTTP errors
//http.Error(w, http.StatusText(500), 500)
//http.Error(w, err.Error(), http.StatusInternalServerError)
