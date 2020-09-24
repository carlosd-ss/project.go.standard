package gjwt

import "crypto/rsa"

var (

	// default, we will load
	// the keys into memory
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey

	// only if it will be used by files
	PathPrivate = "./private.rsa"
	PathPublic  = "./public.rsa.pub"

	ProjectTitle = "jwt web"

	ExpirationHours = 22  // Hours
	DayExpiration   = 200 // Days
)

// Structure of our server configurations
type JsonMsg struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
