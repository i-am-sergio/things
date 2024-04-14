package utils_test

import (
	"auth-microservice/utils"
	"testing"
)

func TestGetIdTokenJWTAuth0(t *testing.T) {
	t.Run("Token válido con campo sub", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dBiq4UXZSKd4EnPifQ0IqcQKY5Pn2BmOMmOZlpJ6cYg"
		expectedSub := "1234567890"

		sub := utils.GetIdTokenJWTAuth0(token)

		if sub != expectedSub {
			t.Errorf("Se esperaba %s pero se obtuvo %s", expectedSub, sub)
		}
	})

	t.Run("Token inválido (malformado)", func(t *testing.T) {
		token := "token_malformado"

		sub := utils.GetIdTokenJWTAuth0(token)

		if sub != "" {
			t.Error("Se esperaba una cadena vacía para un token malformado, pero se obtuvo un valor diferente")
		}
	})

	t.Run("Token válido sin campo sub", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.e30"

		sub := utils.GetIdTokenJWTAuth0(token)

		if sub != "" {
			t.Error("Se esperaba una cadena vacía para un token válido sin campo sub, pero se obtuvo un valor diferente")
		}
	})
}
