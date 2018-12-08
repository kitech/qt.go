package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudioinput.h
// #include <qaudioinput.h>
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
type QAudioInput struct {
	*qtcore.QObject
}
type QAudioInput_ITF interface {
	qtcore.QObject_ITF
	QAudioInput_PTR() *QAudioInput
}

func (ptr *QAudioInput) QAudioInput_PTR() *QAudioInput { return ptr }

func (this *QAudioInput) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAudioInput) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQAudioInputFromPointer(cthis unsafe.Pointer) *QAudioInput {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAudioInput{bcthis0}
}
func (*QAudioInput) NewFromPointer(cthis unsafe.Pointer) *QAudioInput {
	return NewQAudioInputFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAudioInput) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QAudioInput10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioInput(const QAudioFormat &, QObject *)

/*
Construct a new audio input and attach it to parent. The default audio input device is used with the output format parameters.
*/
func (*QAudioInput) NewForInherit(format QAudioFormat_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QAudioInput {
	return NewQAudioInput(format, parent)
}
func NewQAudioInput(format QAudioFormat_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QAudioInput {
	var convArg0 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg0 = format.QAudioFormat_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInputC2ERK12QAudioFormatP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioInputFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioInput")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioInput(const QAudioFormat &, QObject *)

/*
Construct a new audio input and attach it to parent. The default audio input device is used with the output format parameters.
*/
func (*QAudioInput) NewForInheritp() *QAudioInput {
	return NewQAudioInputp()
}
func NewQAudioInputp() *QAudioInput {
	// arg: 0, const QAudioFormat &=LValueReference, QAudioFormat=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInputC2ERK12QAudioFormatP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioInputFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioInput")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioInput(const QAudioFormat &, QObject *)

/*
Construct a new audio input and attach it to parent. The default audio input device is used with the output format parameters.
*/
func (*QAudioInput) NewForInheritp1(format QAudioFormat_ITF) *QAudioInput {
	return NewQAudioInputp1(format)
}
func NewQAudioInputp1(format QAudioFormat_ITF) *QAudioInput {
	var convArg0 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg0 = format.QAudioFormat_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInputC2ERK12QAudioFormatP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioInputFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioInput")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAudioInput(const QAudioDeviceInfo &, const QAudioFormat &, QObject *)

/*
Construct a new audio input and attach it to parent. The default audio input device is used with the output format parameters.
*/
func (*QAudioInput) NewForInherit1(audioDeviceInfo QAudioDeviceInfo_ITF, format QAudioFormat_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QAudioInput {
	return NewQAudioInput1(audioDeviceInfo, format, parent)
}
func NewQAudioInput1(audioDeviceInfo QAudioDeviceInfo_ITF, format QAudioFormat_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QAudioInput {
	var convArg0 unsafe.Pointer
	if audioDeviceInfo != nil && audioDeviceInfo.QAudioDeviceInfo_PTR() != nil {
		convArg0 = audioDeviceInfo.QAudioDeviceInfo_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg1 = format.QAudioFormat_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg2 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInputC2ERK16QAudioDeviceInfoRK12QAudioFormatP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioInputFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioInput")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAudioInput(const QAudioDeviceInfo &, const QAudioFormat &, QObject *)

/*
Construct a new audio input and attach it to parent. The default audio input device is used with the output format parameters.
*/
func (*QAudioInput) NewForInherit1p(audioDeviceInfo QAudioDeviceInfo_ITF) *QAudioInput {
	return NewQAudioInput1p(audioDeviceInfo)
}
func NewQAudioInput1p(audioDeviceInfo QAudioDeviceInfo_ITF) *QAudioInput {
	var convArg0 unsafe.Pointer
	if audioDeviceInfo != nil && audioDeviceInfo.QAudioDeviceInfo_PTR() != nil {
		convArg0 = audioDeviceInfo.QAudioDeviceInfo_PTR().GetCthis()
	}
	// arg: 1, const QAudioFormat &=LValueReference, QAudioFormat=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInputC2ERK16QAudioDeviceInfoRK12QAudioFormatP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioInputFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioInput")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAudioInput(const QAudioDeviceInfo &, const QAudioFormat &, QObject *)

/*
Construct a new audio input and attach it to parent. The default audio input device is used with the output format parameters.
*/
func (*QAudioInput) NewForInherit1p1(audioDeviceInfo QAudioDeviceInfo_ITF, format QAudioFormat_ITF) *QAudioInput {
	return NewQAudioInput1p1(audioDeviceInfo, format)
}
func NewQAudioInput1p1(audioDeviceInfo QAudioDeviceInfo_ITF, format QAudioFormat_ITF) *QAudioInput {
	var convArg0 unsafe.Pointer
	if audioDeviceInfo != nil && audioDeviceInfo.QAudioDeviceInfo_PTR() != nil {
		convArg0 = audioDeviceInfo.QAudioDeviceInfo_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg1 = format.QAudioFormat_PTR().GetCthis()
	}
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInputC2ERK16QAudioDeviceInfoRK12QAudioFormatP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioInputFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioInput")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAudioInput()

/*

 */
func DeleteQAudioInput(this *QAudioInput) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInputD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QAudioFormat format() const

/*
Returns the QAudioFormat being used.
*/
func (this *QAudioInput) Format() *QAudioFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QAudioInput6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioFormat)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(QIODevice *)

/*
Starts transferring audio data from the system's audio input to the device. The device must have been opened in the WriteOnly, Append or ReadWrite modes.

If the QAudioInput is able to successfully get audio data, state() returns either QAudio::ActiveState or QAudio::IdleState, error() returns QAudio::NoError and the stateChanged() signal is emitted.

If a problem occurs during this process, error() returns QAudio::OpenError, state() returns QAudio::StoppedState and the stateChanged() signal is emitted.

See also QIODevice.
*/
func (this *QAudioInput) Start(device qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInput5startEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:72
// index:1
// Public Visibility=Default Availability=Available
// [8] QIODevice * start()

/*
Starts transferring audio data from the system's audio input to the device. The device must have been opened in the WriteOnly, Append or ReadWrite modes.

If the QAudioInput is able to successfully get audio data, state() returns either QAudio::ActiveState or QAudio::IdleState, error() returns QAudio::NoError and the stateChanged() signal is emitted.

If a problem occurs during this process, error() returns QAudio::OpenError, state() returns QAudio::StoppedState and the stateChanged() signal is emitted.

See also QIODevice.
*/
func (this *QAudioInput) Start1() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInput5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*
Stops the audio input, detaching from the system resource.

Sets error() to QAudio::NoError, state() to QAudio::StoppedState and emit stateChanged() signal.
*/
func (this *QAudioInput) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInput4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()

/*
Drops all audio data in the buffers, resets buffers to zero.
*/
func (this *QAudioInput) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInput5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void suspend()

/*
Stops processing audio data, preserving buffered audio data.

Sets error() to QAudio::NoError, state() to QAudio::SuspendedState and emit stateChanged() signal.
*/
func (this *QAudioInput) Suspend() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInput7suspendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resume()

/*
Resumes processing audio data after a suspend().

Sets error() to QAudio::NoError. Sets state() to QAudio::ActiveState if you previously called start(QIODevice*). Sets state() to QAudio::IdleState if you previously called start(). emits stateChanged() signal.
*/
func (this *QAudioInput) Resume() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInput6resumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBufferSize(int)

/*
Sets the audio buffer size to value bytes.

Note: This function can be called anytime before start(), calls to this are ignored after start(). It should not be assumed that the buffer size set is the actual buffer size used, calling bufferSize() anytime after start() will return the actual buffer size being used.

See also bufferSize().
*/
func (this *QAudioInput) SetBufferSize(bytes int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInput13setBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bytes)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:80
// index:0
// Public Visibility=Default Availability=Available
// [4] int bufferSize() const

/*
Returns the audio buffer size in bytes.

If called before start(), returns platform default value. If called before start() but setBufferSize() was called prior, returns value set by setBufferSize(). If called after start(), returns the actual buffer size being used. This may not be what was set previously by setBufferSize().

See also setBufferSize().
*/
func (this *QAudioInput) BufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QAudioInput10bufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int bytesReady() const

/*
Returns the amount of audio data available to read in bytes.

Note: returned value is only valid while in QAudio::ActiveState or QAudio::IdleState state, otherwise returns zero.
*/
func (this *QAudioInput) BytesReady() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QAudioInput10bytesReadyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] int periodSize() const

/*
Returns the period size in bytes.

Note: This is the recommended read size in bytes.
*/
func (this *QAudioInput) PeriodSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QAudioInput10periodSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNotifyInterval(int)

/*
Sets the interval for notify() signal to be emitted. This is based on the ms of audio data processed not on actual real-time. The minimum resolution of the timer is platform specific and values should be checked with notifyInterval() to confirm actual value being used.

See also notifyInterval().
*/
func (this *QAudioInput) SetNotifyInterval(milliSeconds int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInput17setNotifyIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), milliSeconds)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] int notifyInterval() const

/*
Returns the notify interval in milliseconds.

See also setNotifyInterval().
*/
func (this *QAudioInput) NotifyInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QAudioInput14notifyIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVolume(qreal)

/*
Sets the input volume to volume.

The volume is scaled linearly from 0.0 (silence) to 1.0 (full volume). Values outside this range will be clamped.

If the device does not support adjusting the input volume then volume will be ignored and the input volume will remain at 1.0.

The default volume is 1.0.

Note: Adjustments to the volume will change the volume of this audio stream, not the global volume.

See also volume().
*/
func (this *QAudioInput) SetVolume(volume float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInput9setVolumeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), volume)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal volume() const

/*
Returns the input volume.

If the device does not support adjusting the input volume the returned value will be 1.0.

See also setVolume().
*/
func (this *QAudioInput) Volume() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QAudioInput6volumeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 processedUSecs() const

/*
Returns the amount of audio data processed since start() was called in microseconds.
*/
func (this *QAudioInput) ProcessedUSecs() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QAudioInput14processedUSecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 elapsedUSecs() const

/*
Returns the microseconds since start() was called, including time in Idle and Suspend states.
*/
func (this *QAudioInput) ElapsedUSecs() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QAudioInput12elapsedUSecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] QAudio::Error error() const

/*
Returns the error state.
*/
func (this *QAudioInput) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QAudioInput5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] QAudio::State state() const

/*
Returns the state of audio processing.
*/
func (this *QAudioInput) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QAudioInput5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QAudio::State)

/*
This signal is emitted when the device state has changed.
*/
func (this *QAudioInput) StateChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInput12stateChangedEN6QAudio5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioinput.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void notify()

/*
This signal is emitted when x ms of audio data has been processed the interval set by setNotifyInterval(x).
*/
func (this *QAudioInput) Notify() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAudioInput6notifyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
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
