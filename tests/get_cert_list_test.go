package tests

import "testing"

// TODO: fix the test.
func TestGetCertList(t *testing.T) {
	l, n, err := mod.GetCertList()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("l: %s n: %d", l, n)
}
