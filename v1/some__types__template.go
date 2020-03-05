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

func getWorkOfSomeToOther() *WorkOfSomeToOther {
	return pool_of_WorkOfSomeToOther.Get().(*WorkOfSomeToOther)
}
func putWorkOfSomeToOther(d *WorkOfSomeToOther) {
	d.Value = zero_of_WorkOfSomeToOther_Value
	d.ReturnCh = nil
	pool_of_WorkOfSomeToOther.Put(d)
}

func getWorkContextOfSomeToOther() *WorkContextOfSomeToOther {
	return pool_of_WorkOfSomeToOtherContext.Get().(*WorkContextOfSomeToOther)
}
func putWorkContextOfSomeToOther(d *WorkContextOfSomeToOther) {
	d.Context = nil
	d.Value = zero_of_WorkOfSomeToOther_Value
	d.ReturnCh = nil
	pool_of_WorkOfSomeToOtherContext.Put(d)
}

type _SomeToOther struct{}

func (_ _SomeToOther) GetWork() *WorkOfSomeToOther {
	return getWorkOfSomeToOther()
}
func (__ _SomeToOther) GetWorkWith(value Some, returnCh chan<- *ReturnOfOther) *WorkOfSomeToOther {
	work := __.GetWork()
	work.Value = value
	work.ReturnCh = returnCh
	return work
}

func (_ _SomeToOther) PutWork(d *WorkOfSomeToOther) {
	putWorkOfSomeToOther(d)
}

func (_ _SomeToOther) GetReturn() *ReturnOfOther {
	return getReturnOfOther()
}

func (__ _SomeToOther) GetReturnWith(ctx context.Context, value Other, err error) *ReturnOfOther {
	rtn := getReturnOfOther()
	rtn.Context = ctx
	rtn.Value = value
	rtn.Error = err
	return rtn
}

func (_ _SomeToOther) PutReturn(d *ReturnOfOther) {
	putReturnOfOther(d)
}

func (_ _SomeToOther) GetReturnCh() chan *ReturnOfOther {
	return getReturnChOfOther()
}

func (_ _SomeToOther) PutReturnCh(d chan *ReturnOfOther) {
	putReturnChOfOther(d)
}

func (_ _SomeToOther) GetWorkContext() *WorkContextOfSomeToOther {
	return getWorkContextOfSomeToOther()
}
func (__ _SomeToOther) GetWorkContextWith(ctx context.Context, work *WorkOfSomeToOther) *WorkContextOfSomeToOther {
	work_ctx := __.GetWorkContext()
	work_ctx.Context = ctx
	work_ctx.WorkOfSomeToOther = work
	return work_ctx
}

func (_ _SomeToOther) PutWorkContext(d *WorkContextOfSomeToOther) {
	putWorkContextOfSomeToOther(d)
}

func (__ _SomeToOther) CallAsSync(ctx context.Context, value Some, push func(ctx context.Context, value Some, returnCh chan<- *ReturnOfOther)) (context.Context, Other, error) {
	ch := __.GetReturnCh()
	defer __.PutReturnCh(ch)

	push(ctx, value, ch)
	rtn := <-ch
	return rtn.Context, rtn.Value, rtn.Error
}

func (__ _SomeToOther) CallAsAsync(ctx context.Context, value Some, returnCh chan<- *ReturnOfOther, h func(ctx context.Context, arg Some) (Other, error), defered func()) {
	go func() {
		defer defered()

		res, err := h(ctx, value)
		rtn := __.GetReturnWith(ctx, res, err)
		returnCh <- rtn
	}()
}

func (__ _SomeToOther) WithReturnChStack(ctx context.Context, n int) context.Context {
	return insertStackOfReturnChOfOther(ctx, n)
}

func (__ _SomeToOther) PopReturnCh(ctx Valuable) chan<- *ReturnOfOther {
	return popReturnChOfOther(ctx)
}

func (__ _SomeToOther) TopReturnCh(ctx Valuable) chan<- *ReturnOfOther {
	return topReturnChOfOther(ctx)
}

func (__ _SomeToOther) PushReturnCh(ctx Valuable, ch chan<- *ReturnOfOther) {
	pushReturnChOfOther(ctx, ch)
}

func (__ _SomeToOther) NotifyOnReturnCh(ctx Valuable, rtn *ReturnOfOther) bool {
	ch := __.TopReturnCh(ctx)
	if ch == nil {
		return false
	}
	ch <- rtn
	return true
}

var SomeToOther = _SomeToOther{}
