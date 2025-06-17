package tests

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"pki.gov.kz/go/kalkan"
)

// Library name.
const libName = "libkalkancryptwr-64.so.2"

var (
	// Error variable.
	err error
	// Module variable.
	mod *kalkan.Module
	//go:embed GOST512.p12
	testKeyGOST1 []byte
)

// //go:embed gost1.cer
// var testCertGOST1 string

func TestMain(m *testing.M) {
	mod, err = kalkan.Open(libName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = mod.Init(); err != nil {
		fmt.Println(err)
	}
	m.Run()
	mod.XMLFinalize()
	mod.Finalize()

	if err = mod.Close(); err != nil {
		fmt.Println(err)
	}
}
