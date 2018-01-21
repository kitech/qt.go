package qtgui

// /usr/include/qt/QtGui/qgenericpluginfactory.h
// #include <qgenericpluginfactory.h>
// #include <QtGui>

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQGenericPluginFactoryFromPointer(cthis unsafe.Pointer) *QGenericPluginFactory {
	return &QGenericPluginFactory{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qgenericpluginfactory.h:56
// index:0
// Public static
// QObject * create(const class QString &, const class QString &)
func (this *QGenericPluginFactory) Create(arg0 *qtcore.QString, arg1 *qtcore.QString) *qtcore.QObject /*444 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGenericPluginFactory6createERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, arg0, arg1)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QGenericPluginFactory_Create(arg0 *qtcore.QString, arg1 *qtcore.QString) *qtcore.QObject /*444 QObject **/ {
	var nilthis *QGenericPluginFactory
	rv := nilthis.Create(arg0, arg1)
	return rv
}

//  body block end
