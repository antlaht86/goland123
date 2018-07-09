package main

import (
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"log"
	"fmt"
	"database/sql"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}
func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	rows, err := db.Query("SELECT * FROM books;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	for _, bk := range bks {
		// fmt.Println(bk.isbn, bk.title, bk.author, bk.price)
		fmt.Printf("%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}

	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}

}
