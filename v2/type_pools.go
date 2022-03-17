package async

import (
	"context"
	"reflect"
)

func PoolOf[Some any, Other any]() *TypePoolSet[Some, Other] {
	var pool *TypePoolSet[Some, Other]
	{
		type_p := reflect.TypeOf(type_param[Some, Other]{})
		pool_lock.Lock()
		pool_i, ok := pools[type_p]
		if !ok {
			// pool = &TypePoolSet[Some, Other]{}
			pool = NewTypePoolSet[Some, Other]()
			pools[type_p] = pool
		} else {
			pool = pool_i.(*TypePoolSet[Some, Other])
		}
		pool_lock.Unlock()
	}
	return pool
}

func CallAsAsync[Some any, Other any](ctx context.Context, value Some, returnCh chan<- *Return[Other], h func(ctx context.Context, arg Some) (context.Context, Other, error), defered func()) {
	PoolOf[Some, Other]().CallAsAsync(ctx, value, returnCh, h, defered)
}

func CallAsSync[Some any, Other any](ctx context.Context, value Some, push func(ctx context.Context, value Some, returnCh chan<- *Return[Other])) (context.Context, Other, error) {
	return PoolOf[Some, Other]().CallAsSync(ctx, value, push)
}
