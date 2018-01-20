//  header block begin
// /usr/include/qt/QtGui/qpaintdevice.h
// #include <qpaintdevice.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 54
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
type QPaintDevice struct {
	*qtrt.CObject
}

func (this *QPaintDevice) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qpaintdevice.h:72
// index:0
// virtual
// void ~QPaintDevice()
func DeleteQPaintDevice(*QPaintDevice) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintDeviceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:74
// index:0
// virtual
// int devType()
func (this *QPaintDevice) DevType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice7devTypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:75
// index:0
// bool paintingActive()
func (this *QPaintDevice) PaintingActive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice14paintingActiveEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:76
// index:0
// pure virtual
// QPaintEngine * paintEngine()
func (this *QPaintDevice) PaintEngine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice11paintEngineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:78
// index:0
// inline
// int width()
func (this *QPaintDevice) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice5widthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:79
// index:0
// inline
// int height()
func (this *QPaintDevice) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice6heightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:80
// index:0
// inline
// int widthMM()
func (this *QPaintDevice) WidthMM() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice7widthMMEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:81
// index:0
// inline
// int heightMM()
func (this *QPaintDevice) HeightMM() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice8heightMMEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:82
// index:0
// inline
// int logicalDpiX()
func (this *QPaintDevice) LogicalDpiX() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice11logicalDpiXEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:83
// index:0
// inline
// int logicalDpiY()
func (this *QPaintDevice) LogicalDpiY() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice11logicalDpiYEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:84
// index:0
// inline
// int physicalDpiX()
func (this *QPaintDevice) PhysicalDpiX() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice12physicalDpiXEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:85
// index:0
// inline
// int physicalDpiY()
func (this *QPaintDevice) PhysicalDpiY() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice12physicalDpiYEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:86
// index:0
// inline
// int devicePixelRatio()
func (this *QPaintDevice) DevicePixelRatio() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice16devicePixelRatioEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:87
// index:0
// inline
// qreal devicePixelRatioF()
func (this *QPaintDevice) DevicePixelRatioF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice17devicePixelRatioFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:88
// index:0
// inline
// int colorCount()
func (this *QPaintDevice) ColorCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice10colorCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:89
// index:0
// inline
// int depth()
func (this *QPaintDevice) Depth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice5depthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:91
// index:0
// static inline
// qreal devicePixelRatioFScale()
func (this *QPaintDevice) DevicePixelRatioFScale() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintDevice22devicePixelRatioFScaleEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPaintDevice_DevicePixelRatioFScale() {
	// 0: (), ()
	var nilthis *QPaintDevice
	nilthis.DevicePixelRatioFScale()
}

// /usr/include/qt/QtGui/qpaintdevice.h:93
// index:0
// void QPaintDevice()
func NewQPaintDevice() *QPaintDevice {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintDeviceC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPaintDeviceFromPointer(cthis)
	return gothis
}
func NewQPaintDeviceFromPointer(cthis unsafe.Pointer) *QPaintDevice {
	return &QPaintDevice{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qpaintdevice.h:101
// index:1
// inline
// void QPaintDevice(const class QPaintDevice &)
func NewQPaintDevice_1(arg0 unsafe.Pointer) *QPaintDevice {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPaintDeviceC1ERKS_", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPaintDeviceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpaintdevice.h:94
// index:0
// virtual
// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPaintDevice) Metric(metric int) {
	// 0: (, metric QPaintDevice::PaintDeviceMetric), (&metric)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice6metricENS_17PaintDeviceMetricE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &metric)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:95
// index:0
// virtual
// void initPainter(class QPainter *)
func (this *QPaintDevice) InitPainter(painter unsafe.Pointer) {
	// 0: (, painter QPainter *), (painter)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice11initPainterEP8QPainter", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:96
// index:0
// virtual
// QPaintDevice * redirected(class QPoint *)
func (this *QPaintDevice) Redirected(offset unsafe.Pointer) {
	// 0: (, offset QPoint *), (offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice10redirectedEP6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:97
// index:0
// virtual
// QPainter * sharedPainter()
func (this *QPaintDevice) SharedPainter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPaintDevice13sharedPainterEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
