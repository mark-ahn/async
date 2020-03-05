package async_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/mark-ahn/async/v1"
)

func TestReqtype(t *testing.T) {
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
