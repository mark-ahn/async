package async

import (
	"context"
	"sync"
)

type Return[Out any] struct {
	Context context.Context
	Value   Out
	Error   error
}

func (__ *Return[Out]) Unpack() (context.Context, Out, error) {
	return __.Context, __.Value, __.Error
}

type pool_return[Out any] sync.Pool

func new_pool_return[Out any]() *pool_return[Out] {
	return (*pool_return[Out])(&sync.Pool{
		New: func() any {
			var zero Return[Out]
			return &zero
		},
	})
}

func (__ *pool_return[Out]) pool() *sync.Pool { return (*sync.Pool)(__) }

func (__ *pool_return[Out]) Get() *Return[Out] {
	return __.pool().Get().(*Return[Out])
}

func (__ *pool_return[Out]) Put(d *Return[Out]) {
	var zero Out
	d.Context = nil
	d.Value = zero
	d.Error = nil
	__.pool().Put(d)
}

func (__ pool_return[Out]) Pack(ctx context.Context, value Out, err error) *Return[Out] {
	rtn := __.Get()
	rtn.Context = ctx
	rtn.Value = value
	rtn.Error = err
	return rtn
}

func (__ pool_return[Out]) Collect(d *Return[Out]) (context.Context, Out, error) {
	ctx, value, err := d.Unpack()
	__.Put(d)
	return ctx, value, err
}

type pool_ch_return[Out any] sync.Pool

func new_pool_ch_return[Out any]() *pool_ch_return[Out] {
	return (*pool_ch_return[Out])(&sync.Pool{
		New: func() any {
			zero := make(chan *Return[Out], 1)
			return zero
		},
	})
}

func (__ *pool_ch_return[Out]) pool() *sync.Pool { return (*sync.Pool)(__) }

func (__ *pool_ch_return[Out]) Get() chan *Return[Out] {
	return __.pool().Get().(chan *Return[Out])
}
func (__ *pool_ch_return[Out]) Put(d chan *Return[Out]) {
	__.pool().Put(d)
}

type StackOfChanReturn[Out any] struct {
	chans []chan<- *Return[Out]
	sync.Mutex
}

func NewStackOfChanReturn[Out any](n int) *StackOfChanReturn[Out] {
	return &StackOfChanReturn[Out]{
		chans: make([]chan<- *Return[Out], 0, n),
	}
}

func (__ *StackOfChanReturn[Out]) Push(ch chan<- *Return[Out]) {
	__.Lock()
	__.chans = append(__.chans, ch)
	__.Unlock()
}
func (__ *StackOfChanReturn[Out]) Pop() chan<- *Return[Out] {
	var ch chan<- *Return[Out]
	__.Lock()
	defer __.Unlock()
	last := len(__.chans) - 1
	if last < 0 {
		return ch
	}
	__.chans, ch = __.chans[:last], __.chans[last]
	return ch
}

func (__ *StackOfChanReturn[Out]) Top() chan<- *Return[Out] {
	var ch chan<- *Return[Out]
	__.Lock()
	defer __.Unlock()
	last := len(__.chans) - 1
	if last < 0 {
		return ch
	}
	return __.chans[last]
}

type async_key_for_Other int

const (
	async_key_for_Other_return_ch_stack async_key_for_Other = iota
)

type ReturnPoolSet[Out any] struct {
	Value pool_return[Out]
	Chan  pool_ch_return[Out]
}

func NewReturnPoolSet[Out any]() *ReturnPoolSet[Out] {
	return &ReturnPoolSet[Out]{
		Value: *new_pool_return[Out](),
		Chan:  *new_pool_ch_return[Out](),
	}
}

// func (__ *ReturnPoolSet[Out]) GetChan() chan *Return[Out] { return __.Chan.Get() }
// func (__ *ReturnPoolSet[Out]) GetValue() *Return[Out]     { return __.Value.Get() }
// func (__ *ReturnPoolSet[Out]) PackValue(ctx context.Context, value Out, err error) *Return[Out] {
// 	return __.Value.Pack(ctx, value, err)
// }

func (_ ReturnPoolSet[Out]) stack_of(ctx Valuable) (*StackOfChanReturn[Out], bool) {
	stack, ok := ctx.Value(async_key_for_Other_return_ch_stack).(*StackOfChanReturn[Out])
	return stack, ok
}

func (_ ReturnPoolSet[Out]) WithStack(ctx context.Context, n int) context.Context {
	stack := NewStackOfChanReturn[Out](n)
	return context.WithValue(ctx, async_key_for_Other_return_ch_stack, stack)
}

func (__ ReturnPoolSet[Out]) Pop(ctx Valuable) chan<- *Return[Out] {
	stack, ok := __.stack_of(ctx)
	if !ok {
		return nil
	}
	return stack.Pop()
}

func (__ ReturnPoolSet[Out]) Top(ctx Valuable) chan<- *Return[Out] {
	stack, ok := __.stack_of(ctx)
	if !ok {
		return nil
	}
	return stack.Top()
}

func (__ ReturnPoolSet[Out]) Push(ctx Valuable, ch chan<- *Return[Out]) bool {
	stack, ok := __.stack_of(ctx)
	if !ok {
		return false
	}
	stack.Push(ch)
	return true
}

func (__ ReturnPoolSet[Out]) Notify(ctx Valuable, rtn *Return[Out]) bool {
	ch := __.Pop(ctx)
	if ch == nil {
		return false
	}
	ch <- rtn
	return true
}

func (__ ReturnPoolSet[Out]) NotifyWith(ctx context.Context, value Out, err error) bool {
	return __.Notify(ctx, __.Value.Pack(ctx, value, err))
}
