package user

import "github.com/jmoiron/sqlx"

const userTable = "users"

type Repo struct {
	DB *sqlx.DB
}

func NewRepo(DB *sqlx.DB) *Repo {
	return &Repo{
		DB: DB,
	}
}
