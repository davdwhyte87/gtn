package utils

// import "fmt"
import (
	"fmt"
	"net/http"

	"reflect"
	DAO "github.com/davdwhyte87/gtn/dao"
	"regexp"
	"github.com/thedevsaddam/govalidator"
)

var dao = DAO.UserDAO{}

func init() {
	govalidator.AddCustomRule("user_exists", func(field string, rule string, message string, value interface{}) error {
		valSlice := value.(string)
		println(valSlice)
		user, _ := dao.FindByEmail(valSlice)
		if user.Email != "" {
			return fmt.Errorf("This user email exists")
		}
		// if err != nil {
		// 	return fmt.Errorf(err.Error())

		// }
		return nil
	})
}

// CreateUserValidator ...
func CreateUserValidator(r *http.Request) (bool, interface{}) {
	rules := govalidator.MapData{
		"UserName": []string{"required", "between:3,50"},
		"Email":    []string{"required", "min:4", "max:100", "email", "user_exists"},
		"Password": []string{"required", "min:4", "max:20"},
	}
	// var user Models.User
	data := make(map[string]interface{}, 0)
	// messages := govalidator.MapData{
	// 	"Name": []string{"Name field is required", "Name should be between 3 to 50 charachers"},
	// 	"Email":    []string{"Email field is required", "", "", "A valid email is required"},
	// 	"Password": []string {"Password is required", "", ""},
	// }

	opts := govalidator.Options{
		Request: r,     // request object
		Rules:   rules, // rules map
		Data:    &data,
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	err := map[string]interface{}{"validationError": e}

	if len(e) == 0 {
		return true, err
	}
	return false, err
}

// LoginUserValidator ...
func LoginUserValidator(r *http.Request) (bool, interface{}) {
	rules := govalidator.MapData{
		"Password": []string{"required", "min:4", "max:20"},
		"Email":    []string{"required", "min:4", "max:100", "email"},
	}
	// var user Models.User
	data := make(map[string]interface{}, 0)

	opts := govalidator.Options{
		Request: r,     // request object
		Rules:   rules, // rules map
		Data:    &data,
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	err := map[string]interface{}{"validationError": e}

	if len(e) == 0 {
		return true, err
	}
	return false, err
}


// ValidateEmail ... checks if an email is written correctly
func ValidateEmail(email string) bool{

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

// ValidateInt .. checks int type
func ValidateInt(x int) bool{
	xt := reflect.TypeOf(x).Kind()
	if xt == reflect.Int {
		return true
	}
	return false
}