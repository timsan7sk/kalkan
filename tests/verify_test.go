package tests

import "testing"

func TestVerifyData(t *testing.T) {
	sData := testSignData(t)
	t.Log(sData)
	vData, err := mod.VerifyData("", sData, "", flags)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Verified data: ", vData)
}
