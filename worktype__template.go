package async

import (
	"context"
)

type WorkOfSomeThenOther struct {
	Argument Some
	ReturnCh chan<- *ReturnOfOther
}

type WorkWithContextOfSomeThenOther struct {
	context.Context
	*WorkOfSomeThenOther
}

type WorkerOfPushSomeThenOther interface {
	Push(context.Context, *WorkOfSomeThenOther)
	DoneNotify() <-chan struct{}
}
