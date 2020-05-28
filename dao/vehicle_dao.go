package dao

import (
	Models "github.com/davdwhyte87/gtn/models"
	"gopkg.in/mgo.v2/bson"
)

// VehicleDAO ... 
type VehicleDAO struct {

}

// FindAll ... get list of vehicles
func (m *VehicleDAO) FindAll() ([]Models.Vehicle, error) {
	var vehicles []Models.Vehicle
	err := db.C(COLLECTION).Find(bson.M{}).All(&vehicles)
	return vehicles, err
}

// FindByID ... get a vehicle by its id
func (m *VehicleDAO) FindByID(id string) (Models.Vehicle, error) {
	var vehicle Models.Vehicle
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&vehicle)
	return vehicle, err
}

// FindByEmail ... get a vehicle by the email
func (m *VehicleDAO) FindByEmail(email string) (Models.Vehicle, error) {
	var vehicle Models.Vehicle

	err := db.C(COLLECTION).Find(bson.M{"email":email}).One(&vehicle)
	return vehicle, err
}

// Insert a vehicle into database
func (m *VehicleDAO) Insert(vehicle Models.Vehicle) error {
	err := db.C(COLLECTION).Insert(&vehicle)
	return err
}

// Delete an existing vehicle
func (m *VehicleDAO) Delete(vehicle Models.Vehicle) error {
	err := db.C(COLLECTION).Remove(&vehicle)
	return err
}

// Update an existing vehicle
func (m *VehicleDAO) Update(vehicle Models.Vehicle) error {
	err := db.C(COLLECTION).Update(bson.M{"_id":vehicle.ID}, &vehicle)
	return err
}