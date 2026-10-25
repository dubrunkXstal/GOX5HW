package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrEmptyMsg = errors.New("сообщение не может быть пустым")

type Notifier interface {
	Send(string) error
}

type BaseNotifier struct {
	destination string
}

func (b BaseNotifier) Log(msg string) {
	log.Println(b.destination, msg)
}

type EmailNotifier struct{ BaseNotifier }

func (n EmailNotifier) Send(msg string) error {
	if msg == "" {
		return ErrEmptyMsg
	}

	n.Log(msg)
	_, err := fmt.Printf("Email:\n\tmail to: %s\n\tmessage: %s\n\n", n.destination, msg)

	return err
}

type TelegramNotifier struct{ BaseNotifier }

func (n TelegramNotifier) Send(msg string) error {
	if msg == "" {
		return ErrEmptyMsg
	}

	n.Log(msg)
	_, err := fmt.Printf("Telegram:\n\tuser   : %s\n\tmessage: %s\n\n", n.destination, msg)

	return err
}

func Dispatch(all []Notifier, msg string) error {
	for _, n := range all {
		if err := n.Send(msg); err != nil {
			return err
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

	_ = tg.Send("Верни стену")
	_ = mail.Send("Love your videos")

	listing := []Notifier{tg, mail}

	_ = Dispatch(listing, "Ping 123")
}
