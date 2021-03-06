package controllers

import (
	"encoding/json"
	"github.com/RainerGevers/tasker/config"
	"github.com/RainerGevers/tasker/lib/argon"
	"github.com/RainerGevers/tasker/models"
	"github.com/RainerGevers/tasker/serializers"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type user struct {
	User userParams `json:"user"`
}

type userParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type loginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type apiError struct {
	Reference string `json:"reference"`
	Message   string `json:"message"`
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
		err := apiError{Reference: "user_already_registered", Message: "This user email has already registered"}
		response, _ := json.Marshal(err)
		w.WriteHeader(409)
		_, _ = w.Write(response)
		return
	}

	passwordHash, err := argon.GeneratePassword(userData.User.Password)
	if err != nil {
		err := apiError{Reference: "password_hash_failure", Message: "There was a problem while trying to hash the password."}
		response, _ := json.Marshal(err)
		w.WriteHeader(422)
		_, _ = w.Write(response)
		return
	}

	user = models.User{Password: passwordHash, Username: userData.User.Username, Email: userData.User.Email, Uuid: uuid.NewV4().String()}

	dbResult = env.Database.Create(&user)

	if dbResult.Error != nil {
		err := apiError{Reference: "user_insert_error", Message: "There was a problem inserting the user into the database"}
		response, _ := json.Marshal(err)
		w.WriteHeader(422)
		_, _ = w.Write(response)
		return
	}

	userSerialized := serializers.UserSerializer{Username: user.Username, Email: user.Email, Uuid: user.Uuid}

	response, _ := json.Marshal(userSerialized)

	w.WriteHeader(201)
	_, _ = w.Write(response)
}

func UserLogin(env *config.Env, w http.ResponseWriter, r *http.Request) {
	var loginData loginParams

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &loginData)
	if err != nil {
		log.Fatal(err)
		return
	}

	if loginData.Email == "" || loginData.Password == "" {
		err := apiError{Reference: "email_or_password_blank", Message: "The email or password appears to be blank."}
		response, _ := json.Marshal(err)
		w.WriteHeader(422)
		_, _ = w.Write(response)
		return
	}

	user := models.User{}
	dbResult := env.Database.Where("email = ?", loginData.Email).First(&user)
	if dbResult.Error != nil {
		err := apiError{Reference: "email_or_password_incorrect", Message: "The email or password appears to be incorrect."}
		response, _ := json.Marshal(err)
		w.WriteHeader(422)
		_, _ = w.Write(response)
		return
	}

	passwordMatch, err := argon.ComparePassword(loginData.Password, user.Password)
	if err != nil || passwordMatch == false {
		err := apiError{Reference: "email_or_password_incorrect", Message: "The email or password appears to be incorrect."}
		response, _ := json.Marshal(err)
		w.WriteHeader(422)
		_, _ = w.Write(response)
		return
	}

	session := models.Session{User: user, Uuid: uuid.NewV4().String(), ExpiresAt: time.Now().Add(time.Hour * 24), RefreshCode: uuid.NewV4().String()}
	env.Database.Create(&session)

	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = session.Uuid
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		env.Logger.Println(err)
		err := apiError{Reference: "token_error", Message: "There was a error while trying to generate token."}
		response, _ := json.Marshal(err)
		w.WriteHeader(422)
		_, _ = w.Write(response)
		return
	}

	refreshToken := jwt.New(jwt.SigningMethodHS512)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["uuid"] = session.RefreshCode
	refreshClaims["exp"] = time.Now().Add(time.Hour * 24 * 14)

	tr, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		env.Logger.Println(err)
		err := apiError{Reference: "token_error", Message: "There was a error while trying to generate token."}
		response, _ := json.Marshal(err)
		w.WriteHeader(422)
		_, _ = w.Write(response)
		return
	}

	loginSerialized := serializers.LoginSerializer{Uuid: user.Uuid, Email: user.Email, Username: user.Username, AuthToken: t, RefreshToken: tr}

	response, _ := json.Marshal(loginSerialized)

	w.WriteHeader(200)
	_, _ = w.Write(response)

}
