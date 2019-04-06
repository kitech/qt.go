//  header block begin

// +build !minimal

package qtqml

// /usr/include/qt/QtQml/qqmlnetworkaccessmanagerfactory.h
// #include <qqmlnetworkaccessmanagerfactory.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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

// /usr/include/qt/QtQml/qqmlnetworkaccessmanagerfactory.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlNetworkAccessManagerFactory()

/*

 */
func DeleteQQmlNetworkAccessManagerFactory(this *QQmlNetworkAccessManagerFactory) {
	rv, err := qtrt.InvokeQtFunc6("_ZN31QQmlNetworkAccessManagerFactoryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlnetworkaccessmanagerfactory.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QNetworkAccessManager * create(QObject *)

/*
Creates and returns a network access manager with the specified parent. This method must return a new QNetworkAccessManager instance each time it is called.

Note: this method may be called by multiple threads, so ensure the implementation of this method is reentrant.
*/
func (this *QQmlNetworkAccessManagerFactory) Create(parent qtcore.QObject_ITF /*777 QObject **/) *qtnetwork.QNetworkAccessManager /*777 QNetworkAccessManager **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN31QQmlNetworkAccessManagerFactory6createEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtnetwork.NewQNetworkAccessManagerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_11540() {
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
