package models

import "gopkg.in/mgo.v2/bson"

// Trip ... Data representation of a vehicle
type Trip struct {
	ID bson.ObjectId 
	DriverID bson.ObjectId
	Riders []bson.ObjectId
	CreatedAt string
	UpdatedAt string
}