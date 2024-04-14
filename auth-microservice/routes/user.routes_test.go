package routes_test

import (
	"auth-microservice/controllers"
	"auth-microservice/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var tokenJWT = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IktVWUlsME5QWlp1bXVuR2JEcVFpTSJ9.eyJpc3MiOiJodHRwczovL2Rldi1zNThybXZ0anczbnc4eWs0LnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJnb29nbGUtb2F1dGgyfDExMzU0OTIzNjYyMzM2MjQ1NTAyNSIsImF1ZCI6WyJodHRwczovL2FwaS1nb2xhbmctdGVzdCIsImh0dHBzOi8vZGV2LXM1OHJtdnRqdzNudzh5azQudXMuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTcxMjg5NzgyMSwiZXhwIjoxNzEyOTg0MjIxLCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIiwiYXpwIjoiUFFTNjNuSjY5N1VpTU01dlBFYlF5Q28yTjJ1WFZBVTMifQ.LlJYUo_M0HRYBqBiECcIRiqtPFb9QsQauAmh8RTKaDEwXl4t2Yh9XeuPmqSuKT2cYvEV8YOPb6PLNrGA5JhnHTzt-nHY3Srn1EWhP96LefnNnMC9QEf-smmhVRbD-EeXm_yugQ_a3b2rIE6_BI229gw4ZRdJ7ewxraiuuwKfzfAuzi-BtbyxhFZ7QXF4UizZ84u3DDTP8yuk3nv6xUMMXEt1PKTNbegmKYvT_5Z7AQAoljcyHrjyqYWGGkcYIvKSjF6IAsp07qMy11-d74sCfyj2KzwK6_4XYcgkkebIFwAclNZiWtGxuwa0XJF4Z2HD60Ha9cuBFnXYWm72nqqfMA"

func TestGetAllUsers(t *testing.T) {
	e := echo.New()
	mockcontroller := &controllers.UserController{}
	routes.UsersRoutes(e, mockcontroller)
	rec := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/users", nil)

	e.ServeHTTP(rec, req)
	assert.Nil(t, err)
}

func TestGetOneUsers(t *testing.T) {
	e := echo.New()
	mockcontroller := &controllers.UserController{}
	routes.UsersRoutes(e, mockcontroller)
	rec := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/users/1", nil)
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.Nil(t, err)
}

func TestPost(t *testing.T) {
	e := echo.New()
	mockcontroller := &controllers.UserController{}
	routes.UsersRoutes(e, mockcontroller)
	respRec := httptest.NewRecorder()

	body := strings.NewReader(`{"name": "John", "email": "john@example.com", "password": "secret"}`)

	req, _ := http.NewRequest(http.MethodPost, "/users/", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+tokenJWT)

	e.ServeHTTP(respRec, req)

	assert.Equal(t, http.StatusUnauthorized, respRec.Code)
}
