//  header block begin
// /usr/include/qt/QtGui/qgenericpluginfactory.h
// #include <qgenericpluginfactory.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QGenericPluginFactory struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qgenericpluginfactory.h:55
// index:0
// static
// QStringList keys()
func (this *QGenericPluginFactory) Keys() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGenericPluginFactory4keysEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGenericPluginFactory_Keys() {
	// 0: (), ()
	var nilthis *QGenericPluginFactory
	nilthis.Keys()
}

// /usr/include/qt/QtGui/qgenericpluginfactory.h:56
// index:0
// static
// QObject * create(const class QString &, const class QString &)
func (this *QGenericPluginFactory) Create(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (const QString & arg0, const QString & arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGenericPluginFactory6createERK7QStringS2_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QGenericPluginFactory_Create(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (const QString & arg0, const QString & arg1), (arg0, arg1)
	var nilthis *QGenericPluginFactory
	nilthis.Create(arg0, arg1)
}

//  body block end
