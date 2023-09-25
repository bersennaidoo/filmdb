package rest

import (
	"github.com/bersennaidoo/lib/pkg/infrastructure/config"
	"github.com/bersennaidoo/lib/pkg/middleware"
	"github.com/rs/zerolog"
)

const version = "1.0.0"

type Application struct {
	Config     config.AppConfig
	Logger     zerolog.Logger
	Middleware *middleware.MiddleWare
	Status     int
	Err        error
}
