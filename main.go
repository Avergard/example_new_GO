package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Avergard/example_new_GO/handlers"
	"github.com/Avergard/example_new_GO/helpers"
	_ "github.com/lib/pq"
)

type SiteInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Author      string `json:"author"`
}

func infoSiteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	siteInfo := SiteInfo{
		Name:        "Сайт для задания",
		Description: "Этот сайт создан для демонстрации.",
		Version:     "1.0.0",
		Author:      "Александ Азизов",
	}

	if err := json.NewEncoder(w).Encode(siteInfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func infoPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	siteInfo := SiteInfo{
		Name:        "Сайт для задания",
		Description: "Этот сайт создан для демонстрации.",
		Version:     "1.0.0",
		Author:      "Александ Азизов",
	}

	json.NewEncoder(w).Encode(siteInfo)
}

func setupCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	helpers.InitDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/car/get_all", handlers.GetAllCars)
	mux.HandleFunc("/api/car/get_one", handlers.GetCar)
	mux.HandleFunc("/api/car/delete", handlers.DeleteCar)
	mux.HandleFunc("/api/car/add", handlers.AddCar)

	mux.HandleFunc("/api/seller/get_all", handlers.GetAllSellers)
	mux.HandleFunc("/api/seller/get", handlers.GetSeller)
	mux.HandleFunc("/api/seller/delete", handlers.DeleteSalesman)
	mux.HandleFunc("/api/seller/add", handlers.AddSeller)

	mux.HandleFunc("/api/info-site", infoSiteHandler)

	mux.HandleFunc("/info", infoPageHandler)

	err := http.ListenAndServe(":8080", setupCORS(mux))
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

/* команды для постмана
1)удалить позицию по id  http://localhost:8080/api/seller/delete_seller?id_of_seller=5
2)получить все позиции в базе данных  http://localhost:8080/api/seller/get_all
3)получить одну позицию по id http://localhost:8080/api/seller/get_seller?id_of_seller=2
4)добавить позицию http://localhost:8080/api/seller/add_seller, Body - raw - json, дальше записываешь в {} характеристики позиции
пример:
{
  "name": "Иван",
  "surname": "Иванов",
  "age": 30,
  "experience": 5,
  "sales": 100
}
module github.com/Avergard/backend
require github.com/lib/pq v1.10.9
*/
