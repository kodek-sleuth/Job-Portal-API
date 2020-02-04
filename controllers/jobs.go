package controllers

// Task return json data.
// Receive and read json data
// Finesse error handling
// json.Marshal accepts an interface(type/struct) and returns the data(encoded) and error
// json.Decode accepts two parameters, the JSON object to decode and where its result will be stored(pointer)
// create job creation function and fix error handling
// create job struct
// create jobs array

import (
	"JobPortalBackend/helpers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

type Job struct {
	Id string `json:"id"`
	Company string `json:"company"`
	Criteria string `json:"criteria"`
	Location string `json:"location"`
	Description string `json:"description"`
	Salary string `json:"salary"`
}

var jobs []Job

func CreateJob(res http.ResponseWriter, req *http.Request){
	var job Job

	err := json.NewDecoder(req.Body).Decode(&job)

	if req.ContentLength == 0 {
		helpers.ErrorResponse(res, http.StatusBadRequest, "empty json body")
		return
	}

	// Error handling
	if err != nil {
		fmt.Println(err)
		helpers.ErrorResponse(res, 500, "failed to create job")
		return
	}

	helpers.ValidateUserInput(res, job)

	job.Id = strconv.Itoa(rand.Intn(12000000))

	jobs = append(jobs, job)

	var payload = make(map[string]interface{})
	payload["message"] = "successfully created job"
	payload["job"] = job

	helpers.JSONResponse(res, 201, payload)
	return
}

func GetJobs(res http.ResponseWriter, req *http.Request){
	var payload = make(map[string]interface{})
	jobs = append(jobs, Job{ Id: "3244444", Company: "Andela", Criteria: "Full-time", Location: "Uganda",
		Description: "JEFF", Salary: "12,444" })

	if len(jobs) < 1{
		helpers.ErrorResponse(res, 404, "no jobs found")
		return
	}
	payload["message"] = "successfully fetched jobs"
	payload["jobs"] = jobs

	helpers.JSONResponse(res, 200, payload)
}

func GetJob(res http.ResponseWriter, req *http.Request){
	var payload = make(map[string]interface{})

	params := mux.Vars(req)
	id := params["id"]
	fmt.Println(jobs[:1])
	for _, job := range jobs{
		if job.Id == id {
			payload["message"] = "successfully fetched job"
			payload["job"] = job
			helpers.JSONResponse(res, 200, payload)
			return
		}
	}

	helpers.ErrorResponse(res, 404, "no job found")
}

func UpdateJob(res http.ResponseWriter, req *http.Request){
	var payload = make(map[string]interface{})
	var jobAdd Job

	err := json.NewDecoder(req.Body).Decode(&jobAdd)

	if req.ContentLength == 0 {
		helpers.ErrorResponse(res, http.StatusBadRequest, "empty json body")
		return
	}

	// Error handling
	if err != nil {
		fmt.Println(err)
		helpers.ErrorResponse(res, 500, "failed to update job")
		return
	}

	params := mux.Vars(req)
	id := params["id"]



	for index, job := range jobs{
		if job.Id == id {
			jobAdd.Id = jobs[index].Id
			jobs[index].Company = jobAdd.Company
			jobs[index].Criteria = jobAdd.Criteria
			jobs[index].Description = jobAdd.Description
			jobs[index].Location = jobAdd.Location
			jobs[index].Salary = jobAdd.Salary

			payload["message"] = "successfully updated job"
			payload["job"] = jobAdd

			helpers.JSONResponse(res, 200, payload)
			return
		}
	}

	helpers.ErrorResponse(res, 404, "no job found")
}

func DeleteJob(res http.ResponseWriter, req *http.Request){
	var payload = make(map[string]interface{})

	params := mux.Vars(req)
	id := params["id"]

	for index, job := range jobs{
		if job.Id == id {
			jobs = append(jobs[:index], jobs[index+1:]...)
			payload["message"] = "successfully deleted job"

			helpers.JSONResponse(res, 204, payload)
			return
		}
	}

	helpers.ErrorResponse(res, 404, "no job found")
}

