package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
)

func GenerateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func GenerateResetURL(r *http.Request, resetToken string) string {
	protocol := "http"
	if r.TLS != nil {
		protocol = "https"
	}

	host := r.Host

	resetURL := fmt.Sprintf("%s://%s/api/resetPassword/%s", protocol, host, resetToken)
	return resetURL
}
