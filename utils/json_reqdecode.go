package utils

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"bytes"
	"fmt"
)


// DecodeReq ... This helps decode a json request body into an interface
func DecodeReq(r *http.Request, model interface{}) interface{} {
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, model)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	if  err != nil {
		fmt.Printf("%+v\n", err.Error())
		return err
	}
	return err
}