package dispatch

import (
	"errors"
	"fmt"

	"example.com/notifier/internal/notifiers"
)

type Notifier interface {
	Send(string) error
}

func SafeSend(n Notifier, msg string) (err error) {
	defer func() {
		if value := recover(); value != nil {
			err = fmt.Errorf("intercepted panic at send: %v", value)
		}
	}()

	return n.Send(msg)
}

func Dispatch(all []Notifier, msg string) error {
	for _, n := range all {
		if err := SafeSend(n, msg); err != nil {
			if errors.Is(err, notifiers.ErrEmptyMsg) || errors.Is(err, notifiers.ErrLongMsg) {
				return fmt.Errorf("dispatch %v: %w", n, err)
			} else {
				var target *notifiers.EmailError

				if errors.As(err, &target) {
					return fmt.Errorf("dispatch inaccessible server of %s", target.Destination())
				} else {
					return fmt.Errorf("dispatch %v unknown error: %w", n, err)
				}
			}
		}
	}

	return nil
}
