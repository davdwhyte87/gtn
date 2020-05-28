package dao

import (
	"github.com/davdwhyte87/gtn/models"
	"gopkg.in/mgo.v2/bson"
)

// UserDAO ...
type UserDAO struct {

}

// FindAll ... get list of users
func (m *UserDAO) FindAll() ([]models.User, error) {
	var users []models.User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// FindByID ... get a user by its id
func (m *UserDAO) FindByID(id string) (models.User, error) {
	var user models.User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// FindByEmail ... get a user by the email
func (m *UserDAO) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := db.C(COLLECTION).Find(bson.M{"email":email}).One(&user)
	return user, err
}

// Insert a user into database
func (m *UserDAO) Insert(user models.User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

// Delete an existing user
func (m *UserDAO) Delete(user models.User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

// Update an existing user
func (m *UserDAO) Update(user models.User) error {
	err := db.C(COLLECTION).Update(bson.M{"_id":user.ID}, &user)
	return err
}