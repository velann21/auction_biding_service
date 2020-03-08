package helpers

import "errors"

var (
	// ErrParamMissing required field missing error
	NotValidRequestBody = errors.New("Invalid Request")
	NotValidPrice = errors.New("Invalid Price")
)

