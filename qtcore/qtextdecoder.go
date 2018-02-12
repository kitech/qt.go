package qtcore

// /usr/include/qt/QtCore/qtextcodec.h
// #include <qtextcodec.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QTextDecoder struct {
	*qtrt.CObject
}

func (this *QTextDecoder) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextDecoder) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextDecoderFromPointer(cthis unsafe.Pointer) *QTextDecoder {
	return &QTextDecoder{&qtrt.CObject{cthis}}
}
func (*QTextDecoder) NewFromPointer(cthis unsafe.Pointer) *QTextDecoder {
	return NewQTextDecoderFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtextcodec.h:158
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTextDecoder(const QTextCodec *)
func NewQTextDecoder(codec *QTextCodec /*777 const QTextCodec **/) *QTextDecoder {
	var convArg0 = codec.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextDecoderC2EPK10QTextCodec", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDecoderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextDecoder)
	return gothis
}

// /usr/include/qt/QtCore/qtextcodec.h:159
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextDecoder(const QTextCodec *, QTextCodec::ConversionFlags)
func NewQTextDecoder_1(codec *QTextCodec /*777 const QTextCodec **/, flags int) *QTextDecoder {
	var convArg0 = codec.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextDecoderC2EPK10QTextCodec6QFlagsINS0_14ConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDecoderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextDecoder)
	return gothis
}

// /usr/include/qt/QtCore/qtextcodec.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextDecoder()
func DeleteQTextDecoder(this *QTextDecoder) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextDecoderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtextcodec.h:161
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toUnicode(const char *, int)
func (this *QTextDecoder) ToUnicode(chars string, len int) string {
	var convArg0 = qtrt.CString(chars)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextDecoder9toUnicodeEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextcodec.h:162
// index:1
// Public Visibility=Default Availability=Available
// [8] QString toUnicode(const QByteArray &)
func (this *QTextDecoder) ToUnicode_1(ba *QByteArray) string {
	var convArg0 = ba.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextDecoder9toUnicodeERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextcodec.h:163
// index:2
// Public Visibility=Default Availability=Available
// [-2] void toUnicode(QString *, const char *, int)
func (this *QTextDecoder) ToUnicode_2(target string, chars string, len int) {
	var tmpArg0 = NewQString_5(target)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(chars)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextDecoder9toUnicodeEP7QStringPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, len)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFailure()
func (this *QTextDecoder) HasFailure() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextDecoder10hasFailureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
		qtrt.KeepMe()
	}
}

//  keep block end
