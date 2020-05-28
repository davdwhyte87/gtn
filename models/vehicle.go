package models

import "gopkg.in/mgo.v2/bson"

// Vehicle ... Data representation of a vehicle
type Vehicle struct {
	ID bson.ObjectId 
	PlateNumber string 
	Images[] string 
	DriverID bson.ObjectId
	CreatedAt string
	UpdatedAt string
}