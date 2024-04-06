package test

import (
	"auth-microservice/controllers"
	"auth-microservice/db"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IktVWUlsME5QWlp1bXVuR2JEcVFpTSJ9.eyJpc3MiOiJodHRwczovL2Rldi1zNThybXZ0anczbnc4eWs0LnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJnb29nbGUtb2F1dGgyfDExMTMxODM4MDM0Mzc3MDg0MzU3MCIsImF1ZCI6WyJodHRwczovL2FwaS1nb2xhbmctdGVzdCIsImh0dHBzOi8vZGV2LXM1OHJtdnRqdzNudzh5azQudXMuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTcxMjM3OTkwNCwiZXhwIjoxNzEyNDY2MzA0LCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIiwiYXpwIjoiUFFTNjNuSjY5N1VpTU01dlBFYlF5Q28yTjJ1WFZBVTMifQ.nTRJ3QznolRs-X4FtWLP3Ten17Km7eAoCA47kQS6axBGSnzSSv4oBFIOAuAlHrEHYiBp8hcHuI3-xmXwExe_PCOmeQzVYDt0LJ85iM7bYYPjBPVCtnFTbR_uxZqQFDWJ3_HITk6BugAfMRavkvTA6YWj12Tu7kzpoKxp6IBO1X4jM5wG-MKVNIItpju2KHfKl9oV8v5yyY2iu-ZEonYwktosA9-RxLs7BrJx8tLi7g9z6B1lcoYvY7sORJWy5djeM0Fu_afwmdAhy98IvBQkK9JM102_vq6z5eTfpUeCc_r1t-X723ADWZUsqE8uLuqMsYOzw7kribf7xFlIpe_OYw"

func TestMain(t *testing.T) {

	t.Run("Retorna p ctm", func(t *testing.T) {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatalf("Error loading the .env file: %v", err)
		}

		db.DBConnection()

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		//req.Header.Set("Authorization", "Bearer "+token)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		log.Println(c.Request().URL)

		h := controllers.GetUserHandler(c)

		// Assertions
		if assert.NoError(t, h) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}

	})
}
