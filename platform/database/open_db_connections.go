package database

import (
	"github.com/kaleemubarok/fiber/app/queries"
)

type Queries struct {
	*queries.BookQueries
}

func OpenDBConnection() (*Queries, error) {
	db, err := PostgresSQLConnetion()
	if err != nil {
		return nil, err
	}

	return &Queries{
		BookQueries: &queries.BookQueries{
			DB: db,
		},
	}, nil
}
