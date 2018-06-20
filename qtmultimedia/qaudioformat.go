package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudioformat.h
// #include <qaudioformat.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QAudioFormat struct {
	*qtrt.CObject
}
type QAudioFormat_ITF interface {
	QAudioFormat_PTR() *QAudioFormat
}

func (ptr *QAudioFormat) QAudioFormat_PTR() *QAudioFormat { return ptr }

func (this *QAudioFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAudioFormat) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAudioFormatFromPointer(cthis unsafe.Pointer) *QAudioFormat {
	return &QAudioFormat{&qtrt.CObject{cthis}}
}
func (*QAudioFormat) NewFromPointer(cthis unsafe.Pointer) *QAudioFormat {
	return NewQAudioFormatFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioFormat()

/*
Construct a new audio format.

Values are initialized as follows:


sampleRate() = -1
channelCount() = -1
sampleSize() = -1
byteOrder() = QAudioFormat::Endian(QSysInfo::ByteOrder)
sampleType() = QAudioFormat::Unknown codec() = ""
*/
func NewQAudioFormat() *QAudioFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioFormatC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAudioFormat)
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QAudioFormat()

/*

 */
func DeleteQAudioFormat(this *QAudioFormat) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioFormatD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] QAudioFormat & operator=(const QAudioFormat &)

/*

 */
func (this *QAudioFormat) Operator_equal(other QAudioFormat_ITF) *QAudioFormat {
	var convArg0 unsafe.Pointer
	if other != nil && other.QAudioFormat_PTR() != nil {
		convArg0 = other.QAudioFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioFormataSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioFormat)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QAudioFormat &) const

/*

 */
func (this *QAudioFormat) Operator_equal_equal(other QAudioFormat_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QAudioFormat_PTR() != nil {
		convArg0 = other.QAudioFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormateqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:66
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QAudioFormat &) const

/*

 */
func (this *QAudioFormat) Operator_not_equal(other QAudioFormat_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QAudioFormat_PTR() != nil {
		convArg0 = other.QAudioFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormatneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if all of the parameters are valid.
*/
func (this *QAudioFormat) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormat7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSampleRate(int)

/*
Sets the sample rate to samplerate Hertz.

See also sampleRate().
*/
func (this *QAudioFormat) SetSampleRate(sampleRate int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioFormat13setSampleRateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sampleRate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:71
// index:0
// Public Visibility=Default Availability=Available
// [4] int sampleRate() const

/*
Returns the current sample rate in Hertz.

See also setSampleRate().
*/
func (this *QAudioFormat) SampleRate() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormat10sampleRateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChannelCount(int)

/*
Sets the channel count to channels.

See also channelCount().
*/
func (this *QAudioFormat) SetChannelCount(channelCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioFormat15setChannelCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), channelCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] int channelCount() const

/*
Returns the current channel count value.

See also setChannelCount().
*/
func (this *QAudioFormat) ChannelCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormat12channelCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSampleSize(int)

/*
Sets the sample size to the sampleSize specified, in bits.

This is typically 8 or 16, but some systems may support higher sample sizes.

See also sampleSize().
*/
func (this *QAudioFormat) SetSampleSize(sampleSize int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioFormat13setSampleSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sampleSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int sampleSize() const

/*
Returns the current sample size value, in bits.

See also setSampleSize() and bytesPerFrame().
*/
func (this *QAudioFormat) SampleSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormat10sampleSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCodec(const QString &)

/*
Sets the codec to codec.

The parameter to this function should be one of the types reported by the QAudioDeviceInfo::supportedCodecs() function for the audio device you are working with.

See also codec() and QAudioDeviceInfo::supportedCodecs().
*/
func (this *QAudioFormat) SetCodec(codec string) {
	var tmpArg0 = qtcore.NewQString_5(codec)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioFormat8setCodecERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QString codec() const

/*
Returns the current codec identifier.

See also setCodec() and QAudioDeviceInfo::supportedCodecs().
*/
func (this *QAudioFormat) Codec() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormat5codecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setByteOrder(QAudioFormat::Endian)

/*
Sets the byteOrder to byteOrder.

See also byteOrder().
*/
func (this *QAudioFormat) SetByteOrder(byteOrder int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioFormat12setByteOrderENS_6EndianE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), byteOrder)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] QAudioFormat::Endian byteOrder() const

/*
Returns the current byteOrder value.

See also setByteOrder().
*/
func (this *QAudioFormat) ByteOrder() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormat9byteOrderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSampleType(QAudioFormat::SampleType)

/*
Sets the sampleType to sampleType.

See also sampleType().
*/
func (this *QAudioFormat) SetSampleType(sampleType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioFormat13setSampleTypeENS_10SampleTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sampleType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] QAudioFormat::SampleType sampleType() const

/*
Returns the current SampleType value.

See also setSampleType().
*/
func (this *QAudioFormat) SampleType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormat10sampleTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] qint32 bytesForDuration(qint64) const

/*
Returns the number of bytes required for this audio format for duration microseconds.

Returns 0 if this format is not valid.

Note that some rounding may occur if duration is not an exact fraction of the sampleRate().

See also durationForBytes().
*/
func (this *QAudioFormat) BytesForDuration(duration int64) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormat16bytesForDurationEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), duration)
	qtrt.ErrPrint(err, rv)
	return int(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 durationForBytes(qint32) const

/*
Returns the number of microseconds represented by bytes in this format.

Returns 0 if this format is not valid.

Note that some rounding may occur if bytes is not an exact multiple of the number of bytes per frame.

See also bytesForDuration().
*/
func (this *QAudioFormat) DurationForBytes(byteCount int) int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormat16durationForBytesEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), byteCount)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] qint32 bytesForFrames(qint32) const

/*
Returns the number of bytes required for frameCount frames of this format.

Returns 0 if this format is not valid.

See also bytesForDuration().
*/
func (this *QAudioFormat) BytesForFrames(frameCount int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormat14bytesForFramesEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), frameCount)
	qtrt.ErrPrint(err, rv)
	return int(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] qint32 framesForBytes(qint32) const

/*
Returns the number of frames represented by byteCount in this format.

Note that some rounding may occur if byteCount is not an exact multiple of the number of bytes per frame.

Each frame has one sample per channel.

See also framesForDuration().
*/
func (this *QAudioFormat) FramesForBytes(byteCount int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormat14framesForBytesEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), byteCount)
	qtrt.ErrPrint(err, rv)
	return int(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] qint32 framesForDuration(qint64) const

/*
Returns the number of frames required to represent duration microseconds in this format.

Note that some rounding may occur if duration is not an exact fraction of the sampleRate().
*/
func (this *QAudioFormat) FramesForDuration(duration int64) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormat17framesForDurationEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), duration)
	qtrt.ErrPrint(err, rv)
	return int(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 durationForFrames(qint32) const

/*
Return the number of microseconds represented by frameCount frames in this format.
*/
func (this *QAudioFormat) DurationForFrames(frameCount int) int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormat17durationForFramesEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), frameCount)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudioformat.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] int bytesPerFrame() const

/*
Returns the number of bytes required to represent one frame (a sample in each channel) in this format.

Returns 0 if this format is invalid.
*/
func (this *QAudioFormat) BytesPerFrame() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioFormat13bytesPerFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

/*

 */
type QAudioFormat__SampleType = int

// Not Set
const QAudioFormat__Unknown QAudioFormat__SampleType = 0

// Samples are signed integers
const QAudioFormat__SignedInt QAudioFormat__SampleType = 1

// Samples are unsigned intergers
const QAudioFormat__UnSignedInt QAudioFormat__SampleType = 2

// Samples are floats
const QAudioFormat__Float QAudioFormat__SampleType = 3

/*
QAudioFormat::BigEndianQSysInfo::BigEndianSamples are big endian byte order
QAudioFormat::LittleEndianQSysInfo::LittleEndianSamples are little endian byte order

*/
type QAudioFormat__Endian = int

//
const QAudioFormat__BigEndian QAudioFormat__Endian = 0

//
const QAudioFormat__LittleEndian QAudioFormat__Endian = 1

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
