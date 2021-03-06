package async

import (
	"context"
	"fmt"
	"sync"
)

type FuncWorkerOfSomeToOther struct {
	handler func(context.Context, Some) (context.Context, Other, error)

	ctx context.Context

	threads  sync.WaitGroup
	work_ch  chan *WorkContextOfSomeToOther
	done_ch  chan struct{}
	reset_ch chan chan error
}

func newFuncWorkerOfSomeToOther(ctx context.Context, h func(context.Context, Some) (context.Context, Other, error), n int) *FuncWorkerOfSomeToOther {
	__ := &FuncWorkerOfSomeToOther{
		handler: h,

		ctx: ctx,

		work_ch:  make(chan *WorkContextOfSomeToOther, n),
		done_ch:  make(chan struct{}),
		reset_ch: make(chan chan error, n),
	}

	go func() {
		defer func() {
			__.threads.Wait()
			__.reset_queue()
			for i := 0; i < len(__.reset_ch); i += 1 {
				rst_req := <-__.reset_ch
				rst_req <- fmt.Errorf("terminated worker")
				close(rst_req)
			}
			close(__.done_ch)
		}()

	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			case work := <-__.work_ch:
				__.threads.Add(1)
				ctx, value, rtn_ch := work.Unpack()
				go SomeToOther.CallAsAsync(ctx, value, rtn_ch, __.handler, func() {
					__.threads.Done()
				})
				SomeToOther.WorkContext.Pool.Puts(work)
			case reset_done_ch := <-__.reset_ch:
				__.reset_queue()
				close(reset_done_ch)
			}
		}
	}()

	return __
}

func (__ *FuncWorkerOfSomeToOther) reset_queue() {
	for i := 0; i < len(__.work_ch); i += 1 {
		req := <-__.work_ch
		rtn := Others.Return.Pool.Get()
		rtn.Error = fmt.Errorf("canceled by reset")
		req.ReturnCh <- rtn
	}
}

func (__ *FuncWorkerOfSomeToOther) Push(ctx context.Context, value Some, returnCh chan<- *ReturnOfOther) {
	__.threads.Add(1)
	defer __.threads.Done()

	work_ctx := SomeToOther.WorkContext.Pool.GetWiths(ctx, value, returnCh)
	__.work_ch <- work_ctx
}

func (__ *FuncWorkerOfSomeToOther) DoneNotify() <-chan struct{} {
	return __.done_ch
}

func (__ *FuncWorkerOfSomeToOther) Reset(ctx context.Context) <-chan error {
	__.threads.Add(1)
	defer __.threads.Done()

	done_ch := make(chan error, 1)
	select {
	case <-__.ctx.Done():
		done_ch <- fmt.Errorf("contex done")
		close(done_ch)
	default:
		__.reset_ch <- done_ch
	}

	return done_ch
}
