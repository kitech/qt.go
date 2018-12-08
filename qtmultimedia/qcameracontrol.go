package qtmultimedia

// /usr/include/qt/QtMultimedia/qcameracontrol.h
// #include <qcameracontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QCameraControl struct {
	*QMediaControl
}
type QCameraControl_ITF interface {
	QMediaControl_ITF
	QCameraControl_PTR() *QCameraControl
}

func (ptr *QCameraControl) QCameraControl_PTR() *QCameraControl { return ptr }

func (this *QCameraControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QCameraControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQCameraControlFromPointer(cthis unsafe.Pointer) *QCameraControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QCameraControl{bcthis0}
}
func (*QCameraControl) NewFromPointer(cthis unsafe.Pointer) *QCameraControl {
	return NewQCameraControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcameracontrol.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QCameraControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcameracontrol.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCameraControl()

/*

 */
func DeleteQCameraControl(this *QCameraControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QCameraControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcameracontrol.h:68
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QCamera::State state() const

/*
Returns the state of the camera service.

See also setState() and QCamera::state.
*/
func (this *QCameraControl) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QCameraControl5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameracontrol.h:69
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setState(QCamera::State)

/*
Sets the camera state.

State changes are synchronous and indicate user intention, while camera status is used as a feedback mechanism to inform application about backend status. Status changes are reported asynchronously with QCameraControl::statusChanged() signal.

See also state() and QCamera::State.
*/
func (this *QCameraControl) SetState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QCameraControl8setStateEN7QCamera5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameracontrol.h:71
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QCamera::Status status() const

/*
Returns the status of the camera service.

See also QCamera::state.
*/
func (this *QCameraControl) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QCameraControl6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameracontrol.h:73
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QCamera::CaptureModes captureMode() const

/*
Returns the current capture mode.

See also setCaptureMode().
*/
func (this *QCameraControl) CaptureMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QCameraControl11captureModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameracontrol.h:74
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setCaptureMode(QCamera::CaptureModes)

/*
Sets the current capture mode.

The capture mode changes are synchronous and allowed in any camera state.

If the capture mode is changed while camera is active, it's recommended to change status to QCamera::LoadedStatus and start activating the camera in the next event loop with the status changed to QCamera::StartingStatus. This allows the capture settings to be applied before camera is started. Than change the status to QCamera::StartedStatus when the capture mode change is done.

See also captureMode().
*/
func (this *QCameraControl) SetCaptureMode(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QCameraControl14setCaptureModeE6QFlagsIN7QCamera11CaptureModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameracontrol.h:75
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isCaptureModeSupported(QCamera::CaptureModes) const

/*
Returns true if the capture mode is suported.
*/
func (this *QCameraControl) IsCaptureModeSupported(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QCameraControl22isCaptureModeSupportedE6QFlagsIN7QCamera11CaptureModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameracontrol.h:77
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool canChangeProperty(QCameraControl::PropertyChangeType, QCamera::Status) const

/*
Returns true if backend can effectively apply changing camera properties of changeType type while the camera state is QCamera::Active and camera status matches status parameter.

If backend doesn't support applying this change in the active state, it will be stopped before the settings are changed and restarted after. Otherwise the backend should apply the change in the current state, with the camera status indicating the progress, if necessary.
*/
func (this *QCameraControl) CanChangeProperty(changeType int, status int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QCameraControl17canChangePropertyENS_18PropertyChangeTypeEN7QCamera6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), changeType, status)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameracontrol.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QCamera::State)

/*
Signal emitted when the camera state changes.

In most cases the state chage is caused by QCameraControl::setState(), but if critical error has occurred the state changes to QCamera::UnloadedState.
*/
func (this *QCameraControl) StateChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QCameraControl12stateChangedEN7QCamera5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameracontrol.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void statusChanged(QCamera::Status)

/*
Signal emitted when the camera status changes.
*/
func (this *QCameraControl) StatusChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QCameraControl13statusChangedEN7QCamera6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameracontrol.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void error(int, const QString &)

/*
Signal emitted when an error occurs with error code error and a description of the error errorString.
*/
func (this *QCameraControl) Error(error int, errorString string) {
	var tmpArg1 = qtcore.NewQString5(errorString)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QCameraControl5errorEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameracontrol.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void captureModeChanged(QCamera::CaptureModes)

/*
Signal emitted when the camera capture mode changes.
*/
func (this *QCameraControl) CaptureModeChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QCameraControl18captureModeChangedE6QFlagsIN7QCamera11CaptureModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameracontrol.h:86
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraControl(QObject *)

/*
Constructs a camera control object with parent.
*/
func (*QCameraControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraControl {
	return NewQCameraControl(parent)
}
func NewQCameraControl(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QCameraControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcameracontrol.h:86
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraControl(QObject *)

/*
Constructs a camera control object with parent.
*/
func (*QCameraControl) NewForInheritp() *QCameraControl {
	return NewQCameraControlp()
}
func NewQCameraControlp() *QCameraControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QCameraControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraControl")
	return gothis
}

/*

 */
type QCameraControl__PropertyChangeType = int

// Indicates the capture mode is changed.
const QCameraControl__CaptureMode QCameraControl__PropertyChangeType = 1

// Image encoder settings are changed, including resolution.
const QCameraControl__ImageEncodingSettings QCameraControl__PropertyChangeType = 2

// Video encoder settings are changed, including audio, video and container settings.
const QCameraControl__VideoEncodingSettings QCameraControl__PropertyChangeType = 3

// Viewfinder is changed.
const QCameraControl__Viewfinder QCameraControl__PropertyChangeType = 4

// Viewfinder settings are changed.
const QCameraControl__ViewfinderSettings QCameraControl__PropertyChangeType = 5

func (this *QCameraControl) PropertyChangeTypeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCameraControl_PropertyChangeTypeItemName(val int) string {
	var nilthis *QCameraControl
	return nilthis.PropertyChangeTypeItemName(val)
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
