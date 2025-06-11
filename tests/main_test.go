package tests

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"pki.gov.kz/go/kalkan"
)

//go:embed gost1.p12
var testKeyGOST1 []byte

//go:embed gost1.cer
var testCertGOST1 string

//go:embed gost2.p12
var testKeyGOST2 []byte

//go:embed gost2.cer
var testCertGOST2 string

func TestMain(m *testing.M) {
	mod, err := kalkan.Open("libkalkancryptwr-64.so.2")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := mod.Init(); err != nil {
		fmt.Println(err)
	}
}
