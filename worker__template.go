package async

import (
	"context"
	"fmt"
	"sync"
)

type PoolOfSomeThenOther interface {
	// GetSomeArg() Some
	// PutSomeArg(Some)
	// GetOtherReturn() Other
	// PutOtherReturn(Other)

	GetWorkOfSomeThenOther() *WorkOfSomeThenOther
	GetReturnOfOther() *ReturnOfOther
	PutWorkOfSomeThenOther(*WorkOfSomeThenOther)
	PutReturnOfOther(*ReturnOfOther)

	GetChReturnOfOther() chan *ReturnOfOther
	PutChReturnOfOther(chan *ReturnOfOther)
}

type PoolOfSomeThenOtherImpl struct {
	SomeArg         sync.Pool
	OtherRtn        sync.Pool
	SomeWork        sync.Pool
	OtherReturn     sync.Pool
	ChReturnOfOther sync.Pool
}

func (__ *PoolOfSomeThenOtherImpl) GetSomeArg() *Some {
	return __.SomeArg.Get().(*Some)
}
func (__ *PoolOfSomeThenOtherImpl) PutSomeArg(p *Some) {
	__.SomeArg.Put(p)
}
func (__ *PoolOfSomeThenOtherImpl) GetOtherReturn() *Other {
	return __.OtherRtn.Get().(*Other)
}
func (__ *PoolOfSomeThenOtherImpl) PutOtherReturn(p *Other) {
	__.OtherRtn.Put(p)
}

func (__ *PoolOfSomeThenOtherImpl) GetWorkOfSomeThenOther() *WorkOfSomeThenOther {
	return __.SomeWork.Get().(*WorkOfSomeThenOther)
}
func (__ *PoolOfSomeThenOtherImpl) PutWorkOfSomeThenOther(p *WorkOfSomeThenOther) {
	__.SomeWork.Put(p)
}
func (__ *PoolOfSomeThenOtherImpl) GetReturnOfOther() *ReturnOfOther {
	return __.OtherReturn.Get().(*ReturnOfOther)
}
func (__ *PoolOfSomeThenOtherImpl) PutReturnOfOther(p *ReturnOfOther) {
	__.OtherReturn.Put(p)
}

func (__ *PoolOfSomeThenOtherImpl) GetChReturnOfOther() chan *ReturnOfOther {
	return __.ChReturnOfOther.Get().(chan *ReturnOfOther)
}
func (__ *PoolOfSomeThenOtherImpl) PutChReturnOfOther(p chan *ReturnOfOther) {
	__.ChReturnOfOther.Put(p)
}

func NewPoolOfSomeThenOtherImpl() *PoolOfSomeThenOtherImpl {
	return &PoolOfSomeThenOtherImpl{
		SomeArg: sync.Pool{
			New: func() interface{} {
				return new(Some)
			},
		},
		SomeWork: sync.Pool{
			New: func() interface{} {
				return new(WorkOfSomeThenOther)
			},
		},
		OtherRtn: sync.Pool{
			New: func() interface{} {
				return new(Other)
			},
		},
		OtherReturn: sync.Pool{
			New: func() interface{} {
				return new(ReturnOfOther)
			},
		},
		ChReturnOfOther: sync.Pool{
			New: func() interface{} {
				return make(chan *ReturnOfOther, 1)
			},
		},
	}
}

type ApiDecouplerOfSomeThenOther struct {
	pool PoolOfSomeThenOther
}

func NewApiDecouplerOfSomeThenOther() *ApiDecouplerOfSomeThenOther {
	return &ApiDecouplerOfSomeThenOther{
		pool: NewPoolOfSomeThenOtherImpl(),
	}
}

func (__ *ApiDecouplerOfSomeThenOther) ReturnError(rtn_ch chan<- *ReturnOfOther, err error) {
	rtn := __.pool.GetReturnOfOther()
	var other Other
	rtn.Value = other
	rtn.Error = err
	rtn_ch <- rtn
}

func (__ *ApiDecouplerOfSomeThenOther) Handle(ctx context.Context, req *WorkOfSomeThenOther, h func(ctx context.Context, arg Some) (Other, error), defered func()) {
	defer defered()

	rtn := __.pool.GetReturnOfOther()

	res, err := h(ctx, req.Argument)
	if err != nil {
		rtn.Error = err
		req.ReturnCh <- rtn
		return
	}

	rtn.Value = res
	rtn.Error = nil
	req.ReturnCh <- rtn
}

func (__ *ApiDecouplerOfSomeThenOther) Call(ctx context.Context, worker WorkerOfPushSomeThenOther, arg Some) (Other, error) {
	ch := __.pool.GetChReturnOfOther()
	defer __.pool.PutChReturnOfOther(ch)

	req := __.pool.GetWorkOfSomeThenOther()
	req.Argument = arg
	req.ReturnCh = ch

	worker.Push(ctx, req)
	defer func() {
		req.ReturnCh = nil
		__.pool.PutWorkOfSomeThenOther(req)
	}()

	res := <-ch
	defer func() {
		res.Error = nil
		__.pool.PutReturnOfOther(res)
	}()
	return res.Value, res.Error
}

type WorkHandlerOfSomeThenOther = func(ctx context.Context, arg Some) (Other, error)

// type CallerOfSomeThenOther interface {
// 	Call(ctx context.Context, arg Some) (Other, error)
// }

type WorkerOfSomeThenOther struct {
	api_syncer *ApiDecouplerOfSomeThenOther
	handler    WorkHandlerOfSomeThenOther

	ctx context.Context

	threads  sync.WaitGroup
	req_ch   chan *WorkWithContextOfSomeThenOther
	done_ch  chan struct{}
	reset_ch chan chan error

	pool sync.Pool
}

func NewWorkerOfSomeThenOther(ctx context.Context, h WorkHandlerOfSomeThenOther, n int) *WorkerOfSomeThenOther {
	__ := &WorkerOfSomeThenOther{
		api_syncer: NewApiDecouplerOfSomeThenOther(),
		handler:    h,

		ctx: ctx,

		req_ch:   make(chan *WorkWithContextOfSomeThenOther, n),
		done_ch:  make(chan struct{}),
		reset_ch: make(chan chan error, n),

		pool: sync.Pool{
			New: func() interface{} {
				return &WorkWithContextOfSomeThenOther{}
			},
		},
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
			case req := <-__.req_ch:
				__.threads.Add(1)
				go __.api_syncer.Handle(req.Context, req.WorkOfSomeThenOther, __.handler, func() {
					__.pool.Put(req)
					__.threads.Done()
				})
			case reset_done_ch := <-__.reset_ch:
				__.reset_queue()
				close(reset_done_ch)
			}
		}
	}()

	return __
}

func (__ *WorkerOfSomeThenOther) reset_queue() {
	for i := 0; i < len(__.req_ch); i += 1 {
		req := <-__.req_ch
		__.api_syncer.ReturnError(req.ReturnCh, fmt.Errorf("canceled by reset"))
	}
}

func (__ *WorkerOfSomeThenOther) Push(ctx context.Context, req *WorkOfSomeThenOther) {
	__.threads.Add(1)
	defer __.threads.Done()

	select {
	case <-__.ctx.Done():
		__.api_syncer.ReturnError(req.ReturnCh, fmt.Errorf("terminated worker"))
		return
	default:
	}

	req_n_ctx := __.pool.Get().(*WorkWithContextOfSomeThenOther)
	req_n_ctx.Context = ctx
	req_n_ctx.WorkOfSomeThenOther = req
	__.req_ch <- req_n_ctx
}

func (__ *WorkerOfSomeThenOther) DoneNotify() <-chan struct{} {
	return __.done_ch
}

func (__ *WorkerOfSomeThenOther) Reset(ctx context.Context) <-chan error {
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

func (__ *WorkerOfSomeThenOther) Call(ctx context.Context, arg Some) (Other, error) {
	return __.api_syncer.Call(ctx, __, arg)
}
