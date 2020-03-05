package async

type DoneNotifier interface {
	DoneNotify() <-chan struct{}
}
type Valuable interface {
	Value(interface{}) interface{}
}

// type pool struct {
// 	ReturnOfOther     pool_ReturnOfOther
// 	ChanReturnOfOther pool_ChanReturnOfOther
// }

// var Pool = pool{}
