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

func getReturnOfOther() *ReturnOfOther {
	return pool_of_ReturnOfOther.Get().(*ReturnOfOther)
}
func putReturnOfOther(d *ReturnOfOther) {
	d.Context = nil
	d.Value = zero_of_ReturnOfOther_value
	d.Error = nil
	pool_of_ReturnOfOther.Put(d)
}

func getReturnChOfOther() chan *ReturnOfOther {
	return pool_of_ReturnOfOther_ch.Get().(chan *ReturnOfOther)
}
func putReturnChOfOther(d chan *ReturnOfOther) {
	pool_of_ReturnOfOther_ch.Put(d)
}

type StackOfReturnChOfOther struct {
	chans []chan<- *ReturnOfOther
	sync.Mutex
}

func NewStackOfReturnChOfOther(n int) *StackOfReturnChOfOther {
	return &StackOfReturnChOfOther{
		chans: make([]chan<- *ReturnOfOther, 0, n),
	}
}

func (__ *StackOfReturnChOfOther) Push(ch chan<- *ReturnOfOther) {
	__.Lock()
	__.chans = append(__.chans, ch)
	__.Unlock()
}
func (__ *StackOfReturnChOfOther) Pop() chan<- *ReturnOfOther {
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

func (__ *StackOfReturnChOfOther) Top() chan<- *ReturnOfOther {
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

func insertStackOfReturnChOfOther(ctx context.Context, n int) context.Context {
	stack := NewStackOfReturnChOfOther(n)
	return context.WithValue(ctx, async_key_for_Other_return_ch_stack, stack)
}

func popReturnChOfOther(ctx Valuable) chan<- *ReturnOfOther {
	stack, ok := ctx.Value(async_key_for_Other_return_ch_stack).(*StackOfReturnChOfOther)
	if !ok {
		return nil
	}
	return stack.Pop()
}

func topReturnChOfOther(ctx Valuable) chan<- *ReturnOfOther {
	stack, ok := ctx.Value(async_key_for_Other_return_ch_stack).(*StackOfReturnChOfOther)
	if !ok {
		return nil
	}
	return stack.Top()
}

func pushReturnChOfOther(ctx Valuable, ch chan<- *ReturnOfOther) bool {
	stack, ok := ctx.Value(async_key_for_Other_return_ch_stack).(*StackOfReturnChOfOther)
	if !ok {
		return false
	}
	stack.Push(ch)
	return true
}
