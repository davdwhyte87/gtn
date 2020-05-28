package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

// DatabaseDAO ... This is the the data access model
type DatabaseDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// database collection
const (
	COLLECTION = "verxv"
)

// Connect ... Establish a connection to database
func (m *DatabaseDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}
