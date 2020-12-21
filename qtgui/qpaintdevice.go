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
// extern C begin: 45
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

/*
 */
// size 24
type QPaintDevice struct {
	*qtrt.CObject
}
type QPaintDevice_ITF interface {
	QPaintDevice_PTR() *QPaintDevice
}

func (ptr *QPaintDevice) QPaintDevice_PTR() *QPaintDevice { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QPaintDeviceFromptr(cthis Voidptr) *QPaintDevice {
	return &QPaintDevice{&qtrt.CObject{cthis}}
}
func (*QPaintDevice) Fromptr(cthis Voidptr) *QPaintDevice {
	return QPaintDeviceFromptr(cthis)
}

// /usr/include/qt/QtGui/qpaintdevice.h:74
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [4] int devType() const

/*
 */
func (this *QPaintDevice) DevType() int {
	rv, err := qtrt.Qtcc3(2740044500, "_ZNK12QPaintDevice7devTypeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:75
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool paintingActive() const

/*
 */
func (this *QPaintDevice) PaintingActive() bool {
	rv, err := qtrt.Qtcc3(4098809866, "_ZNK12QPaintDevice14paintingActiveEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpaintdevice.h:76
// index:0
// Public purevirtual virtual Direct Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine() const

/*
 */
func (this *QPaintDevice) PaintEngine() *QPaintEngine /*777 QPaintEngine **/ {
	rv, err := qtrt.Qtcc3(1815356186, "_ZNK12QPaintDevice11paintEngineEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QPaintEngineFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpaintdevice.h:78
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int width() const

/*
 */
func (this *QPaintDevice) Width() int {
	rv, err := qtrt.Qtcc3(4251465228, "_ZNK12QPaintDevice5widthEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:79
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int height() const

/*
 */
func (this *QPaintDevice) Height() int {
	rv, err := qtrt.Qtcc3(1938128564, "_ZNK12QPaintDevice6heightEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:80
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int widthMM() const

/*
 */
func (this *QPaintDevice) WidthMM() int {
	rv, err := qtrt.Qtcc3(4290881768, "_ZNK12QPaintDevice7widthMMEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:81
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int heightMM() const

/*
 */
func (this *QPaintDevice) HeightMM() int {
	rv, err := qtrt.Qtcc3(4250480684, "_ZNK12QPaintDevice8heightMMEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:82
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int logicalDpiX() const

/*
 */
func (this *QPaintDevice) LogicalDpiX() int {
	rv, err := qtrt.Qtcc3(1690059126, "_ZNK12QPaintDevice11logicalDpiXEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:83
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int logicalDpiY() const

/*
 */
func (this *QPaintDevice) LogicalDpiY() int {
	rv, err := qtrt.Qtcc3(1702767425, "_ZNK12QPaintDevice11logicalDpiYEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:84
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int physicalDpiX() const

/*
 */
func (this *QPaintDevice) PhysicalDpiX() int {
	rv, err := qtrt.Qtcc3(1606783227, "_ZNK12QPaintDevice12physicalDpiXEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:85
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int physicalDpiY() const

/*
 */
func (this *QPaintDevice) PhysicalDpiY() int {
	rv, err := qtrt.Qtcc3(1577581260, "_ZNK12QPaintDevice12physicalDpiYEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:86
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int devicePixelRatio() const

/*
 */
func (this *QPaintDevice) DevicePixelRatio() int {
	rv, err := qtrt.Qtcc3(2416087545, "_ZNK12QPaintDevice16devicePixelRatioEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:87
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [8] qreal devicePixelRatioF() const

/*
 */
func (this *QPaintDevice) DevicePixelRatioF() float64 {
	rv, err := qtrt.Qtcc3(1021285250, "_ZNK12QPaintDevice17devicePixelRatioFEv", qtrt.FFITO_DOUBLE,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:88
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int colorCount() const

/*
 */
func (this *QPaintDevice) ColorCount() int {
	rv, err := qtrt.Qtcc3(3927565789, "_ZNK12QPaintDevice10colorCountEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:89
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int depth() const

/*
 */
func (this *QPaintDevice) Depth() int {
	rv, err := qtrt.Qtcc3(1241268582, "_ZNK12QPaintDevice5depthEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpaintdevice.h:91
// index:0
// Public static inline Direct Visibility=Default Availability=Available
// [8] qreal devicePixelRatioFScale()

/*
 */
func (this *QPaintDevice) DevicePixelRatioFScale() float64 {
	rv, err := qtrt.Qtcc3(3846552345, "_ZN12QPaintDevice22devicePixelRatioFScaleEv", qtrt.FFITO_POINTER)
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}
func QPaintDevice_DevicePixelRatioFScale() float64 {
	var nilthis *QPaintDevice
	rv := nilthis.DevicePixelRatioFScale()
	return rv
}

func DeleteQPaintDevice(this *QPaintDevice) {
	rv, err := qtrt.Qtcc1(0, "_ZN12QPaintDeviceD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QPaintDevice__PaintDeviceMetric = int

//
const QPaintDevice__PdmWidth QPaintDevice__PaintDeviceMetric = 1

//
const QPaintDevice__PdmHeight QPaintDevice__PaintDeviceMetric = 2

//
const QPaintDevice__PdmWidthMM QPaintDevice__PaintDeviceMetric = 3

//
const QPaintDevice__PdmHeightMM QPaintDevice__PaintDeviceMetric = 4

//
const QPaintDevice__PdmNumColors QPaintDevice__PaintDeviceMetric = 5

//
const QPaintDevice__PdmDepth QPaintDevice__PaintDeviceMetric = 6

//
const QPaintDevice__PdmDpiX QPaintDevice__PaintDeviceMetric = 7

//
const QPaintDevice__PdmDpiY QPaintDevice__PaintDeviceMetric = 8

//
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

func init_unused_10051() {
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
