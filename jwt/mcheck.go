package jwt

// The method by which an organization has obtained confirmation of access
// to personal data from a personal data subject.
type MCheck string

const (
	BIO MCheck = "BIO" // Organizational biometrics system.
	DS  MCheck = "DS"  // Digital signature.
	DID MCheck = "DID" // Digital ID system.
	OTP MCheck = "OTP" // One-time password.
	PC  MCheck = "PC"  // Paper information carrier.
)
