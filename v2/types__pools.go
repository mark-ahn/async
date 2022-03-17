package async

import (
	"context"
	"reflect"
	"sync"
)

type type_param[In any, Out any] struct{}

var pool_lock sync.Mutex
var pools = map[reflect.Type]interface{}{}

func PoolOf[In any, Out any]() *AsyncPoolSet[In, Out] {
	var pool *AsyncPoolSet[In, Out]
	{
		type_p := reflect.TypeOf(type_param[In, Out]{})
		pool_lock.Lock()
		pool_i, ok := pools[type_p]
		if !ok {
			pool = NewAsyncPoolSet[In, Out]()
			pools[type_p] = pool
		} else {
			pool = pool_i.(*AsyncPoolSet[In, Out])
		}
		pool_lock.Unlock()
	}
	return pool
}

func CallAsAsync[In any, Out any](ctx context.Context, value In, returnCh chan<- *Return[Out], h func(ctx context.Context, arg In) (context.Context, Out, error), defered func()) {
	PoolOf[In, Out]().CallAsAsync(ctx, value, returnCh, h, defered)
}

func CallAsSync[In any, Out any](ctx context.Context, value In, push func(ctx context.Context, value In, returnCh chan<- *Return[Out])) (context.Context, Out, error) {
	return PoolOf[In, Out]().CallAsSync(ctx, value, push)
}
