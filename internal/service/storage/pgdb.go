package storage

import (
	"database/sql"

	"github.com/bersennaidoo/filmdb/internal/domain/models"
	"github.com/lib/pq"
)

type PGStore struct {
	DB *sql.DB
}

func NewPGStore(db *sql.DB) *PGStore {

	return &PGStore{
		DB: db,
	}
}

func (m *PGStore) Insert(movie *models.Movie) error {

	query := `
             INSERT INTO movies (title, year, runtime, genres)
			 VALUES ($1, $2, $3, $4)
			 RETURNING id, created_at, version`

	args := []interface{}{movie.Title, movie.Year,
		movie.Runtime, pq.Array(movie.Genres)}

	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt,
		&movie.Version)
}

func (m *PGStore) Get(id int64) (*models.Movie, error) {

	return nil, nil
}

func (m *PGStore) Update(movie *models.Movie) error {

	return nil
}

func (m *PGStore) Delete(id int64) error {

	return nil
}
