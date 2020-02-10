package async

//go:generate genny -in _restype__template.go -out restype__template__gen.go gen "_Prefix_=Of Other=BUILTINS,interface{},struct{}"
//go:generate genny -in _reqtype__template.go -out reqtype__template__gen.go gen "_Prefix_=Of Some=BUILTINS,interface{},struct{} Other=BUILTINS,interface{},struct{}"

import "github.com/cheekybits/genny/generic"

type Some generic.Type
type Other generic.Type
