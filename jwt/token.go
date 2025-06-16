package jwt

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type header map[string]any
type signature string

// Represents a JWT Token.  Different fields will be used depending on whether you're
// creating or parsing/verifying a token.
type Token struct {
	Raw       string    `json:"-"`                // The raw token.  Populated when you Parse a token
	Method    Method    `json:"-"`                // The signing method used or to be used
	Header    header    `json:"header,omitempty"` // The first segment of the token
	Claims    Claims    `json:"claims,omitempty"` // The second segment of the token
	Signature signature `json:"sign,omitempty"`   // The third segment of the token.  Populated when you Parse a token
	Valid     bool      `json:"-"`                // Is the token valid?  Populated when you Parse/Verify a token
}

// Creates a new Token with the specified signing method and an empty map of claims.
func New(method Method) *Token {
	return NewWithClaims(method, Claims{})
}

// Creates a new Token with the specified signing method and claims.
func NewWithClaims(method Method, claims Claims) *Token {
	return &Token{
		Header: header{
			"typ": "JWT",
			"alg": method.Alg(),
		},
		Claims: claims,
		Method: method,
	}
}

func (t *Token) Parse() {

}

// Obtains Header and Claims fields of Token in Base64 string
func (t Token) StringBase64() (string, error) {
	h, err := json.Marshal(t.Header)
	if err != nil {
		return "", err
	}
	c, err := json.Marshal(t.Claims)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h) + "." + base64.StdEncoding.EncodeToString(c), nil
}

// Replaces "/" to "_" and "+" to "-" and "=" to "".
func (t Token) ReplaceAll() (out string) {

	out = strings.ReplaceAll(string(t.Signature), "/", "_")
	out = strings.ReplaceAll(out, "+", "-")
	out = strings.ReplaceAll(out, "=", "")
	return
}
