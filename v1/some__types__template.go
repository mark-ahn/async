package async

import (
	"context"
	"sync"
)

type WorkerOfSomeToOther interface {
	Push(ctx context.Context, value Some, returnCh chan<- *ReturnOfOther)
}

type WorkOfSomeToOther struct {
	Value    Some
	ReturnCh chan<- *ReturnOfOther
}

type WorkContextOfSomeToOther struct {
	Context context.Context
	*WorkOfSomeToOther
}

func (__ *WorkContextOfSomeToOther) Unpack() (context.Context, Some, chan<- *ReturnOfOther) {
	return __.Context, __.Value, __.ReturnCh
}

var (
	zero_of_WorkOfSomeToOther       WorkOfSomeToOther
	zero_of_WorkOfSomeToOther_Value Some
	pool_of_WorkOfSomeToOther       = sync.Pool{
		New: func() interface{} {
			return &WorkOfSomeToOther{}
		},
	}
	pool_of_WorkOfSomeToOtherContext = sync.Pool{
		New: func() interface{} {
			return &WorkContextOfSomeToOther{}
		},
	}
)

type pool_WorkOfSomeToOther struct{}

func (_ pool_WorkOfSomeToOther) Get() *WorkOfSomeToOther {
	return pool_of_WorkOfSomeToOther.Get().(*WorkOfSomeToOther)
}
func (_ pool_WorkOfSomeToOther) Put(d *WorkOfSomeToOther) {
	d.Value = zero_of_WorkOfSomeToOther_Value
	d.ReturnCh = nil
	pool_of_WorkOfSomeToOther.Put(d)
}

func (__ pool_WorkOfSomeToOther) GetWith(value Some, returnCh chan<- *ReturnOfOther) *WorkOfSomeToOther {
	work := __.Get()
	work.Value = value
	work.ReturnCh = returnCh
	return work
}

type pool_WorkContextOfSomeToOther struct{}

func (_ pool_WorkContextOfSomeToOther) Get() *WorkContextOfSomeToOther {
	return pool_of_WorkOfSomeToOtherContext.Get().(*WorkContextOfSomeToOther)
}
func (_ pool_WorkContextOfSomeToOther) Put(d *WorkContextOfSomeToOther) {
	d.Context = nil
	d.WorkOfSomeToOther = nil
	pool_of_WorkOfSomeToOtherContext.Put(d)
}

func (__ pool_WorkContextOfSomeToOther) GetWith(ctx context.Context, work *WorkOfSomeToOther) *WorkContextOfSomeToOther {
	work_ctx := __.Get()
	work_ctx.Context = ctx
	work_ctx.WorkOfSomeToOther = work
	return work_ctx
}

func (__ pool_WorkContextOfSomeToOther) GetWiths(ctx context.Context, value Some, returnCh chan<- *ReturnOfOther) *WorkContextOfSomeToOther {
	work := pool_of_WorkOfSomeToOther.Get().(*WorkOfSomeToOther)
	work.Value = value
	work.ReturnCh = returnCh
	return __.GetWith(ctx, work)
}

func (__ pool_WorkContextOfSomeToOther) Collect(d *WorkContextOfSomeToOther) (context.Context, Some, chan<- *ReturnOfOther) {
	ctx, value, rtn_ch := d.Unpack()
	__.Puts(d)
	return ctx, value, rtn_ch
}

func (__ pool_WorkContextOfSomeToOther) Puts(d *WorkContextOfSomeToOther) {
	work := d.WorkOfSomeToOther
	work.Value = zero_of_WorkOfSomeToOther_Value
	work.ReturnCh = nil
	pool_of_WorkOfSomeToOther.Put(work)

	__.Put(d)
}

type func_worker_SomeToOther struct{}

func (_ func_worker_SomeToOther) New(ctx context.Context, h func(context.Context, Some) (context.Context, Other, error), n int) *FuncWorkerOfSomeToOther {
	return newFuncWorkerOfSomeToOther(ctx, h, n)
}

type _SomeToOther struct {
	Work struct {
		Pool pool_WorkOfSomeToOther
	}
	WorkContext struct {
		Pool pool_WorkContextOfSomeToOther
	}
	FuncWorker func_worker_SomeToOther
	// Pool struct {
	// 	Work        pool_WorkOfSomeToOther
	// 	WorkContext pool_WorkContextOfSomeToOther
	// }
}

func (__ _SomeToOther) CallAsSync(ctx context.Context, value Some, push func(ctx context.Context, value Some, returnCh chan<- *ReturnOfOther)) (context.Context, Other, error) {
	ch := Others.ChanReturn.Pool.Get()
	defer Others.ChanReturn.Pool.Put(ch)

	push(ctx, value, ch)
	rtn := <-ch
	return rtn.Context, rtn.Value, rtn.Error
}

func (__ _SomeToOther) CallAsAsync(ctx context.Context, value Some, returnCh chan<- *ReturnOfOther, h func(ctx context.Context, arg Some) (context.Context, Other, error), defered func()) {
	go func() {
		defer defered()

		ctx, res, err := h(ctx, value)
		rtn := Others.Return.Pool.GetWith(ctx, res, err)
		returnCh <- rtn
	}()
}

var SomeToOther = _SomeToOther{}
