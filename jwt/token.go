package jwt

import (
	"encoding/json"
)

type header map[string]any

// Token represents a JWT Token.  Different fields will be used depending on whether you're
// creating or parsing/verifying a token.
type Token struct {
	Raw       string `json:"-"`                // The raw token.  Populated when you Parse a token
	Method    Method `json:"-"`                // The signing method used or to be used
	Header    header `json:"header,omitempty"` // The first segment of the token
	Claims    Claims `json:"claims,omitempty"` // The second segment of the token
	Signature string `json:"sign,omitempty"`   // The third segment of the token.  Populated when you Parse a token
	Valid     bool   `json:"-"`                // Is the token valid?  Populated when you Parse/Verify a token
}

// New creates a new Token with the specified signing method and an empty map of claims.
func New(method Method) *Token {
	return NewWithClaims(method, Claims{})
}

// NewWithClaims creates a new Token with the specified signing method and claims.
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

func (t Token) String() (out string) {

	b, err := json.Marshal(t)
	if err == nil {
		out = string(b)
	}
	return
}
