package test

// import (
// 	"auth-microservice/controllers"
// 	"auth-microservice/db"
// 	"auth-microservice/middleware"
// 	"bytes"
// 	"encoding/json"
// 	"io"
// 	"log"
// 	"mime/multipart"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"strings"
// 	"testing"

// 	"github.com/joho/godotenv"
// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// )

// // var token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IktVWUlsME5QWlp1bXVuR2JEcVFpTSJ9.eyJpc3MiOiJodHRwczovL2Rldi1zNThybXZ0anczbnc4eWs0LnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJnb29nbGUtb2F1dGgyfDExMTMxODM4MDM0Mzc3MDg0MzU3MCIsImF1ZCI6WyJodHRwczovL2FwaS1nb2xhbmctdGVzdCIsImh0dHBzOi8vZGV2LXM1OHJtdnRqdzNudzh5azQudXMuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTcxMjM3OTkwNCwiZXhwIjoxNzEyNDY2MzA0LCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIiwiYXpwIjoiUFFTNjNuSjY5N1VpTU01dlBFYlF5Q28yTjJ1WFZBVTMifQ.nTRJ3QznolRs-X4FtWLP3Ten17Km7eAoCA47kQS6axBGSnzSSv4oBFIOAuAlHrEHYiBp8hcHuI3-xmXwExe_PCOmeQzVYDt0LJ85iM7bYYPjBPVCtnFTbR_uxZqQFDWJ3_HITk6BugAfMRavkvTA6YWj12Tu7kzpoKxp6IBO1X4jM5wG-MKVNIItpju2KHfKl9oV8v5yyY2iu-ZEonYwktosA9-RxLs7BrJx8tLi7g9z6B1lcoYvY7sORJWy5djeM0Fu_afwmdAhy98IvBQkK9JM102_vq6z5eTfpUeCc_r1t-X723ADWZUsqE8uLuqMsYOzw7kribf7xFlIpe_OYw"

// var tokenAdminJWT = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IktVWUlsME5QWlp1bXVuR2JEcVFpTSJ9.eyJpc3MiOiJodHRwczovL2Rldi1zNThybXZ0anczbnc4eWs0LnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJnb29nbGUtb2F1dGgyfDExMTMxODM4MDM0Mzc3MDg0MzU3MCIsImF1ZCI6WyJodHRwczovL2FwaS1nb2xhbmctdGVzdCIsImh0dHBzOi8vZGV2LXM1OHJtdnRqdzNudzh5azQudXMuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTcxMjU5ODM0OSwiZXhwIjoxNzEyNjg0NzQ5LCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIiwiYXpwIjoiUFFTNjNuSjY5N1VpTU01dlBFYlF5Q28yTjJ1WFZBVTMifQ.Ex2MEavArK2zBqMAAcZd7D5A1T5G4SDi49rtZtXwOLnviiTTiTPUZHbwVbCHX1r3X9eH-hl_7LFPm-khrFdCMDbJhj72S39UxB-yVJ0TJp8IdOKKdzZh4ay5wVZcBDnxE03R84AXXkx38qWnlQxDOmZKfDSekVvB0NtLH2M5qbcSVjTSIwdJohU39rJ4FFmxX2Bs5UwAT9BST67FPHOSZ5mG5_h-Q8K0bXg07WUBWkWuRxHqcY0eDOjWZkEmqz4gcjWdLVLKukAAg4zZHNTTPPB5Q_YVfFUY6O39cpEajcS30OTvDdCvSw7xYsdbgRgWsR8Y4jXSJPgtWzK-MEIhow"

// func TestMain(t *testing.T) {
// 	if err := godotenv.Load("../.env"); err != nil {
// 		log.Fatalf("Error loading the .env file: %v", err)
// 	}

// 	db.DBConnection()
// 	e := echo.New()

// 	t.Run("Get one user that exist", func(t *testing.T) {

// 		e.GET("/users/:id", controllers.GetUserHandler)
// 		idUser := "google-oauth2|111318380343770843570"
// 		req := httptest.NewRequest(http.MethodGet, "/users/"+idUser, nil)
// 		rec := httptest.NewRecorder()

// 		assert.Equal(t, http.StatusOK, rec.Code)

// 		e.ServeHTTP(rec, req)

// 		var userResponse struct {
// 			IdAuth string `json:"IdAuth"`
// 		}

// 		// Deserializa el cuerpo de la respuesta en la estructura definida
// 		if err := json.Unmarshal(rec.Body.Bytes(), &userResponse); err != nil {
// 			t.Fatalf("Error al deserializar la respuesta del usuario: %v", err)
// 		}

// 		// Comprueba si el campo IdAuth de la respuesta coincide con el esperado
// 		assert.Equal(t, "google-oauth2|111318380343770843570", userResponse.IdAuth)
// 	})

// 	t.Run("Get one user that not exist", func(t *testing.T) {

// 		e.GET("/users/:id", controllers.GetUserHandler)

// 		idUser := "dame_un_beso"
// 		req := httptest.NewRequest(http.MethodGet, "/users/"+idUser, nil)
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		// Crear el grabador de respuesta
// 		rec := httptest.NewRecorder()

// 		// Simular la solicitud HTTP
// 		e.ServeHTTP(rec, req)
// 		// Comprueba el código de estado de la respuesta
// 		assert.Equal(t, http.StatusNotFound, rec.Code)

// 	})

// 	t.Run("Change rol of user with middleware", func(t *testing.T) {
// 		// Crear el cuerpo de la solicitud como un lector de bytes

// 		// Define la ruta y el controlador que deseas probar
// 		e.PUT("/users/role/:id", controllers.ChangeRoleHandler, middleware.RoleMiddleware)

// 		// Crear el cuerpo de la solicitud como un lector de bytes
// 		reqBody := strings.NewReader(`"USER"`)

// 		// Crear la solicitud con el cuerpo adecuado
// 		idUser := "google-oauth2|111318380343770843570"
// 		req := httptest.NewRequest(http.MethodPut, "/users/role/"+idUser, reqBody)
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+tokenAdminJWT)
// 		// Crear el grabador de respuesta
// 		rec := httptest.NewRecorder()

// 		// Simular la solicitud HTTP
// 		e.ServeHTTP(rec, req)
// 		// Comprueba el código de estado de la respuesta
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 	})

// 	t.Run("Update user", func(t *testing.T) {
// 		// Define la ruta y el controlador que deseas probar
// 		e.PUT("/users/:id", controllers.UpdateUserHandler)

// 		body := new(bytes.Buffer)
// 		writer := multipart.NewWriter(body)
// 		writer.WriteField("name", "pepito")
// 		writer.WriteField("email", "soyguapo@gmail.com")
// 		writer.Close()

// 		// Crear la solicitud con el cuerpo adecuado
// 		idUser := "google-oauth2|111318380343770843570"
// 		req := httptest.NewRequest(http.MethodPut, "/users/"+idUser, body)
// 		req.Header.Set("Content-Type", writer.FormDataContentType())

// 		// Crear el grabador de respuesta
// 		rec := httptest.NewRecorder()

// 		// Simular la solicitud HTTP
// 		e.ServeHTTP(rec, req)

// 		var userResponse struct {
// 			Name  string `json:"name"`
// 			Email string `json:"email"`
// 		}

// 		// Deserializa el cuerpo de la respuesta en la estructura definida
// 		if err := json.Unmarshal(rec.Body.Bytes(), &userResponse); err != nil {
// 			t.Fatalf("Error al deserializar la respuesta del usuario: %v", err)
// 		}
// 		assert.Equal(t, "pepito", userResponse.Name)
// 		assert.Equal(t, "soyguapo@gmail.com", userResponse.Email)

// 		// Comprueba el código de estado de la respuesta
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 	})

// 	t.Run("Update user with image", func(t *testing.T) {
// 		// Define la ruta y el controlador que deseas probar
// 		e.PUT("/users/:id", controllers.UpdateUserHandler)

// 		// Crear el cuerpo del formulario multipart
// 		body := new(bytes.Buffer)
// 		writer := multipart.NewWriter(body)
// 		writer.WriteField("name", "arroba")
// 		writer.WriteField("email", "arroba@gmail.com")

// 		part, err := writer.CreateFormFile("image", "florcita.jpg")
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		file, err := os.Open("florcita.jpg")
// 		if err != nil {
// 			t.Fatalf("Error al abrir el archivo de imagen: %v", err)
// 		}
// 		io.Copy(part, file)
// 		writer.Close()

// 		// Crear la solicitud con el cuerpo adecuado
// 		idUser := "google-oauth2|111318380343770843570"
// 		req := httptest.NewRequest(http.MethodPut, "/users/"+idUser, body)
// 		req.Header.Set("Content-Type", writer.FormDataContentType())

// 		// Crear el grabador de respuesta
// 		rec := httptest.NewRecorder()

// 		// Simular la solicitud HTTP
// 		e.ServeHTTP(rec, req)

// 		var userResponse struct {
// 			Name  string `json:"name"`
// 			Email string `json:"email"`
// 		}

// 		// Deserializa el cuerpo de la respuesta en la estructura definida
// 		if err := json.Unmarshal(rec.Body.Bytes(), &userResponse); err != nil {
// 			t.Fatalf("Error al deserializar la respuesta del usuario: %v", err)
// 		}
// 		assert.Equal(t, "pepito", userResponse.Name)
// 		assert.Equal(t, "soyguapo@gmail.com", userResponse.Email)

// 		// Comprueba el código de estado de la respuesta
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 	})

// }
