package qtmultimedia

// /usr/include/qt/QtMultimedia/qcamerainfocontrol.h
// #include <qcamerainfocontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QCameraInfoControl struct {
	*QMediaControl
}
type QCameraInfoControl_ITF interface {
	QMediaControl_ITF
	QCameraInfoControl_PTR() *QCameraInfoControl
}

func (ptr *QCameraInfoControl) QCameraInfoControl_PTR() *QCameraInfoControl { return ptr }

func (this *QCameraInfoControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QCameraInfoControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQCameraInfoControlFromPointer(cthis unsafe.Pointer) *QCameraInfoControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QCameraInfoControl{bcthis0}
}
func (*QCameraInfoControl) NewFromPointer(cthis unsafe.Pointer) *QCameraInfoControl {
	return NewQCameraInfoControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcamerainfocontrol.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraInfoControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCameraInfoControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcamerainfocontrol.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCameraInfoControl()

/*

 */
func DeleteQCameraInfoControl(this *QCameraInfoControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCameraInfoControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcamerainfocontrol.h:57
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QCamera::Position cameraPosition(const QString &) const

/*
Returns the physical position of the camera named deviceName on the hardware system.
*/
func (this *QCameraInfoControl) CameraPosition(deviceName string) int {
	var tmpArg0 = qtcore.NewQString_5(deviceName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCameraInfoControl14cameraPositionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamerainfocontrol.h:58
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int cameraOrientation(const QString &) const

/*
Returns the physical orientation of the sensor for the camera named deviceName.

The value is the orientation angle (clockwise, in steps of 90 degrees) of the camera sensor in relation to the display in its natural orientation.
*/
func (this *QCameraInfoControl) CameraOrientation(deviceName string) int {
	var tmpArg0 = qtcore.NewQString_5(deviceName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCameraInfoControl17cameraOrientationERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qcamerainfocontrol.h:61
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraInfoControl(QObject *)

/*
Constructs a camera info control with the given parent.
*/
func NewQCameraInfoControl(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraInfoControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCameraInfoControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraInfoControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraInfoControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamerainfocontrol.h:61
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraInfoControl(QObject *)

/*
Constructs a camera info control with the given parent.
*/
func NewQCameraInfoControl__() *QCameraInfoControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCameraInfoControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraInfoControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraInfoControl")
	return gothis
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
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
