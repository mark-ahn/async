package async

import (
	"context"
	"sync"
)

type WorkerChain[Item any] struct {
	ctx context.Context

	pool    *AsyncPoolSet[Item, Item]
	doneCh  chan struct{}
	chains  []Worker[Item, Item]
	threads sync.WaitGroup
}

func NewWorkerChain[Item any](ctx context.Context, workers ...Worker[Item, Item]) *WorkerChain[Item] {
	__ := &WorkerChain[Item]{
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

func (__ *WorkerChain[Item]) Push(ctx context.Context, value Item, returnCh chan<- *Return[Item]) {
	__.threads.Add(1)
	select {
	case <-__.ctx.Done():
		defer __.threads.Done()
		var zero Item
		rtn := __.pool.Return.Value.Pack(ctx, zero, context.Canceled)
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
		ch := __.pool.Return.Chan.Get()
		defer __.pool.Return.Chan.Put(ch)
		var rvalue Item = arg
		for _, worker := range __.chains {
			worker.Push(ctx, rvalue, ch)
			rtn := <-ch
			ctx, rvalue, err = rtn.Unpack()
			if err != nil {
				returnCh <- rtn
				return
			}
			__.pool.Return.Value.Put(rtn)
		}
		rtn := __.pool.Return.Value.Pack(ctx, rvalue, err)
		returnCh <- rtn
	}()
}

func (__ *WorkerChain[Item]) DoneNotify() <-chan struct{} {
	return __.doneCh
}

func (__ *WorkerChain[Item]) Reset(_ context.Context) <-chan error {
	ch := make(chan error)
	close(ch)
	return ch
}
