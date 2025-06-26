package tests

import (
	"testing"
	"time"

	"pki.gov.kz/go/kalkan/jwt"
)

func TestNewToken(t *testing.T) {
	issuedAt := time.Now().Local().Unix() - 120
	expiresAt := time.Now().Local().Unix() + 900
	var claims = jwt.Claims{
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
		BIN:       "161140016747",
		MCheck:    jwt.OTP,
	}
	token := jwt.New(claims, mod)
	err := token.Sign()
	if err != nil {
		t.Fatal(err)
	}
}
