package main

import (
	"log"
	"net/http"

	"github.com/davdwhyte87/gtn/controllers"
	"github.com/davdwhyte87/gtn/dao"
	"github.com/gorilla/mux"

	"os"

	"github.com/davdwhyte87/gtn/middlewares"

	"github.com/joho/godotenv"
)

var databaseDao = dao.DatabaseDAO{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	Server, _ := os.LookupEnv("SERVER")
	database, _ := os.LookupEnv("DATABASE")
	databaseDao.Server = Server
	databaseDao.Database = database
	databaseDao.Connect()
}

func main() {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()
	userRouter := v1.PathPrefix("/user").Subrouter()
	// vehicleRouter := v1.PathPrefix("/vehicle").Subrouter()
	// becomeAdriverRouter := r.PathPrefix("/api/v1/").Subrouter()

	userRouter.HandleFunc("/signup", controllers.CreateUser).Methods("POST")
	// signUpRouter.Use(InputValidationMiddleware)
	userRouter.HandleFunc("/signin", controllers.LoginUser).Methods("POST")
	userRouter.HandleFunc("/confirm", controllers.ConfrimUser).Methods("POST")
	userRouter.HandleFunc("/send_code", controllers.GetFPassCode).Methods("POST")
	userRouter.HandleFunc("/reset_pass", controllers.ResetPass).Methods("POST")
	userRouter.HandleFunc("/me",
		middlewares.MultipleMiddleware(controllers.GetUser, middlewares.AuthenticationMiddleware)).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
