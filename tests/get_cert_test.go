package tests

import "testing"

func testGetCertFromXML(t *testing.T) string {
	sData := testSignXML(t)
	c, err := mod.GetCertFromXML(sData, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Arg: %s\n", c)
	return c

}
func TestGetCertFromXML(t *testing.T) {
	_ = testGetCertFromXML(t)
}
