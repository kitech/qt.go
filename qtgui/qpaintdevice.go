package qtgui

// /usr/include/qt/QtGui/qpaintdevice.h
// #include <qpaintdevice.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 54
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPaintDevice) InheritMetric(f func(metric int) int) {
	qtrt.SetAllInheritCallback(this, "metric", f)
}

// void initPainter(class QPainter *)
func (this *QPaintDevice) InheritInitPainter(f func(painter *QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initPainter", f)
}

// QPaintDevice * redirected(class QPoint *)
func (this *QPaintDevice) InheritRedirected(f func(offset *qtcore.QPoint /*777 QPoint **/) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "redirected", f)
}

// QPainter * sharedPainter()
func (this *QPaintDevice) InheritSharedPainter(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "sharedPainter", f)
}

type QPaintDevice struct {
	*qtrt.CObject
}
type QPaintDevice_ITF interface {
	QPaintDevice_PTR() *QPaintDevice
}

func (ptr *QPaintDevice) QPaintDevice_PTR() *QPaintDevice { return ptr }

func (this *QPaintDevice) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPaintDevice) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPaintDeviceFromPointer(cthis unsafe.Pointer) *QPaintDevice {
	return &QPaintDevice{&qtrt.CObject{cthis}}
}
func (*QPaintDevice) NewFromPointer(cthis unsafe.Pointer) *QPaintDevice {
	return NewQPaintDeviceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpaintdevice.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPaintDevice()
func DeleteQPaintDevice(this *QPaintDevice) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintDeviceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpaintdevice.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int devType()
func (this *QPaintDevice) DevType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice7devTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool paintingActive()
func (this *QPaintDevice) PaintingActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice14paintingActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintdevice.h:76
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine()
func (this *QPaintDevice) PaintEngine() *QPaintEngine /*777 QPaintEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice11paintEngineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpaintdevice.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int width()
func (this *QPaintDevice) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int height()
func (this *QPaintDevice) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int widthMM()
func (this *QPaintDevice) WidthMM() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice7widthMMEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int heightMM()
func (this *QPaintDevice) HeightMM() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice8heightMMEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int logicalDpiX()
func (this *QPaintDevice) LogicalDpiX() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice11logicalDpiXEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:83
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int logicalDpiY()
func (this *QPaintDevice) LogicalDpiY() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice11logicalDpiYEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int physicalDpiX()
func (this *QPaintDevice) PhysicalDpiX() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice12physicalDpiXEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:85
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int physicalDpiY()
func (this *QPaintDevice) PhysicalDpiY() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice12physicalDpiYEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int devicePixelRatio()
func (this *QPaintDevice) DevicePixelRatio() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice16devicePixelRatioEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:87
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal devicePixelRatioF()
func (this *QPaintDevice) DevicePixelRatioF() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice17devicePixelRatioFEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:88
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int colorCount()
func (this *QPaintDevice) ColorCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice10colorCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int depth()
func (this *QPaintDevice) Depth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice5depthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:91
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] qreal devicePixelRatioFScale()
func (this *QPaintDevice) DevicePixelRatioFScale() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintDevice22devicePixelRatioFScaleEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}
func QPaintDevice_DevicePixelRatioFScale() float64 {
	var nilthis *QPaintDevice
	rv := nilthis.DevicePixelRatioFScale()
	return rv
}

// /usr/include/qt/QtGui/qpaintdevice.h:93
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QPaintDevice()
func NewQPaintDevice() *QPaintDevice {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintDeviceC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPaintDevice)
	return gothis
}

// /usr/include/qt/QtGui/qpaintdevice.h:101
// index:1
// Private inline Visibility=Default Availability=NotAvailable
// [-2] void QPaintDevice(const QPaintDevice &)
func NewQPaintDevice_1(arg0 *QPaintDevice) *QPaintDevice {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintDeviceC2ERKS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPaintDevice)
	return gothis
}

// /usr/include/qt/QtGui/qpaintdevice.h:94
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPaintDevice) Metric(metric int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice6metricENS_17PaintDeviceMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), metric)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:95
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void initPainter(QPainter *)
func (this *QPaintDevice) InitPainter(painter *QPainter /*777 QPainter **/) {
	var convArg0 = painter.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice11initPainterEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:96
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QPaintDevice * redirected(QPoint *)
func (this *QPaintDevice) Redirected(offset *qtcore.QPoint /*777 QPoint **/) *QPaintDevice /*777 QPaintDevice **/ {
	var convArg0 = offset.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice10redirectedEP6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpaintdevice.h:97
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QPainter * sharedPainter()
func (this *QPaintDevice) SharedPainter() *QPainter /*777 QPainter **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice13sharedPainterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPainterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

type QPaintDevice__PaintDeviceMetric = int

const QPaintDevice__PdmWidth QPaintDevice__PaintDeviceMetric = 1
const QPaintDevice__PdmHeight QPaintDevice__PaintDeviceMetric = 2
const QPaintDevice__PdmWidthMM QPaintDevice__PaintDeviceMetric = 3
const QPaintDevice__PdmHeightMM QPaintDevice__PaintDeviceMetric = 4
const QPaintDevice__PdmNumColors QPaintDevice__PaintDeviceMetric = 5
const QPaintDevice__PdmDepth QPaintDevice__PaintDeviceMetric = 6
const QPaintDevice__PdmDpiX QPaintDevice__PaintDeviceMetric = 7
const QPaintDevice__PdmDpiY QPaintDevice__PaintDeviceMetric = 8
const QPaintDevice__PdmPhysicalDpiX QPaintDevice__PaintDeviceMetric = 9
const QPaintDevice__PdmPhysicalDpiY QPaintDevice__PaintDeviceMetric = 10
const QPaintDevice__PdmDevicePixelRatio QPaintDevice__PaintDeviceMetric = 11
const QPaintDevice__PdmDevicePixelRatioScaled QPaintDevice__PaintDeviceMetric = 12

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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
