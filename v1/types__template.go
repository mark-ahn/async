package async

import (
	"context"
	"sync"
)

type RequestOfSomeToOther struct {
	Context  context.Context
	Value    Some
	ReturnCh chan<- ReturnOfSome
}

var (
	pool_of_RequestOfSomeToOther = sync.Pool{
		New: func() interface{} {
			return &RequestOfSomeToOther{}
		},
	}
)

func GetRequestOfSomeToOther() *RequestOfSomeToOther {
	return pool_of_RequestOfSomeToOther.Get().(*RequestOfSomeToOther)
}
func PutRequestOfSomeToOther(d *RequestOfSomeToOther) {
	pool_of_RequestOfSomeToOther.Put(d)
}
