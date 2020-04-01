package controllers

import (
	"JobPortalBackend/helpers"
	"JobPortalBackend/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUsers(res http.ResponseWriter, req *http.Request){
	var user models.Users

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		helpers.ErrorResponse(res, http.StatusInternalServerError, err.Error())
		return
	}

	if req.ContentLength == 0 {
		helpers.ErrorResponse(res, http.StatusBadRequest, "empty json body")
		return
	}

	if message, isError := helpers.ValidateUserInput(user); !isError {
		helpers.ErrorResponse(res, http.StatusBadRequest, message)
		return
	}

	result, err := user.CreateUsersCollection()

	fmt.Println(result)

	if err != nil {
		helpers.ErrorResponse(res, http.StatusBadRequest, "failed to sign up, please try again later ")
		return
	}

	var payload = make(map[string]interface{})
	payload["message"] = "successfully signed up"
	payload["user"] = user

	helpers.JSONResponse(res, http.StatusCreated, payload)
}
