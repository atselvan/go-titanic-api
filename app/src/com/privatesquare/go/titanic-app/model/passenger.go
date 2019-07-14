package model

// PassengerInfo represents the information of a passenger
//
// swagger:parameters addPassenger
type PassengerInfo struct {
	// uuid
	Uuid string `json:"uuid"`
	// survived
	Survived bool `json:"survived"`
	// passenger class: 1: Luxury, 2: Economy
	Pclass string `json:"pclass"`
	// name
	Name string `json:"name"`
	// sex
	Sex string `json:"sex"`
	// age
	Age float64 `json:"Age"`
	// ssa
	SSA int `json:"ssa"`
	// pca
	PCA int `json:"pca"`
	// fare
	Fare float64 `json:"fare"`
}
