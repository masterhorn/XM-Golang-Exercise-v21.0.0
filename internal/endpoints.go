package internal

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/masterhorn/XM-Golang-Exercise-v21.0.0/api/rest"
)

func Handlers() http.Handler {

	r := mux.NewRouter()

	authR := r.NewRoute().Subrouter()
	authR.Use(AccessHeaderMiddleware)

	// Companies
	r.HandleFunc("/company", rest.ListCompanies).Methods("GET")
	r.HandleFunc("/company/{id}", rest.GetCompany).Methods("GET")
	authR.HandleFunc("/company", rest.CreateCompany).Methods("POST")
	authR.HandleFunc("/company/{id}", rest.DeleteCompany).Methods("DELETE")
	r.HandleFunc("/company/{id}", rest.UpdateCompany).Methods("PUT")

	return r
}
