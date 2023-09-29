package storage_test

import (
	"testing"

	"github.com/bersennaidoo/filmdb/internal/domain/models"
	"github.com/bersennaidoo/filmdb/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateUnit(t *testing.T) {

	t.Run("updates movie given a movie return nil for error", func(t *testing.T) {

		movie := models.Movie{}

		pgstore := mocks.NewStorer(t)
		pgstore.On("Update", &movie).Return(nil).Once()

		err := pgstore.Update(&movie)

		require.Nil(t, err)
		pgstore.AssertExpectations(t)
	})
}

func TestInsertUnit(t *testing.T) {

	t.Run("inserts a movie and return nil for error", func(t *testing.T) {

		//Arrange
		movie := &models.Movie{}
		pgstore := mocks.NewStorer(t)

		pgstore.On("Insert", movie).Return(nil).Once()

		//Act
		err := pgstore.Insert(movie)

		//Assert
		require.Nil(t, err)
		pgstore.AssertExpectations(t)
	})
}

func TestDeleteUnit(t *testing.T) {

	t.Run("deletes movie and returns nil for error", func(t *testing.T) {

		id := int64(1)

		pgstore := mocks.NewStorer(t)

		pgstore.On("Delete", id).Return(nil).Once()

		err := pgstore.Delete(id)

		require.Nil(t, err)
		pgstore.AssertExpectations(t)
	})
}

func TestGetUnit(t *testing.T) {

	t.Run("returns movie and nil for error", func(t *testing.T) {

		expectedMovie := models.Movie{}

		var id int64

		pgstore := mocks.NewStorer(t)

		pgstore.On("Get", id).Return(&expectedMovie, nil).Once()

		movie, err := pgstore.Get(id)

		require.Nil(t, err)
		assert.Equal(t, *movie, expectedMovie, "they should be equal")
		pgstore.AssertExpectations(t)
	})
}
