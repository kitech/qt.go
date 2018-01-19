//  header block begin
// /usr/include/qt/QtGui/qrasterwindow.h
// #include <qrasterwindow.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QRasterWindow struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qrasterwindow.h:52
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QRasterWindow) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QRasterWindow10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrasterwindow.h:56
// index:0
// void QRasterWindow(class QWindow *)
func NewQRasterWindow(parent unsafe.Pointer) *QRasterWindow {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QRasterWindowC2EP7QWindow", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QRasterWindow{cthis}
}

// /usr/include/qt/QtGui/qrasterwindow.h:57
// index:0
// virtual
// void ~QRasterWindow()
func DeleteQRasterWindow(*QRasterWindow) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QRasterWindowD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
