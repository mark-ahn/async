package async_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/mark-ahn/async/v2"
)

type type_param[T any, U any] struct{}

func Test_TypeOf(t *testing.T) {
	type_1 := reflect.TypeOf(type_param[any, any]{})
	type_2 := reflect.TypeOf(type_param[int, float32]{})
	type_3 := reflect.TypeOf(type_param[int, float32]{})
	fmt.Printf("%+v %+v %+v\n", type_1, type_2, type_2 == type_3)
}

func new_square_brackets() async.Worker[any, any] {
	return async.NewFuncWorker(context.TODO(), func(ctx context.Context, str interface{}) (context.Context, interface{}, error) {
		// fmt.Println("square", str)
		return ctx, fmt.Sprintf("[%v]", str), nil
	}, 10)
}

func new_brackets() async.Worker[any, any] {
	return async.NewFuncWorker(context.TODO(), func(ctx context.Context, str interface{}) (context.Context, interface{}, error) {
		// fmt.Println("curl", str)
		return ctx, fmt.Sprintf("(%v)", str), nil
	}, 10)
}
