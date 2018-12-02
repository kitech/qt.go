package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h
// #include <qaudiodecodercontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QAudioDecoderControl struct {
	*QMediaControl
}
type QAudioDecoderControl_ITF interface {
	QMediaControl_ITF
	QAudioDecoderControl_PTR() *QAudioDecoderControl
}

func (ptr *QAudioDecoderControl) QAudioDecoderControl_PTR() *QAudioDecoderControl { return ptr }

func (this *QAudioDecoderControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QAudioDecoderControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQAudioDecoderControlFromPointer(cthis unsafe.Pointer) *QAudioDecoderControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QAudioDecoderControl{bcthis0}
}
func (*QAudioDecoderControl) NewFromPointer(cthis unsafe.Pointer) *QAudioDecoderControl {
	return NewQAudioDecoderControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAudioDecoderControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAudioDecoderControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAudioDecoderControl()

/*

 */
func DeleteQAudioDecoderControl(this *QAudioDecoderControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QAudioDecoder::State state() const

/*
Returns the state of a player control.
*/
func (this *QAudioDecoderControl) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAudioDecoderControl5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:62
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString sourceFilename() const

/*
Returns the current media source filename, or a null QString if none (or a device)

See also setSourceFilename().
*/
func (this *QAudioDecoderControl) SourceFilename() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAudioDecoderControl14sourceFilenameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setSourceFilename(const QString &)

/*
Sets the current source to fileName. Changing the source will stop any current decoding and discard any buffers.

Sources are exclusive, so only one can be set.

See also sourceFilename().
*/
func (this *QAudioDecoderControl) SetSourceFilename(fileName string) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControl17setSourceFilenameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:65
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QIODevice * sourceDevice() const

/*
Returns the current media source QIODevice, or 0 if none (or a file).

See also setSourceDevice().
*/
func (this *QAudioDecoderControl) SourceDevice() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAudioDecoderControl12sourceDeviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setSourceDevice(QIODevice *)

/*
Sets the current source to device. Changing the source will stop any current decoding and discard any buffers.

Sources are exclusive, so only one can be set.

See also sourceDevice().
*/
func (this *QAudioDecoderControl) SetSourceDevice(device qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControl15setSourceDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:68
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void start()

/*
Starts decoding the current media.

If successful the player control will immediately enter the decoding state.

See also state() and read().
*/
func (this *QAudioDecoderControl) Start() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControl5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:69
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void stop()

/*
Stops playback of the current media and discards any buffers.

If successful the player control will immediately enter the stopped state.
*/
func (this *QAudioDecoderControl) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControl4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:71
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAudioFormat audioFormat() const

/*
Returns the current audio format of the decoded stream.

Any buffers returned should have this format.

See also setAudioFormat() and formatChanged().
*/
func (this *QAudioDecoderControl) AudioFormat() *QAudioFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAudioDecoderControl11audioFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioFormat)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:72
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setAudioFormat(const QAudioFormat &)

/*
Set the desired audio format for decoded samples to format.

If the decoder does not support this format, error() will be set to FormatError.

If you do not specify a format, the format of the decoded audio itself will be used. Otherwise, some format conversion will be applied.

If you wish to reset the decoded format to that of the original audio file, you can specify an invalid format.

See also audioFormat().
*/
func (this *QAudioDecoderControl) SetAudioFormat(format QAudioFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg0 = format.QAudioFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControl14setAudioFormatERK12QAudioFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:74
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAudioBuffer read()

/*
Attempts to read a buffer from the decoder, without blocking. Returns invalid buffer if there are no decoded buffers available, or on error.
*/
func (this *QAudioDecoderControl) Read() *QAudioBuffer /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControl4readEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioBufferFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioBuffer)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:75
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool bufferAvailable() const

/*
Returns true if a buffer is available to be read, and false otherwise.
*/
func (this *QAudioDecoderControl) BufferAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAudioDecoderControl15bufferAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:77
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qint64 position() const

/*
Returns position (in milliseconds) of the last buffer read from the decoder or -1 if no buffers have been read.
*/
func (this *QAudioDecoderControl) Position() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAudioDecoderControl8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:78
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qint64 duration() const

/*
Returns total duration (in milliseconds) of the audio stream or -1 if not available.
*/
func (this *QAudioDecoderControl) Duration() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QAudioDecoderControl8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QAudioDecoder::State)

/*
Signals that the state of a player control has changed.

See also state().
*/
func (this *QAudioDecoderControl) StateChanged(newState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControl12stateChangedEN13QAudioDecoder5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void formatChanged(const QAudioFormat &)

/*
Signals that the current audio format of the decoder has changed to format.

See also audioFormat() and setAudioFormat().
*/
func (this *QAudioDecoderControl) FormatChanged(format QAudioFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg0 = format.QAudioFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControl13formatChangedERK12QAudioFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sourceChanged()

/*
Signals that the current source of the decoder has changed.

See also sourceFilename() and sourceDevice().
*/
func (this *QAudioDecoderControl) SourceChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControl13sourceChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void error(int, const QString &)

/*
Signals that an error has occurred. The errorString provides a more detailed explanation.
*/
func (this *QAudioDecoderControl) Error(error int, errorString string) {
	var tmpArg1 = qtcore.NewQString_5(errorString)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControl5errorEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void bufferReady()

/*
Signals that a new buffer is ready for reading.
*/
func (this *QAudioDecoderControl) BufferReady() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControl11bufferReadyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void bufferAvailableChanged(bool)

/*
Signals that the bufferAvailable property has changed to available.
*/
func (this *QAudioDecoderControl) BufferAvailableChanged(available bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControl22bufferAvailableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), available)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished()

/*
Signals that the decoding has finished successfully. If decoding fails, error signal is emitted instead.

See also start(), stop(), and error().
*/
func (this *QAudioDecoderControl) Finished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControl8finishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void positionChanged(qint64)

/*
Signals that the current position of the decoder has changed.

See also durationChanged().
*/
func (this *QAudioDecoderControl) PositionChanged(position int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControl15positionChangedEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void durationChanged(qint64)

/*
Signals that the estimated duration of the decoded data has changed.

See also positionChanged().
*/
func (this *QAudioDecoderControl) DurationChanged(duration int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControl15durationChangedEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), duration)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:95
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAudioDecoderControl(QObject *)

/*
Constructs a new audio decoder control with the given parent.
*/
func (*QAudioDecoderControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioDecoderControl {
	return NewQAudioDecoderControl(parent)
}
func NewQAudioDecoderControl(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioDecoderControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioDecoderControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioDecoderControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiodecodercontrol.h:95
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAudioDecoderControl(QObject *)

/*
Constructs a new audio decoder control with the given parent.
*/
func (*QAudioDecoderControl) NewForInherit__() *QAudioDecoderControl {
	return NewQAudioDecoderControl__()
}
func NewQAudioDecoderControl__() *QAudioDecoderControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN20QAudioDecoderControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioDecoderControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioDecoderControl")
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
