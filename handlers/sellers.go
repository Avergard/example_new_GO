package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Avergard/example_new_GO/helpers"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

type Seller struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Age        int    `json:"age"`
	Experience int    `json:"experience"`
	Sales      int    `json:"sales"`
}

// получение всех продавцов
func GetAllSellers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := helpers.DB.Query("SELECT id, name, surname, age, experience, sales FROM sellers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sellers []Seller
	for rows.Next() {
		var seller Seller
		if err := rows.Scan(&seller.ID, &seller.Name, &seller.Surname, &seller.Age, &seller.Experience, &seller.Sales); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sellers = append(sellers, seller)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytesBody, err := json.Marshal(sellers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(bytesBody)
	if err != nil {
		fmt.Println(err)
	}
}

// продавец по id
func GetSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idOfSeller, err := strconv.Atoi(r.URL.Query().Get("id_of_seller"))
	if err != nil {
		http.Error(w, "неверный id", http.StatusBadRequest)
		return
	}

	var seller Seller
	err = helpers.DB.QueryRow("SELECT id, name, surname, age, experience, sales FROM sellers WHERE id = $1", idOfSeller).Scan(
		&seller.ID, &seller.Name, &seller.Surname, &seller.Age, &seller.Experience, &seller.Sales)
	if err == sql.ErrNoRows {
		http.Error(w, "продавец не найден", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytesBody, err := json.Marshal(seller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(bytesBody)
	if err != nil {
		fmt.Println(err)
	}
}

// удаление по id
func DeleteSalesman(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idOfSeller, err := strconv.Atoi(r.URL.Query().Get("id_of_seller"))
	if err != nil {
		http.Error(w, "неверный id", http.StatusBadRequest)
		return
	}

	_, err = helpers.DB.Exec("DELETE FROM sellers WHERE id = $1", idOfSeller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// добавить продавца
func AddSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var seller Seller
	if err := json.NewDecoder(r.Body).Decode(&seller); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := helpers.DB.QueryRow(
		"INSERT INTO sellers (name, surname, age, experience, sales) VALUES ($1, $2, $3, $4, $5) RETURNING iD", seller.Name, seller.Surname, seller.Age, seller.Experience, seller.Sales).Scan(&seller.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	bytesBody, err := json.Marshal(seller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(bytesBody)
	if err != nil {
		fmt.Println(err)
	}
}

/*http error заменить на log.error
инсерт заменить, аппенд заменить на общее получение инфы
ID, не id
*/
