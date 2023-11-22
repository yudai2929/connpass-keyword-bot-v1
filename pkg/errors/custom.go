package errors

import "github.com/cockroachdb/errors"

type customError interface {
	Error() string
}

func New(msg string) customError {
	return errors.New(msg)
}

func Wrap(err error, msg string) customError {
	return errors.Wrap(err, msg)
}
