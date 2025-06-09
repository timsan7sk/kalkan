package tests

import (
	"fmt"
	"os"
	"testing"

	"pki.gov.kz/go/kalkan"
)

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
