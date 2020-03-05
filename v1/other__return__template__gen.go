// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package async

import (
	"context"
	"sync"
)

type ReturnOfBool struct {
	Context context.Context
	Value   bool
	Error   error
}

var (
	zero_of_ReturnOfBool       ReturnOfBool
	zero_of_ReturnOfBool_value bool
	pool_of_ReturnOfBool       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfBool{}
		},
	}
	pool_of_ReturnOfBool_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfBool, 1)
		},
	}
)

func getReturnOfBool() *ReturnOfBool {
	return pool_of_ReturnOfBool.Get().(*ReturnOfBool)
}
func putReturnOfBool(d *ReturnOfBool) {
	d.Context = nil
	d.Value = zero_of_ReturnOfBool_value
	d.Error = nil
	pool_of_ReturnOfBool.Put(d)
}

func getReturnChOfBool() chan *ReturnOfBool {
	return pool_of_ReturnOfBool_ch.Get().(chan *ReturnOfBool)
}
func putReturnChOfBool(d chan *ReturnOfBool) {
	pool_of_ReturnOfBool_ch.Put(d)
}

type ReturnOfByte struct {
	Context context.Context
	Value   byte
	Error   error
}

var (
	zero_of_ReturnOfByte       ReturnOfByte
	zero_of_ReturnOfByte_value byte
	pool_of_ReturnOfByte       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfByte{}
		},
	}
	pool_of_ReturnOfByte_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfByte, 1)
		},
	}
)

func getReturnOfByte() *ReturnOfByte {
	return pool_of_ReturnOfByte.Get().(*ReturnOfByte)
}
func putReturnOfByte(d *ReturnOfByte) {
	d.Context = nil
	d.Value = zero_of_ReturnOfByte_value
	d.Error = nil
	pool_of_ReturnOfByte.Put(d)
}

func getReturnChOfByte() chan *ReturnOfByte {
	return pool_of_ReturnOfByte_ch.Get().(chan *ReturnOfByte)
}
func putReturnChOfByte(d chan *ReturnOfByte) {
	pool_of_ReturnOfByte_ch.Put(d)
}

type ReturnOfComplex128 struct {
	Context context.Context
	Value   complex128
	Error   error
}

var (
	zero_of_ReturnOfComplex128       ReturnOfComplex128
	zero_of_ReturnOfComplex128_value complex128
	pool_of_ReturnOfComplex128       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfComplex128{}
		},
	}
	pool_of_ReturnOfComplex128_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfComplex128, 1)
		},
	}
)

func getReturnOfComplex128() *ReturnOfComplex128 {
	return pool_of_ReturnOfComplex128.Get().(*ReturnOfComplex128)
}
func putReturnOfComplex128(d *ReturnOfComplex128) {
	d.Context = nil
	d.Value = zero_of_ReturnOfComplex128_value
	d.Error = nil
	pool_of_ReturnOfComplex128.Put(d)
}

func getReturnChOfComplex128() chan *ReturnOfComplex128 {
	return pool_of_ReturnOfComplex128_ch.Get().(chan *ReturnOfComplex128)
}
func putReturnChOfComplex128(d chan *ReturnOfComplex128) {
	pool_of_ReturnOfComplex128_ch.Put(d)
}

type ReturnOfComplex64 struct {
	Context context.Context
	Value   complex64
	Error   error
}

var (
	zero_of_ReturnOfComplex64       ReturnOfComplex64
	zero_of_ReturnOfComplex64_value complex64
	pool_of_ReturnOfComplex64       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfComplex64{}
		},
	}
	pool_of_ReturnOfComplex64_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfComplex64, 1)
		},
	}
)

func getReturnOfComplex64() *ReturnOfComplex64 {
	return pool_of_ReturnOfComplex64.Get().(*ReturnOfComplex64)
}
func putReturnOfComplex64(d *ReturnOfComplex64) {
	d.Context = nil
	d.Value = zero_of_ReturnOfComplex64_value
	d.Error = nil
	pool_of_ReturnOfComplex64.Put(d)
}

func getReturnChOfComplex64() chan *ReturnOfComplex64 {
	return pool_of_ReturnOfComplex64_ch.Get().(chan *ReturnOfComplex64)
}
func putReturnChOfComplex64(d chan *ReturnOfComplex64) {
	pool_of_ReturnOfComplex64_ch.Put(d)
}

type ReturnOfError struct {
	Context context.Context
	Value   error
	Error   error
}

var (
	zero_of_ReturnOfError       ReturnOfError
	zero_of_ReturnOfError_value error
	pool_of_ReturnOfError       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfError{}
		},
	}
	pool_of_ReturnOfError_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfError, 1)
		},
	}
)

func getReturnOfError() *ReturnOfError {
	return pool_of_ReturnOfError.Get().(*ReturnOfError)
}
func putReturnOfError(d *ReturnOfError) {
	d.Context = nil
	d.Value = zero_of_ReturnOfError_value
	d.Error = nil
	pool_of_ReturnOfError.Put(d)
}

func getReturnChOfError() chan *ReturnOfError {
	return pool_of_ReturnOfError_ch.Get().(chan *ReturnOfError)
}
func putReturnChOfError(d chan *ReturnOfError) {
	pool_of_ReturnOfError_ch.Put(d)
}

type ReturnOfFloat32 struct {
	Context context.Context
	Value   float32
	Error   error
}

var (
	zero_of_ReturnOfFloat32       ReturnOfFloat32
	zero_of_ReturnOfFloat32_value float32
	pool_of_ReturnOfFloat32       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfFloat32{}
		},
	}
	pool_of_ReturnOfFloat32_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfFloat32, 1)
		},
	}
)

func getReturnOfFloat32() *ReturnOfFloat32 {
	return pool_of_ReturnOfFloat32.Get().(*ReturnOfFloat32)
}
func putReturnOfFloat32(d *ReturnOfFloat32) {
	d.Context = nil
	d.Value = zero_of_ReturnOfFloat32_value
	d.Error = nil
	pool_of_ReturnOfFloat32.Put(d)
}

func getReturnChOfFloat32() chan *ReturnOfFloat32 {
	return pool_of_ReturnOfFloat32_ch.Get().(chan *ReturnOfFloat32)
}
func putReturnChOfFloat32(d chan *ReturnOfFloat32) {
	pool_of_ReturnOfFloat32_ch.Put(d)
}

type ReturnOfFloat64 struct {
	Context context.Context
	Value   float64
	Error   error
}

var (
	zero_of_ReturnOfFloat64       ReturnOfFloat64
	zero_of_ReturnOfFloat64_value float64
	pool_of_ReturnOfFloat64       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfFloat64{}
		},
	}
	pool_of_ReturnOfFloat64_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfFloat64, 1)
		},
	}
)

func getReturnOfFloat64() *ReturnOfFloat64 {
	return pool_of_ReturnOfFloat64.Get().(*ReturnOfFloat64)
}
func putReturnOfFloat64(d *ReturnOfFloat64) {
	d.Context = nil
	d.Value = zero_of_ReturnOfFloat64_value
	d.Error = nil
	pool_of_ReturnOfFloat64.Put(d)
}

func getReturnChOfFloat64() chan *ReturnOfFloat64 {
	return pool_of_ReturnOfFloat64_ch.Get().(chan *ReturnOfFloat64)
}
func putReturnChOfFloat64(d chan *ReturnOfFloat64) {
	pool_of_ReturnOfFloat64_ch.Put(d)
}

type ReturnOfInt struct {
	Context context.Context
	Value   int
	Error   error
}

var (
	zero_of_ReturnOfInt       ReturnOfInt
	zero_of_ReturnOfInt_value int
	pool_of_ReturnOfInt       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfInt{}
		},
	}
	pool_of_ReturnOfInt_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfInt, 1)
		},
	}
)

func getReturnOfInt() *ReturnOfInt {
	return pool_of_ReturnOfInt.Get().(*ReturnOfInt)
}
func putReturnOfInt(d *ReturnOfInt) {
	d.Context = nil
	d.Value = zero_of_ReturnOfInt_value
	d.Error = nil
	pool_of_ReturnOfInt.Put(d)
}

func getReturnChOfInt() chan *ReturnOfInt {
	return pool_of_ReturnOfInt_ch.Get().(chan *ReturnOfInt)
}
func putReturnChOfInt(d chan *ReturnOfInt) {
	pool_of_ReturnOfInt_ch.Put(d)
}

type ReturnOfInt16 struct {
	Context context.Context
	Value   int16
	Error   error
}

var (
	zero_of_ReturnOfInt16       ReturnOfInt16
	zero_of_ReturnOfInt16_value int16
	pool_of_ReturnOfInt16       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfInt16{}
		},
	}
	pool_of_ReturnOfInt16_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfInt16, 1)
		},
	}
)

func getReturnOfInt16() *ReturnOfInt16 {
	return pool_of_ReturnOfInt16.Get().(*ReturnOfInt16)
}
func putReturnOfInt16(d *ReturnOfInt16) {
	d.Context = nil
	d.Value = zero_of_ReturnOfInt16_value
	d.Error = nil
	pool_of_ReturnOfInt16.Put(d)
}

func getReturnChOfInt16() chan *ReturnOfInt16 {
	return pool_of_ReturnOfInt16_ch.Get().(chan *ReturnOfInt16)
}
func putReturnChOfInt16(d chan *ReturnOfInt16) {
	pool_of_ReturnOfInt16_ch.Put(d)
}

type ReturnOfInt32 struct {
	Context context.Context
	Value   int32
	Error   error
}

var (
	zero_of_ReturnOfInt32       ReturnOfInt32
	zero_of_ReturnOfInt32_value int32
	pool_of_ReturnOfInt32       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfInt32{}
		},
	}
	pool_of_ReturnOfInt32_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfInt32, 1)
		},
	}
)

func getReturnOfInt32() *ReturnOfInt32 {
	return pool_of_ReturnOfInt32.Get().(*ReturnOfInt32)
}
func putReturnOfInt32(d *ReturnOfInt32) {
	d.Context = nil
	d.Value = zero_of_ReturnOfInt32_value
	d.Error = nil
	pool_of_ReturnOfInt32.Put(d)
}

func getReturnChOfInt32() chan *ReturnOfInt32 {
	return pool_of_ReturnOfInt32_ch.Get().(chan *ReturnOfInt32)
}
func putReturnChOfInt32(d chan *ReturnOfInt32) {
	pool_of_ReturnOfInt32_ch.Put(d)
}

type ReturnOfInt64 struct {
	Context context.Context
	Value   int64
	Error   error
}

var (
	zero_of_ReturnOfInt64       ReturnOfInt64
	zero_of_ReturnOfInt64_value int64
	pool_of_ReturnOfInt64       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfInt64{}
		},
	}
	pool_of_ReturnOfInt64_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfInt64, 1)
		},
	}
)

func getReturnOfInt64() *ReturnOfInt64 {
	return pool_of_ReturnOfInt64.Get().(*ReturnOfInt64)
}
func putReturnOfInt64(d *ReturnOfInt64) {
	d.Context = nil
	d.Value = zero_of_ReturnOfInt64_value
	d.Error = nil
	pool_of_ReturnOfInt64.Put(d)
}

func getReturnChOfInt64() chan *ReturnOfInt64 {
	return pool_of_ReturnOfInt64_ch.Get().(chan *ReturnOfInt64)
}
func putReturnChOfInt64(d chan *ReturnOfInt64) {
	pool_of_ReturnOfInt64_ch.Put(d)
}

type ReturnOfInt8 struct {
	Context context.Context
	Value   int8
	Error   error
}

var (
	zero_of_ReturnOfInt8       ReturnOfInt8
	zero_of_ReturnOfInt8_value int8
	pool_of_ReturnOfInt8       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfInt8{}
		},
	}
	pool_of_ReturnOfInt8_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfInt8, 1)
		},
	}
)

func getReturnOfInt8() *ReturnOfInt8 {
	return pool_of_ReturnOfInt8.Get().(*ReturnOfInt8)
}
func putReturnOfInt8(d *ReturnOfInt8) {
	d.Context = nil
	d.Value = zero_of_ReturnOfInt8_value
	d.Error = nil
	pool_of_ReturnOfInt8.Put(d)
}

func getReturnChOfInt8() chan *ReturnOfInt8 {
	return pool_of_ReturnOfInt8_ch.Get().(chan *ReturnOfInt8)
}
func putReturnChOfInt8(d chan *ReturnOfInt8) {
	pool_of_ReturnOfInt8_ch.Put(d)
}

type ReturnOfRune struct {
	Context context.Context
	Value   rune
	Error   error
}

var (
	zero_of_ReturnOfRune       ReturnOfRune
	zero_of_ReturnOfRune_value rune
	pool_of_ReturnOfRune       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfRune{}
		},
	}
	pool_of_ReturnOfRune_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfRune, 1)
		},
	}
)

func getReturnOfRune() *ReturnOfRune {
	return pool_of_ReturnOfRune.Get().(*ReturnOfRune)
}
func putReturnOfRune(d *ReturnOfRune) {
	d.Context = nil
	d.Value = zero_of_ReturnOfRune_value
	d.Error = nil
	pool_of_ReturnOfRune.Put(d)
}

func getReturnChOfRune() chan *ReturnOfRune {
	return pool_of_ReturnOfRune_ch.Get().(chan *ReturnOfRune)
}
func putReturnChOfRune(d chan *ReturnOfRune) {
	pool_of_ReturnOfRune_ch.Put(d)
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

type ReturnOfUint struct {
	Context context.Context
	Value   uint
	Error   error
}

var (
	zero_of_ReturnOfUint       ReturnOfUint
	zero_of_ReturnOfUint_value uint
	pool_of_ReturnOfUint       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfUint{}
		},
	}
	pool_of_ReturnOfUint_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfUint, 1)
		},
	}
)

func getReturnOfUint() *ReturnOfUint {
	return pool_of_ReturnOfUint.Get().(*ReturnOfUint)
}
func putReturnOfUint(d *ReturnOfUint) {
	d.Context = nil
	d.Value = zero_of_ReturnOfUint_value
	d.Error = nil
	pool_of_ReturnOfUint.Put(d)
}

func getReturnChOfUint() chan *ReturnOfUint {
	return pool_of_ReturnOfUint_ch.Get().(chan *ReturnOfUint)
}
func putReturnChOfUint(d chan *ReturnOfUint) {
	pool_of_ReturnOfUint_ch.Put(d)
}

type ReturnOfUint16 struct {
	Context context.Context
	Value   uint16
	Error   error
}

var (
	zero_of_ReturnOfUint16       ReturnOfUint16
	zero_of_ReturnOfUint16_value uint16
	pool_of_ReturnOfUint16       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfUint16{}
		},
	}
	pool_of_ReturnOfUint16_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfUint16, 1)
		},
	}
)

func getReturnOfUint16() *ReturnOfUint16 {
	return pool_of_ReturnOfUint16.Get().(*ReturnOfUint16)
}
func putReturnOfUint16(d *ReturnOfUint16) {
	d.Context = nil
	d.Value = zero_of_ReturnOfUint16_value
	d.Error = nil
	pool_of_ReturnOfUint16.Put(d)
}

func getReturnChOfUint16() chan *ReturnOfUint16 {
	return pool_of_ReturnOfUint16_ch.Get().(chan *ReturnOfUint16)
}
func putReturnChOfUint16(d chan *ReturnOfUint16) {
	pool_of_ReturnOfUint16_ch.Put(d)
}

type ReturnOfUint32 struct {
	Context context.Context
	Value   uint32
	Error   error
}

var (
	zero_of_ReturnOfUint32       ReturnOfUint32
	zero_of_ReturnOfUint32_value uint32
	pool_of_ReturnOfUint32       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfUint32{}
		},
	}
	pool_of_ReturnOfUint32_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfUint32, 1)
		},
	}
)

func getReturnOfUint32() *ReturnOfUint32 {
	return pool_of_ReturnOfUint32.Get().(*ReturnOfUint32)
}
func putReturnOfUint32(d *ReturnOfUint32) {
	d.Context = nil
	d.Value = zero_of_ReturnOfUint32_value
	d.Error = nil
	pool_of_ReturnOfUint32.Put(d)
}

func getReturnChOfUint32() chan *ReturnOfUint32 {
	return pool_of_ReturnOfUint32_ch.Get().(chan *ReturnOfUint32)
}
func putReturnChOfUint32(d chan *ReturnOfUint32) {
	pool_of_ReturnOfUint32_ch.Put(d)
}

type ReturnOfUint64 struct {
	Context context.Context
	Value   uint64
	Error   error
}

var (
	zero_of_ReturnOfUint64       ReturnOfUint64
	zero_of_ReturnOfUint64_value uint64
	pool_of_ReturnOfUint64       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfUint64{}
		},
	}
	pool_of_ReturnOfUint64_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfUint64, 1)
		},
	}
)

func getReturnOfUint64() *ReturnOfUint64 {
	return pool_of_ReturnOfUint64.Get().(*ReturnOfUint64)
}
func putReturnOfUint64(d *ReturnOfUint64) {
	d.Context = nil
	d.Value = zero_of_ReturnOfUint64_value
	d.Error = nil
	pool_of_ReturnOfUint64.Put(d)
}

func getReturnChOfUint64() chan *ReturnOfUint64 {
	return pool_of_ReturnOfUint64_ch.Get().(chan *ReturnOfUint64)
}
func putReturnChOfUint64(d chan *ReturnOfUint64) {
	pool_of_ReturnOfUint64_ch.Put(d)
}

type ReturnOfUint8 struct {
	Context context.Context
	Value   uint8
	Error   error
}

var (
	zero_of_ReturnOfUint8       ReturnOfUint8
	zero_of_ReturnOfUint8_value uint8
	pool_of_ReturnOfUint8       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfUint8{}
		},
	}
	pool_of_ReturnOfUint8_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfUint8, 1)
		},
	}
)

func getReturnOfUint8() *ReturnOfUint8 {
	return pool_of_ReturnOfUint8.Get().(*ReturnOfUint8)
}
func putReturnOfUint8(d *ReturnOfUint8) {
	d.Context = nil
	d.Value = zero_of_ReturnOfUint8_value
	d.Error = nil
	pool_of_ReturnOfUint8.Put(d)
}

func getReturnChOfUint8() chan *ReturnOfUint8 {
	return pool_of_ReturnOfUint8_ch.Get().(chan *ReturnOfUint8)
}
func putReturnChOfUint8(d chan *ReturnOfUint8) {
	pool_of_ReturnOfUint8_ch.Put(d)
}

type ReturnOfUintptr struct {
	Context context.Context
	Value   uintptr
	Error   error
}

var (
	zero_of_ReturnOfUintptr       ReturnOfUintptr
	zero_of_ReturnOfUintptr_value uintptr
	pool_of_ReturnOfUintptr       = sync.Pool{
		New: func() interface{} {
			return &ReturnOfUintptr{}
		},
	}
	pool_of_ReturnOfUintptr_ch = sync.Pool{
		New: func() interface{} {
			return make(chan *ReturnOfUintptr, 1)
		},
	}
)

func getReturnOfUintptr() *ReturnOfUintptr {
	return pool_of_ReturnOfUintptr.Get().(*ReturnOfUintptr)
}
func putReturnOfUintptr(d *ReturnOfUintptr) {
	d.Context = nil
	d.Value = zero_of_ReturnOfUintptr_value
	d.Error = nil
	pool_of_ReturnOfUintptr.Put(d)
}

func getReturnChOfUintptr() chan *ReturnOfUintptr {
	return pool_of_ReturnOfUintptr_ch.Get().(chan *ReturnOfUintptr)
}
func putReturnChOfUintptr(d chan *ReturnOfUintptr) {
	pool_of_ReturnOfUintptr_ch.Put(d)
}

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
