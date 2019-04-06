//  header block begin

// +build !minimal

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
// extern C begin: 109
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

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
If enabled is true, QTextStream will attempt to detect Unicode encoding by peeking into the stream data to see if it can find the UTF-8, UTF-16, or UTF-32 Byte Order Mark (BOM). If this mark is found, QTextStream will replace the current codec with the UTF codec.

This function can be used together with setCodec(). It is common to set the codec to UTF-8, and then enable UTF-16 detection.

See also autoDetectUnicode(), setCodec(), and QTextCodec::codecForUtfText().
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

See also setAutoDetectUnicode(), setCodec(), and QTextCodec::codecForUtfText().
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

//  body block end

//  keep block begin

func init_unused_10328() {
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
