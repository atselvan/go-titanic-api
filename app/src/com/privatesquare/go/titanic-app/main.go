package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

const (
	host     = "titanic-db"
	port     = 5432
	user     = "postgres"
	password = "BXJSKzj8cLScg7Zg"
	dbname   = "titanic"
)

var (
	psqlInfo string
)

type PassengerInfo struct {
	Survived bool `json:"survived"`
	Pclass  string `json:"pclass"`
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  float64 `json:"Age"`
	SSA  int `json:"ssa"`
	PCA  int `json:"pca"`
	fare float64 `json:"fare"`
}

func main() {
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	router := mux.NewRouter()
	router.HandleFunc("/passengers", GetPassengersHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPassengersHandler(w http.ResponseWriter, r *http.Request){
	passengers := GetPassengersService()
	json.NewEncoder(w).Encode(passengers)
}


func GetPassengersService() []PassengerInfo{
	db, err := sql.Open("postgres", psqlInfo)
	HandleError(err, "Error making a connection to the database")
	defer db.Close()
	err = db.Ping()
	HandleError(err, "Error pinging to the database")
	var passengers []PassengerInfo
	rows, err := db.Query("SELECT * FROM passengers")
	HandleError(err, "There was a error getting passengers Info from the database")

	for rows.Next() {
		var pi PassengerInfo
		err = rows.Scan(&pi.Survived,&pi.Pclass,&pi.Name,&pi.Sex,&pi.Age,&pi.SSA,&pi.PCA,&pi.fare)
		HandleError(err, "There was a error getting passengers Info from the databas")
		passengers = append(passengers, pi)
	}
	return passengers
}

func HandleError(err error, errorMessage string) {
	if err != nil {
		log.Println(errorMessage)
		log.Fatal(err)
	}
}