//  header block begin
// /usr/include/qt/QtGui/qpaintdevicewindow.h
// #include <qpaintdevicewindow.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
	*QWindow
	*QPaintDevice
}

func (this *QPaintDeviceWindow) GetCthis() unsafe.Pointer {
	return this.QWindow.GetCthis()
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QPaintDeviceWindow) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QPaintDeviceWindow10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:58
// index:0
// void update(const class QRect &)
func (this *QPaintDeviceWindow) Update(rect unsafe.Pointer) {
	// 0: (, rect const QRect &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindow6updateERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:59
// index:1
// void update(const class QRegion &)
func (this *QPaintDeviceWindow) Update_1(region unsafe.Pointer) {
	// 1: (, region const QRegion &), (region)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindow6updateERK7QRegion", ffiqt.FFI_TYPE_VOID, this.GetCthis(), region)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:66
// index:2
// void update()
func (this *QPaintDeviceWindow) Update_2() {
	// 2: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindow6updateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:69
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QPaintDeviceWindow) PaintEvent(event unsafe.Pointer) {
	// 0: (, event QPaintEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindow10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:71
// index:0
// virtual
// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPaintDeviceWindow) Metric(metric int) {
	// 0: (, metric QPaintDevice::PaintDeviceMetric), (&metric)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QPaintDeviceWindow6metricEN12QPaintDevice17PaintDeviceMetricE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &metric)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:72
// index:0
// virtual
// void exposeEvent(class QExposeEvent *)
func (this *QPaintDeviceWindow) ExposeEvent(arg0 unsafe.Pointer) {
	// 0: (, QExposeEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindow11exposeEventEP12QExposeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:73
// index:0
// virtual
// bool event(class QEvent *)
func (this *QPaintDeviceWindow) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindow5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:75
// index:0
// void QPaintDeviceWindow(class QPaintDeviceWindowPrivate &, class QWindow *)
func NewQPaintDeviceWindow(dd unsafe.Pointer, parent unsafe.Pointer) *QPaintDeviceWindow {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindowC2ER25QPaintDeviceWindowPrivateP7QWindow", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPaintDeviceWindowFromPointer(cthis)
	return gothis
}
func NewQPaintDeviceWindowFromPointer(cthis unsafe.Pointer) *QPaintDeviceWindow {
	bcthis0 := NewQWindowFromPointer(cthis)
	bcthis1 := NewQPaintDeviceFromPointer(cthis)
	return &QPaintDeviceWindow{bcthis0, bcthis1}
}

//  body block end
