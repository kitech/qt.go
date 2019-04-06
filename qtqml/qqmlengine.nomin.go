//  header block begin

// +build !minimal

package qtqml

// /usr/include/qt/QtQml/qqmlengine.h
// #include <qqmlengine.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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

// /usr/include/qt/QtQml/qqmlengine.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNetworkAccessManagerFactory(QQmlNetworkAccessManagerFactory *)

/*
Sets the factory to use for creating QNetworkAccessManager(s).

QNetworkAccessManager is used for all network access by QML. By implementing a factory it is possible to create custom QNetworkAccessManager with specialized caching, proxy and cookie support.

The factory must be set before executing the engine.

Note: QQmlEngine does not take ownership of the factory.

See also networkAccessManagerFactory().
*/
func (this *QQmlEngine) SetNetworkAccessManagerFactory(arg0 QQmlNetworkAccessManagerFactory_ITF /*777 QQmlNetworkAccessManagerFactory **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlNetworkAccessManagerFactory_PTR() != nil {
		convArg0 = arg0.QQmlNetworkAccessManagerFactory_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine30setNetworkAccessManagerFactoryEP31QQmlNetworkAccessManagerFactory", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlNetworkAccessManagerFactory * networkAccessManagerFactory() const

/*
Returns the current QQmlNetworkAccessManagerFactory.

See also setNetworkAccessManagerFactory().
*/
func (this *QQmlEngine) NetworkAccessManagerFactory() *QQmlNetworkAccessManagerFactory /*777 QQmlNetworkAccessManagerFactory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine27networkAccessManagerFactoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlNetworkAccessManagerFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlengine.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkAccessManager * networkAccessManager() const

/*
Returns a common QNetworkAccessManager which can be used by any QML type instantiated by this engine.

If a QQmlNetworkAccessManagerFactory has been set and a QNetworkAccessManager has not yet been created, the QQmlNetworkAccessManagerFactory will be used to create the QNetworkAccessManager; otherwise the returned QNetworkAccessManager will have no proxy or cache set.

See also setNetworkAccessManagerFactory().
*/
func (this *QQmlEngine) NetworkAccessManager() *qtnetwork.QNetworkAccessManager /*777 QNetworkAccessManager **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine20networkAccessManagerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtnetwork.NewQNetworkAccessManagerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_11512() {
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
