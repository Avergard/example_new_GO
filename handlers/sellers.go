package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Seller struct {
	Name       string `json:"Name"`
	Surname    string `json:"Surname"`
	Age        int    `json:"Age"`
	Experience int    `json:"Experience"`
	Sales      int    `json:"Sales"`
}

var sellerBase = map[int]Seller{
	1: {
		Name:       "Nikita",
		Surname:    "Telitsyn",
		Age:        35,
		Experience: 11,
		Sales:      21,
	},
	2: {
		Name:       "Sasha",
		Surname:    "Azizov",
		Age:        46,
		Experience: 22,
		Sales:      48,
	},
	3: {
		Name:       "Denis",
		Surname:    "Lisenkov",
		Age:        56,
		Experience: 30,
		Sales:      55,
	},
	4: {
		Name:       "Kirill",
		Surname:    "Bilibeev",
		Age:        61,
		Experience: 35,
		Sales:      62,
	},
	5: {
		Name:       "Fedya",
		Surname:    "Lazarev",
		Age:        28,
		Experience: 4,
		Sales:      7,
	},
}

func GetAllSellers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bytesBody, err := json.Marshal(sellerBase)
	if err != nil {
		fmt.Println(err)
		SendError(w, "LOL", "server out", "server out", http.StatusBadRequest)
	}

	_, err = w.Write(bytesBody)
	if err != nil {
		fmt.Println(err)
	}
}

func GetSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idOfSeller, err := strconv.Atoi(r.URL.Query().Get("id_of_seller"))
	if err != nil {
		fmt.Println(err)
	}

	seller, ok := sellerBase[idOfSeller]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		SendError(w, "Sorry about this", "Not found", "Seller out", http.StatusNotFound)
		return
	}

	bytesBody, err := json.Marshal(seller)
	if err != nil {
		fmt.Println(err)
	}

	_, err = w.Write(bytesBody)
	if err != nil {
		fmt.Println(err)
	}
}

func BecomeASalesman(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println(r.Body)
}
