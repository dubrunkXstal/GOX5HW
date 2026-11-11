package notifiers

import "fmt"

type EmailNotifier struct{ baseNotifier }

func NewEmailNotifier(destination string) *EmailNotifier {
	return &EmailNotifier{baseNotifier: *newBaseNotifier(destination)}
}

func (n EmailNotifier) Send(msg string) error {
	if msg == "" {
		return fmt.Errorf("email send: %w", ErrEmptyMsg)
	}
	if len(msg) > 30 {
		return fmt.Errorf("email send: %w", ErrLongMsg)
	}
	if msg == "FORCE_EMAIL_ERROR" {
		return NewEmailError(n.destination, ErrBadServer)
	}

	n.log(msg)
	_, err := fmt.Printf("Email:\n\tmail to: %s\n\tmessage: %s\n\n", n.destination, msg)

	return err
}
