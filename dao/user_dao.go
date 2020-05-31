package dao

import (
	"github.com/davdwhyte87/gtn/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// gbson "gopkg.in/mgo.v2/bson"
	// "fmt"
)

// UserDAO ...
type UserDAO struct {
}

// FindAll ... get list of users
func (m *UserDAO) FindAll() ([]models.User, error) {
	var users []models.User
	cursor, err := UserCollection.Find(Ctx, bson.M{})
	if err != nil {
		print(err.Error())
	}
	err = cursor.All(Ctx, &users)
	return users, err
}

// FindByID ... get a user by its id
func (m *UserDAO) FindByID(id string) (models.User, error) {
	var user models.User
	docID, _ := primitive.ObjectIDFromHex(id)
	err := UserCollection.FindOne(Ctx, bson.M{"_id": docID}).Decode(&user)
	// print(err.Error())
	// err := db.C(COLLECTION).FindId(gbson.ObjectIdHex(id)).One(&user)
	return user, err
}

// FindByEmail ... get a user by the email
func (m *UserDAO) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := UserCollection.FindOne(Ctx, bson.M{"email": email}).Decode(&user)
	// err := db.C(COLLECTION).Find(bson.M{"email": email}).One(&user)
	return user, err
}

// Insert a user into database
func (m *UserDAO) Insert(user models.User) error {
	userb, _ := bson.Marshal(user)
	_, err := UserCollection.InsertOne(Ctx, userb)
	return err
}

// Delete an existing user
func (m *UserDAO) Delete(user models.User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

// Update an existing user
func (m *UserDAO) Update(user models.User) error {
	docID, _ := primitive.ObjectIDFromHex(user.ID.Hex())
	userb, _ := bson.Marshal(user)
	_, err := UserCollection.UpdateOne(Ctx, bson.M{"_id": docID}, bson.M{"$set":userb})
	return err
}
