package tests

import (
	"JobPortalBackend/controllers"
	"bytes"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/jobs", controllers.CreateJob).Methods("POST")
	router.HandleFunc("/api/jobs", controllers.GetJobs).Methods("GET")
	router.HandleFunc("/api/jobs/{id}", controllers.GetJob).Methods("GET")
	router.HandleFunc("/api/jobs/{id}", controllers.UpdateJob).Methods("PUT")
	router.HandleFunc("/api/jobs/{id}", controllers.DeleteJob).Methods("DELETE")

	return router
}

func TestGetJobs(t *testing.T){
	expected := `{"jobs":[{"id":"3244444","company":"Andela","criteria":"Full-time","location":"Uganda","description":"JEFF","salary":"12,444"}],"message":"successfully fetched jobs"}`
	request, err := http.NewRequest("GET", "/api/jobs", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder() // Record Response
	Router().ServeHTTP(response, request) // Start server
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")
	assert.Equal(t, response.Body.String(), expected)
}

func TestCreateJobs(t *testing.T){
	jsonStr := []byte(`{"company": "Google","criteria": "full-time","salary": "25,0000","description": "hello","location": "Kampala, Uganda"}`)
	request, _ := http.NewRequest("POST", "/api/jobs", bytes.NewBuffer(jsonStr))
	response := httptest.NewRecorder() // Record Response
	Router().ServeHTTP(response, request) // Start server
	assert.Equal(t, http.StatusCreated, response.Code, "CREATED response is expected")
}

func TestGetJob(t *testing.T){
	request, _ := http.NewRequest("GET", "/api/jobs/3244444", nil)
	response := httptest.NewRecorder() // Record Response
	Router().ServeHTTP(response, request) // Start server
	assert.Equal(t, http.StatusOK, response.Code, "Ol response is expected")
}

func TestUpdateJob(t *testing.T){
	jsonStr := []byte(`{"company": "Googe","criteria": "full-time","salary": "25,0000","description": "hello","location": "Kampala, Uganda"}`)
	request, _ := http.NewRequest("PUT", "/api/jobs/3244444", bytes.NewBuffer(jsonStr))
	response := httptest.NewRecorder() // Record Response
	Router().ServeHTTP(response, request) // Start server
	assert.Equal(t, http.StatusOK, response.Code, "Ok response is expected")
}

func TestDeleteJob(t *testing.T){
	request, _ := http.NewRequest("DELETE", "/api/jobs/3244444", nil)
	response := httptest.NewRecorder() // Record Response
	Router().ServeHTTP(response, request) // Start server
	assert.Equal(t, http.StatusNoContent, response.Code)
}



