package async

type DoneNotifier interface {
	DoneNotify() <-chan struct{}
}
