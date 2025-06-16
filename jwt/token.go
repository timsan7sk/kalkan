package jwt

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type header map[string]any
type signature string

// Represents a JWT Token.  Different fields will be used depending on whether you're
// creating or verifying a token.
type Token struct {
	Finish    string    `json:"-"`                // The complite token. Populated after ReplaceAll method.
	Method    Method    `json:"-"`                // The signing method used or to be used
	Header    header    `json:"header,omitempty"` // The first segment of the token
	Claims    Claims    `json:"claims,omitempty"` // The second segment of the token
	Signature signature `json:"sign,omitempty"`   // The third segment of the token.
	Valid     bool      `json:"-"`                // Is the token valid?  Populated when you Verify a token
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

// Marshal to JSON the t.Header and t.Claims, then encode them into a Base64 string.
func (t Token) StringBase64() string {
	h, _ := json.Marshal(t.Header)
	c, _ := json.Marshal(t.Claims)
	return base64.StdEncoding.EncodeToString(h) + "." + base64.StdEncoding.EncodeToString(c)
}

// Replaces "/" to "_" and "+" to "-" and "=" to ""
// and assigns a value to t.Finish.
func (t *Token) ReplaceAll() {
	data := t.StringBase64() + "." + string(t.Signature)
	out := strings.ReplaceAll(data, "/", "_")
	out = strings.ReplaceAll(out, "+", "-")
	out = strings.ReplaceAll(out, "=", "")
	t.Finish = out
}
