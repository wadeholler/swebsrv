package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pborman/uuid"
	"html"
	"log"
	"net/http"
)

func base(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}
func douuid(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, uuid.New())
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", base).Methods("GET")
	myRouter.HandleFunc("/hi", hi).Methods("GET")
	myRouter.HandleFunc("/douuid", douuid).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
