package models

import 	"go.mongodb.org/mongo-driver/bson/primitive"

// User ... this is a representation of a user on the server
type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	UserName  string        `bson:"user_name"`
	Email     string        `bson:"email"`
	Password  string        `bson:"password"`
	Confirmed bool          `bson:"confirmed"`
	PassCode  int
	CreatedAt string
	UpdatedAt string
}

// ReturnUser ... this returns a user safely
func (user User) ReturnUser() User {
	user.PassCode = 0
	user.Password = ""
	return user
}
