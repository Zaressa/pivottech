package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var products []product

type product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	readProducts()

	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	db, err := sql.Open("sqlite3", "./zip.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// sqlStmt := `
	// create table products (id integer primary key, name text, price integer);
	// delete from products;
	// `
	// _, err = db.Exec(sqlStmt)
	// if err != nil {
	// 	log.Printf("%q: %s\n", err, sqlStmt)
	// 	return
	// }

	err = seed(db)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p product
		err = rows.Scan(&p.Id, &p.Name, &p.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(products)
}

func seed(db *sql.DB) error { // Path: cmd/Seeder/main.go

	// products := []product{
	// 	{Id: 1, Name: "Milk", Price: 2.99},
	// 	{Id: 2, Name: "Eggs", Price: 3.99},
	// 	{Id: 3, Name: "Bread", Price: 2.49},
	// }

	tx, err := db.Begin()
	if err != nil {
		return err

	}
	stmt, err := tx.Prepare("insert INTO products (id, name, price) values(?, ?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i := 0; i < 5; i++ {
		_, err = stmt.Exec(fmt.Sprintf("id-%d", i), fmt.Sprintf("name-%d", i), fmt.Sprintf("price-%d", i))
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var price int
		err = rows.Scan(&id, &name, &price)
		if err != nil {
			return err
		}
		fmt.Println(id, name, price)
	}
	err = rows.Err()
	if err != nil {
		return err
	}

	stmt, err = db.Prepare("insert INTO products (id, name, price) values(?, ?,?)")

	if err != nil {
		return err
	}
	defer stmt.Close()
	var id int
	var name string
	var price int
	err = stmt.QueryRow(id, name, price).Scan(&id, &name, &price)
	if err != nil {
		return err
	}
	fmt.Println(id, name, price)

	_, err = db.Exec("delete from products")
	if err != nil {
		return err
	}

	_, err = db.Exec("insert into products (id, name, price) values (?, ?, ?)", 1, "Milk", 2.99)
	if err != nil {
		return err
	}

	rows, err = db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var price int
		err = rows.Scan(&id, &name, &price)
		if err != nil {
			return err
		}
		fmt.Println(id, name, price)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
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
	w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusNotFound)
	w.WriteHeader(http.StatusCreated)
	w.WriteHeader(http.StatusAccepted)
	w.WriteHeader(http.StatusNoContent)
	w.WriteHeader(http.StatusUnauthorized)
w.Header().Set("Content-Type", "application/json")
resp := make(map[string]string)
resp["message"] = "Products not found"
_, err := json.Marshal(resp)
if err != nil {
	log.Fatal(err)
	if err := json.NewEncoder(w).Encode(products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	return

}
}

func deleteProductsHandler(w http.ResponseWriter, r *http.Request) {
	os.Remove ("products.json")
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