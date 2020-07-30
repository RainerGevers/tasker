package main

import (
	"github.com/RainerGevers/tasker/db"
	"github.com/RainerGevers/tasker/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Up")
	database, err := db.NewConnection()
	if err != nil {
		log.Fatal(err)
		return
	}
	database.AutoMigrate(&models.Migration{})
	db.RunMigrations(database)
	r := mux.NewRouter()
	r.HandleFunc("/", TestRoute)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func TestRoute(w http.ResponseWriter, _r *http.Request) {
	w.Write([]byte("Hello WORLD!!!\n"))
}