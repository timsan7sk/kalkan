package jwt

import (
	"pki.gov.kz/go/kalkan"
)

// Available methods of working with the token.
type Methods interface {
	Open(libName string) error                    // Opens library or returns error.
	Init() error                                  // Initialize the library.
	LoadKeyStore(path, pwd string) error          // Loads key storage.
	Sign(inData string) (string, error)           // Returns encoded signature or error.
	Finialize()                                   // Finializes work with library and frees the memory.
	Close() error                                 // Closes the library.
	Verify(signingString, signature string) error // Returns nil if signature is valid.
	Alg() string                                  // returns the alg identifier for this method (example: 'HS256').
}

const (
	// Flags for singing JWT.
	flags kalkan.Flag = kalkan.FlagSignDraft | kalkan.FlagOutBase64
)

// Error variable.
var err error

// Method implements the algorithm family of signing methods.
type Method struct {
	Name   string
	module *kalkan.Module
}

// Opens library or returns error.
func (m *Method) Open(libName string) error {
	m.module, err = kalkan.Open(libName)
	if err != nil {
		return err
	}
	return nil
}

// Initialize the library.
func (m *Method) Init() error {
	if err := m.module.Init(); err != nil {
		return err
	}
	return nil
}
func (m *Method) LoadKeyStore(path, pwd string) error {
	return m.module.LoadKeyStore(path, pwd, "", kalkan.StoreTypePKCS12)
}

// Obtains the algorithm name.
func (m *Method) Alg() string {
	return m.Name
}

// Verifies the signature.
func (m *Method) Verify(signingString, signature string) error {
	return nil
}

// Implements signing.
func (m *Method) Sign(inData string) (signature, error) {
	v, err := m.module.SignData("", inData, "", flags)
	if err != nil {
		return "", err
	}
	return signature(v), nil
}

// Finializes work with library and frees the memory.
func (m *Method) Finialize() {
	m.module.Finalize()
}

// Closes the library.
func (m *Method) Close() error {
	return m.module.Close()
}
