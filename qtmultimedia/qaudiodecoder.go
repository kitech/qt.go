package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudiodecoder.h
// #include <qaudiodecoder.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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
type QAudioDecoder struct {
	*QMediaObject
}
type QAudioDecoder_ITF interface {
	QMediaObject_ITF
	QAudioDecoder_PTR() *QAudioDecoder
}

func (ptr *QAudioDecoder) QAudioDecoder_PTR() *QAudioDecoder { return ptr }

func (this *QAudioDecoder) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaObject.GetCthis()
	}
}
func (this *QAudioDecoder) SetCthis(cthis unsafe.Pointer) {
	this.QMediaObject = NewQMediaObjectFromPointer(cthis)
}
func NewQAudioDecoderFromPointer(cthis unsafe.Pointer) *QAudioDecoder {
	bcthis0 := NewQMediaObjectFromPointer(cthis)
	return &QAudioDecoder{bcthis0}
}
func (*QAudioDecoder) NewFromPointer(cthis unsafe.Pointer) *QAudioDecoder {
	return NewQAudioDecoderFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAudioDecoder) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QAudioDecoder10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioDecoder(QObject *)

/*
Construct an QAudioDecoder instance parented to parent.
*/
func NewQAudioDecoder(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioDecoder {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoderC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioDecoderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioDecoder")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioDecoder(QObject *)

/*
Construct an QAudioDecoder instance parented to parent.
*/
func NewQAudioDecoder__() *QAudioDecoder {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoderC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioDecoderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioDecoder")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAudioDecoder()

/*

 */
func DeleteQAudioDecoder(this *QAudioDecoder) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:81
// index:0
// Public static Visibility=Default Availability=Available
// [4] QMultimedia::SupportEstimate hasSupport(const QString &, const QStringList &)

/*
Returns the level of support an audio decoder has for a mimeType and a set of codecs.
*/
func (this *QAudioDecoder) HasSupport(mimeType string, codecs qtcore.QStringList_ITF) int {
	var tmpArg0 = qtcore.NewQString_5(mimeType)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if codecs != nil && codecs.QStringList_PTR() != nil {
		convArg1 = codecs.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder10hasSupportERK7QStringRK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QAudioDecoder_HasSupport(mimeType string, codecs qtcore.QStringList_ITF) int {
	var nilthis *QAudioDecoder
	rv := nilthis.HasSupport(mimeType, codecs)
	return rv
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:81
// index:0
// Public static Visibility=Default Availability=Available
// [4] QMultimedia::SupportEstimate hasSupport(const QString &, const QStringList &)

/*
Returns the level of support an audio decoder has for a mimeType and a set of codecs.
*/
func (this *QAudioDecoder) HasSupport__(mimeType string) int {
	var tmpArg0 = qtcore.NewQString_5(mimeType)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder10hasSupportERK7QStringRK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] QAudioDecoder::State state() const

/*

 */
func (this *QAudioDecoder) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QAudioDecoder5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sourceFilename() const

/*
Returns the current file name to decode. If setSourceDevice was called, this will be empty.

Note: Getter function for property sourceFilename.

See also setSourceFilename().
*/
func (this *QAudioDecoder) SourceFilename() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QAudioDecoder14sourceFilenameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSourceFilename(const QString &)

/*
Sets the current audio file name to fileName.

When this property is set any current decoding is stopped, and any audio buffers are discarded.

You can only specify either a source filename or a source QIODevice. Setting one will unset the other.

Note: Setter function for property sourceFilename.

See also sourceFilename().
*/
func (this *QAudioDecoder) SetSourceFilename(fileName string) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder17setSourceFilenameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * sourceDevice() const

/*
Returns the current source QIODevice, if one was set. If setSourceFilename() was called, this will be 0.

See also setSourceDevice().
*/
func (this *QAudioDecoder) SourceDevice() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QAudioDecoder12sourceDeviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSourceDevice(QIODevice *)

/*
Sets the current audio QIODevice to device.

When this property is set any current decoding is stopped, and any audio buffers are discarded.

You can only specify either a source filename or a source QIODevice. Setting one will unset the other.

See also sourceDevice().
*/
func (this *QAudioDecoder) SetSourceDevice(device qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder15setSourceDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QAudioFormat audioFormat() const

/*
Returns the current audio format of the decoded stream.

Any buffers returned should have this format.

See also setAudioFormat() and formatChanged().
*/
func (this *QAudioDecoder) AudioFormat() *QAudioFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QAudioDecoder11audioFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioFormat)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAudioFormat(const QAudioFormat &)

/*
Set the desired audio format for decoded samples to format.

This property can only be set while the decoder is stopped. Setting this property at other times will be ignored.

If the decoder does not support this format, error() will be set to FormatError.

If you do not specify a format, the format of the decoded audio itself will be used. Otherwise, some format conversion will be applied.

If you wish to reset the decoded format to that of the original audio file, you can specify an invalid format.

See also audioFormat().
*/
func (this *QAudioDecoder) SetAudioFormat(format QAudioFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg0 = format.QAudioFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder14setAudioFormatERK12QAudioFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] QAudioDecoder::Error error() const

/*
Returns the current error state.
*/
func (this *QAudioDecoder) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QAudioDecoder5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:115
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QAudioDecoder::Error)

/*
Returns the current error state.
*/
func (this *QAudioDecoder) Error_1(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder5errorENS_5ErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*

 */
func (this *QAudioDecoder) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QAudioDecoder11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QAudioBuffer read() const

/*
Read a buffer from the decoder, if one is available. Returns an invalid buffer if there are no decoded buffers currently available, or on failure. In both cases this function will not block.

You should either respond to the bufferReady() signal or check the bufferAvailable() function before calling read() to make sure you get useful data.
*/
func (this *QAudioDecoder) Read() *QAudioBuffer /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QAudioDecoder4readEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioBufferFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioBuffer)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool bufferAvailable() const

/*
Returns true if a buffer is available to be read, and false otherwise. If there is no buffer available, calling the read() function will return an invalid buffer.

Note: Getter function for property bufferAvailable.
*/
func (this *QAudioDecoder) BufferAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QAudioDecoder15bufferAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 position() const

/*
Returns position (in milliseconds) of the last buffer read from the decoder or -1 if no buffers have been read.
*/
func (this *QAudioDecoder) Position() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QAudioDecoder8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 duration() const

/*
Returns total duration (in milliseconds) of the audio stream or -1 if not available.
*/
func (this *QAudioDecoder) Duration() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QAudioDecoder8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start()

/*
Starts decoding the audio resource.

As data gets decoded, the bufferReady() signal will be emitted when enough data has been decoded. Calling read() will then return an audio buffer without blocking.

If you call read() before a buffer is ready, an invalid buffer will be returned, again without blocking.

See also read().
*/
func (this *QAudioDecoder) Start() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*
Stop decoding audio. Calling start() again will resume decoding from the beginning.
*/
func (this *QAudioDecoder) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void bufferAvailableChanged(bool)

/*
Signals the availability (if available is true) of a new buffer.

If available is false, there are no buffers available.

Note: Notifier signal for property bufferAvailable.

See also bufferAvailable() and bufferReady().
*/
func (this *QAudioDecoder) BufferAvailableChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder22bufferAvailableChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void bufferReady()

/*
Signals that a new decoded audio buffer is available to be read.

See also read() and bufferAvailable().
*/
func (this *QAudioDecoder) BufferReady() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder11bufferReadyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished()

/*
Signals that the decoding has finished successfully. If decoding fails, error signal is emitted instead.

See also start(), stop(), and error().
*/
func (this *QAudioDecoder) Finished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder8finishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QAudioDecoder::State)

/*
Signal the state of the decoder object has changed.

Note: Notifier signal for property state.
*/
func (this *QAudioDecoder) StateChanged(newState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder12stateChangedENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void formatChanged(const QAudioFormat &)

/*
Signals that the current audio format of the decoder has changed to format.

See also audioFormat() and setAudioFormat().
*/
func (this *QAudioDecoder) FormatChanged(format QAudioFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg0 = format.QAudioFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder13formatChangedERK12QAudioFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sourceChanged()

/*
Signals that the current source of the decoder has changed.

Note: Notifier signal for property sourceFilename.

See also sourceFilename() and sourceDevice().
*/
func (this *QAudioDecoder) SourceChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder13sourceChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void positionChanged(qint64)

/*
Signals that the current position of the decoder has changed.

See also durationChanged().
*/
func (this *QAudioDecoder) PositionChanged(position int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder15positionChangedEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void durationChanged(qint64)

/*
Signals that the estimated duration of the decoded data has changed.

See also positionChanged().
*/
func (this *QAudioDecoder) DurationChanged(duration int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder15durationChangedEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), duration)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:123
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool bind(QObject *)

/*

 */
func (this *QAudioDecoder) Bind(arg0 qtcore.QObject_ITF /*777 QObject **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder4bindEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudiodecoder.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void unbind(QObject *)

/*

 */
func (this *QAudioDecoder) Unbind(arg0 qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QAudioDecoder6unbindEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
Defines the current state of a media player.


*/
type QAudioDecoder__State = int

// The decoder is not decoding. Decoding will start at the start of the media.
const QAudioDecoder__StoppedState QAudioDecoder__State = 0

// The audio player is currently decoding media.
const QAudioDecoder__DecodingState QAudioDecoder__State = 1

func (this *QAudioDecoder) StateItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAudioDecoder_StateItemName(val int) string {
	var nilthis *QAudioDecoder
	return nilthis.StateItemName(val)
}

/*
Defines a media player error condition.


*/
type QAudioDecoder__Error = int

// No error has occurred.
const QAudioDecoder__NoError QAudioDecoder__Error = 0

// A media resource couldn't be resolved.
const QAudioDecoder__ResourceError QAudioDecoder__Error = 1

// The format of a media resource isn't supported.
const QAudioDecoder__FormatError QAudioDecoder__Error = 2

// There are not the appropriate permissions to play a media resource.
const QAudioDecoder__AccessDeniedError QAudioDecoder__Error = 3

// A valid playback service was not found, playback cannot proceed.
const QAudioDecoder__ServiceMissingError QAudioDecoder__Error = 4

func (this *QAudioDecoder) ErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAudioDecoder_ErrorItemName(val int) string {
	var nilthis *QAudioDecoder
	return nilthis.ErrorItemName(val)
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
