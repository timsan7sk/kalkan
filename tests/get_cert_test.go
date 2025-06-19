package tests

import "testing"

func TestGetCertFromXML(t *testing.T) {
	sData := testSignXML(t)
	_, err := mod.GetCertFromXML(sData, 1)
	if err != nil {
		t.Fatal(err)
	}
}
