package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Lewvy/serv/parser"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	log.Println("Reached HandleRoot")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hello")
	return
}

func getFile(w http.ResponseWriter, r *http.Request) {
	filePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file := filePath + "/md/" + r.PathValue("file")
	str := parser.Lexer(file)
	fmt.Println(file)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, str)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("GET /md/{file}", getFile)
	fmt.Println("Server listening to :8080")
	http.ListenAndServe(":8080", mux)
}
