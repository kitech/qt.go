package qtgui

// /usr/include/qt/QtGui/qpaintdevicewindow.h
// #include <qpaintdevicewindow.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// void paintEvent(class QPaintEvent *)
func (this *QPaintDeviceWindow) InheritPaintEvent(f func(event *QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPaintDeviceWindow) InheritMetric(f func(metric int) int) {
	qtrt.SetAllInheritCallback(this, "metric", f)
}

// void exposeEvent(class QExposeEvent *)
func (this *QPaintDeviceWindow) InheritExposeEvent(f func(arg0 *QExposeEvent /*777 QExposeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "exposeEvent", f)
}

// bool event(class QEvent *)
func (this *QPaintDeviceWindow) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QPaintDeviceWindow struct {
	*QWindow
	*QPaintDevice
}
type QPaintDeviceWindow_ITF interface {
	QWindow_ITF
	QPaintDevice_ITF
	QPaintDeviceWindow_PTR() *QPaintDeviceWindow
}

func (ptr *QPaintDeviceWindow) QPaintDeviceWindow_PTR() *QPaintDeviceWindow { return ptr }

func (this *QPaintDeviceWindow) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWindow.GetCthis()
	}
}
func (this *QPaintDeviceWindow) SetCthis(cthis unsafe.Pointer) {
	this.QWindow = NewQWindowFromPointer(cthis)
	this.QPaintDevice = NewQPaintDeviceFromPointer(cthis)
}
func NewQPaintDeviceWindowFromPointer(cthis unsafe.Pointer) *QPaintDeviceWindow {
	bcthis0 := NewQWindowFromPointer(cthis)
	bcthis1 := NewQPaintDeviceFromPointer(cthis)
	return &QPaintDeviceWindow{bcthis0, bcthis1}
}
func (*QPaintDeviceWindow) NewFromPointer(cthis unsafe.Pointer) *QPaintDeviceWindow {
	return NewQPaintDeviceWindowFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QPaintDeviceWindow) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QPaintDeviceWindow10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update(const QRect &)
func (this *QPaintDeviceWindow) Update(rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPaintDeviceWindow6updateERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void update(const QRegion &)
func (this *QPaintDeviceWindow) Update_1(region QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if region != nil && region.QRegion_PTR() != nil {
		convArg0 = region.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPaintDeviceWindow6updateERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:66
// index:2
// Public Visibility=Default Availability=Available
// [-2] void update()
func (this *QPaintDeviceWindow) Update_2() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPaintDeviceWindow6updateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:69
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QPaintDeviceWindow) PaintEvent(event QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QPaintEvent_PTR() != nil {
		convArg0 = event.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPaintDeviceWindow10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:71
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(enum QPaintDevice::PaintDeviceMetric) const
func (this *QPaintDeviceWindow) Metric(metric int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QPaintDeviceWindow6metricEN12QPaintDevice17PaintDeviceMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), metric)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:72
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void exposeEvent(QExposeEvent *)
func (this *QPaintDeviceWindow) ExposeEvent(arg0 QExposeEvent_ITF /*777 QExposeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QExposeEvent_PTR() != nil {
		convArg0 = arg0.QExposeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPaintDeviceWindow11exposeEventEP12QExposeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevicewindow.h:73
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QPaintDeviceWindow) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPaintDeviceWindow5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQPaintDeviceWindow(this *QPaintDeviceWindow) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPaintDeviceWindowD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
