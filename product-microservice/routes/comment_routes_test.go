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

type MockCommentController struct {
	mock.Mock
	controllers.CommentController
}

func (m *MockCommentController) CreateComment(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockCommentController) GetComments(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockCommentController) GetCommentByID(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockCommentController) GetCommentsByProductID(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockCommentController) UpdateComment(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockCommentController) DeleteComment(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func TestCommentRoutes(t *testing.T) {
	e := echo.New()
	mockController := new(MockCommentController)
	CommentRoutes(e, mockController)
	formData := bytes.NewBufferString("name=TestProduct&description=JustATest")
	req := httptest.NewRequest(http.MethodPost, "/comments", formData)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	mockController.On("CreateComment", mock.Anything).Return(nil)
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockController.AssertExpectations(t)
}