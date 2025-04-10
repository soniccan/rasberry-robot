package myerr

import "errors"

// ErrorCause はエラーの原因を返す
func ErrorCause(err error) error {
	if unwrapped := errors.Unwrap(err); unwrapped != nil {
		return ErrorCause(unwrapped)
	}
	return err
}

// As はエラーを型アサーションする
func As[T any](err error) (T, bool) {
	var target T
	if errors.As(err, &target) {
		return target, true
	}
	return target, false
}
