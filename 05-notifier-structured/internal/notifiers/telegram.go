package notifiers

import (
	"fmt"

	"example.com/notifier/internal/model"
)

type TelegramNotifier struct{ model.BaseNotifier }

func NewTelegramNotifier(destination string) *TelegramNotifier {
	return &TelegramNotifier{BaseNotifier: *model.NewBaseNotifier(destination)}
}

func (n TelegramNotifier) Send(msg string) error {
	if msg == "" {
		return fmt.Errorf("telegram send: %w", model.ErrEmptyMsg)
	}
	if len(msg) > 20 {
		return fmt.Errorf("telegram send: %w", model.ErrLongMsg)
	}
	if msg == "FORCE_EMAIL_ERROR" {
		return nil
	}

	n.Log(msg)
	_, err := fmt.Printf("Telegram:\n\tuser   : %s\n\tmessage: %s\n\n", n.Destination(), msg)

	return err
}
