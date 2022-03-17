package async

import (
	"context"
	"sync"
)

type Return[Other any] struct {
	Context context.Context
	Value   Other
	Error   error
}

func (__ *Return[Other]) Unpack() (context.Context, Other, error) {
	return __.Context, __.Value, __.Error
}

var (
// zero_of_ReturnOfOther Return[Other]
// pool_of_ReturnOfOther = sync.Pool{
// 	New: func() interface{} {
// 		return &Return[Other]{}
// 	},
// }
// pool_of_ReturnOfOther_ch = sync.Pool{
// 	New: func() interface{} {
// 		return make(chan *Return[Other], 1)
// 	},
// }
// pool_map = map[reflect.Type]interface{}{
// 	reflect.TypeOf(interface{}{}): pool_Return[interface{}]{},
// }
)

type pool_Return[Other any] struct {
	pool sync.Pool
}

func (__ *pool_Return[Other]) Get() *Return[Other] {
	return __.pool.Get().(*Return[Other])
}
func (__ *pool_Return[Other]) Put(d *Return[Other]) {
	var zero Other
	d.Context = nil
	d.Value = zero
	d.Error = nil
	__.pool.Put(d)
}

func (__ pool_Return[Other]) GetWith(ctx context.Context, value Other, err error) *Return[Other] {
	rtn := __.Get()
	rtn.Context = ctx
	rtn.Value = value
	rtn.Error = err
	return rtn
}

func (__ pool_Return[Other]) Collect(d *Return[Other]) (context.Context, Other, error) {
	ctx, value, err := d.Unpack()
	__.Put(d)
	return ctx, value, err
}

type pool_ChanReturn[Other any] struct {
	pool sync.Pool
}

func (__ *pool_ChanReturn[Other]) Get() chan *Return[Other] {
	return __.pool.Get().(chan *Return[Other])
}
func (__ *pool_ChanReturn[Other]) Put(d chan *Return[Other]) {
	__.pool.Put(d)
}

type StackOfChanReturn[Other any] struct {
	chans []chan<- *Return[Other]
	sync.Mutex
}

func NewStackOfChanReturn[Other any](n int) *StackOfChanReturn[Other] {
	return &StackOfChanReturn[Other]{
		chans: make([]chan<- *Return[Other], 0, n),
	}
}

func (__ *StackOfChanReturn[Other]) Push(ch chan<- *Return[Other]) {
	__.Lock()
	__.chans = append(__.chans, ch)
	__.Unlock()
}
func (__ *StackOfChanReturn[Other]) Pop() chan<- *Return[Other] {
	var ch chan<- *Return[Other]
	__.Lock()
	defer __.Unlock()
	last := len(__.chans) - 1
	if last < 0 {
		return ch
	}
	__.chans, ch = __.chans[:last], __.chans[last]
	return ch
}

func (__ *StackOfChanReturn[Other]) Top() chan<- *Return[Other] {
	var ch chan<- *Return[Other]
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

type ReturnPoolSet[Other any] struct {
	Return struct {
		Pool pool_Return[Other]
	}
	ChanReturn struct {
		Pool pool_ChanReturn[Other]
	}
}

func NewReturnPoolSet[Other any]() *ReturnPoolSet[Other] {
	return &ReturnPoolSet[Other]{
		Return: struct {
			Pool pool_Return[Other]
		}{
			Pool: pool_Return[Other]{
				pool: sync.Pool{
					New: func() any {
						var zero Return[Other]
						return &zero
					},
				},
			},
		},
		ChanReturn: struct{ Pool pool_ChanReturn[Other] }{
			Pool: pool_ChanReturn[Other]{
				pool: sync.Pool{
					New: func() any {
						zero := make(chan *Return[Other], 1)
						return zero
					},
				},
			},
		},
	}
}

// type _OtherSet[Other any] struct {
// 	Return struct {
// 		Pool pool_Return[Other]
// 	}
// 	ChanReturn struct {
// 		Pool    pool_ChanReturn[Other]
// 		Context chanReturn[Other]
// 	}
// }

func (_ ReturnPoolSet[Other]) stack_of(ctx Valuable) (*StackOfChanReturn[Other], bool) {
	stack, ok := ctx.Value(async_key_for_Other_return_ch_stack).(*StackOfChanReturn[Other])
	return stack, ok
}

func (_ ReturnPoolSet[Other]) WithStack(ctx context.Context, n int) context.Context {
	stack := NewStackOfChanReturn[Other](n)
	return context.WithValue(ctx, async_key_for_Other_return_ch_stack, stack)
}

func (__ ReturnPoolSet[Other]) Pop(ctx Valuable) chan<- *Return[Other] {
	stack, ok := __.stack_of(ctx)
	if !ok {
		return nil
	}
	return stack.Pop()
}

func (__ ReturnPoolSet[Other]) Top(ctx Valuable) chan<- *Return[Other] {
	stack, ok := __.stack_of(ctx)
	if !ok {
		return nil
	}
	return stack.Top()
}

func (__ ReturnPoolSet[Other]) Push(ctx Valuable, ch chan<- *Return[Other]) bool {
	stack, ok := __.stack_of(ctx)
	if !ok {
		return false
	}
	stack.Push(ch)
	return true
}

func (__ ReturnPoolSet[Other]) Notify(ctx Valuable, rtn *Return[Other]) bool {
	ch := __.Pop(ctx)
	if ch == nil {
		return false
	}
	ch <- rtn
	return true
}

func (__ ReturnPoolSet[Other]) NotifyWith(ctx context.Context, value Other, err error) bool {
	return __.Notify(ctx, __.Return.Pool.GetWith(ctx, value, err))
}

// type _OtherSet[Other any] struct {
// 	Return struct {
// 		Pool pool_Return[Other]
// 	}
// 	ChanReturn struct {
// 		Pool    pool_ChanReturn[Other]
// 		Context chanReturn[Other]
// 	}
// }

// var Others = _Other{}
