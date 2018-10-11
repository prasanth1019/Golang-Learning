package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("This is main \n ")
	http.HandleFunc("/", handlerfun)
	http.ListenAndServe(":3243", nil)
}

func handlerfun(w http.ResponseWriter, req *http.Request) {
	log.Printf("This is handlerfun \n ")
	// w.Header().set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server. \n"))
}
