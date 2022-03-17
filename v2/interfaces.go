package async

import "context"

type Valuable interface {
	Value(interface{}) interface{}
}

type Worker[Some any, Other any] interface {
	Push(ctx context.Context, value Some, returnCh chan<- *Return[Other])
}
