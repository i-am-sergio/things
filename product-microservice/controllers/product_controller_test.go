package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"product-microservice/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockProductService struct {
	mock.Mock
}
func (m *MockProductService) CreateProductService(form *multipart.Form, image *multipart.FileHeader) (models.Product, error) {
	args := m.Called(form, image)
	return args.Get(0).(models.Product), args.Error(1)
}
func (m *MockProductService) UpdateProductService(productID uint, form *multipart.Form, image *multipart.FileHeader) (models.Product, error) {
	args := m.Called(productID, form, image)
	return args.Get(0).(models.Product), args.Error(1)
}
func (m *MockProductService) GetProductsService() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}
func (m *MockProductService) GetProductByIDService(productID uint) (models.Product, error) {
	args := m.Called(productID)
	return args.Get(0).(models.Product), args.Error(1)
}
func (m *MockProductService) GetProductsByCategoryService(category string) ([]models.Product, error) {
	args := m.Called(category)
	return args.Get(0).([]models.Product), args.Error(1)
}
func (m *MockProductService) SearchProductsService(searchTerm string) ([]models.Product, error) {
	args := m.Called(searchTerm)
	return args.Get(0).([]models.Product), args.Error(1)
}
func (m *MockProductService) DeleteProductService(productID uint) error {
	args := m.Called(productID)
	return args.Error(0)
}
func (m *MockProductService) PremiumService(productID uint) (models.Product, error) {
	args := m.Called(productID)
	return args.Get(0).(models.Product), args.Error(1)
}
func (m *MockProductService) GetProductsPremiumService() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func TestCreateProduct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		e := echo.New()
		formData := new(bytes.Buffer)
		writer := multipart.NewWriter(formData)
		_ = writer.WriteField("name", "Test Product")
		file, _ := writer.CreateFormFile("image", "test.png")
		file.Write([]byte("fake image data"))
		writer.Close()
		req := httptest.NewRequest(http.MethodPost, "/products", formData)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		expectedProduct := models.Product{ID: 1, Name: "Test Product"}
		mockProductService := new(MockProductService)
		form, _ := c.MultipartForm()
		fileHeader, _ := c.FormFile("image")
		mockProductService.On("CreateProductService", form, fileHeader).Return(expectedProduct, nil)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.CreateProduct(c))
		assert.Equal(t, http.StatusCreated, rec.Code)
		var product models.Product
		json.Unmarshal(rec.Body.Bytes(), &product)
		assert.Equal(t, expectedProduct, product)
		mockProductService.AssertExpectations(t)
	})
	t.Run("MultipartFormError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/products", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		mockProductService := new(MockProductService)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.CreateProduct(c))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		var responseBody map[string]string
		json.Unmarshal(rec.Body.Bytes(), &responseBody)
		require.Contains(t, responseBody, "error")
	})
	t.Run("FormFileError", func(t *testing.T) {
		e := echo.New()
		formData := new(bytes.Buffer)
		writer := multipart.NewWriter(formData)
		_ = writer.WriteField("name", "Test Product")
		writer.Close()
		req := httptest.NewRequest(http.MethodPost, "/products", formData)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		mockProductService := new(MockProductService)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.CreateProduct(c))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		var responseBody map[string]string
		json.Unmarshal(rec.Body.Bytes(), &responseBody)
		require.Contains(t, responseBody, "error")
		assert.Equal(t, "http: no such file", responseBody["error"])
	})
	t.Run("ServiceError", func(t *testing.T) {
		e := echo.New()
		formData := new(bytes.Buffer)
		writer := multipart.NewWriter(formData)
		_ = writer.WriteField("name", "Test Product")
		file, _ := writer.CreateFormFile("image", "test.png")
		file.Write([]byte("fake image data"))
		writer.Close()
		req := httptest.NewRequest(http.MethodPost, "/products", formData)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		mockProductService := new(MockProductService)
		form, _ := c.MultipartForm()
		fileHeader, _ := c.FormFile("image")
		expectedError := errors.New("failed to create product")
		mockProductService.On("CreateProductService", form, fileHeader).Return(models.Product{}, expectedError)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.CreateProduct(c))
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		var responseBody map[string]string
		json.Unmarshal(rec.Body.Bytes(), &responseBody)
		require.Contains(t, responseBody, "error")
		assert.Equal(t, "failed to create product", responseBody["error"])
		mockProductService.AssertExpectations(t)
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
        e := echo.New()
        formData := new(bytes.Buffer)
        writer := multipart.NewWriter(formData)
        _ = writer.WriteField("name", "Updated Product")
        file, _ := writer.CreateFormFile("image", "update.png")
        file.Write([]byte("updated image data"))
        writer.Close()
        req := httptest.NewRequest(http.MethodPut, "/products/1", formData)
        req.Header.Set("Content-Type", writer.FormDataContentType())
        rec := httptest.NewRecorder()
        c := e.NewContext(req, rec)
        c.SetParamNames("id")
        c.SetParamValues("1")
        expectedProduct := models.Product{ID: 1, Name: "Updated Product"}
        mockProductService := new(MockProductService)
        form, _ := c.MultipartForm()
        fileHeader, _ := c.FormFile("image")
        mockProductService.On("UpdateProductService", uint(1), form, fileHeader).Return(expectedProduct, nil)
        controller := NewProductController(mockProductService)
        require.NoError(t, controller.UpdateProduct(c))
        assert.Equal(t, http.StatusOK, rec.Code)
        var product models.Product
        json.Unmarshal(rec.Body.Bytes(), &product)
        assert.Equal(t, expectedProduct, product)
        mockProductService.AssertExpectations(t)
    })
	t.Run("IDConversionError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/products/invalid-id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("invalid-id")
		controller := NewProductController(nil)
		err := controller.UpdateProduct(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		var responseBody map[string]string
		json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.Contains(t, responseBody, "error")
		assert.Equal(t, invalidCommentIDError, responseBody["error"])
	})
	t.Run("MultipartFormError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/products/1", nil)
		req.Header.Set("Content-Type", "multipart/form-data")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")
		mockProductService := new(MockProductService)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.UpdateProduct(c))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		var responseBody map[string]string
		json.Unmarshal(rec.Body.Bytes(), &responseBody)
		require.Contains(t, responseBody, "error")
	})
	t.Run("ServiceError", func(t *testing.T) {
		e := echo.New()
		formData := new(bytes.Buffer)
		writer := multipart.NewWriter(formData)
		_ = writer.WriteField("name", "Updated Product")
		file, _ := writer.CreateFormFile("image", "update.png")
		file.Write([]byte("updated image data"))
		writer.Close()
		req := httptest.NewRequest(http.MethodPut, "/products/1", formData)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")
		expectedError := errors.New("failed to update product")
		mockProductService := new(MockProductService)
		form, _ := c.MultipartForm()
		fileHeader, _ := c.FormFile("image")
		mockProductService.On("UpdateProductService", uint(1), form, fileHeader).Return(models.Product{}, expectedError)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.UpdateProduct(c))
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		var responseBody map[string]string
		json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.Contains(t, responseBody, "error")
		assert.Equal(t, "failed to update product", responseBody["error"])
		mockProductService.AssertExpectations(t)
	})
}

func TestGetProducts(t *testing.T) {
    t.Run("Success", func(t *testing.T) {
        e := echo.New()
        req := httptest.NewRequest(http.MethodGet, "/products", nil)
        rec := httptest.NewRecorder()
        c := e.NewContext(req, rec)
        expectedProducts := []models.Product{
            {ID: 1, Name: "Product 1"},
            {ID: 2, Name: "Product 2"},
        }
        mockProductService := new(MockProductService)
        mockProductService.On("GetProductsService").Return(expectedProducts, nil)
        controller := NewProductController(mockProductService)
        require.NoError(t, controller.GetProducts(c))
        assert.Equal(t, http.StatusOK, rec.Code)
        var products []models.Product
        require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &products))
        assert.Equal(t, expectedProducts, products)
        mockProductService.AssertExpectations(t)
    })
	t.Run("ServiceError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		expectedError := errors.New("failed to fetch products")
		mockProductService := new(MockProductService)
		var nilProducts []models.Product
		mockProductService.On("GetProductsService").Return(nilProducts, expectedError)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.GetProducts(c))
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		var responseBody map[string]string
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &responseBody))
		assert.Contains(t, responseBody, "error")
		assert.Equal(t, "failed to fetch products", responseBody["error"])
		mockProductService.AssertExpectations(t)
	})
}

func TestGetProductsById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")
		expectedProduct := models.Product{ID: 1, Name: "Product 1"}
		mockProductService := new(MockProductService)
		mockProductService.On("GetProductByIDService", uint(1)).Return(expectedProduct, nil)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.GetProductsById(c))
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("IDConversionError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products/invalid-id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("invalid-id")
		controller := NewProductController(nil)
		err := controller.GetProductsById(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		var responseBody map[string]string
		json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.Contains(t, responseBody, "error")
		assert.Equal(t, invalidCommentIDError, responseBody["error"])
	})
	t.Run("ServiceError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")
		expectedError := errors.New("failed to fetch product")
		mockProductService := new(MockProductService)
		mockProductService.On("GetProductByIDService", uint(1)).Return(models.Product{}, expectedError)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.GetProductsById(c))
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		var responseBody map[string]string
		json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.Contains(t, responseBody, "error")
		assert.Equal(t, "failed to fetch product", responseBody["error"])
		mockProductService.AssertExpectations(t)
	})
}

func TestGetProductsByCategory(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
        e := echo.New()
        req := httptest.NewRequest(http.MethodGet, "/products?category=test", nil)
        rec := httptest.NewRecorder()
        c := e.NewContext(req, rec)
        expectedProducts := []models.Product{
            {ID: 1, Name: "Product 1", Category: "test"},
            {ID: 2, Name: "Product 2", Category: "test"},
        }
        mockProductService := new(MockProductService)
        mockProductService.On("GetProductsByCategoryService", "test").Return(expectedProducts, nil)
        controller := NewProductController(mockProductService)
        require.NoError(t, controller.GetProductsByCategory(c))
        assert.Equal(t, http.StatusOK, rec.Code)
        var products []models.Product
        require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &products))
        assert.Equal(t, expectedProducts, products)
        mockProductService.AssertExpectations(t)
    })
	t.Run("ServiceError", func(t *testing.T) {
        e := echo.New()
        req := httptest.NewRequest(http.MethodGet, "/products?category=test", nil)
        rec := httptest.NewRecorder()
        c := e.NewContext(req, rec)
        expectedError := errors.New("failed to fetch products by category")
        mockProductService := new(MockProductService)
		var nilProducts []models.Product
        mockProductService.On("GetProductsByCategoryService", "test").Return(nilProducts, expectedError)
        controller := NewProductController(mockProductService)
        require.NoError(t, controller.GetProductsByCategory(c))
        assert.Equal(t, http.StatusInternalServerError, rec.Code)
        var responseBody map[string]string
        require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &responseBody))
        assert.Contains(t, responseBody, "error")
        assert.Equal(t, expectedError.Error(), responseBody["error"])
        mockProductService.AssertExpectations(t)
    })
}

func TestSearchProducts(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products?q=test", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		expectedProducts := []models.Product{
			{ID: 1, Name: "Product 1", Description: "test description"},
			{ID: 2, Name: "Product 2", Description: "test description"},
		}
		mockProductService := new(MockProductService)
		mockProductService.On("SearchProductsService", "test").Return(expectedProducts, nil)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.SearchProducts(c))
		assert.Equal(t, http.StatusOK, rec.Code)
		var products []models.Product
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &products))
		assert.Equal(t, expectedProducts, products)
		mockProductService.AssertExpectations(t)
	})
	t.Run("ServiceError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products?q=test", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		expectedError := errors.New("failed to search products")
		mockProductService := new(MockProductService)
		var nilProducts []models.Product
		mockProductService.On("SearchProductsService", "test").Return(nilProducts, expectedError)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.SearchProducts(c))
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		var responseBody map[string]string
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &responseBody))
		assert.Contains(t, responseBody, "error")
		assert.Equal(t, expectedError.Error(), responseBody["error"])
		mockProductService.AssertExpectations(t)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")
		mockProductService := new(MockProductService)
		mockProductService.On("DeleteProductService", uint(1)).Return(nil)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.DeleteProduct(c))
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("IDConversionError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/products/invalid-id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("invalid-id")
		controller := NewProductController(nil)
		err := controller.DeleteProduct(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		var responseBody map[string]string
		json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.Contains(t, responseBody, "error")
		assert.Equal(t, invalidCommentIDError, responseBody["error"])
	})
	t.Run("ServiceError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")
		expectedError := errors.New("failed to delete product")
		mockProductService := new(MockProductService)
		mockProductService.On("DeleteProductService", uint(1)).Return(expectedError)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.DeleteProduct(c))
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		var responseBody map[string]string
		json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.Contains(t, responseBody, "error")
		assert.Equal(t, expectedError.Error(), responseBody["error"])
		mockProductService.AssertExpectations(t)
	})
}

func TestPremium(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/products/premium/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")
		expectedProduct := models.Product{ID: 1, Name: "Product 1"}
		mockProductService := new(MockProductService)
		mockProductService.On("PremiumService", uint(1)).Return(expectedProduct, nil)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.Premium(c))
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("IDConversionError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/products/premium/invalid-id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("invalid-id")
		controller := NewProductController(nil)
		err := controller.Premium(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		var responseBody map[string]string
		json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.Contains(t, responseBody, "error")
		assert.Equal(t, invalidCommentIDError, responseBody["error"])
	})
	t.Run("ServiceError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/products/premium/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")
		expectedError := errors.New("failed to premium product")
		mockProductService := new(MockProductService)
		mockProductService.On("PremiumService", uint(1)).Return(models.Product{}, expectedError)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.Premium(c))
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		var responseBody map[string]string
		json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.Contains(t, responseBody, "error")
		assert.Equal(t, expectedError.Error(), responseBody["error"])
		mockProductService.AssertExpectations(t)
	})
}

func TestGetProductsPremium(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products/premium", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		expectedProducts := []models.Product{
			{ID: 1, Name: "Product 1", Status: true},
			{ID: 2, Name: "Product 2", Status: true},
		}
		mockProductService := new(MockProductService)
		mockProductService.On("GetProductsPremiumService").Return(expectedProducts, nil)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.GetProductsPremium(c))
		assert.Equal(t, http.StatusOK, rec.Code)
		var products []models.Product
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &products))
		assert.Equal(t, expectedProducts, products)
		mockProductService.AssertExpectations(t)
	})
	t.Run("ServiceError", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products/premium", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		expectedError := errors.New("failed to fetch premium products")
		mockProductService := new(MockProductService)
		var nilProducts []models.Product
		mockProductService.On("GetProductsPremiumService").Return(nilProducts, expectedError)
		controller := NewProductController(mockProductService)
		require.NoError(t, controller.GetProductsPremium(c))
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		var responseBody map[string]string
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &responseBody))
		assert.Contains(t, responseBody, "error")
		assert.Equal(t, expectedError.Error(), responseBody["error"])
		mockProductService.AssertExpectations(t)
	})
}