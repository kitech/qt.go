package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudiorolecontrol.h
// #include <qaudiorolecontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QAudioRoleControl struct {
	*QMediaControl
}
type QAudioRoleControl_ITF interface {
	QMediaControl_ITF
	QAudioRoleControl_PTR() *QAudioRoleControl
}

func (ptr *QAudioRoleControl) QAudioRoleControl_PTR() *QAudioRoleControl { return ptr }

func (this *QAudioRoleControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QAudioRoleControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQAudioRoleControlFromPointer(cthis unsafe.Pointer) *QAudioRoleControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QAudioRoleControl{bcthis0}
}
func (*QAudioRoleControl) NewFromPointer(cthis unsafe.Pointer) *QAudioRoleControl {
	return NewQAudioRoleControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudiorolecontrol.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAudioRoleControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAudioRoleControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiorolecontrol.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAudioRoleControl()

/*

 */
func DeleteQAudioRoleControl(this *QAudioRoleControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAudioRoleControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qaudiorolecontrol.h:58
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QAudio::Role audioRole() const

/*
Returns the audio role of the media played by the media service.

See also setAudioRole().
*/
func (this *QAudioRoleControl) AudioRole() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAudioRoleControl9audioRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qaudiorolecontrol.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setAudioRole(QAudio::Role)

/*
Sets the audio role of the media played by the media service.

See also audioRole().
*/
func (this *QAudioRoleControl) SetAudioRole(role int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAudioRoleControl12setAudioRoleEN6QAudio4RoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiorolecontrol.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void audioRoleChanged(QAudio::Role)

/*
Signal emitted when the audio role has changed.
*/
func (this *QAudioRoleControl) AudioRoleChanged(role int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAudioRoleControl16audioRoleChangedEN6QAudio4RoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiorolecontrol.h:67
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAudioRoleControl(QObject *)

/*
Construct a QAudioRoleControl with the given parent.
*/
func (*QAudioRoleControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioRoleControl {
	return NewQAudioRoleControl(parent)
}
func NewQAudioRoleControl(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioRoleControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAudioRoleControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioRoleControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioRoleControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiorolecontrol.h:67
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAudioRoleControl(QObject *)

/*
Construct a QAudioRoleControl with the given parent.
*/
func (*QAudioRoleControl) NewForInherit__() *QAudioRoleControl {
	return NewQAudioRoleControl__()
}
func NewQAudioRoleControl__() *QAudioRoleControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAudioRoleControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioRoleControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioRoleControl")
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
