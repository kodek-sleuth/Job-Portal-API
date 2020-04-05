package models

import (
	//"fmt"
	_ "github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Name string `sql:"type:VARCHAR(100);not null"`
	Email string `sql:"unique;unique_index;not null"`
	Password string `sql:"type:VARCHAR(255);not null"`
	Base
}
//

type User struct {
	Name string
	Email string
	Password string
}


func (u *Users) CreateUsers()(*Users, error) {

	db, err := SQLConnection()
	if err != nil{
		return u, err
	}

	// Hash Password
	password := []byte(u.Password)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return u, err
	}
	u.Password = string(hash)

	// Create user
	if err := db.Create(u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (u *Users) CheckLoginCredentials()(*Users, error){
	var user User
	db, err := SQLConnection()
	if err != nil{
		return u, err
	}

	if err = db.Where("email = ?", u.Email).First(&user).Error; err != nil {
		return u, err
	}

	// ComparePasswords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return u, err
	}
	return u, nil
}

//func GetJobs() ([]Job, interface{}) {
//	var jobs []Job
//
//	cursor, err := collection.Find(context.TODO(), bson.M{})
//
//	if err != nil {
//		return jobs, err
//	}
//
//	defer cursor.Close(context.TODO())
//
//	for cursor.Next(context.TODO()) {
//		var job Job
//		err = cursor.Decode(&job)
//		if err != nil {
//			return jobs, err
//		}
//		jobs = append(jobs, job)
//	}
//
//	return jobs, err
//}
//
//func (j *Job) UpdateJobCollection(id string) (*Job, interface{}) {
//	var updatedDocument bson.M
//	objID, err := primitive.ObjectIDFromHex(id) // change string to object ID
//	if err != nil{
//		return j, err
//	}
//
//	filter := bson.M{ "_id": objID }
//	update := bson.M{"$set": bson.M{ "company": j.Company,  "salary": j.Salary,
//		"criteria": j.Criteria , "location": j.Location, "description": j.Description } }
//
//	err = collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&updatedDocument)
//
//	if err != nil {
//		if err.Error() == "mongo: no documents in result" {
//			return j, "no job found to update"
//		}
//		return j, err
//	}
//	return j, err
//}
//
//func GetJob(id string) (interface{}, interface{}) {
//	var updatedDocument bson.M
//	objID, err := primitive.ObjectIDFromHex(id) // change string to object ID
//
//	if err != nil{
//		return updatedDocument, err
//	}
//
//	filter := bson.M{ "_id": objID }
//
//	err = collection.FindOne(context.TODO(), filter).Decode(&updatedDocument)
//
//	if err != nil{
//		if err.Error() == "mongo: no documents in result" {
//			return updatedDocument, "no job found"
//		}
//		return updatedDocument, err
//	}
//
//	return updatedDocument, err
//}
//
//func DeleteJob(id string) (interface{}, interface{}){
//	var updatedDocument bson.M
//	objID, err := primitive.ObjectIDFromHex(id) // change string to object ID
//
//	if err != nil{
//		return updatedDocument, err
//	}
//
//	filter := bson.M{ "_id": objID }
//
//	err = collection.FindOneAndDelete(context.TODO(), filter).Decode(&updatedDocument)
//
//	if err != nil{
//		if err.Error() == "mongo: no documents in result" {
//			return updatedDocument, "no job found"
//		}
//		return updatedDocument, err
//	}
//
//	return updatedDocument, err
//}
