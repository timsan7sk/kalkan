package tests

import (
	"testing"

	"pki.gov.kz/go/kalkan"
)

func TestHashData(t *testing.T) {
	_ = testHashData(t)
}

func testHashData(t *testing.T) string {
	flags := kalkan.FlagHashGOST95 | kalkan.FlagOutBase64
	h, err := mod.HashData(kalkan.HashAlgGOST95, flags, "TestDataToHash")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(h)
	return h
}
