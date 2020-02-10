package async

import (
	"context"
)

type WorkOfSomeThenOther struct {
	Argument Some
	ReturnCh chan<- *ReturnOfOther
}

type WorkerOfPushSomeThenOther interface {
	Push(context.Context, *WorkOfSomeThenOther) error
	DoneNotify() <-chan struct{}
}
