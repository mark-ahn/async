package async

type WorkOfSomeThenOther struct {
	Argument Some
	ReturnCh chan<- *ReturnOfOther
}

type WorkerOfPushSomeThenOther interface {
	Push(*WorkOfSomeThenOther)
	DoneNotify() <-chan struct{}
}
