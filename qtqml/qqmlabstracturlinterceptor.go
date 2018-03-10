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

/*

 */
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

/*
Constructor for QQmlAbstractUrlInterceptor.
*/
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

/*

 */
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

/*
A pure virtual function where you can intercept the url. The returned value is taken as the new value for the url. The type of url being intercepted is given by the type variable.

Your implementation of this function must be thread-safe, as it can be called from multiple threads at the same time.
*/
func (this *QQmlAbstractUrlInterceptor) Intercept(path qtcore.QUrl_ITF, type_ int) *qtcore.QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if path != nil && path.QUrl_PTR() != nil {
		convArg0 = path.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QQmlAbstractUrlInterceptor9interceptERK4QUrlNS_8DataTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

/*
Specifies where URL interception is taking place.

Because QML loads qmldir files for locating types, there are two URLs involved in loading a QML type. The URL of the (possibly implicit) qmldir used for locating the type and the URL of the file which defines the type. Intercepting both leads to either complex URL replacement or double URL replacements for the same file.


*/
type QQmlAbstractUrlInterceptor__DataType = int

// The URL being intercepted is for a Qml file. Intercepting this, but not the Qmldir file, leaves the base dir of a QML file untouched and acts like replacing the file with another file.
const QQmlAbstractUrlInterceptor__QmlFile QQmlAbstractUrlInterceptor__DataType = 0

// The URL being intercepted is an import for a Javascript file.
const QQmlAbstractUrlInterceptor__JavaScriptFile QQmlAbstractUrlInterceptor__DataType = 1

// The URL being intercepted is for a Qmldir file. Intercepting this, but not the QmlFile, allows for swapping out entire sub trees.
const QQmlAbstractUrlInterceptor__QmldirFile QQmlAbstractUrlInterceptor__DataType = 2

//
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
