package tests

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"pki.gov.kz/go/kalkan"
)

const (
	libName = "libkalkancryptwr-64.so.2" // Library name.
	path    = "GOST512.p12"
)

var (
	// Error variable.
	err error
	// Module variable.
	mod *kalkan.Module
	//go:embed GOST512.p12
	testKeyGOST1 []byte
	//go:embed GOST512.cer
	testCertGOST1 string
	//go:embed GOST512.crt
	testCertGOST2 string
	//go:embed pass.txt
	pwd string
)

func TestMain(m *testing.M) {
	mod, err = kalkan.Open(libName, kalkan.TestOpts...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = mod.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(err)
	if err = mod.LoadKeyStore(path, pwd, "", kalkan.StoreTypePKCS12); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	m.Run()
	mod.XMLFinalize()
	mod.Finalize()

	if err = mod.Close(); err != nil {
		fmt.Println(err)
	}
}
