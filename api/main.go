package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age" validate:"gte=0"`
	Address string `json:"address"`
	Country string `json:"country"`
	Alive   bool   `json:"is_alive"`
}

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
	Expired bool `json:"is_expired"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:    "John",
		Age:     20,
		Address: "123 Main St",
		Country: "USA",
		Alive:   true,
	}
	userData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(userData)
	if err != nil {
		log.Fatal(err)
	}
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	product := Product{
		Name:  "iPhone",
		Price: 1000,
		Stock: 10,
		Expired: false,
	}
	productData, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(productData)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(product)
	w.WriteHeader(http.StatusOK)
	
	productData, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(productData)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("GET /user", GetUser)
	http.HandleFunc("GET /product", GetProduct)
	http.HandleFunc("POST /product", SaveProduct)

	fmt.Println("Server started on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
