package qtmultimedia

// /usr/include/qt/QtMultimedia/qcamerafeedbackcontrol.h
// #include <qcamerafeedbackcontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QCameraFeedbackControl struct {
	*QMediaControl
}
type QCameraFeedbackControl_ITF interface {
	QMediaControl_ITF
	QCameraFeedbackControl_PTR() *QCameraFeedbackControl
}

func (ptr *QCameraFeedbackControl) QCameraFeedbackControl_PTR() *QCameraFeedbackControl { return ptr }

func (this *QCameraFeedbackControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QCameraFeedbackControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQCameraFeedbackControlFromPointer(cthis unsafe.Pointer) *QCameraFeedbackControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QCameraFeedbackControl{bcthis0}
}
func (*QCameraFeedbackControl) NewFromPointer(cthis unsafe.Pointer) *QCameraFeedbackControl {
	return NewQCameraFeedbackControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcamerafeedbackcontrol.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraFeedbackControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraFeedbackControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcamerafeedbackcontrol.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCameraFeedbackControl()

/*

 */
func DeleteQCameraFeedbackControl(this *QCameraFeedbackControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraFeedbackControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcamerafeedbackcontrol.h:77
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isEventFeedbackLocked(QCameraFeedbackControl::EventType) const

/*
Returns true if the feedback setting for event is locked. This may be true because of legal compliance issues, or because configurability of this event's feedback is not supported.

This function was introduced in  Qt 5.0.
*/
func (this *QCameraFeedbackControl) IsEventFeedbackLocked(arg0 int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraFeedbackControl21isEventFeedbackLockedENS_9EventTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamerafeedbackcontrol.h:79
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isEventFeedbackEnabled(QCameraFeedbackControl::EventType) const

/*
Returns true if the feedback for event is enabled.

This function was introduced in  Qt 5.0.
*/
func (this *QCameraFeedbackControl) IsEventFeedbackEnabled(arg0 int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QCameraFeedbackControl22isEventFeedbackEnabledENS_9EventTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamerafeedbackcontrol.h:81
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool setEventFeedbackEnabled(QCameraFeedbackControl::EventType, bool)

/*
Turns on feedback for the specific event if enabled is true, otherwise disables the feedback. Returns true if the feedback could be modified, or false otherwise (e.g. this feedback type is locked).

This function was introduced in  Qt 5.0.

See also isEventFeedbackEnabled().
*/
func (this *QCameraFeedbackControl) SetEventFeedbackEnabled(arg0 int, arg1 bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraFeedbackControl23setEventFeedbackEnabledENS_9EventTypeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamerafeedbackcontrol.h:82
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void resetEventFeedback(QCameraFeedbackControl::EventType)

/*
Restores the feedback setting for this event to its default setting.

This function was introduced in  Qt 5.0.
*/
func (this *QCameraFeedbackControl) ResetEventFeedback(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraFeedbackControl18resetEventFeedbackENS_9EventTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafeedbackcontrol.h:84
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool setEventFeedbackSound(QCameraFeedbackControl::EventType, const QString &)

/*
When the given event occurs, the sound effect referenced by filePath will be played instead of the default sound.

If this feedback type is locked, or if the supplied path is inaccessible, this function will return false. In addition, some forms of feedback may be non-auditory (e.g. a red light, or a vibration), and false may be returned in this case.

The file referenced should be linear PCM (WAV format).

Note: In the case that a valid file path to an unsupported file is given, this function will return true but the feedback will use the original setting.

This function was introduced in  Qt 5.0.
*/
func (this *QCameraFeedbackControl) SetEventFeedbackSound(arg0 int, filePath string) bool {
	var tmpArg1 = qtcore.NewQString_5(filePath)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraFeedbackControl21setEventFeedbackSoundENS_9EventTypeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamerafeedbackcontrol.h:87
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraFeedbackControl(QObject *)

/*
Constructs a camera feedback control object with parent.
*/
func (*QCameraFeedbackControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraFeedbackControl {
	return NewQCameraFeedbackControl(parent)
}
func NewQCameraFeedbackControl(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraFeedbackControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraFeedbackControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFeedbackControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraFeedbackControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamerafeedbackcontrol.h:87
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraFeedbackControl(QObject *)

/*
Constructs a camera feedback control object with parent.
*/
func (*QCameraFeedbackControl) NewForInherit__() *QCameraFeedbackControl {
	return NewQCameraFeedbackControl__()
}
func NewQCameraFeedbackControl__() *QCameraFeedbackControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN22QCameraFeedbackControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFeedbackControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraFeedbackControl")
	return gothis
}

/*
This enumeration describes certain events that occur during camera usage. You can associate some form of feedback to be given when the event occurs, or check whether feedback for this event is enabled or locked so that changes cannot be made.


*/
type QCameraFeedbackControl__EventType = int

// The viewfinder stream was started (even if not visible)
const QCameraFeedbackControl__ViewfinderStarted QCameraFeedbackControl__EventType = 1

// The viewfinder stream was stopped
const QCameraFeedbackControl__ViewfinderStopped QCameraFeedbackControl__EventType = 2

// An image was captured but not yet fully processed
const QCameraFeedbackControl__ImageCaptured QCameraFeedbackControl__EventType = 3

// An image is fully available and saved somewhere.
const QCameraFeedbackControl__ImageSaved QCameraFeedbackControl__EventType = 4

// An error occurred while capturing an image
const QCameraFeedbackControl__ImageError QCameraFeedbackControl__EventType = 5

// Video recording has started
const QCameraFeedbackControl__RecordingStarted QCameraFeedbackControl__EventType = 6

// Video recording is in progress
const QCameraFeedbackControl__RecordingInProgress QCameraFeedbackControl__EventType = 7

// Video recording has stopped
const QCameraFeedbackControl__RecordingStopped QCameraFeedbackControl__EventType = 8

// The camera is trying to automatically focus
const QCameraFeedbackControl__AutoFocusInProgress QCameraFeedbackControl__EventType = 9

//
const QCameraFeedbackControl__AutoFocusLocked QCameraFeedbackControl__EventType = 10

//
const QCameraFeedbackControl__AutoFocusFailed QCameraFeedbackControl__EventType = 11

func (this *QCameraFeedbackControl) EventTypeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCameraFeedbackControl_EventTypeItemName(val int) string {
	var nilthis *QCameraFeedbackControl
	return nilthis.EventTypeItemName(val)
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
