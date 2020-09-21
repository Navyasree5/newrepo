package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type books struct {
	BookName string `json:"BookName,omitempty"`
	Author   string `json:"Author,omitempty"`
	Pages    int    `json:"Pages,omitempty"`
	Price    int    `json:"Price,omitempty"`
}

var book []books

func getbook(c echo.Context) error {
	Author, _ := strconv.Atoi(c.Param("Author"))
	return c.JSON(http.StatusOK, book[Author])
}
func insertbook(c echo.Context) error {
	books := books{}
	err := c.Bind(&books)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	book = append(book, books)
	return c.JSON(http.StatusCreated, book)

}
func updatebook(c echo.Context) error {
	books := new(books)
	err := c.Bind(books)
	if err != nil {
		return err
	}
	Author, _ := strconv.Atoi(c.Param("Author"))
	book[Author].BookName = books.BookName
	return c.JSON(http.StatusOK, book)
}
func deletebook(c echo.Context) error {
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/books", getbook)
	e.POST("/books", insertbook)
	e.PUT("/books", updatebook)
	e.DELETE("/books", deletebook)
	e.Logger.Fatal(e.Start(":9002"))
	return
}
