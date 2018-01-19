//  header block begin
// /usr/include/qt/QtGui/qpaintdevicewindow.h
// #include <qpaintdevicewindow.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QPaintDeviceWindow struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QPaintDeviceWindow) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QPaintDeviceWindow10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:58
// index:0
// void update(const class QRect &)
func (this *QPaintDeviceWindow) Update(rect unsafe.Pointer) {
	// 0: (, const QRect & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindow6updateERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:59
// index:1
// void update(const class QRegion &)
func (this *QPaintDeviceWindow) Update_1(region unsafe.Pointer) {
	// 1: (, const QRegion & region), (region)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindow6updateERK7QRegion", ffiqt.FFI_TYPE_VOID, this.cthis, region)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:66
// index:2
// void update()
func (this *QPaintDeviceWindow) Update_2() {
	// 2: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindow6updateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
