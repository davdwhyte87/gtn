package utils

import "time"


// CurrentDate ... return the current date and time
func CurrentDate() string{
	dt := time.Now()
	datec := dt.Format("01-02-2006 15:04:05")
	return datec
}