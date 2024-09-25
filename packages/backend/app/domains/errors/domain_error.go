package errors

import "fmt"

// DomainError is an interface for domain error
type DomainError interface {
	error
	// DomainErrorType returns the type of domain error
	DomainErrorType() DomainErrorType
}

type domainError struct {
	orgError  error
	errorType DomainErrorType
	message   string
}

// NewDomainError returns a new DomainError
func NewDomainError(orgError error, errorType DomainErrorType, msg string) DomainError {
	return domainError{
		errorType: errorType,
		orgError:  orgError,
		message:   msg,
	}
}

func (d domainError) DomainErrorType() DomainErrorType {
	return d.errorType
}

func (d domainError) Error() string {
	return fmt.Sprintf("%s: %s", d.message, d.orgError.Error())
}
