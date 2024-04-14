package routes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"product-microservice/controllers"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductController struct {
	mock.Mock
	controllers.ProductController
}

func (m *MockProductController) CreateProduct(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockProductController) UpdateProduct(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockProductController) GetProducts(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockProductController) GetProductsById(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockProductController) GetProductsByCategory(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockProductController) SearchProducts(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockProductController) DeleteProduct(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockProductController) Premium(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockProductController) GetProductsPremium(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func TestProductRoutes(t *testing.T) {
	e := echo.New()
	mockController := new(MockProductController)
	ProductRoutes(e, mockController)
	t.Run("CreateProduct", func(t *testing.T) {
		formData := bytes.NewBufferString("name=TestProduct&description=JustATest")
		req := httptest.NewRequest(http.MethodPost, "/products", formData)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		e.NewContext(req, rec)
		mockController.On("CreateProduct", mock.Anything).Return(nil)
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockController.AssertExpectations(t)
	})
	t.Run("GetProducts", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		rec := httptest.NewRecorder()
		e.NewContext(req, rec)
		mockController.On("GetProducts", mock.AnythingOfType("*echo.context")).Return(nil).Once()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockController.AssertExpectations(t)
	})
	
}
