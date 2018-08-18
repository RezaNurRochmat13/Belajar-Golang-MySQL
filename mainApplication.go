package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Book struct
type Book struct {
	idBook          string
	bookName        string
	bookWriter      string
	bookPublisher   string
	bookDescription string
	idCategoryBook  string
}

// dbConnection function does config connection to database
func dbConnection() (db *sql.DB) {
	// Configure MySQL Connetion
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bookstore")

	if err != nil {
		panic(err.Error())
	}

	return db
}

// Main function
func main() {
	fmt.Println("test")
	getAllBooks() // Method GET All Books From DB
}

// getAllBooks does function get all data from book table
func getAllBooks() {
	db := dbConnection()
	results, err := db.Query("SELECT * FROM bookstore.book;")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var book Book
		err = results.Scan(&book.idBook, &book.bookName, &book.bookWriter,
			&book.bookPublisher, &book.bookDescription, &book.idCategoryBook)

		if err != nil {
			panic(err.Error())
		}

		log.Println("Books name : ", book.bookName)
		log.Println("Books writer : ", book.bookWriter)
		log.Println("Books publisher : ", book.bookPublisher)
		log.Println("Books description : ", book.bookDescription)
	}
}
