package notifiers

import "fmt"

type TelegramNotifier struct{ baseNotifier }

func NewTelegramNotifier(destination string) *TelegramNotifier {
	return &TelegramNotifier{baseNotifier: *newBaseNotifier(destination)}
}

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

	n.log(msg)
	_, err := fmt.Printf("Telegram:\n\tuser   : %s\n\tmessage: %s\n\n", n.destination, msg)

	return err
}
