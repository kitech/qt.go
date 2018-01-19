package qtrt

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
)

func KeepMe() {}

// 类似C++的overload name resolve
func SymbolResolve(args []interface{}, vtys map[uint8]map[uint8]reflect.Type) int32 {
	// cycle 1
	if ret := symbolResolveComplete(args, vtys); ret != -1 {
		return ret
	}

	// cycle 2
	if ret := symbolResolveConvert(args, vtys); ret != -1 {
		return ret
	}

	// cycle others
	argc := len(args)
	if argc < 0 {
	}

	matsyms := make(map[uint8]bool, 0)
	for symidx := 0; symidx < len(vtys); symidx++ {
		vty := vtys[uint8(symidx)]
		pnum := len(vty)

		// TODO maybe have default values
		if argc != pnum {
			continue
		}

		matp := make(map[int]bool, len(vty))
		for idx := 0; idx < len(vty); idx++ {
			matp[idx] = false

			ety := vty[uint8(idx)]
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

		// a folder/reduce
		ismat := true
		for idx := 0; idx < len(matp); idx++ {
			ismat = ismat && matp[idx]
		}

		// 查找到第一个match的symbol即返回
		if ismat {
			matsyms[uint8(symidx)] = true // 准备多级match
			return int32(symidx)
		}
	}

	return -1
}

func FillDefaultValues(args []interface{}, dvals []interface{}) []interface{} {
	return args
}

// 完全匹配，参数个数，参数类型
func symbolResolveComplete(args []interface{}, vtys map[uint8]map[uint8]reflect.Type) int32 {
	argc := len(args)
	if argc < 0 {
	}

	matsyms := make(map[uint8]bool, 0)
	for symidx := 0; symidx < len(vtys); symidx++ {
		vty := vtys[uint8(symidx)]
		pnum := len(vty)

		if argc != pnum {
			continue
		}

		// 参数类型匹配
		matp := make(map[int]bool, len(vty))
		for idx := 0; idx < len(vty); idx++ {
			matp[idx] = false

			ety := vty[uint8(idx)]
			aty := reflect.TypeOf(args[idx])

			if aty.AssignableTo(ety) {
				matp[idx] = true
			}
		}

		// a folder/reduce
		ismat := true
		for idx := 0; idx < len(matp); idx++ {
			ismat = ismat && matp[idx]
		}

		// 查找到第一个match的symbol即返回
		if ismat {
			matsyms[uint8(symidx)] = true // 准备多级match
			return int32(symidx)
		}
	}

	return -1
}

// 参数个数匹配
func symbolResolveCount() {

}

// 参数类型直接匹配？是否是symbolResolveCompelete？
func symbolResolveType() {
}

// 参数个数匹配
// 参数类型转换匹配
func symbolResolveConvert(args []interface{}, vtys map[uint8]map[uint8]reflect.Type) int32 {
	argc := len(args)
	if argc < 0 {
	}

	matsyms := make(map[int32]bool, 0)
	for symidx := 0; symidx < len(vtys); symidx++ {
		vty := vtys[uint8(symidx)]
		pnum := len(vty)

		if argc != pnum {
			continue
		}

		// 参数类型匹配
		matp := make(map[int]bool, len(vty))
		for idx := 0; idx < len(vty); idx++ {
			matp[idx] = false

			ety := vty[uint8(idx)]
			aty := reflect.TypeOf(args[idx])

			// 关键点？
			if aty.ConvertibleTo(ety) {
				matp[idx] = true
			}
		}

		// a folder/reduce
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

	log.Println(strings.Join(infos, ", "))

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
func getFuncTypes(f interface{}) map[int32]reflect.Type {
	if f == nil {
		return nil
	}

	vf := reflect.TypeOf(f)
	if vf.Kind() != reflect.Func {
		log.Fatalln(vf)
		return nil
	}

	rets := make(map[int32]reflect.Type, 0)
	for idx := 0; idx < vf.NumIn(); idx++ {
		rets[int32(idx)] = vf.In(idx)
	}
	return rets
}
