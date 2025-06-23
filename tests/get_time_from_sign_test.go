package tests

import (
	"testing"

	"pki.gov.kz/go/kalkan"
)

// TODO: fix text
func TestGetTimeFromSign(t *testing.T) {
	var flags = kalkan.FlagSignCMS | kalkan.FlagWithTimestamp
	s := testSignData(t, flags)
	v, err := mod.GetTimeFromSign(s, 0, flags)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}
