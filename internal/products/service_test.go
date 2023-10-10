package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceGetAllBySellerUnitTests(t *testing.T) {
	t.Run("TestServiceGetAllBySellerOK", func(t *testing.T) {
		// Arrange
		sellerId := "Seller1"
		repository := &MockRepository{
			Products: []Product{
				{
					ID:          "123",
					SellerID:    sellerId,
					Description: "Description",
					Price:       100.0,
				},
			},
			Err: nil,
		}
		service := NewService(repository)
		expectedResult := []Product{
			{
				ID:          "123",
				SellerID:    sellerId,
				Description: "Description",
				Price:       100.0,
			},
		}

		// Act
		result, resultErr := service.GetAllBySeller(sellerId)

		// Assert
		assert.Nil(t, resultErr, "error should be nil")
		assert.Equal(t, expectedResult, result, "product should be retrieved from service successfully")
	})

	t.Run("TestServiceGetAllBySellerNotFound", func(t *testing.T) {
		// Arrange
		sellerId := "Seller1"
		repository := &MockRepository{Products: []Product{}, Err: nil}
		service := NewService(repository)
		expectedResult := []Product{}
		expectedLen := 0

		// Act
		result, err := service.GetAllBySeller(sellerId)

		// Assert
		assert.Nil(t, err, "error should be nil")
		assert.Equal(t, expectedLen, len(result), "result should be an empty slice of products")
		assert.Equal(t, expectedResult, result, "result should be an empty slice of products")
	})

	t.Run("TestServiceGetAllBySellerError", func(t *testing.T) {
		// Arrange
		sellerId := "Seller1"
		repository := &MockRepository{Products: []Product{}, Err: errors.New("")}
		service := NewService(repository)

		// Act
		result, err := service.GetAllBySeller(sellerId)

		// Assert
		assert.NotNil(t, err, "service should return error")
		assert.Nil(t, result, "service should return nil slice pointer")
	})
}

func TestServiceGetAllBySellerIntegrationTests(t *testing.T) {
	t.Run("TestServiceGetAllBySellerOK", func(t *testing.T) {
		// Arrange
		repository := NewRepository([]Product{
			{
				ID:          "123",
				SellerID:    "Seller1",
				Description: "Description1",
				Price:       100.0,
			},
			{
				ID:          "456",
				SellerID:    "Seller2",
				Description: "Description2",
				Price:       200.0,
			},
		})
		service := NewService(repository)
		expectedResult := []Product{
			{
				ID:          "456",
				SellerID:    "Seller2",
				Description: "Description2",
				Price:       200.0,
			},
		}

		// Act
		result, resultErr := service.GetAllBySeller("Seller2")

		// Assert
		assert.Nil(t, resultErr, "error should be nil")
		assert.Equal(t, expectedResult, result, "product should be retrieved from service successfully")
	})

	t.Run("TestServiceGetAllBySellerNotFound", func(t *testing.T) {
		// Arrange
		repository := NewRepository([]Product{
			{
				ID:          "123",
				SellerID:    "Seller1",
				Description: "Description1",
				Price:       100.0,
			},
			{
				ID:          "456",
				SellerID:    "Seller2",
				Description: "Description2",
				Price:       200.0,
			},
		})
		service := NewService(repository)
		expectedResult := []Product{}
		expectedLen := 0

		// Act
		result, err := service.GetAllBySeller("Seller3")

		// Assert
		assert.Nil(t, err, "error should be nil")
		assert.Equal(t, expectedLen, len(result), "result should be an empty slice of products")
		assert.Equal(t, expectedResult, result, "result should be an empty slice of products")
	})

	t.Run("TestServiceGetAllBySellerInvalidIdError", func(t *testing.T) {
		// Arrange
		repository := NewRepository([]Product{
			{
				ID:          "123",
				SellerID:    "Seller1",
				Description: "Description1",
				Price:       100.0,
			},
			{
				ID:          "456",
				SellerID:    "Seller2",
				Description: "Description2",
				Price:       200.0,
			},
		})
		service := NewService(repository)

		// Act
		result, err := service.GetAllBySeller("")

		// Assert
		assert.NotNil(t, err, "service should return error")
		assert.Nil(t, result, "service should return nil slice pointer")
	})
}
