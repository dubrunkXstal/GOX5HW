package main

import (
	"fmt"

	"example.com/notifier/internal/model"
	"example.com/notifier/internal/notifiers"
	"example.com/notifier/internal/senders"
)

func main() {
	tg := notifiers.NewTelegramNotifier("@durov")
	mail := notifiers.NewEmailNotifier("grant@3blue1brown.com")
	broken := notifiers.NewBrokenNotifier("£$@%4**.")

	listing := []model.Notifier{tg, mail}

	err := senders.SafeSend(broken, "Адхщба")
	fmt.Printf("1. Паника в нотификаторе:                 %s\n", err)

	err = tg.Send("0123456789012345678911111")
	fmt.Printf("2. Длинное для тг сообщение напрямую:     %s\n", err)

	err = senders.Dispatch(listing, "0123456789012345678911111")
	fmt.Printf("3. Длинное для тг сообщение в диспетчере: %s\n", err)

	err = mail.Send("FORCE_EMAIL_ERROR")
	fmt.Printf("4. Ошибка почтового сервера напрямую:     %s\n", err)

	err = senders.Dispatch(listing, "FORCE_EMAIL_ERROR")
	fmt.Printf("5. Ошибка почтового сервера в диспетчере: %s\n", err)

	err = senders.SafeSend(tg, "FORCE_LOG_PANIC")
	fmt.Printf("6. Паника в логгере:                      %s\n", err)
}
