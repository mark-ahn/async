package async

import (
	"context"
	"sync"
)

type ReturnOfSome struct {
	Context context.Context
	Value   Some
	Error   error
}

var pool_of_ReturnOfSome = sync.Pool{
	New: func() interface{} {
		return &ReturnOfSome{}
	},
}

func GetReturnOfSome() *ReturnOfSome {
	return pool_of_ReturnOfSome.Get().(*ReturnOfSome)
}
func PutReturnOfSome(d *ReturnOfSome) {
	pool_of_ReturnOfSome.Put(d)
}
