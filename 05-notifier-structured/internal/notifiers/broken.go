package notifiers

type BrokenNotifier struct{ baseNotifier }

func NewBrokenNotifier(destination string) *BrokenNotifier {
	return &BrokenNotifier{baseNotifier: *newBaseNotifier(destination)}
}

func (BrokenNotifier) Send(msg string) error {
	panic("Panic at notifier")
}
