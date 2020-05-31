package models


import 	"go.mongodb.org/mongo-driver/bson/primitive"
// Wallet ... Data representation of a users wallet
type Wallet struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserID primitive.ObjectID `bson:"userId"`
	// Riders []bson.ObjectId
	Balance   int
	CreatedAt string
	UpdatedAt string
}
