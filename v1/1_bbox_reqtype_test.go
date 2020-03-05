package async_test

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/mark-ahn/async/v1"
)

func TestBasic(t *testing.T) {
	worker_ctx, stop_worker := context.WithCancel(context.Background())
	var worker async.WorkerOfInterfaceToInterface = async.NewChainOfInterface(worker_ctx, new_squere_brackets(), new_brackets())

	_, value, err := async.InterfaceToInterface.CallAsSync(context.TODO(), "50", worker.Push)
	if err != nil {
		t.Fatal(err)
	}

	expect := "([50])"
	if !reflect.DeepEqual(value, expect) {
		t.Errorf("expect %v, got %v", expect, value)
	}
	stop_worker()
	fmt.Printf("done\n")
}

func push(ctx context.Context, value interface{}, rtnCh chan<- *async.ReturnOfInterface) *async.WorkContextOfInterfaceToInterface {
	return async.InterfaceToInterface.Pool.WorkContext.GetWith(ctx, async.InterfaceToInterface.Pool.Work.GetWith(value, rtnCh))
}

func TestAsyncLogic(t *testing.T) {
	worker_ctx, stop_worker := context.WithCancel(context.Background())
	square := new_squere_brackets()
	round := new_brackets()

	requests := make(chan *async.WorkContextOfInterfaceToInterface, 2)
	go func() {
		squared := async.Interfaces.ChanReturn.Pool.Get()
		rounded := async.Interfaces.ChanReturn.Pool.Get()
		defer func() {
			async.Interfaces.ChanReturn.Pool.Put(squared)
			async.Interfaces.ChanReturn.Pool.Put(rounded)
			fmt.Println("exit thread")
		}()
	loop:
		for {
			select {
			case request := <-requests:
				func() {
					ctx, value, returnCh := request.Unpack()
					defer async.InterfaceToInterface.Pool.WorkContext.Put(request)
					ctx = async.Interfaces.ChanReturn.Context.WithStack(ctx, 2)
					async.Interfaces.ChanReturn.Context.Push(ctx, returnCh)
					square.Push(ctx, value, squared)
				}()
			case square := <-squared:
				ctx, value, err := square.Unpack()
				if err != nil {
					async.Interfaces.ChanReturn.Context.Notify(ctx, square)
					continue
				}
				async.Interfaces.Return.Pool.Put(square)
				round.Push(ctx, value, rounded)
			case round := <-rounded:
				ctx, _, err := round.Unpack()
				if err != nil {
					async.Interfaces.ChanReturn.Context.Notify(ctx, round)
					continue
				}
				rtn_ch := async.Interfaces.ChanReturn.Context.Top(ctx)
				rtn_ch <- round

			case <-worker_ctx.Done():
				break loop
			}
		}
	}()

	expect := "([50])"
	rtn_ch := async.Interfaces.ChanReturn.Pool.Get()
	// requests <- async.InterfaceToInterface.GetWorkContextWith(context.TODO(), async.InterfaceToInterface.GetWorkWith("50", rtn_ch))
	requests <- push(context.TODO(), "50", rtn_ch)
	rtn := <-rtn_ch
	value := rtn.Value
	if !reflect.DeepEqual(value, expect) {
		t.Errorf("expect %v, got %v", expect, value)
	}

	for i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0} {
		requests <- push(context.TODO(), strconv.FormatInt(int64(i), 10), rtn_ch)
		rtn := <-rtn_ch
		fmt.Println(rtn.Value)
	}

	stop_worker()
	fmt.Printf("done\n")
}
