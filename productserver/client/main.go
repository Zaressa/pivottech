package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://localhost:8080/products")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	bs, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(bs))
}
