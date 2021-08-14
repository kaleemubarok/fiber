package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kaleemubarok/fiber/app/models"
)

type IAccountQueries interface {
	GetAccounts() ([]models.Account, error)
	GetAccount(id uuid.UUID) (models.Account, error)
	GetAccountByEmail(email string) (models.Account, error)
	CreateAccount(account models.Account) error
}

type AccountQueries struct {
	*sqlx.DB
}

func (db AccountQueries) GetAccounts() ([]models.Account, error) {
	account := []models.Account{}

	query := `SELECT * FROM accounts`

	err := db.Select(&account, query)
	if err != nil {
		return account, err
	}

	return account, nil
}

func (db AccountQueries) GetAccount(id uuid.UUID) (models.Account, error) {
	account := models.Account{}

	query := `SELECT * FROM accounts WHERE id=$1`
	err := db.Get(&account, query, id)
	if err != nil {
		return account, err
	}

	return account, nil
}

func (db AccountQueries) CreateAccount(account models.Account) error {
	query := `INSERT INTO accounts (id, email, salt, password, account_status, created_at)
	VALUES ( $1, $2, $3, $4, $5, $6 )`

	_, err := db.Exec(query, account.ID, account.Email, account.Salt, account.Password, account.AccountStatus, account.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (db AccountQueries) GetAccountByEmail(email string) (models.Account, error) {
	account := models.Account{}

	query := `SELECT * FROM accounts WHERE email=$1`
	err := db.Get(&account, query, email)
	if err != nil {
		return account, err
	}

	return account, nil
}

func NewAccountQueries(db sqlx.DB) *AccountQueries {
	return &AccountQueries{&db}
}
