package main

import (
	"fmt"
	"net/http"
	"new_go_example/handlers"
)

func main() {

	handlers.InitDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/car/get_cars", handlers.GetAllCars)
	mux.HandleFunc("/api/car/get_one", handlers.GetCar)
	mux.HandleFunc("/api/car/delete", handlers.DeleteCar)
	mux.HandleFunc("/api/car/change", handlers.ChangeCar)
	mux.HandleFunc("/api/car/add", handlers.AddCar)

	mux.HandleFunc("/api/seller/get_all", handlers.GetAllSellers)
	mux.HandleFunc("/api/seller/get_seller", handlers.GetSeller)
	mux.HandleFunc("/api/seller/delete_seller", handlers.DeleteSalesman)
	mux.HandleFunc("/api/seller/add_seller", handlers.AddSeller)
	mux.HandleFunc("/api/seller/become_a", handlers.BecomeASalesman)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}
