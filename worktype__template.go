package async

import (
	"context"
)

type _Prefix_WorkOfSomeThenOther struct {
	Argument Some
	ReturnCh chan<- *_Prefix_ReturnOfOther
}

type _Prefix_WorkerOfPushSomeThenOther interface {
	Push(context.Context, *_Prefix_WorkOfSomeThenOther) error
	DoneNotify() <-chan struct{}
}
