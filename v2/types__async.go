package async

import "context"

type func_worker[In any, Out any] struct{}

func (_ func_worker[In, Out]) New(ctx context.Context, h func(context.Context, In) (context.Context, Out, error), n int) *FuncWorker[In, Out] {
	return NewFuncWorker(ctx, h, n)
}

type AsyncPoolSet[In any, Out any] struct {
	Param  ParamPoolSet[In, Out]
	Return ReturnPoolSet[Out]
}

func NewAsyncPoolSet[In any, Out any]() *AsyncPoolSet[In, Out] {
	return &AsyncPoolSet[In, Out]{
		Param:  *NewParamPoolSet[In, Out](),
		Return: *NewReturnPoolSet[Out](),
	}
}

func (__ AsyncPoolSet[In, Out]) CallAsSync(ctx context.Context, value In, push func(ctx context.Context, value In, returnCh chan<- *Return[Out])) (context.Context, Out, error) {
	ch := __.Return.Chan.Get()
	defer __.Return.Chan.Put(ch)

	push(ctx, value, ch)
	rtn := <-ch
	return rtn.Context, rtn.Value, rtn.Error
}

func (__ AsyncPoolSet[In, Out]) CallAsAsync(ctx context.Context, value In, returnCh chan<- *Return[Out], h func(ctx context.Context, arg In) (context.Context, Out, error), defered func()) {
	go func() {
		defer defered()

		ctx, res, err := h(ctx, value)
		rtn := __.Return.Value.Pack(ctx, res, err)
		returnCh <- rtn
	}()
}
