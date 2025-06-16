package jwt

import (
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
	module *kalkan.Module
}

var (
	err   error
	path  string      = "GOST512.p12"
	pwd   string      = "Qazwsx!@#123"
	flags kalkan.Flag = kalkan.FlagSignDraft | kalkan.FlagOutBase64
)

func (m *Method) Init() error {
	m.module, err = kalkan.Open("libkalkancryptwr-64.so.2")
	if err != nil {
		return err
	}
	if err := m.module.Init(); err != nil {
		return err
	}
	return nil
}

func (m *Method) Alg() string {
	return m.Name
}
func (m *Method) Verify(signingString, signature string) error {
	return nil
}
func (m *Method) LoadKeyStore() error {
	return m.module.LoadKeyStore(path, pwd, "", kalkan.StoreTypePKCS12)
}

// Sign implements token signing for the Method.
func (m *Method) Sign(inData string) (signature, error) {
	if err := m.LoadKeyStore(); err != nil {
		return "", err
	}
	v, err := m.module.SignData("", inData, "", flags)
	if err != nil {
		return "", err
	}
	defer m.module.Finalize()
	return signature(v), nil
}
