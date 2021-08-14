package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID            uuid.UUID    `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt     time.Time    `db:"created_at" json:"created_at"`
	UpdateAt      sql.NullTime `db:"updated_at" json:"updated_at"`
	Email         string       `db:"email" json:"email" validate:"required,lte=255"`
	Salt          string       `db:"salt" json:"salt" validate:"required,lte=255"`
	Password      string       `db:"password" json:"password" validate:"required,lte=255"`
	AccountStatus int          `db:"account_status" json:"account_status" validate:"required,len=1"`
}

type AccountSwag struct {
	ID            uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdateAt      time.Time `db:"updated_at" json:"updated_at"`
	Email         string    `db:"email" json:"email" validate:"required,lte=255"`
	Salt          string    `db:"salt" json:"salt" validate:"required,lte=255"`
	Password      string    `db:"password" json:"password" validate:"required,lte=255"`
	AccountStatus int       `db:"account_status" json:"account_status" validate:"required,len=1"`
}
