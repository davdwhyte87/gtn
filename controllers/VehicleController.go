package controllers

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	DAO "github.com/davdwhyte87/gtn/dao"
	Models "github.com/davdwhyte87/gtn/models"
	Utils  "github.com/davdwhyte87/gtn/utils"
	"time"
)


var vehicleDao = DAO.VehicleDAO{}

// CreateVehicle ... This creates a new Vehicle
func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle Models.Vehicle

	requestUserID := r.Context().Value("user_id")
	requestUserIDString := requestUserID.(string)
	if requestUserID == "" {
		Utils.RespondWithError(w, http.StatusNotFound, "An error occured")
	}

	err := Utils.DecodeReq(r, &vehicle)
	// fmt.Printf("%+v\n", err)
	if err != nil {
		Utils.RespondWithError(w, http.StatusInternalServerError, "This is an invalid request object. Cannot decode on server")
		return
	}

	// Validate input data 
	ok, errInput := Utils.CreateVehicleValidator(r)
	if ok == false {
		Utils.RespondWithJSON(w, http.StatusBadRequest, errInput)
		return
	}

	// other fields
	dt := time.Now()
	vehicle.DriverID = bson.ObjectId(requestUserIDString)
	vehicle.CreatedAt = dt.Format("01-02-2006 15:04:05")
	vehicle.UpdatedAt = vehicle.CreatedAt 


	// save vehicle to database
	if err := vehicleDao.Insert(vehicle); err != nil {
		Utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	Utils.RespondWithOk(w, "Vehicle Created")
	return
}

// UpdateVehicle ... This creates a new Vehicle
func UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle Models.Vehicle

	err := Utils.DecodeReq(r, &vehicle)
	// fmt.Printf("%+v\n", err)
	if err != nil {
		Utils.RespondWithError(w, http.StatusInternalServerError, "This is an invalid request object. Cannot decode on server")
		return
	}

	// Validate input data 
	ok, errInput := Utils.CreateVehicleValidator(r)
	if ok == false {
		Utils.RespondWithJSON(w, http.StatusBadRequest, errInput)
		return
	}

	// update time
	dt := time.Now()
	vehicle.UpdatedAt = dt.Format("01-02-2006 15:04:05")


	// save vehicle to database
	if err := vehicleDao.Update(vehicle); err != nil {
		Utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	Utils.RespondWithJSON(w, http.StatusOK, vehicle )
	return
}
