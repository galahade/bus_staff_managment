package service

import "errors"

const (
	DateString string = "2006-01-02"
	TimeString string = "2006-01-02T15:04:05.000Z"
)
var ServiceError error  = errors.New("service: serivce has error.")

func hasError(err error) (error, bool)  {
	if err != nil {
		return ServiceError, true
	}
	return nil, false
}


