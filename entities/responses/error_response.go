package responses

import (
	"encoding/json"
	"github.com/auction_biding/helpers"
	"net/http"
)

type ErrorResponse struct {
	Data    []interface{} `json:"data"`
	Status  string        `json:"status"`
	Message string        `json:"message"`
}

// HandleError handles error and send response
func HandleError(rw http.ResponseWriter, err error) {
	// build default response
	var response *ErrorResponse
	response = &ErrorResponse{Data: make([]interface{}, 0), Message: "somethingWentWrong",
		Status: http.StatusText(http.StatusInternalServerError)}
	rw.Header().Set("Content-Type", "application/json")
	// set header, message and status
	switch err {
	case helpers.NotValidRequestBody:
		rw.WriteHeader(http.StatusBadRequest)
		response.Message = "invalidRequest"
		response.Status = http.StatusText(http.StatusBadRequest)
	case helpers.NotValidPrice:
		rw.WriteHeader(http.StatusBadRequest)
		response.Message = "invalidprice for bid"
		response.Status = http.StatusText(http.StatusBadRequest)
	default:
		rw.WriteHeader(http.StatusInternalServerError)
	}

	// send response
	err = json.NewEncoder(rw).Encode(response)
	if err != nil{

	}
	return
}

