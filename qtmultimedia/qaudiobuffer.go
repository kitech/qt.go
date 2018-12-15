package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudiobuffer.h
// #include <qaudiobuffer.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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
type QAudioBuffer struct {
	*qtrt.CObject
}
type QAudioBuffer_ITF interface {
	QAudioBuffer_PTR() *QAudioBuffer
}

func (ptr *QAudioBuffer) QAudioBuffer_PTR() *QAudioBuffer { return ptr }

func (this *QAudioBuffer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAudioBuffer) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAudioBufferFromPointer(cthis unsafe.Pointer) *QAudioBuffer {
	return &QAudioBuffer{&qtrt.CObject{cthis}}
}
func (*QAudioBuffer) NewFromPointer(cthis unsafe.Pointer) *QAudioBuffer {
	return NewQAudioBufferFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioBuffer()

/*
Create a new, empty, invalid buffer.
*/
func (*QAudioBuffer) NewForInherit() *QAudioBuffer {
	return NewQAudioBuffer()
}
func NewQAudioBuffer() *QAudioBuffer {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioBufferC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioBufferFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAudioBuffer)
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:61
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAudioBuffer(const QByteArray &, const QAudioFormat &, qint64)

/*
Create a new, empty, invalid buffer.
*/
func (*QAudioBuffer) NewForInherit1(data qtcore.QByteArray_ITF, format QAudioFormat_ITF, startTime int64) *QAudioBuffer {
	return NewQAudioBuffer1(data, format, startTime)
}
func NewQAudioBuffer1(data qtcore.QByteArray_ITF, format QAudioFormat_ITF, startTime int64) *QAudioBuffer {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg1 = format.QAudioFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioBufferC2ERK10QByteArrayRK12QAudioFormatx", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, startTime)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioBufferFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAudioBuffer)
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:61
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAudioBuffer(const QByteArray &, const QAudioFormat &, qint64)

/*
Create a new, empty, invalid buffer.
*/
func (*QAudioBuffer) NewForInherit1p(data qtcore.QByteArray_ITF, format QAudioFormat_ITF) *QAudioBuffer {
	return NewQAudioBuffer1p(data, format)
}
func NewQAudioBuffer1p(data qtcore.QByteArray_ITF, format QAudioFormat_ITF) *QAudioBuffer {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg1 = format.QAudioFormat_PTR().GetCthis()
	}
	// arg: 2, qint64=Typedef, qint64=Typedef, long long, LongLong
	startTime := int64(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioBufferC2ERK10QByteArrayRK12QAudioFormatx", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, startTime)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioBufferFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAudioBuffer)
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:62
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QAudioBuffer(int, const QAudioFormat &, qint64)

/*
Create a new, empty, invalid buffer.
*/
func (*QAudioBuffer) NewForInherit2(numFrames int, format QAudioFormat_ITF, startTime int64) *QAudioBuffer {
	return NewQAudioBuffer2(numFrames, format, startTime)
}
func NewQAudioBuffer2(numFrames int, format QAudioFormat_ITF, startTime int64) *QAudioBuffer {
	var convArg1 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg1 = format.QAudioFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioBufferC2EiRK12QAudioFormatx", qtrt.FFI_TYPE_POINTER, numFrames, convArg1, startTime)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioBufferFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAudioBuffer)
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:62
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QAudioBuffer(int, const QAudioFormat &, qint64)

/*
Create a new, empty, invalid buffer.
*/
func (*QAudioBuffer) NewForInherit2p(numFrames int, format QAudioFormat_ITF) *QAudioBuffer {
	return NewQAudioBuffer2p(numFrames, format)
}
func NewQAudioBuffer2p(numFrames int, format QAudioFormat_ITF) *QAudioBuffer {
	var convArg1 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg1 = format.QAudioFormat_PTR().GetCthis()
	}
	// arg: 2, qint64=Typedef, qint64=Typedef, long long, LongLong
	startTime := int64(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioBufferC2EiRK12QAudioFormatx", qtrt.FFI_TYPE_POINTER, numFrames, convArg1, startTime)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioBufferFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAudioBuffer)
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] QAudioBuffer & operator=(const QAudioBuffer &)

/*

 */
func (this *QAudioBuffer) Operator_equal(other QAudioBuffer_ITF) *QAudioBuffer {
	var convArg0 unsafe.Pointer
	if other != nil && other.QAudioBuffer_PTR() != nil {
		convArg0 = other.QAudioBuffer_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioBufferaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioBufferFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioBuffer)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QAudioBuffer()

/*

 */
func DeleteQAudioBuffer(this *QAudioBuffer) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioBufferD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if this is a valid buffer. A valid buffer has more than zero frames in it and a valid format.
*/
func (this *QAudioBuffer) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioBuffer7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QAudioFormat format() const

/*
Returns the format of this buffer.

Several properties of this format influence how the duration() or byteCount() are calculated from the frameCount().
*/
func (this *QAudioBuffer) Format() *QAudioFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioBuffer6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioFormat)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:72
// index:0
// Public Visibility=Default Availability=Available
// [4] int frameCount() const

/*
Returns the number of complete audio frames in this buffer.

An audio frame is an interleaved set of one sample per channel for the same instant in time.
*/
func (this *QAudioBuffer) FrameCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioBuffer10frameCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] int sampleCount() const

/*
Returns the number of samples in this buffer.

If the format of this buffer has multiple channels, then this count includes all channels. This means that a stereo buffer with 1000 samples in total will have 500 left samples and 500 right samples (interleaved), and this function will return 1000.

See also frameCount().
*/
func (this *QAudioBuffer) SampleCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioBuffer11sampleCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] int byteCount() const

/*
Returns the size of this buffer, in bytes.
*/
func (this *QAudioBuffer) ByteCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioBuffer9byteCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 duration() const

/*
Returns the duration of audio in this buffer, in microseconds.

This depends on the /l format(), and the frameCount().
*/
func (this *QAudioBuffer) Duration() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioBuffer8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 startTime() const

/*
Returns the time in a stream that this buffer starts at (in microseconds).

If this buffer is not part of a stream, this will return -1.
*/
func (this *QAudioBuffer) StartTime() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioBuffer9startTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] const void * constData() const

/*
Returns a pointer to this buffer's data. You can only read it.

This method is preferred over the const version of data() to prevent unnecessary copying.

There is also a templatized version of this constData() function that allows you to retrieve a specific type of read-only pointer to the data. Note that there is no checking done on the format of the audio buffer - this is simply a convenience function.


  // With a 16bit sample buffer:
  const quint16 *data = buffer->constData<quint16>();
*/
func (this *QAudioBuffer) ConstData() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioBuffer9constDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] const void * data() const

/*
Returns a pointer to this buffer's data. You can only read it.

You should use the constData() function rather than this to prevent accidental deep copying.

There is also a templatized version of this data() function that allows you to retrieve a specific type of read-only pointer to the data. Note that there is no checking done on the format of the audio buffer - this is simply a convenience function.


  // With a 16bit sample const buffer:
  const quint16 *data = buffer->data<quint16>();
*/
func (this *QAudioBuffer) Data() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QAudioBuffer4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtMultimedia/qaudiobuffer.h:88
// index:1
// Public Visibility=Default Availability=Available
// [8] void * data()

/*
Returns a pointer to this buffer's data. You can only read it.

You should use the constData() function rather than this to prevent accidental deep copying.

There is also a templatized version of this data() function that allows you to retrieve a specific type of read-only pointer to the data. Note that there is no checking done on the format of the audio buffer - this is simply a convenience function.


  // With a 16bit sample const buffer:
  const quint16 *data = buffer->data<quint16>();
*/
func (this *QAudioBuffer) Data1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QAudioBuffer4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
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
