package qtmultimedia

// /usr/include/qt/QtMultimedia/qcameralockscontrol.h
// #include <qcameralockscontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QCameraLocksControl struct {
	*QMediaControl
}
type QCameraLocksControl_ITF interface {
	QMediaControl_ITF
	QCameraLocksControl_PTR() *QCameraLocksControl
}

func (ptr *QCameraLocksControl) QCameraLocksControl_PTR() *QCameraLocksControl { return ptr }

func (this *QCameraLocksControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QCameraLocksControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQCameraLocksControlFromPointer(cthis unsafe.Pointer) *QCameraLocksControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QCameraLocksControl{bcthis0}
}
func (*QCameraLocksControl) NewFromPointer(cthis unsafe.Pointer) *QCameraLocksControl {
	return NewQCameraLocksControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcameralockscontrol.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraLocksControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraLocksControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcameralockscontrol.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCameraLocksControl()

/*

 */
func DeleteQCameraLocksControl(this *QCameraLocksControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraLocksControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcameralockscontrol.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QCamera::LockTypes supportedLocks() const

/*
Returns the lock types, the camera supports.
*/
func (this *QCameraLocksControl) SupportedLocks() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraLocksControl14supportedLocksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameralockscontrol.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QCamera::LockStatus lockStatus(QCamera::LockType) const

/*
Returns the camera lock status.
*/
func (this *QCameraLocksControl) LockStatus(lock int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraLocksControl10lockStatusEN7QCamera8LockTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), lock)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameralockscontrol.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void searchAndLock(QCamera::LockTypes)

/*
Request camera locks.
*/
func (this *QCameraLocksControl) SearchAndLock(locks int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraLocksControl13searchAndLockE6QFlagsIN7QCamera8LockTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), locks)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameralockscontrol.h:64
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void unlock(QCamera::LockTypes)

/*
Unlock camera locks.
*/
func (this *QCameraLocksControl) Unlock(locks int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraLocksControl6unlockE6QFlagsIN7QCamera8LockTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), locks)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameralockscontrol.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lockStatusChanged(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)

/*
Signals the lock type status was changed with the specified reason.
*/
func (this *QCameraLocksControl) LockStatusChanged(type_ int, status int, reason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraLocksControl17lockStatusChangedEN7QCamera8LockTypeENS0_10LockStatusENS0_16LockChangeReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, status, reason)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameralockscontrol.h:70
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraLocksControl(QObject *)

/*
Constructs a camera locks control object with parent.
*/
func (*QCameraLocksControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraLocksControl {
	return NewQCameraLocksControl(parent)
}
func NewQCameraLocksControl(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraLocksControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraLocksControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraLocksControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraLocksControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcameralockscontrol.h:70
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraLocksControl(QObject *)

/*
Constructs a camera locks control object with parent.
*/
func (*QCameraLocksControl) NewForInheritp() *QCameraLocksControl {
	return NewQCameraLocksControlp()
}
func NewQCameraLocksControlp() *QCameraLocksControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraLocksControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraLocksControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraLocksControl")
	return gothis
}

//  body block end

//  keep block begin

func init_unused_11835() {
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
