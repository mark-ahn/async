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

type WorkerOfCallSomeThenOther interface {
	Call(context.Context, Some) (Other, error)
}

type WorkerOfPushSomeThenOther interface {
	Push(context.Context, *WorkOfSomeThenOther)
}
