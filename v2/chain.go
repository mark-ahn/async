package async

import (
	"context"
	"sync"
)

type Chain[Item any] struct {
	ctx context.Context

	pool    *TypePoolSet[Item, Item]
	doneCh  chan struct{}
	chains  []Worker[Item, Item]
	threads sync.WaitGroup
}

func NewChain[Item any](ctx context.Context, workers ...Worker[Item, Item]) *Chain[Item] {
	__ := &Chain[Item]{
		pool:   PoolOf[Item, Item](),
		ctx:    ctx,
		doneCh: make(chan struct{}),
		chains: workers,
	}

	go func() {
		defer func() {
			__.threads.Wait()
			close(__.doneCh)
		}()
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			}
		}
	}()
	return __
}

func (__ *Chain[Item]) Push(ctx context.Context, value Item, returnCh chan<- *Return[Item]) {
	__.threads.Add(1)
	select {
	case <-__.ctx.Done():
		defer __.threads.Done()
		var zero Item
		rtn := __.pool.Return.Return.Pool.GetWith(ctx, zero, context.Canceled)
		returnCh <- rtn
		return
	default:
	}

	go func() {
		defer __.threads.Done()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		__.threads.Add(1)
		go func() {
			defer __.threads.Done()
			select {
			case <-__.ctx.Done():
				cancel()
			case <-ctx.Done():
			}
		}()

		var err error
		arg := value
		ch := __.pool.Return.ChanReturn.Pool.Get()
		defer __.pool.Return.ChanReturn.Pool.Put(ch)
		var rvalue Item = arg
		for _, worker := range __.chains {
			worker.Push(ctx, rvalue, ch)
			rtn := <-ch
			ctx, rvalue, err = rtn.Unpack()
			if err != nil {
				returnCh <- rtn
				return
			}
			__.pool.Return.Return.Pool.Put(rtn)
		}
		rtn := __.pool.Return.Return.Pool.GetWith(ctx, rvalue, err)
		returnCh <- rtn
	}()
}

func (__ *Chain[Item]) DoneNotify() <-chan struct{} {
	return __.doneCh
}

func (__ *Chain[Item]) Reset(_ context.Context) <-chan error {
	ch := make(chan error)
	close(ch)
	return ch
}
