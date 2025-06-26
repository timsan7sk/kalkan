package tests

import (
	"testing"

	"pki.gov.kz/go/kalkan"
)

// TODO: fix test.
func TestGetToken(t *testing.T) {
	_, err := mod.GetToken(kalkan.StoreTypeKazToken)
	if err != nil {
		t.Fatal(err)
	}

}
