package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"reflect"

	"pki.gov.kz/go/kalkan"
)

const (
	// Flags for singing JWT.
	flags kalkan.Flag = kalkan.FlagSignDraft | kalkan.FlagOutBase64
)

// Represents a JWT Token.  Different fields will be used depending on whether you're
// creating or verifying a token
type Token struct {
	// Finish    string         // The complite token. Populated after ReplaceAll method
	Header    Header         // The first segment of the token
	Claims    Claims         // The second segment of the token
	Module    *kalkan.Module // The KalkanCrypt module
	Signature string         // The third segment of the token
	// Valid     bool           // Is the token valid?  Populated when you Verify a token
}

// Creates a new Token with the specified claims and KalkanCrypt module.
func New(claims Claims, mod *kalkan.Module) *Token {
	return &Token{
		Header: Header{
			"typ": "JWT",
			"alg": "GOST15",
		},
		Claims: claims,
		Module: mod,
	}
}

// Marshal to JSON the t.Header and t.Claims, then encode them into a "URL and Filename safe" Base64 string er error.
//
// see https://datatracker.ietf.org/doc/html/rfc4648#page-8
func (t Token) stringForSign() (string, error) {
	if reflect.DeepEqual(t.Header, Header{}) || reflect.DeepEqual(t.Claims, Claims{}) {
		return "", errors.New("header or claims are empty")
	}
	// set encoder.
	h, err := json.Marshal(t.Header)
	if err != nil {
		return "", err
	}
	c, err := json.Marshal(t.Claims)
	if err != nil {
		return "", err
	}
	e := base64.URLEncoding.WithPadding(base64.NoPadding)
	return e.EncodeToString(h) + "." + e.EncodeToString(c), nil
}

// Signs the t.Header and t.Claims elements encoded into a "URL and Filename safe" Base64 string
// and assigns the result to the t.Signature filed.
//
// see https://datatracker.ietf.org/doc/html/rfc4648#page-8
func (t *Token) Sign() error {
	if reflect.DeepEqual(t.Header, Header{}) || reflect.DeepEqual(t.Claims, Claims{}) {
		return errors.New("header, claims or signature are empty")
	}
	// get data for singing
	data, err := t.stringForSign()
	if err != nil {
		return err
	}
	// singing data
	s, err := t.Module.SignData("", data, "", flags)
	if err != nil {
		return err
	}
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	e := base64.URLEncoding.WithPadding(base64.NoPadding)
	// assignment of the result
	t.Signature = e.EncodeToString(b)
	return nil
}

// Obtains JWT encoded into "URL and Filename safe" Base64 string or error.
//
// see https://datatracker.ietf.org/doc/html/rfc4648#page-8
func (t *Token) GetToken() (string, error) {
	if reflect.DeepEqual(t.Header, Header{}) || reflect.DeepEqual(t.Claims, Claims{}) || t.Signature == "" {
		return "", errors.New("header, claims or signature are empty")
	}
	hc, err := t.stringForSign()
	if err != nil {
		return "", err
	}
	return hc + "." + t.Signature, nil
}
