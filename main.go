package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	Error     string
	ErrorCode int
}

type Questions []struct {
	ID          int    `json:"id"`
	Question    string `json:"question"`
	Placeholder string `json:"placeholder"`
	Answers     []struct {
		Text string `json:"text"`
		ID   string `json:"id"`
	} `json:"answers"`
	Response string `json:"response"`
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// TODO: serve frontend

	http.HandleFunc("/", frontend)
	// TODO: handle answers
	http.HandleFunc("/api/answer", answer)
	http.ListenAndServe(":3000", nil)
}

func answer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		w.Write([]byte(""))
		return
	}
	content, _ := ioutil.ReadAll(r.Body)
	_ = os.WriteFile("./answers.json", content, 0644)
	fmt.Println("stored answers in answers.json")

	response := Response{}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func frontend(w http.ResponseWriter, r *http.Request) {
	response := Response{}

	js, err := json.Marshal(response)
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
