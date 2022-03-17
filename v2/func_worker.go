package async

import (
	"context"
	"fmt"
	"sync"
)

type FuncWorker[Some any, Other any] struct {
	pool    *AsyncPoolSet[Some, Other]
	handler func(context.Context, Some) (context.Context, Other, error)

	ctx context.Context

	threads  sync.WaitGroup
	work_ch  chan *WorkContext[Some, Other]
	done_ch  chan struct{}
	reset_ch chan chan error
}

func NewFuncWorker[Some any, Other any](ctx context.Context, h func(context.Context, Some) (context.Context, Other, error), n int) *FuncWorker[Some, Other] {
	pool := PoolOf[Some, Other]()

	__ := &FuncWorker[Some, Other]{
		pool:    pool,
		handler: h,

		ctx: ctx,

		work_ch:  make(chan *WorkContext[Some, Other], n),
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
				go pool.CallAsAsync(ctx, value, rtn_ch, __.handler, func() {
					__.threads.Done()
				})
				pool.Param.puts(work)
			case reset_done_ch := <-__.reset_ch:
				__.reset_queue()
				close(reset_done_ch)
			}
		}
	}()

	return __
}

func (__ *FuncWorker[Some, Other]) reset_queue() {
	for i := 0; i < len(__.work_ch); i += 1 {
		req := <-__.work_ch
		rtn := __.pool.Return.Value.Get()
		rtn.Error = fmt.Errorf("canceled by reset")
		req.ReturnCh <- rtn
	}
}

func (__ *FuncWorker[Some, Other]) Push(ctx context.Context, value Some, returnCh chan<- *Return[Other]) {
	__.threads.Add(1)
	defer __.threads.Done()

	work_ctx := __.pool.Param.Pack(ctx, value, returnCh)
	__.work_ch <- work_ctx
}

func (__ *FuncWorker[Some, Other]) DoneNotify() <-chan struct{} {
	return __.done_ch
}

func (__ *FuncWorker[Some, Other]) Reset(ctx context.Context) <-chan error {
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
