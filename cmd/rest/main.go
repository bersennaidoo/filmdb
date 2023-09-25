package main

import (
	"github.com/bersennaidoo/filmdb/internal/transport/rest"
	"github.com/bersennaidoo/lib/pkg/infrastructure/config"
	"github.com/bersennaidoo/lib/pkg/infrastructure/logger"
	"github.com/bersennaidoo/lib/pkg/middleware"
)

func main() {
	cfg, _ := config.InitConfig()

	mid := middleware.New()

	trace := logger.TRACE

	zcfg := logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    trace,
	}

	zlog := logger.NewZeroLoggerSrv(zcfg)

	app := rest.Application{
		Config:     cfg,
		Logger:     zlog,
		Middleware: mid,
	}

	app.ServerRun()
}
