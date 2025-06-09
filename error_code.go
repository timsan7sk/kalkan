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

const (
	ErrorCodeOK                           ErrorCode = 0         // No error.
	ErrorCodeReadPKCS12                   ErrorCode = 149946370 // Unable to read file format pkcs#12.
	ErrorCodeOpenPKCS12                   ErrorCode = 149946371 // Unable to open file format pkcs#12.
	ErrorCodeInvalidExtID                 ErrorCode = 149946372 // Invalid certificate extension identifier.
	ErrorCodeBufferTooSmall               ErrorCode = 149946373 // Buffer size is too small.
	ErrorCodeCertParse                    ErrorCode = 149946374 // Unable to parse the certificate.
	ErrorCodeInvalidFlag                  ErrorCode = 149946375 // Invalid flag.
	ErrorCodeOpenFile                     ErrorCode = 149946376 // Can't open file.
	ErrorCodeInvalidPassword              ErrorCode = 149946377 // Wrong password.
	ErrorCodeAllocMemory                  ErrorCode = 149946381 // Unable to allocate memory
	ErrorCodeCheckChain                   ErrorCode = 149946382 // CA certificate or user certificate not found when checking chain.
	ErrorCodeCertValidType                ErrorCode = 149946384 // Invalid certificate validation type.
	ErrorCodeBadCRLFormat                 ErrorCode = 149946385 // Incorrect CRL format.
	ErrorCodeLoadCRL                      ErrorCode = 149946386 // Unable to load CRL.
	ErrorCodeLoadCRLs                     ErrorCode = 149946387 // Unable to load CRLs.
	ErrorCodeUnknownAlg                   ErrorCode = 149946389 // Unknown signature algorithm.
	ErrorCodeKeyNotFound                  ErrorCode = 149946390 // User's private key not found.
	ErrorCodeSignInit                     ErrorCode = 149946391 // Unable to initialize signature manager.
	ErrorCodeSign                         ErrorCode = 149946392 // Failed to generate digital signature.
	ErrorCodeEncrypt                      ErrorCode = 149946393 // Encryption error.
	ErrorCodeInvalidFlags                 ErrorCode = 149946394 // Invalid flags.
	ErrorCodeCertNotFound                 ErrorCode = 149946395 // User certificate not found.
	ErrorCodeVerifySign                   ErrorCode = 149946396 // Error verifying XML signature.
	ErrorCodeBase64Decrypt                ErrorCode = 149946397 // Error decripting from Base64.
	ErrorCodeUnknownCMSFormat             ErrorCode = 149946398 // Unknown CMS format.
	ErrorCodeCACertNotFound               ErrorCode = 149946400 // CA certificate not found.
	ErrorCodeXMLSecInit                   ErrorCode = 149946401 // Error initializing XMLSec.
	ErrorCodeLoadTrustedCerts             ErrorCode = 149946402 // Error loading trusted certificates.
	ErrorCodeSignInvalid                  ErrorCode = 149946403 // Invalid XML signature.
	ErrorCodeNoSignFound                  ErrorCode = 149946404 // Signature not found in input data.
	ErrorCodeDecrypt                      ErrorCode = 149946405 // Decryption error.
	ErrorCodeXMLParse                     ErrorCode = 149946406 // Unable to parse XML.
	ErrorCodeXMLAddID                     ErrorCode = 149946407 // Failed to add ID attribute.
	ErrorCodeXMLInternal                  ErrorCode = 149946408 // Error while working with XML.
	ErrorCodeXMLSetSign                   ErrorCode = 149946409 // Failed to sign XML.
	ErrorCodeOpenSSL                      ErrorCode = 149946410 // OpenSSL error.
	ErrorCodeNoTokenFound                 ErrorCode = 149946412 // Token not found.
	ErrorCodeOCSPAddCert                  ErrorCode = 149946413 // Failed to add certificate to OCSP.
	ErrorCodeOCSPParseURL                 ErrorCode = 149946414 // Unable to parse URL.
	ErrorCodeOCSPAddHost                  ErrorCode = 149946415 // Failed to add host.
	ErrorCodeOCSPReq                      ErrorCode = 149946416 // Failed to add current time to request.
	ErrorCodeOCSPConnecting               ErrorCode = 149946417 // Error connecting to OCSP responder.
	ErrorCodeVerifyNoData                 ErrorCode = 149946418 // No input data for verification.
	ErrorCodeIDAttrNotFound               ErrorCode = 149946419 // ID attribute not found.
	ErrorCodeIDRange                      ErrorCode = 149946420 // Invalid identifier.
	ErrorCodeReaderNotFound               ErrorCode = 149946423 // Reader not found.
	ErrorCodeGetCertProp                  ErrorCode = 149946424 // Failed to get attribute value of certificate.
	ErrorCodeSignFormat                   ErrorCode = 149946425 // Unknown signature format.
	ErrorCodeInDataFormat                 ErrorCode = 149946426 // Unknown input data format.
	ErrorCodeOutDataFormat                ErrorCode = 149946427 // Unknown output format.
	ErrorCodeVerifyInit                   ErrorCode = 149946428 // Unable to initialize signature verification manager.
	ErrorCodeVerify                       ErrorCode = 149946429 // Failed to verify digital signature.
	ErrorCodeHash                         ErrorCode = 149946430 // Failed to hash the data.
	ErrorCodeSignHash                     ErrorCode = 149946431 // Failed to sign hashed data.
	ErrorCodeCACertsNotFound              ErrorCode = 149946432 // The CA certificate was not found in the certificate store.
	ErrorCodeCertTimeInvalid              ErrorCode = 149946434 // The certificate has expired or has not yet begin.
	ErrorCodeConvert                      ErrorCode = 149946435 // Error writing certificate to X509 structure.
	ErrorCodeGenQueryTSA                  ErrorCode = 149946436 // Error generating Timestamp request.
	ErrorCodeCreateObj                    ErrorCode = 149946437 // Ошибка записи OID в ASN1 структуру
	ErrorCodeCreateNoNce                  ErrorCode = 149946438 // Error generating unique number.
	ErrorCodeHTTP                         ErrorCode = 149946439 // HTTP error.
	ErrorCodeCAdESBESFailed               ErrorCode = 149946440 // Error checking CAdES-BES extension in CMS.
	ErrorCodeCAdESTFailed                 ErrorCode = 149946441 // Error verifying TSA token signature.
	ErrorCodeNoTSAToken                   ErrorCode = 149946442 // The signature does not contain a TSA mark.
	ErrorCodeInvalidHashLen               ErrorCode = 149946443 // Invalid hash length.
	ErrorCodeGenRand                      ErrorCode = 149946444 // Random number generation error.
	ErrorCodeSoapNS                       ErrorCode = 149946445 // No SOAP message headers found.
	ErrorCodeGetPubKey                    ErrorCode = 149946446 // Error exporting public key.
	ErrorCodeGetCertInfo                  ErrorCode = 149946447 // Error retrieving certificate information.
	ErrorCodeFileRead                     ErrorCode = 149946448 // Error reading file.
	ErrorCodeCheckHash                    ErrorCode = 149946449 // The hash does not match.
	ErrorCodeZipExtract                   ErrorCode = 149946450 // Unable to open archive.
	ErrorCodeNoManifestFile               ErrorCode = 149946451 // Manifest file not found.
	ErrorCodeVerifyTSHash                 ErrorCode = 149946452 // Failed to verify TS signature hash.
	ErrorCodeXADESTFailed                 ErrorCode = 149946453 // XAdES-T signature verification error.
	ErrorCodeOCSPRespStatMalformedRequest ErrorCode = 149946454 // Invalid request.
	ErrorCodeOCSPRespStatInternal         ErrorCode = 149946455 // Internal error.
	ErrorCodeOCSPRespStatTryLater         ErrorCode = 149946456 // Try later.
	ErrorCodeOCSPRespStatSigRequired      ErrorCode = 149946457 // Request signature required.
	ErrorCodeOCSPRespStatUnauthorized     ErrorCode = 149946458 // Request not authorized.
	ErrorCodeVerifyIssuerSerialV2         ErrorCode = 149946459 // Failed to verify IssuerSerialV2 in XAdES.
	ErrorCodeOCSPCheckCertFromResp        ErrorCode = 149946460 // Error verifying OCSP responder certificate.
	ErrorCodeExpiredCRL                   ErrorCode = 149946461 // Certificate Revocation List has expared.
	ErrorCodeLibraryNotInitialized        ErrorCode = 149946625 // The library is not initialized.
	ErrorCodeEngineLoad                   ErrorCode = 149946880 // Error connecting (loading) module (engine).
	ErrorCodeInData                       ErrorCode = 149947136 // Incorrect input data.
	ErrorCodeCertStatusOK                 ErrorCode = 149947392 // Certificate status is valid. Used when checking the certificate via OCSP. (not an error, an entry is made in the log).
	ErrorCodeCertStatusRevoked            ErrorCode = 149947393 // Certificate status - revoked. Used when checking the certificate via OCSP.
	ErrorCodeCertStatusUnknown            ErrorCode = 149947394 // Certificate status is unknown. Used when checking a certificate via OCSP. For example, the certificate issuer could not be established.
)

var errorLabels = map[ErrorCode]string{
	ErrorCodeOK:                           "Нет ошибки",
	ErrorCodeReadPKCS12:                   "Невозможно прочитать файл формата pkcs#12",
	ErrorCodeOpenPKCS12:                   "Невозможно открыть файл формата pkcs12",
	ErrorCodeInvalidExtID:                 "Недопустимый идентификатор расширения сертификата",
	ErrorCodeBufferTooSmall:               "Размер буфера слишком мал",
	ErrorCodeCertParse:                    "Невозможно разобрать (распарсить) сертификат",
	ErrorCodeInvalidFlag:                  "Недопустимый флаг",
	ErrorCodeOpenFile:                     "Невозможно открыть файл",
	ErrorCodeInvalidPassword:              "Неправильный пароль",
	ErrorCodeAllocMemory:                  "Невозможно выделить память",
	ErrorCodeCheckChain:                   "Не найден сертификат УЦ или сертификат пользователя при проверки цепочки",
	ErrorCodeCertValidType:                "Недопустимый тип валидации сертификата",
	ErrorCodeBadCRLFormat:                 "Некорректный формат CRL",
	ErrorCodeLoadCRL:                      "Невозможно загрузить CRL",
	ErrorCodeLoadCRLs:                     "Невозможно загрузить CRL-ы",
	ErrorCodeUnknownAlg:                   "Неизвестный алгоритм подписи",
	ErrorCodeKeyNotFound:                  "Не найден приватный ключ пользователя",
	ErrorCodeSignInit:                     "Невозможно инициализировать менеджера подписи",
	ErrorCodeSign:                         "Не удалось сгенерировать цифровую подпись",
	ErrorCodeEncrypt:                      "Ошибка шифрования",
	ErrorCodeInvalidFlags:                 "Недопустимые флаги",
	ErrorCodeCertNotFound:                 "Не найден сертификат пользователя",
	ErrorCodeVerifySign:                   "Ошибка верификации подписи xml",
	ErrorCodeBase64Decrypt:                "Ошибка дешифровки из Base 64",
	ErrorCodeUnknownCMSFormat:             "Неизвестный формат CMS",
	ErrorCodeCACertNotFound:               "Не найден сертификат УЦ",
	ErrorCodeXMLSecInit:                   "Ошибка инициализации xmlsec",
	ErrorCodeLoadTrustedCerts:             "Ошибка загрузки доверенных сертификатов",
	ErrorCodeSignInvalid:                  "Недопустимая подпись xml",
	ErrorCodeNoSignFound:                  "Не найдена подпись во входных данных",
	ErrorCodeDecrypt:                      "Ошибка дешифрования",
	ErrorCodeXMLParse:                     "Невозможно разобрать (распарсить) xml",
	ErrorCodeXMLAddID:                     "Не удалось добавить атрибут ID",
	ErrorCodeXMLInternal:                  "Ошибка при работе с xml",
	ErrorCodeXMLSetSign:                   "Не удалось подписать xml",
	ErrorCodeOpenSSL:                      "Ошибка openssl",
	ErrorCodeNoTokenFound:                 "Не найден токен",
	ErrorCodeOCSPAddCert:                  "Не удалось добавить сертификат в ocsp",
	ErrorCodeOCSPParseURL:                 "Не удалось разобрать url",
	ErrorCodeOCSPAddHost:                  "Не удалось добавить хост",
	ErrorCodeOCSPReq:                      "Не удалось добавить текущее время в запрос",
	ErrorCodeOCSPConnecting:               "Ошибка подключения к OCSP респондеру",
	ErrorCodeVerifyNoData:                 "Нет входных данных для верификации",
	ErrorCodeIDAttrNotFound:               "Не найден атрибут ID",
	ErrorCodeIDRange:                      "Некорректный идентификатор",
	ErrorCodeReaderNotFound:               "Не найден ридер",
	ErrorCodeGetCertProp:                  "Не удалось получить значение атрибута",
	ErrorCodeSignFormat:                   "Неизвестный формат подписи",
	ErrorCodeInDataFormat:                 "Неизвестный формат входных данных",
	ErrorCodeOutDataFormat:                "Неизвестный формат выходных данных",
	ErrorCodeVerifyInit:                   "Невозможно инициализировать менеджера верификации подписи",
	ErrorCodeVerify:                       "Не удалось верифицировать цифровую подпись",
	ErrorCodeHash:                         "Не удалось хэшировать данные",
	ErrorCodeSignHash:                     "Не удалось подписать хэшированные данные",
	ErrorCodeCACertsNotFound:              "Не найден сертификат УЦ в хранилище сертификатов",
	ErrorCodeCertTimeInvalid:              "Срок действия сертификата истек либо еще не наступил",
	ErrorCodeConvert:                      "Ошибка записи сертификата в структуру X509",
	ErrorCodeGenQueryTSA:                  "Ошибка генерации запроса timestamp",
	ErrorCodeCreateObj:                    "Ошибка записи OID в ASN1 структуру",
	ErrorCodeCreateNoNce:                  "Ошибка генерации уникального числа",
	ErrorCodeHTTP:                         "Ошибка протокола http",
	ErrorCodeCAdESBESFailed:               "Ошибка проверки расширения CADESBES в CMS",
	ErrorCodeCAdESTFailed:                 "Ошибка проверки подписи токена TSA",
	ErrorCodeNoTSAToken:                   "В подписи не присутствует метка TSA",
	ErrorCodeInvalidHashLen:               "Неправильная длина хэша",
	ErrorCodeGenRand:                      "Ошибка генерации случайного числа",
	ErrorCodeSoapNS:                       "Не найдены заголовки SOAP-сообщений",
	ErrorCodeGetPubKey:                    "Ошибка экспорта публичного ключа",
	ErrorCodeGetCertInfo:                  "Ошибка получения информации о сертификате",
	ErrorCodeFileRead:                     "Ошибка чтения файла",
	ErrorCodeCheckHash:                    "Хэш не совпадает",
	ErrorCodeZipExtract:                   "Невозможно открыть архив",
	ErrorCodeNoManifestFile:               "Не найден MANIFEST",
	ErrorCodeVerifyTSHash:                 "не удалось проверить Хэш подписи TS",
	ErrorCodeXADESTFailed:                 "XAdES-T: Ошибка проверки подписи",
	ErrorCodeOCSPRespStatMalformedRequest: "Неправильный запрос",
	ErrorCodeOCSPRespStatInternal:         "Внутренняя ошибка",
	ErrorCodeOCSPRespStatTryLater:         "Попробуйте позже",
	ErrorCodeOCSPRespStatSigRequired:      "Должны подписать запрос",
	ErrorCodeOCSPRespStatUnauthorized:     "Запрос не авторизован",
	ErrorCodeVerifyIssuerSerialV2:         "не удалось проверить IssuerSerialV2 в XAdES",
	ErrorCodeOCSPCheckCertFromResp:        "Ошибка проверки сертификата OCSP-респондера",
	ErrorCodeExpiredCRL:                   "CRL-файл просрочен",
	ErrorCodeLibraryNotInitialized:        "Библиотека не инициализирована",
	ErrorCodeEngineLoad:                   "Ошибка подключения (загрузки) модуля (engine)",
	ErrorCodeInData:                       "Некорректные входные данные",
	ErrorCodeCertStatusOK:                 "Статус сертификата – валидный. Используется при проверке сертификата по OCSP. (не является ошибкой, делается запись в лог)",
	ErrorCodeCertStatusRevoked:            "Статус сертификата – отозван. Используется при проверке сертификата по OCSP.",
	ErrorCodeCertStatusUnknown:            "Статус сертификата – неизвестен. Используется при проверке сертификата по OCSP. Например, не удалось установить издателя сертификата.",
}
