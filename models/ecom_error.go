package models

import "fmt"

type EcomError struct {
	Code    int64  `json:"code"`
	Message error  `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param,omitempty"`
}

func (r *EcomError) Error() string {
	return fmt.Sprintf("%v", r.Message)
}

type EcomErrorBody struct {
	Type    string `json:"type"`
	Code    int64  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Param   string `json:"param,omitempty"`
	TraceId string `json:"traceId,omitempty"`
}

type EcomErrorResponse struct {
	ErrorType EcomErrorBody `json:"error"`
}
