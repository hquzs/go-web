package util

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Error is the core.Error
type Error struct {
	Err  error
	Code string
}

// NewError is the constructor of Error
// Codes can be omit, and if you set many codes, only the first one is useful
func NewError(err error, codes ...string) *Error {
	if len(codes) == 0 {
		code, e := parseCodeAndErr(err)
		return &Error{
			Code: code,
			Err:  e,
		}
	}
	return &Error{
		Code: codes[0],
		Err:  err,
	}
}

func (e *Error) Error() string {
	if e.Code == UnknownError {
		return e.Err.Error()
	}
	return fmt.Sprintf("[%s]%s", e.Code, e.Err.Error())
}

// NewStructError new struct error
func NewStructError(err error) *Error {
	return &Error{
		Code: StructError,
		Err:  err,
	}
}

func parseCodeAndErr(err error) (string, error) {
	if isRPCError(err) {
		err = formatRPCError(err)
	}

	if strings.HasPrefix(err.Error(), "[") {
		pattern := regexp.MustCompile(`^\[(.*?)\](.*)$`)
		params := pattern.FindStringSubmatch(err.Error())
		if len(params) == 3 {
			return params[1], errors.New(params[2])
		}
	}
	return UnknownError, err
}

func isRPCError(err error) bool {
	return strings.Contains(err.Error(), "rpc error: ")
}

func formatRPCError(err error) error {
	pattern := regexp.MustCompile(`^rpc error: code = .*? desc = `)
	return errors.New(pattern.ReplaceAllString(err.Error(), ""))
}
