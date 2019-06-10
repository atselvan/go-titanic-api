package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"regexp"
)

const (
	port = 5432
)

var (
	psqlInfo string
	host     = os.Getenv("DB_HOST")
	dbName   = os.Getenv("DB_NAME")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
)

type PassengerInfo struct {
	Uuid     string  `json:"uuid"`
	Survived bool    `json:"survived"`
	Pclass   string  `json:"pclass"`
	Name     string  `json:"name"`
	Sex      string  `json:"sex"`
	Age      float64 `json:"Age"`
	SSA      int     `json:"ssa"`
	PCA      int     `json:"pca"`
	Fare     float64 `json:"fare"`
}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/passengers", GetPassengersHandler).Methods("GET")
	router.HandleFunc("/passenger/{uuid}", GetPassengerHandler).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(PageNotFoundHandler)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	logger("handling GET /")
	w.Header().Set("Content-Type", "application/json")
	message, err := json.Marshal(Message{Status: "200 OK", Message: "Welcome to the Titanic Passengers Application"})
	HandleError(err, "JSON Unmarshal Error")
	w.Write(message)
	logger("finished handling GET /")
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
			message, err := json.Marshal(Message{Status: "404 Not Found", Message: fmt.Sprintf("Passenger with uuid %s does not exist", params["uuid"])})
			HandleError(err, "JSON Unmarshal Error")
			//http.Error(w, string(message), http.StatusNotFound)
			w.WriteHeader(404)
			w.Write(message)
		}
	} else {
		message, err := json.Marshal(Message{Status: "400 Bad Request", Message: fmt.Sprintf("The UUID : %s is not valid", params["uuid"])})
		HandleError(err, "JSON Unmarshal Error")
		//http.Error(w, string(message), http.StatusBadRequest)
		w.WriteHeader(400)
		w.Write(message)
	}
	logger(fmt.Sprintf("finished handling GET /passenger/%s", uuid))
}

func PageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	logger("handling page not found")
	w.Header().Set("Content-Type", "application/json")
	message, err := json.Marshal(Message{Status: "404 Not Found", Message: "Page not found"})
	HandleError(err, "JSON Unmarshal Error")
	//http.Error(w, string(message), http.StatusNotFound)
	w.WriteHeader(404)
	w.Write(message)
	logger("finished handling page not found")
}

func DBErrorHandler(w http.ResponseWriter, r *http.Request) {
	logger("handling database error")
	message, err := json.Marshal(Message{Status: "500 Internal Server Error", Message: "Database issue. Check Server logs for more details"})
	HandleError(err, "JSON Marshal Error")
	//http.Error(w, string(message), http.StatusInternalServerError)
	w.WriteHeader(500)
	w.Write(message)
	logger("finished handling database error")
	return
}

func GetPassengersService() ([]PassengerInfo, error) {
	db, err := sql.Open("postgres", psqlInfo)
	HandleError(err, "Error making a connection to the database")
	defer db.Close()
	err = db.Ping()
	HandleError(err, "Error pinging to the database")
	var passengers []PassengerInfo
	rows, err := db.Query("SELECT * FROM passengers")
	HandleError(err, "There was a error getting passengers Info from the database")
	if err == nil {
		for rows.Next() {
			var pi PassengerInfo
			err = rows.Scan(&pi.Uuid, &pi.Survived, &pi.Pclass, &pi.Name, &pi.Sex, &pi.Age, &pi.SSA, &pi.PCA, &pi.Fare)
			HandleError(err, "There was a error getting passengers Info from the database")
			passengers = append(passengers, pi)
		}
	}
	return passengers, err
}

func GetPassengerService(uuid string) (PassengerInfo, error) {
	db, err := sql.Open("postgres", psqlInfo)
	HandleError(err, "Error making a connection to the database")
	defer db.Close()
	err = db.Ping()
	HandleError(err, "Error pinging to the database")
	var pi PassengerInfo
	rows, err := db.Query(fmt.Sprintf("select * from passengers where uuid='%s'", uuid))
	HandleError(err, "There was a error getting passengers Info from the database")
	if err == nil {
		for rows.Next() {
			err = rows.Scan(&pi.Uuid, &pi.Survived, &pi.Pclass, &pi.Name, &pi.Sex, &pi.Age, &pi.SSA, &pi.PCA, &pi.Fare)
			HandleError(err, "There was a error getting passengers Info from the database")
		}
	}
	return pi, err
}

func logger(message string) {
	logger := log.New(os.Stdout, "[INFO]: ", log.LstdFlags)
	logger.Println(message)
}

func HandleError(err error, errorMessage string) {
	if err != nil {
		logger(errorMessage)
		logger(err.Error())
	}
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
