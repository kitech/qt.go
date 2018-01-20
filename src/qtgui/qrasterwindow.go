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
	*QPaintDeviceWindow
}

func (this *QRasterWindow) GetCthis() unsafe.Pointer {
	return this.QPaintDeviceWindow.GetCthis()
}

// /usr/include/qt/QtGui/qrasterwindow.h:52
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QRasterWindow) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QRasterWindow10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrasterwindow.h:56
// index:0
// void QRasterWindow(class QWindow *)
func NewQRasterWindow(parent unsafe.Pointer) *QRasterWindow {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QRasterWindowC2EP7QWindow", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQRasterWindowFromPointer(cthis)
	return gothis
}
func NewQRasterWindowFromPointer(cthis unsafe.Pointer) *QRasterWindow {
	bcthis0 := NewQPaintDeviceWindowFromPointer(cthis)
	return &QRasterWindow{bcthis0}
}

// /usr/include/qt/QtGui/qrasterwindow.h:57
// index:0
// virtual
// void ~QRasterWindow()
func DeleteQRasterWindow(*QRasterWindow) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QRasterWindowD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrasterwindow.h:60
// index:0
// virtual
// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QRasterWindow) Metric(metric int) {
	// 0: (, metric QPaintDevice::PaintDeviceMetric), (&metric)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QRasterWindow6metricEN12QPaintDevice17PaintDeviceMetricE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &metric)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrasterwindow.h:61
// index:0
// virtual
// QPaintDevice * redirected(class QPoint *)
func (this *QRasterWindow) Redirected(arg0 unsafe.Pointer) {
	// 0: (, QPoint * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QRasterWindow10redirectedEP6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
