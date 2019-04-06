//  header block begin

// +build !minimal

package qtgui

// /usr/include/qt/QtGui/qtextdocumentwriter.h
// #include <qtextdocumentwriter.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// /usr/include/qt/QtGui/qtextdocumentwriter.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCodec(QTextCodec *)

/*
Sets the codec for this stream to codec. The codec is used for encoding any data that is written. By default, QTextDocumentWriter uses UTF-8.

See also codec().
*/
func (this *QTextDocumentWriter) SetCodec(codec qtcore.QTextCodec_ITF /*777 QTextCodec **/) {
	var convArg0 unsafe.Pointer
	if codec != nil && codec.QTextCodec_PTR() != nil {
		convArg0 = codec.QTextCodec_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QTextDocumentWriter8setCodecEP10QTextCodec", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentwriter.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCodec * codec() const

/*
Returns the codec that is currently assigned to the writer.

See also setCodec().
*/
func (this *QTextDocumentWriter) Codec() *qtcore.QTextCodec /*777 QTextCodec **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QTextDocumentWriter5codecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_10962() {
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
}

//  keep block end
