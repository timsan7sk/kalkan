package tests

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"pki.gov.kz/go/kalkan"
)

//go:embed GOST512.p12
var testKeyGOST1 []byte

// //go:embed gost1.cer
// var testCertGOST1 string

func TestMain(m *testing.M) {
	mod, err := kalkan.Open("libkalkancryptwr-64.so.2")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := mod.Init(); err != nil {
		fmt.Println(err)
	}
	m.Run()
}
