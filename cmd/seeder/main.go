package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var products []product

type product struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Descrition string `json:"description"`
}

func main() {

	var dbProducts string
	var jsonProducts string

	flag.StringVar(&dbProducts, "db", "products.db", "Locate: zip.db")
	flag.StringVar(&jsonProducts, "json", "products.json", "read: products.json")
	flag.Parse()

	db, err := sql.Open("sqlite3", dbProducts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	

	//err = seed(db)

	prod, err := os.ReadFile(jsonProducts)
	if err := seed(db); err != nil {
		log.Fatal(err)
	}

	var payload Products //change payload to products
	err = json.Unmarshal(prod, &payload)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert INTO products (id, name, price) values(?, ?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, p := range payload.Products {
		_, err = stmt.Exec(p.Id, p.Name, p.Price)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var price int
		if err := rows.Scan(&id, &name, &price); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, price)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}

func readProducts() {
	bs, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bs, &products)
	if err != nil {
		log.Fatal(err)
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
		if p.Id == id {
			return &p
		}
	}
	return nil
}

func deleteProductsHandler(w http.ResponseWriter, r *http.Request) {
	os.Remove("products.json")
	w.WriteHeader(http.StatusOK)

}

func addProductHandler(w http.ResponseWriter, r *http.Request) {
	var p product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	products = append(products, p)
	w.WriteHeader(http.StatusCreated)
}
