package controllers

import (
	"github.com/RainerGevers/tasker/config"
	"github.com/RainerGevers/tasker/models"
	"net/http"
)

func MeGet(env *config.Env, w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value("current_user").(models.User)
	env.Logger.Printf("context: %#v\n %T\n", currentUser, currentUser)
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}
