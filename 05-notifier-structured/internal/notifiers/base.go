package notifiers

import "log"

type baseNotifier struct {
	destination string
}

func newBaseNotifier(destination string) *baseNotifier {
	return &baseNotifier{destination: destination}
}

func (b baseNotifier) log(msg string) {
	if msg == "FORCE_LOG_PANIC" {
		panic("Panic at logger")
	}

	log.Println(b.destination, msg)
}
