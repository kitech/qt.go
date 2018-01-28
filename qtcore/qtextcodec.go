package qtcore

// /usr/include/qt/QtCore/qtextcodec.h
// #include <qtextcodec.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
type QTextCodec struct {
	*qtrt.CObject
}

func (this *QTextCodec) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextCodec) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQTextCodecFromPointer(cthis unsafe.Pointer) *QTextCodec {
	return &QTextCodec{&qtrt.CObject{cthis}}
}
func (*QTextCodec) NewFromPointer(cthis unsafe.Pointer) *QTextCodec {
	return NewQTextCodecFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtextcodec.h:61
// index:0
// Public static
// QTextCodec * codecForName(const QByteArray &)
func (this *QTextCodec) CodecForName(name *QByteArray) *QTextCodec /*777 QTextCodec **/ {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec12codecForNameERK10QByteArray", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QTextCodec_CodecForName(name *QByteArray) *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForName(name)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:62
// index:1
// Public static inline
// QTextCodec * codecForName(const char *)
func (this *QTextCodec) CodecForName_1(name string) *QTextCodec /*777 QTextCodec **/ {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec12codecForNameEPKc", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QTextCodec_CodecForName_1(name string) *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForName_1(name)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:63
// index:0
// Public static
// QTextCodec * codecForMib(int)
func (this *QTextCodec) CodecForMib(mib int) *QTextCodec /*777 QTextCodec **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec11codecForMibEi", ffiqt.FFI_TYPE_POINTER, mib)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QTextCodec_CodecForMib(mib int) *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForMib(mib)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:68
// index:0
// Public static
// QTextCodec * codecForLocale()
func (this *QTextCodec) CodecForLocale() *QTextCodec /*777 QTextCodec **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec14codecForLocaleEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QTextCodec_CodecForLocale() *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForLocale()
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:69
// index:0
// Public static
// void setCodecForLocale(QTextCodec *)
func (this *QTextCodec) SetCodecForLocale(c *QTextCodec /*777 QTextCodec **/) {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec17setCodecForLocaleEPS_", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QTextCodec_SetCodecForLocale(c *QTextCodec /*777 QTextCodec **/) {
	var nilthis *QTextCodec
	nilthis.SetCodecForLocale(c)
}

// /usr/include/qt/QtCore/qtextcodec.h:75
// index:0
// Public static
// QTextCodec * codecForHtml(const QByteArray &)
func (this *QTextCodec) CodecForHtml(ba *QByteArray) *QTextCodec /*777 QTextCodec **/ {
	var convArg0 = ba.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec12codecForHtmlERK10QByteArray", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QTextCodec_CodecForHtml(ba *QByteArray) *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForHtml(ba)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:76
// index:1
// Public static
// QTextCodec * codecForHtml(const QByteArray &, QTextCodec *)
func (this *QTextCodec) CodecForHtml_1(ba *QByteArray, defaultCodec *QTextCodec /*777 QTextCodec **/) *QTextCodec /*777 QTextCodec **/ {
	var convArg0 = ba.GetCthis()
	var convArg1 = defaultCodec.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec12codecForHtmlERK10QByteArrayPS_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QTextCodec_CodecForHtml_1(ba *QByteArray, defaultCodec *QTextCodec /*777 QTextCodec **/) *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForHtml_1(ba, defaultCodec)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:78
// index:0
// Public static
// QTextCodec * codecForUtfText(const QByteArray &)
func (this *QTextCodec) CodecForUtfText(ba *QByteArray) *QTextCodec /*777 QTextCodec **/ {
	var convArg0 = ba.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec15codecForUtfTextERK10QByteArray", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QTextCodec_CodecForUtfText(ba *QByteArray) *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForUtfText(ba)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:79
// index:1
// Public static
// QTextCodec * codecForUtfText(const QByteArray &, QTextCodec *)
func (this *QTextCodec) CodecForUtfText_1(ba *QByteArray, defaultCodec *QTextCodec /*777 QTextCodec **/) *QTextCodec /*777 QTextCodec **/ {
	var convArg0 = ba.GetCthis()
	var convArg1 = defaultCodec.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec15codecForUtfTextERK10QByteArrayPS_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QTextCodec_CodecForUtfText_1(ba *QByteArray, defaultCodec *QTextCodec /*777 QTextCodec **/) *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForUtfText_1(ba, defaultCodec)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:81
// index:0
// Public
// bool canEncode(QChar)
func (this *QTextCodec) CanEncode(arg0 *QChar /*123*/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9canEncodeE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtextcodec.h:83
// index:1
// Public
// bool canEncode(const QString &)
func (this *QTextCodec) CanEncode_1(arg0 *QString) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9canEncodeERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtextcodec.h:85
// index:2
// Public
// bool canEncode(QStringView)
func (this *QTextCodec) CanEncode_2(arg0 *QStringView /*123*/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9canEncodeE11QStringView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtextcodec.h:87
// index:0
// Public
// QString toUnicode(const QByteArray &)
func (this *QTextCodec) ToUnicode(arg0 *QByteArray) *QString /*123*/ {
	var convArg0 = arg0.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9toUnicodeERK10QByteArray", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:88
// index:1
// Public
// QString toUnicode(const char *)
func (this *QTextCodec) ToUnicode_1(chars string) *QString /*123*/ {
	var convArg0 = qtrt.CString(chars)
	defer qtrt.FreeMem(convArg0)
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9toUnicodeEPKc", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:90
// index:0
// Public
// QByteArray fromUnicode(const QString &)
func (this *QTextCodec) FromUnicode(uc *QString) *QByteArray /*123*/ {
	var convArg0 = uc.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec11fromUnicodeERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:92
// index:1
// Public
// QByteArray fromUnicode(QStringView)
func (this *QTextCodec) FromUnicode_1(uc *QStringView /*123*/) *QByteArray /*123*/ {
	var convArg0 = uc.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec11fromUnicodeE11QStringView", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:119
// index:0
// Public
// QTextDecoder * makeDecoder(QTextCodec::ConversionFlags)
func (this *QTextCodec) MakeDecoder(flags int) *QTextDecoder /*777 QTextDecoder **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec11makeDecoderE6QFlagsINS_14ConversionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextDecoderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:120
// index:0
// Public
// QTextEncoder * makeEncoder(QTextCodec::ConversionFlags)
func (this *QTextCodec) MakeEncoder(flags int) *QTextEncoder /*777 QTextEncoder **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec11makeEncoderE6QFlagsINS_14ConversionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextEncoderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:122
// index:0
// Public pure virtual
// QByteArray name()
func (this *QTextCodec) Name() *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec4nameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:124
// index:0
// Public pure virtual
// int mibEnum()
func (this *QTextCodec) MibEnum() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec7mibEnumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextcodec.h:130
// index:0
// Protected
// void QTextCodec()
func NewQTextCodec() *QTextCodec {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodecC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextCodecFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextcodec.h:131
// index:0
// Protected virtual
// void ~QTextCodec()
func DeleteQTextCodec(*QTextCodec) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodecD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

type QTextCodec__ConversionFlag = int

const QTextCodec__DefaultConversion QTextCodec__ConversionFlag = 0
const QTextCodec__ConvertInvalidToNull QTextCodec__ConversionFlag = 2147483648
const QTextCodec__IgnoreHeader QTextCodec__ConversionFlag = 1
const QTextCodec__FreeFunction QTextCodec__ConversionFlag = 2

//  body block end
