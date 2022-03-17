package async

import (
	"context"
	"sync"
)

type Work[In any, Out any] struct {
	Value    In
	ReturnCh chan<- *Return[Out]
}

type WorkContext[In any, Out any] struct {
	Context context.Context
	*Work[In, Out]
}

func (__ *WorkContext[In, Out]) Unpack() (context.Context, In, chan<- *Return[Out]) {
	return __.Context, __.Value, __.ReturnCh
}

type pool_work[In any, Out any] sync.Pool

func new_pool_work[In any, Out any]() *pool_work[In, Out] {
	pool := &sync.Pool{
		New: func() any {
			var zero Work[In, Out]
			return &zero
		},
	}
	return (*pool_work[In, Out])(pool)
}
func (__ *pool_work[In, Out]) pool() *sync.Pool { return (*sync.Pool)(__) }

func (__ *pool_work[In, Out]) Get() *Work[In, Out] {
	return __.pool().Get().(*Work[In, Out])
}
func (__ *pool_work[In, Out]) Put(d *Work[In, Out]) {
	var zero In
	d.Value = zero
	d.ReturnCh = nil
	__.pool().Put(d)
}

func (__ pool_work[In, Out]) Pack(value In, returnCh chan<- *Return[Out]) *Work[In, Out] {
	work := __.Get()
	work.Value = value
	work.ReturnCh = returnCh
	return work
}

type pool_work_context[In any, Out any] sync.Pool

func new_pool_work_context[In any, Out any]() *pool_work_context[In, Out] {
	return (*pool_work_context[In, Out])(&sync.Pool{
		New: func() any {
			var zero WorkContext[In, Out]
			return &zero
		},
	})
}
func (__ *pool_work_context[In, Out]) pool() *sync.Pool { return (*sync.Pool)(__) }

func (__ *pool_work_context[In, Out]) Get() *WorkContext[In, Out] {
	return __.pool().Get().(*WorkContext[In, Out])
}
func (__ *pool_work_context[In, Out]) Put(d *WorkContext[In, Out]) {
	d.Context = nil
	d.Work = nil
	__.pool().Put(d)
}

func (__ *pool_work_context[In, Out]) Pack(ctx context.Context, work *Work[In, Out]) *WorkContext[In, Out] {
	work_ctx := __.Get()
	work_ctx.Context = ctx
	work_ctx.Work = work
	return work_ctx
}

func NewParamPoolSet[In any, Out any]() *ParamPoolSet[In, Out] {
	return &ParamPoolSet[In, Out]{
		Work:        new_pool_work[In, Out](),
		WorkContext: new_pool_work_context[In, Out](),
	}
}

type ParamPoolSet[In any, Out any] struct {
	Work        *pool_work[In, Out]
	WorkContext *pool_work_context[In, Out]
	FuncWorker  func_worker[In, Out]
}

func (__ ParamPoolSet[In, Out]) Pack(ctx context.Context, value In, returnCh chan<- *Return[Out]) *WorkContext[In, Out] {
	return __.WorkContext.Pack(ctx, __.Work.Pack(value, returnCh))
}

func (__ ParamPoolSet[In, Out]) Collect(d *WorkContext[In, Out]) (context.Context, In, chan<- *Return[Out]) {
	ctx, value, rtn_ch := d.Unpack()
	__.puts(d)
	return ctx, value, rtn_ch
}

func (__ ParamPoolSet[In, Out]) puts(d *WorkContext[In, Out]) {
	__.Work.Put(d.Work)
	__.WorkContext.Put(d)
}
