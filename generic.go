package async

//go:generate genny -in restype__template.go -out restype__template__gen.go gen "_Prefix_=Of Other=Bytes,string,interface{}"
//go:generate genny -in worktype__template.go -out worktype__template__gen.go gen "_Prefix_=Of Some=Bytes,string,interface{} Other=Bytes,string,interface{}"
//go:generate genny -in worker__template.go -out worker__template__gen.go gen "_Prefix_=Of Some=Bytes,string,interface{} Other=Bytes,string,interface{}"

import "github.com/cheekybits/genny/generic"

type Some generic.Type
type Other generic.Type
type Another generic.Type

type Bytes = []byte
