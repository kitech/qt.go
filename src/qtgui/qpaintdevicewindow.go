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
func NewQPaintDeviceWindowFromPointer(cthis unsafe.Pointer) *QPaintDeviceWindow {
	bcthis0 := NewQWindowFromPointer(cthis)
	bcthis1 := NewQPaintDeviceFromPointer(cthis)
	return &QPaintDeviceWindow{bcthis0, bcthis1}
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QPaintDeviceWindow) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QPaintDeviceWindow10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:58
// index:0
// Public
// void update(const class QRect &)
func (this *QPaintDeviceWindow) Update(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindow6updateERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:59
// index:1
// Public
// void update(const class QRegion &)
func (this *QPaintDeviceWindow) Update_1(region *QRegion) {
	var convArg0 = region.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindow6updateERK7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:66
// index:2
// Public
// void update()
func (this *QPaintDeviceWindow) Update_2() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindow6updateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:69
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QPaintDeviceWindow) PaintEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindow10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:71
// index:0
// Protected virtual
// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPaintDeviceWindow) Metric(metric int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QPaintDeviceWindow6metricEN12QPaintDevice17PaintDeviceMetricE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &metric)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:72
// index:0
// Protected virtual
// void exposeEvent(class QExposeEvent *)
func (this *QPaintDeviceWindow) ExposeEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindow11exposeEventEP12QExposeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:73
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QPaintDeviceWindow) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPaintDeviceWindow5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
