package errors

// DomainErrorType is a type for domain error type
type DomainErrorType int

const (
	_ DomainErrorType = iota
	// DomainErrorTypeNotFound is a domain error type for not found error
	DomainErrorTypeNotFound
	// DomainErrorTypeInvalidData is a domain error type for invalid data error
	DomainErrorTypeInvalidData
	// DomainErrorTypeDuplicateEntry is a domain error type for duplicate entry
	DomainErrorTypeDuplicateEntry
)
