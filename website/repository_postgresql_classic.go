package website

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgconn"
)

type PostgreSQLClassicRepository struct {
	db *sql.DB
}

func NewPostgreSQLClassicRepository(db *sql.DB) *PostgreSQLClassicRepository {
	return &PostgreSQLClassicRepository{
		db: db,
	}

}

func (r *PostgreSQLClassicRepository) Migrate(ctx context.Context) error {
	query := `
		CREATE TABLE IF NOT EXISTS websites (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			url 	TEXT NOT NULL,
			rank 	INT NOT NULL
		);
	`
	_, err := r.db.ExecContext(ctx, query)
	return err
}

func (r *PostgreSQLClassicRepository) Create(ctx context.Context, website Website) (*Website, error) {
	var id int64
	err := r.db.QueryRowContext(ctx, "INSERT INTO websites(name, url, rank) values($1, $2, $3) RETURNING id", website.Name, website.URL, website.Rank).Scan(&id)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}
	website.ID = id
	return &website, nil

}

func (r *PostgreSQLClassicRepository) All(ctx context.Context) ([]Website, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM websites")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var all []Website
	for rows.Next() {
		var website Wet
	}
}
