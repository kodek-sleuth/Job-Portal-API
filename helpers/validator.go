package helpers

import (
	_ "fmt"
	"github.com/fatih/structs"
	//"golang.org/x/crypto/bcrypt"
)

func ValidateJobInput(payload interface{})(string, bool){
	// Convert struct to map
	newJob := structs.Map(payload)
	company, criteria, salary, description, location := newJob["Company"].(string), newJob["Criteria"].(string),
	newJob["Salary"].(string), newJob["Description"].(string), newJob["Location"].(string)

	if len(company) < 1 || len(salary) < 1 || len(description) < 1 || len(location) < 1 || len(criteria)  < 1{
		return "Company, criteria, location, description and salary should not be left empty", false
	}

	return "", true
}

func ValidateUserInput(payload interface{})(string, bool) {
	// Convert struct to map
	newUser := structs.Map(payload)
	name, email, password := newUser["Name"].(string), newUser["Email"].(string),
		newUser["Password"].(string)
	if len(name) < 1 || len(email) < 1 || len(password) < 1 {
		return "Name, email and password are required", false
	}

	return "", true
}

func ValidateLoginInput(payload interface{})(string, bool){
	newUser := structs.Map(payload)
	email, password := newUser["Email"].(string), newUser["Password"].(string)
	if len(email) < 1 || len(password) < 1 {
		return "Email and password are required", false
	}
	return "", true
}

//func CheckLoginCredentials()(*Users, error) {
//	db, err := SQLConnection()
//	if err != nil {
//		return u, err
//	}
//
//	// ComparePasswords
//	password := []byte(u.Password)
//	err := bcrypt.CompareHashAndPassword()
//	if err != nil {
//		return u, err
//	}
//
//	if err := db.Where("email = ? AND password = ?", u.Email, u.Password).First(u).Error; err != nil {
//		return u, err
//	}
//
//	return u, err
//}

