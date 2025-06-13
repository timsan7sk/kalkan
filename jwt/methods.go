package jwt

import "pki.gov.kz/go/kalkan"

// SigningMethod can be used add new methods for signing or verifying tokens.
type Methods interface {
	Verify(signingString, signature string) error      // Returns nil if signature is valid
	Sign(inData string, sm signMethod) (string, error) // Returns encoded signature or error
	Alg() string                                       // returns the alg identifier for this method (example: 'HS256')
}

type signMethod func(inSign, inData, alias string, flag kalkan.Flag) (string, error)

// SigningMethodGOST15 implements the GOST 2015 family of signing methods.
// Expects key type of []byte for both signing and validation
type Method struct {
	Name string
}

func (m *Method) Alg() string {
	return m.Name
}
func (m *Method) Verify(signingString, signature string) error {
	return nil
}

// Sign implements token signing for the SigningMethod.
// Key must be []byte
func (m *Method) Sign(inData string, sm signMethod) (string, error) {
	return "", ErrInvalidKeyType
}

func init() {

}
