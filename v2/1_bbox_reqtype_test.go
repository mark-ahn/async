package async_test

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/mark-ahn/async/v2"
)

func TestBasic(t *testing.T) {
	worker_ctx, stop_worker := context.WithCancel(context.Background())
	var worker async.Worker[any, any] = async.NewChain(worker_ctx, new_square_brackets(), new_brackets())

	_, value, err := async.CallAsSync(context.TODO(), "50", worker.Push)
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

func push(ctx context.Context, value interface{}, rtnCh chan<- *async.Return[any]) *async.WorkContext[any, any] {
	pool := async.PoolOf[any, any]()
	return pool.Param.WorkContext.Pool.GetWith(ctx, pool.Param.Work.Pool.GetWith(value, rtnCh))
}

func TestAsyncLogic(t *testing.T) {
	worker_ctx, stop_worker := context.WithCancel(context.Background())
	square := new_square_brackets()
	round := new_brackets()

	pool := async.PoolOf[any, any]()

	requests := make(chan *async.WorkContext[any, any], 2)
	go func() {
		squared := pool.Return.ChanReturn.Pool.Get()
		rounded := pool.Return.ChanReturn.Pool.Get()
		defer func() {
			pool.Return.ChanReturn.Pool.Put(squared)
			pool.Return.ChanReturn.Pool.Put(rounded)
			fmt.Println("exit thread")
		}()
	loop:
		for {
			select {
			case request := <-requests:
				func() {
					ctx, value, returnCh := request.Unpack()
					defer pool.Param.WorkContext.Pool.Put(request)
					ctx = pool.Return.WithStack(ctx, 2)
					pool.Return.Push(ctx, returnCh)
					square.Push(ctx, value, squared)
				}()
			case square := <-squared:
				ctx, value, err := square.Unpack()
				if err != nil {
					pool.Return.Notify(ctx, square)
					continue
				}
				pool.Return.Return.Pool.Put(square)
				round.Push(ctx, value, rounded)
			case round := <-rounded:
				ctx, _, err := round.Unpack()
				if err != nil {
					pool.Return.Notify(ctx, round)
					continue
				}
				rtn_ch := pool.Return.Top(ctx)
				rtn_ch <- round

			case <-worker_ctx.Done():
				break loop
			}
		}
	}()

	expect := "([50])"
	rtn_ch := pool.Return.ChanReturn.Pool.Get()
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
