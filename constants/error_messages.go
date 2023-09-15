package constants

var ErrorMessage = struct {
	AUTHENTICATION_REQUIRED                string
	INVALID_AUTHENTICATION                 string
	SERVICE_PROVIDER_AUTHENTICATION_FAILED string
	JSON_UNKNOWN_FIELD                     string
	JSON_CANNOT_UNMARSHAL                  string
}{
	AUTHENTICATION_REQUIRED:                "authentication required",
	INVALID_AUTHENTICATION:                 "invalid authentication",
	SERVICE_PROVIDER_AUTHENTICATION_FAILED: "service provider authentication failed",
	JSON_UNKNOWN_FIELD:                     "json: unknown field",
	JSON_CANNOT_UNMARSHAL:                  "json: cannot unmarshal",
}
