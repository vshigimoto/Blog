package repository

import "github.com/jmoiron/sqlx"

type Repo struct {
	main *sqlx.DB
	repo *sqlx.DB
}

func New(repo, main *sqlx.DB) *Repo {
	return &Repo{main: main, repo: repo}
}
