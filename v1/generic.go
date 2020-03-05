package async

//go:generate genny -in return__template.go -out return__template__gen.go gen "Some=BUILTINS,Bytes,interface{}"
//go:generate genny -in types__template.go -out types__template__gen.go gen "Some=BUILTINS,Bytes,interface{} Other=BUILTINS,Bytes,interface{}"
// async package is a handy functions for handling async logic using channel.
// mainly ties data types (context, data, error...) to send/receive it through single channel

import "github.com/cheekybits/genny/generic"

type Some generic.Type
type Other generic.Type

type Bytes = []byte
