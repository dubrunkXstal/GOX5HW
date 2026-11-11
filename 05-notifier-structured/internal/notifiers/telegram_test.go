package notifiers_test

import (
	"errors"
	"testing"

	"example.com/notifier/internal/notifiers"
)

func TestTgMsg(t *testing.T) {
	n := notifiers.NewTelegramNotifier("@durov")

	result := n.Send("Hello Pasha")
	expect := error(nil)

	if !errors.Is(result, expect) {
		t.Errorf("корректное сообщение по телеграму падает с ошибкой")
	}
}
