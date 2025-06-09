package kalkan

// Defines the type of storage/media.
type StoreType int

const (
	StoreTypePKCS12     StoreType = 0x00000001 // File system.
	StoreTypeKZIDCard   StoreType = 0x00000002 // Identity card of a citizen of the Republic of Kazakhstan.
	StoreTypeKazToken   StoreType = 0x00000004 // KazToken.
	StoreTypeEToken     StoreType = 0x00000008 // eToken 72k.
	StoreTypeJaCarta    StoreType = 0x00000010 // JaCarta.
	StoreTypeX509Cert   StoreType = 0x00000020 // X509 Certificate.
	StoreTypeAKey       StoreType = 0x00000040 // aKey.
	StoreTypeEToken5110 StoreType = 0x00000080 // eToken 5110.
)
