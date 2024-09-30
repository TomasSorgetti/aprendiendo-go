package main

import (
	"GolangTest/src/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)



func main() {
	router := mux.NewRouter()

	routes.SetupRoutes(router)

	log.Println("Servidor corriendo en http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}