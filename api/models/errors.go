package models

import (
	"fmt"
)

// http error type
type HTTPError struct {
	Code    int    `json:"-"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e HTTPError) Error() string {
	return e.Message
}

// When req body has wrong format
type FormatValidationError struct {
	Message string
}

func (e FormatValidationError) Error() string {
	return e.Message
}

// request format correct, but data is invalid
type DataValidationError struct {
	Message string
}

func (e DataValidationError) Error() string {
	return e.Message
}

type InvalidJSONError struct {
	Message string
}

func (e InvalidJSONError) Error() string {
	return e.Message
}

// resource or route is not found
type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	if e.Message == "" {
		return "resource not found"
	}
	return e.Message
}

// custom area type
func WrapError(customErr string, originalErr error) error {
	err := fmt.Errorf("%s: %v", customErr, originalErr)
	return err
}
