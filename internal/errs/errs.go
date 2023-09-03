package errs

import (
	"fmt"
	"strconv"
	"strings"
)

type ErrCode string

// Prefix
const (
	GenericError    ErrCode = "1"
	HandlerError    ErrCode = "2"
	ServiceError    ErrCode = "3"
	ValidationError ErrCode = "4"
)

// Type
const (
	ErrUnknown            ErrCode = "000"
	ErrDecoder            ErrCode = "001"
	ErrRequiredValidation ErrCode = "002"
	ErrInvalidValidation  ErrCode = "003"
)

func GenerateErrorCode(prefix ErrCode, errType ErrCode) int {
	combinedCode := fmt.Sprintf("%s%s", prefix, errType)
	code, err := strconv.Atoi(combinedCode)
	if err != nil {
		return -1
	}

	if code > 0 {
		return code
	}

	return 1000
}

type RequiredFieldError struct {
	RequiredFields []string
}

func (e RequiredFieldError) Error() string {
	return fmt.Sprintf("The following fields is required: %s", strings.Join(e.RequiredFields[:], ","))
}

type InvalidFieldError struct {
	InvalidFields []string
}

func (e InvalidFieldError) Error() string {
	return fmt.Sprintf("The following fields is invalid: %s", strings.Join(e.InvalidFields[:], ","))
}

type InternalError struct {
	Message string
	Code    int
}

func (e InternalError) Error() string {
	return e.Message
}

type InvalidLoginErr struct {
	Message string
}

func (e InvalidLoginErr) Error() string {
	return e.Message
}
