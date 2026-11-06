package notifiers

import (
	"example.com/notifier/internal/model"
)

type BrokenNotifier struct{ model.BaseNotifier }

func NewBrokenNotifier(destination string) *BrokenNotifier {
	return &BrokenNotifier{BaseNotifier: *model.NewBaseNotifier(destination)}
}

func (BrokenNotifier) Send(msg string) error {
	panic("Panic at notifier")
}
