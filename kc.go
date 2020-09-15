package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	tmpDB, err := sql.Open("postgres", "dbname=books_database user=postgres password=postgres host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB
    	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}
	e.Logger.Fatal(e.Start(":1323"))
}
func getBook(bookID int) (Book, error) {
	//Retrieve
	res := Book{}

	var id int
	var name string
	var author string
	var price int
	
	err := db.QueryRow(`SELECT id, name, author, price FROM books where id = $1`, bookID).Scan(&id, &name, &author, &price)
	if err == nil {
		res = Book{ID: id, Name: name, Author: author, Price: price}
	}

	return res, err
}

func insertBook(name, author string,price int ) (int, error) {
	//Create
	var bookID int
	err := db.QueryRow(`INSERT INTO books(name, author, price) VALUES($1, $2, $3) RETURNING id`, name, author, pages, publicationDate).Scan(&bookID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Last inserted ID: %v\n", Books)
	return bookID, err
}

func updateBook(id int, name, author string, pages int, publicationDate time.Time) (int, error) {
	//Create
	res, err := db.Exec(`UPDATE books set name=$1, author=$2, price=$3 RETURNING id`, name, author, pages, publicationDate, id)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}

func removeBook(bookID int) (int, error) {
	//Delete
	res, err := db.Exec(`delete from books where id = $1`, bookID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}
