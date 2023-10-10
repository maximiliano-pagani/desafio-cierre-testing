package products

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createTestServer(mockRepository *MockRepository) *gin.Engine {
	r := gin.Default()

	service := NewService(mockRepository)
	handler := NewHandler(service)

	g := r.Group("/api/v1/products")
	g.GET("", handler.GetProducts)

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestHandlerGetProducts(t *testing.T) {
	t.Run("TestHandlerGetProductsOK", func(t *testing.T) {
		// Arrange
		sellerId := "Seller1"
		server := createTestServer(&MockRepository{
			Products: []Product{
				{
					ID:          "123",
					SellerID:    sellerId,
					Description: "Description",
					Price:       100.0,
				},
			},
			Err: nil,
		})
		req, respRec := createRequestTest(http.MethodGet, "/api/v1/products?seller_id="+sellerId, "")
		expectedCode := http.StatusOK
		expectedBody, err := json.Marshal(
			[]Product{
				{
					ID:          "123",
					SellerID:    sellerId,
					Description: "Description",
					Price:       100.0,
				},
			},
		)
		assert.Nil(t, err)

		// Act
		server.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, expectedCode, respRec.Code)
		assert.Equal(t, expectedBody, respRec.Body.Bytes())
	})

	t.Run("TestHandlerGetProductsNotFound", func(t *testing.T) {
		// Arrange
		sellerId := "Seller1"
		server := createTestServer(&MockRepository{
			Products: []Product{},
			Err:      nil,
		})
		req, respRec := createRequestTest(http.MethodGet, "/api/v1/products?seller_id="+sellerId, "")
		expectedCode := http.StatusOK
		expectedBody, err := json.Marshal([]Product{})
		assert.Nil(t, err)

		// Act
		server.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, expectedCode, respRec.Code)
		assert.Equal(t, expectedBody, respRec.Body.Bytes())
	})

	t.Run("TestHandlerGetProductsError", func(t *testing.T) {
		// Arrange
		sellerId := "Seller1"
		server := createTestServer(&MockRepository{
			Products: []Product{},
			Err:      errors.New(""),
		})
		req, respRec := createRequestTest(http.MethodGet, "/api/v1/products?seller_id="+sellerId, "")
		expectedCode := http.StatusInternalServerError

		// Act
		server.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, expectedCode, respRec.Code)
	})

	t.Run("TestHandlerGetProductsBadRequest", func(t *testing.T) {
		// Arrange
		server := createTestServer(&MockRepository{})
		req, respRec := createRequestTest(http.MethodGet, "/api/v1/products", "")
		expectedCode := http.StatusBadRequest

		// Act
		server.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, expectedCode, respRec.Code)
	})
}
