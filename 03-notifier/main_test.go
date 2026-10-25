package main

import (
	"testing"
)

func TestEmptyMsg(t *testing.T) {
	n := EmailNotifier{
		BaseNotifier: BaseNotifier{"grant@3blue1brown.com"},
	}

	result := n.Send("")
	expect := ErrEmptyMsg

	if result != expect {
		t.Errorf("пустое сообщение должно являться ошибкой")
	}
}

func TestMailMsg(t *testing.T) {
	n := EmailNotifier{
		BaseNotifier: BaseNotifier{"grant@3blue1brown.com"},
	}

	result := n.Send("Hello Grant")
	expect := error(nil)

	if result != expect {
		t.Errorf("корректное сообщение по почте падает с ошибкой")
	}
}

func TestTgMsg(t *testing.T) {
	n := TelegramNotifier{
		BaseNotifier: BaseNotifier{"@durov"},
	}

	result := n.Send("Hello Pasha")
	expect := error(nil)

	if result != expect {
		t.Errorf("корректное сообщение по телеграму падает с ошибкой")
	}
}

func TestDispatcher(t *testing.T) {
	tg := TelegramNotifier{
		BaseNotifier: BaseNotifier{destination: "@durov"},
	}

	mail := EmailNotifier{
		BaseNotifier: BaseNotifier{destination: "grant@3blue1brown.com"},
	}

	listing := []Notifier{tg, mail}

	result := Dispatch(listing, "Ping 47")
	expect := error(nil)

	if result != expect {
		t.Errorf("корректное сообщение через диспетчер падает с ошибкой")
	}

	result = Dispatch(listing, "")
	expect = ErrEmptyMsg

	if result != expect {
		t.Errorf("пустое сообщение через диспетчер должно являться ошибкой")
	}
}
