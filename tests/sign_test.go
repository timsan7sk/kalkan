package tests

import (
	"testing"

	"pki.gov.kz/go/kalkan"
)

var flags = kalkan.FlagSignCMS | kalkan.FlagInBase64 | kalkan.FlagOutBase64

func TestSignData(t *testing.T) {
	_ = testSignData(t)
}
func TestSignXML(t *testing.T) {
	_ = testSignXML(t)
}

func testSignData(t *testing.T) string {
	data := "QOlCN2P7QVHU5mqcH"
	sData, err := mod.SignData("", data, "", flags)
	if err != nil {
		t.Fatal(err)
	}
	return sData
}

func testSignXML(t *testing.T) string {
	data := "<test>data</test>"
	sData, err := mod.SignXML(data, "", 0, "", "", "")
	if err != nil {
		t.Fatal(err)
	}
	return sData
}
