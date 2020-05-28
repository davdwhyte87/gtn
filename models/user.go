package models

import "gopkg.in/mgo.v2/bson"

// User ... this is a representation of a user on the server
type User struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	UserName  string        `bson:"name" json:"name" validate:"required,min=2,max=100"`
	Email     string        `bson:"email" json:"email" validate:"required,email"`
	Password  string        `bson:"password" json:"password" validate:"required,alpha"`
	Confirmed bool          `bson:"confirmed" json:"confirmed"`
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
