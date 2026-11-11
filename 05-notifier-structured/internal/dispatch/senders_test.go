package dispatch_test

import (
	"errors"
	"testing"

	"example.com/notifier/internal/dispatch"
	"example.com/notifier/internal/notifiers"
)

func TestDispatcher(t *testing.T) {
	tg := notifiers.NewTelegramNotifier("@durov")
	mail := notifiers.NewEmailNotifier("grant@3blue1brown.com")

	listing := []dispatch.Notifier{tg, mail}

	result := dispatch.Dispatch(listing, "Ping 47")
	expect := error(nil)

	if result != expect {
		t.Errorf("корректное сообщение через диспетчер падает с ошибкой")
	}

	result = dispatch.Dispatch(listing, "")
	expect = notifiers.ErrEmptyMsg

	if !errors.Is(result, expect) {
		t.Errorf("пустое сообщение через диспетчер должно являться ошибкой %v, но было получено: %v", expect, result)
	}
}
