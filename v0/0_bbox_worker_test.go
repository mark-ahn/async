package async_test

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mark-ahn/async/v0"
)

type some_worker struct {
	threads sync.WaitGroup
	work_ch chan *async.WorkOfStringThenString
	done_ch chan struct{}
}

func new_some_worker(ctx context.Context) *some_worker {
	__ := &some_worker{
		work_ch: make(chan *async.WorkOfStringThenString, 1),
		done_ch: make(chan struct{}),
	}
	go func() {
		defer func() {
			for i := 0; i < len(__.work_ch); i += 1 {
				work := <-__.work_ch
				work.ReturnCh <- &async.ReturnOfString{
					Error: fmt.Errorf("worker is terminated"),
				}
			}
			__.threads.Wait()
			close(__.done_ch)
		}()

	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			case work := <-__.work_ch:
				__.threads.Add(1)
				go func(s string) {
					defer __.threads.Done()
					<-time.After(time.Second)
					work.ReturnCh <- &async.ReturnOfString{
						Value: "echo " + s,
						Error: nil,
					}
				}(work.Argument)
			}
		}
	}()
	return __
}

func (__ *some_worker) Push(ctx context.Context, req *async.WorkOfStringThenString) {
	select {
	case <-ctx.Done():
		req.ReturnCh <- &async.ReturnOfString{
			Error: fmt.Errorf("context done"),
		}
	case __.work_ch <- req:
	}
}
func (__ *some_worker) DoneNotify() <-chan struct{} {
	return __.done_ch
}
