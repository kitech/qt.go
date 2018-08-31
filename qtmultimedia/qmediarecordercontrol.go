package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h
// #include <qmediarecordercontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
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
type QMediaRecorderControl struct {
	*QMediaControl
}
type QMediaRecorderControl_ITF interface {
	QMediaControl_ITF
	QMediaRecorderControl_PTR() *QMediaRecorderControl
}

func (ptr *QMediaRecorderControl) QMediaRecorderControl_PTR() *QMediaRecorderControl { return ptr }

func (this *QMediaRecorderControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QMediaRecorderControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQMediaRecorderControlFromPointer(cthis unsafe.Pointer) *QMediaRecorderControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QMediaRecorderControl{bcthis0}
}
func (*QMediaRecorderControl) NewFromPointer(cthis unsafe.Pointer) *QMediaRecorderControl {
	return NewQMediaRecorderControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaRecorderControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QMediaRecorderControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaRecorderControl()

/*

 */
func DeleteQMediaRecorderControl(this *QMediaRecorderControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMediaRecorderControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QUrl outputLocation() const

/*
Returns the current output location being used.

See also setOutputLocation().
*/
func (this *QMediaRecorderControl) OutputLocation() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QMediaRecorderControl14outputLocationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:64
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool setOutputLocation(const QUrl &)

/*
Sets the output location and returns if this operation is successful. If file at the output location already exists, it should be overwritten.

The location can be relative or empty; in this case the service should use the system specific place and file naming scheme.

After recording has started, the backend should report the actual file location with actualLocationChanged() signal.

See also outputLocation().
*/
func (this *QMediaRecorderControl) SetOutputLocation(location qtcore.QUrl_ITF) bool {
	var convArg0 unsafe.Pointer
	if location != nil && location.QUrl_PTR() != nil {
		convArg0 = location.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMediaRecorderControl17setOutputLocationERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QMediaRecorder::State state() const

/*
Return the current recording state.

See also setState().
*/
func (this *QMediaRecorderControl) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QMediaRecorderControl5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:67
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QMediaRecorder::Status status() const

/*
Return the current recording status.
*/
func (this *QMediaRecorderControl) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QMediaRecorderControl6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:69
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qint64 duration() const

/*
Return the current duration in milliseconds.
*/
func (this *QMediaRecorderControl) Duration() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QMediaRecorderControl8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:71
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isMuted() const

/*
Returns true if the recorder is muted, and false if it is not.
*/
func (this *QMediaRecorderControl) IsMuted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QMediaRecorderControl7isMutedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:72
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qreal volume() const

/*
Returns the audio volume of a media recorder control.

See also setVolume().
*/
func (this *QMediaRecorderControl) Volume() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QMediaRecorderControl6volumeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:74
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void applySettings()

/*
Commits the encoder settings and performs pre-initialization to reduce delays when recording is started.
*/
func (this *QMediaRecorderControl) ApplySettings() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMediaRecorderControl13applySettingsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QMediaRecorder::State)

/*
Signals that the state of a media recorder has changed.
*/
func (this *QMediaRecorderControl) StateChanged(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMediaRecorderControl12stateChangedEN14QMediaRecorder5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void statusChanged(QMediaRecorder::Status)

/*
Signals that the status of a media recorder has changed.
*/
func (this *QMediaRecorderControl) StatusChanged(status int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMediaRecorderControl13statusChangedEN14QMediaRecorder6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), status)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void durationChanged(qint64)

/*
Signals that the duration of the recorded media has changed.

This only emitted when there is a discontinuous change in the duration such as being reset to 0.
*/
func (this *QMediaRecorderControl) DurationChanged(position int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMediaRecorderControl15durationChangedEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mutedChanged(bool)

/*
Signals that the muted state of a media recorder has changed.
*/
func (this *QMediaRecorderControl) MutedChanged(muted bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMediaRecorderControl12mutedChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), muted)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void volumeChanged(qreal)

/*
Signals that the audio gain value has changed.
*/
func (this *QMediaRecorderControl) VolumeChanged(volume float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMediaRecorderControl13volumeChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), volume)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void actualLocationChanged(const QUrl &)

/*
Signals that the actual media location has changed. This signal should be emitted at start of recording.
*/
func (this *QMediaRecorderControl) ActualLocationChanged(location qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if location != nil && location.QUrl_PTR() != nil {
		convArg0 = location.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMediaRecorderControl21actualLocationChangedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void error(int, const QString &)

/*
Signals that an error has occurred. The errorString describes the error.
*/
func (this *QMediaRecorderControl) Error(error int, errorString string) {
	var tmpArg1 = qtcore.NewQString_5(errorString)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMediaRecorderControl5errorEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:86
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setState(QMediaRecorder::State)

/*
Set the media recorder state.

See also state().
*/
func (this *QMediaRecorderControl) SetState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMediaRecorderControl8setStateEN14QMediaRecorder5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:87
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setMuted(bool)

/*
Sets the muted state of a media recorder.

See also isMuted().
*/
func (this *QMediaRecorderControl) SetMuted(muted bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMediaRecorderControl8setMutedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), muted)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:88
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setVolume(qreal)

/*
Sets the audio volume of a media recorder control.

The volume is scaled linearly, ranging from 0 (silence) to 100 (full volume).

See also volume().
*/
func (this *QMediaRecorderControl) SetVolume(volume float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMediaRecorderControl9setVolumeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), volume)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:91
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaRecorderControl(QObject *)

/*
Constructs a media recorder control with the given parent.
*/
func (*QMediaRecorderControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaRecorderControl {
	return NewQMediaRecorderControl(parent)
}
func NewQMediaRecorderControl(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaRecorderControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMediaRecorderControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaRecorderControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaRecorderControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediarecordercontrol.h:91
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaRecorderControl(QObject *)

/*
Constructs a media recorder control with the given parent.
*/
func (*QMediaRecorderControl) NewForInherit__() *QMediaRecorderControl {
	return NewQMediaRecorderControl__()
}
func NewQMediaRecorderControl__() *QMediaRecorderControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMediaRecorderControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaRecorderControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaRecorderControl")
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
