package jwt

import (
	"fmt"

	"pki.gov.kz/go/kalkan"
)

// SigningMethod can be used add new methods for signing or verifying tokens.
type Methods interface {
	Verify(signingString, signature string) error // Returns nil if signature is valid
	Sign(inData string) (string, error)           // Returns encoded signature or error
	Alg() string                                  // returns the alg identifier for this method (example: 'HS256')
}

// Method implements the algorithm family of signing methods.
// Expects key type of []byte for both signing and validation
type Method struct {
	Name   string
	Module *kalkan.Module
}

var (
	err  error
	mod  *kalkan.Module
	path string
	pwd  string
	flag kalkan.Flag
)

func (m *Method) Alg() string {
	return m.Name
}
func (m *Method) Verify(signingString, signature string) error {
	return nil
}
func (m *Method) LoadKeyStore() error {
	return m.Module.LoadKeyStore(path, pwd, "", kalkan.StoreTypePKCS12)
}

// Sign implements token signing for the Method.
func (m *Method) Sign(inData string) (string, error) {

	if err := m.LoadKeyStore(); err != nil {
		return "", err
	}
	v, err := m.Module.SignData("", inData, "", flag)
	if err != nil {
		return "", err
	}
	defer m.Module.Finalize()
	return v, nil

}

func init() {
	mod, err = kalkan.Open("libkalkancryptwr-64.so.2")
	if err != nil {
		fmt.Println(err)
	}
	if err := mod.Init(); err != nil {
		fmt.Println(err)
	}
}
