package qtqml

// /usr/include/qt/QtQml/qqmlscriptstring.h
// #include <qqmlscriptstring.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

type QQmlScriptString struct {
	*qtrt.CObject
}
type QQmlScriptString_ITF interface {
	QQmlScriptString_PTR() *QQmlScriptString
}

func (ptr *QQmlScriptString) QQmlScriptString_PTR() *QQmlScriptString { return ptr }

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
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlScriptString & operator=(const QQmlScriptString &)
func (this *QQmlScriptString) Operator_equal(arg0 QQmlScriptString_ITF) *QQmlScriptString {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlScriptString_PTR() != nil {
		convArg0 = arg0.QQmlScriptString_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QQmlScriptStringaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlScriptStringFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlScriptString)
	return rv2
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:67
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QQmlScriptString &) const
func (this *QQmlScriptString) Operator_equal_equal(arg0 QQmlScriptString_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlScriptString_PTR() != nil {
		convArg0 = arg0.QQmlScriptString_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlScriptStringeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QQmlScriptString &) const
func (this *QQmlScriptString) Operator_not_equal(arg0 QQmlScriptString_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlScriptString_PTR() != nil {
		convArg0 = arg0.QQmlScriptString_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlScriptStringneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const
func (this *QQmlScriptString) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlScriptString7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUndefinedLiteral() const
func (this *QQmlScriptString) IsUndefinedLiteral() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlScriptString18isUndefinedLiteralEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNullLiteral() const
func (this *QQmlScriptString) IsNullLiteral() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlScriptString13isNullLiteralEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QString stringLiteral() const
func (this *QQmlScriptString) StringLiteral() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlScriptString13stringLiteralEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal numberLiteral(_Bool *) const
func (this *QQmlScriptString) NumberLiteral(ok *bool) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlScriptString13numberLiteralEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQml/qqmlscriptstring.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool booleanLiteral(_Bool *) const
func (this *QQmlScriptString) BooleanLiteral(ok *bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlScriptString14booleanLiteralEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
