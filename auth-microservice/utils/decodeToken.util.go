package utils

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type TokenData struct {
	Sub string `json:"sub"`
}

func GetIdTokenJWTAuth0(token string) string {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return ""
	}
	payload, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return ""
	}
	var tokenData TokenData
	if err := json.Unmarshal(payload, &tokenData); err != nil {
		return ""
	}
	return tokenData.Sub
}
