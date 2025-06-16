package tests

import (
	"testing"
	"time"

	"pki.gov.kz/go/kalkan/jwt"
)

var err error

func TestNewToken(t *testing.T) {
	issuedAt := time.Now().Add(-120 * time.Second)
	expiresAt := issuedAt.Add(900 * time.Second)
	var claims = jwt.Claims{
		IssuedAt:  &issuedAt,
		ExpiresAt: &expiresAt,
		BIN:       "161140016747",
		Check:     "OTP",
	}
	var method = jwt.Method{
		Name:   "GOST15",
		Module: nil,
	}
	token := jwt.NewWithClaims(method, claims)
	token.Method.Init()
	if err := token.Method.LoadKeyStore(); err != nil {
		t.Log(err)
	}
	token.Signature, err = token.Method.Sign(token.Header.String() + "." + token.Claims.String())
	if err != nil {
		t.Log(err)
	} else {
		t.Logf("sign: %+s\n", token.Signature)

	}
}
