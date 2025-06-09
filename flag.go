package kalkan

// Flag for the KalcanCrypt library.
type Flag int

const (
	FlagSignDraft       Flag = 0x00000001 // Draft sign.
	FlagSignCMS         Flag = 0x00000002 // Signature in CMS format.
	FlagInPEM           Flag = 0x00000004 // Input data in PEM format.
	FlagInDER           Flag = 0x00000008 // Input data in DER encoding.
	FlagInBase64        Flag = 0x00000010 // Input data in Base64 encoding.
	FlagExInBase64      Flag = 0x00000020 // Extra input data in BASE64 encoding.
	FlagDetachedSign    Flag = 0x00000040 // Detached signature.
	FlagAttachCert      Flag = 0x00000080 // Attach the certificate to the signature.
	FlagWithTimestamp   Flag = 0x00000100 // Add a timestamp to the signature.
	FlagOutPEM          Flag = 0x00000200 // Output data in PEM format.
	FlagOutDER          Flag = 0x00000400 // Output data in DER encoding.
	FlagOutBase64       Flag = 0x00000800 // Output data in Base64 encoding.
	FlagProxyOff        Flag = 0x00001000 // Disable proxy server usage and clear settings.
	FlagProxyOn         Flag = 0x00002000 // Enable and set proxy server settings (address and port).
	FlagProxyAuth       Flag = 0x00004000 // The proxy server requires authorization (login/password).
	FlagInFile          Flag = 0x00008000 // Use if the inData/outData parameter contains an absolute path to a file.
	FlagNoCheckCertTime Flag = 0x00010000 // Do not check the certificate expiration date when building a chain to the root (to check old signatures with an expired certificate).
	FlagHashSHA256      Flag = 0x00020000 // SHA256 hashing algorithm.
	FlagHashGOST95      Flag = 0x00040000 // GOST 34311-95 hashing algorithm.
	FlagGetOCSPResponse Flag = 0x00080000 // Get a response from the OCSP service.
	FlagHashGOST15      Flag = 0x00200000 // GOST 34311-2015 hashing algorithm.
)
