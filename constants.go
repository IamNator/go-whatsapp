package go_whatsapp

import "time"

type Constant interface {
	string | int | int64 | float64 | bool
}

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
