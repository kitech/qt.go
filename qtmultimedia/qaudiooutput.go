package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudiooutput.h
// #include <qaudiooutput.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QAudioOutput struct {
	*qtcore.QObject
}
type QAudioOutput_ITF interface {
	qtcore.QObject_ITF
	QAudioOutput_PTR() *QAudioOutput
}

func (ptr *QAudioOutput) QAudioOutput_PTR() *QAudioOutput { return ptr }

func (this *QAudioOutput) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAudioOutput) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQAudioOutputFromPointer(cthis unsafe.Pointer) *QAudioOutput {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAudioOutput{bcthis0}
}
func (*QAudioOutput) NewFromPointer(cthis unsafe.Pointer) *QAudioOutput {
	return NewQAudioOutputFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAudioOutput) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioOutput10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioOutput(const QAudioFormat &, QObject *)

/*
Construct a new audio output and attach it to parent. The default audio output device is used with the output format parameters.
*/
func (*QAudioOutput) NewForInherit(format QAudioFormat_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QAudioOutput {
	return NewQAudioOutput(format, parent)
}
func NewQAudioOutput(format QAudioFormat_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QAudioOutput {
	var convArg0 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg0 = format.QAudioFormat_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutputC2ERK12QAudioFormatP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioOutputFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioOutput")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioOutput(const QAudioFormat &, QObject *)

/*
Construct a new audio output and attach it to parent. The default audio output device is used with the output format parameters.
*/
func (*QAudioOutput) NewForInherit__() *QAudioOutput {
	return NewQAudioOutput__()
}
func NewQAudioOutput__() *QAudioOutput {
	// arg: 0, const QAudioFormat &=LValueReference, QAudioFormat=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutputC2ERK12QAudioFormatP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioOutputFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioOutput")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioOutput(const QAudioFormat &, QObject *)

/*
Construct a new audio output and attach it to parent. The default audio output device is used with the output format parameters.
*/
func (*QAudioOutput) NewForInherit__1(format QAudioFormat_ITF) *QAudioOutput {
	return NewQAudioOutput__1(format)
}
func NewQAudioOutput__1(format QAudioFormat_ITF) *QAudioOutput {
	var convArg0 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg0 = format.QAudioFormat_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutputC2ERK12QAudioFormatP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioOutputFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioOutput")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAudioOutput(const QAudioDeviceInfo &, const QAudioFormat &, QObject *)

/*
Construct a new audio output and attach it to parent. The default audio output device is used with the output format parameters.
*/
func (*QAudioOutput) NewForInherit_1(audioDeviceInfo QAudioDeviceInfo_ITF, format QAudioFormat_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QAudioOutput {
	return NewQAudioOutput_1(audioDeviceInfo, format, parent)
}
func NewQAudioOutput_1(audioDeviceInfo QAudioDeviceInfo_ITF, format QAudioFormat_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QAudioOutput {
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
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutputC2ERK16QAudioDeviceInfoRK12QAudioFormatP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioOutputFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioOutput")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAudioOutput(const QAudioDeviceInfo &, const QAudioFormat &, QObject *)

/*
Construct a new audio output and attach it to parent. The default audio output device is used with the output format parameters.
*/
func (*QAudioOutput) NewForInherit_1_(audioDeviceInfo QAudioDeviceInfo_ITF) *QAudioOutput {
	return NewQAudioOutput_1_(audioDeviceInfo)
}
func NewQAudioOutput_1_(audioDeviceInfo QAudioDeviceInfo_ITF) *QAudioOutput {
	var convArg0 unsafe.Pointer
	if audioDeviceInfo != nil && audioDeviceInfo.QAudioDeviceInfo_PTR() != nil {
		convArg0 = audioDeviceInfo.QAudioDeviceInfo_PTR().GetCthis()
	}
	// arg: 1, const QAudioFormat &=LValueReference, QAudioFormat=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutputC2ERK16QAudioDeviceInfoRK12QAudioFormatP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioOutputFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioOutput")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAudioOutput(const QAudioDeviceInfo &, const QAudioFormat &, QObject *)

/*
Construct a new audio output and attach it to parent. The default audio output device is used with the output format parameters.
*/
func (*QAudioOutput) NewForInherit_1_1(audioDeviceInfo QAudioDeviceInfo_ITF, format QAudioFormat_ITF) *QAudioOutput {
	return NewQAudioOutput_1_1(audioDeviceInfo, format)
}
func NewQAudioOutput_1_1(audioDeviceInfo QAudioDeviceInfo_ITF, format QAudioFormat_ITF) *QAudioOutput {
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
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutputC2ERK16QAudioDeviceInfoRK12QAudioFormatP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioOutputFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioOutput")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAudioOutput()

/*

 */
func DeleteQAudioOutput(this *QAudioOutput) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutputD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QAudioFormat format() const

/*
Returns the QAudioFormat being used.
*/
func (this *QAudioOutput) Format() *QAudioFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioOutput6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioFormat)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(QIODevice *)

/*
Starts transferring audio data from the device to the system's audio output. The device must have been opened in the ReadOnly or ReadWrite modes.

If the QAudioOutput is able to successfully output audio data, state() returns QAudio::ActiveState, error() returns QAudio::NoError and the stateChanged() signal is emitted.

If a problem occurs during this process, error() returns QAudio::OpenError, state() returns QAudio::StoppedState and the stateChanged() signal is emitted.

See also QIODevice.
*/
func (this *QAudioOutput) Start(device qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutput5startEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:72
// index:1
// Public Visibility=Default Availability=Available
// [8] QIODevice * start()

/*
Starts transferring audio data from the device to the system's audio output. The device must have been opened in the ReadOnly or ReadWrite modes.

If the QAudioOutput is able to successfully output audio data, state() returns QAudio::ActiveState, error() returns QAudio::NoError and the stateChanged() signal is emitted.

If a problem occurs during this process, error() returns QAudio::OpenError, state() returns QAudio::StoppedState and the stateChanged() signal is emitted.

See also QIODevice.
*/
func (this *QAudioOutput) Start_1() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutput5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*
Stops the audio output, detaching from the system resource.

Sets error() to QAudio::NoError, state() to QAudio::StoppedState and emit stateChanged() signal.
*/
func (this *QAudioOutput) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutput4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()

/*
Drops all audio data in the buffers, resets buffers to zero.
*/
func (this *QAudioOutput) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutput5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void suspend()

/*
Stops processing audio data, preserving buffered audio data.

Sets error() to QAudio::NoError, state() to QAudio::SuspendedState and emits stateChanged() signal.
*/
func (this *QAudioOutput) Suspend() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutput7suspendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resume()

/*
Resumes processing audio data after a suspend().

Sets error() to QAudio::NoError. Sets state() to QAudio::ActiveState if you previously called start(QIODevice*). Sets state() to QAudio::IdleState if you previously called start(). emits stateChanged() signal.
*/
func (this *QAudioOutput) Resume() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutput6resumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBufferSize(int)

/*
Sets the audio buffer size to value in bytes.

Note: This function can be called anytime before start(). Calls to this are ignored after start(). It should not be assumed that the buffer size set is the actual buffer size used - call bufferSize() anytime after start() to return the actual buffer size being used.

See also bufferSize().
*/
func (this *QAudioOutput) SetBufferSize(bytes int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutput13setBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bytes)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:80
// index:0
// Public Visibility=Default Availability=Available
// [4] int bufferSize() const

/*
Returns the audio buffer size in bytes.

If called before start(), returns platform default value. If called before start() but setBufferSize() was called prior, returns value set by setBufferSize(). If called after start(), returns the actual buffer size being used. This may not be what was set previously by setBufferSize().

See also setBufferSize().
*/
func (this *QAudioOutput) BufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioOutput10bufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int bytesFree() const

/*
Returns the number of free bytes available in the audio buffer.

Note: The returned value is only valid while in QAudio::ActiveState or QAudio::IdleState state, otherwise returns zero.
*/
func (this *QAudioOutput) BytesFree() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioOutput9bytesFreeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] int periodSize() const

/*
Returns the period size in bytes. This is the amount of data required each period to prevent buffer underrun, and to ensure uninterrupted playback.

Note: It is recommended to provide at least enough data for a full period with each write operation.
*/
func (this *QAudioOutput) PeriodSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioOutput10periodSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNotifyInterval(int)

/*
Sets the interval for notify() signal to be emitted. This is based on the ms of audio data processed, not on wall clock time. The minimum resolution of the timer is platform specific and values should be checked with notifyInterval() to confirm the actual value being used.

See also notifyInterval().
*/
func (this *QAudioOutput) SetNotifyInterval(milliSeconds int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutput17setNotifyIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), milliSeconds)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] int notifyInterval() const

/*
Returns the notify interval in milliseconds.

See also setNotifyInterval().
*/
func (this *QAudioOutput) NotifyInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioOutput14notifyIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 processedUSecs() const

/*
Returns the amount of audio data processed since start() was called (in microseconds).
*/
func (this *QAudioOutput) ProcessedUSecs() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioOutput14processedUSecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 elapsedUSecs() const

/*
Returns the microseconds since start() was called, including time in Idle and Suspend states.
*/
func (this *QAudioOutput) ElapsedUSecs() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioOutput12elapsedUSecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:91
// index:0
// Public Visibility=Default Availability=Available
// [4] QAudio::Error error() const

/*
Returns the error state.
*/
func (this *QAudioOutput) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioOutput5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] QAudio::State state() const

/*
Returns the state of audio processing.
*/
func (this *QAudioOutput) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioOutput5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVolume(qreal)

/*
Sets the output volume to volume.

The volume is scaled linearly from 0.0 (silence) to 1.0 (full volume). Values outside this range will be clamped.

The default volume is 1.0.

Note: Adjustments to the volume will change the volume of this audio stream, not the global volume.

UI volume controls should usually be scaled nonlinearly. For example, using a logarithmic scale will produce linear changes in perceived loudness, which is what a user would normally expect from a volume control. See QAudio::convertVolume() for more details.

See also volume().
*/
func (this *QAudioOutput) SetVolume(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutput9setVolumeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal volume() const

/*
Returns the volume between 0.0 and 1.0 inclusive.

See also setVolume().
*/
func (this *QAudioOutput) Volume() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioOutput6volumeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QString category() const

/*
Returns the audio category of this audio stream.

Some platforms can group audio streams into categories and manage their volumes independently, or display them in a system mixer control. You can set this property to allow the platform to distinguish the purpose of your streams.

See also setCategory().
*/
func (this *QAudioOutput) Category() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioOutput8categoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCategory(const QString &)

/*
Sets the audio category of this audio stream to category.

Some platforms can group audio streams into categories and manage their volumes independently, or display them in a system mixer control. You can set this property to allow the platform to distinguish the purpose of your streams.

Not all platforms support audio stream categorization. In this case, the function call will be ignored.

Changing an audio output stream's category while it is opened will not take effect until it is reopened.

See also category().
*/
func (this *QAudioOutput) SetCategory(category string) {
	var tmpArg0 = qtcore.NewQString_5(category)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutput11setCategoryERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QAudio::State)

/*
This signal is emitted when the device state has changed. This is the current state of the audio output.
*/
func (this *QAudioOutput) StateChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutput12stateChangedEN6QAudio5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiooutput.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void notify()

/*
This signal is emitted when a certain interval of milliseconds of audio data has been processed. The interval is set by setNotifyInterval().
*/
func (this *QAudioOutput) Notify() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioOutput6notifyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
