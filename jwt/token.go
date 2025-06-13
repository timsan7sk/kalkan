package jwt

import "encoding/base64"

// Token represents a JWT Token.  Different fields will be used depending on whether you're
// creating or parsing/verifying a token.
type Token struct {
	Raw       string         // The raw token.  Populated when you Parse a token
	Method    Method         // The signing method used or to be used
	Header    map[string]any // The first segment of the token
	Claims    Claims         // The second segment of the token
	Signature string         // The third segment of the token.  Populated when you Parse a token
	Valid     bool           // Is the token valid?  Populated when you Parse/Verify a token
}

// New creates a new Token with the specified signing method and an empty map of claims.
func New(method Method) *Token {
	return NewWithClaims(method, MapClaims{})
}

// NewWithClaims creates a new Token with the specified signing method and claims.
func NewWithClaims(method Method, claims Claims) *Token {
	return &Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": method.Alg(),
		},
		Claims: claims,
		Method: method,
	}
}

// EncodeSegment encodes a JWT specific base64url encoding with padding stripped
//
// Deprecated: In a future release, we will demote this function to a non-exported function, since it
// should only be used internally
func EncodeSegment(seg []byte) string {
	return base64.RawURLEncoding.EncodeToString(seg)
}
