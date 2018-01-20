//  header block begin
// /usr/include/qt/QtCore/qtextcodec.h
// #include <qtextcodec.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTextEncoder struct {
	*qtrt.CObject
}

func (this *QTextEncoder) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQTextEncoderFromPointer(cthis unsafe.Pointer) *QTextEncoder {
	return &QTextEncoder{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qtextcodec.h:141
// index:0
// Public inline
// void QTextEncoder(const class QTextCodec *)
func NewQTextEncoder(codec unsafe.Pointer) *QTextEncoder {
	cthis := qtrt.Calloc(1, 256) // 40
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextEncoderC2EPK10QTextCodec", ffiqt.FFI_TYPE_VOID, cthis, codec)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextEncoderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextcodec.h:143
// index:0
// Public
// void ~QTextEncoder()
func DeleteQTextEncoder(*QTextEncoder) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextEncoderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:145
// index:0
// Public
// QByteArray fromUnicode(const class QString &)
func (this *QTextEncoder) FromUnicode(str *QString) interface{} {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextEncoder11fromUnicodeERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:147
// index:1
// Public
// QByteArray fromUnicode(class QStringView)
func (this *QTextEncoder) FromUnicode_1(str *QStringView) interface{} {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextEncoder11fromUnicodeE11QStringView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:148
// index:2
// Public
// QByteArray fromUnicode(const class QChar *, int)
func (this *QTextEncoder) FromUnicode_2(uc unsafe.Pointer, len int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextEncoder11fromUnicodeEPK5QChari", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), uc, &len)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:149
// index:0
// Public
// bool hasFailure()
func (this *QTextEncoder) HasFailure() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextEncoder10hasFailureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
