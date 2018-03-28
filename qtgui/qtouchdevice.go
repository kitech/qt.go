package qtgui

// /usr/include/qt/QtGui/qtouchdevice.h
// #include <qtouchdevice.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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
type QTouchDevice struct {
	*qtrt.CObject
}
type QTouchDevice_ITF interface {
	QTouchDevice_PTR() *QTouchDevice
}

func (ptr *QTouchDevice) QTouchDevice_PTR() *QTouchDevice { return ptr }

func (this *QTouchDevice) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTouchDevice) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTouchDeviceFromPointer(cthis unsafe.Pointer) *QTouchDevice {
	return &QTouchDevice{&qtrt.CObject{cthis}}
}
func (*QTouchDevice) NewFromPointer(cthis unsafe.Pointer) *QTouchDevice {
	return NewQTouchDeviceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtouchdevice.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTouchDevice()

/*
Creates a new touch device instance. By default the name is empty, the only capability is Position and type is TouchScreen.
*/
func NewQTouchDevice() *QTouchDevice {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTouchDeviceC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTouchDeviceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTouchDevice)
	return gothis
}

// /usr/include/qt/QtGui/qtouchdevice.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTouchDevice()

/*

 */
func DeleteQTouchDevice(this *QTouchDevice) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTouchDeviceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtouchdevice.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*
Returns the touch device name.

This string may often be empty. It is however useful for systems that have more than one touch input device because there it can be used to differentiate between the devices (i.e. to tell from which device a QTouchEvent originates from).

See also setName().
*/
func (this *QTouchDevice) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTouchDevice4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtouchdevice.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] QTouchDevice::DeviceType type() const

/*
Returns the touch device type.

See also setType().
*/
func (this *QTouchDevice) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTouchDevice4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:80
// index:0
// Public Visibility=Default Availability=Available
// [4] QTouchDevice::Capabilities capabilities() const

/*
Returns the touch device capabilities.

See also setCapabilities().
*/
func (this *QTouchDevice) Capabilities() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTouchDevice12capabilitiesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumTouchPoints() const

/*
Returns the maximum number of simultaneous touch points (fingers) that can be detected.

This function was introduced in  Qt 5.2.

See also setMaximumTouchPoints().
*/
func (this *QTouchDevice) MaximumTouchPoints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTouchDevice18maximumTouchPointsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtouchdevice.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setName(const QString &)

/*
Sets the name (a unique identifier) for the device. In most systems it is enough to leave this unset and keep the default empty name. This identifier becomes important when having multiple touch devices and a need to differentiate between them.

See also name().
*/
func (this *QTouchDevice) SetName(name string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTouchDevice7setNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setType(QTouchDevice::DeviceType)

/*
Sets the device type devType.

See also type().
*/
func (this *QTouchDevice) SetType(devType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTouchDevice7setTypeENS_10DeviceTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), devType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCapabilities(QTouchDevice::Capabilities)

/*
Sets the capabilities caps supported by the device and its driver.

See also capabilities().
*/
func (this *QTouchDevice) SetCapabilities(caps int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTouchDevice15setCapabilitiesE6QFlagsINS_14CapabilityFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), caps)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumTouchPoints(int)

/*
Sets the maximum number of simultaneous touchpoints max supported by the device and its driver.

See also maximumTouchPoints().
*/
func (this *QTouchDevice) SetMaximumTouchPoints(max int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTouchDevice21setMaximumTouchPointsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), max)
	qtrt.ErrPrint(err, rv)
}

/*
This enum represents the type of device that generated a QTouchEvent.


*/
type QTouchDevice__DeviceType = int

// In this type of device, the touch surface and display are integrated. This means the surface and display typically have the same size, such that there is a direct relationship between the touch points' physical positions and the coordinate reported by QTouchEvent::TouchPoint. As a result, Qt allows the user to interact directly with multiple QWidgets and QGraphicsItems at the same time.
const QTouchDevice__TouchScreen QTouchDevice__DeviceType = 0

// In this type of device, the touch surface is separate from the display. There is not a direct relationship between the physical touch location and the on-screen coordinates. Instead, they are calculated relative to the current mouse position, and the user must use the touch-pad to move this reference point. Unlike touch-screens, Qt allows users to only interact with a single QWidget or QGraphicsItem at a time.
const QTouchDevice__TouchPad QTouchDevice__DeviceType = 1

/*


 */
type QTouchDevice__CapabilityFlag = int

//
const QTouchDevice__Position QTouchDevice__CapabilityFlag = 1

//
const QTouchDevice__Area QTouchDevice__CapabilityFlag = 2

//
const QTouchDevice__Pressure QTouchDevice__CapabilityFlag = 4

//
const QTouchDevice__Velocity QTouchDevice__CapabilityFlag = 8

//
const QTouchDevice__RawPositions QTouchDevice__CapabilityFlag = 16

//
const QTouchDevice__NormalizedPosition QTouchDevice__CapabilityFlag = 32

//
const QTouchDevice__MouseEmulation QTouchDevice__CapabilityFlag = 64

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
