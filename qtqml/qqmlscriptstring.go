package qtqml

// /usr/include/qt/QtQml/qqmlscriptstring.h
// #include <qqmlscriptstring.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin
type QQmlScriptString struct {
	*qtrt.CObject
}

func (this *QQmlScriptString) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlScriptString) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQQmlScriptStringFromPointer(cthis unsafe.Pointer) *QQmlScriptString {
	return &QQmlScriptString{&qtrt.CObject{cthis}}
}
func (*QQmlScriptString) NewFromPointer(cthis unsafe.Pointer) *QQmlScriptString {
	return NewQQmlScriptStringFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:61
// index:0
// Public
// void QQmlScriptString()
func NewQQmlScriptString() *QQmlScriptString {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QQmlScriptStringC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlScriptStringFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:63
// index:0
// Public
// void ~QQmlScriptString()
func DeleteQQmlScriptString(*QQmlScriptString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QQmlScriptStringD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:70
// index:0
// Public
// bool isEmpty()
func (this *QQmlScriptString) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QQmlScriptString7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:72
// index:0
// Public
// bool isUndefinedLiteral()
func (this *QQmlScriptString) IsUndefinedLiteral() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QQmlScriptString18isUndefinedLiteralEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:73
// index:0
// Public
// bool isNullLiteral()
func (this *QQmlScriptString) IsNullLiteral() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QQmlScriptString13isNullLiteralEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:74
// index:0
// Public
// QString stringLiteral()
func (this *QQmlScriptString) StringLiteral() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QQmlScriptString13stringLiteralEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:75
// index:0
// Public
// qreal numberLiteral(bool *)
func (this *QQmlScriptString) NumberLiteral(ok unsafe.Pointer /*666*/) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QQmlScriptString13numberLiteralEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:76
// index:0
// Public
// bool booleanLiteral(bool *)
func (this *QQmlScriptString) BooleanLiteral(ok unsafe.Pointer /*666*/) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QQmlScriptString14booleanLiteralEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
