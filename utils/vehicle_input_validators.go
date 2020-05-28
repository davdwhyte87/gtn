package utils
// import "fmt"
import (

	// "fmt"
	"net/http"
	// "reflect"
	"github.com/thedevsaddam/govalidator"
)

// CreateVehicleValidator ...
func CreateVehicleValidator(r *http.Request) (bool, interface{}){
	rules := govalidator.MapData{
		"Images": []string{"required", "between:3,50"},
		"PlateNumber":    []string{"required", "min:4", "max:100"},
	}

	data := make(map[string]interface{}, 0)
	// messages := govalidator.MapData{
	// 	"Name": []string{"Name field is required", "Name should be between 3 to 50 charachers"},
	// 	"Email":    []string{"Email field is required", "", "", "A valid email is required"},
	// 	"Password": []string {"Password is required", "", ""},
	// }

	opts := govalidator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Data:            &data,
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	err := map[string]interface{}{"validationError": e}
	
	if len(e) == 0 {
		return true, err
	}
	return false, err
}

