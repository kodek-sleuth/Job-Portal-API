package models

import (
	_"github.com/jinzhu/gorm"
)

type Users struct {
	Base
	Name string `json:"company"`
	Email string `json:"criteria"`
	Password string `json:"location"`
}
//


func (u *Users) CreateUsersCollection()(*Users, interface{}) {
	db, err := SQLConnection()
	if err != nil{
		return u, err
	}
	db.Create(u)
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
