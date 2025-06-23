|[![NCA](nca_icon.png)](https://pki.gov.kz/)| **Go wrapper for the [KalkanCrypt](https://sdk.pki.gov.kz//) library.** |
|:-----------------------------------------------------:|:---------------------------------------------------------------|

#### Functionality ####


#### Example ####

```go
func init() {
	mod, err := kalkan.Open("libkalkancryptwr-64.so.2")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := mod.Init(); err != nil {
		fmt.Println(err)
	}
}

```

| KC_Function_List					| C | Go | Test | Description																									|
|:----------------------------------|:-:|:--:|:----:|:--------------------------------------------------------------------------------------------------------------|
| KC_Init							| + | +  | +    | Initializes the library. 																						|
| KC_InitDebug						| + | +  | +    | Enables writing the library's log to a file.																	|
| KC_GetTokens						| + | +  | -    | Obtains a pointer to a string of connected storage devices and their number.									|
| KC_GetCertificatesList			| + | +  | +    | Obtains list of certificates as a string and their number.													|
| KC_LoadKeyStore					| + | +  | -    | Loads keys/certificates from storage.																			|
| X509LoadCertificateFromFile		| + | +  | -    | Loads a certificate from a file.																				|
| X509LoadCertificateFromBuffer		| + | +  | -    | Loads a certificate from memory.																				|
| X509ExportCertificateFromStore	| + | +  | -    | Exports a certificate from the storage.																		|
| X509CertificateGetInfo    		| + | +  | -    | Obtains field/extension values ​​from the certificate.														  |
| X509ValidateCertificate			| + | +  | -    | Performs certificate verification, validity check, certificate chain building, OCSP or CRL revocation check.	|
| HashData							| + | +  | -    | Hashes data.																									|
| SignData							| + | +  | +	| Signs the data.																								|
| SignDataArchive					| - | -  | -    | |
| SignHash							| - | -  | -    | Signs the input hashed data.																					|
| SignWSSE							| + | +  | -    | Signs a SOAP message according to the WS-Security specification.												|
| SignXML							| + | +  | -    | Signs XML file.																								|
| VerifyData						| + | +  | -    | Provides signature verification.																				|
| VerifyXML							| + | +  | -    | Provides signature verification of XML data.																	|
| UVerifyData						| + | +  | -    | Provides signature verification.																				|
| KC_GetLastError					| + | +  | -    | Obtains error code of the functions operation.|
| KC_GetLastErrorString				| + | +  | -    | Obtains protocol of the functions operation.																	|
| KC_getSigAlgFromXML				| - | -  | -    | Obtains XML signature algorithm.																				|
| KC_getCertFromXML					| + | +  | -    | Obtains certificate from XML.																					|
| KC_XMLFinalize					| + | +  | -    | Frees memory and terminates the library with modules responsible for parsing, signing and verifying XML data. |
| KC_Finalize						| + | +  | -    | Frees resources occupied by the KalkanCrypt crypto provider and terminates the library.						|
| TSASetUrl							| + | +  | -    | Sets TSA service addresses.																					|
| KC_GetTimeFromSig					| + | +  | -    | Obtains the signature time.																					|
| KC_SetProxy						| + | +  | -    | Sets proxy server settings.																					|
| KC_InsertTStoCMS					| + | +  | -    | |
| KC_GetSignInfosFromDoc			| - | -  | -    | |
| KC_CreateCMSfromDraftSign			| - | -  | -    | |
| KC_GetTSfromDraftSign				| - | -  | -    | |
| KC_CreateXAdESfromCloudStorage	| - | -  | -    | |
| KC_GenerateX509ReqFL				| - | -  | -    | |
| KC_AddSignatureToX509Req			| - | -  | -    | |
| ZipConSign						| - | -  | -    | Signs a file and then places it in a zip container.															|
| ZipConVerify						| - | -  | -    | Verifies the signature of a zip container.																	|
| KC_GetCertFromCMS					| + | +  | -    | Obtains certificate from CMS format data.																		|
| KC_getCertFromZipFile				| - | -  | -    | Obtains certificate from ZIP file.																			|
