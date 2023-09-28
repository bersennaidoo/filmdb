package storage

import (
	"database/sql"
	"errors"

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

	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
             SELECT id, created_at, title, year, runtime, genres, version
			 FROM movies
			 WHERE id = $1`

	var movie models.Movie

	err := m.DB.QueryRow(query, id).Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.Title,
		&movie.Year,
		&movie.Runtime,
		pq.Array(&movie.Genres),
		&movie.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &movie, nil
}

func (m *PGStore) Update(movie *models.Movie) error {

	query := `
             UPDATE movies
			 SET title = $1, year = $2, runtime = $3, genres = $4, version = version + 1
			 WHERE id = $5
			 RETURNING version`

	args := []interface{}{
		movie.Title,
		movie.Year,
		movie.Runtime,
		pq.Array(movie.Genres),
		movie.ID,
	}

	return m.DB.QueryRow(query, args...).Scan(&movie.Version)
}

func (m *PGStore) Delete(id int64) error {

	query := `
             DELETE FROM movies
			 WHERE id = $1`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
