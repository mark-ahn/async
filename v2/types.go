package async

import (
	"context"
	"sync"
)

type Worker[Some any, Other any] interface {
	Push(ctx context.Context, value Some, returnCh chan<- *Return[Other])
}

type Work[Some any, Other any] struct {
	Value    Some
	ReturnCh chan<- *Return[Other]
}

type WorkContext[Some any, Other any] struct {
	Context context.Context
	*Work[Some, Other]
}

func (__ *WorkContext[Some, Other]) Unpack() (context.Context, Some, chan<- *Return[Other]) {
	return __.Context, __.Value, __.ReturnCh
}

// var (
// 	zero_of_WorkOfSomeToOther       Work
// 	zero_of_WorkOfSomeToOther_Value Some
// 	pool_of_WorkOfSomeToOther       = sync.Pool{
// 		New: func() interface{} {
// 			return &Work{}
// 		},
// 	}
// 	pool_of_WorkOfSomeToOtherContext = sync.Pool{
// 		New: func() interface{} {
// 			return &WorkContext{}
// 		},
// 	}
// )

type pool_Work[Some any, Other any] struct {
	pool sync.Pool
}

func (__ *pool_Work[Some, Other]) Get() *Work[Some, Other] {
	return __.pool.Get().(*Work[Some, Other])
}
func (__ *pool_Work[Some, Other]) Put(d *Work[Some, Other]) {
	var zero Some
	d.Value = zero
	d.ReturnCh = nil
	__.pool.Put(d)
}

func (__ pool_Work[Some, Other]) GetWith(value Some, returnCh chan<- *Return[Other]) *Work[Some, Other] {
	work := __.Get()
	work.Value = value
	work.ReturnCh = returnCh
	return work
}

type pool_WorkContext[Some any, Other any] struct {
	pool sync.Pool
}

func (__ *pool_WorkContext[Some, Other]) Get() *WorkContext[Some, Other] {
	return __.pool.Get().(*WorkContext[Some, Other])
}
func (__ *pool_WorkContext[Some, Other]) Put(d *WorkContext[Some, Other]) {
	d.Context = nil
	d.Work = nil
	__.pool.Put(d)
}

func (__ *pool_WorkContext[Some, Other]) GetWith(ctx context.Context, work *Work[Some, Other]) *WorkContext[Some, Other] {
	work_ctx := __.Get()
	work_ctx.Context = ctx
	work_ctx.Work = work
	return work_ctx
}

func NewParamPoolSet[Some any, Other any]() *ParamPoolSet[Some, Other] {
	return &ParamPoolSet[Some, Other]{
		Work: struct{ Pool pool_Work[Some, Other] }{
			Pool: pool_Work[Some, Other]{
				pool: sync.Pool{
					New: func() any {
						var zero Work[Some, Other]
						return &zero
					},
				},
			},
		},
		WorkContext: struct{ Pool pool_WorkContext[Some, Other] }{
			Pool: pool_WorkContext[Some, Other]{
				pool: sync.Pool{
					New: func() any {
						var zero WorkContext[Some, Other]
						return &zero
					},
				},
			},
		},
	}
}

type ParamPoolSet[Some any, Other any] struct {
	Work struct {
		Pool pool_Work[Some, Other]
	}
	WorkContext struct {
		Pool pool_WorkContext[Some, Other]
	}
	FuncWorker func_worker[Some, Other]
}

func (__ ParamPoolSet[Some, Other]) GetWiths(ctx context.Context, value Some, returnCh chan<- *Return[Other]) *WorkContext[Some, Other] {
	work := __.Work.Pool.Get()
	work.Value = value
	work.ReturnCh = returnCh
	return __.WorkContext.Pool.GetWith(ctx, work)
}

func (__ ParamPoolSet[Some, Other]) Collect(d *WorkContext[Some, Other]) (context.Context, Some, chan<- *Return[Other]) {
	ctx, value, rtn_ch := d.Unpack()
	__.Puts(d)
	return ctx, value, rtn_ch
}

func (__ ParamPoolSet[Some, Other]) Puts(d *WorkContext[Some, Other]) {
	var zero Some
	work := d.Work
	work.Value = zero
	work.ReturnCh = nil
	__.Work.Pool.Put(work)

	__.WorkContext.Pool.Put(d)
}

type func_worker[Some any, Other any] struct{}

func (_ func_worker[Some, Other]) New(ctx context.Context, h func(context.Context, Some) (context.Context, Other, error), n int) *FuncWorker[Some, Other] {
	return NewFuncWorker(ctx, h, n)
}

// type _SomeToOther[Some any, Other any] struct {
// 	Work struct {
// 		Pool pool_Work
// 	}
// 	WorkContext struct {
// 		Pool pool_WorkContext
// 	}
// 	FuncWorker func_worker[Some, Other]
// }

type TypePoolSet[Some any, Other any] struct {
	Param  ParamPoolSet[Some, Other]
	Return ReturnPoolSet[Other]
}

func NewTypePoolSet[Some any, Other any]() *TypePoolSet[Some, Other] {
	return &TypePoolSet[Some, Other]{
		Param:  *NewParamPoolSet[Some, Other](),
		Return: *NewReturnPoolSet[Other](),
	}
}

func (__ TypePoolSet[Some, Other]) CallAsSync(ctx context.Context, value Some, push func(ctx context.Context, value Some, returnCh chan<- *Return[Other])) (context.Context, Other, error) {
	ch := __.Return.ChanReturn.Pool.Get()
	defer __.Return.ChanReturn.Pool.Put(ch)

	push(ctx, value, ch)
	rtn := <-ch
	return rtn.Context, rtn.Value, rtn.Error
}

func (__ TypePoolSet[Some, Other]) CallAsAsync(ctx context.Context, value Some, returnCh chan<- *Return[Other], h func(ctx context.Context, arg Some) (context.Context, Other, error), defered func()) {
	go func() {
		defer defered()

		ctx, res, err := h(ctx, value)
		rtn := __.Return.Return.Pool.GetWith(ctx, res, err)
		returnCh <- rtn
	}()
}

// var SomeToOther = _SomeToOther[Some, Other]{}
