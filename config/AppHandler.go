package config

import (
	"gorm.io/gorm"
	"net/http"
)

type Env struct {
	Database *gorm.DB
}

type AppHandler struct {
	Env *Env
	Handler func(env *Env, w http.ResponseWriter, r *http.Request)
}

func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ah.Handler(ah.Env, w, r)
}