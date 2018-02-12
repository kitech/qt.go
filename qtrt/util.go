package qtrt

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"syscall"
)

func printq(v interface{}, args ...interface{}) string {
	msg := fmt.Sprintf("%+v", v)
	for _, arg := range args {
		msg += fmt.Sprintf(" %+v", arg)
	}
	return msg
}

func ErrPrint(err error, args ...interface{}) error {
	if err != nil {
		log.Output(2, printq(err, args...))
	}
	return err
}

func NilPrint(v interface{}, args ...interface{}) interface{} {
	if v == nil {
		log.Output(2, printq(v, args...))
	}
	return v
}

func FileExist(fname string) bool {
	if _, err := os.Stat(fname); err != nil {
		if err.(*os.PathError).Err == syscall.ENOENT {
			return false
		}
	}
	return true
}

// TODO 要是侯选可以惰性求值就好了，否则在只能一个求值的场景则会有问题
// 简单的三元去处模拟函数
func IfElse(q bool, tv interface{}, fv interface{}) interface{} {
	if q == true {
		return tv
	} else {
		return fv
	}
}

func IfElseInt(q bool, tv int, fv int) int {
	return IfElse(q, tv, fv).(int)
}

func IfElseStr(q bool, tv string, fv string) string {
	return IfElse(q, tv, fv).(string)
}

func Assert(v interface{}, info string, args ...interface{}) {
	fmtv := fmt.Sprintf("%+v, %+v", v, info)
	for _, arg := range args {
		fmtv += fmt.Sprintf(", %+v", arg)
	}
	if v == nil {
		panic(fmtv)
	}

	tv := reflect.TypeOf(v)
	if tv.Kind() == reflect.Bool && v.(bool) == false {
		panic(fmtv)
	}

	vv := reflect.ValueOf(v)
	switch tv.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uint8, reflect.Int8:
		if vv.Int() == 0 {
			panic(fmtv)
		}
	case reflect.String:
		if v.(string) == "" {
			panic(fmtv)
		}
	}
}

var Uint8Ty = reflect.TypeOf(uint8(0))
var IntTy = reflect.TypeOf(int(0))
