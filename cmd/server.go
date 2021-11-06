package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"httpArithmProg/internal"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/putTask", internal.PutTask).Methods("POST")
	router.HandleFunc("/api/getList", internal.GetListAndStatus).Methods("GET")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
