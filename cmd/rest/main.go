package main

import (
	"github.com/bersennaidoo/filmdb/internal/service/storage"
	"github.com/bersennaidoo/filmdb/internal/transport/rest"
	"github.com/bersennaidoo/lib/pkg/infrastructure/config"
	"github.com/bersennaidoo/lib/pkg/infrastructure/connections"
	"github.com/bersennaidoo/lib/pkg/infrastructure/logger"
	"github.com/bersennaidoo/lib/pkg/middleware"
)

func main() {
	cfg, _ := config.InitConfig()

	valid := middleware.NewValidator()

	mid := middleware.New(valid)

	trace := logger.TRACE

	zcfg := logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    trace,
	}

	zlog := logger.NewZeroLoggerSrv(zcfg)

	db, err := connections.OpenPGDB(cfg)
	if err != nil {
		zlog.Error().Err(err).Msg("")
		panic("can't connect to database")
	}

	defer db.Close()

	if err == nil {

		zlog.Info().Msg("database connection pool established")
	}

	pgstore := storage.NewPGStore(db)

	storage := storage.New(pgstore)

	app := rest.Application{
		Config:     cfg,
		Logger:     zlog,
		Middleware: mid,
		Storage:    storage,
	}

	app.ServerRun()
}
