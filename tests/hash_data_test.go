package tests

import (
	"encoding/base64"
	"testing"

	rand "math/rand/v2"

	"pki.gov.kz/go/kalkan"
)

func TestHashData(t *testing.T) {
	var h string

	for i := 0; i <= 100; i++ {
		h = testHashData(t)
		t.Log("hashed: ", h)
		raw, err := base64.StdEncoding.DecodeString(h)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("hash len: ", len(raw))
		t.Log("ittertion: ", i)
	}
}

func testHashData(t *testing.T) string {
	flags := kalkan.FlagInBase64 | kalkan.FlagOutBase64 | kalkan.FlagHashGOST15
	r := make([]byte, rand.UintN(8)+1)
	raw := base64.StdEncoding.EncodeToString(r)
	t.Log("to hash: ", raw)
	h, err := mod.HashData("", flags, raw)
	if err != nil {
		t.Fatal(err)
	}
	return h
}

func BenchmarkNashSameData(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		flags := kalkan.FlagInBase64 | kalkan.FlagOutBase64 | kalkan.FlagHashGOST15
		raw := base64.StdEncoding.EncodeToString([]byte{1, 2})
		h, err := mod.HashData("", flags, raw)
		if err != nil {
			b.Fatal(err)
		}
		b.Log("Hashed Data: ", h)
	}
}

func BenchmarkDiffData(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		flags := kalkan.FlagInBase64 | kalkan.FlagOutBase64 | kalkan.FlagHashGOST15
		r := make([]byte, rand.UintN(8)+1)
		raw := base64.StdEncoding.EncodeToString(r)
		b.Log("Raw Data: ", raw)
		h, err := mod.HashData("", flags, raw)
		if err != nil {
			b.Fatal(err)
		}
		b.Log("Hashed Data: ", h)
	}
}
