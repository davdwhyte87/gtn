package dao

import (
	"github.com/davdwhyte87/gtn/models"
	"gopkg.in/mgo.v2/bson"
)

// WalletDAO ...
type WalletDAO struct {
}

// FindByUserID ... get a user by the email
func (m *WalletDAO) FindByUserID(id bson.ObjectId) (models.Wallet, error) {
	var wallet models.Wallet

	err := db.C(COLLECTION).Find(bson.M{"userId": id}).One(&wallet)
	return wallet, err
}

// FindByEmail ... get a user by the email
func (m *WalletDAO) FindByEmail(email string) (models.Wallet, error) {
	var wallet models.Wallet

	err := db.C(COLLECTION).Find(bson.M{"email": email}).One(&wallet)
	return wallet, err
}

// Insert a user into database
func (m *WalletDAO) Insert(wallet models.Wallet) error {
	_, err := WalletCollection.InsertOne(Ctx, wallet)
	return err
}

// FindAll ... get list of users
func (m *WalletDAO) FindAll() ([]models.Wallet, error) {
	var wallets []models.Wallet
	err := db.C(COLLECTION).Find(bson.M{}).All(&wallets)
	return wallets, err
}

// Update an existing user
func (m *WalletDAO) Update(wallet models.Wallet) error {
	err := db.C(COLLECTION).Update(bson.M{"_id": wallet.ID}, &wallet)
	return err
}
