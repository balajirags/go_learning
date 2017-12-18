package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/", myHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func myHandler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"URL Path is %q", r.URL.Path)
}