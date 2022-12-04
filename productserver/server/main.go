package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// this can be used across any func
var products []product

type product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p := productlookup(id)
	if p == nil {
		log.Printf("Product not found: %d\n", id)
		return
	}
	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Printf("Error encoding product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func productlookup(id int) *product {
	for _, p := range products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

func initProducts() {
	bs, err := os.ReadFile("cmd/seeder/products.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bs, &products)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	initProducts()
	r := mux.NewRouter()
	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products/{id}", getProductsHandler).Methods("GET")

	r.HandleFunc("/products", getProductsHandler)

	// Bind to a port and pass our router in
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
