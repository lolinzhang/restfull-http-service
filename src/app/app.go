package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"handler"
)

func main(){
	r := mux.NewRouter()
	// r.HeadersRegexp("Content-Type", "application/(text|json)")

	r.HandleFunc("/users", handler.UsersGetHandler).Methods("GET")
	r.HandleFunc("/users", handler.UsersPOSTHandler).Methods("POST")
	r.HandleFunc("/users/{user_id:[0-9]+}/relationships", handler.UsersRelationshipsGetHandler).Methods("GET")
	r.HandleFunc("/users/{user_id:[0-9]+}/relationships/{other_user_id:[0-9]+}", handler.UsersRelationshipsPutHandler).Methods("PUT")

	fmt.Println("http listen on 8000.");

	http.ListenAndServe(":8000", r)
}
