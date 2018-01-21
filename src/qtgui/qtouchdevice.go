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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQTouchDeviceFromPointer(cthis unsafe.Pointer) *QTouchDevice {
	return &QTouchDevice{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qtouchdevice.h:73
// index:0
// Public
// void QTouchDevice()
func NewQTouchDevice() *QTouchDevice {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTouchDeviceC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTouchDeviceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtouchdevice.h:74
// index:0
// Public
// void ~QTouchDevice()
func DeleteQTouchDevice(*QTouchDevice) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTouchDeviceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:78
// index:0
// Public
// QString name()
func (this *QTouchDevice) Name() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTouchDevice4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtouchdevice.h:79
// index:0
// Public
// QTouchDevice::DeviceType type()
func (this *QTouchDevice) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTouchDevice4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:80
// index:0
// Public
// QTouchDevice::Capabilities capabilities()
func (this *QTouchDevice) Capabilities() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTouchDevice12capabilitiesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:81
// index:0
// Public
// int maximumTouchPoints()
func (this *QTouchDevice) MaximumTouchPoints() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTouchDevice18maximumTouchPointsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtouchdevice.h:83
// index:0
// Public
// void setName(const class QString &)
func (this *QTouchDevice) SetName(name *qtcore.QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTouchDevice7setNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:84
// index:0
// Public
// void setType(enum QTouchDevice::DeviceType)
func (this *QTouchDevice) SetType(devType int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTouchDevice7setTypeENS_10DeviceTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &devType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtouchdevice.h:86
// index:0
// Public
// void setMaximumTouchPoints(int)
func (this *QTouchDevice) SetMaximumTouchPoints(max int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTouchDevice21setMaximumTouchPointsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &max)
	gopp.ErrPrint(err, rv)
}

//  body block end
