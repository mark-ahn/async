package async_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/mark-ahn/async"
)

func TestReqtype(t *testing.T) {
	worker_ctx, stop_worker := context.WithCancel(context.Background())
	var worker async.WorkerOfPushStringThenString = new_some_worker(worker_ctx)

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
