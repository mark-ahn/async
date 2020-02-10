package async

type _Prefix_WorkOfSomeThenOther struct {
	Argument Some
	ReturnCh chan<- *_Prefix_ReturnOfOther
}

type WorkerOfPushSomeThenOther interface {
	Push(*_Prefix_WorkOfSomeThenOther)
	DoneNotify() <-chan struct{}
}
