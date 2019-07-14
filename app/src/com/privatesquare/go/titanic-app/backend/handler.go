package backend

import (
	m "com/privatesquare/go/titanic-app/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

const (
	successStatusCode             = 200
	createdStatusCode             = 201
	notFoundStatusCode            = 404
	badRequestStatusCode          = 400
	internalServerErrorStatusCode = 500

	successStatus             = "200 OK"
	createdStatus             = "201 Created"
	notFoundStatus            = "404 Not Found"
	badReqStatus              = "400 Bad Request"
	internalServerErrorStatus = "500 Internal Server Error"

	jsonMarshalErrorMessage   = "JSON Marshal Error"
	jsonUnmarshalErrorMessage = "JSON Unmarshal Error"
)

func writeHTTPResponse(w http.ResponseWriter, successCode int, successStatus, messageString string) {
	w.Header().Set("Content-Type", "application/json")
	message, err := json.Marshal(m.SuccessResp{Status: successStatus, Message: messageString})
	HandleError(err, jsonMarshalErrorMessage)
	w.WriteHeader(successCode)
	w.Write(message)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	logger("handling GET /")
	message := "Welcome to the Titanic Passengers Application"
	writeHTTPResponse(w, successStatusCode, successStatus, message)
	logger("finished handling GET /")
}

func PageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	logger("handling page not found")
	message := "Page not found"
	writeHTTPResponse(w, notFoundStatusCode, notFoundStatus, message)
	logger("finished handling page not found")
}

func DBErrorHandler(w http.ResponseWriter, r *http.Request) {
	logger("handling database error")
	message := "Database Error. Check Server logs for more details"
	writeHTTPResponse(w, internalServerErrorStatusCode, internalServerErrorStatus, message)
	logger("finished handling database error")
	return
}

func GetPassengersHandler(w http.ResponseWriter, r *http.Request) {
	logger("handling GET /passengers")
	w.Header().Set("Content-Type", "application/json")
	passengers, error := GetPassengersService()
	if error == nil {
		json.NewEncoder(w).Encode(passengers)
	} else {
		DBErrorHandler(w, r)
	}
	logger("finished handling GET /passengers")
}

func GetPassengerHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuid := params["uuid"]
	logger(fmt.Sprintf("handling GET /passenger/%s", uuid))
	w.Header().Set("Content-Type", "application/json")
	if IsValidUUID(uuid) {
		passengerInfo, error := GetPassengerService(uuid)
		if error != nil {
			DBErrorHandler(w, r)
		} else if passengerInfo.Uuid != "" {
			json.NewEncoder(w).Encode(passengerInfo)
		} else {
			message := fmt.Sprintf("Passenger with uuid %s does not exist", params["uuid"])
			writeHTTPResponse(w, notFoundStatusCode, notFoundStatus, message)
		}
	} else {
		message := fmt.Sprintf("The UUID : %s is not valid", params["uuid"])
		writeHTTPResponse(w, badRequestStatusCode, badReqStatus, message)
	}
	logger(fmt.Sprintf("finished handling GET /passenger/%s", uuid))
}

func AddPassengerHandler(w http.ResponseWriter, r *http.Request) {
	logger("handling POST /passenger")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var passengerInfo m.PassengerInfo
	err := json.Unmarshal(reqBody, &passengerInfo)
	HandleError(err, jsonUnmarshalErrorMessage)
	message := fmt.Sprintf("Passenger info is added with uuid %s", AddPassengerService(passengerInfo))
	writeHTTPResponse(w, createdStatusCode, createdStatus, message)
	logger("finished handling POST /passenger")
}

//TODO
func UpdatePassengerHandher(w http.ResponseWriter, r *http.Request) {

}

//TODO
func DeletePassengerHandler(w http.ResponseWriter, r *http.Request) {

}
