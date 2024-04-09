package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func CreateCommentTest(t *testing.T) {
	t.Run("Should return status 200", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/comments", strings.NewReader(`{"comment": "This is a comment"}`))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		controller := Controller{}
		controller.CreateComment(c)
		assert.Equal(t, http.StatusCreated, rec.Code)
	})
}