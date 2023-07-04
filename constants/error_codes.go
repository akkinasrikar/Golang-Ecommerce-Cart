package constants

var ErrorCode = struct {
	INTERNAL_SYSTEM_ERROR        string
	INVALID_REQUEST              string
	PARAMETER_MISSING_OR_INVALID string
	SERVICE_PROVIDER_ERROR       string
	PARAMETER_UNKNOWN            string
	HEADER_MISSING_OR_INVALID    string
}{
	INTERNAL_SYSTEM_ERROR:        "internal_system_error",
	INVALID_REQUEST:              "invalid_request",
	PARAMETER_MISSING_OR_INVALID: "parameter_missing_or_invalid",
	SERVICE_PROVIDER_ERROR:       "service_provider_error",
	PARAMETER_UNKNOWN:            "parameter_unknown",
	HEADER_MISSING_OR_INVALID:    "header_missing_or_invalid",
}
