package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kaleemubarok/fiber/app/models"
)

type BookQueries struct {
	*sqlx.DB
}

type IBookQueries interface {
	GetBooks() ([]models.Book, error)
	GetBook(id uuid.UUID) (models.Book, error)
	CreateBook(b *models.Book) error
	UpdateBook(id uuid.UUID, b *models.Book) error
	DeleteBook(id uuid.UUID) error
}

func (q *BookQueries) GetBooks() ([]models.Book, error) {
	books := []models.Book{}

	query := `SELECT * FROM books`

	err := q.Select(&books, query)
	if err != nil {
		return books, err
	}

	return books, nil
}

func (q *BookQueries) GetBook(id uuid.UUID) (models.Book, error) {
	book := models.Book{}

	query := `SELECT * FROM books WHERE id = $1`

	err := q.Get(&book, query, id)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (q *BookQueries) CreateBook(b *models.Book) error {
	query := `INSERT INTO books 
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdateAt, b.Title, b.Author, b.BookStatus, b.BookAttrs)
	if err != nil {
		return err
	}

	return nil
}

func (q *BookQueries) UpdateBook(id uuid.UUID, b *models.Book) error {
	query := `UPDATE books
	SET 
	 	update_at = $2,
		title = $3,
		author = $4,
		book_status = $5,
		book_attrs = $6,
	WHERE
		id = $1`

	_, err := q.Exec(query, id, b.UpdateAt, b.Title, b.Author, b.BookStatus, b.BookAttrs)
	if err != nil {
		return err
	}

	return nil
}

func (q *BookQueries) DeleteBook(id uuid.UUID) error {
	query := `DELETE FROM books WHERE id = $1`

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
