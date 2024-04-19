package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"auth-microservice/middleware" // Asegúrate de ajustar la importación según la estructura de tu proyecto
)

func TestJWTMiddleware_ValidToken(t *testing.T) {
	// Configurar las variables de entorno necesarias para el test
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatalf("Error loading the .env file: %v", err)
	}

	// Preparar un servidor de prueba de Echo
	e := echo.New()

	// Mockear el contexto de Echo con un token JWT válido
	validToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IktVWUlsME5QWlp1bXVuR2JEcVFpTSJ9.eyJpc3MiOiJodHRwczovL2Rldi1zNThybXZ0anczbnc4eWs0LnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJnb29nbGUtb2F1dGgyfDExMzU0OTIzNjYyMzM2MjQ1NTAyNSIsImF1ZCI6WyJodHRwczovL2FwaS1nb2xhbmctdGVzdCIsImh0dHBzOi8vZGV2LXM1OHJtdnRqdzNudzh5azQudXMuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTcxMzExMTAwMiwiZXhwIjoxNzEzMTk3NDAyLCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIiwiYXpwIjoiUFFTNjNuSjY5N1VpTU01dlBFYlF5Q28yTjJ1WFZBVTMifQ.N7mhTk0TGUA67ZDh3bXY_42aTCI0AIL69b5pa6wKgjmo_N7WXtG2iKWdldGXB80iIRgahXY_mlSyoXwaSNTm9dymv16NT_TEQyzWltRtNIHfi2zZBAmuX3VV9XAVH_TnhKsN-MnWpiR8DpRZO0nc6we46sFnN-i92BsMlG7aR5rakTfb_3X4G2EvFTpDGlAXD1E01ClCy38I6ZWIVCFVeZI-Ut8lRcj9ccHrM4q6MPKHS2ph2BYrZEcIRkBHqeJB8vow1fJy2b5gLJ1GHG7CmrqJeAKh5i5YLD6TcRTItWDAfIhmAeZjHzLdVSxHYLs6sQrCUv1spJDg4OaIHJzTRg"
	reqValidToken := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidToken.Header.Set(echo.HeaderAuthorization, "Bearer "+validToken)
	recValidToken := httptest.NewRecorder()
	cValidToken := e.NewContext(reqValidToken, recValidToken)

	// Ejecutar el middleware
	handler := middleware.JWTMiddleware(func(c echo.Context) error {
		return c.String(http.StatusOK, "Token válido")
	})
	err := handler(cValidToken)

	// Verificar que la solicitud fue manejada correctamente
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, recValidToken.Code)
	assert.Equal(t, "Token válido", recValidToken.Body.String())
}

func TestJWTMiddleware_InvalidToken(t *testing.T) {
	// Configurar las variables de entorno necesarias para el test
	// Preparar un servidor de prueba de Echo
	e := echo.New()

	// Mockear el contexto de Echo con un token JWT válido
	validToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCasdasdIsImtpZCI6IktVWUlsME5QWlp1bXVuR2JEcVFpTSJ9.eyJpc3MiOiJodHRwczovL2Rldi1zNThybXZ0anczbnc4eWs0LnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJnb29nbGUtb2F1dGgyfDExMTMxODM4MDM0Mzc3MDg0MzU3MCIsImF1ZCI6WyJodHRwczovL2FwaS1nb2xhbmctdGVzdCIsImh0dHBzOi8vZGV2LXM1OHJtdnRqdzNudzh5azQudXMuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTcxMjk4NjIyMSwiZXhwIjoxNzEzMDcyNjIxLCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIiwiYXpwIjoiUFFTNjNuSjY5N1VpTU01dlBFYlF5Q28yTjJ1WFZBVTMifQ.YlS0bYweI1m3ywmoISPDWa3JRRh-Cj9hytRZWYfFYkZvhsGUW44h1QZCmC4MHkzv3c8C1h3UC7LTK84I8MC51GbqwL6SoOdeqJzxR81NPi83M-NYh4CjK62Ws49z1kHfX5wLN6piiSebG82Ru4oCR6Kad9dBSnJDJcdRQnHpQaPqJzG6_ZM2IXWqiYb61FJT81YGvb8lqq5EeAl39BY046lGN6o3qq1Cr0CI2sCs-uBcwncG6uhizbke4ZCdCmUGYwWcrlbFyxqKpwbzL74ZXluR4Ykq8X5t9TnBVjsVhhuTdPaYEUcTbcyplV1NZItgb-D-c_CFtFSQcQuVTHnbhg"
	reqValidToken := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidToken.Header.Set(echo.HeaderAuthorization, "Bearer "+validToken)
	recValidToken := httptest.NewRecorder()
	cValidToken := e.NewContext(reqValidToken, recValidToken)

	// Ejecutar el middleware
	handler := middleware.JWTMiddleware(func(c echo.Context) error {
		return c.String(http.StatusUnauthorized, "{\"message\":\"Failed to validate JWT.\"}")
	})
	err := handler(cValidToken)

	// Verificar que la solicitud fue manejada correctamente
	assert.Error(t, err)
	assert.Equal(t, http.StatusUnauthorized, recValidToken.Code)
}
