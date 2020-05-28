package controllers

import (
	"fmt"
	"net/http"
	"time"

	DAO "github.com/davdwhyte87/gtn/dao"
	Models "github.com/davdwhyte87/gtn/models"
	Utils "github.com/davdwhyte87/gtn/utils"
	"gopkg.in/mgo.v2/bson"
)

var wallerDao = DAO.WalletDAO{}

// GetWallet ... This function gets a users wallet with the ID
func GetWallet(w http.ResponseWriter, r *http.Request) {
	var wallet Models.Wallet

	// get user id from request
	requestUserID := r.Context().Value("user_id")
	requestUserIDString := requestUserID.(string)
	// fmt.Print("UserId", requestUserIDString)
	if requestUserID == "" {
		Utils.RespondWithError(w, http.StatusNotFound, "An error occured")
	}

	// get wallet from database
	var getWalletError error
	wallet, getWalletError = wallerDao.FindByUserID(bson.ObjectIdHex(requestUserIDString))
	// fmt.Printf("WalletData : %v\n", wallet)
	if getWalletError != nil {
		Utils.RespondWithError(w, http.StatusNotFound, "Wallet not found")
		return
	}

	Utils.RespondWithJSON(w, http.StatusOK, wallet)
	return
}

// GetAllWallets ... Gets all the wallets available in the system
func GetAllWallets(w http.ResponseWriter, r *http.Request) {

	wallet, getWalletError := wallerDao.FindAll()
	// fmt.Printf("WalletData : %v\n", wallet)
	if getWalletError != nil {
		Utils.RespondWithError(w, http.StatusNotFound, getWalletError.Error())
		return
	}

	Utils.RespondWithJSON(w, http.StatusOK, wallet)
	return
}

// LoadWallet ... Allows a user to load his wallet
func LoadWallet(w http.ResponseWriter, r *http.Request) {
	type LoadReq struct {
		Amount int
	}
	var reqData LoadReq
	// populate the user object with data from requests
	err := Utils.DecodeReq(r, &reqData)
	if err != nil {
		Utils.RespondWithError(w, http.StatusBadRequest, "This is an invalid request object. Cannot decode on server")
		return
	}

	if reqData.Amount < 0 {
		Utils.RespondWithError(w, http.StatusBadRequest, "Invalid amount")
		return
	}

	// get user id from request
	requestUserID := r.Context().Value("user_id")
	requestUserIDString := requestUserID.(string)
	// get money from reserve and give to user wallet

	wallet, getWalletErr := wallerDao.FindByUserID(bson.ObjectIdHex(requestUserIDString))
	if getWalletErr != nil {
		Utils.RespondWithError(w, http.StatusInternalServerError, "Cannot get wallet successfully")
		return
	}
	wallet.Balance = reqData.Amount
	dt := time.Now()
	wallet.UpdatedAt = dt.Format("01-02-2006 15:04:05")
	walletUpdateErr := wallerDao.Update(wallet)
	if walletUpdateErr != nil {
		fmt.Printf("%v\n", walletUpdateErr)
		Utils.RespondWithError(w, http.StatusInternalServerError, "Unable to complete transaction")
		return
	}

	// return success
	Utils.RespondWithOk(w, "Transaction successfull")
	return
}

// Withdraw ... users to get their cash
func Withdraw() {

}
