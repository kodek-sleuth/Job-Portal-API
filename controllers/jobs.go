package controllers

import (
	"JobPortalBackend/helpers"
	"JobPortalBackend/models"
	"encoding/json"
	"fmt"
	"net/http"
)

//
//// Task return json data.
//// Receive and read json data
//// Finesse error handling.
//// json.Marshal accepts an interface(type/struct) and returns the data(encoded) and error.
//// json.Decode accepts two parameters, the JSON object to decode and where its result will be stored(pointer).
//// create job creation function and fix error handling.
//// create job struct
//// create jobs array.
//
//import (
//	"JobPortalBackend/helpers"
//	"JobPortalBackend/models"
//	"encoding/json"
//	"github.com/gorilla/mux"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"net/http"
//)
//
func CreateJob(res http.ResponseWriter, req *http.Request){
	var job models.Job

	if err := json.NewDecoder(req.Body).Decode(&job); err != nil {
		helpers.ErrorResponse(res, http.StatusInternalServerError, err.Error())
		return
	}

	if req.ContentLength == 0 {
		helpers.ErrorResponse(res, http.StatusBadRequest, "empty json body")
		return
	}

	if message, isError := helpers.ValidateJobInput(job); !isError {
		helpers.ErrorResponse(res, http.StatusBadRequest, message)
		return
	}

	result, err := job.CreateJobCollection()

	fmt.Println(result)

	if err != nil {
		helpers.ErrorResponse(res, http.StatusBadRequest, "failed to create job")
		return
	}

	var payload = make(map[string]interface{})
	payload["message"] = "successfully created job"
	payload["job"] = job

	helpers.JSONResponse(res, http.StatusCreated, payload)
}
//
//func GetJobs(res http.ResponseWriter, req *http.Request){
//	jobs, err := models.GetJobs()
//
//	if err != nil {
//		helpers.ErrorResponse(res, http.StatusBadRequest, "failed to get jobs")
//		return
//	}
//
//	var payload = make(map[string]interface{})
//	payload["message"] = "successfully fetched jobs"
//	payload["jobs"] = jobs
//
//	helpers.JSONResponse(res, http.StatusOK, payload)
//}
//
//func GetJob(res http.ResponseWriter, req *http.Request){
//	var payload = make(map[string]interface{})
//
//	params := mux.Vars(req)
//	id := params["id"]
//
//	result, err := models.GetJob(id)
//	if err != nil {
//		if err == "no job found"{
//			helpers.ErrorResponse(res, http.StatusNotFound, "no job found")
//			return
//		}
//		helpers.ErrorResponse(res, http.StatusInternalServerError, "failed to fetch job")
//		return
//	}
//
//	payload["message"] = "successfully fetched job"
//	payload["job"] = result
//
//	helpers.JSONResponse(res, http.StatusOK, payload)
//}
//
//func UpdateJob(res http.ResponseWriter, req *http.Request){
//	var job models.Job
//
//	err := json.NewDecoder(req.Body).Decode(&job)
//
//	if err != nil {
//		helpers.ErrorResponse(res, http.StatusInternalServerError, "failed to update job")
//		return
//	}
//
//	if req.ContentLength == 0 {
//		helpers.ErrorResponse(res, http.StatusBadRequest, "empty json body")
//		return
//	}
//
//	params := mux.Vars(req)
//	id := params["id"]
//
//	result, errr := job.UpdateJobCollection(id)
//	if errr != nil {
//		if errr == "no job found to update"{
//			helpers.ErrorResponse(res, http.StatusNotFound, "no job found to update")
//			return
//		}
//		helpers.ErrorResponse(res, http.StatusInternalServerError, "failed to update job")
//		return
//	}
//
//	var payload = make(map[string]interface{})
//	payload["message"] = "successfully updated job"
//	payload["job"] = map[string]string{
//		"_id": id,
//		"company": result.Company,
//		"criteria": result.Criteria,
//		"location": result.Location,
//		"description": result.Description,
//		"salary": result.Salary,
//	}
//
//	helpers.JSONResponse(res, http.StatusOK, payload)
//}
//
//func DeleteJob(res http.ResponseWriter, req *http.Request){
//	var payload = make(map[string]interface{})
//
//	params := mux.Vars(req)
//	id := params["id"]
//
//	result, err := models.DeleteJob(id)
//	if err != nil {
//		if err == "no job found"{
//			helpers.ErrorResponse(res, http.StatusNotFound, "no job found to delete")
//			return
//		}
//		helpers.ErrorResponse(res, http.StatusInternalServerError, "failed to delete job")
//		return
//	}
//
//	payload["message"] = "successfully deleted job"
//	payload["job"] = result
//
//	helpers.JSONResponse(res, http.StatusOK, payload)
//}
