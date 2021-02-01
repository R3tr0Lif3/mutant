package handlers

import (
	"log"

	"mutant/api"

	"net/http"

	"os"

	"github.com/gorilla/mux"

	"github.com/rs/cors"
)

//Handlers main handler
func Handlers() {

	router := mux.NewRouter()
	router.HandleFunc("/mutant", api.IsMutant).Methods("POST")
	router.HandleFunc("/stats", api.GetStats).Methods("GET")
	PORT := os.Getenv("PORT")
	if PORT == "" {

		PORT = "8080"

	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal((http.ListenAndServe(":"+PORT, handler)))

}
