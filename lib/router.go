package lib

import (
	"github.com/RainerGevers/tasker/config"
	"github.com/RainerGevers/tasker/controllers"
	"github.com/gorilla/mux"
)

func AddRoutes(r *mux.Router, env *config.Env) *mux.Router {

	/*#################  VERSION 1  #################*/

	v1 := r.PathPrefix("/v1").Subrouter()

	maintenance := v1.PathPrefix("/maintenance").Subrouter()
	maintenance.HandleFunc("/alive", config.AppHandler{Env: env, Handler: controllers.AliveCheck}.ServeHTTP).Methods("GET")

	users := v1.PathPrefix("/users").Subrouter()
	users.HandleFunc("/register", config.AppHandler{Env: env, Handler: controllers.UsersRegister}.ServeHTTP).Methods("POST")
	users.HandleFunc("/login", config.AppHandler{Env: env, Handler: controllers.UserLogin}.ServeHTTP).Methods("POST")

	/*######################################################################################################################################*/

	return r
}
