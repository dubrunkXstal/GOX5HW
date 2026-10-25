package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrEmptyMsg  = errors.New("сообщение не может быть пустым")
	ErrLongMsg   = errors.New("слишком длинное сообщение")
	ErrBadServer = errors.New("недоступный сервер")
)

type Notifier interface {
	Send(string) error
}

type BaseNotifier struct {
	destination string
}

func (b BaseNotifier) Log(msg string) {
	if msg == "FORCE_LOG_PANIC" {
		panic("Panic at logger")
	}

	log.Println(b.destination, msg)
}

type EmailError struct {
	destination string
	err         error
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

type EmailNotifier struct{ BaseNotifier }

func (n EmailNotifier) Send(msg string) error {
	if msg == "" {
		return fmt.Errorf("email send: %w", ErrEmptyMsg)
	}
	if len(msg) > 30 {
		return fmt.Errorf("email send: %w", ErrLongMsg)
	}
	if msg == "FORCE_EMAIL_ERROR" {
		return &EmailError{
			destination: n.destination,
			err:         ErrBadServer,
		}
	}

	n.Log(msg)
	_, err := fmt.Printf("Email:\n\tmail to: %s\n\tmessage: %s\n\n", n.destination, msg)

	return err
}

type TelegramNotifier struct{ BaseNotifier }

func (n TelegramNotifier) Send(msg string) error {
	if msg == "" {
		return fmt.Errorf("telegram send: %w", ErrEmptyMsg)
	}
	if len(msg) > 20 {
		return fmt.Errorf("telegram send: %w", ErrLongMsg)
	}
	if msg == "FORCE_EMAIL_ERROR" {
		return nil
	}

	n.Log(msg)
	_, err := fmt.Printf("Telegram:\n\tuser   : %s\n\tmessage: %s\n\n", n.destination, msg)

	return err
}

type BrokenNotifier struct{ BaseNotifier }

func (BrokenNotifier) Send(msg string) error {
	panic("Panic at notifier")
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
			if errors.Is(err, ErrEmptyMsg) || errors.Is(err, ErrLongMsg) {
				return fmt.Errorf("dispatch %v: %w", n, err)
			} else {
				var target *EmailError

				if errors.As(err, &target) {
					return fmt.Errorf("dispatch inaccessible server of %s", target.destination)
				} else {
					return fmt.Errorf("dispatch %v unknown error: %w", n, err)
				}
			}
		}
	}

	return nil
}

func main() {
	tg := TelegramNotifier{
		BaseNotifier: BaseNotifier{destination: "@durov"},
	}

	mail := EmailNotifier{
		BaseNotifier: BaseNotifier{destination: "grant@3blue1brown.com"},
	}

	broken := BrokenNotifier{
		BaseNotifier: BaseNotifier{destination: "£$@%4**."},
	}

	listing := []Notifier{tg, mail}

	err := SafeSend(broken, "Адхщба")
	fmt.Printf("1. Паника в нотификаторе:                 %s\n", err)

	err = tg.Send("0123456789012345678911111")
	fmt.Printf("2. Длинное для тг сообщение напрямую:     %s\n", err)

	err = Dispatch(listing, "0123456789012345678911111")
	fmt.Printf("3. Длинное для тг сообщение в диспетчере: %s\n", err)

	err = mail.Send("FORCE_EMAIL_ERROR")
	fmt.Printf("4. Ошибка почтового сервера напрямую:     %s\n", err)

	err = Dispatch(listing, "FORCE_EMAIL_ERROR")
	fmt.Printf("5. Ошибка почтового сервера в диспетчере: %s\n", err)

	err = SafeSend(tg, "FORCE_LOG_PANIC")
	fmt.Printf("6. Паника в логгере:                      %s\n", err)
}
