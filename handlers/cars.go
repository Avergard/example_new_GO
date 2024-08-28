package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Car struct {
	Car_id              int    `json:"car_id"`
	Mark                string `json:"mark"`
	Kilometerage        int    `json:"kilometerage"`
	Number_of_owners    int    `json:"number_of_owners"`
	Technical_condition string `json:"technical_condition"`
	Traffic_accidents   bool   `json:"traffic_accidents"`
}

var CarBase = map[int]Car{
	1: {
		Car_id:              1,
		Mark:                "Niva",
		Kilometerage:        3750,
		Number_of_owners:    2,
		Technical_condition: "In good condition",
		Traffic_accidents:   false,
	},
	2: {
		Car_id:              2,
		Mark:                "Lada",
		Kilometerage:        78000,
		Number_of_owners:    6,
		Technical_condition: "in poor condition",
		Traffic_accidents:   true,
	},
	3: {
		Car_id:              3,
		Mark:                "Nissan",
		Kilometerage:        29000,
		Number_of_owners:    1,
		Technical_condition: "In good condition",
		Traffic_accidents:   false,
	},
	4: {
		Car_id:              4,
		Mark:                "BMW",
		Kilometerage:        42000,
		Number_of_owners:    2,
		Technical_condition: "in poor condition",
		Traffic_accidents:   false,
	},
	5: {
		Car_id:              5,
		Mark:                "Toyota",
		Kilometerage:        17500,
		Number_of_owners:    2,
		Technical_condition: "In good condition",
		Traffic_accidents:   true,
	},
}

func GetAllCars(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println(r.URL.Query().Get("mark"))
	fmt.Println(r.URL.Query().Get("kilometerage"))
	fmt.Println(r.URL.Query().Get("number_of_owners"))
	fmt.Println(r.URL.Query().Get("technical_conditions"))
	fmt.Println(r.URL.Query().Get("traffic_accidents"))
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func AddCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func ChangeCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
