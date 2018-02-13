package qtqml

// /usr/include/qt/QtQml/qqmlabstracturlinterceptor.h
// #include <qqmlabstracturlinterceptor.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

type QQmlAbstractUrlInterceptor struct {
	*qtrt.CObject
}
type QQmlAbstractUrlInterceptor_ITF interface {
	QQmlAbstractUrlInterceptor_PTR() *QQmlAbstractUrlInterceptor
}

func (ptr *QQmlAbstractUrlInterceptor) QQmlAbstractUrlInterceptor_PTR() *QQmlAbstractUrlInterceptor {
	return ptr
}

func (this *QQmlAbstractUrlInterceptor) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlAbstractUrlInterceptor) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlAbstractUrlInterceptorFromPointer(cthis unsafe.Pointer) *QQmlAbstractUrlInterceptor {
	return &QQmlAbstractUrlInterceptor{&qtrt.CObject{cthis}}
}
func (*QQmlAbstractUrlInterceptor) NewFromPointer(cthis unsafe.Pointer) *QQmlAbstractUrlInterceptor {
	return NewQQmlAbstractUrlInterceptorFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlabstracturlinterceptor.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QQmlAbstractUrlInterceptor()
func NewQQmlAbstractUrlInterceptor() *QQmlAbstractUrlInterceptor {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QQmlAbstractUrlInterceptorC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlAbstractUrlInterceptorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlAbstractUrlInterceptor)
	return gothis
}

// /usr/include/qt/QtQml/qqmlabstracturlinterceptor.h:59
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void ~QQmlAbstractUrlInterceptor()
func DeleteQQmlAbstractUrlInterceptor(this *QQmlAbstractUrlInterceptor) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QQmlAbstractUrlInterceptorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlabstracturlinterceptor.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QUrl intercept(const QUrl &, enum QQmlAbstractUrlInterceptor::DataType)
func (this *QQmlAbstractUrlInterceptor) Intercept(path *qtcore.QUrl, type_ int) *qtcore.QUrl /*123*/ {
	var convArg0 = path.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QQmlAbstractUrlInterceptor9interceptERK4QUrlNS_8DataTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

type QQmlAbstractUrlInterceptor__DataType = int

const QQmlAbstractUrlInterceptor__QmlFile QQmlAbstractUrlInterceptor__DataType = 0
const QQmlAbstractUrlInterceptor__JavaScriptFile QQmlAbstractUrlInterceptor__DataType = 1
const QQmlAbstractUrlInterceptor__QmldirFile QQmlAbstractUrlInterceptor__DataType = 2
const QQmlAbstractUrlInterceptor__UrlString QQmlAbstractUrlInterceptor__DataType = 4096

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
