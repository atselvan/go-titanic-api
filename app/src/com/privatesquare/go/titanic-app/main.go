// Copyright 2019 Allan Tony Selvan. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Titanic Passengers API
//
// This documentation describes the API to manage passengers of Titanic
//
//     Schemes: http, https
//     Host: localhost:8000
//     BasePath: /
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Allan Tony Selvan <atselvan99@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basicAuth: []
//
//     SecurityDefinitions:
//       basicAuth:
//	       type: basic
//
// swagger:meta
package main

import (
	b "com/privatesquare/go/titanic-app/backend"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", b.HomeHandler).Methods("GET")

	// swagger:operation GET /passengers passengers getPassengers
	// ---
	// summary: Returns info of all passengers
	// description: Returns info of all passengers
	// responses:
	//   "200":
	//     "$ref": "#/responses/success"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "500":
	//     "$ref": "#/responses/internalServerError"
	//
	router.HandleFunc("/passengers", b.GetPassengersHandler).Methods("GET")

	// swagger:operation GET /passenger/{uuid} passenger getPassenger
	// ---
	// summary: Returns passenger's info
	// description: UUID should be valid else the request will return 400 Bad Request status
	// parameters:
	// - name: uuid
	//   in: path
	//   description: uuid of the passenger
	//   type: string
	//   required:
	// responses:
	//   "200":
	//     "$ref": "#/responses/success"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "500":
	//     "$ref": "#/responses/internalServerError"

	router.HandleFunc("/passenger/{uuid}", b.GetPassengerHandler).Methods("GET")

	// swagger:operation POST /passenger passenger addPassenger
	// ---
	// summary: Add passenger's info to the record
	// description: UUID should be valid else the request will return 400 Bad Request status
	// responses:
	//   "200":
	//     "$ref": "#/responses/success"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "500":
	//     "$ref": "#/responses/internalServerError"
	router.HandleFunc("/passenger", b.AddPassengerHandler).Methods("POST")

	// swagger:operation PUT /passenger passenger updatePassenger
	// ---
	// summary: Update a passenger's info
	// responses:
	//   "200":
	//     "$ref": "#/responses/success"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "500":
	//     "$ref": "#/responses/internalServerError"
	//
	router.HandleFunc("/passenger/{uuid}", b.UpdatePassengerHandher).Methods("PUT")

	// swagger:operation DELETE /passenger/{uuid} passenger deletePassenger
	// ---
	// summary: Delete a passenger's info from the record
	// responses:
	//   "200":
	//     "$ref": "#/responses/success"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "500":
	//     "$ref": "#/responses/internalServerError"
	//
	router.HandleFunc("/passenger/{uuid}", b.DeletePassengerHandler).Methods("DELETE")

	router.NotFoundHandler = http.HandlerFunc(b.PageNotFoundHandler)

	fs := http.FileServer(http.Dir("./swaggerui"))
	router.PathPrefix("/swaggerui").Handler(http.StripPrefix("/swaggerui", fs))

	log.Fatal(http.ListenAndServe(":8000", router))
}
