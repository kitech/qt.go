//  header block begin

// +build !minimal

package qtcore

// /usr/include/qt/QtCore/qsettings.h
// #include <qsettings.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// /usr/include/qt/QtCore/qsettings.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIniCodec(QTextCodec *)

/*
Sets the codec for accessing INI files (including .conf files on Unix) to codec. The codec is used for decoding any data that is read from the INI file, and for encoding any data that is written to the file. By default, no codec is used, and non-ASCII characters are encoded using standard INI escape sequences.

Warning: The codec must be set immediately after creating the QSettings object, before accessing any data.

This function was introduced in  Qt 4.5.

See also iniCodec().
*/
func (this *QSettings) SetIniCodec(codec QTextCodec_ITF /*777 QTextCodec **/) {
	var convArg0 unsafe.Pointer
	if codec != nil && codec.QTextCodec_PTR() != nil {
		convArg0 = codec.QTextCodec_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings11setIniCodecEP10QTextCodec", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:180
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setIniCodec(const char *)

/*
Sets the codec for accessing INI files (including .conf files on Unix) to codec. The codec is used for decoding any data that is read from the INI file, and for encoding any data that is written to the file. By default, no codec is used, and non-ASCII characters are encoded using standard INI escape sequences.

Warning: The codec must be set immediately after creating the QSettings object, before accessing any data.

This function was introduced in  Qt 4.5.

See also iniCodec().
*/
func (this *QSettings) SetIniCodec1(codecName string) {
	var convArg0 = qtrt.CString(codecName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings11setIniCodecEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:181
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCodec * iniCodec() const

/*
Returns the codec that is used for accessing INI files. By default, no codec is used, so a null pointer is returned.

This function was introduced in  Qt 4.5.

See also setIniCodec().
*/
func (this *QSettings) IniCodec() *QTextCodec /*777 QTextCodec **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings8iniCodecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_10538() {
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
