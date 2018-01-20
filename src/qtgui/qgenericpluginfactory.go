//  header block begin
// /usr/include/qt/QtGui/qgenericpluginfactory.h
// #include <qgenericpluginfactory.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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
	*qtrt.CObject
}

func (this *QGenericPluginFactory) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQGenericPluginFactoryFromPointer(cthis unsafe.Pointer) *QGenericPluginFactory {
	return &QGenericPluginFactory{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qgenericpluginfactory.h:55
// index:0
// Public static
// QStringList keys()
func (this *QGenericPluginFactory) Keys() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGenericPluginFactory4keysEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGenericPluginFactory_Keys() {
	var nilthis *QGenericPluginFactory
	nilthis.Keys()
}

// /usr/include/qt/QtGui/qgenericpluginfactory.h:56
// index:0
// Public static
// QObject * create(const class QString &, const class QString &)
func (this *QGenericPluginFactory) Create(arg0 *qtcore.QString, arg1 *qtcore.QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGenericPluginFactory6createERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, arg0, arg1)
	gopp.ErrPrint(err, rv)
	return rv
}
func QGenericPluginFactory_Create(arg0 *qtcore.QString, arg1 *qtcore.QString) {
	var nilthis *QGenericPluginFactory
	nilthis.Create(arg0, arg1)
}

//  body block end
