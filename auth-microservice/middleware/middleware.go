package middleware

import (
	"auth-microservice/services"
	"auth-microservice/utils"
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func EnsureValidToken() func(next http.Handler) http.Handler {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return func(next http.Handler) http.Handler {
		return middleware.CheckJWT(next)
	}
}

// HasScope checks whether our claims have a specific scope.
func (c CustomClaims) HasScope(expectedScope string) bool {
	result := strings.Split(c.Scope, " ")
	for i := range result {
		if result[i] == expectedScope {
			return true
		}
	}

	return false
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Aplica el middleware de validación de JWT
		handler := EnsureValidToken()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Set("user", c.Get("user"))
			// No es necesario llamar next(c) aquí
		}))
		handler.ServeHTTP(c.Response(), c.Request())

		// Si hay un error durante la validación del token JWT, retorna un error
		if c.Response().Status == http.StatusUnauthorized {
			return echo.NewHTTPError(http.StatusUnauthorized, "No se proporcionó un token de autorización válido")
		}

		// Si no hay errores, llama al siguiente controlador en la cadena
		return next(c)
	}
}

func RoleMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Obtener el encabezado de autorización de la solicitud
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			// Si no se proporciona un token de autorización, retorna un error
			return echo.NewHTTPError(http.StatusUnauthorized, "Token de autorización faltante")
		}

		// Dividir el encabezado de autorización para obtener el token JWT
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			// Si el formato del encabezado de autorización no es válido, retorna un error
			return echo.NewHTTPError(http.StatusBadRequest, "Formato de token de autorización inválido")
		}

		// Obtener el token JWT del encabezado de autorización
		token := authParts[1]

		// Obtener el sub (subject) del token JWT
		sub := utils.GetIdTokenJWTAuth0(token)

		// Obtener el usuario desde la base de datos utilizando el sub (subject)
		user, err := services.GetUserByIdAuth(sub)
		if err != nil {
			return err
		}

		// Verificar el rol del usuario
		if user.Role != "ADMIN" {
			// Si el usuario no tiene el rol de "ADMIN", retornar un error 403 (Forbidden)
			return echo.NewHTTPError(http.StatusForbidden, "No tienes permiso para acceder a este recurso")
		}

		// Si el usuario tiene el rol de "ADMIN", continuar con la siguiente función de middleware o controlador en la cadena
		return next(c)
	}
}
