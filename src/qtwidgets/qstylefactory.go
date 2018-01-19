//  header block begin
// /usr/include/qt/QtWidgets/qstylefactory.h
// #include <qstylefactory.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qstylefactory.h:54
// index:0
// static
// QStringList keys()
func (this *QStyleFactory) Keys() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStyleFactory4keysEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStyleFactory_Keys() {
	// 0: (), ()
	var nilthis *QStyleFactory
	nilthis.Keys()
}

// /usr/include/qt/QtWidgets/qstylefactory.h:55
// index:0
// static
// QStyle * create(const class QString &)
func (this *QStyleFactory) Create(arg0 unsafe.Pointer) {
	// 0: (const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QStyleFactory6createERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStyleFactory_Create(arg0 unsafe.Pointer) {
	// 0: (const QString & arg0), (arg0)
	var nilthis *QStyleFactory
	nilthis.Create(arg0)
}

//  body block end
