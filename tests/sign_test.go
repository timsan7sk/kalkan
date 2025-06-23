package tests

import (
	"testing"

	"pki.gov.kz/go/kalkan"
)

func TestSignData(t *testing.T) {
	var flags = kalkan.FlagSignCMS | kalkan.FlagInBase64 | kalkan.FlagOutBase64
	_ = testSignData(t, flags)
}
func TestSignXML(t *testing.T) {
	_ = testSignXML(t)
}

func TestSignWSSE(t *testing.T) {
	data, err := kalkan.WrapSOAP("<test>data</test>", "id-ASDASDAS7786SA9S6D78A6SD9876ASD6")
	if err != nil {
		t.Fatal(err)
	}
	_, err = mod.SignWSSE("", data, "", 0)
	if err != nil {
		t.Fatal(err)
	}
}

func testSignData(t *testing.T, flags kalkan.Flag) string {

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
