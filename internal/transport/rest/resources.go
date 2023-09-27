package rest

import "github.com/bersennaidoo/lib/pkg/middleware"

type ProduceRequest struct {
	Title   string             `json:"title"`
	Year    int32              `json:"year"`
	Runtime middleware.Runtime `json:"runtime"`
	Genres  []string           `json:"genres"`
}
