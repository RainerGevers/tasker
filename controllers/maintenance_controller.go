package controllers

import (
	"github.com/RainerGevers/tasker/config"
	"net/http"
)

func AliveCheck(env *config.Env, w http.ResponseWriter, r *http.Request) {
	env.Logger.Println("V1::AliveCheck")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}
