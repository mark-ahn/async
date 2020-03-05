// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package async

import (
	"context"
	"sync"
)

type ReturnOfBytes struct {
	Context context.Context
	Value   Bytes
	Error   error
}

var (
	zero_of_ReturnOfBytes       ReturnOfBytes
	zero_of_ReturnOfBytes_value Bytes
	pool_of_ReturnOfBytes       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfBytes{}
		},
	}
	pool_of_ReturnOfBytes_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfBytes, 1)
		},
	}
)

func getReturnOfBytes() *ReturnOfBytes {
	return pool_of_ReturnOfBytes.Get().(*ReturnOfBytes)
}
func putReturnOfBytes(d *ReturnOfBytes) {
	d.Context = nil
	d.Value = zero_of_ReturnOfBytes_value
	d.Error = nil
	pool_of_ReturnOfBytes.Put(d)
}

func getReturnChOfBytes() chan *ReturnOfBytes {
	return pool_of_ReturnOfBytes_ch.Get().(chan *ReturnOfBytes)
}
func putReturnChOfBytes(d chan *ReturnOfBytes) {
	pool_of_ReturnOfBytes_ch.Put(d)
}

type ReturnOfString struct {
	Context context.Context
	Value   string
	Error   error
}

var (
	zero_of_ReturnOfString       ReturnOfString
	zero_of_ReturnOfString_value string
	pool_of_ReturnOfString       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfString{}
		},
	}
	pool_of_ReturnOfString_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfString, 1)
		},
	}
)

func getReturnOfString() *ReturnOfString {
	return pool_of_ReturnOfString.Get().(*ReturnOfString)
}
func putReturnOfString(d *ReturnOfString) {
	d.Context = nil
	d.Value = zero_of_ReturnOfString_value
	d.Error = nil
	pool_of_ReturnOfString.Put(d)
}

func getReturnChOfString() chan *ReturnOfString {
	return pool_of_ReturnOfString_ch.Get().(chan *ReturnOfString)
}
func putReturnChOfString(d chan *ReturnOfString) {
	pool_of_ReturnOfString_ch.Put(d)
}

type ReturnOfInterface struct {
	Context context.Context
	Value   interface{}
	Error   error
}

var (
	zero_of_ReturnOfInterface       ReturnOfInterface
	zero_of_ReturnOfInterface_value interface{}
	pool_of_ReturnOfInterface       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfInterface{}
		},
	}
	pool_of_ReturnOfInterface_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfInterface, 1)
		},
	}
)

func getReturnOfInterface() *ReturnOfInterface {
	return pool_of_ReturnOfInterface.Get().(*ReturnOfInterface)
}
func putReturnOfInterface(d *ReturnOfInterface) {
	d.Context = nil
	d.Value = zero_of_ReturnOfInterface_value
	d.Error = nil
	pool_of_ReturnOfInterface.Put(d)
}

func getReturnChOfInterface() chan *ReturnOfInterface {
	return pool_of_ReturnOfInterface_ch.Get().(chan *ReturnOfInterface)
}
func putReturnChOfInterface(d chan *ReturnOfInterface) {
	pool_of_ReturnOfInterface_ch.Put(d)
}
