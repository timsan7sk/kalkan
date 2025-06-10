package kalkan

import "strconv"

type ErrorCode int64 // Runtime error code.

// Returns the error text associated with the given code.
func (t ErrorCode) String() string {
	if v, ok := errorLabels[t]; ok {
		return v
	}
	return "Unknown error code."
}

// Returns the string representation of t in the hex base,
func (t ErrorCode) Hex() string {
	return strconv.FormatInt(int64(t), 16)
}

// Error codes.
const (
	ErrorCodeOK                           ErrorCode = 0x00000000 // No error.
	ErrorCodeReadPKCS12                   ErrorCode = 0x08F00002 // Unable to read file format PKCS#12.
	ErrorCodeOpenPKCS12                   ErrorCode = 0x08F00003 // Unable to open file format PKCS#12.
	ErrorCodeInvalidExtID                 ErrorCode = 0x08F00004 // Invalid certificate extension identifier.
	ErrorCodeBufferTooSmall               ErrorCode = 0x08F00005 // Buffer size is too small.
	ErrorCodeCertParse                    ErrorCode = 0x08F00006 // Unable to parse the certificate.
	ErrorCodeInvalidFlag                  ErrorCode = 0x08F00007 // Invalid flag.
	ErrorCodeOpenFile                     ErrorCode = 0x08F00008 // Can't open file.
	ErrorCodeInvalidPassword              ErrorCode = 0x08F00009 // Wrong password.
	ErrorCodeAllocMemory                  ErrorCode = 0x08F0000d // Unable to allocate memory.
	ErrorCodeCheckChain                   ErrorCode = 0x08F0000e // CA certificate or user certificate not found when checking chain.
	ErrorCodeCertValidType                ErrorCode = 0x08F00010 // Invalid certificate validation type.
	ErrorCodeBadCRLFormat                 ErrorCode = 0x08F00011 // Incorrect CRL format.
	ErrorCodeLoadCRL                      ErrorCode = 0x08F00012 // Unable to load CRL.
	ErrorCodeLoadCRLs                     ErrorCode = 0x08F00013 // Unable to load CRLs.
	ErrorCodeUnknownAlg                   ErrorCode = 0x08F00015 // Unknown signature algorithm.
	ErrorCodeKeyNotFound                  ErrorCode = 0x08F00016 // User's private key not found.
	ErrorCodeSignInit                     ErrorCode = 0x08F00017 // Unable to initialize signature manager.
	ErrorCodeSign                         ErrorCode = 0x08F00018 // Failed to generate digital signature.
	ErrorCodeEncrypt                      ErrorCode = 0x08F00019 // Encryption error.
	ErrorCodeInvalidFlags                 ErrorCode = 0x08F0001a // Invalid flags.
	ErrorCodeCertNotFound                 ErrorCode = 0x08F0001b // User certificate not found.
	ErrorCodeVerifySign                   ErrorCode = 0x08F0001c // Error verifying XML signature.
	ErrorCodeBase64Decrypt                ErrorCode = 0x08F0001d // Error decrypting from Base64.
	ErrorCodeUnknownCMSFormat             ErrorCode = 0x08F0001e // Unknown CMS format.
	ErrorCodeCACertNotFound               ErrorCode = 0x08F00020 // CA certificate not found.
	ErrorCodeXMLSecInit                   ErrorCode = 0x08F00021 // Error initializing XMLSec.
	ErrorCodeLoadTrustedCerts             ErrorCode = 0x08F00022 // Error loading trusted certificates.
	ErrorCodeSignInvalid                  ErrorCode = 0x08F00023 // Invalid XML signature.
	ErrorCodeNoSignFound                  ErrorCode = 0x08F00024 // Signature not found in input data.
	ErrorCodeDecrypt                      ErrorCode = 0x08F00025 // Decryption error.
	ErrorCodeXMLParse                     ErrorCode = 0x08F00026 // Unable to parse XML.
	ErrorCodeXMLAddID                     ErrorCode = 0x08F00027 // Failed to add ID attribute.
	ErrorCodeXMLInternal                  ErrorCode = 0x08F00028 // Error while working with XML.
	ErrorCodeXMLSetSign                   ErrorCode = 0x08F00029 // Failed to sign XML.
	ErrorCodeOpenSSL                      ErrorCode = 0x08F0002a // OpenSSL error.
	ErrorCodeNoTokenFound                 ErrorCode = 0x08F0002c // Token not found.
	ErrorCodeOCSPAddCert                  ErrorCode = 0x08F0002d // Failed to add certificate to OCSP.
	ErrorCodeOCSPParseURL                 ErrorCode = 0x08F0002e // Unable to parse URL.
	ErrorCodeOCSPAddHost                  ErrorCode = 0x08F0002f // Failed to add host.
	ErrorCodeOCSPReq                      ErrorCode = 0x08F00030 // Failed to add current time to request.
	ErrorCodeOCSPConnecting               ErrorCode = 0x08F00031 // Error connecting to OCSP responder.
	ErrorCodeVerifyNoData                 ErrorCode = 0x08F00032 // No input data for verification.
	ErrorCodeIDAttrNotFound               ErrorCode = 0x08F00033 // ID attribute not found.
	ErrorCodeIDRange                      ErrorCode = 0x08F00034 // Invalid identifier.
	ErrorCodeReaderNotFound               ErrorCode = 0x08F00037 // Reader not found.
	ErrorCodeGetCertProp                  ErrorCode = 0x08F00038 // Failed to get attribute value of certificate.
	ErrorCodeSignFormat                   ErrorCode = 0x08F00039 // Unknown signature format.
	ErrorCodeInDataFormat                 ErrorCode = 0x08F0003a // Unknown input data format.
	ErrorCodeOutDataFormat                ErrorCode = 0x08F0003b // Unknown output data format.
	ErrorCodeVerifyInit                   ErrorCode = 0x08F0003c // Unable to initialize signature verification manager.
	ErrorCodeVerify                       ErrorCode = 0x08F0003d // Failed to verify digital signature.
	ErrorCodeHash                         ErrorCode = 0x08F0003e // Failed to hash the data.
	ErrorCodeSignHash                     ErrorCode = 0x08F0003f // Failed to sign hashed data.
	ErrorCodeCACertsNotFound              ErrorCode = 0x08F00040 // The CA certificate was not found in the certificate store.
	ErrorCodeCertTimeInvalid              ErrorCode = 0x08F00042 // The certificate has expired or has not yet begin.
	ErrorCodeConvert                      ErrorCode = 0x08F00043 // Error writing certificate to X509 structure.
	ErrorCodeGenQueryTSA                  ErrorCode = 0x08F00044 // Error generating Timestamp request.
	ErrorCodeCreateObj                    ErrorCode = 0x08F00045 // Error writing OID to ASN.1 structure.
	ErrorCodeCreateNoNce                  ErrorCode = 0x08F00046 // Error generating unique number.
	ErrorCodeHTTP                         ErrorCode = 0x08F00047 // HTTP error.
	ErrorCodeCAdESBESFailed               ErrorCode = 0x08F00048 // Error checking CAdES-BES extension in CMS.
	ErrorCodeCAdESTFailed                 ErrorCode = 0x08F00049 // Error verifying TSA token signature.
	ErrorCodeNoTSAToken                   ErrorCode = 0x08F0004a // The signature does not contain a TSA mark.
	ErrorCodeInvalidHashLen               ErrorCode = 0x08F0004b // Invalid hash length.
	ErrorCodeGenRand                      ErrorCode = 0x08F0004c // Random number generation error.
	ErrorCodeSoapNS                       ErrorCode = 0x08F0004d // No SOAP message headers found.
	ErrorCodeGetPubKey                    ErrorCode = 0x08F0004e // Error exporting public key.
	ErrorCodeGetCertInfo                  ErrorCode = 0x08F0004f // Error retrieving certificate information.
	ErrorCodeFileRead                     ErrorCode = 0x08F00050 // Error reading file.
	ErrorCodeCheckHash                    ErrorCode = 0x08F00051 // The hash does not match.
	ErrorCodeZipExtract                   ErrorCode = 0x08F00052 // Unable to open archive.
	ErrorCodeNoManifestFile               ErrorCode = 0x08F00053 // Manifest file not found.
	ErrorCodeVerifyTSHash                 ErrorCode = 0x08F00054 // Failed to verify TS signature hash.
	ErrorCodeXADESTFailed                 ErrorCode = 0x08F00055 // XAdES-T signature verification error.
	ErrorCodeOCSPRespStatMalformedRequest ErrorCode = 0x08F00056 // Invalid request.
	ErrorCodeOCSPRespStatInternal         ErrorCode = 0x08F00057 // Internal error.
	ErrorCodeOCSPRespStatTryLater         ErrorCode = 0x08F00058 // Try later.
	ErrorCodeOCSPRespStatSigRequired      ErrorCode = 0x08F00059 // Request signature required.
	ErrorCodeOCSPRespStatUnauthorized     ErrorCode = 0x08F0005a // Request not authorized.
	ErrorCodeVerifyIssuerSerialV2         ErrorCode = 0x08F0005b // Failed to verify IssuerSerialV2 in XAdES.
	ErrorCodeOCSPCheckCertFromResp        ErrorCode = 0x08F0005c // Error verifying OCSP responder certificate.
	ErrorCodeExpiredCRL                   ErrorCode = 0x08F0005d // The certificate revocation list file has been expired.
	ErrorCodeLibraryNotInit               ErrorCode = 0x08F00101 // The library is not initialized.
	ErrorCodeEngineLoad                   ErrorCode = 0x08F00200 // Error connecting (loading) module (engine).
	ErrorCodeInData                       ErrorCode = 0x08F00300 // Incorrect input data.
	ErrorCodeCertStatusOK                 ErrorCode = 0x08F00400 // Certificate status is valid. Used when checking the certificate via OCSP. (not an error, an entry is made in the log).
	ErrorCodeCertStatusRevoked            ErrorCode = 0x08F00401 // Certificate status - revoked. Used when checking the certificate via OCSP.
	ErrorCodeCertStatusUnknown            ErrorCode = 0x08F00402 // Certificate status is unknown. Used when checking a certificate via OCSP. For example, the certificate issuer could not be established.
)

// Error codes map.
var errorLabels = map[ErrorCode]string{
	ErrorCodeOK:                           "No error.",
	ErrorCodeReadPKCS12:                   "Unable to read file format PKCS#12.",
	ErrorCodeOpenPKCS12:                   "Unable to open file format pkcs#12.",
	ErrorCodeInvalidExtID:                 "Invalid certificate extension identifier.",
	ErrorCodeBufferTooSmall:               "Buffer size is too small.",
	ErrorCodeCertParse:                    "Unable to parse the certificate.",
	ErrorCodeInvalidFlag:                  "Invalid flag.",
	ErrorCodeOpenFile:                     "Can't open file.",
	ErrorCodeInvalidPassword:              "Wrong password.",
	ErrorCodeAllocMemory:                  "Unable to allocate memory.",
	ErrorCodeCheckChain:                   "CA certificate or user certificate not found when checking chain.",
	ErrorCodeCertValidType:                "Invalid certificate validation type.",
	ErrorCodeBadCRLFormat:                 "Incorrect CRL format.",
	ErrorCodeLoadCRL:                      "Unable to load CRL.",
	ErrorCodeLoadCRLs:                     "Unable to load CRLs.",
	ErrorCodeUnknownAlg:                   "Unknown signature algorithm.",
	ErrorCodeKeyNotFound:                  "User's private key not found.",
	ErrorCodeSignInit:                     "Unable to initialize signature manager.",
	ErrorCodeSign:                         "Failed to generate digital signature.",
	ErrorCodeEncrypt:                      "Encryption error.",
	ErrorCodeInvalidFlags:                 "Invalid flags.",
	ErrorCodeCertNotFound:                 "User certificate not found.",
	ErrorCodeVerifySign:                   "Error verifying XML signature",
	ErrorCodeBase64Decrypt:                "Error decrypting from Base64.",
	ErrorCodeUnknownCMSFormat:             "Unknown CMS format.",
	ErrorCodeCACertNotFound:               "CA certificate not found.",
	ErrorCodeXMLSecInit:                   "Error initializing XMLSec.",
	ErrorCodeLoadTrustedCerts:             "Error loading trusted certificates.",
	ErrorCodeSignInvalid:                  "Invalid XML signature.",
	ErrorCodeNoSignFound:                  "Signature not found in input data.",
	ErrorCodeDecrypt:                      "Decryption error.",
	ErrorCodeXMLParse:                     "Unable to parse XML.",
	ErrorCodeXMLAddID:                     "Failed to add ID attribute.",
	ErrorCodeXMLInternal:                  "Error while working with XML.",
	ErrorCodeXMLSetSign:                   "Failed to sign XML.",
	ErrorCodeOpenSSL:                      "OpenSSL error.",
	ErrorCodeNoTokenFound:                 "Token not found.",
	ErrorCodeOCSPAddCert:                  "Failed to add certificate to OCSP.",
	ErrorCodeOCSPParseURL:                 "Unable to parse URL.",
	ErrorCodeOCSPAddHost:                  "Failed to add host.",
	ErrorCodeOCSPReq:                      "Failed to add current time to request.",
	ErrorCodeOCSPConnecting:               "Error connecting to OCSP responder.",
	ErrorCodeVerifyNoData:                 "No input data for verification.",
	ErrorCodeIDAttrNotFound:               "ID attribute not found.",
	ErrorCodeIDRange:                      "Invalid identifier.",
	ErrorCodeReaderNotFound:               "Reader not found.",
	ErrorCodeGetCertProp:                  "Failed to get attribute value of certificate.",
	ErrorCodeSignFormat:                   "Unknown signature format",
	ErrorCodeInDataFormat:                 "Unknown input data format.",
	ErrorCodeOutDataFormat:                "Unknown output data format.",
	ErrorCodeVerifyInit:                   "Unable to initialize signature verification manager.",
	ErrorCodeVerify:                       "Failed to verify digital signature.",
	ErrorCodeHash:                         "Failed to hash the data.",
	ErrorCodeSignHash:                     "Failed to sign hashed data.",
	ErrorCodeCACertsNotFound:              "The CA certificate was not found in the certificate store.",
	ErrorCodeCertTimeInvalid:              "The certificate has expired or has not yet begin.",
	ErrorCodeConvert:                      "Error writing certificate to X509 structure.",
	ErrorCodeGenQueryTSA:                  "Error generating Timestamp request.",
	ErrorCodeCreateObj:                    "Error writing OID to ASN.1 structure.",
	ErrorCodeCreateNoNce:                  "Error generating unique number.",
	ErrorCodeHTTP:                         "HTTP error.",
	ErrorCodeCAdESBESFailed:               "Error checking CAdES-BES extension in CMS.",
	ErrorCodeCAdESTFailed:                 "Error verifying TSA token signature.",
	ErrorCodeNoTSAToken:                   "The signature does not contain a TSA mark.",
	ErrorCodeInvalidHashLen:               "Invalid hash length.",
	ErrorCodeGenRand:                      "Random number generation error.",
	ErrorCodeSoapNS:                       "No SOAP message headers found.",
	ErrorCodeGetPubKey:                    "Error exporting public key.",
	ErrorCodeGetCertInfo:                  "Error retrieving certificate information.",
	ErrorCodeFileRead:                     "Error reading file.",
	ErrorCodeCheckHash:                    "The hash does not match.",
	ErrorCodeZipExtract:                   "Unable to open archive.",
	ErrorCodeNoManifestFile:               "Manifest file not found.",
	ErrorCodeVerifyTSHash:                 "Failed to verify TS signature hash.",
	ErrorCodeXADESTFailed:                 "XAdES-T signature verification error.",
	ErrorCodeOCSPRespStatMalformedRequest: "Invalid request.",
	ErrorCodeOCSPRespStatInternal:         "Internal error.",
	ErrorCodeOCSPRespStatTryLater:         "Try later.",
	ErrorCodeOCSPRespStatSigRequired:      "Request signature required.",
	ErrorCodeOCSPRespStatUnauthorized:     "Request not authorized.",
	ErrorCodeVerifyIssuerSerialV2:         "Failed to verify IssuerSerialV2 in XAdES.",
	ErrorCodeOCSPCheckCertFromResp:        "Error verifying OCSP responder certificate.",
	ErrorCodeExpiredCRL:                   "The certificate revocation list file has been expired.",
	ErrorCodeLibraryNotInit:               "The library is not initialized.",
	ErrorCodeEngineLoad:                   "Error connecting (loading) module (engine).",
	ErrorCodeInData:                       "Incorrect input data.",
	ErrorCodeCertStatusOK:                 "Certificate status is valid. Used when checking the certificate via OCSP. (not an error, an entry is made in the log).",
	ErrorCodeCertStatusRevoked:            "Certificate status - revoked. Used when checking the certificate via OCSP.",
	ErrorCodeCertStatusUnknown:            "Certificate status is unknown. Used when checking a certificate via OCSP. For example, the certificate issuer could not be established.",
}
