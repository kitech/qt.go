package qtrt

/*
#cgo CFLAGS: -std=c11
*/
import "C"

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
	"unsafe"
)

func KeepMe() {}

// 类似C++的overload name resolve
func SymbolResolve(args []interface{}, vtys map[int32]map[int32]reflect.Type) int32 {
	argc := len(args)
	if argc < 0 {
	}

	matsyms := make(map[int32]bool, 0)
	for symidx := 0; symidx < len(vtys); symidx++ {
		vty := vtys[int32(symidx)]
		pnum := len(vty)

		if argc != pnum {
			continue
		}

		matp := make(map[int]bool, len(vty))
		for idx := 0; idx < len(vty); idx++ {
			matp[idx] = false

			ety := vty[int32(idx)]
			aty := reflect.TypeOf(args[idx])

			if ety.Kind() != reflect.Struct && (aty.ConvertibleTo(ety) || aty.AssignableTo(ety)) {
				matp[idx] = true
			}
			if matp[idx] == false {
				if canHandyConvert(aty, ety) {
					matp[idx] = true
				}
			}

		}

		ismat := true
		for idx := 0; idx < len(matp); idx++ {
			ismat = ismat && matp[idx]
		}

		// 查找到第一个match的symbol即返回
		if ismat {
			matsyms[int32(symidx)] = true // 准备多级match
			return int32(symidx)
		}
	}

	return -1
}

// 参数个数匹配
func symbolResolveCount() {

}

// 参数类型直接匹配
func symbolResolveType() {

}

// 参数类型转换匹配
func symbolResolveConvert() {

}

// 无临时对象的参数类型转换匹配
func symbolResolveNotemp() {

}

func canHandyConvert(from reflect.Type, to reflect.Type) bool {
	infos := make([]string, 0)

	infos = append(infos, from.Kind().String())
	switch from.Kind() {
	case reflect.Ptr:
		infos = append(infos, from.Elem().Kind().String(), from.Elem().Name())
	case reflect.Struct:
		infos = append(infos, from.Name())
	default:
	}

	infos = append(infos, "===")

	infos = append(infos, to.Kind().String())
	switch to.Kind() {
	case reflect.Ptr:
		infos = append(infos, to.Elem().Kind().String(), to.Elem().Name())
	case reflect.Struct:
		infos = append(infos, to.Name())
	default:
	}

	fmt.Println(strings.Join(infos, ", "))

	//
	switch {
	// string => char *
	case from.Kind() == reflect.String &&
		to.Kind() == reflect.Ptr && to.Elem().Kind() == reflect.Uint8:
		return true
		// &QXxx => QXxx
	case from.Kind() == reflect.Ptr &&
		from.Elem().Kind() == to.Kind() && from.Elem().Name() == to.Name():
		return true
		// QXxx => QXxx
	case from.Kind() == reflect.Struct &&
		from.Kind() == to.Kind() && from.Name() == to.Name():
		return true
	}
	return false
}

func ErrorResolve(class, method string, args []interface{}) {
	pcs := make([]uintptr, 3)
	npc := runtime.Callers(0, pcs)
	if npc == -1 {
	}
	rtf := runtime.FuncForPC(pcs[2])
	file, line := rtf.FileLine(pcs[2])
	fmt.Println(rtf.Name(), file, line, "Unresolved VT", class, method, args)
}

/////////
// runtime.SetFinalizer(x, UniverseFree)
func UniverseFree(this interface{}) {
	oty := reflect.TypeOf(this)
	oval := reflect.ValueOf(this)

	_, ok := oty.MethodByName("Free")
	if ok {
		// in := []reflect.Value{oval}
		// mth.Func.Call(in)
		oval.MethodByName("Free").Call([]reflect.Value{})
		// log.Println(this, "freed", oty.Elem().Name())
	} else {
		log.Println(this, "has no Free method.", oty.Elem().Name())
	}
}

/////////

func BoolTy(pointer bool) reflect.Type {
	if pointer {
		var v = true
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(true)
}

func ByteTy(pointer bool) reflect.Type {
	if pointer {
		var v = byte(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(byte(0))
}

func StringTy(pointer bool) reflect.Type {
	var s = "foo"
	if pointer {
		return reflect.TypeOf(&s)
	}
	return reflect.TypeOf(s)
}

func RuneTy(pointer bool) reflect.Type {
	var s rune = '\000'
	if pointer {
		return reflect.TypeOf(&s)
	}
	return reflect.TypeOf(s)
}

func Int16Ty(pointer bool) reflect.Type {
	if pointer {
		var v = int16(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(int16(0))
}

func UInt16Ty(pointer bool) reflect.Type {
	if pointer {
		var v = uint16(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(uint16(0))
}

func Int32Ty(pointer bool) reflect.Type {
	if pointer {
		var v = int32(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(int32(0))
}

func UInt32Ty(pointer bool) reflect.Type {
	if pointer {
		var v = uint32(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(uint32(0))
}

func Int64Ty(pointer bool) reflect.Type {
	if pointer {
		var v = int64(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(int64(0))
}

func UInt64Ty(pointer bool) reflect.Type {
	if pointer {
		var v = uint64(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(uint64(0))
}

func FloatTy(pointer bool) reflect.Type {
	if pointer {
		var v = float32(0.0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(float32(0.0))
}
func DoubleTy(pointer bool) reflect.Type {
	if pointer {
		var v = float64(0.0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(float64(0.0))
}

func VoidpTy() reflect.Type {
	var v unsafe.Pointer = nil
	return reflect.TypeOf(v)
}
