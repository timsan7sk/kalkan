package tests

import (
	"testing"

	"pki.gov.kz/go/kalkan"
)

func TestHashData(t *testing.T) {
	_ = testHashData(t)
}

func testHashData(t *testing.T) string {
	flags := kalkan.FlagHashGOST15 | kalkan.FlagOutBase64
	h, err := mod.HashData("", flags, "TestDataToHash")
	if err != nil {
		t.Fatal(err)
	}
	return h
}
