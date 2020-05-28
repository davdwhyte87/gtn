package models

import "gopkg.in/mgo.v2/bson"

// Transaction ... Data representation of a transaction
type Transaction struct {
	ID     bson.ObjectId `bson:"_id"`
	SenderEmail string`bson:"sender_email"`
	RecieverEmail string `bson:"reciever_email"`
	Amount   int
	CreatedAt string
	UpdatedAt string
}
