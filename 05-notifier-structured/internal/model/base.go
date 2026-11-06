package model

import "log"

type Notifier interface {
	Send(string) error
}

type BaseNotifier struct {
	destination string
}

func NewBaseNotifier(destination string) *BaseNotifier {
	return &BaseNotifier{destination: destination}
}

func (b BaseNotifier) Destination() string {
	return b.destination
}

func (b BaseNotifier) Log(msg string) {
	if msg == "FORCE_LOG_PANIC" {
		panic("Panic at logger")
	}

	log.Println(b.destination, msg)
}
