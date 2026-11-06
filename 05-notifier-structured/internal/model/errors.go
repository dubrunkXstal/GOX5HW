package model

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyMsg  = errors.New("сообщение не может быть пустым")
	ErrLongMsg   = errors.New("слишком длинное сообщение")
	ErrBadServer = errors.New("недоступный сервер")
	ErrNotFound  = errors.New("сообщение не найдено")
)

type EmailError struct {
	destination string
	err         error
}

func NewEmailError(destination string, err error) *EmailError {
	return &EmailError{
		destination: destination,
		err:         err,
	}
}

func (e EmailError) Destination() string {
	return e.destination
}

func (e EmailError) Error() string {
	return fmt.Sprintf("server of %v doesn't responding because of %v", e.destination, e.err)
}

func (e EmailError) Unwrap() error {
	return e.err
}

func (e EmailError) Is(target error) bool {
	t, ok := target.(*EmailError)
	if !ok {
		return false
	}
	return (e.destination == t.destination)
}
