package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedro0402/go-mod/config"
	"github.com/pedro0402/go-mod/handlers"
	"github.com/pedro0402/go-mod/models"
)

func main() {

	dbConnection := config.SetupDB()

	_, err := dbConnection.Exec(models.CreateTableSQL)

	if err != nil {
		log.Fatal(err)
	}

	defer dbConnection.Close()

	router := mux.NewRouter()

	taskHandler := handlers.NewTaskHandler(dbConnection)

	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
