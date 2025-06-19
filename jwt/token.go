package jwt

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
)

// Represents a JWT Token.  Different fields will be used depending on whether you're
// creating or verifying a token.
type Token struct {
	Method    Method         `json:"-"` // The signing method used or to be used
	Header    map[string]any `json:""`  // The first segment of the token
	Claims    Claims         `json:""`  // The second segment of the token
	Signature string         `json:""`  // The third segment of the token.
	Valid     bool           `json:"-"` // Is the token valid?  Populated when you Verify a token
}

// Creates a new Token with the specified signing method and an empty claims.
func New(method Method) *Token {
	return NewWithClaims(method, Claims{})
}

// Creates a new Token with the specified signing method and claims.
func NewWithClaims(method Method, claims Claims) *Token {
	return &Token{
		Header: map[string]any{
			"typ": "JWT",
			"alg": method.Alg(),
		},
		Claims: claims,
		Method: method,
	}
}

// Marshal to JSON the t.Header and t.Claims, then encode them into a "URL and Filename safe" Base64 string.
// see [RFC4648]https://datatracker.ietf.org/doc/html/rfc4648#page-8
func (t Token) String() string {
	if reflect.DeepEqual(t.Header, map[string]any{}) && reflect.DeepEqual(t.Claims, Claims{}) {
		return ""
	}
	h, _ := json.Marshal(t.Header)
	c, _ := json.Marshal(t.Claims)
	e := base64.URLEncoding.WithPadding(base64.NoPadding)

	return e.EncodeToString(h) + "." + e.EncodeToString(c)
}
