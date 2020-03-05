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
