package tests

import "testing"

func TestGetCertFromXML(t *testing.T) {
	sData := testSignXML(t)
	gCert, err := mod.GetCertFromXML(sData, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(gCert))
}
