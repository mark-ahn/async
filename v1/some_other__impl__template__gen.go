// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package async

import (
	"context"
	"fmt"
	"sync"
)

type FuncWorkerOfBytesToBytes struct {
	handler func(context.Context, Bytes) (context.Context, Bytes, error)

	ctx context.Context

	threads  sync.WaitGroup
	work_ch  chan *WorkContextOfBytesToBytes
	done_ch  chan struct{}
	reset_ch chan chan error
}

func newFuncWorkerOfBytesToBytes(ctx context.Context, h func(context.Context, Bytes) (context.Context, Bytes, error), n int) *FuncWorkerOfBytesToBytes {
	__ := &FuncWorkerOfBytesToBytes{
		handler: h,

		ctx: ctx,

		work_ch:  make(chan *WorkContextOfBytesToBytes, n),
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
				go BytesToBytes.CallAsAsync(ctx, value, rtn_ch, __.handler, func() {
					__.threads.Done()
				})
				BytesToBytes.WorkContext.Pool.Puts(work)
			case reset_done_ch := <-__.reset_ch:
				__.reset_queue()
				close(reset_done_ch)
			}
		}
	}()

	return __
}

func (__ *FuncWorkerOfBytesToBytes) reset_queue() {
	for i := 0; i < len(__.work_ch); i += 1 {
		req := <-__.work_ch
		rtn := Bytess.Return.Pool.Get()
		rtn.Error = fmt.Errorf("canceled by reset")
		req.ReturnCh <- rtn
	}
}

func (__ *FuncWorkerOfBytesToBytes) Push(ctx context.Context, value Bytes, returnCh chan<- *ReturnOfBytes) {
	__.threads.Add(1)
	defer __.threads.Done()

	work_ctx := BytesToBytes.WorkContext.Pool.GetWiths(ctx, value, returnCh)
	__.work_ch <- work_ctx
}

func (__ *FuncWorkerOfBytesToBytes) DoneNotify() <-chan struct{} {
	return __.done_ch
}

func (__ *FuncWorkerOfBytesToBytes) Reset(ctx context.Context) <-chan error {
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

type FuncWorkerOfBytesToString struct {
	handler func(context.Context, Bytes) (context.Context, string, error)

	ctx context.Context

	threads  sync.WaitGroup
	work_ch  chan *WorkContextOfBytesToString
	done_ch  chan struct{}
	reset_ch chan chan error
}

func newFuncWorkerOfBytesToString(ctx context.Context, h func(context.Context, Bytes) (context.Context, string, error), n int) *FuncWorkerOfBytesToString {
	__ := &FuncWorkerOfBytesToString{
		handler: h,

		ctx: ctx,

		work_ch:  make(chan *WorkContextOfBytesToString, n),
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
				go BytesToString.CallAsAsync(ctx, value, rtn_ch, __.handler, func() {
					__.threads.Done()
				})
				BytesToString.WorkContext.Pool.Puts(work)
			case reset_done_ch := <-__.reset_ch:
				__.reset_queue()
				close(reset_done_ch)
			}
		}
	}()

	return __
}

func (__ *FuncWorkerOfBytesToString) reset_queue() {
	for i := 0; i < len(__.work_ch); i += 1 {
		req := <-__.work_ch
		rtn := Strings.Return.Pool.Get()
		rtn.Error = fmt.Errorf("canceled by reset")
		req.ReturnCh <- rtn
	}
}

func (__ *FuncWorkerOfBytesToString) Push(ctx context.Context, value Bytes, returnCh chan<- *ReturnOfString) {
	__.threads.Add(1)
	defer __.threads.Done()

	work_ctx := BytesToString.WorkContext.Pool.GetWiths(ctx, value, returnCh)
	__.work_ch <- work_ctx
}

func (__ *FuncWorkerOfBytesToString) DoneNotify() <-chan struct{} {
	return __.done_ch
}

func (__ *FuncWorkerOfBytesToString) Reset(ctx context.Context) <-chan error {
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

type FuncWorkerOfBytesToInterface struct {
	handler func(context.Context, Bytes) (context.Context, interface{}, error)

	ctx context.Context

	threads  sync.WaitGroup
	work_ch  chan *WorkContextOfBytesToInterface
	done_ch  chan struct{}
	reset_ch chan chan error
}

func newFuncWorkerOfBytesToInterface(ctx context.Context, h func(context.Context, Bytes) (context.Context, interface{}, error), n int) *FuncWorkerOfBytesToInterface {
	__ := &FuncWorkerOfBytesToInterface{
		handler: h,

		ctx: ctx,

		work_ch:  make(chan *WorkContextOfBytesToInterface, n),
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
				go BytesToInterface.CallAsAsync(ctx, value, rtn_ch, __.handler, func() {
					__.threads.Done()
				})
				BytesToInterface.WorkContext.Pool.Puts(work)
			case reset_done_ch := <-__.reset_ch:
				__.reset_queue()
				close(reset_done_ch)
			}
		}
	}()

	return __
}

func (__ *FuncWorkerOfBytesToInterface) reset_queue() {
	for i := 0; i < len(__.work_ch); i += 1 {
		req := <-__.work_ch
		rtn := Interfaces.Return.Pool.Get()
		rtn.Error = fmt.Errorf("canceled by reset")
		req.ReturnCh <- rtn
	}
}

func (__ *FuncWorkerOfBytesToInterface) Push(ctx context.Context, value Bytes, returnCh chan<- *ReturnOfInterface) {
	__.threads.Add(1)
	defer __.threads.Done()

	work_ctx := BytesToInterface.WorkContext.Pool.GetWiths(ctx, value, returnCh)
	__.work_ch <- work_ctx
}

func (__ *FuncWorkerOfBytesToInterface) DoneNotify() <-chan struct{} {
	return __.done_ch
}

func (__ *FuncWorkerOfBytesToInterface) Reset(ctx context.Context) <-chan error {
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

type FuncWorkerOfStringToBytes struct {
	handler func(context.Context, string) (context.Context, Bytes, error)

	ctx context.Context

	threads  sync.WaitGroup
	work_ch  chan *WorkContextOfStringToBytes
	done_ch  chan struct{}
	reset_ch chan chan error
}

func newFuncWorkerOfStringToBytes(ctx context.Context, h func(context.Context, string) (context.Context, Bytes, error), n int) *FuncWorkerOfStringToBytes {
	__ := &FuncWorkerOfStringToBytes{
		handler: h,

		ctx: ctx,

		work_ch:  make(chan *WorkContextOfStringToBytes, n),
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
				go StringToBytes.CallAsAsync(ctx, value, rtn_ch, __.handler, func() {
					__.threads.Done()
				})
				StringToBytes.WorkContext.Pool.Puts(work)
			case reset_done_ch := <-__.reset_ch:
				__.reset_queue()
				close(reset_done_ch)
			}
		}
	}()

	return __
}

func (__ *FuncWorkerOfStringToBytes) reset_queue() {
	for i := 0; i < len(__.work_ch); i += 1 {
		req := <-__.work_ch
		rtn := Bytess.Return.Pool.Get()
		rtn.Error = fmt.Errorf("canceled by reset")
		req.ReturnCh <- rtn
	}
}

func (__ *FuncWorkerOfStringToBytes) Push(ctx context.Context, value string, returnCh chan<- *ReturnOfBytes) {
	__.threads.Add(1)
	defer __.threads.Done()

	work_ctx := StringToBytes.WorkContext.Pool.GetWiths(ctx, value, returnCh)
	__.work_ch <- work_ctx
}

func (__ *FuncWorkerOfStringToBytes) DoneNotify() <-chan struct{} {
	return __.done_ch
}

func (__ *FuncWorkerOfStringToBytes) Reset(ctx context.Context) <-chan error {
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

type FuncWorkerOfStringToString struct {
	handler func(context.Context, string) (context.Context, string, error)

	ctx context.Context

	threads  sync.WaitGroup
	work_ch  chan *WorkContextOfStringToString
	done_ch  chan struct{}
	reset_ch chan chan error
}

func newFuncWorkerOfStringToString(ctx context.Context, h func(context.Context, string) (context.Context, string, error), n int) *FuncWorkerOfStringToString {
	__ := &FuncWorkerOfStringToString{
		handler: h,

		ctx: ctx,

		work_ch:  make(chan *WorkContextOfStringToString, n),
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
				go StringToString.CallAsAsync(ctx, value, rtn_ch, __.handler, func() {
					__.threads.Done()
				})
				StringToString.WorkContext.Pool.Puts(work)
			case reset_done_ch := <-__.reset_ch:
				__.reset_queue()
				close(reset_done_ch)
			}
		}
	}()

	return __
}

func (__ *FuncWorkerOfStringToString) reset_queue() {
	for i := 0; i < len(__.work_ch); i += 1 {
		req := <-__.work_ch
		rtn := Strings.Return.Pool.Get()
		rtn.Error = fmt.Errorf("canceled by reset")
		req.ReturnCh <- rtn
	}
}

func (__ *FuncWorkerOfStringToString) Push(ctx context.Context, value string, returnCh chan<- *ReturnOfString) {
	__.threads.Add(1)
	defer __.threads.Done()

	work_ctx := StringToString.WorkContext.Pool.GetWiths(ctx, value, returnCh)
	__.work_ch <- work_ctx
}

func (__ *FuncWorkerOfStringToString) DoneNotify() <-chan struct{} {
	return __.done_ch
}

func (__ *FuncWorkerOfStringToString) Reset(ctx context.Context) <-chan error {
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

type FuncWorkerOfStringToInterface struct {
	handler func(context.Context, string) (context.Context, interface{}, error)

	ctx context.Context

	threads  sync.WaitGroup
	work_ch  chan *WorkContextOfStringToInterface
	done_ch  chan struct{}
	reset_ch chan chan error
}

func newFuncWorkerOfStringToInterface(ctx context.Context, h func(context.Context, string) (context.Context, interface{}, error), n int) *FuncWorkerOfStringToInterface {
	__ := &FuncWorkerOfStringToInterface{
		handler: h,

		ctx: ctx,

		work_ch:  make(chan *WorkContextOfStringToInterface, n),
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
				go StringToInterface.CallAsAsync(ctx, value, rtn_ch, __.handler, func() {
					__.threads.Done()
				})
				StringToInterface.WorkContext.Pool.Puts(work)
			case reset_done_ch := <-__.reset_ch:
				__.reset_queue()
				close(reset_done_ch)
			}
		}
	}()

	return __
}

func (__ *FuncWorkerOfStringToInterface) reset_queue() {
	for i := 0; i < len(__.work_ch); i += 1 {
		req := <-__.work_ch
		rtn := Interfaces.Return.Pool.Get()
		rtn.Error = fmt.Errorf("canceled by reset")
		req.ReturnCh <- rtn
	}
}

func (__ *FuncWorkerOfStringToInterface) Push(ctx context.Context, value string, returnCh chan<- *ReturnOfInterface) {
	__.threads.Add(1)
	defer __.threads.Done()

	work_ctx := StringToInterface.WorkContext.Pool.GetWiths(ctx, value, returnCh)
	__.work_ch <- work_ctx
}

func (__ *FuncWorkerOfStringToInterface) DoneNotify() <-chan struct{} {
	return __.done_ch
}

func (__ *FuncWorkerOfStringToInterface) Reset(ctx context.Context) <-chan error {
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

type FuncWorkerOfInterfaceToBytes struct {
	handler func(context.Context, interface{}) (context.Context, Bytes, error)

	ctx context.Context

	threads  sync.WaitGroup
	work_ch  chan *WorkContextOfInterfaceToBytes
	done_ch  chan struct{}
	reset_ch chan chan error
}

func newFuncWorkerOfInterfaceToBytes(ctx context.Context, h func(context.Context, interface{}) (context.Context, Bytes, error), n int) *FuncWorkerOfInterfaceToBytes {
	__ := &FuncWorkerOfInterfaceToBytes{
		handler: h,

		ctx: ctx,

		work_ch:  make(chan *WorkContextOfInterfaceToBytes, n),
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
				go InterfaceToBytes.CallAsAsync(ctx, value, rtn_ch, __.handler, func() {
					__.threads.Done()
				})
				InterfaceToBytes.WorkContext.Pool.Puts(work)
			case reset_done_ch := <-__.reset_ch:
				__.reset_queue()
				close(reset_done_ch)
			}
		}
	}()

	return __
}

func (__ *FuncWorkerOfInterfaceToBytes) reset_queue() {
	for i := 0; i < len(__.work_ch); i += 1 {
		req := <-__.work_ch
		rtn := Bytess.Return.Pool.Get()
		rtn.Error = fmt.Errorf("canceled by reset")
		req.ReturnCh <- rtn
	}
}

func (__ *FuncWorkerOfInterfaceToBytes) Push(ctx context.Context, value interface{}, returnCh chan<- *ReturnOfBytes) {
	__.threads.Add(1)
	defer __.threads.Done()

	work_ctx := InterfaceToBytes.WorkContext.Pool.GetWiths(ctx, value, returnCh)
	__.work_ch <- work_ctx
}

func (__ *FuncWorkerOfInterfaceToBytes) DoneNotify() <-chan struct{} {
	return __.done_ch
}

func (__ *FuncWorkerOfInterfaceToBytes) Reset(ctx context.Context) <-chan error {
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

type FuncWorkerOfInterfaceToString struct {
	handler func(context.Context, interface{}) (context.Context, string, error)

	ctx context.Context

	threads  sync.WaitGroup
	work_ch  chan *WorkContextOfInterfaceToString
	done_ch  chan struct{}
	reset_ch chan chan error
}

func newFuncWorkerOfInterfaceToString(ctx context.Context, h func(context.Context, interface{}) (context.Context, string, error), n int) *FuncWorkerOfInterfaceToString {
	__ := &FuncWorkerOfInterfaceToString{
		handler: h,

		ctx: ctx,

		work_ch:  make(chan *WorkContextOfInterfaceToString, n),
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
				go InterfaceToString.CallAsAsync(ctx, value, rtn_ch, __.handler, func() {
					__.threads.Done()
				})
				InterfaceToString.WorkContext.Pool.Puts(work)
			case reset_done_ch := <-__.reset_ch:
				__.reset_queue()
				close(reset_done_ch)
			}
		}
	}()

	return __
}

func (__ *FuncWorkerOfInterfaceToString) reset_queue() {
	for i := 0; i < len(__.work_ch); i += 1 {
		req := <-__.work_ch
		rtn := Strings.Return.Pool.Get()
		rtn.Error = fmt.Errorf("canceled by reset")
		req.ReturnCh <- rtn
	}
}

func (__ *FuncWorkerOfInterfaceToString) Push(ctx context.Context, value interface{}, returnCh chan<- *ReturnOfString) {
	__.threads.Add(1)
	defer __.threads.Done()

	work_ctx := InterfaceToString.WorkContext.Pool.GetWiths(ctx, value, returnCh)
	__.work_ch <- work_ctx
}

func (__ *FuncWorkerOfInterfaceToString) DoneNotify() <-chan struct{} {
	return __.done_ch
}

func (__ *FuncWorkerOfInterfaceToString) Reset(ctx context.Context) <-chan error {
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

type FuncWorkerOfInterfaceToInterface struct {
	handler func(context.Context, interface{}) (context.Context, interface{}, error)

	ctx context.Context

	threads  sync.WaitGroup
	work_ch  chan *WorkContextOfInterfaceToInterface
	done_ch  chan struct{}
	reset_ch chan chan error
}

func newFuncWorkerOfInterfaceToInterface(ctx context.Context, h func(context.Context, interface{}) (context.Context, interface{}, error), n int) *FuncWorkerOfInterfaceToInterface {
	__ := &FuncWorkerOfInterfaceToInterface{
		handler: h,

		ctx: ctx,

		work_ch:  make(chan *WorkContextOfInterfaceToInterface, n),
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
				go InterfaceToInterface.CallAsAsync(ctx, value, rtn_ch, __.handler, func() {
					__.threads.Done()
				})
				InterfaceToInterface.WorkContext.Pool.Puts(work)
			case reset_done_ch := <-__.reset_ch:
				__.reset_queue()
				close(reset_done_ch)
			}
		}
	}()

	return __
}

func (__ *FuncWorkerOfInterfaceToInterface) reset_queue() {
	for i := 0; i < len(__.work_ch); i += 1 {
		req := <-__.work_ch
		rtn := Interfaces.Return.Pool.Get()
		rtn.Error = fmt.Errorf("canceled by reset")
		req.ReturnCh <- rtn
	}
}

func (__ *FuncWorkerOfInterfaceToInterface) Push(ctx context.Context, value interface{}, returnCh chan<- *ReturnOfInterface) {
	__.threads.Add(1)
	defer __.threads.Done()

	work_ctx := InterfaceToInterface.WorkContext.Pool.GetWiths(ctx, value, returnCh)
	__.work_ch <- work_ctx
}

func (__ *FuncWorkerOfInterfaceToInterface) DoneNotify() <-chan struct{} {
	return __.done_ch
}

func (__ *FuncWorkerOfInterfaceToInterface) Reset(ctx context.Context) <-chan error {
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
