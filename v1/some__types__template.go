package async

import (
	"context"
	"sync"
)

type WorkerOfSomeToOther interface {
	Push(ctx context.Context, work *WorkOfSomeToOther, returnCh chan<- *ReturnOfOther)
}

type WorkOfSomeToOther struct {
	Value    Some
	ReturnCh chan<- *ReturnOfOther
}

type WorkOfSomeToOtherWithContext struct {
	Context context.Context
	WorkOfSomeToOther
}

var (
	zero_of_WorkOfSomeToOther       WorkOfSomeToOther
	zero_of_WorkOfSomeToOther_Value Some
	pool_of_WorkOfSomeToOther       = sync.Pool{
		New: func() interface{} {
			return &WorkOfSomeToOther{}
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

func (__ _SomeToOther) CallSync(ctx context.Context, value Some, push func(ctx context.Context, work *WorkOfSomeToOther, returnCh chan<- *ReturnOfOther)) (context.Context, Other, error) {
	ch := __.GetReturnCh()
	defer __.PutReturnCh(ch)

	work := __.GetWorkWith(value, ch)
	push(ctx, work, ch)
	rtn := <-ch
	defer __.PutReturn(rtn)
	return rtn.Context, rtn.Value, rtn.Error
}

func (__ _SomeToOther) CallAsync(ctx context.Context, work *WorkOfSomeToOther, h func(ctx context.Context, arg Some) (Other, error), defered func()) {
	go func() {
		defer defered()

		res, err := h(ctx, work.Value)
		rtn := __.GetReturnWith(ctx, res, err)
		work.ReturnCh <- rtn
	}()
}
