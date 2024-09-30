package routes

import (
    "github.com/gorilla/mux"
)


func SetupRoutes(r *mux.Router) {
    userRouter := r.PathPrefix("/users").Subrouter()   
    UserRoutes(userRouter)

}