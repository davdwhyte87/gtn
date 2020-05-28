package middlewares

import (
	"net/http"
	"strings"
	"encoding/json"
	. "github.com/davdwhyte87/gtn/models"
	"gopkg.in/go-playground/validator.v9"
	. "github.com/davdwhyte87/gtn/utils"
	"io/ioutil"
	"bytes"
)

// HandleValidation ... This handles all model validation
func handleValidation(w http.ResponseWriter, r *http.Request) bool {
	b, bodyReaderErr := ioutil.ReadAll(r.Body)
	if bodyReaderErr != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error reading request body")
	}
	var user User
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return false
	}
	// re insert the data read from r.body because r.Body is a read once byte stream
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	validate := validator.New()
  err := validate.Struct(user)
  var valerrors = []string{}
  if err != nil {
	  if strings.Contains(err.Error(), "Description") {
		  valerrors=  append(valerrors, "Description is invalid")
	  }
	  RespondWithError(w, http.StatusBadRequest, err.Error())
	  return false
	}
	return true
}

// InputValidationMiddleware ... This middle ware validates data coming into controllers through http requests
func InputValidationMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validation := handleValidation(w,r)
		if handleValidation(w, r)  {
			nextHandler.ServeHTTP(w, r)
		}
	})
}

