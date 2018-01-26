package qtqml

// /usr/include/qt/QtQml/qqmlabstracturlinterceptor.h
// #include <qqmlabstracturlinterceptor.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
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
type QQmlAbstractUrlInterceptor struct {
	*qtrt.CObject
}

func (this *QQmlAbstractUrlInterceptor) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlAbstractUrlInterceptor) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQQmlAbstractUrlInterceptorFromPointer(cthis unsafe.Pointer) *QQmlAbstractUrlInterceptor {
	return &QQmlAbstractUrlInterceptor{&qtrt.CObject{cthis}}
}
func (*QQmlAbstractUrlInterceptor) NewFromPointer(cthis unsafe.Pointer) *QQmlAbstractUrlInterceptor {
	return NewQQmlAbstractUrlInterceptorFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlabstracturlinterceptor.h:58
// index:0
// Public inline
// void QQmlAbstractUrlInterceptor()
func NewQQmlAbstractUrlInterceptor() *QQmlAbstractUrlInterceptor {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QQmlAbstractUrlInterceptorC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlAbstractUrlInterceptorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlabstracturlinterceptor.h:59
// index:0
// Public inline virtual
// void ~QQmlAbstractUrlInterceptor()
func DeleteQQmlAbstractUrlInterceptor(*QQmlAbstractUrlInterceptor) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QQmlAbstractUrlInterceptorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlabstracturlinterceptor.h:60
// index:0
// Public pure virtual
// QUrl intercept(const class QUrl &, enum QQmlAbstractUrlInterceptor::DataType)
func (this *QQmlAbstractUrlInterceptor) Intercept(path *qtcore.QUrl, type_ int) *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QQmlAbstractUrlInterceptor9interceptERK4QUrlNS_8DataTypeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, type_)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

type QQmlAbstractUrlInterceptor__DataType = int

const QQmlAbstractUrlInterceptor__QmlFile QQmlAbstractUrlInterceptor__DataType = 0
const QQmlAbstractUrlInterceptor__JavaScriptFile QQmlAbstractUrlInterceptor__DataType = 1
const QQmlAbstractUrlInterceptor__QmldirFile QQmlAbstractUrlInterceptor__DataType = 2
const QQmlAbstractUrlInterceptor__UrlString QQmlAbstractUrlInterceptor__DataType = 4096

//  body block end
