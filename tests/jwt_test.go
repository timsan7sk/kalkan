package tests

import (
	"testing"
	"time"

	"pki.gov.kz/go/kalkan/jwt"
)

func TestNewToken(t *testing.T) {
	issuedAt := time.Now().Add(-120 * time.Second)
	expiresAt := issuedAt.Add(900 * time.Second)
	var claims = jwt.Claims{
		IssuedAt:  &issuedAt,
		ExpiresAt: &expiresAt,
		BIN:       "000000000000",
		Check:     "OTP",
	}
	var method = jwt.Method{
		Name:   "GOST15",
		Module: nil,
	}
	token := jwt.NewWithClaims(method, claims)
	t.Logf("token: %+s\n", token.String())
	token.Method.Init()
	if err := token.Method.LoadKeyStore(); err != nil {
		t.Log(err)
	}
	s, err := token.Method.Sign(token.String())
	if err != nil {
		t.Log(err)
	} else {
		t.Logf("sign: %+s\n", s)

	}
}
