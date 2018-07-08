package main

import (
	_ "github.com/lib/pq"
	"net/http"
	"io"
	"os"
	"log"
)

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}


func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(port, nil)
	//db, err := sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//
	//if err = db.Ping(); err != nil {
	//	panic(err)
	//}
	//fmt.Println("You connected to your database.")
	//
	//rows, err := db.Query("SELECT * FROM books;")
	//if err != nil {
	//	panic(err)
	//}
	//defer rows.Close()
	//
	//bks := make([]Book, 0)
	//for rows.Next() {
	//	bk := Book{}
	//	err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
	//	if err != nil {
	//		panic(err)
	//	}
	//	bks = append(bks, bk)
	//}
	//if err = rows.Err(); err != nil {
	//	panic(err)
	//}
	//
	//for _, bk := range bks {
	//	// fmt.Println(bk.isbn, bk.title, bk.author, bk.price)
	//	fmt.Printf("%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	//}
}
