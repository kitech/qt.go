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
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QTextDecoder struct {
	*qtrt.CObject
}
type QTextDecoder_ITF interface {
	QTextDecoder_PTR() *QTextDecoder
}

func (ptr *QTextDecoder) QTextDecoder_PTR() *QTextDecoder { return ptr }

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

/*

 */
func (*QTextDecoder) NewForInherit(codec QTextCodec_ITF /*777 const QTextCodec **/) *QTextDecoder {
	return NewQTextDecoder(codec)
}
func NewQTextDecoder(codec QTextCodec_ITF /*777 const QTextCodec **/) *QTextDecoder {
	var convArg0 unsafe.Pointer
	if codec != nil && codec.QTextCodec_PTR() != nil {
		convArg0 = codec.QTextCodec_PTR().GetCthis()
	}
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

/*

 */
func (*QTextDecoder) NewForInherit1(codec QTextCodec_ITF /*777 const QTextCodec **/, flags int) *QTextDecoder {
	return NewQTextDecoder1(codec, flags)
}
func NewQTextDecoder1(codec QTextCodec_ITF /*777 const QTextCodec **/, flags int) *QTextDecoder {
	var convArg0 unsafe.Pointer
	if codec != nil && codec.QTextCodec_PTR() != nil {
		convArg0 = codec.QTextCodec_PTR().GetCthis()
	}
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

/*

 */
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

/*
Converts a from the encoding of this codec to Unicode, and returns the result in a QString.
*/
func (this *QTextDecoder) ToUnicode(chars string, len_ int) string {
	var convArg0 = qtrt.CString(chars)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextDecoder9toUnicodeEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextcodec.h:162
// index:1
// Public Visibility=Default Availability=Available
// [8] QString toUnicode(const QByteArray &)

/*
Converts a from the encoding of this codec to Unicode, and returns the result in a QString.
*/
func (this *QTextDecoder) ToUnicode1(ba QByteArray_ITF) string {
	var convArg0 unsafe.Pointer
	if ba != nil && ba.QByteArray_PTR() != nil {
		convArg0 = ba.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextDecoder9toUnicodeERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextcodec.h:163
// index:2
// Public Visibility=Default Availability=Available
// [-2] void toUnicode(QString *, const char *, int)

/*
Converts a from the encoding of this codec to Unicode, and returns the result in a QString.
*/
func (this *QTextDecoder) ToUnicode2(target string, chars string, len_ int) {
	var tmpArg0 = NewQString5(target)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(chars)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextDecoder9toUnicodeEP7QStringPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, len_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFailure() const

/*

 */
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
