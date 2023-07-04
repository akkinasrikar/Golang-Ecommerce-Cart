package constants

type ErrorValue struct {
	Code int
	Name string
}

var ErrorType = struct {
	INVALID_REQUEST_ERROR                 ErrorValue
	DOWNSTREAM_ERROR                      ErrorValue
	INTERNAL_SYSTEM_ERROR                 ErrorValue
	AUTHENTICATION_ERROR                  ErrorValue
	SERVICE_PROVIDER_ERROR                ErrorValue
	SERVICE_PROVIDER_AUTHENTICATION_ERROR ErrorValue
}{
	INVALID_REQUEST_ERROR:                 ErrorValue{422, "invalid_request_error"},
	DOWNSTREAM_ERROR:                      ErrorValue{550, "downstream_error"},
	INTERNAL_SYSTEM_ERROR:                 ErrorValue{500, "internal_system_error"},
	AUTHENTICATION_ERROR:                  ErrorValue{401, "authentication_error"},
	SERVICE_PROVIDER_ERROR:                ErrorValue{550, "service_provider_error"},
	SERVICE_PROVIDER_AUTHENTICATION_ERROR: ErrorValue{550, "service_provider_authentication_error"},
}
