package main

import (
	"github.com/RainerGevers/tasker/config"
	"github.com/RainerGevers/tasker/db"
	"github.com/RainerGevers/tasker/lib"
	"github.com/RainerGevers/tasker/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init(){
	err := godotenv.Load()
	if err != nil {
	}
}

func main() {
	log.Println("Starting Up")
	database, err := db.NewConnection()
	if err != nil {
		log.Fatal(err)
		return
	}
	database.AutoMigrate(&models.Migration{})
	db.RunMigrations(database)
	env := config.Env{Database: database}
	r := lib.AddRoutes(mux.NewRouter(), &env)
	log.Fatal(http.ListenAndServe("127.0.0.1:4500", r))
}