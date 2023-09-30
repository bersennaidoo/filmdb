package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bersennaidoo/filmdb/mocks"
	"github.com/stretchr/testify/assert"
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
		req, _ := http.NewRequest("DELETE", "/v1/movies/:id", nil)
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
		req, _ := http.NewRequest("GET", "/v1/movies/:id", nil)
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
		req, _ := http.NewRequest("PUT", "/v1/movies/:id", nil)
		appmock.On("UpdateMovieHandler", rr, req).Return(expected).Once()

		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, expected)
		appmock.AssertExpectations(t)
	})
}
