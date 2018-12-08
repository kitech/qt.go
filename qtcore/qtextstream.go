package qtcore

// /usr/include/qt/QtCore/qtextstream.h
// #include <qtextstream.h>
// #include <QtCore>

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

//  ext block end

//  body block begin

/*

 */
type QTextStream struct {
	*qtrt.CObject
}
type QTextStream_ITF interface {
	QTextStream_PTR() *QTextStream
}

func (ptr *QTextStream) QTextStream_PTR() *QTextStream { return ptr }

func (this *QTextStream) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextStream) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextStreamFromPointer(cthis unsafe.Pointer) *QTextStream {
	return &QTextStream{&qtrt.CObject{cthis}}
}
func (*QTextStream) NewFromPointer(cthis unsafe.Pointer) *QTextStream {
	return NewQTextStreamFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtextstream.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextStream()

/*
Constructs a QTextStream. Before you can use it for reading or writing, you must assign a device or a string.

See also setDevice() and setString().
*/
func (*QTextStream) NewForInherit() *QTextStream {
	return NewQTextStream()
}
func NewQTextStream() *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:94
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextStream(QIODevice *)

/*
Constructs a QTextStream. Before you can use it for reading or writing, you must assign a device or a string.

See also setDevice() and setString().
*/
func (*QTextStream) NewForInherit1(device QIODevice_ITF /*777 QIODevice **/) *QTextStream {
	return NewQTextStream1(device)
}
func NewQTextStream1(device QIODevice_ITF /*777 QIODevice **/) *QTextStream {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2EP9QIODevice", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:96
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextStream(QString *, QIODevice::OpenMode)

/*
Constructs a QTextStream. Before you can use it for reading or writing, you must assign a device or a string.

See also setDevice() and setString().
*/
func (*QTextStream) NewForInherit2(string string, openMode int) *QTextStream {
	return NewQTextStream2(string, openMode)
}
func NewQTextStream2(string string, openMode int) *QTextStream {
	var tmpArg0 = NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2EP7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, openMode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:96
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextStream(QString *, QIODevice::OpenMode)

/*
Constructs a QTextStream. Before you can use it for reading or writing, you must assign a device or a string.

See also setDevice() and setString().
*/
func (*QTextStream) NewForInherit2p(string string) *QTextStream {
	return NewQTextStream2p(string)
}
func NewQTextStream2p(string string) *QTextStream {
	var tmpArg0 = NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QIODevice::OpenMode=Elaborated, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>, Unexposed
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2EP7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, openMode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:97
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTextStream(QByteArray *, QIODevice::OpenMode)

/*
Constructs a QTextStream. Before you can use it for reading or writing, you must assign a device or a string.

See also setDevice() and setString().
*/
func (*QTextStream) NewForInherit3(array QByteArray_ITF /*777 QByteArray **/, openMode int) *QTextStream {
	return NewQTextStream3(array, openMode)
}
func NewQTextStream3(array QByteArray_ITF /*777 QByteArray **/, openMode int) *QTextStream {
	var convArg0 unsafe.Pointer
	if array != nil && array.QByteArray_PTR() != nil {
		convArg0 = array.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2EP10QByteArray6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, openMode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:97
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTextStream(QByteArray *, QIODevice::OpenMode)

/*
Constructs a QTextStream. Before you can use it for reading or writing, you must assign a device or a string.

See also setDevice() and setString().
*/
func (*QTextStream) NewForInherit3p(array QByteArray_ITF /*777 QByteArray **/) *QTextStream {
	return NewQTextStream3p(array)
}
func NewQTextStream3p(array QByteArray_ITF /*777 QByteArray **/) *QTextStream {
	var convArg0 unsafe.Pointer
	if array != nil && array.QByteArray_PTR() != nil {
		convArg0 = array.QByteArray_PTR().GetCthis()
	}
	// arg: 1, QIODevice::OpenMode=Elaborated, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>, Unexposed
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2EP10QByteArray6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, openMode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:98
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QTextStream(const QByteArray &, QIODevice::OpenMode)

/*
Constructs a QTextStream. Before you can use it for reading or writing, you must assign a device or a string.

See also setDevice() and setString().
*/
func (*QTextStream) NewForInherit4(array QByteArray_ITF, openMode int) *QTextStream {
	return NewQTextStream4(array, openMode)
}
func NewQTextStream4(array QByteArray_ITF, openMode int) *QTextStream {
	var convArg0 unsafe.Pointer
	if array != nil && array.QByteArray_PTR() != nil {
		convArg0 = array.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2ERK10QByteArray6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, openMode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:98
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QTextStream(const QByteArray &, QIODevice::OpenMode)

/*
Constructs a QTextStream. Before you can use it for reading or writing, you must assign a device or a string.

See also setDevice() and setString().
*/
func (*QTextStream) NewForInherit4p(array QByteArray_ITF) *QTextStream {
	return NewQTextStream4p(array)
}
func NewQTextStream4p(array QByteArray_ITF) *QTextStream {
	var convArg0 unsafe.Pointer
	if array != nil && array.QByteArray_PTR() != nil {
		convArg0 = array.QByteArray_PTR().GetCthis()
	}
	// arg: 1, QIODevice::OpenMode=Elaborated, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>, Unexposed
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamC2ERK10QByteArray6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, openMode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextStream)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTextStream()

/*

 */
func DeleteQTextStream(this *QTextStream) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtextstream.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCodec(QTextCodec *)

/*
Sets the codec for this stream to codec. The codec is used for decoding any data that is read from the assigned device, and for encoding any data that is written. By default, QTextCodec::codecForLocale() is used, and automatic unicode detection is enabled.

If QTextStream operates on a string, this function does nothing.

Warning: If you call this function while the text stream is reading from an open sequential socket, the internal buffer may still contain text decoded using the old codec.

See also codec(), setAutoDetectUnicode(), and setLocale().
*/
func (this *QTextStream) SetCodec(codec QTextCodec_ITF /*777 QTextCodec **/) {
	var convArg0 unsafe.Pointer
	if codec != nil && codec.QTextCodec_PTR() != nil {
		convArg0 = codec.QTextCodec_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream8setCodecEP10QTextCodec", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:103
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setCodec(const char *)

/*
Sets the codec for this stream to codec. The codec is used for decoding any data that is read from the assigned device, and for encoding any data that is written. By default, QTextCodec::codecForLocale() is used, and automatic unicode detection is enabled.

If QTextStream operates on a string, this function does nothing.

Warning: If you call this function while the text stream is reading from an open sequential socket, the internal buffer may still contain text decoded using the old codec.

See also codec(), setAutoDetectUnicode(), and setLocale().
*/
func (this *QTextStream) SetCodec1(codecName string) {
	var convArg0 = qtrt.CString(codecName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream8setCodecEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCodec * codec() const

/*
Returns the codec that is current assigned to the stream.

See also setCodec(), setAutoDetectUnicode(), and locale().
*/
func (this *QTextStream) Codec() *QTextCodec /*777 QTextCodec **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream5codecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qtextstream.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoDetectUnicode(bool)

/*
If enabled is true, QTextStream will attempt to detect Unicode encoding by peeking into the stream data to see if it can find the UTF-16 or UTF-32 BOM (Byte Order Mark). If this mark is found, QTextStream will replace the current codec with the UTF codec.

This function can be used together with setCodec(). It is common to set the codec to UTF-8, and then enable UTF-16 detection.

See also autoDetectUnicode() and setCodec().
*/
func (this *QTextStream) SetAutoDetectUnicode(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream20setAutoDetectUnicodeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoDetectUnicode() const

/*
Returns true if automatic Unicode detection is enabled, otherwise returns false. Automatic Unicode detection is enabled by default.

See also setAutoDetectUnicode() and setCodec().
*/
func (this *QTextStream) AutoDetectUnicode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream17autoDetectUnicodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGenerateByteOrderMark(bool)

/*
If generate is true and a UTF codec is used, QTextStream will insert the BOM (Byte Order Mark) before any data has been written to the device. If generate is false, no BOM will be inserted. This function must be called before any data is written. Otherwise, it does nothing.

See also generateByteOrderMark() and bom().
*/
func (this *QTextStream) SetGenerateByteOrderMark(generate bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream24setGenerateByteOrderMarkEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), generate)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:108
// index:0
// Public Visibility=Default Availability=Available
// [1] bool generateByteOrderMark() const

/*
Returns true if QTextStream is set to generate the UTF BOM (Byte Order Mark) when using a UTF codec; otherwise returns false. UTF BOM generation is set to false by default.

See also setGenerateByteOrderMark().
*/
func (this *QTextStream) GenerateByteOrderMark() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream21generateByteOrderMarkEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocale(const QLocale &)

/*
Sets the locale for this stream to locale. The specified locale is used for conversions between numbers and their string representations.

The default locale is C and it is a special case - the thousands group separator is not used for backward compatibility reasons.

This function was introduced in  Qt 4.5.

See also locale().
*/
func (this *QTextStream) SetLocale(locale QLocale_ITF) {
	var convArg0 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg0 = locale.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream9setLocaleERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QLocale locale() const

/*
Returns the locale for this stream. The default locale is C.

This function was introduced in  Qt 4.5.

See also setLocale().
*/
func (this *QTextStream) Locale() *QLocale /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream6localeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLocale)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)

/*
Sets the current device to device. If a device has already been assigned, QTextStream will call flush() before the old device is replaced.

Note: This function resets locale to the default locale ('C') and codec to the default codec, QTextCodec::codecForLocale().

See also device() and setString().
*/
func (this *QTextStream) SetDevice(device QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device() const

/*
Returns the current device associated with the QTextStream, or 0 if no device has been assigned.

See also setDevice() and string().
*/
func (this *QTextStream) Device() *QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qtextstream.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setString(QString *, QIODevice::OpenMode)

/*
Sets the current string to string, using the given openMode. If a device has already been assigned, QTextStream will call flush() before replacing it.

See also string() and setDevice().
*/
func (this *QTextStream) SetString(string string, openMode int) {
	var tmpArg0 = NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream9setStringEP7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, openMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setString(QString *, QIODevice::OpenMode)

/*
Sets the current string to string, using the given openMode. If a device has already been assigned, QTextStream will call flush() before replacing it.

See also string() and setDevice().
*/
func (this *QTextStream) SetStringp(string string) {
	var tmpArg0 = NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QIODevice::OpenMode=Elaborated, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>, Unexposed
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream9setStringEP7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, openMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QString * string() const

/*
Returns the current string assigned to the QTextStream, or 0 if no string has been assigned.

See also setString() and device().
*/
func (this *QTextStream) String() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream6stringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextstream.h:120
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextStream::Status status() const

/*
Returns the status of the text stream.

See also QTextStream::Status, setStatus(), and resetStatus().
*/
func (this *QTextStream) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtextstream.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStatus(QTextStream::Status)

/*
Sets the status of the text stream to the status given.

Subsequent calls to setStatus() are ignored until resetStatus() is called.

This function was introduced in  Qt 4.1.

See also Status, status(), and resetStatus().
*/
func (this *QTextStream) SetStatus(status int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream9setStatusENS_6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), status)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetStatus()

/*
Resets the status of the text stream.

This function was introduced in  Qt 4.1.

See also QTextStream::Status, status(), and setStatus().
*/
func (this *QTextStream) ResetStatus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream11resetStatusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:124
// index:0
// Public Visibility=Default Availability=Available
// [1] bool atEnd() const

/*
Returns true if there is no more data to be read from the QTextStream; otherwise returns false. This is similar to, but not the same as calling QIODevice::atEnd(), as QTextStream also takes into account its internal Unicode buffer.
*/
func (this *QTextStream) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()

/*
Resets QTextStream's formatting options, bringing it back to its original constructed state. The device, string and any buffered data is left untouched.
*/
func (this *QTextStream) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void flush()

/*
Flushes any buffered data waiting to be written to the device.

If QTextStream operates on a string, this function does nothing.
*/
func (this *QTextStream) Flush() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:127
// index:0
// Public Visibility=Default Availability=Available
// [1] bool seek(qint64)

/*
Seeks to the position pos in the device. Returns true on success; otherwise returns false.
*/
func (this *QTextStream) Seek(pos int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream4seekEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 pos() const

/*
Returns the device position corresponding to the current position of the stream, or -1 if an error occurs (e.g., if there is no device or string, or if there's a device error).

Because QTextStream is buffered, this function may have to seek the device to reconstruct a valid device position. This operation can be expensive, so you may want to avoid calling this function in a tight loop.

This function was introduced in  Qt 4.2.

See also seek().
*/
func (this *QTextStream) Pos() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qtextstream.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void skipWhiteSpace()

/*
Reads and discards whitespace from the stream until either a non-space character is detected, or until atEnd() returns true. This function is useful when reading a stream character by character.

Whitespace characters are all characters for which QChar::isSpace() returns true.

See also operator>>().
*/
func (this *QTextStream) SkipWhiteSpace() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream14skipWhiteSpaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QString readLine(qint64)

/*
Reads one line of text from the stream, and returns it as a QString. The maximum allowed line length is set to maxlen. If the stream contains lines longer than this, then the lines will be split after maxlen characters and returned in parts.

If maxlen is 0, the lines can be of any length.

The returned line has no trailing end-of-line characters ("\n" or "\r\n"), so calling QString::trimmed() can be unnecessary.

If the stream has read to the end of the file, readLine() will return a null QString. For strings, or for devices that support it, you can explicitly test for the end of the stream using atEnd().

See also readAll() and QIODevice::readLine().
*/
func (this *QTextStream) ReadLine(maxlen int64) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream8readLineEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxlen)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextstream.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QString readLine(qint64)

/*
Reads one line of text from the stream, and returns it as a QString. The maximum allowed line length is set to maxlen. If the stream contains lines longer than this, then the lines will be split after maxlen characters and returned in parts.

If maxlen is 0, the lines can be of any length.

The returned line has no trailing end-of-line characters ("\n" or "\r\n"), so calling QString::trimmed() can be unnecessary.

If the stream has read to the end of the file, readLine() will return a null QString. For strings, or for devices that support it, you can explicitly test for the end of the stream using atEnd().

See also readAll() and QIODevice::readLine().
*/
func (this *QTextStream) ReadLinep() string {
	// arg: 0, qint64=Typedef, qint64=Typedef, long long, LongLong
	maxlen := int64(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream8readLineEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxlen)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextstream.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool readLineInto(QString *, qint64)

/*
Reads one line of text from the stream into line. If line is 0, the read line is not stored.

The maximum allowed line length is set to maxlen. If the stream contains lines longer than this, then the lines will be split after maxlen characters and returned in parts.

If maxlen is 0, the lines can be of any length.

The resulting line has no trailing end-of-line characters ("\n" or "\r\n"), so calling QString::trimmed() can be unnecessary.

If line has sufficient capacity for the data that is about to be read, this function may not need to allocate new memory. Because of this, it can be faster than readLine().

Returns false if the stream has read to the end of the file or an error has occurred; otherwise returns true. The contents in line before the call are discarded in any case.

This function was introduced in  Qt 5.5.

See also readAll() and QIODevice::readLine().
*/
func (this *QTextStream) ReadLineInto(line string, maxlen int64) bool {
	var tmpArg0 = NewQString5(line)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream12readLineIntoEP7QStringx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool readLineInto(QString *, qint64)

/*
Reads one line of text from the stream into line. If line is 0, the read line is not stored.

The maximum allowed line length is set to maxlen. If the stream contains lines longer than this, then the lines will be split after maxlen characters and returned in parts.

If maxlen is 0, the lines can be of any length.

The resulting line has no trailing end-of-line characters ("\n" or "\r\n"), so calling QString::trimmed() can be unnecessary.

If line has sufficient capacity for the data that is about to be read, this function may not need to allocate new memory. Because of this, it can be faster than readLine().

Returns false if the stream has read to the end of the file or an error has occurred; otherwise returns true. The contents in line before the call are discarded in any case.

This function was introduced in  Qt 5.5.

See also readAll() and QIODevice::readLine().
*/
func (this *QTextStream) ReadLineIntop(line string) bool {
	var tmpArg0 = NewQString5(line)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, qint64=Typedef, qint64=Typedef, long long, LongLong
	maxlen := int64(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream12readLineIntoEP7QStringx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QString readAll()

/*
Reads the entire content of the stream, and returns it as a QString. Avoid this function when working on large files, as it will consume a significant amount of memory.

Calling readLine() is better if you do not know how much data is available.

See also readLine().
*/
func (this *QTextStream) ReadAll() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream7readAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextstream.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QString read(qint64)

/*
Reads at most maxlen characters from the stream, and returns the data read as a QString.

This function was introduced in  Qt 4.1.

See also readAll(), readLine(), and QIODevice::read().
*/
func (this *QTextStream) Read(maxlen int64) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream4readEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxlen)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextstream.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFieldAlignment(QTextStream::FieldAlignment)

/*
Sets the field alignment to mode. When used together with setFieldWidth(), this function allows you to generate formatted output with text aligned to the left, to the right or center aligned.

See also fieldAlignment() and setFieldWidth().
*/
func (this *QTextStream) SetFieldAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream17setFieldAlignmentENS_14FieldAlignmentE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:138
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextStream::FieldAlignment fieldAlignment() const

/*
Returns the current field alignment.

See also setFieldAlignment() and fieldWidth().
*/
func (this *QTextStream) FieldAlignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream14fieldAlignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtextstream.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPadChar(QChar)

/*
Sets the pad character to ch. The default value is the ASCII space character (' '), or QChar(0x20). This character is used to fill in the space in fields when generating text.

Example:


  QString s;
  QTextStream out(&s);
  out.setFieldWidth(10);
  out.setFieldAlignment(QTextStream::AlignCenter);
  out.setPadChar('-');
  out << "Qt" << "rocks!";



The string s contains:


  ----Qt------rocks!--



See also padChar() and setFieldWidth().
*/
func (this *QTextStream) SetPadChar(ch QChar_ITF /*123*/) {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream10setPadCharE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:141
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar padChar() const

/*
Returns the current pad character.

See also setPadChar() and setFieldWidth().
*/
func (this *QTextStream) PadChar() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream7padCharEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFieldWidth(int)

/*
Sets the current field width to width. If width is 0 (the default), the field width is equal to the length of the generated text.

Note: The field width applies to every element appended to this stream after this function has been called (e.g., it also pads endl). This behavior is different from similar classes in the STL, where the field width only applies to the next element.

See also fieldWidth() and setPadChar().
*/
func (this *QTextStream) SetFieldWidth(width int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream13setFieldWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:144
// index:0
// Public Visibility=Default Availability=Available
// [4] int fieldWidth() const

/*
Returns the current field width.

See also setFieldWidth().
*/
func (this *QTextStream) FieldWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream10fieldWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextstream.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNumberFlags(QTextStream::NumberFlags)

/*
Sets the current number flags to flags. flags is a set of flags from the NumberFlag enum, and describes options for formatting generated code (e.g., whether or not to always write the base or sign of a number).

See also numberFlags(), setIntegerBase(), and setRealNumberNotation().
*/
func (this *QTextStream) SetNumberFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream14setNumberFlagsE6QFlagsINS_10NumberFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:147
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextStream::NumberFlags numberFlags() const

/*
Returns the current number flags.

See also setNumberFlags(), integerBase(), and realNumberNotation().
*/
func (this *QTextStream) NumberFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream11numberFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtextstream.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIntegerBase(int)

/*
Sets the base of integers to base, both for reading and for generating numbers. base can be either 2 (binary), 8 (octal), 10 (decimal) or 16 (hexadecimal). If base is 0, QTextStream will attempt to detect the base by inspecting the data on the stream. When generating numbers, QTextStream assumes base is 10 unless the base has been set explicitly.

See also integerBase(), QString::number(), and setNumberFlags().
*/
func (this *QTextStream) SetIntegerBase(base int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream14setIntegerBaseEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), base)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:150
// index:0
// Public Visibility=Default Availability=Available
// [4] int integerBase() const

/*
Returns the current base of integers. 0 means that the base is detected when reading, or 10 (decimal) when generating numbers.

See also setIntegerBase(), QString::number(), and numberFlags().
*/
func (this *QTextStream) IntegerBase() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream11integerBaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextstream.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRealNumberNotation(QTextStream::RealNumberNotation)

/*
Sets the real number notation to notation (SmartNotation, FixedNotation, ScientificNotation). When reading and generating numbers, QTextStream uses this value to detect the formatting of real numbers.

See also realNumberNotation(), setRealNumberPrecision(), setNumberFlags(), and setIntegerBase().
*/
func (this *QTextStream) SetRealNumberNotation(notation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream21setRealNumberNotationENS_18RealNumberNotationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), notation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:153
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextStream::RealNumberNotation realNumberNotation() const

/*
Returns the current real number notation.

See also setRealNumberNotation(), realNumberPrecision(), numberFlags(), and integerBase().
*/
func (this *QTextStream) RealNumberNotation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream18realNumberNotationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtextstream.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRealNumberPrecision(int)

/*
Sets the precision of real numbers to precision. This value describes the number of fraction digits QTextStream should write when generating real numbers.

The precision cannot be a negative value. The default value is 6.

See also realNumberPrecision() and setRealNumberNotation().
*/
func (this *QTextStream) SetRealNumberPrecision(precision int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStream22setRealNumberPrecisionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), precision)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:156
// index:0
// Public Visibility=Default Availability=Available
// [4] int realNumberPrecision() const

/*
Returns the current real number precision, or the number of fraction digits QTextStream will write when generating real numbers.

See also setRealNumberPrecision(), setRealNumberNotation(), realNumberNotation(), numberFlags(), and integerBase().
*/
func (this *QTextStream) RealNumberPrecision() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextStream19realNumberPrecisionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextstream.h:158
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(QChar &)

/*

 */
func (this *QTextStream) Operator_right_shift(ch QChar_ITF) *QTextStream {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsER5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:159
// index:1
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(char &)

/*

 */
func (this *QTextStream) Operator_right_shift1(ch byte) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ch)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:160
// index:2
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(short &)

/*

 */
func (this *QTextStream) Operator_right_shift2(i int16) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERs", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:161
// index:3
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(unsigned short &)

/*

 */
func (this *QTextStream) Operator_right_shift3(i uint16) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:162
// index:4
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(int &)

/*

 */
func (this *QTextStream) Operator_right_shift4(i int) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:163
// index:5
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(unsigned int &)

/*

 */
func (this *QTextStream) Operator_right_shift5(i uint) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:164
// index:6
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(long &)

/*

 */
func (this *QTextStream) Operator_right_shift6(i int) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:165
// index:7
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(unsigned long &)

/*

 */
func (this *QTextStream) Operator_right_shift7(i uint) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERm", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:166
// index:8
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(qlonglong &)

/*

 */
func (this *QTextStream) Operator_right_shift8(i int64) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:167
// index:9
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(qulonglong &)

/*

 */
func (this *QTextStream) Operator_right_shift9(i uint64) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:168
// index:10
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(float &)

/*

 */
func (this *QTextStream) Operator_right_shift10(f float32) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:169
// index:11
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(double &)

/*

 */
func (this *QTextStream) Operator_right_shift11(f float64) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsERd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:170
// index:12
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(QString &)

/*

 */
func (this *QTextStream) Operator_right_shift12(s string) *QTextStream {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsER7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:171
// index:13
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(QByteArray &)

/*

 */
func (this *QTextStream) Operator_right_shift13(array QByteArray_ITF) *QTextStream {
	var convArg0 unsafe.Pointer
	if array != nil && array.QByteArray_PTR() != nil {
		convArg0 = array.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsER10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:172
// index:14
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator>>(char *)

/*

 */
func (this *QTextStream) Operator_right_shift14(c string) *QTextStream {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamrsEPc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:174
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(QChar)

/*

 */
func (this *QTextStream) Operator_left_shift(ch QChar_ITF /*123*/) *QTextStream {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:175
// index:1
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(char)

/*

 */
func (this *QTextStream) Operator_left_shift1(ch byte) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ch)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:176
// index:2
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(short)

/*

 */
func (this *QTextStream) Operator_left_shift2(i int16) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEs", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:177
// index:3
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(unsigned short)

/*

 */
func (this *QTextStream) Operator_left_shift3(i uint16) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:178
// index:4
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(int)

/*

 */
func (this *QTextStream) Operator_left_shift4(i int) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:179
// index:5
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(unsigned int)

/*

 */
func (this *QTextStream) Operator_left_shift5(i uint) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:180
// index:6
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(long)

/*

 */
func (this *QTextStream) Operator_left_shift6(i int) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:181
// index:7
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(unsigned long)

/*

 */
func (this *QTextStream) Operator_left_shift7(i uint) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEm", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:182
// index:8
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(qlonglong)

/*

 */
func (this *QTextStream) Operator_left_shift8(i int64) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:183
// index:9
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(qulonglong)

/*

 */
func (this *QTextStream) Operator_left_shift9(i uint64) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:184
// index:10
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(float)

/*

 */
func (this *QTextStream) Operator_left_shift10(f float32) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:185
// index:11
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(double)

/*

 */
func (this *QTextStream) Operator_left_shift11(f float64) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:186
// index:12
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(const QString &)

/*

 */
func (this *QTextStream) Operator_left_shift12(s string) *QTextStream {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:187
// index:13
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(QLatin1String)

/*

 */
func (this *QTextStream) Operator_left_shift13(s QLatin1String_ITF /*123*/) *QTextStream {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:188
// index:14
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(const QStringRef &)

/*

 */
func (this *QTextStream) Operator_left_shift14(s QStringRef_ITF) *QTextStream {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsERK10QStringRef", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:189
// index:15
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(const QByteArray &)

/*

 */
func (this *QTextStream) Operator_left_shift15(array QByteArray_ITF) *QTextStream {
	var convArg0 unsafe.Pointer
	if array != nil && array.QByteArray_PTR() != nil {
		convArg0 = array.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:190
// index:16
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(const char *)

/*

 */
func (this *QTextStream) Operator_left_shift16(c string) *QTextStream {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:191
// index:17
// Public Visibility=Default Availability=Available
// [16] QTextStream & operator<<(const void *)

/*

 */
func (this *QTextStream) Operator_left_shift17(ptr unsafe.Pointer /*666*/) *QTextStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextStreamlsEPKv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ptr)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStream)
	return rv2
}

/*
This enum specifies which notations to use for expressing float and double as strings.



See also setRealNumberNotation().

*/
type QTextStream__RealNumberNotation = int

// Scientific or fixed-point notation, depending on which makes most sense (printf()'s %g flag).
const QTextStream__SmartNotation QTextStream__RealNumberNotation = 0

// Fixed-point notation (printf()'s %f flag).
const QTextStream__FixedNotation QTextStream__RealNumberNotation = 1

// Scientific notation (printf()'s %e flag).
const QTextStream__ScientificNotation QTextStream__RealNumberNotation = 2

func (this *QTextStream) RealNumberNotationItemName(val int) string {
	switch val {
	case QTextStream__SmartNotation: // 0
		return "SmartNotation"
	case QTextStream__FixedNotation: // 1
		return "FixedNotation"
	case QTextStream__ScientificNotation: // 2
		return "ScientificNotation"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextStream_RealNumberNotationItemName(val int) string {
	var nilthis *QTextStream
	return nilthis.RealNumberNotationItemName(val)
}

/*
This enum specifies how to align text in fields when the field is wider than the text that occupies it.



See also setFieldAlignment().

*/
type QTextStream__FieldAlignment = int

// Pad on the right side of fields.
const QTextStream__AlignLeft QTextStream__FieldAlignment = 0

// Pad on the left side of fields.
const QTextStream__AlignRight QTextStream__FieldAlignment = 1

// Pad on both sides of field.
const QTextStream__AlignCenter QTextStream__FieldAlignment = 2

// Same as AlignRight, except that the sign of a number is flush left.
const QTextStream__AlignAccountingStyle QTextStream__FieldAlignment = 3

func (this *QTextStream) FieldAlignmentItemName(val int) string {
	switch val {
	case QTextStream__AlignLeft: // 0
		return "AlignLeft"
	case QTextStream__AlignRight: // 1
		return "AlignRight"
	case QTextStream__AlignCenter: // 2
		return "AlignCenter"
	case QTextStream__AlignAccountingStyle: // 3
		return "AlignAccountingStyle"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextStream_FieldAlignmentItemName(val int) string {
	var nilthis *QTextStream
	return nilthis.FieldAlignmentItemName(val)
}

/*
This enum describes the current status of the text stream.



See also status().

*/
type QTextStream__Status = int

// The text stream is operating normally.
const QTextStream__Ok QTextStream__Status = 0

// The text stream has read past the end of the data in the underlying device.
const QTextStream__ReadPastEnd QTextStream__Status = 1

// The text stream has read corrupt data.
const QTextStream__ReadCorruptData QTextStream__Status = 2

// The text stream cannot write to the underlying device.
const QTextStream__WriteFailed QTextStream__Status = 3

func (this *QTextStream) StatusItemName(val int) string {
	switch val {
	case QTextStream__Ok: // 0
		return "Ok"
	case QTextStream__ReadPastEnd: // 1
		return "ReadPastEnd"
	case QTextStream__ReadCorruptData: // 2
		return "ReadCorruptData"
	case QTextStream__WriteFailed: // 3
		return "WriteFailed"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextStream_StatusItemName(val int) string {
	var nilthis *QTextStream
	return nilthis.StatusItemName(val)
}

/*


 */
type QTextStream__NumberFlag = int

//
const QTextStream__ShowBase QTextStream__NumberFlag = 1

//
const QTextStream__ForcePoint QTextStream__NumberFlag = 2

//
const QTextStream__ForceSign QTextStream__NumberFlag = 4

//
const QTextStream__UppercaseBase QTextStream__NumberFlag = 8

//
const QTextStream__UppercaseDigits QTextStream__NumberFlag = 16

func (this *QTextStream) NumberFlagItemName(val int) string {
	switch val {
	case QTextStream__ShowBase: // 1
		return "ShowBase"
	case QTextStream__ForcePoint: // 2
		return "ForcePoint"
	case QTextStream__ForceSign: // 4
		return "ForceSign"
	case QTextStream__UppercaseBase: // 8
		return "UppercaseBase"
	case QTextStream__UppercaseDigits: // 16
		return "UppercaseDigits"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextStream_NumberFlagItemName(val int) string {
	var nilthis *QTextStream
	return nilthis.NumberFlagItemName(val)
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
}

//  keep block end
