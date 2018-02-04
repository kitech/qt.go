package qtgui

// /usr/include/qt/QtGui/qtouchdevice.h
// #include <qtouchdevice.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QTouchDevice struct {
	*qtrt.CObject
}

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
func NewQTouchDevice() *QTouchDevice {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTouchDeviceC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTouchDeviceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTouchDevice)
	return gothis
}

// /usr/include/qt/QtGui/qtouchdevice.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTouchDevice()
func DeleteQTouchDevice(this *QTouchDevice) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTouchDeviceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtouchdevice.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QTouchDevice) Name() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTouchDevice4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qtouchdevice.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] QTouchDevice::DeviceType type()
func (this *QTouchDevice) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTouchDevice4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:80
// index:0
// Public Visibility=Default Availability=Available
// [4] QTouchDevice::Capabilities capabilities()
func (this *QTouchDevice) Capabilities() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTouchDevice12capabilitiesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumTouchPoints()
func (this *QTouchDevice) MaximumTouchPoints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTouchDevice18maximumTouchPointsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtouchdevice.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setName(const QString &)
func (this *QTouchDevice) SetName(name *qtcore.QString) {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTouchDevice7setNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setType(enum QTouchDevice::DeviceType)
func (this *QTouchDevice) SetType(devType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTouchDevice7setTypeENS_10DeviceTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), devType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCapabilities(QTouchDevice::Capabilities)
func (this *QTouchDevice) SetCapabilities(caps int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTouchDevice15setCapabilitiesE6QFlagsINS_14CapabilityFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), caps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumTouchPoints(int)
func (this *QTouchDevice) SetMaximumTouchPoints(max int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTouchDevice21setMaximumTouchPointsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), max)
	gopp.ErrPrint(err, rv)
}

type QTouchDevice__DeviceType = int

const QTouchDevice__TouchScreen QTouchDevice__DeviceType = 0
const QTouchDevice__TouchPad QTouchDevice__DeviceType = 1

type QTouchDevice__CapabilityFlag = int

const QTouchDevice__Position QTouchDevice__CapabilityFlag = 1
const QTouchDevice__Area QTouchDevice__CapabilityFlag = 2
const QTouchDevice__Pressure QTouchDevice__CapabilityFlag = 4
const QTouchDevice__Velocity QTouchDevice__CapabilityFlag = 8
const QTouchDevice__RawPositions QTouchDevice__CapabilityFlag = 16
const QTouchDevice__NormalizedPosition QTouchDevice__CapabilityFlag = 32
const QTouchDevice__MouseEmulation QTouchDevice__CapabilityFlag = 64

//  body block end
