package kalkan

import "runtime"

// Obtains the library name depending on the OS.
func getLibraryName() string {
	switch runtime.GOOS {
	case "freebsd":
		return "libkalkancryptwr-64.so.2"
	case "linux":
		return "libkalkancryptwr-64.so"
	default:
		panic("GOOS=" + runtime.GOOS + " is not supported")
	}
}
