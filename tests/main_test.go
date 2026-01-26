package tests

import (
	_ "embed"
	"log"
	"os"
	"testing"

	"pki.gov.kz/go/kalkan"
)

const (
	libName = "libkalkancryptwr-64.so.2" // Library name.
	path    = "nca_is.p12"
)

var (
	// Error variable.
	err error
	// Module variable.
	mod *kalkan.Module
	//go:embed kaznachei.p12
	testKeyKAZNACHEI []byte
	//go:embed GOST512.cer
	testCertGOST1 string
	//go:embed GOST512.crt
	testCertGOST2 string
	//go:embed pwd_is.txt
	pwd string
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}

func TestMain(m *testing.M) {
	mod, err = kalkan.Open(libName, kalkan.TestOpts...)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if err = mod.Init(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Printf("pwd: %s path: %s\n", pwd, path)
	if err = mod.LoadKeyStore(path, pwd, "", kalkan.StoreTypePKCS12); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// code, s := mod.GetLastErrorString()
	// log.Println(code)
	// log.Println(s)
	x := m.Run()
	os.Exit(x)
	// defer mod.XMLFinalize()
	// defer mod.Finalize()

	// if err = mod.Close(); err != nil {
	// 	fmt.Println(err)
	// }
}
