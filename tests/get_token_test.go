package tests

import (
	"fmt"
	"testing"

	"pki.gov.kz/go/kalkan"
)

// TODO: fix test.
func TestGetToken(t *testing.T) {
	v, err := mod.GetToken(kalkan.StoreTypeKazToken)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("TRA TA TA")
	t.Logf("v: %s \n", v)
}
