package controllers

import (
	"JobPortalBackend/helpers"
	"JobPortalBackend/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUsers(res http.ResponseWriter, req *http.Request){
	// append error
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

	//if _, boolean := user.GetUser(); boolean == false {
	//	helpers.ErrorResponse(res, http.StatusConflict, "email used is already registered on this platform" )
	//	return
	//}

	if _, err := user.CreateUsers(); err != nil {
		helpers.ErrorResponse(res, http.StatusBadRequest, fmt.Sprintf("failed to sign up, please try again later %+v", err.Error()))
		return
	}



	var payload = make(map[string]interface{})
	payload["message"] = "successfully signed up"
	payload["user"] = map[string]string{
		"name": user.Name,
		"email": user.Email,
	}

	helpers.JSONResponse(res, http.StatusCreated, payload)
}

func LoginUsers(res http.ResponseWriter, req *http.Request){
	var user models.Users

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		helpers.ErrorResponse(res, http.StatusInternalServerError, err.Error())
		return
	}

	if req.ContentLength == 0 {
		helpers.ErrorResponse(res, http.StatusBadRequest, "empty json body")
		return
	}

	if message, isError := helpers.ValidateLoginInput(user); !isError {
		helpers.ErrorResponse(res, http.StatusBadRequest, message)
		return
	}

	if _, err := user.CheckLoginCredentials(); err != nil {
		helpers.ErrorResponse(res, http.StatusBadRequest, fmt.Sprintf("wrong username or password %+v", err.Error()))
		return
	}

	token, err := helpers.GenerateJWT()
	if err != nil{
		helpers.ErrorResponse(res, http.StatusBadRequest, fmt.Sprintf("wrong username or password %+v", err.Error()))
		return
	}

	var payload = make(map[string]interface{})
	payload["message"] = "successfully logged in"
	payload["user"] = map[string]string{
		"name": user.Name,
		"email": user.Email,
		"token": token,
	}

	helpers.JSONResponse(res, http.StatusCreated, payload)
}




