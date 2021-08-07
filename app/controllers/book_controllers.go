package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kaleemubarok/fiber/app/models"
	appUtils "github.com/kaleemubarok/fiber/app/utils"
	"github.com/kaleemubarok/fiber/pkg/utils"
	"github.com/kaleemubarok/fiber/platform/database"
)

// GetBooks func get all existis books.
// @Desction Fet all exists books.
// @Summary get all exists books.
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /v1/books [get]
func GetBooks(c *fiber.Ctx) error {
	//create DB connection

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	books, err := db.GetBooks()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "books were not found",
			"count": 0,
			"books": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(books),
		"books": books,
	})
}

// GetBook func get book by given ID or 404 not found.
// @Description Get book by given ID.
// @Summary get get book by given ID.
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {array} models.Book
// @Router /v1/book/{id} [get]
func GetBook(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	book, err := db.GetBook(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "book with given ID is not found",
			"book":  nil,
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"book":  book,
	})
}

// CreeateBook func for cretae new book.
// @Description create a new book.
// @Summary create a new book
// @Tags Book
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param book_attrs model.BookAttrs true "Book attributes"
// @Success 200 {object} models.book
// @Security ApiKeyAuth
// @Router /v1/book [post]
func CreateBook(c *fiber.Ctx) error {
	// get claims for JWT
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// get expires time from JWT
	expires := claims.Expires

	now := time.Now().Unix()
	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	book := &models.Book{}

	// validate if JSON is valid
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := appUtils.NewValidator()

	book.ID = uuid.New()
	book.CreatedAt = time.Now()
	book.BookStatus = 1 // 0 == draft, 1 == active

	// validate book field
	if err := validate.Struct(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// create a new book
	if err := db.CreateBook(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"book":  book,
	})

}
