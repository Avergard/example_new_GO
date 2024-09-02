package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

type Car struct {
	Id                  int    `json:"id"`
	Mark                string `json:"mark"`
	Technical_condition string `json:"technical_condition"`
	Kilometerage        int    `json:"kilometerage"`
	Number_of_owners    int    `json:"number_of_owners"`
	Traffic_accidents   bool   `json:"traffic_accidents"`
}

func GetAllCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, mark, technical_condition,kilometerage,number_of_owners,traffic_accidents FROM cars")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var cars []Car
	for rows.Next() {
		var car Car
		if err := rows.Scan(&car.Id, &car.Mark, &car.Technical_condition, &car.Kilometerage, &car.Number_of_owners, &car.Traffic_accidents); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cars = append(cars, car)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytesBody, err := json.Marshal(cars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(bytesBody)
	if err != nil {
		fmt.Println(err)
	}
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idOfCar, err := strconv.Atoi(r.URL.Query().Get("id_of_car"))
	if err != nil {
		http.Error(w, "Неправильный id машины", http.StatusInternalServerError)
		return
	}

	var car Car
	err = db.QueryRow("SELECT id, mark, technical_condition, kilometerage, number_of_owners, traffic_accidents FROM cars WHERE id = $1", idOfCar).Scan(
		&car.Id, &car.Mark, &car.Technical_condition, &car.Kilometerage, &car.Number_of_owners, &car.Traffic_accidents)
	if err == sql.ErrNoRows {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytesBody, err := json.Marshal(car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(bytesBody)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idOfCar, err := strconv.Atoi(r.URL.Query().Get("id_of_car"))
	if err != nil {
		http.Error(w, "invalid car id", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM cars WHERE id = $1", idOfCar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func AddCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	car.Id = 0

	err := db.QueryRow(
		"INSERT INTO cars(mark,technical_condition, kilometerage, number_of_owners, traffic_accidents) VALUES ($1,$2,$3,$4,$5) RETURNING id", car.Mark, car.Technical_condition, car.Kilometerage, car.Number_of_owners, car.Traffic_accidents).Scan(&car.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	bytesBody, err := json.Marshal(car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(bytesBody)
	if err != nil {
		fmt.Println(err)
	}
}
