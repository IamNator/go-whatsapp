package go_whatsapp

import "time"

const (
	// PackageName name of this package
	PackageName = "go_whatsapp"

	// DefaultAPIVersion default api version
	DefaultAPIVersion = "v14"

	// DefaultAPIURL default api url
	DefaultAPIURL = "https://api.whatsapp.com"

	// DefaultRateLimit default rate limit
	DefaultRateLimit = 200 // 200 requests per second -> max is about 500/second

	// DefaultTimeout default timeout
	DefaultTimeout = 2 * time.Second
)

// ---------------------------------------- API VERSION

type (
	MetaAPIVersion string
)

const (
	V13 MetaAPIVersion = "13.0" //previous version
	V14 MetaAPIVersion = "14.0" // previous version

	V15 MetaAPIVersion = "15.0" //latest version
	V16 MetaAPIVersion = "16.0" //reserved
)

func (m MetaAPIVersion) String() string {
	return string(m)
}
