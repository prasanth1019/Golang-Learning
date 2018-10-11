package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "flag"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}

func YourHandler3(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
    router := mux.Vars(r)
    fmt.Println(router["username"])
}

func YourHandler2(w http.ResponseWriter, r *http.Request)  {
  reqVariables := mux.Vars(r)
  w.WriteHeader(500)
  key_details := reqVariables["key"]
  fmt.Println(key_details)
}

func main() {
    var dirStr string
    flag.StringVar( &dirStr, "dirStr", ".", "the directory to serve files from. Defaults to the current dir")
    flag.Parse()
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", YourHandler)
    r.HandleFunc("/fruit/{key}", YourHandler2)
    // r.HandleFunc("/auth", YourHandler3).Queries("username")

    //Serving Htto files
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dirStr))))
    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}
