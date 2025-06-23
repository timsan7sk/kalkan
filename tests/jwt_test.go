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
		Check:     "OTP",
	}
	var method = jwt.Method{
		Name: "ES256",
	}
	token := jwt.NewWithClaims(method, claims)
	if err = token.Method.Open(libName); err != nil {
		t.Fatal(err)
	}
	if err = token.Method.Init(); err != nil {
		t.Fatal(err)
	}
	if err = token.Method.LoadKeyStore(path, pwd); err != nil {
		t.Fatal(err)
	}
	token.Signature, err = token.Method.Sign(token.StringBase64())
	if err != nil {
		t.Fatal(err)
	}
	token.ReplaceAll()
	// token.Method.Finialize()
	if err = token.Method.Close(); err != nil {
		t.Fatal(err)
	}

}
