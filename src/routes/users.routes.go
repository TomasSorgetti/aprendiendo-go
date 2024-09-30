package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
    r.HandleFunc("", GetUsersHandler).Methods("GET")
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Lista de usuarios"))
}
