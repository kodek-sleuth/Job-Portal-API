package main

import (
	"JobPortalBackend/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/api/jobs", controllers.CreateJob).Methods("POST")
	router.HandleFunc("/api/jobs", controllers.GetJobs).Methods("GET")
	router.HandleFunc("/api/jobs/{id}", controllers.GetJob).Methods("GET")
	router.HandleFunc("/api/jobs/{id}", controllers.UpdateJob).Methods("PUT")
	router.HandleFunc("/api/jobs/{id}", controllers.DeleteJob).Methods("DELETE")
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
