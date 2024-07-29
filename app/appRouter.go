package app

import (
	mux2 "github.com/gorilla/mux"
)

func Router() *mux2.Router {
	mux := mux2.NewRouter()

	mux.HandleFunc("/api/users", GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", GetUser).Methods("GET")
	mux.HandleFunc("/api/user/", CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", DeleteUser).Methods("DELETE")

	return mux
}
