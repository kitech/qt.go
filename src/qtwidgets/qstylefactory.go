//  header block begin
// /usr/include/qt/QtWidgets/qstylefactory.h
// #include <qstylefactory.h>
// #include <QtWidgets>
package qtwidgets

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QStyleFactory struct {
	*qtrt.CObject
}

func (this *QStyleFactory) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQStyleFactoryFromPointer(cthis unsafe.Pointer) *QStyleFactory {
	return &QStyleFactory{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qstylefactory.h:54
// index:0
// Public static
// QStringList keys()
func (this *QStyleFactory) Keys() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStyleFactory4keysEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QStyleFactory_Keys() {
	var nilthis *QStyleFactory
	nilthis.Keys()
}

// /usr/include/qt/QtWidgets/qstylefactory.h:55
// index:0
// Public static
// QStyle * create(const class QString &)
func (this *QStyleFactory) Create(arg0 *qtcore.QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStyleFactory6createERK7QString", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	return rv
}
func QStyleFactory_Create(arg0 *qtcore.QString) {
	var nilthis *QStyleFactory
	nilthis.Create(arg0)
}

//  body block end
