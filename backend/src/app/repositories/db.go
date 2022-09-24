package repositories

import (
	"tmp_project_name/app/infrastructures/postgresql"

	"github.com/jmoiron/sqlx"
)

func createDB() *postgresql.Postgresql {
	return postgresql.NewPostgresql()
}

type db interface {
	WithDbContext(fn func(db *sqlx.DB) error) error
}
