package async

type DoneNotifier interface {
	DoneNotify() <-chan struct{}
}
type Valuable interface {
	Value(interface{}) interface{}
}
