package rest_test

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bersennaidoo/filmdb/internal/service/storage"
	"github.com/bersennaidoo/filmdb/internal/transport/rest"
	"github.com/bersennaidoo/filmdb/mocks"
	"github.com/bersennaidoo/lib/pkg/infrastructure/config"
	"github.com/bersennaidoo/lib/pkg/infrastructure/logger"
	"github.com/bersennaidoo/lib/pkg/middleware"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func TestCreateMovieHandlerUnit(t *testing.T) {

	t.Run("creates a movie", func(t *testing.T) {

		appmock := mocks.NewRester(t)

		handler := http.HandlerFunc(appmock.CreateMovieHandler)

		expected := http.StatusOK

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/v1/movies", nil)
		appmock.On("CreateMovieHandler", rr, req).Return(expected).Once()

		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, expected)
		appmock.AssertExpectations(t)
	})
}

func TestDeleteMovieHandlerUnit(t *testing.T) {

	t.Run("deletes a movie by id", func(t *testing.T) {

		appmock := mocks.NewRester(t)

		handler := http.HandlerFunc(appmock.DeleteMovieHandler)

		expected := http.StatusOK

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/v1/movies/1", nil)
		appmock.On("DeleteMovieHandler", rr, req).Return(expected).Once()

		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, expected)
		appmock.AssertExpectations(t)
	})
}

func TestHealthcheckHandlerUnit(t *testing.T) {

	t.Run("checks health of system and return 200", func(t *testing.T) {

		appmock := mocks.NewRester(t)

		handler := http.HandlerFunc(appmock.HealthcheckHandler)

		expected := http.StatusOK

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/healthcheck", nil)
		appmock.On("HealthcheckHandler", rr, req).Return(expected).Once()

		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, expected)
		appmock.AssertExpectations(t)
	})
}

func TestShowMovieHandlerUnit(t *testing.T) {

	t.Run("shows movie and returns 200", func(t *testing.T) {

		appmock := mocks.NewRester(t)

		handler := http.HandlerFunc(appmock.ShowMovieHandler)

		expected := http.StatusOK

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/movies/1", nil)
		appmock.On("ShowMovieHandler", rr, req).Return(expected).Once()

		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, expected)
		appmock.AssertExpectations(t)
	})
}

func TestUpdateMovieHandlerUnit(t *testing.T) {

	t.Run("updates movie and returns 200", func(t *testing.T) {

		appmock := mocks.NewRester(t)

		handler := http.HandlerFunc(appmock.UpdateMovieHandler)

		expected := http.StatusOK

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/v1/movies/1", nil)
		appmock.On("UpdateMovieHandler", rr, req).Return(expected).Once()

		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, expected)
		appmock.AssertExpectations(t)
	})
}

// Integration Tests

func TestCreateMovieHandlerIntegration(t *testing.T) {

	t.Skip("skipping test")

	t.Run("creates a movie", func(t *testing.T) {

		cfg, _ := config.InitConfig()

		valid := middleware.NewValidator()

		mid := middleware.New(valid)

		trace := logger.TRACE

		zcfg := logger.LogConfig{
			Environment: cfg.Environment,
			LogLevel:    trace,
		}

		zlog := logger.NewZeroLoggerSrv(zcfg)

		db, err := sql.Open("postgres", "postgresql://bersen:bersen@localhost/filmdb")
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

		newMovie := `{
                        "title": "Bersen",
						"year": 1972,
						"runtime": "108 mins",
						"genres": ["animation", "adventure"]
				     }`

		handler := http.HandlerFunc(app.CreateMovieHandler)

		req := httptest.NewRequest(http.MethodPost, "/v1/movies", bytes.NewBuffer([]byte(newMovie)))

		r := httptest.NewRecorder()

		handler.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})
}

func TestDeleteMovieHandlerIntegration(t *testing.T) {

	t.Skip("Skipping test")

	t.Run("deletes a movie", func(t *testing.T) {

		cfg, _ := config.InitConfig()

		valid := middleware.NewValidator()

		mid := middleware.New(valid)

		trace := logger.TRACE

		zcfg := logger.LogConfig{
			Environment: cfg.Environment,
			LogLevel:    trace,
		}

		zlog := logger.NewZeroLoggerSrv(zcfg)

		db, err := sql.Open("postgres", "postgresql://bersen:bersen@localhost/filmdb")
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

		router := httprouter.New()

		router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", http.HandlerFunc(app.DeleteMovieHandler))

		req := httptest.NewRequest(http.MethodDelete, "/v1/movies/7", nil)

		r := httptest.NewRecorder()

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestShowMovieHandlerIntegration(t *testing.T) {

	t.Skip("Skipping ShowMovieHandlerIntegration test")

	t.Run("shows a movie", func(t *testing.T) {

		cfg, _ := config.InitConfig()

		valid := middleware.NewValidator()

		mid := middleware.New(valid)

		trace := logger.TRACE

		zcfg := logger.LogConfig{
			Environment: cfg.Environment,
			LogLevel:    trace,
		}

		zlog := logger.NewZeroLoggerSrv(zcfg)

		db, err := sql.Open("postgres", "postgresql://bersen:bersen@localhost/filmdb")
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

		router := httprouter.New()

		router.HandlerFunc(http.MethodGet, "/v1/movies/:id", http.HandlerFunc(app.ShowMovieHandler))

		req := httptest.NewRequest(http.MethodGet, "/v1/movies/11", nil)

		r := httptest.NewRecorder()

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestUpdateMovieHandlerIntegration(t *testing.T) {

	t.Skip("Skipping UpdateMovieHandlerIntegration test")

	t.Run("updates a movie", func(t *testing.T) {

		cfg, _ := config.InitConfig()

		valid := middleware.NewValidator()

		mid := middleware.New(valid)

		trace := logger.TRACE

		zcfg := logger.LogConfig{
			Environment: cfg.Environment,
			LogLevel:    trace,
		}

		zlog := logger.NewZeroLoggerSrv(zcfg)

		db, err := sql.Open("postgres", "postgresql://bersen:bersen@localhost/filmdb")
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

		newMovie := `{
                        "title": "Tatie",
						"year": 1972,
						"runtime": "108 mins",
						"genres": ["animation", "adventure"]
				     }`

		payload := bytes.NewBuffer([]byte(newMovie))

		router := httprouter.New()

		router.HandlerFunc(http.MethodPut, "/v1/movies/:id", http.HandlerFunc(app.UpdateMovieHandler))

		req := httptest.NewRequest(http.MethodPut, "/v1/movies/11", payload)

		r := httptest.NewRecorder()

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}
