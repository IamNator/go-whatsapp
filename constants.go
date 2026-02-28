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

type APIVersion string

const (
	V13 APIVersion = "v13.0" // Graph API v13.0 (2020)
	V14 APIVersion = "v14.0" // Graph API v14.0 (2021)

	V15 APIVersion = "v15.0" // Graph API v15.0 (2022) – WhatsApp Cloud API first supported w/ this major version
	V16 APIVersion = "v16.0" // Graph API v16.0 (2023)
	V17 APIVersion = "v17.0" // Graph API v17.0 (2023-2024)
	V18 APIVersion = "v18.0" // Graph API v18.0 (2024)
	V19 APIVersion = "v19.0" // Graph API v19.0 (2024)
	V20 APIVersion = "v20.0" // Graph API v20.0 (2024)
	V21 APIVersion = "v21.0" // Graph API v21.0 (2024-2025) – current stable for Cloud API as of early 2026
	V22 APIVersion = "v22.0" // Graph API v22.0 (2025)
	V23 APIVersion = "v23.0" // Graph API v23.0 (2025)
	V24 APIVersion = "v24.0" // Graph API v24.0 (2025)
	V25 APIVersion = "v25.0" // Graph API v25.0 (early 2026)
)

func (m APIVersion) String() string {
	return string(m)
}
