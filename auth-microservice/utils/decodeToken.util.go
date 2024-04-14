package utils

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

// Define una estructura para almacenar los datos del token decodificado
type TokenData struct {
	Sub string `json:"sub"`
}

func GetIdTokenJWTAuth0(token string) string {
	// Dividir el token en sus partes: encabezado, carga útil y firma
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return ""
	}

	// Decodificar la carga útil del token
	payload, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return ""
	}

	// Analizar el JSON decodificado en una estructura TokenData
	var tokenData TokenData
	if err := json.Unmarshal(payload, &tokenData); err != nil {
		return ""
	}

	// Retornar el campo "sub" del token decodificado
	return tokenData.Sub
}
