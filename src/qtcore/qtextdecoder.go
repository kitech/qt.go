//  header block begin
// /usr/include/qt/QtCore/qtextcodec.h
// #include <qtextcodec.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QTextDecoder struct {
	*qtrt.CObject
}

func (this *QTextDecoder) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qtextcodec.h:158
// index:0
// inline
// void QTextDecoder(const class QTextCodec *)
func NewQTextDecoder(codec unsafe.Pointer) *QTextDecoder {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextDecoderC2EPK10QTextCodec", ffiqt.FFI_TYPE_VOID, cthis, codec)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDecoderFromPointer(cthis)
	return gothis
}
func NewQTextDecoderFromPointer(cthis unsafe.Pointer) *QTextDecoder {
	return &QTextDecoder{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qtextcodec.h:159
// index:1
// void QTextDecoder(const class QTextCodec *, class QTextCodec::ConversionFlags)
func NewQTextDecoder_1(codec unsafe.Pointer, flags int) *QTextDecoder {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextDecoderC2EPK10QTextCodec6QFlagsINS0_14ConversionFlagEE", ffiqt.FFI_TYPE_VOID, cthis, codec, &flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDecoderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextcodec.h:160
// index:0
// void ~QTextDecoder()
func DeleteQTextDecoder(*QTextDecoder) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextDecoderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:161
// index:0
// QString toUnicode(const char *, int)
func (this *QTextDecoder) ToUnicode(chars unsafe.Pointer, len int) {
	// 0: (, chars const char *, len int), (chars, &len)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextDecoder9toUnicodeEPKci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), chars, &len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:162
// index:1
// QString toUnicode(const class QByteArray &)
func (this *QTextDecoder) ToUnicode_1(ba unsafe.Pointer) {
	// 1: (, ba const QByteArray &), (ba)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextDecoder9toUnicodeERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ba)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:163
// index:2
// void toUnicode(class QString *, const char *, int)
func (this *QTextDecoder) ToUnicode_2(target unsafe.Pointer, chars unsafe.Pointer, len int) {
	// 2: (, target QString *, chars const char *, len int), (target, chars, &len)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextDecoder9toUnicodeEP7QStringPKci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), target, chars, &len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:164
// index:0
// bool hasFailure()
func (this *QTextDecoder) HasFailure() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextDecoder10hasFailureEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
