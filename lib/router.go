package lib

import (
	"github.com/RainerGevers/tasker/config"
	"github.com/RainerGevers/tasker/controllers"
	"github.com/gorilla/mux"
)

func AddRoutes(r *mux.Router, env *config.Env) *mux.Router {

	users := r.PathPrefix("/users").Subrouter()
	users.HandleFunc("/register", config.AppHandler{Env: env, Handler: controllers.UsersRegister}.ServeHTTP).Methods("POST")

	return r
}


