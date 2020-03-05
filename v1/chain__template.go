package async

import (
	"context"
	"sync"
)

type ChainOfInterface struct {
	ctx     context.Context
	doneCh  chan struct{}
	chains  []WorkerOfInterfaceToInterface
	threads sync.WaitGroup
}

func NewChainOfInterface(ctx context.Context, workers ...WorkerOfInterfaceToInterface) *ChainOfInterface {
	__ := &ChainOfInterface{
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
func (__ *ChainOfInterface) Push(ctx context.Context, value interface{}, returnCh chan<- *ReturnOfInterface) {
	__.threads.Add(1)
	select {
	case <-__.ctx.Done():
		defer __.threads.Done()
		rtn := Interfaces.Return.Pool.GetWith(ctx, nil, context.Canceled)
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
		ch := Interfaces.ChanReturn.Pool.Get()
		defer Interfaces.ChanReturn.Pool.Put(ch)
		for _, worker := range __.chains {
			worker.Push(ctx, arg, ch)
			rtn := <-ch
			ctx, arg, err = rtn.Unpack()
			if err != nil {
				returnCh <- rtn
				return
			}
			Interfaces.Return.Pool.Put(rtn)
		}
		rtn := Interfaces.Return.Pool.GetWith(ctx, arg, err)
		returnCh <- rtn
	}()
}

func (__ *ChainOfInterface) DoneNotify() <-chan struct{} {
	return __.doneCh
}

func (__ *ChainOfInterface) Reset(_ context.Context) <-chan error {
	ch := make(chan error)
	close(ch)
	return ch
}
