package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"httpArithmProg/internal"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {

	var flagMaxNum uint
	flag.UintVar(&flagMaxNum, "N", 3, "max amount of processes to work at the same time")
	flag.Parse()

	wg := sync.WaitGroup{}
	for i := uint(0); i < flagMaxNum; i++ {
		wg.Add(1)
		go internal.TaskCompletion(&wg, i)
	}

	go func() {
		wg.Wait()
		log.Println("exiting successfully")
		os.Exit(0)
	}()

	router := mux.NewRouter()

	router.HandleFunc("/api/putTask", internal.PutTask).Methods("POST")
	router.HandleFunc("/api/getList", internal.GetListAndStatus).Methods("GET")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
