package dao

import (
	"github.com/davdwhyte87/gtn/models"
	"gopkg.in/mgo.v2/bson"
	
)

// TransDAO ...
type TransDAO struct {
}

// FindByID ... get a transaction by the ID
func (m *TransDAO) FindByID(id bson.ObjectId) (models.Transaction, error) {
	var trans models.Transaction

	err := db.C(COLLECTION).Find(bson.M{"_id": id}).One(&trans)
	return trans, err
}

// FindByEmail ... get a transaction by the email
func (m *TransDAO) FindByEmail(email string) (models.Transaction, error) {
	var trans models.Transaction

	err := db.C(COLLECTION).Find(bson.M{"sender_email": email, "reciever_email": email}).One(&trans)
	return trans, err
}

// Insert a transaction into database
func (m *TransDAO) Insert(trans models.Transaction) error {
	err := db.C(COLLECTION).Insert(&trans)
	return err
}

// FindAll ... get all transactions
func (m *TransDAO) FindAll() ([]models.Transaction, error) {
	var trans []models.Transaction
	err := db.C(COLLECTION).Find(bson.M{}).All(&trans)
	return trans, err
}

// Update an existing user
// func (m *WalletDAO) Update(wallet models.Wallet) error {
// 	err := db.C(COLLECTION).Update(bson.M{"_id": wallet.ID}, &wallet)
// 	return err
// }
