package tests

import (
	"testing"

	"pki.gov.kz/go/kalkan"
)

// TODO: fix test.
func TestGetTokens(t *testing.T) {
	v, n, err := mod.GetTokens(kalkan.StoreTypePKCS12)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("v: %s n: %d\n", v, n)
}
