//  header block begin
// /usr/include/qt/QtGui/qtouchdevice.h
// #include <qtouchdevice.h>
// #include <QtGui>
package qtgui

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
type QTouchDevice struct {
	*qtrt.CObject
}

func (this *QTouchDevice) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qtouchdevice.h:73
// index:0
// void QTouchDevice()
func NewQTouchDevice() *QTouchDevice {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTouchDeviceC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTouchDeviceFromPointer(cthis)
	return gothis
}
func NewQTouchDeviceFromPointer(cthis unsafe.Pointer) *QTouchDevice {
	return &QTouchDevice{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qtouchdevice.h:74
// index:0
// void ~QTouchDevice()
func DeleteQTouchDevice(*QTouchDevice) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTouchDeviceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:76
// index:0
// static
// QList<const QTouchDevice *> devices()
func (this *QTouchDevice) Devices() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTouchDevice7devicesEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTouchDevice_Devices() {
	// 0: (), ()
	var nilthis *QTouchDevice
	nilthis.Devices()
}

// /usr/include/qt/QtGui/qtouchdevice.h:78
// index:0
// QString name()
func (this *QTouchDevice) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTouchDevice4nameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:79
// index:0
// QTouchDevice::DeviceType type()
func (this *QTouchDevice) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTouchDevice4typeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:80
// index:0
// QTouchDevice::Capabilities capabilities()
func (this *QTouchDevice) Capabilities() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTouchDevice12capabilitiesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:81
// index:0
// int maximumTouchPoints()
func (this *QTouchDevice) MaximumTouchPoints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTouchDevice18maximumTouchPointsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:83
// index:0
// void setName(const class QString &)
func (this *QTouchDevice) SetName(name unsafe.Pointer) {
	// 0: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTouchDevice7setNameERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:84
// index:0
// void setType(enum QTouchDevice::DeviceType)
func (this *QTouchDevice) SetType(devType int) {
	// 0: (, devType QTouchDevice::DeviceType), (&devType)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTouchDevice7setTypeENS_10DeviceTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &devType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:85
// index:0
// void setCapabilities(QTouchDevice::Capabilities)
func (this *QTouchDevice) SetCapabilities(caps int) {
	// 0: (, QFlags<QTouchDevice::CapabilityFlag> caps), (caps)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTouchDevice15setCapabilitiesE6QFlagsINS_14CapabilityFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), caps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:86
// index:0
// void setMaximumTouchPoints(int)
func (this *QTouchDevice) SetMaximumTouchPoints(max int) {
	// 0: (, max int), (&max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTouchDevice21setMaximumTouchPointsEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &max)
	gopp.ErrPrint(err, rv)
}

//  body block end
