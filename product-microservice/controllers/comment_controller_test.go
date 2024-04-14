package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"product-microservice/models"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockCommentService struct {
	mock.Mock
}

func (m *MockCommentService) UpdateProductRating(productID uint) error {
	args := m.Called(productID)
	return args.Error(0)
}

func (m *MockCommentService) CreateCommentService(comment models.Comment) error {
	args := m.Called(comment)
	return args.Error(0)
}

func (m *MockCommentService) GetCommentsService() ([]models.Comment, error) {
	args := m.Called()
	return args.Get(0).([]models.Comment), args.Error(1)
}

func (m *MockCommentService) GetCommentByIDService(id uint) (models.Comment, error) {
	args := m.Called(id)
	return args.Get(0).(models.Comment), args.Error(1)
}

func (m *MockCommentService) GetCommentsByProductIDService(id uint) ([]models.Comment, error) {
	args := m.Called(id)
	return args.Get(0).([]models.Comment), args.Error(1)
}

func (m *MockCommentService) UpdateCommentService(comment models.Comment, id uint) error {
	args := m.Called(comment, id)
	return args.Error(0)
}

func (m *MockCommentService) DeleteCommentService(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateComment(t *testing.T){
	t.Run("Success", func(t *testing.T){
		comment := models.Comment{
			ID:        1,
			ProductID: 2,
			Rating:    5,
		}
		jsonBody, _ := json.Marshal(comment)
		req := httptest.NewRequest(http.MethodPost, "/comments", bytes.NewBuffer(jsonBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		mockCommentService := new(MockCommentService)
		mockCommentService.On("CreateCommentService", mock.AnythingOfType("models.Comment")).Return(nil)
		controller := NewCommentController(mockCommentService)
		require.NoError(t, controller.CreateComment(c))
		assert.Equal(t, http.StatusCreated, rec.Code)
		var returnedComment models.Comment
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &returnedComment))
		assert.Equal(t, comment.ID, returnedComment.ID)
		mockCommentService.AssertExpectations(t)
	})
	t.Run("Error creating comment", func(t *testing.T) {
		comment := models.Comment{
			ID:        1,
			ProductID: 2,
			Rating:    5,
		}
		jsonBody, _ := json.Marshal(comment)
		req := httptest.NewRequest(http.MethodPost, "/comments", bytes.NewBuffer(jsonBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		mockCommentService := new(MockCommentService)
		mockCommentService.On("CreateCommentService", mock.AnythingOfType("models.Comment")).Return(assert.AnError) // Simulate an error from the service
		controller := NewCommentController(mockCommentService)
		
		require.NoError(t, controller.CreateComment(c))
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		responseMap := make(map[string]string)
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &responseMap))
		assert.Contains(t, responseMap["error"], "assert.AnError general error for testing")
		mockCommentService.AssertExpectations(t)
	})
	t.Run("Error binding comment", func(t *testing.T) {
		invalidJSON := []byte(`{"ID": "abc", "ProductID": "xyz", "Rating": "not a number"}`)
		req := httptest.NewRequest(http.MethodPost, "/comments", bytes.NewBuffer(invalidJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		mockCommentService := new(MockCommentService)
		controller := NewCommentController(mockCommentService)
		controller.CreateComment(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		responseMap := make(map[string]string)
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &responseMap))
		assert.Contains(t, responseMap["error"], "Unmarshal type error")
		assert.Contains(t, responseMap["error"], "cannot unmarshal string into Go struct field Comment.ID of type uint")
		mockCommentService.AssertNotCalled(t, "CreateCommentService")
	})	
}

func TestGetComments(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		comments := []models.Comment{
			{
				ID:        1,
				ProductID: 2,
				Rating:    5,
			},
			{
				ID:        2,
				ProductID: 2,
				Rating:    4,
			},
		}
		req := httptest.NewRequest(http.MethodGet, "/comments", nil)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		mockCommentService := new(MockCommentService)
		mockCommentService.On("GetCommentsService").Return(comments, nil)
		controller := NewCommentController(mockCommentService)
		require.NoError(t, controller.GetComments(c))
		require.Equal(t, http.StatusOK, rec.Code)
		var returnedComments []models.Comment
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &returnedComments))
		require.Equal(t, comments, returnedComments)
		mockCommentService.AssertExpectations(t)
	})
	t.Run("Error getting comments", func(t *testing.T) {
        req := httptest.NewRequest(http.MethodGet, "/comments", nil)
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		var emptyComments []models.Comment
		mockCommentService := new(MockCommentService)
		mockCommentService.On("GetCommentsService").Return(emptyComments, assert.AnError) // Simulate an error from the service
		controller := NewCommentController(mockCommentService)
		controller.GetComments(c)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
		responseMap := make(map[string]string)
		err := json.Unmarshal(rec.Body.Bytes(), &responseMap)
		require.NoError(t, err)
		require.Contains(t, responseMap["error"], "assert.AnError")
		mockCommentService.AssertExpectations(t)
    })
}

func TestGetCommentByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/comments/1", nil)
		rec := httptest.NewRecorder()
		e.GET("/comments/:id", func(c echo.Context) error { return nil })
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")
		mockCommentService := new(MockCommentService)
		expectedComment := models.Comment{
			ID:        1,
			ProductID: 2,
			Rating:    5,
		}
		mockCommentService.On("GetCommentByIDService", uint(1)).Return(expectedComment, nil)
		controller := NewCommentController(mockCommentService)
		err := controller.GetCommentByID(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var returnedComment models.Comment
		err = json.Unmarshal(rec.Body.Bytes(), &returnedComment)
		require.NoError(t, err)
		assert.Equal(t, expectedComment, returnedComment)
		mockCommentService.AssertExpectations(t)
	})
	t.Run("Error parsing ID", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/comments/invalid-id", nil)
		rec := httptest.NewRecorder()
		e.GET("/comments/:id", func(c echo.Context) error { return nil })
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("invalid-id")
		controller := NewCommentController(nil)
		err := controller.GetCommentByID(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		responseMap := make(map[string]string)
		err = json.Unmarshal(rec.Body.Bytes(), &responseMap)
		require.NoError(t, err)
		assert.Contains(t, responseMap["error"], invalidCommentIDError)
	})
	t.Run("Error from service when fetching comment", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/comments/1", nil)
		rec := httptest.NewRecorder()
		e.GET("/comments/:id", func(c echo.Context) error { return nil })
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")
		mockCommentService := new(MockCommentService)
		mockCommentService.On("GetCommentByIDService", uint(1)).Return(models.Comment{}, errors.New("database error"))
		controller := NewCommentController(mockCommentService)
		err := controller.GetCommentByID(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		responseMap := make(map[string]string)
		err = json.Unmarshal(rec.Body.Bytes(), &responseMap)
		require.NoError(t, err)
		assert.Contains(t, responseMap["error"], "database error")
		mockCommentService.AssertExpectations(t)
	})
}

func TestGetCommentsByProductID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/comments/1", nil)
		rec := httptest.NewRecorder()
		e.GET("/comments/:id", func(c echo.Context) error { return nil })
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")
		mockCommentService := new(MockCommentService)
		expectedComments := []models.Comment{
			{
				ID:        1,
				ProductID: 1,
				Rating:    5,
			},
			{
				ID:        2,
				ProductID: 1,
				Rating:    4,
			},
		}
		mockCommentService.On("GetCommentsByProductIDService", uint(1)).Return(expectedComments, nil)
		controller := NewCommentController(mockCommentService)
		err := controller.GetCommentsByProductID(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var returnedComments []models.Comment
		err = json.Unmarshal(rec.Body.Bytes(), &returnedComments)
		require.NoError(t, err)
		assert.Equal(t, expectedComments, returnedComments)
		mockCommentService.AssertExpectations(t)
	})
	t.Run("Error parsing ID", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/comments/invalid-id", nil)
		rec := httptest.NewRecorder()
		e.GET("/comments/:id", func(c echo.Context) error { return nil })
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("invalid-id")
		controller := NewCommentController(nil)
		controller.GetCommentsByProductID(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		responseMap := make(map[string]string)
		if err := json.Unmarshal(rec.Body.Bytes(), &responseMap); assert.NoError(t, err) {
			assert.Contains(t, responseMap["error"], invalidCommentIDError)
		}
	})
	t.Run("Error from service when fetching comments", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/comments/1", nil)
		rec := httptest.NewRecorder()
		e.GET("/comments/:id", func(c echo.Context) error { return nil })
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")
		mockCommentService := new(MockCommentService)
		mockCommentService.On("GetCommentsByProductIDService", uint(1)).Return([]models.Comment{}, errors.New("database error"))
		controller := NewCommentController(mockCommentService)
		controller.GetCommentsByProductID(c)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
		responseMap := make(map[string]string)
		err := json.Unmarshal(rec.Body.Bytes(), &responseMap)
		require.NoError(t, err)
		require.Contains(t, responseMap["error"], "database error")
		mockCommentService.AssertExpectations(t)
	})
}

func TestUpdateComment(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/comments/1", strings.NewReader(`{"ProductID":2,"Rating":5}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/comments/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		updatedComment := models.Comment{ID: 1, ProductID: 2, Rating: 5}
		mockCommentService := new(MockCommentService)
		mockCommentService.On("GetCommentByIDService", uint(1)).Return(models.Comment{ID: 1, ProductID: 2, Rating: 4}, nil)
		mockCommentService.On("UpdateCommentService", updatedComment, uint(2)).Return(nil)
		controller := NewCommentController(mockCommentService)
		controller.UpdateComment(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		var response models.Comment
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, updatedComment, response)
		mockCommentService.AssertExpectations(t)
	})
	t.Run("Error parsing ID", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/comments/invalid", strings.NewReader(`{"ProductID":2,"Rating":5}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/comments/:id")
		c.SetParamNames("id")
		c.SetParamValues("invalid")
		controller := NewCommentController(nil)
		controller.UpdateComment(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		responseMap := make(map[string]string)
		json.Unmarshal(rec.Body.Bytes(), &responseMap)
		assert.Contains(t, responseMap["error"], invalidCommentIDError)
	})
	t.Run("Error fetching comment", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/comments/1", strings.NewReader(`{"ProductID":2,"Rating":5}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/comments/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		mockCommentService := new(MockCommentService)
		mockCommentService.On("GetCommentByIDService", uint(1)).Return(models.Comment{}, errors.New("not found"))
		controller := NewCommentController(mockCommentService)
		controller.UpdateComment(c)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		responseMap := make(map[string]string)
		json.Unmarshal(rec.Body.Bytes(), &responseMap)
		assert.Contains(t, responseMap["error"], "not found")
		mockCommentService.AssertExpectations(t)
	})
	t.Run("Error updating comment", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/comments/1", strings.NewReader(`{"ProductID":2,"Rating":5}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/comments/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		existingComment := models.Comment{ID: 1, ProductID: 2, Rating: 4}
		mockCommentService := new(MockCommentService)
		mockCommentService.On("GetCommentByIDService", uint(1)).Return(existingComment, nil)
		mockCommentService.On("UpdateCommentService", mock.Anything, uint(2)).Return(errors.New("update error"))
		controller := NewCommentController(mockCommentService)
		controller.UpdateComment(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		responseMap := make(map[string]string)
		json.Unmarshal(rec.Body.Bytes(), &responseMap)
		assert.Contains(t, responseMap["error"], "update error")
		mockCommentService.AssertExpectations(t)
	})
	t.Run("Error during binding", func(t *testing.T) {
		e := echo.New()
		requestBody := `{"ProductID": "2", "Rating": "five"}`
		req := httptest.NewRequest(http.MethodPut, "/comments/1", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/comments/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		mockCommentService := new(MockCommentService)
		mockCommentService.On("GetCommentByIDService", uint(1)).Return(models.Comment{ID: 1, ProductID: 2, Rating: 4}, nil)
		controller := NewCommentController(mockCommentService)
		controller.UpdateComment(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		responseMap := make(map[string]string)
		json.Unmarshal(rec.Body.Bytes(), &responseMap)
		assert.Contains(t, responseMap["error"], "cannot unmarshal string into Go struct field Comment.rating of type float64")
		mockCommentService.AssertExpectations(t)
	})
}

func TestDeleteComment(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/comments/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/comments/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		mockCommentService := new(MockCommentService)
		mockCommentService.On("DeleteCommentService", uint(1)).Return(nil)
		controller := NewCommentController(mockCommentService)
		err := controller.DeleteComment(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var responseMap map[string]string
		json.Unmarshal(rec.Body.Bytes(), &responseMap)
		assert.Equal(t, "Comment deleted", responseMap["message"])
		mockCommentService.AssertExpectations(t)
	})
	t.Run("Error parsing ID", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/comments/invalid-id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/comments/:id")
		c.SetParamNames("id")
		c.SetParamValues("invalid-id")
		controller := NewCommentController(nil)
		controller.DeleteComment(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		responseMap := make(map[string]string)
		json.Unmarshal(rec.Body.Bytes(), &responseMap)
		assert.Contains(t, responseMap["error"], invalidCommentIDError)
	})
	t.Run("Error deleting comment", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/comments/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/comments/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		mockCommentService := new(MockCommentService)
		mockCommentService.On("DeleteCommentService", uint(1)).Return(errors.New("database error"))
		controller := NewCommentController(mockCommentService)
		err := controller.DeleteComment(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		var responseMap map[string]string
		json.Unmarshal(rec.Body.Bytes(), &responseMap)
		assert.Contains(t, responseMap["error"], "database error")
		mockCommentService.AssertExpectations(t)
	})
}