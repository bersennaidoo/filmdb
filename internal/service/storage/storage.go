package storage

import (
	"errors"

	"github.com/bersennaidoo/filmdb/internal/domain"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Storage struct {
	PGStore domain.Storer
}

func New(pgstore domain.Storer) *Storage {

	return &Storage{
		PGStore: pgstore,
	}
}
