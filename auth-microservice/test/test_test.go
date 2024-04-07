package test

import (
	"auth-microservice/controllers"
	"auth-microservice/db"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

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
		c.SetParamValues("1")

		_ = controllers.GetUserHandler(c)

		// Assertions
		assert.Equal(t, http.StatusOK, rec.Code)

		// Define la estructura esperada del usuario
		expectedUser := struct {
			ID        int       `json:"ID"`
			CreatedAt string    `json:"CreatedAt"`
			UpdatedAt string    `json:"UpdatedAt"`
			DeletedAt *struct{} `json:"DeletedAt"`
			Name      string    `json:"Name"`
			Email     string    `json:"Email"`
			Password  string    `json:"Password"`
			Image     string    `json:"Image"`
			Ubication string    `json:"Ubication"`
			Role      string    `json:"Role"`
		}{
			ID:        1,
			CreatedAt: "2024-04-05T21:59:44.395459-05:00",
			UpdatedAt: "2024-04-06T12:25:18.94784-05:00",
			DeletedAt: nil,
			Name:      "pedrito",
			Email:     "ldavis@unsa.edu.www",
			Password:  "rg4l",
			Image:     "href",
			Ubication: "contigo pero sin ti xdddd",
			Role:      "ENTERPRISE",
		}

		// Serializa la estructura esperada a JSON
		expectedUserJSON, err := json.Marshal(expectedUser)
		if err != nil {
			t.Fatalf("Error al serializar el usuario esperado: %v", err)
		}

		// Compara el cuerpo de la respuesta con el JSON esperado
		assert.JSONEq(t, string(expectedUserJSON), rec.Body.String())
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

}
