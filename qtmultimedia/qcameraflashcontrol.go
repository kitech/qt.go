package qtmultimedia

// /usr/include/qt/QtMultimedia/qcameraflashcontrol.h
// #include <qcameraflashcontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QCameraFlashControl struct {
	*QMediaControl
}
type QCameraFlashControl_ITF interface {
	QMediaControl_ITF
	QCameraFlashControl_PTR() *QCameraFlashControl
}

func (ptr *QCameraFlashControl) QCameraFlashControl_PTR() *QCameraFlashControl { return ptr }

func (this *QCameraFlashControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QCameraFlashControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQCameraFlashControlFromPointer(cthis unsafe.Pointer) *QCameraFlashControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QCameraFlashControl{bcthis0}
}
func (*QCameraFlashControl) NewFromPointer(cthis unsafe.Pointer) *QCameraFlashControl {
	return NewQCameraFlashControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcameraflashcontrol.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraFlashControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraFlashControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcameraflashcontrol.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCameraFlashControl()

/*

 */
func DeleteQCameraFlashControl(this *QCameraFlashControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraFlashControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcameraflashcontrol.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QCameraExposure::FlashModes flashMode() const

/*
Returns the current flash mode.

See also setFlashMode().
*/
func (this *QCameraFlashControl) FlashMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraFlashControl9flashModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameraflashcontrol.h:62
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setFlashMode(QCameraExposure::FlashModes)

/*
Set the current flash mode.

Usually a single QCameraExposure::FlashMode flag is used, but some non conflicting flags combination are also allowed, like QCameraExposure::FlashManual | QCameraExposure::FlashSlowSyncRearCurtain.

See also flashMode().
*/
func (this *QCameraFlashControl) SetFlashMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraFlashControl12setFlashModeE6QFlagsIN15QCameraExposure9FlashModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraflashcontrol.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isFlashModeSupported(QCameraExposure::FlashModes) const

/*
Return true if the reqested flash mode is supported. Some QCameraExposure::FlashMode values can be combined, for example QCameraExposure::FlashManual | QCameraExposure::FlashSlowSyncRearCurtain
*/
func (this *QCameraFlashControl) IsFlashModeSupported(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraFlashControl20isFlashModeSupportedE6QFlagsIN15QCameraExposure9FlashModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraflashcontrol.h:65
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isFlashReady() const

/*
Returns true if flash is charged.
*/
func (this *QCameraFlashControl) IsFlashReady() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraFlashControl12isFlashReadyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraflashcontrol.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void flashReady(bool)

/*
Signal emitted when flash state changes to ready.
*/
func (this *QCameraFlashControl) FlashReady(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraFlashControl10flashReadyEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraflashcontrol.h:71
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraFlashControl(QObject *)

/*
Constructs a camera flash control object with parent.
*/
func NewQCameraFlashControl(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraFlashControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraFlashControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFlashControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraFlashControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcameraflashcontrol.h:71
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraFlashControl(QObject *)

/*
Constructs a camera flash control object with parent.
*/
func NewQCameraFlashControl__() *QCameraFlashControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraFlashControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFlashControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraFlashControl")
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
