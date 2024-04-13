package services

import (
	"mime/multipart"
	"product-microservice/db"
	"product-microservice/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockDBClient struct {
	mock.Mock
}

func (m *MockDBClient) AutoMigrate(models ...interface{}) error {
    args := m.Called(models)
    return args.Error(0)
}
func (m *MockDBClient) First(value interface{}, conditions ...interface{}) error {
    args := m.Called(value, conditions)
    return args.Error(0)
}
func (m *MockDBClient) Save(value interface{}) error {
    args := m.Called(value)
    return args.Error(0)
}
func (m *MockDBClient) Create(value interface{}) error {
    args := m.Called(value)
    return args.Error(0)
}
func (m *MockDBClient) FindPreloaded(relation string, value interface{}, conditions ...interface{}) error {
	args := m.Called(relation, value, conditions)
	return args.Error(0)
}
func (m *MockDBClient) Find(value interface{}, conditions ...interface{}) error {
	args := m.Called(value, conditions)
    return args.Error(0)
}
func (m *MockDBClient) FindWithCondition(value interface{}, condition string, args ...interface{}) error {
    mockArgs := m.Called(value, condition, args)
    return mockArgs.Error(0)
}
func (m *MockDBClient) Delete(value interface{}) error {
	args := m.Called(value)
	return args.Error(0)
}
func (m *MockDBClient) DeleteWithCondition(model interface{}, query string, args ...interface{}) error {
    mockArgs := m.Called(model, query, args)
    return mockArgs.Error(0)
}
func (m *MockDBClient) DeleteByID(model interface{}, id interface{}) error {
    mockArgs := m.Called(model, id)
    return mockArgs.Error(0)
}

type MockCloudinaryClient struct {
    mock.Mock
}
func (m *MockCloudinaryClient) InitCloudinary(envLoader db.EnvLoader) error {
	args := m.Called(envLoader)
	return args.Error(0)
}
func (m *MockCloudinaryClient) UploadImage(file db.FileHeaderWrapper) (string, error) {
    args := m.Called(file)
    return args.String(0), args.Error(1)
}

func TestCreateProductService(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockDB := new(MockDBClient)
		mockDB.On("Create", mock.AnythingOfType("*models.Product")).Return(nil)
		mockCloudinary := new(MockCloudinaryClient)
        mockCloudinary.On("UploadImage", mock.Anything).Return("cloudinary_url", nil)
		service := NewProductService(mockDB, mockCloudinary)
		form := &multipart.Form{
            Value: map[string][]string{
                "UserID":      {"1"},
                "Price":       {"10.5"},
                "Name":        {"Test Product"},
                "Description": {"Test Description"},
                "Category":    {"Test Category"},
                "Ubication":   {"Test Ubication"},
            },
        }
		image := &multipart.FileHeader{}
		product, err := service.CreateProductService(form, image)
        require.NoError(t, err)
		assert.Equal(t, uint(1), product.UserID)
        assert.Equal(t, "Test Product", product.Name)
        assert.Equal(t, "Test Description", product.Description)
        assert.Equal(t, "Test Category", product.Category)
        assert.Equal(t, 10.5, product.Price)
        assert.Equal(t, "cloudinary_url", product.Image)
        assert.Equal(t, "Test Ubication", product.Ubication)
	})
	t.Run("Missing required field", func(t *testing.T) {
		mockDB := new(MockDBClient)
		mockCloudinary := new(MockCloudinaryClient)
		service := NewProductService(mockDB, mockCloudinary)
		form := &multipart.Form{
			Value: map[string][]string{
				"UserID":      {"1"},
				"Price":       {"10.5"},
				"Name":        {"Test Product"},
				"Description": {"Test Description"},
				"Category":    {"Test Category"},
			},
		}
		image := &multipart.FileHeader{}
		product, err := service.CreateProductService(form, image)
		require.Error(t, err)
		assert.Equal(t, "missing required field: Ubication", err.Error())
		assert.Equal(t, uint(0), product.UserID)
	})
	t.Run("Invalid UserID", func(t *testing.T) {
		mockDB := new(MockDBClient)
		mockCloudinary := new(MockCloudinaryClient)
		service := NewProductService(mockDB, mockCloudinary)
		form := &multipart.Form{
			Value: map[string][]string{
				"UserID":      {"invalid"},
				"Price":       {"10.5"},
				"Name":        {"Test Product"},
				"Description": {"Test Description"},
				"Category":    {"Test Category"},
				"Ubication":   {"Test Ubication"},
			},
		}
		image := &multipart.FileHeader{}
		product, err := service.CreateProductService(form, image)
		require.Error(t, err)
		assert.Equal(t, "strconv.ParseUint: parsing \"invalid\": invalid syntax", err.Error())
		assert.Equal(t, uint(0), product.UserID)
	})
	t.Run("Invalid Price", func(t *testing.T) {
		mockDB := new(MockDBClient)
		mockCloudinary := new(MockCloudinaryClient)
		service := NewProductService(mockDB, mockCloudinary)
		form := &multipart.Form{
			Value: map[string][]string{
				"UserID":      {"1"},
				"Price":       {"invalid"},
				"Name":        {"Test Product"},
				"Description": {"Test Description"},
				"Category":    {"Test Category"},
				"Ubication":   {"Test Ubication"},
			},
		}
		image := &multipart.FileHeader{}
		product, err := service.CreateProductService(form, image)
		require.Error(t, err)
		assert.Equal(t, "strconv.ParseFloat: parsing \"invalid\": invalid syntax", err.Error())
		assert.Equal(t, uint(0), product.UserID)
	})
	t.Run("Error uploading image", func(t *testing.T) {
		mockDB := new(MockDBClient)
		mockCloudinary := new(MockCloudinaryClient)
		mockCloudinary.On("UploadImage", mock.Anything).Return("", assert.AnError)
		service := NewProductService(mockDB, mockCloudinary)
		form := &multipart.Form{
			Value: map[string][]string{
				"UserID":      {"1"},
				"Price":       {"10.5"},
				"Name":        {"Test Product"},
				"Description": {"Test Description"},
				"Category":    {"Test Category"},
				"Ubication":   {"Test Ubication"},
			},
		}
		image := &multipart.FileHeader{}
		product, err := service.CreateProductService(form, image)
		require.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		assert.Equal(t, uint(0), product.UserID)
	})
	t.Run("Error creating product", func(t *testing.T) {
		mockDB := new(MockDBClient)
		mockDB.On("Create", mock.AnythingOfType("*models.Product")).Return(assert.AnError)
		mockCloudinary := new(MockCloudinaryClient)
		mockCloudinary.On("UploadImage", mock.Anything).Return("cloudinary_url", nil)
		service := NewProductService(mockDB, mockCloudinary)
		form := &multipart.Form{
			Value: map[string][]string{
				"UserID":      {"1"},
				"Price":       {"10.5"},
				"Name":        {"Test Product"},
				"Description": {"Test Description"},
				"Category":    {"Test Category"},
				"Ubication":   {"Test Ubication"},
			},
		}
		image := &multipart.FileHeader{}
		product, err := service.CreateProductService(form, image)
		require.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		assert.Equal(t, uint(1), product.UserID)
	})
}

func TestUpdateProductService(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		formUpdate := &multipart.Form{
            Value: map[string][]string{
                "UserID":      {"2"},
                "Price":       {"11.5"},
                "Name":        {"Test Product Update"},
                "Description": {"Test Description Update"},
                "Category":    {"Test Category Update"},
                "Ubication":   {"Test Ubication Update"},
            },
        }
        image := &multipart.FileHeader{}
		product	:= models.Product{
			UserID:      1,
			Name:        "Test Product",
			Description: "Test Description",
			Category:    "Test Category",
			Price:       10.5,
			Rate:        0.0,
			Ubication:   "Test Ubication",
			Image:       "cloudinary_url",
		}
		mockDB := new(MockDBClient)
		mockDB.On("First", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*models.Product)
			*arg = product
		}).Return(nil)
		mockDB.On("Save", mock.AnythingOfType("*models.Product")).Return(nil)
		mockCloudinary := new(MockCloudinaryClient)
		mockCloudinary.On("UploadImage", mock.Anything).Return("cloudinary_url", nil)
		service := NewProductService(mockDB, mockCloudinary)
		product, err := service.UpdateProductService(1, formUpdate, image)
		require.NoError(t, err)
		assert.Equal(t, uint(2), product.UserID)
		assert.Equal(t, "Test Product Update", product.Name)
		assert.Equal(t, "Test Description Update", product.Description)
		assert.Equal(t, "Test Category Update", product.Category)
		assert.Equal(t, 11.5, product.Price)
		assert.Equal(t, "cloudinary_url", product.Image)
		assert.Equal(t, "Test Ubication Update", product.Ubication)
	})
	t.Run("Product not found", func(t *testing.T) {
		mockDB := new(MockDBClient)
		mockDB.On("First", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Return(assert.AnError)
		mockCloudinary := new(MockCloudinaryClient)
		service := NewProductService(mockDB, mockCloudinary)
		formUpdate := &multipart.Form{}
		image := &multipart.FileHeader{}
		product, err := service.UpdateProductService(1, formUpdate, image)
		require.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		assert.Equal(t, uint(0), product.UserID)
	})
	t.Run("Missing required field", func(t *testing.T) {
		product	:= models.Product{
			UserID:      1,
			Name:        "Test Product",
			Description: "Test Description",
			Category:    "Test Category",
			Price:       10.5,
			Rate:        0.0,
			Ubication:   "Test Ubication",
			Image:       "cloudinary_url",
		}
		mockDB := new(MockDBClient)
		mockDB.On("First", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*models.Product)
			*arg = product
		}).Return(nil)
		mockCloudinary := new(MockCloudinaryClient)
		service := NewProductService(mockDB, mockCloudinary)
		formUpdate := &multipart.Form{
			Value: map[string][]string{
				"UserID":      {"2"},
				"Price":       {"11.5"},
				"Name":        {"Test Product Update"},
				"Description": {"Test Description Update"},
				"Category":    {"Test Category Update"},
			},
		}
		image := &multipart.FileHeader{}
		product, err := service.UpdateProductService(1, formUpdate, image)
		require.Error(t, err)
		assert.Equal(t, "missing required field: Ubication", err.Error())
		assert.Equal(t, uint(1), product.UserID)
	})
	t.Run("Invalid UserID", func(t *testing.T) {
		product	:= models.Product{
			UserID:      1,
			Name:        "Test Product",
			Description: "Test Description",
			Category:    "Test Category",
			Price:       10.5,
			Rate:        0.0,
			Ubication:   "Test Ubication",
			Image:       "cloudinary_url",
		}
		mockDB := new(MockDBClient)
		mockDB.On("First", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*models.Product)
			*arg = product
		}).Return(nil)
		mockCloudinary := new(MockCloudinaryClient)
		service := NewProductService(mockDB, mockCloudinary)
		formUpdate := &multipart.Form{
			Value: map[string][]string{
				"UserID":      {"invalid"},
				"Price":       {"11.5"},
				"Name":        {"Test Product Update"},
				"Description": {"Test Description Update"},
				"Category":    {"Test Category Update"},
				"Ubication":   {"Test Ubication Update"},
			},
		}
		image := &multipart.FileHeader{}
		product, err := service.UpdateProductService(1, formUpdate, image)
		require.Error(t, err)
		assert.Equal(t, "strconv.ParseUint: parsing \"invalid\": invalid syntax", err.Error())
		assert.Equal(t, uint(1), product.UserID)
	})
	t.Run("Invalid Price", func(t *testing.T) {
		product	:= models.Product{
			UserID:      1,
			Name:        "Test Product",
			Description: "Test Description",
			Category:    "Test Category",
			Price:       10.5,
			Rate:        0.0,
			Ubication:   "Test Ubication",
			Image:       "cloudinary_url",
		}
		mockDB := new(MockDBClient)
		mockDB.On("First", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*models.Product)
			*arg = product
		}).Return(nil)
		mockCloudinary := new(MockCloudinaryClient)
		service := NewProductService(mockDB, mockCloudinary)
		formUpdate := &multipart.Form{
			Value: map[string][]string{
				"UserID":      {"2"},
				"Price":       {"invalid"},
				"Name":        {"Test Product Update"},
				"Description": {"Test Description Update"},
				"Category":    {"Test Category Update"},
				"Ubication":   {"Test Ubication Update"},
			},
		}
		image := &multipart.FileHeader{}
		product, err := service.UpdateProductService(1, formUpdate, image)
		require.Error(t, err)
		assert.Equal(t, "strconv.ParseFloat: parsing \"invalid\": invalid syntax", err.Error())
		assert.Equal(t, uint(1), product.UserID)
	})
	t.Run("Error uploading image", func(t *testing.T) {
		product	:= models.Product{
			UserID:      1,
			Name:        "Test Product",
			Description: "Test Description",
			Category:    "Test Category",
			Price:       10.5,
			Rate:        0.0,
			Ubication:   "Test Ubication",
			Image:       "cloudinary_url",
		}
		mockDB := new(MockDBClient)
		mockDB.On("First", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*models.Product)
			*arg = product
		}).Return(nil)
		mockCloudinary := new(MockCloudinaryClient)
		mockCloudinary.On("UploadImage", mock.Anything).Return("", assert.AnError)
		service := NewProductService(mockDB, mockCloudinary)
		formUpdate := &multipart.Form{
			Value: map[string][]string{
				"UserID":      {"2"},
				"Price":       {"11.5"},
				"Name":        {"Test Product Update"},
				"Description": {"Test Description Update"},
				"Category":    {"Test Category Update"},
				"Ubication":   {"Test Ubication Update"},
			},
		}
		image := &multipart.FileHeader{}
		product, err := service.UpdateProductService(1, formUpdate, image)
		require.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		assert.Equal(t, uint(2), product.UserID)
	})
	t.Run("Error updating product", func(t *testing.T) {
		product	:= models.Product{
			UserID:      1,
			Name:        "Test Product",
			Description: "Test Description",
			Category:    "Test Category",
			Price:       10.5,
			Rate:        0.0,
			Ubication:   "Test Ubication",
			Image:       "cloudinary_url",
		}
		mockDB := new(MockDBClient)
		mockDB.On("First", mock.AnythingOfType("*models.Product"), []interface{}{uint(1)}).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*models.Product)
			*arg = product
		}).Return(nil)
		mockDB.On("Save", mock.AnythingOfType("*models.Product")).Return(assert.AnError)
		mockCloudinary := new(MockCloudinaryClient)
		mockCloudinary.On("UploadImage", mock.Anything).Return("cloudinary_url", nil)
		service := NewProductService(mockDB, mockCloudinary)
		formUpdate := &multipart.Form{
			Value: map[string][]string{
				"UserID":      {"2"},
				"Price":       {"11.5"},
				"Name":        {"Test Product Update"},
				"Description": {"Test Description Update"},
				"Category":    {"Test Category Update"},
				"Ubication":   {"Test Ubication Update"},
			},
		}
		image := &multipart.FileHeader{}
		product, err := service.UpdateProductService(1, formUpdate, image)
		require.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		assert.Equal(t, uint(2), product.UserID)
	})
}