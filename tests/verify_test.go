package tests

import "testing"

func TestVerifyData(t *testing.T) {
	sData := testSignData(t)
	_, err := mod.VerifyData("", sData, "", flags)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyXML(t *testing.T) {
	sData := testSignXML(t)
	_, err := mod.VerifyXML("", sData, 0)
	if err != nil {
		t.Fatal(err)
	}
}
