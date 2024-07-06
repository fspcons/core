package errs

import (
	"errors"
	"strings"
)

var (
	ErrInvalidID = NewError("invalid id")
	ErrNoChanges = NewError("no changes")
	ErrNotFound  = NewError("not found")
)

type Error struct {
	errors []error
}

func (e Error) Errors() []error {
	return e.errors
}

// Error implements error interface
func (e *Error) Error() string {
	var sb strings.Builder
	if len(e.errors) > 1 {
		_, _ = sb.WriteString("many errors: ")
	}
	for i, err := range e.errors {
		if i > 0 {
			_, _ = sb.WriteString(", ")
		}
		_, _ = sb.WriteString(err.Error())
	}
	return sb.String()
}

func ToError(errs ...error) error {
	if len(errs) == 0 {
		return nil
	}
	var e *Error
	for _, err := range errs {
		if err == nil {
			continue
		}
		var e2 *Error
		if errors.As(err, &e2) {
			if e == nil {
				e = e2
			} else {
				e.errors = append(e.errors, e2.errors...)
			}
		} else {
			if e == nil {
				e = &Error{}
			}
			e.errors = append(e.errors, err)
		}
	}
	if e == nil {
		return nil
	}
	return e
}

func AppendNewError(err error, ms ...string) error {
	errs := []error{err}
	for _, msg := range ms {
		errs = append(errs, errors.New(msg))
	}
	return ToError(errs...)
}

// NewError builds a new *Error
func NewError(ms ...string) error {
	return AppendNewError(nil, ms...)
}
