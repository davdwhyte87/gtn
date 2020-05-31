package dao

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mgo "gopkg.in/mgo.v2"
)

// DatabaseDAO ... This is the the data access model
type DatabaseDAO struct {
	Server   string
	Database string
}

var database *mongo.Database
var db *mgo.Database

// UserCollection ...
var UserCollection mongo.Collection
// WalletCollection ...
var WalletCollection mongo.Collection
// Ctx ...
var Ctx context.Context

// database collection
const (
	COLLECTION = "verxv"
)

//Connect ... Establish a connection to database
func (m *DatabaseDAO) Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("SERVER")))
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		// panic(err)
		print((err.Error()))
	}
	// defer ctxcancel()
	
	Ctx = ctx
	database = client.Database("gtn")
	UserCollection = *database.Collection("user")
	WalletCollection = *database.Collection("wallet")
	print("connected")
	// defer client.Disconnect(ctx)
}

// func (m *DatabaseDAO)Connect() {
// 	dialInfo, err := mgo.ParseURL(m.Server)
// 	if err != nil {
// 		print(err.Error())
// 	}
// 	dialInfo.Timeout = time.Duration(30)
// 	session, err := mgo.DialWithInfo(dialInfo)
// 	if err != nil {
// 		// log.Fatal(err)
// 	}
// 	db = session.DB(m.Database)
// }
