package controllers

import (
	// "fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	DAO "github.com/davdwhyte87/gtn/dao"
	Models "github.com/davdwhyte87/gtn/models"
	Utils "github.com/davdwhyte87/gtn/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2/bson"
)

var userDao = DAO.UserDAO{}
var walletDao = DAO.WalletDAO{}

var secreteKey string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	secreteKey, _ = os.LookupEnv("SECRETE_KEY")
}

// CreateUser ... This function crea
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// defer r.Body.Close()

	var user Models.User
	// populate the user object with data from requests
	err := Utils.DecodeReq(r, &user)
	// fmt.Printf("%+v\n", user)
	// fmt.Printf("%+v\n", err)
	if err != nil {
		Utils.RespondWithError(w, http.StatusBadRequest, "This is an invalid request object. Cannot decode on server")
		return
	}

	// Validate input data
	ok, errInput := Utils.CreateUserValidator(r)
	if ok == false {
		Utils.RespondWithJSON(w, http.StatusBadRequest, errInput)
		return
	}

	user.ID = bson.NewObjectId()
	user.Confirmed = false
	user.CreatedAt = Utils.CurrentDate()
	user.UpdatedAt = Utils.CurrentDate()
	rand.Seed(time.Now().UnixNano())
	user.PassCode = rand.Intn(9999-939) + 939
	// hash password
	var hashError error
	user.Password, hashError = Utils.HashPassword(user.Password)
	if hashError != nil {
		Utils.RespondWithError(w, http.StatusBadRequest, "Error encrypting password")
	}

	// save user to database
	if err := userDao.Insert(user); err != nil {
		Utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// create a users wallet
	var wallet Models.Wallet
	wallet.ID = bson.NewObjectId()
	wallet.CreatedAt = Utils.CurrentDate()
	wallet.UpdatedAt = wallet.CreatedAt
	wallet.Balance = 0
	wallet.UserID = user.ID
	if creatWalletErr := walletDao.Insert(wallet); err != nil {
		Utils.RespondWithError(w, http.StatusInternalServerError, creatWalletErr.Error())
		return
	}

	// send the user a welcome email
	var emialData Utils.EmailData
	emialData.EmailTo = user.Email
	emialData.ContentData = map[string]interface{}{"Name": user.UserName, "Code": user.PassCode}
	emialData.Template = "hello.html"
	emialData.Title = "Welcome!"
	mailSent := Utils.SendEmail(emialData)
	print(mailSent)
	if mailSent {
		print("mail sent")
	} else {
		print("email not sent")
	}
	// return response
	Utils.RespondWithJSON(w, http.StatusCreated, user.ReturnUser())
	return
}

// ConfrimUser ... allows users to convirm their email
func ConfrimUser(w http.ResponseWriter, r *http.Request) {
	// get code
	type Req struct {
		Code  int
		Email string
	}
	var reqData Req
	// populate the req object with data from requests
	err := Utils.DecodeReq(r, &reqData)
	if err != nil {
		Utils.RespondWithError(w, http.StatusBadRequest, "This is an invalid request object. Cannot decode on server")
		return
	}
	// validate input
	if !Utils.ValidateEmail(reqData.Email) {
		Utils.RespondWithError(w, http.StatusBadRequest, "Bad email")
		return
	}

	// get user

	user, errGetUser := userDao.FindByEmail(reqData.Email)
	if errGetUser != nil {
		Utils.RespondWithError(w, http.StatusBadRequest, errGetUser.Error())
		return
	}

	if user.PassCode == reqData.Code {
		user.Confirmed = true
	} else {
		Utils.RespondWithError(w, http.StatusBadRequest, "Invalid details")
		return
	}
	// update user
	errUpdateUser := userDao.Update(user)
	if errUpdateUser != nil {
		Utils.RespondWithError(w, http.StatusBadRequest, errUpdateUser.Error())
		return
	}

	Utils.RespondWithOk(w, "Account has been activated")
	return
}

// LoginUser ... This function validates a users identity and then gives the user an auth token
func LoginUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user Models.User
	// populate the user object with data from requests
	err := Utils.DecodeReq(r, &user)
	// fmt.Printf("%+v\n", user)
	// fmt.Printf("%+v\n", err)
	if err != nil {
		Utils.RespondWithError(w, http.StatusBadRequest, "This is an invalid request object. Cannot decode on server")
		return
	}

	// Validate input data
	ok, errInput := Utils.LoginUserValidator(r)
	if ok == false {
		Utils.RespondWithJSON(w, http.StatusBadRequest, errInput)
		return
	}

	// get the user from database
	userData, userDataError := userDao.FindByEmail(user.Email)
	if userDataError != nil {
		Utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	// fmt.Printf("%+v\n", userData)
	if Utils.CheckPasswordHash(user.Password, userData.Password) {
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims = jwt.MapClaims{
			"exp":       time.Now().Add(time.Hour * 72).Unix(),
			"email":     userData.Email,
			"id":        userData.ID,
			"user_name": userData.UserName,
		}
		signedString, tokenErr := token.SignedString([]byte(secreteKey))
		if tokenErr != nil {
			Utils.RespondWithError(w, http.StatusInternalServerError, "Error generating token")
			return
		}
		returnData := map[string]string{"token": signedString, "message": "Successful"}
		Utils.RespondWithJSON(w, http.StatusOK, returnData)
	} else {
		Utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
	}
	return
}

// GetFPassCode ... generates new code for reset pass action
func GetFPassCode(w http.ResponseWriter, r *http.Request) {
	// generate a random code
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(9999-939) + 939

	type Req struct {
		Email string
	}
	var reqData Req
	err := Utils.DecodeReq(r, &reqData)
	if err != nil {
		Utils.RespondWithError(w, http.StatusBadRequest, "This is an invalid request object. Cannot decode on server")
		return
	}

	// update the user model on new code
	user, errGetUser := userDao.FindByEmail(reqData.Email)
	if errGetUser != nil {
		Utils.RespondWithError(w, http.StatusNotFound, errGetUser.Error())
		return
	}
	user.PassCode = code

	errUpdateUser := userDao.Update(user)
	if errUpdateUser != nil {
		Utils.RespondWithError(w, http.StatusNotFound, errGetUser.Error())
		return
	}

	// send the user a welcome email
	var emialData Utils.EmailData
	emialData.EmailTo = user.Email
	emialData.ContentData = map[string]interface{}{"Name": user.UserName, "Code": code}
	emialData.Template = "new_code.html"
	emialData.Title = "Welcome!"
	mailSent := Utils.SendEmail(emialData)
	if mailSent {
		print("mail sent")
	} else {
		print("email not sent")
	}

	Utils.RespondWithJSON(w, http.StatusOK, "Code has been sent to your email")
	return
	// send cod to user
}

// ResetPass ... allows user reset password
func ResetPass(w http.ResponseWriter, r *http.Request) {
	// get code
	type Req struct {
		Code        int
		NewPassword string
		Email string
	}
	var reqData Req
	// populate the req object with data from requests
	err := Utils.DecodeReq(r, &reqData)
	if err != nil {
		Utils.RespondWithError(w, http.StatusBadRequest, "This is an invalid request object. Cannot decode on server")
		return
	}

	// get user from email
	user, errGetUser := userDao.FindByEmail(reqData.Email)
	if errGetUser != nil {
		Utils.RespondWithError(w, http.StatusNotFound, errGetUser.Error())
		return
	}
	// check passcode validaity
	if reqData.Code != user.PassCode {
		Utils.RespondWithError(w, http.StatusBadRequest, "Invalid code")
		return
	}
	// get new pass and change password
	// hash password
	var hashError error
	user.Password, hashError = Utils.HashPassword(reqData.NewPassword)
	if hashError != nil {
		Utils.RespondWithError(w, http.StatusBadRequest, "Error encrypting password")
		return
	}

	errUpdateUser := userDao.Update(user)
	if errUpdateUser != nil {
		Utils.RespondWithError(w, http.StatusNotFound, errGetUser.Error())
		return
	}

	Utils.RespondWithOk(w, "Reset Done")
	return
}

// GetUser ... this gets an authenticated user
func GetUser(w http.ResponseWriter, r *http.Request) {
	// get the id of user
	requestUserID := r.Context().Value("user_id")
	requestUserIDString := requestUserID.(string)
	// find user in database
	user, errGetUser := userDao.FindByID(requestUserIDString)
	if errGetUser != nil {
		Utils.RespondWithError(w, http.StatusNotFound, errGetUser.Error())
		return
	}

	// return user
	Utils.RespondWithJSON(w, http.StatusOK, user.ReturnUser())
	return
}
