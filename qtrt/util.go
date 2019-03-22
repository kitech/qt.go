package qtrt

import (
	"reflect"

	"github.com/kitech/qt.go/miscutil"
)

var ErrPrint = miscutil.ErrPrint
var NilPrint = miscutil.NilPrint
var FileExist = miscutil.FileExist

var IfElse = miscutil.IfElse
var IfElseInt = miscutil.IfElseInt
var IfElseStr = miscutil.IfElseStr
var Assert = miscutil.Assert

var Uint8Ty = reflect.TypeOf(uint8(0))
var IntTy = reflect.TypeOf(int(0))
var Uint64Ty = reflect.TypeOf(uint64(0))
var Float64Ty = reflect.TypeOf(float64(0))
var Float32Ty = reflect.TypeOf(float32(0))
