package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Response struct {
	Error     string
	ErrorCode int
}

//go:embed frontend/dist
var staticFiles embed.FS

var exPath string

func main() {
	var staticFS = fs.FS(staticFiles)
	htmlContent, err := fs.Sub(staticFS, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.FS(htmlContent))
	http.Handle("/", fs)

	http.HandleFunc("/api/answer", answer)
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath = filepath.Dir(ex)
	err = os.Mkdir(exPath+"/answers", 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	fmt.Println("Please open http://localhost:3000")
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

	d := time.Now().Format("2006-01-02")
	p := exPath + "/answers/" + d + ".json"
	_ = os.WriteFile(p, content, 0644)
	fmt.Println("stored answers in ", p)

	response := Response{}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
