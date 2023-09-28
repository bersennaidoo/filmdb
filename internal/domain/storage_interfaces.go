package domain

import "github.com/bersennaidoo/filmdb/internal/domain/models"

type Storer interface {
	Insert(movie *models.Movie) error
	Get(id int64) (*models.Movie, error)
	Update(movie *models.Movie) error
	Delete(id int64) error
}
