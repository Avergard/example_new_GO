package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Car struct {
	Mark               string `json:"Mark"`
	Kilometerage       int    `json:"Kilometerage"`
	NumberOfOwners     int    `json:"Number_of_owners"`
	TechnicalСondition string `json:"Technical_condition"`
	TrafficAccidents   bool   `json:"Traffic_accidents"`
}

var CarBase = map[int]Car{
	1: {
		Mark:               "Niva",
		Kilometerage:       3750,
		NumberOfOwners:     2,
		TechnicalСondition: "In good condition",
		TrafficAccidents:   false,
	},
	2: {
		Mark:               "Lada",
		Kilometerage:       78000,
		NumberOfOwners:     6,
		TechnicalСondition: "in poor condition",
		TrafficAccidents:   true,
	},
	3: {
		Mark:               "Nissan",
		Kilometerage:       29000,
		NumberOfOwners:     1,
		TechnicalСondition: "In good condition",
		TrafficAccidents:   false,
	},
	4: {
		Mark:               "BMW",
		Kilometerage:       42000,
		NumberOfOwners:     2,
		TechnicalСondition: "in poor condition",
		TrafficAccidents:   false,
	},
	5: {
		Mark:               "Toyota",
		Kilometerage:       17500,
		NumberOfOwners:     2,
		TechnicalСondition: "In good condition",
		TrafficAccidents:   true,
	},
}

func GetCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bytesBody, err := json.Marshal(CarBase)
	if err != nil {
		fmt.Println(err)
	}
	_, err = w.Write(bytesBody)
	if err != nil {
		fmt.Println(err)
	}
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println(r.URL.Query().Get("Mark"))
	fmt.Println(r.URL.Query().Get("Kilometerage"))
	fmt.Println(r.URL.Query().Get("NumberOfOwners"))
	fmt.Println(r.URL.Query().Get("TechnicalСondition"))
	fmt.Println(r.URL.Query().Get("TrafficAccidents"))
}
