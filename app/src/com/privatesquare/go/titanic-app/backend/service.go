package backend

import (
	m "com/privatesquare/go/titanic-app/model"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

const (
	port = 5432
)

var (
	host     = os.Getenv("DB_HOST")
	dbName   = os.Getenv("DB_NAME")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	psqlInfo = fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", host, port, dbName, user, password)
)

func dbConnect() *sql.DB {
	//uncomment below connection details for testing purpose
	//psqlInfo = "host=localhost port=5432 dbname=titanic user=postgres password=password sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	HandleError(err, "Error making a connection to the database")
	err = db.Ping()
	HandleError(err, "Error pinging to the database")
	return db
}

func dbClose(db *sql.DB) {
	db.Close()
}

func GetPassengersService() ([]m.PassengerInfo, error) {
	db := dbConnect()
	var passengers []m.PassengerInfo
	sqlStatement := "SELECT * FROM passengers"
	rows, err := db.Query(sqlStatement)
	HandleError(err, "There was a error getting passengers Info from the database")
	if err == nil {
		for rows.Next() {
			var pi m.PassengerInfo
			err = rows.Scan(&pi.Uuid, &pi.Survived, &pi.Pclass, &pi.Name, &pi.Sex, &pi.Age, &pi.SSA, &pi.PCA, &pi.Fare)
			HandleError(err, "There was a error getting passengers Info from the database")
			passengers = append(passengers, pi)
		}
	}
	dbClose(db)
	return passengers, err
}

func GetPassengerService(uuid string) (m.PassengerInfo, error) {
	db := dbConnect()
	var pi m.PassengerInfo
	sqlStatement := fmt.Sprintf("select * from passengers where uuid='%s'", uuid)
	rows, err := db.Query(sqlStatement)
	HandleError(err, "There was a error getting passengers Info from the database")
	if err == nil {
		for rows.Next() {
			err = rows.Scan(&pi.Uuid, &pi.Survived, &pi.Pclass, &pi.Name, &pi.Sex, &pi.Age, &pi.SSA, &pi.PCA, &pi.Fare)
			HandleError(err, "There was a error getting passengers Info from the database")
		}
	}
	dbClose(db)
	return pi, err
}

func AddPassengerService(pi m.PassengerInfo) string {
	db := dbConnect()
	sqlStatement := fmt.Sprintf("INSERT INTO passengers (survived, pclass, name, sex, age, ssa, pca, fare) VALUES (%v, %s, '%s', '%s', %f, %d, %d, %f) RETURNING uuid", pi.Survived, pi.Pclass, pi.Name, pi.Sex, pi.Age, pi.SSA, pi.PCA, pi.Fare)
	var uuid string
	err := db.QueryRow(sqlStatement).Scan(&uuid)
	HandleError(err, "There was a error adding passenger's Info to the database")
	logger(fmt.Sprintf("Passenger info is added with uuid %s", uuid))
	dbClose(db)
	return uuid
}

// TODO
func UpdatePassengerService() {

}

//TODO
func DeletePassengerService() {

}
