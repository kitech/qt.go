// +build !minimal

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
// extern C begin: 22
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
type QTextEncoder struct {
	*qtrt.CObject
}
type QTextEncoder_ITF interface {
	QTextEncoder_PTR() *QTextEncoder
}

func (ptr *QTextEncoder) QTextEncoder_PTR() *QTextEncoder { return ptr }

func (this *QTextEncoder) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextEncoder) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextEncoderFromPointer(cthis unsafe.Pointer) *QTextEncoder {
	return &QTextEncoder{&qtrt.CObject{cthis}}
}
func (*QTextEncoder) NewFromPointer(cthis unsafe.Pointer) *QTextEncoder {
	return NewQTextEncoderFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtextcodec.h:140
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTextEncoder(const QTextCodec *)

/*

 */
func (*QTextEncoder) NewForInherit(codec QTextCodec_ITF /*777 const QTextCodec **/) *QTextEncoder {
	return NewQTextEncoder(codec)
}
func NewQTextEncoder(codec QTextCodec_ITF /*777 const QTextCodec **/) *QTextEncoder {
	var convArg0 unsafe.Pointer
	if codec != nil && codec.QTextCodec_PTR() != nil {
		convArg0 = codec.QTextCodec_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextEncoderC2EPK10QTextCodec", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextEncoderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextEncoder)
	return gothis
}

// /usr/include/qt/QtCore/qtextcodec.h:141
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextEncoder(const QTextCodec *, QTextCodec::ConversionFlags)

/*

 */
func (*QTextEncoder) NewForInherit1(codec QTextCodec_ITF /*777 const QTextCodec **/, flags int) *QTextEncoder {
	return NewQTextEncoder1(codec, flags)
}
func NewQTextEncoder1(codec QTextCodec_ITF /*777 const QTextCodec **/, flags int) *QTextEncoder {
	var convArg0 unsafe.Pointer
	if codec != nil && codec.QTextCodec_PTR() != nil {
		convArg0 = codec.QTextCodec_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextEncoderC2EPK10QTextCodec6QFlagsINS0_14ConversionFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextEncoderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextEncoder)
	return gothis
}

// /usr/include/qt/QtCore/qtextcodec.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextEncoder()

/*

 */
func DeleteQTextEncoder(this *QTextEncoder) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextEncoderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtextcodec.h:144
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray fromUnicode(const QString &)

/*
Converts str from Unicode to the encoding of this codec, and returns the result in a QByteArray.
*/
func (this *QTextEncoder) FromUnicode(str string) *QByteArray /*123*/ {
	var tmpArg0 = NewQString5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextEncoder11fromUnicodeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:146
// index:1
// Public Visibility=Default Availability=Available
// [8] QByteArray fromUnicode(QStringView)

/*
Converts str from Unicode to the encoding of this codec, and returns the result in a QByteArray.
*/
func (this *QTextEncoder) FromUnicode1(str QStringView_ITF /*123*/) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringView_PTR() != nil {
		convArg0 = str.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextEncoder11fromUnicodeE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:147
// index:2
// Public Visibility=Default Availability=Available
// [8] QByteArray fromUnicode(const QChar *, int)

/*
Converts str from Unicode to the encoding of this codec, and returns the result in a QByteArray.
*/
func (this *QTextEncoder) FromUnicode2(uc QChar_ITF /*777 const QChar **/, len_ int) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if uc != nil && uc.QChar_PTR() != nil {
		convArg0 = uc.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextEncoder11fromUnicodeEPK5QChari", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFailure() const

/*

 */
func (this *QTextEncoder) HasFailure() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextEncoder10hasFailureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

func init_unused_10571() {
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
