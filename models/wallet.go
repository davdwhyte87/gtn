package models

import "gopkg.in/mgo.v2/bson"

// Wallet ... Data representation of a users wallet
type Wallet struct {
	ID     bson.ObjectId `bson:"_id"`
	UserID bson.ObjectId `bson:"userId"`
	// Riders []bson.ObjectId
	Balance   int
	CreatedAt string
	UpdatedAt string
}
