package test

import (
	"auth-microservice/controllers"
	"auth-microservice/db"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IktVWUlsME5QWlp1bXVuR2JEcVFpTSJ9.eyJpc3MiOiJodHRwczovL2Rldi1zNThybXZ0anczbnc4eWs0LnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJnb29nbGUtb2F1dGgyfDExMTMxODM4MDM0Mzc3MDg0MzU3MCIsImF1ZCI6WyJodHRwczovL2FwaS1nb2xhbmctdGVzdCIsImh0dHBzOi8vZGV2LXM1OHJtdnRqdzNudzh5azQudXMuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTcxMjM3OTkwNCwiZXhwIjoxNzEyNDY2MzA0LCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIiwiYXpwIjoiUFFTNjNuSjY5N1VpTU01dlBFYlF5Q28yTjJ1WFZBVTMifQ.nTRJ3QznolRs-X4FtWLP3Ten17Km7eAoCA47kQS6axBGSnzSSv4oBFIOAuAlHrEHYiBp8hcHuI3-xmXwExe_PCOmeQzVYDt0LJ85iM7bYYPjBPVCtnFTbR_uxZqQFDWJ3_HITk6BugAfMRavkvTA6YWj12Tu7kzpoKxp6IBO1X4jM5wG-MKVNIItpju2KHfKl9oV8v5yyY2iu-ZEonYwktosA9-RxLs7BrJx8tLi7g9z6B1lcoYvY7sORJWy5djeM0Fu_afwmdAhy98IvBQkK9JM102_vq6z5eTfpUeCc_r1t-X723ADWZUsqE8uLuqMsYOzw7kribf7xFlIpe_OYw"
var tokenTest = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IktVWUlsME5QWlp1bXVuR2JEcVFpTSJ9.eyJpc3MiOiJodHRwczovL2Rldi1zNThybXZ0anczbnc4eWs0LnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJnb29nbGUtb2F1dGgyfDExMzU0OTIzNjYyMzM2MjQ1NTAyNSIsImF1ZCI6WyJodHRwczovL2FwaS1nb2xhbmctdGVzdCIsImh0dHBzOi8vZGV2LXM1OHJtdnRqdzNudzh5azQudXMuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTcxMjQ2NjU0MCwiZXhwIjoxNzEyNTUyOTQwLCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIiwiYXpwIjoiUFFTNjNuSjY5N1VpTU01dlBFYlF5Q28yTjJ1WFZBVTMifQ.DOt3R9y5sLRoI5vUerZXutvwR-97of06YM7lIlBPUq6mVe6gthdAilqXQKQu5eESSq-X5msMARzll0nz-zaEff2SaauDI4eZqzbAu0x_O1S1k3RkivNG-u4v7kn5rcx50DWeZWwxIkhzdrVPX1aD5iG8wrHP9ewo7pnM5mVsIsdU2CfpF7JHDu-AJecnSuvS8oNXhCY7ALtIJpoIS54TSOALNIGJ1EUrGC-K2qUtsDAGz7gSghJJgpxwiKdiT-QKZhlCSeX_3I4QNPK88BPHMmQ-VXECcCc0XNekC8hDPYN3dQuV81I6asya-6w_lyljDIpMxu46lXDpEFw2bVs41A App.tsx:16:16"

func TestMain(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	db.DBConnection()
	e := echo.New()

	t.Run("Get one user that exist", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("google-oauth2|111318380343770843570")

		_ = controllers.GetUserHandler(c)

		// Assertions
		assert.Equal(t, http.StatusOK, rec.Code)

		// Define la estructura esperada del usuario
		var userResponse struct {
			IdAuth string `json:"IdAuth"`
		}

		// Deserializa el cuerpo de la respuesta en la estructura definida
		if err := json.Unmarshal(rec.Body.Bytes(), &userResponse); err != nil {
			t.Fatalf("Error al deserializar la respuesta del usuario: %v", err)
		}

		// Comprueba si el campo IdAuth de la respuesta coincide con el esperado
		assert.Equal(t, "google-oauth2|111318380343770843570", userResponse.IdAuth)
	})

	t.Run("Get one user that not exist", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("-10")

		_ = controllers.GetUserHandler(c)

		// Assertions
		assert.Equal(t, http.StatusNotFound, rec.Code)

	})

	t.Run("Change rol of user", func(t *testing.T) {
		// Crear el cuerpo de la solicitud como un lector de bytes
		reqBody := bytes.NewReader([]byte("ADMIN"))

		// Crear la solicitud con el cuerpo adecuado
		req := httptest.NewRequest(http.MethodPut, "/", reqBody)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/role/:id")
		c.SetParamNames("id")
		c.SetParamValues(tokenTest)

		_ = controllers.ChangeRoleHandler(c)

		// Assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})

}
