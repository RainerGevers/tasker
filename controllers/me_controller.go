package controllers

import (
	"encoding/json"
	"github.com/RainerGevers/tasker/config"
	"github.com/RainerGevers/tasker/models"
	"github.com/RainerGevers/tasker/serializers"
	"net/http"
)

func MeGet(_ *config.Env, w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value("current_user").(models.User)
	userSerialized := serializers.UserSerializer{
		Email:    currentUser.Email,
		Username: currentUser.Username,
		Uuid:     currentUser.Uuid,
	}
	response, _ := json.Marshal(userSerialized)
	w.WriteHeader(200)
	_, _ = w.Write(response)
}
