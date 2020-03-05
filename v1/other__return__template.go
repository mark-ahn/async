package async

import (
	"context"
	"sync"
)

type ReturnOfOther struct {
	Context context.Context
	Value   Other
	Error   error
}

func (__ *ReturnOfOther) Unpack() (context.Context, Some, error) {
	return __.Context, __.Value, __.Error
}

var (
	zero_of_ReturnOfOther       ReturnOfOther
	zero_of_ReturnOfOther_value Other
	pool_of_ReturnOfOther       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfOther{}
		},
	}
	pool_of_ReturnOfOther_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfOther, 1)
		},
	}
)

type pool_ReturnOfOther struct{}

func (_ pool_ReturnOfOther) Get() *ReturnOfOther {
	return pool_of_ReturnOfOther.Get().(*ReturnOfOther)
}
func (_ pool_ReturnOfOther) Put(d *ReturnOfOther) {
	d.Context = nil
	d.Value = zero_of_ReturnOfOther_value
	d.Error = nil
	pool_of_ReturnOfOther.Put(d)
}

func (__ pool_ReturnOfOther) GetWith(ctx context.Context, value Other, err error) *ReturnOfOther {
	rtn := __.Get()
	rtn.Context = ctx
	rtn.Value = value
	rtn.Error = err
	return rtn
}

type pool_ChanReturnOfOther struct{}

func (_ pool_ChanReturnOfOther) Get() chan *ReturnOfOther {
	return pool_of_ReturnOfOther_ch.Get().(chan *ReturnOfOther)
}
func (_ pool_ChanReturnOfOther) Put(d chan *ReturnOfOther) {
	pool_of_ReturnOfOther_ch.Put(d)
}

type StackOfChanReturnOfOther struct {
	chans []chan<- *ReturnOfOther
	sync.Mutex
}

func NewStackOfChanReturnOfOther(n int) *StackOfChanReturnOfOther {
	return &StackOfChanReturnOfOther{
		chans: make([]chan<- *ReturnOfOther, 0, n),
	}
}

func (__ *StackOfChanReturnOfOther) Push(ch chan<- *ReturnOfOther) {
	__.Lock()
	__.chans = append(__.chans, ch)
	__.Unlock()
}
func (__ *StackOfChanReturnOfOther) Pop() chan<- *ReturnOfOther {
	var ch chan<- *ReturnOfOther
	__.Lock()
	defer __.Unlock()
	last := len(__.chans) - 1
	if last < 0 {
		return ch
	}
	__.chans, ch = __.chans[:last], __.chans[last]
	return ch
}

func (__ *StackOfChanReturnOfOther) Top() chan<- *ReturnOfOther {
	var ch chan<- *ReturnOfOther
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

type chanReturnOfOther struct{}

func (_ chanReturnOfOther) WithStack(ctx context.Context, n int) context.Context {
	stack := NewStackOfChanReturnOfOther(n)
	return context.WithValue(ctx, async_key_for_Other_return_ch_stack, stack)
}

func (_ chanReturnOfOther) Pop(ctx Valuable) chan<- *ReturnOfOther {
	stack, ok := ctx.Value(async_key_for_Other_return_ch_stack).(*StackOfChanReturnOfOther)
	if !ok {
		return nil
	}
	return stack.Pop()
}

func (_ chanReturnOfOther) Top(ctx Valuable) chan<- *ReturnOfOther {
	stack, ok := ctx.Value(async_key_for_Other_return_ch_stack).(*StackOfChanReturnOfOther)
	if !ok {
		return nil
	}
	return stack.Top()
}

func (_ chanReturnOfOther) Push(ctx Valuable, ch chan<- *ReturnOfOther) bool {
	stack, ok := ctx.Value(async_key_for_Other_return_ch_stack).(*StackOfChanReturnOfOther)
	if !ok {
		return false
	}
	stack.Push(ch)
	return true
}

func (__ chanReturnOfOther) Notify(ctx Valuable, rtn *ReturnOfOther) bool {
	ch := __.Top(ctx)
	if ch == nil {
		return false
	}
	ch <- rtn
	return true
}

type _Other struct {
	Return struct {
		Pool pool_ReturnOfOther
	}
	ChanReturn struct {
		Pool    pool_ChanReturnOfOther
		Context chanReturnOfOther
	}
	// Pool struct {
	// 	Return     pool_ReturnOfOther
	// 	ChanReturn pool_ChanReturnOfOther
	// }
	// Context struct {
	// 	ChanReturn chanReturnOfOther
	// }
}

var Others = _Other{}
