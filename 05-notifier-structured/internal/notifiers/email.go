package notifiers

import (
	"fmt"

	"example.com/notifier/internal/model"
)

type EmailNotifier struct{ model.BaseNotifier }

func NewEmailNotifier(destination string) *EmailNotifier {
	return &EmailNotifier{BaseNotifier: *model.NewBaseNotifier(destination)}
}

func (n EmailNotifier) Send(msg string) error {
	if msg == "" {
		return fmt.Errorf("email send: %w", model.ErrEmptyMsg)
	}
	if len(msg) > 30 {
		return fmt.Errorf("email send: %w", model.ErrLongMsg)
	}
	if msg == "FORCE_EMAIL_ERROR" {
		return model.NewEmailError(n.Destination(), model.ErrBadServer)
	}

	n.Log(msg)
	_, err := fmt.Printf("Email:\n\tmail to: %s\n\tmessage: %s\n\n", n.Destination(), msg)

	return err
}
