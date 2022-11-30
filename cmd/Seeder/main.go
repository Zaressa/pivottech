package main
import(
	"database/sql"
	"log"
	_"github.com/mattn/go-sqlite3"
)
func main () {
	db, err := sql.Open("sqlite3", "./zip.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

rows, err := db.Query("SELECT id, name, price FROM products")
if err != nil {
	log.Fatal (err)
}
defer rows.Close()

for rows.Next() {
	var id int
	var name string
	var price float64
	err = rows.Scan(&id, &name, &price)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(id, name, price)
}
err = rows.Err()
if err != nil {
	log.Fatal(err)
}
