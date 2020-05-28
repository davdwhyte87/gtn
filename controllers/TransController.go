package controllers

import (
	DAO "github.com/davdwhyte87/gtn/dao"
	Models "github.com/davdwhyte87/gtn/models"
	Utils "github.com/davdwhyte87/gtn/utils"
	"gopkg.in/mgo.v2/bson"
)

var transDao = DAO.TransDAO{}

// CreateTransaction ... This function creates a new transaction
func CreateTransaction(senderEmail string, recieverEmail string, Amount int) bool {
	var trans Models.Transaction
	if Amount < 0 {
		return false
	}

	trans.ID = bson.NewObjectId()
	trans.Amount = Amount
	trans.CreatedAt = Utils.CurrentDate()
	trans.UpdatedAt = Utils.CurrentDate()
	trans.SenderEmail = senderEmail
	trans.RecieverEmail = recieverEmail

	err := transDao.Insert(trans)
	if err != nil{
		print(err.Error())
		return false
	}

	return true
}
