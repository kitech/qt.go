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
// extern C begin: 64
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

// int metric(QPaintDevice::PaintDeviceMetric)
func (this *QPaintDevice) InheritMetric(f func(metric int) int) {
	qtrt.SetAllInheritCallback(this, "metric", f)
}

// void initPainter(QPainter *)
func (this *QPaintDevice) InheritInitPainter(f func(painter *QPainter /*777 QPainter **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initPainter", f)
}

// QPaintDevice * redirected(QPoint *)
func (this *QPaintDevice) InheritRedirected(f func(offset *qtcore.QPoint /*777 QPoint **/) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "redirected", f)
}

// QPainter * sharedPainter()
func (this *QPaintDevice) InheritSharedPainter(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "sharedPainter", f)
}

/*

 */
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

/*

 */
func DeleteQPaintDevice(this *QPaintDevice) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintDeviceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpaintdevice.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int devType() const

/*

 */
func (this *QPaintDevice) DevType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice7devTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool paintingActive() const

/*
Returns true if the device is currently being painted on, i.e. someone has called QPainter::begin() but not yet called QPainter::end() for this device; otherwise returns false.

See also QPainter::isActive().
*/
func (this *QPaintDevice) PaintingActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice14paintingActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintdevice.h:76
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine() const

/*
Returns a pointer to the paint engine used for drawing on the device.
*/
func (this *QPaintDevice) PaintEngine() *QPaintEngine /*777 QPaintEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice11paintEngineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpaintdevice.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int width() const

/*
Returns the width of the paint device in default coordinate system units (e.g. pixels for QPixmap and QWidget).

See also widthMM().
*/
func (this *QPaintDevice) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int height() const

/*
Returns the height of the paint device in default coordinate system units (e.g. pixels for QPixmap and QWidget).

See also heightMM().
*/
func (this *QPaintDevice) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int widthMM() const

/*
Returns the width of the paint device in millimeters. Due to platform limitations it may not be possible to use this function to determine the actual physical size of a widget on the screen.

See also width().
*/
func (this *QPaintDevice) WidthMM() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice7widthMMEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int heightMM() const

/*
Returns the height of the paint device in millimeters. Due to platform limitations it may not be possible to use this function to determine the actual physical size of a widget on the screen.

See also height().
*/
func (this *QPaintDevice) HeightMM() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice8heightMMEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int logicalDpiX() const

/*
Returns the horizontal resolution of the device in dots per inch, which is used when computing font sizes. For X11, this is usually the same as could be computed from widthMM().

Note that if the logicalDpiX() doesn't equal the physicalDpiX(), the corresponding QPaintEngine must handle the resolution mapping.

See also logicalDpiY() and physicalDpiX().
*/
func (this *QPaintDevice) LogicalDpiX() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice11logicalDpiXEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:83
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int logicalDpiY() const

/*
Returns the vertical resolution of the device in dots per inch, which is used when computing font sizes. For X11, this is usually the same as could be computed from heightMM().

Note that if the logicalDpiY() doesn't equal the physicalDpiY(), the corresponding QPaintEngine must handle the resolution mapping.

See also logicalDpiX() and physicalDpiY().
*/
func (this *QPaintDevice) LogicalDpiY() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice11logicalDpiYEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int physicalDpiX() const

/*
Returns the horizontal resolution of the device in dots per inch. For example, when printing, this resolution refers to the physical printer's resolution. The logical DPI on the other hand, refers to the resolution used by the actual paint engine.

Note that if the physicalDpiX() doesn't equal the logicalDpiX(), the corresponding QPaintEngine must handle the resolution mapping.

See also physicalDpiY() and logicalDpiX().
*/
func (this *QPaintDevice) PhysicalDpiX() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice12physicalDpiXEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:85
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int physicalDpiY() const

/*
Returns the horizontal resolution of the device in dots per inch. For example, when printing, this resolution refers to the physical printer's resolution. The logical DPI on the other hand, refers to the resolution used by the actual paint engine.

Note that if the physicalDpiY() doesn't equal the logicalDpiY(), the corresponding QPaintEngine must handle the resolution mapping.

See also physicalDpiX() and logicalDpiY().
*/
func (this *QPaintDevice) PhysicalDpiY() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice12physicalDpiYEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int devicePixelRatio() const

/*
Returns the device pixel ratio for device.

Common values are 1 for normal-dpi displays and 2 for high-dpi "retina" displays.
*/
func (this *QPaintDevice) DevicePixelRatio() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice16devicePixelRatioEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:87
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal devicePixelRatioF() const

/*
Returns the device pixel ratio for the device as a floating point number.

This function was introduced in  Qt 5.6.
*/
func (this *QPaintDevice) DevicePixelRatioF() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice17devicePixelRatioFEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:88
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int colorCount() const

/*
Returns the number of different colors available for the paint device. If the number of colors available is too great to be represented by the int data type, then INT_MAX will be returned instead.
*/
func (this *QPaintDevice) ColorCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice10colorCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int depth() const

/*
Returns the bit depth (number of bit planes) of the paint device.
*/
func (this *QPaintDevice) Depth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice5depthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:91
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] qreal devicePixelRatioFScale()

/*

 */
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

/*
Constructs a paint device. This constructor can be invoked only from subclasses of QPaintDevice.
*/
func (*QPaintDevice) NewForInherit() *QPaintDevice {
	return NewQPaintDevice()
}
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

/*
Constructs a paint device. This constructor can be invoked only from subclasses of QPaintDevice.
*/
func (*QPaintDevice) NewForInherit1(arg0 QPaintDevice_ITF) *QPaintDevice {
	return NewQPaintDevice1(arg0)
}
func NewQPaintDevice1(arg0 QPaintDevice_ITF) *QPaintDevice {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintDevice_PTR() != nil {
		convArg0 = arg0.QPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPaintDeviceC2ERKS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPaintDevice)
	return gothis
}

// /usr/include/qt/QtGui/qpaintdevice.h:94
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(QPaintDevice::PaintDeviceMetric) const

/*
Returns the metric information for the given paint device metric.

See also PaintDeviceMetric.
*/
func (this *QPaintDevice) Metric(metric int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice6metricENS_17PaintDeviceMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), metric)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:95
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void initPainter(QPainter *) const

/*

 */
func (this *QPaintDevice) InitPainter(painter QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice11initPainterEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpaintdevice.h:96
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QPaintDevice * redirected(QPoint *) const

/*

 */
func (this *QPaintDevice) Redirected(offset qtcore.QPoint_ITF /*777 QPoint **/) *QPaintDevice /*777 QPaintDevice **/ {
	var convArg0 unsafe.Pointer
	if offset != nil && offset.QPoint_PTR() != nil {
		convArg0 = offset.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice10redirectedEP6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpaintdevice.h:97
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QPainter * sharedPainter() const

/*

 */
func (this *QPaintDevice) SharedPainter() *QPainter /*777 QPainter **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QPaintDevice13sharedPainterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPainterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

/*
Describes the various metrics of a paint device.



See also metric() and devicePixelRatioF().

*/
type QPaintDevice__PaintDeviceMetric = int

// The width of the paint device in default coordinate system units (e.g. pixels for QPixmap and QWidget). See also width().
const QPaintDevice__PdmWidth QPaintDevice__PaintDeviceMetric = 1

// The height of the paint device in default coordinate system units (e.g. pixels for QPixmap and QWidget). See also height().
const QPaintDevice__PdmHeight QPaintDevice__PaintDeviceMetric = 2

// The width of the paint device in millimeters. See also widthMM().
const QPaintDevice__PdmWidthMM QPaintDevice__PaintDeviceMetric = 3

// The height of the paint device in millimeters. See also heightMM().
const QPaintDevice__PdmHeightMM QPaintDevice__PaintDeviceMetric = 4

// The number of different colors available for the paint device. See also colorCount().
const QPaintDevice__PdmNumColors QPaintDevice__PaintDeviceMetric = 5

// The bit depth (number of bit planes) of the paint device. See also depth().
const QPaintDevice__PdmDepth QPaintDevice__PaintDeviceMetric = 6

// The horizontal resolution of the device in dots per inch. See also logicalDpiX().
const QPaintDevice__PdmDpiX QPaintDevice__PaintDeviceMetric = 7

// The vertical resolution of the device in dots per inch. See also logicalDpiY().
const QPaintDevice__PdmDpiY QPaintDevice__PaintDeviceMetric = 8

// The horizontal resolution of the device in dots per inch. See also physicalDpiX().
const QPaintDevice__PdmPhysicalDpiX QPaintDevice__PaintDeviceMetric = 9

//
const QPaintDevice__PdmPhysicalDpiY QPaintDevice__PaintDeviceMetric = 10

//
const QPaintDevice__PdmDevicePixelRatio QPaintDevice__PaintDeviceMetric = 11

//
const QPaintDevice__PdmDevicePixelRatioScaled QPaintDevice__PaintDeviceMetric = 12

func (this *QPaintDevice) PaintDeviceMetricItemName(val int) string {
	switch val {
	case QPaintDevice__PdmWidth: // 1
		return "PdmWidth"
	case QPaintDevice__PdmHeight: // 2
		return "PdmHeight"
	case QPaintDevice__PdmWidthMM: // 3
		return "PdmWidthMM"
	case QPaintDevice__PdmHeightMM: // 4
		return "PdmHeightMM"
	case QPaintDevice__PdmNumColors: // 5
		return "PdmNumColors"
	case QPaintDevice__PdmDepth: // 6
		return "PdmDepth"
	case QPaintDevice__PdmDpiX: // 7
		return "PdmDpiX"
	case QPaintDevice__PdmDpiY: // 8
		return "PdmDpiY"
	case QPaintDevice__PdmPhysicalDpiX: // 9
		return "PdmPhysicalDpiX"
	case QPaintDevice__PdmPhysicalDpiY: // 10
		return "PdmPhysicalDpiY"
	case QPaintDevice__PdmDevicePixelRatio: // 11
		return "PdmDevicePixelRatio"
	case QPaintDevice__PdmDevicePixelRatioScaled: // 12
		return "PdmDevicePixelRatioScaled"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPaintDevice_PaintDeviceMetricItemName(val int) string {
	var nilthis *QPaintDevice
	return nilthis.PaintDeviceMetricItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10735() {
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
