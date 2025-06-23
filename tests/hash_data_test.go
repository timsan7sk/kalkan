package tests

import (
	"testing"

	"pki.gov.kz/go/kalkan"
)

func TestHashData(t *testing.T) {
	flags := kalkan.FlagHashGOST95 | kalkan.FlagOutBase64
	_, err := mod.HashData(kalkan.HashAlgGOST95, flags, "TestDataToHash")
	if err != nil {
		t.Fatal(err)
	}
}
