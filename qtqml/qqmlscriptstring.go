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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlScriptStringFromPointer(cthis unsafe.Pointer) *QQmlScriptString {
	return &QQmlScriptString{&qtrt.CObject{cthis}}
}
func (*QQmlScriptString) NewFromPointer(cthis unsafe.Pointer) *QQmlScriptString {
	return NewQQmlScriptStringFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlScriptString()
func NewQQmlScriptString() *QQmlScriptString {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QQmlScriptStringC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlScriptStringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlScriptString)
	return gothis
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QQmlScriptString()
func DeleteQQmlScriptString(this *QQmlScriptString) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QQmlScriptStringD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QQmlScriptString) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlScriptString7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUndefinedLiteral()
func (this *QQmlScriptString) IsUndefinedLiteral() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlScriptString18isUndefinedLiteralEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNullLiteral()
func (this *QQmlScriptString) IsNullLiteral() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlScriptString13isNullLiteralEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QString stringLiteral()
func (this *QQmlScriptString) StringLiteral() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlScriptString13stringLiteralEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal numberLiteral(_Bool *)
func (this *QQmlScriptString) NumberLiteral(ok unsafe.Pointer /*666*/) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlScriptString13numberLiteralEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool booleanLiteral(_Bool *)
func (this *QQmlScriptString) BooleanLiteral(ok unsafe.Pointer /*666*/) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlScriptString14booleanLiteralEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

//  body block end
