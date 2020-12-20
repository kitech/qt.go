package qtrt

// some extra meta info funcs

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// must call in func (this *XObject) XEnumItemName()
// 1. QObject subclass
// 2. non-QObject class
func GetClassEnumItemName(this interface{}, val int) string {
	symname := "C_QMetaObject_getEnumItemName"

	defval := fmt.Sprintf("%d", val)
	// non-QObject
	if !reflect.ValueOf(this).MethodByName("QObject_PTR").IsValid() {
		return defval
	}

	// get class staticMetaObject
	smo := getClassStaticMetaObjectByObject(this)
	if smo == nil {
		return defval
	}

	// get enum name
	pc, _, _, _ := runtime.Caller(1)
	fno := runtime.FuncForPC(pc)
	// fno.Name() format: github.com/kitech/qt.go/qtcore.(*QProcess).ExitStatusItemName
	revpos := strings.LastIndex(fno.Name(), ".")
	optname := fno.Name()[revpos+1 : len(fno.Name())-len("ItemName")]
	optname_c := CString(optname)
	defer CFree(optname_c)
	// log.Println(optname, val)

	rv, err := InvokeQtFunc6(symname, FFI_TYPE_POINTER, smo, optname_c, _Ctype_int(val))
	ErrPrint(err, rv, val, optname, smo, this)
	realval := GoStringI(rv)
	if realval == "" {
		return defval
	}
	return realval
}

func getClassStaticMetaObjectByObject(this interface{}) Voidptr {
	// eg. _ZN11QColumnView16staticMetaObjectE
	clsname := getClassNameByObject(this)
	return GetClassStaticMetaObjectByName(clsname)
}

func GetClassStaticMetaObjectByName(clsname string) Voidptr {
	symname := fmt.Sprintf("_ZN%d%s16staticMetaObjectE", len(clsname), clsname)
	addr := GetQtSymAddrRaw(symname)
	return addr
}

func getClassNameByObject(this interface{}) string {
	oty := reflect.TypeOf(this)
	clsname := strings.Split(oty.Elem().String(), ".")[1]
	return clsname
}

// must a QObject or subclass
func getClassNameByCObject(cthis Voidptr) string {
	rv, err := InvokeQtFunc6("_ZNK7QObject10metaObjectEv", FFI_TYPE_POINTER, cthis)
	rv2, err2 := InvokeQtFunc6("_ZNK11QMetaObject9classNameEv", FFI_TYPE_POINTER, Voidptr(uintptr(rv)))
	ErrPrint(err, cthis)
	ErrPrint(err2, cthis)
	return GoStringI(rv2)
}
