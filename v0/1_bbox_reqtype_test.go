package async_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/mark-ahn/async/v0"
)

type worker interface {
	async.WorkerOfPushStringThenString
	async.DoneNotifier
}

func TestReqtype(t *testing.T) {
	worker_ctx, stop_worker := context.WithCancel(context.Background())
	var worker worker = new_some_worker(worker_ctx)

	res_ch := make(chan *async.ReturnOfString, 2)
	worker.Push(worker_ctx, &async.WorkOfStringThenString{
		Argument: "test",
		ReturnCh: res_ch,
	})
	worker.Push(worker_ctx, &async.WorkOfStringThenString{
		Argument: "how is it?",
		ReturnCh: res_ch,
	})

	stop_worker()
	<-worker.DoneNotify()

	close(res_ch)
	for d := range res_ch {
		fmt.Printf("%v\n", d)
	}
	fmt.Printf("done\n")
}

func TestWorker(t *testing.T) {
	worker_ctx, stop_worker := context.WithCancel(context.Background())
	worker := async.NewWorkerOfStringThenString(worker_ctx, func(ctx context.Context, str string) (string, error) {
		return "echo " + str, nil
	}, 10)
	defer func() {
		fmt.Printf("worker done\n")
	}()

	v, err := worker.Call(worker_ctx, "test")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", v)

	v, err = worker.Call(worker_ctx, "how is it?")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", v)

	res_ch := make(chan *async.ReturnOfString, 2)
	worker.Push(worker_ctx, &async.WorkOfStringThenString{
		Argument: "test",
		ReturnCh: res_ch,
	})

	stop_worker()
	<-worker.DoneNotify()

	close(res_ch)
	for d := range res_ch {
		fmt.Printf("%v\n", d)
	}

	fmt.Printf("done\n")
}
