package notifiers_test

import (
	"errors"
	"testing"

	"example.com/notifier/internal/notifiers"
)

func TestEmptyMsg(t *testing.T) {
	n := notifiers.NewEmailNotifier("grant@3blue1brown.com")

	result := n.Send("")
	expect := notifiers.ErrEmptyMsg

	if !errors.Is(result, expect) {
		t.Errorf("пустое сообщение должно являться ошибкой")
	}
}

func TestMailMsg(t *testing.T) {
	n := notifiers.NewEmailNotifier("grant@3blue1brown.com")

	result := n.Send("Hello Grant")
	expect := error(nil)

	if !errors.Is(result, expect) {
		t.Errorf("корректное сообщение по почте падает с ошибкой")
	}
}
