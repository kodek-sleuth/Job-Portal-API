package main

import (
	"JobPortalBackend/controllers"
	"github.com/gorilla/mux"
	"JobPortalBackend/models"
	"log"
	"net/http"
	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "kevina52"
	dbname   = "authors"
)

func main(){
	router := mux.NewRouter()
	models.SQLConnection()

	router.HandleFunc("/api/jobs", controllers.CreateJob).Methods("POST")
	//router.HandleFunc("/api/jobs", controllers.GetJobs).Methods("GET")
	//router.HandleFunc("/api/jobs/{id}", controllers.GetJob).Methods("GET")
	//router.HandleFunc("/api/jobs/{id}", controllers.UpdateJob).Methods("PUT")
	//router.HandleFunc("/api/jobs/{id}", controllers.DeleteJob).Methods("DELETE")

	router.HandleFunc("/api/jobs", controllers.CreateUsers).Methods("POST")
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
