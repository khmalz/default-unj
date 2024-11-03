package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age" validate:"gte=0"`
	Address  string    `json:"address"`
	Country  string    `json:"country"`
	Alive    bool      `json:"is_alive"`
	ProductS []Product `json:"products"`
}

type Product struct {
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Stock   int    `json:"stock"`
	Expired bool   `json:"is_expired"`
}

var user User
var product []Product

func GetUser(w http.ResponseWriter, r *http.Request) {
	userData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(userData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productData, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(product)

	_, err = w.Write(productData)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveUser(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&user)

	product = user.ProductS

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
	fmt.Println(product)
	w.WriteHeader(http.StatusOK)

	userData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(userData)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveProduct(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("POST /user", SaveUser)
	http.HandleFunc("GET /product", GetProduct)
	http.HandleFunc("POST /product", SaveProduct)

	fmt.Println("Server started on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
