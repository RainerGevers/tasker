package controllers

import (
	"github.com/RainerGevers/tasker/config"
	"net/http"
)

func MeGet(env *config.Env, w http.ResponseWriter, r *http.Request) {

	env.Logger.Printf("context: %#v", r.Context().Value("current_user_uuid"))
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}
