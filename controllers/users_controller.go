package controllers

import (
	"encoding/json"
	"github.com/RainerGevers/tasker/config"
	"github.com/RainerGevers/tasker/lib/argon"
	"github.com/RainerGevers/tasker/models"
	"github.com/RainerGevers/tasker/serializers"
	"io/ioutil"
	"log"
	"net/http"
)

type user struct {
	User userParams `json:"user"`
}

type userParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type error struct {
	Reference string `json:"reference"`
	Message string `json:"message"`
}

func UsersRegister(env *config.Env, w http.ResponseWriter, r *http.Request) {
	var userData user

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &userData)
	if err != nil {
		log.Fatal(err)
		return
	}

	user := models.User{}

	dbResult := env.Database.Where("email = ?", userData.User.Email).First(&user)

	if dbResult.Error == nil {
		err := error{Reference: "user_already_registered", Message: "This user email has already registered"}
		response, _ := json.Marshal(err)
		w.WriteHeader(409)
		w.Write(response)
		return
	}

	passwordHash, err := argon.GeneratePassword(userData.User.Password)
	if err != nil {
		err := error{Reference: "password_hash_failure", Message: "There was a problem while trying to hash the password."}
		response, _ := json.Marshal(err)
		w.WriteHeader(422)
		w.Write(response)
		return
	}

	user = models.User{Password: passwordHash, Username: userData.User.Username, Email: userData.User.Email}

	dbResult = env.Database.Create(&user)

	if dbResult.Error != nil {
		err := error{Reference: "user_insert_error", Message: "There was a problem inserting the user into the database"}
		response, _ := json.Marshal(err)
		w.WriteHeader(422)
		w.Write(response)
		return
	}

	userSerialized := serializers.UserSerializer{Id: user.ID, Username: user.Username, Email: user.Email}

	response, _ := json.Marshal(userSerialized)

	w.WriteHeader(201)
	w.Write(response)
}


