package tests

import (
	"net/url"
	"testing"

	"pki.gov.kz/go/kalkan"
)

func TestSetProxy(t *testing.T) {
	mockURL, err := url.Parse("mockproxy.kz")
	if err != nil {
		t.Fatal(err)
	}
	err = mod.SetProxy(kalkan.FlagProxyOn, mockURL)
	if err != nil {
		t.Fatal(err)
	}
}
