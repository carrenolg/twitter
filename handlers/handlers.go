package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/carrenolg/twitter/middlew"
	"github.com/carrenolg/twitter/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handler() {
	router := mux.NewRouter()

	// middle
	router.HandleFunc("/registry", middlew.Checkdb(routes.Registry)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}