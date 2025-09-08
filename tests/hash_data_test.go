package tests

import (
	"encoding/base64"
	"testing"

	rand "math/rand/v2"

	"pki.gov.kz/go/kalkan"
)

func TestHashData(t *testing.T) {
	h := testHashData(t)
	t.Log("len:", len(h), "hashed: ", h)
}

func testHashData(t *testing.T) string {
	flags := kalkan.FlagInBase64 | kalkan.FlagOutBase64 | kalkan.FlagHashGOST15
	r := make([]byte, rand.UintN(64)+1)
	raw := base64.StdEncoding.EncodeToString(r)
	h, err := mod.HashData("", flags, raw)
	if err != nil {
		t.Fatal(err)
	}
	return h
}

func BenchmarkDiffData(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		flags := kalkan.FlagInBase64 | kalkan.FlagOutDER | kalkan.FlagHashGOST15
		r := make([]byte, rand.UintN(64)+1)
		raw := base64.StdEncoding.EncodeToString(r)
		_, err := mod.HashData("", flags, raw)
		if err != nil {
			b.Fatal(err)
		}
	}
}
