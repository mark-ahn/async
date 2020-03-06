package async_test

import (
	"context"
	"fmt"

	"github.com/mark-ahn/async/v1"
)

func new_squere_brackets() async.WorkerOfInterfaceToInterface {
	return async.InterfaceToInterface.FuncWorker.New(context.TODO(), func(ctx context.Context, str interface{}) (context.Context, interface{}, error) {
		return ctx, fmt.Sprintf("[%v]", str), nil
	}, 10)
}

func new_brackets() async.WorkerOfInterfaceToInterface {
	return async.InterfaceToInterface.FuncWorker.New(context.TODO(), func(ctx context.Context, str interface{}) (context.Context, interface{}, error) {
		return ctx, fmt.Sprintf("(%v)", str), nil
	}, 10)
}
