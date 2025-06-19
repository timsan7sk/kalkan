package tests

import "testing"

func TestVerifyData(t *testing.T) {
	sData := testSignData(t)
	vData, err := mod.VerifyData("", sData, "", flags)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Verified data: ", vData)
}

func TestVerifyXML(t *testing.T) {
	sData := testSignXML(t)
	vData, err := mod.VerifyXML("", sData, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Verified XML: ", vData)

}
