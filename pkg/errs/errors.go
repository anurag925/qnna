package errs

import "errors"

var (
	// error messages
	// application global errors
	ErrInvalidJWT       error = errors.New("invalid JWT given")
	ErrInvalidToken     error = errors.New("invalid token given")
	ErrNotUUID          error = errors.New("invalid UUID given")
	ErrMalformedData    error = errors.New("malformed data given as input")
	ErrNotFound         error = errors.New("resource not found")
	ErrBadRequest       error = errors.New("bad request")
	ErrInvalidInput     error = errors.New("invalid input")
	ErrDuplicateEntry   error = errors.New("duplicate entry")
	ErrInvalidOperation error = errors.New("invalid operation")
	ErrStructValidation error = errors.New("invalid data in struct")
)
