package errbldr

import "github.com/pkg/errors"

// Msg proxies to errors.New and returns ErrorBuilder from the result
func Msg(message string) ErrorBuilder {
	return &errorBuilder{err: errors.New(message)}
}

// Msgf proxies to errors.Errorf and returns ErrorBuilder from the result
func Msgf(format string, args ...interface{}) ErrorBuilder {
	return &errorBuilder{err: errors.Errorf(format, args...)}
}

// Wrap proxies to errors.Wrap and returns ErrorBuilder from the result
func Wrap(err error, message string) ErrorBuilder {
	return &errorBuilder{err: errors.Wrap(err, message)}
}

// Wrapf proxies to errors.Wrapf and returns ErrorBuilder from the result
func Wrapf(err error, format string, args ...interface{}) ErrorBuilder {
	return &errorBuilder{err: errors.Wrapf(err, format, args...)}
}

// ErrorBuilder is a core interface of the library
type ErrorBuilder interface {
	Err() error
	Msg(message string) ErrorBuilder
	Msgf(format string, args ...interface{}) ErrorBuilder
}

type errorBuilder struct {
	err error
}

// Err returns an underlying error. Should be called after all build methods.
func (b *errorBuilder) Err() error {
	return b.err
}

// Msg wraps the underlying err with a message
func (b *errorBuilder) Msg(message string) ErrorBuilder {
	b.err = errors.Wrap(b.err, message)
	return b
}

// Msgf wraps the underlying err with a formatted message
func (b *errorBuilder) Msgf(format string, args ...interface{}) ErrorBuilder {
	b.err = errors.Wrapf(b.err, format, args...)
	return b
}
