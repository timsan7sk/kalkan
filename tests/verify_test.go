package tests

import (
	"testing"

	"pki.gov.kz/go/kalkan"
)

func TestVerifyData(t *testing.T) {
	var flags = kalkan.FlagSignCMS | kalkan.FlagInBase64 | kalkan.FlagOutBase64
	inSign := testSignData(t, flags)
	_, err := mod.VerifyData(inSign, "", "", flags)
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
