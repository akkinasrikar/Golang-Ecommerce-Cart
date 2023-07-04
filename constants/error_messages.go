package constants

var ErrorMessage = struct {
	AUTHENTICATION_REQUIRED                string
	INVALID_AUTHENTICATION                 string
	SERVICE_PROVIDER_AUTHENTICATION_FAILED string
}{
	AUTHENTICATION_REQUIRED:                "authentication required",
	INVALID_AUTHENTICATION:                 "invalid authentication",
	SERVICE_PROVIDER_AUTHENTICATION_FAILED: "service provider authentication failed",
}
