package models

import (
	"context"
	//"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"go.mongodb.org/mongo-driver/mongo"

	//"reflect"

	//"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	//"reflect"

	//"net/http"
)

// Bson.D - D reps a bson doc containing ordered elements
// Bson.M - M reps a bson doc containing unordered elements

type Job struct {
	ID  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Company string `json:"company" bson:"company"`
	Criteria string `json:"criteria" bson:"criteria"`
	Location string `json:"location" bson:"location"`
	Description string `json:"description" bson:"description"`
	Salary string `json:"salary" bson:"salary"`
}

var collection = MongoConnection("jobportal", "jobs")

func (j *Job) CreateJobCollection()(interface{}, interface{}) {
	result, err := collection.InsertOne(context.TODO(), j) // bson.M is representation of BSON data(Unordered map)
	return result.InsertedID, err
}

func GetJobs() ([]Job, interface{}) {
	var jobs []Job

	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		return jobs, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var job Job
		err = cursor.Decode(&job)
		if err != nil {
			return jobs, err
		}
		jobs = append(jobs, job)
	}

	return jobs, err
}

func (j *Job) UpdateJobCollection(id string) (*Job, interface{}) {
	var updatedDocument bson.M
	objID, err := primitive.ObjectIDFromHex(id) // change string to object ID
	if err != nil{
		return j, err
	}

	filter := bson.M{ "_id": objID }
	update := bson.M{"$set": bson.M{ "company": j.Company,  "salary": j.Salary,
		"criteria": j.Criteria , "location": j.Location, "description": j.Description } }

	err = collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&updatedDocument)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return j, "no job found to update"
		}
		return j, err
	}
	return j, err
}

func GetJob(id string) (interface{}, interface{}) {
	var updatedDocument bson.M
	objID, err := primitive.ObjectIDFromHex(id) // change string to object ID

	if err != nil{
		return updatedDocument, err
	}

	filter := bson.M{ "_id": objID }

	err = collection.FindOne(context.TODO(), filter).Decode(&updatedDocument)

	if err != nil{
		if err.Error() == "mongo: no documents in result" {
			return updatedDocument, "no job found"
		}
		return updatedDocument, err
	}

	return updatedDocument, err
}

func DeleteJob(id string) (interface{}, interface{}){
	var updatedDocument bson.M
	objID, err := primitive.ObjectIDFromHex(id) // change string to object ID

	if err != nil{
		return updatedDocument, err
	}

	filter := bson.M{ "_id": objID }

	err = collection.FindOneAndDelete(context.TODO(), filter).Decode(&updatedDocument)

	if err != nil{
		if err.Error() == "mongo: no documents in result" {
			return updatedDocument, "no job found"
		}
		return updatedDocument, err
	}

	return updatedDocument, err
}
