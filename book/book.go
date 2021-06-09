package book

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mthaler/book-managment-system/database"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}


func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	fmt.Println(book)
	db.Create(&book)
	c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	update := new(Book)
	if err := c.BodyParser(update); err != nil {
		c.Status(503).Send(err)
		return
	}
	book.Title = update.Title
	book.Author = update.Author
	book.Rating = update.Rating
	db.Save(book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	if book.Title == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	}
	db.First(&book, id)
	db.Delete(&book)
	c.Send("Book Successfully deleted")
}