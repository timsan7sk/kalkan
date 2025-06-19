package tests

import (
	"testing"

	"pki.gov.kz/go/kalkan"
)

func TestSignData(t *testing.T) {
	_ = signDataTest(t)
}

func signDataTest(t *testing.T) string {
	data := "QOlCN2P7QVHU5mqcH"
	flags := kalkan.FlagSignCMS | kalkan.FlagInBase64 | kalkan.FlagOutBase64
	sData, err := mod.SignData("", data, "", flags)
	if err != nil {
		t.Fatal(err)
	}
	return sData
}
