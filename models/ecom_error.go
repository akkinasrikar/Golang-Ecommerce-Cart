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
