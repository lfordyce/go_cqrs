package db

import (
	"context"
	"database/sql"
	"github.com/lfordyce/hero_cqrs/schema"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err !=  nil {
		return nil, err
	}

	return &PostgresRepository{
		db,
	}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()
}

func (r *PostgresRepository) InsertHero(ctx context.Context, hero schema.Hero) error {
	_, err := r.db.Exec("INSERT INTO heros(id, body, created_at) VALUES($1, $2, $3)", hero.ID, hero.Body, hero.CreatedAt)
	return err
}

func (r *PostgresRepository) ListHero(ctx context.Context, skip uint64, take uint64) ([]schema.Hero, error) {

	//rows, err := r.db.Query("SELECT * FROM meows ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take)

	rows, err := r.db.Query("SELECT * FROM heros ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Pares all rows into an array of Heros
	// var heros []schema.Hero
	heros := []schema.Hero{}
	for rows.Next() {
		hero := schema.Hero{}
		if err = rows.Scan(&hero.ID, &hero.Body, &hero.CreatedAt); err == nil {
			heros = append(heros, hero)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return heros, nil
}