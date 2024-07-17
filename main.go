package main

import (
	"fmt"
	"net/http"
	"new_go_example/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/car/get_cars", handlers.GetCars)
	mux.HandleFunc("/api/car/get_one", handlers.GetCar)

	mux.HandleFunc("/api/seller/get_all", handlers.GetAllSellers)
	mux.HandleFunc("/api/seller/get_seller", handlers.GetSeller)
	mux.HandleFunc("/api/seller/become_a", handlers.BecomeASalesman)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}
