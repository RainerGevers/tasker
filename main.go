package main

import (
	"github.com/RainerGevers/tasker/config"
	"github.com/RainerGevers/tasker/db"
	"github.com/RainerGevers/tasker/lib"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
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
	db.RunMigrations(database)
	logger := config.InitLogger()
	env := config.Env{Database: database, Logger: logger}
	r := lib.AddRoutes(mux.NewRouter(), &env)
	logger.Println("Running on port 4500")
	log.Fatal(http.ListenAndServe("0.0.0.0:4500", r))
}
