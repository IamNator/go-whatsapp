package go_whatsapp

import "time"

const (
	// PackageName name of this package
	PackageName = "go_whatsapp"

	// DefaultAPIURL default api url
	DefaultAPIURL = "https://api.whatsapp.com"

	// DefaultRateLimit default rate limit
	DefaultRateLimit = 200 // 200 requests per second -> max is about 500/second

	// DefaultTimeout default timeout
	DefaultTimeout = 3 * time.Second
)

// ---------------------------------------- API VERSION  ------

type (
	MetaAPIVersion string
)

const (
	V13 MetaAPIVersion = "v13.0"
	V14 MetaAPIVersion = "v14.0"

	V15 MetaAPIVersion = "v15.0" // 2022
	V16 MetaAPIVersion = "v16.0" // 2023
)

func (m MetaAPIVersion) String() string {
	return string(m)
}
