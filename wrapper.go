package kalkan

import (
	"net/url"
	"time"
)

// Wrapper over KalkanCrypt library methods.
type Kalkan interface {
	Init() error // Initializes the library.

	LoadKeyStore(path, pwd, alias string, storeType StoreType) error // Loads keys/certificates from storage.

	HashData(alg HashAlg, flags Flag, inData string) (string, error) // Hashes data.

	SignData(inSign, inData, alias string, flag Flag) (string, error)                                          // Signs data.
	SignXML(xml, alias string, flags Flag, signNodeID, parentSignNode, parentNameSpace string) (string, error) // Signs data in XML format.
	SignWSSE(alias, inData, signNodeID string, flags Flag) (string, error)                                     // Signs SOAP messages according to the WS-Security specification. Required for [SmartBridge]:https://sb.egov.kz/.

	VerifyData(inSign, inData, alias string, flags Flag) (*VerifiedData, error) // Verifies signature.
	VerifyXML(alias, inData string, flags Flag) (string, error)                 // Verifies the signature of XML data.

	GetCertFromCMS(cms string, signID int, flag Flag) (string, error) // Obtains a certificate from a signature in CMS format.
	GetCertFromXML(xml string, signID int) (string, error)            // Obtains a certificate from a data signature in XML format.
	GetCertList() (string, int, error)                                // Obtains a list of certificates as a string and their number.

	GetTimeFromSign(cms string, sigID int, flag Flag) (time.Time, error) // Obtains the signature time.

	GetToken(store StoreType) (string, error) // Obtains a string of connected storage devices and their number.

	GetLastError() ErrorCode                 // Obtains last error code.
	GetLastErrorString() (ErrorCode, string) // Obtains the code and text of the last error.

	GetSignAlgorithmFromXML(xml string) (string, error) // Obtains the signature algorithm from XML.

	SetProxy(flag Flag, url *url.URL) error // Set proxy server settings.
	TSASetURL(tsaURL string) error          // Sets the TSA service address.

	X509CertificateGetInfo(inCert string, prop CertProp) (string, error)                                   // Obtains field/extension values ​​from a certificate.
	X509ExportCertificateFromStore(alias string) (string, error)                                           // Exports a certificate from the store.
	X509LoadCertificateFromBuffer(inCert []byte, flags CertEncodingType) error                             // Loads a certificate from a byte slice.
	X509LoadCertificateFromFile(certPath string, certType CertType) error                                  // Loads cerificate from file.
	X509ValidateCertificate(inCert string, validateType ValidateType, validatePath string) (string, error) // Validates the certificate.

	XMLFinalize() // Frees memory and terminates the library with modules responsible for parsing, signing and checking data in XML format.
	Finalize()    // Frees resources of the KalkanCryptCOM crypto provider and terminates the library.
	Close() error // Closes a Module.
}

var _ Kalkan = (*Module)(nil)
