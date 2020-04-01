package helpers

import (
	_ "fmt"
	"github.com/fatih/structs"
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
		return "Name, email and password", false
	}

	return "", true
}

