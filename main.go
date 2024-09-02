package main

import (
	"fmt"
	"net/http"
	"new_go_example/handlers"
)

func main() {

	handlers.InitDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/car/get_all_cars", handlers.GetAllCars)
	mux.HandleFunc("/api/car/get_one", handlers.GetCar)
	mux.HandleFunc("/api/car/delete_car", handlers.DeleteCar)
	mux.HandleFunc("/api/car/add_car", handlers.AddCar)

	mux.HandleFunc("/api/seller/get_all_sellers", handlers.GetAllSellers)
	mux.HandleFunc("/api/seller/get_seller", handlers.GetSeller)
	mux.HandleFunc("/api/seller/delete_seller", handlers.DeleteSalesman)
	mux.HandleFunc("/api/seller/add_seller", handlers.AddSeller)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
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

*/
