package async

//go:generate genny -in other__return__template.go -out other__return__template__gen.go gen "Other=Bytes,string,interface{}"
//go:generate genny -in some__types__template.go -out some__types__template__gen.go gen "Some=Bytes,string,interface{} Other=Bytes,string,interface{}"
//go:generate genny -in some_other__impl__template.go -out some_other__impl__template__gen.go gen "Some=Bytes,string,interface{} Other=Bytes,string,interface{}"
// async package is a handy functions for handling async logic using channel.
// mainly ties data types (context, data, error...) to send/receive it through single channel

import "github.com/cheekybits/genny/generic"

type Some generic.Type
type Other generic.Type

type Bytes = []byte
