package main

import (
	"context"
	"errors"
	"net/http"
	"product-microservice/db"
	"product-microservice/repository"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockRealDBInit struct {
	mock.Mock
}

func (m *MockRealDBInit) Init(loader db.EnvLoader, connector db.DBConnector) (repository.DBInterface, error) {
	args := m.Called(loader, connector)
	return args.Get(0).(repository.DBInterface), args.Error(1)
}

type MockCloudinary struct {
	mock.Mock
}

func (m *MockCloudinary) InitCloudinary(loader db.EnvLoader) error {
	args := m.Called(loader)
	return args.Error(0)
}

func (m *MockCloudinary) UploadImage(imagePath db.FileHeaderWrapper) (string, error) {
	args := m.Called(imagePath)
	return args.String(0), args.Error(1)
}

type MockDBInterface struct {
	mock.Mock
}

func (m *MockDBInterface) AutoMigrate(dst ...interface{}) error {
	args := m.Called(dst...)
	return args.Error(0)
}

func (m *MockDBInterface) First(dest interface{}, conds ...interface{}) error {
	args := m.Called(dest, conds)
	return args.Error(0)
}

func (m *MockDBInterface) Save(value interface{}) error {
	args := m.Called(value)
	return args.Error(0)
}

func (m *MockDBInterface) Create(value interface{}) error {
	args := m.Called(value)
	return args.Error(0)
}

func (m *MockDBInterface) FindPreloaded(relation string, dest interface{}, conds ...interface{}) error {
	args := m.Called(relation, dest, conds)
	return args.Error(0)
}

func (m *MockDBInterface) Find(dest interface{}, conds ...interface{}) error {
	args := m.Called(dest, conds)
	return args.Error(0)
}

func (m *MockDBInterface) FindWithCondition(dest interface{}, query string, arg ...interface{}) error {
	args := m.Called(dest, query, arg)
	return args.Error(0)
}

func (m *MockDBInterface) Delete(value interface{}) error {
	args := m.Called(value)
	return args.Error(0)
}

func (m *MockDBInterface) DeleteWithCondition(model interface{}, query string, arg ...interface{}) error {
	args := m.Called(model, query, arg)
	return args.Error(0)
}

func (m *MockDBInterface) DeleteByID(model interface{}, id interface{}) error {
	args := m.Called(model, id)
	return args.Error(0)
}

func TestAppInitialize(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockDB := new(MockRealDBInit)
		mockDBInterface := new(MockDBInterface)
		mockCloudinary := new(MockCloudinary)
		e := echo.New()
		app := &App{
			DB:          mockDB,
			Cloudinary:  mockCloudinary,
			HTTPHandler: e,
		}
		mockDB.On("Init", mock.AnythingOfType("*db.DotEnvLoader"), mock.AnythingOfType("*db.GormConnector")).Return(mockDBInterface, nil)
		mockDBInterface.On("AutoMigrate", mock.AnythingOfType("*models.Product"), mock.AnythingOfType("*models.Comment")).Return(nil)
		mockCloudinary.On("InitCloudinary", mock.AnythingOfType("*db.DotEnvLoader")).Return(nil)
		err := app.Initialize()
		require.NoError(t, err)
		mockDB.AssertExpectations(t)
		mockDBInterface.AssertExpectations(t)
		mockCloudinary.AssertExpectations(t)
	})
	t.Run("ErrorDBInit", func(t *testing.T) {
		mockDB := new(MockRealDBInit)
		mockDBInterface := new(MockDBInterface)
		mockCloudinary := new(MockCloudinary)
		e := echo.New()
		app := &App{
			DB:          mockDB,
			Cloudinary:  mockCloudinary,
			HTTPHandler: e,
		}
		mockDB.On("Init", mock.AnythingOfType("*db.DotEnvLoader"), mock.AnythingOfType("*db.GormConnector")).Return(mockDBInterface, errors.New("error al inicializar la base de datos"))
		err := app.Initialize()
		require.Error(t, err)
		mockDB.AssertExpectations(t)
	})
	t.Run("ErrorAutoMigrate", func(t *testing.T) {
		mockDB := new(MockRealDBInit)
		mockDBInterface := new(MockDBInterface)
		mockCloudinary := new(MockCloudinary)
		e := echo.New()
		app := &App{
			DB:          mockDB,
			Cloudinary:  mockCloudinary,
			HTTPHandler: e,
		}
		mockDB.On("Init", mock.AnythingOfType("*db.DotEnvLoader"), mock.AnythingOfType("*db.GormConnector")).Return(mockDBInterface, nil)
		mockDBInterface.On("AutoMigrate", mock.AnythingOfType("*models.Product"), mock.AnythingOfType("*models.Comment")).Return(errors.New("error al hacer AutoMigrate"))
		err := app.Initialize()
		require.Error(t, err)
		mockDB.AssertExpectations(t)
		mockDBInterface.AssertExpectations(t)
	})
	t.Run("ErrorCloudinaryInit", func(t *testing.T) {
		mockDB := new(MockRealDBInit)
		mockDBInterface := new(MockDBInterface)
		mockCloudinary := new(MockCloudinary)
		e := echo.New()
		app := &App{
			DB:          mockDB,
			Cloudinary:  mockCloudinary,
			HTTPHandler: e,
		}
		mockDB.On("Init", mock.AnythingOfType("*db.DotEnvLoader"), mock.AnythingOfType("*db.GormConnector")).Return(mockDBInterface, nil)
		mockDBInterface.On("AutoMigrate", mock.AnythingOfType("*models.Product"), mock.AnythingOfType("*models.Comment")).Return(nil)
		mockCloudinary.On("InitCloudinary", mock.AnythingOfType("*db.DotEnvLoader")).Return(errors.New("error al inicializar cloudinary"))
		err := app.Initialize()
		require.Error(t, err)
		mockDB.AssertExpectations(t)
		mockDBInterface.AssertExpectations(t)
		mockCloudinary.AssertExpectations(t)
	})
}
func TestAppRun(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		e := echo.New()
		app := &App{
			HTTPHandler: e,
		}
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		})
		go func() {
			if err := app.Run(":9999"); err != nil {
				t.Log("Server failed to start:", err)
			}
		}()
		time.Sleep(time.Second)
		resp, err := http.Get("http://localhost:9999/")
		if err != nil {
			t.Fatalf("Failed to make a GET request to the running server: %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected HTTP 200 OK response, got %d", resp.StatusCode)
		}
		if err := e.Shutdown(context.Background()); err != nil {
			t.Fatal("Failed to shut down the server properly:", err)
		}
	})
	t.Run("Error", func(t *testing.T) {
		e := echo.New()
		app := &App{
			HTTPHandler: e,
		}
		err := app.Run(":-1")
		if err == nil {
			t.Fatal("Expected an error when trying to run the server, but got nil")
		}
	})
}