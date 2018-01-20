//  header block begin
// /usr/include/qt/QtCore/qtextcodec.h
// #include <qtextcodec.h>
// #include <QtCore>
package qtcore

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
type QTextCodec struct {
	*qtrt.CObject
}

func (this *QTextCodec) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQTextCodecFromPointer(cthis unsafe.Pointer) *QTextCodec {
	return &QTextCodec{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qtextcodec.h:61
// index:0
// Public static
// QTextCodec * codecForName(const class QByteArray &)
func (this *QTextCodec) CodecForName(name *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec12codecForNameERK10QByteArray", ffiqt.FFI_TYPE_POINTER, name)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTextCodec_CodecForName(name *QByteArray) {
	var nilthis *QTextCodec
	nilthis.CodecForName(name)
}

// /usr/include/qt/QtCore/qtextcodec.h:62
// index:1
// Public static inline
// QTextCodec * codecForName(const char *)
func (this *QTextCodec) CodecForName_1(name string) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec12codecForNameEPKc", ffiqt.FFI_TYPE_POINTER, name)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTextCodec_CodecForName_1(name string) {
	var nilthis *QTextCodec
	nilthis.CodecForName_1(name)
}

// /usr/include/qt/QtCore/qtextcodec.h:63
// index:0
// Public static
// QTextCodec * codecForMib(int)
func (this *QTextCodec) CodecForMib(mib int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec11codecForMibEi", ffiqt.FFI_TYPE_POINTER, mib)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTextCodec_CodecForMib(mib int) {
	var nilthis *QTextCodec
	nilthis.CodecForMib(mib)
}

// /usr/include/qt/QtCore/qtextcodec.h:65
// index:0
// Public static
// QList<QByteArray> availableCodecs()
func (this *QTextCodec) AvailableCodecs() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec15availableCodecsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTextCodec_AvailableCodecs() {
	var nilthis *QTextCodec
	nilthis.AvailableCodecs()
}

// /usr/include/qt/QtCore/qtextcodec.h:66
// index:0
// Public static
// QList<int> availableMibs()
func (this *QTextCodec) AvailableMibs() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec13availableMibsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTextCodec_AvailableMibs() {
	var nilthis *QTextCodec
	nilthis.AvailableMibs()
}

// /usr/include/qt/QtCore/qtextcodec.h:68
// index:0
// Public static
// QTextCodec * codecForLocale()
func (this *QTextCodec) CodecForLocale() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec14codecForLocaleEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTextCodec_CodecForLocale() {
	var nilthis *QTextCodec
	nilthis.CodecForLocale()
}

// /usr/include/qt/QtCore/qtextcodec.h:69
// index:0
// Public static
// void setCodecForLocale(class QTextCodec *)
func (this *QTextCodec) SetCodecForLocale(c unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec17setCodecForLocaleEPS_", ffiqt.FFI_TYPE_POINTER, c)
	gopp.ErrPrint(err, rv)
}
func QTextCodec_SetCodecForLocale(c unsafe.Pointer) {
	var nilthis *QTextCodec
	nilthis.SetCodecForLocale(c)
}

// /usr/include/qt/QtCore/qtextcodec.h:75
// index:0
// Public static
// QTextCodec * codecForHtml(const class QByteArray &)
func (this *QTextCodec) CodecForHtml(ba *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec12codecForHtmlERK10QByteArray", ffiqt.FFI_TYPE_POINTER, ba)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTextCodec_CodecForHtml(ba *QByteArray) {
	var nilthis *QTextCodec
	nilthis.CodecForHtml(ba)
}

// /usr/include/qt/QtCore/qtextcodec.h:76
// index:1
// Public static
// QTextCodec * codecForHtml(const class QByteArray &, class QTextCodec *)
func (this *QTextCodec) CodecForHtml_1(ba *QByteArray, defaultCodec unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec12codecForHtmlERK10QByteArrayPS_", ffiqt.FFI_TYPE_POINTER, ba, defaultCodec)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTextCodec_CodecForHtml_1(ba *QByteArray, defaultCodec unsafe.Pointer) {
	var nilthis *QTextCodec
	nilthis.CodecForHtml_1(ba, defaultCodec)
}

// /usr/include/qt/QtCore/qtextcodec.h:78
// index:0
// Public static
// QTextCodec * codecForUtfText(const class QByteArray &)
func (this *QTextCodec) CodecForUtfText(ba *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec15codecForUtfTextERK10QByteArray", ffiqt.FFI_TYPE_POINTER, ba)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTextCodec_CodecForUtfText(ba *QByteArray) {
	var nilthis *QTextCodec
	nilthis.CodecForUtfText(ba)
}

// /usr/include/qt/QtCore/qtextcodec.h:79
// index:1
// Public static
// QTextCodec * codecForUtfText(const class QByteArray &, class QTextCodec *)
func (this *QTextCodec) CodecForUtfText_1(ba *QByteArray, defaultCodec unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec15codecForUtfTextERK10QByteArrayPS_", ffiqt.FFI_TYPE_POINTER, ba, defaultCodec)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTextCodec_CodecForUtfText_1(ba *QByteArray, defaultCodec unsafe.Pointer) {
	var nilthis *QTextCodec
	nilthis.CodecForUtfText_1(ba, defaultCodec)
}

// /usr/include/qt/QtCore/qtextcodec.h:81
// index:0
// Public
// bool canEncode(class QChar)
func (this *QTextCodec) CanEncode(arg0 *QChar) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9canEncodeE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:83
// index:1
// Public
// bool canEncode(const class QString &)
func (this *QTextCodec) CanEncode_1(arg0 *QString) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9canEncodeERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:85
// index:2
// Public
// bool canEncode(class QStringView)
func (this *QTextCodec) CanEncode_2(arg0 *QStringView) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9canEncodeE11QStringView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:87
// index:0
// Public
// QString toUnicode(const class QByteArray &)
func (this *QTextCodec) ToUnicode(arg0 *QByteArray) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9toUnicodeERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:88
// index:1
// Public
// QString toUnicode(const char *)
func (this *QTextCodec) ToUnicode_1(chars string) interface{} {
	var convArg0 = qtrt.CString(chars)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9toUnicodeEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:114
// index:2
// Public inline
// QString toUnicode(const char *, int, struct QTextCodec::ConverterState *)
func (this *QTextCodec) ToUnicode_2(in string, length int, state unsafe.Pointer) interface{} {
	var convArg0 = qtrt.CString(in)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9toUnicodeEPKciPNS_14ConverterStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &length, state)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:90
// index:0
// Public
// QByteArray fromUnicode(const class QString &)
func (this *QTextCodec) FromUnicode(uc *QString) interface{} {
	var convArg0 = uc.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec11fromUnicodeERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:92
// index:1
// Public
// QByteArray fromUnicode(class QStringView)
func (this *QTextCodec) FromUnicode_1(uc *QStringView) interface{} {
	var convArg0 = uc.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec11fromUnicodeE11QStringView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:116
// index:2
// Public inline
// QByteArray fromUnicode(const class QChar *, int, struct QTextCodec::ConverterState *)
func (this *QTextCodec) FromUnicode_2(in unsafe.Pointer, length int, state unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec11fromUnicodeEPK5QChariPNS_14ConverterStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), in, &length, state)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:122
// index:0
// Public pure virtual
// QByteArray name()
func (this *QTextCodec) Name() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:123
// index:0
// Public virtual
// QList<QByteArray> aliases()
func (this *QTextCodec) Aliases() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec7aliasesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:124
// index:0
// Public pure virtual
// int mibEnum()
func (this *QTextCodec) MibEnum() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec7mibEnumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:127
// index:0
// Protected pure virtual
// QString convertToUnicode(const char *, int, struct QTextCodec::ConverterState *)
func (this *QTextCodec) ConvertToUnicode(in string, length int, state unsafe.Pointer) interface{} {
	var convArg0 = qtrt.CString(in)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec16convertToUnicodeEPKciPNS_14ConverterStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &length, state)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:128
// index:0
// Protected pure virtual
// QByteArray convertFromUnicode(const class QChar *, int, struct QTextCodec::ConverterState *)
func (this *QTextCodec) ConvertFromUnicode(in unsafe.Pointer, length int, state unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec18convertFromUnicodeEPK5QChariPNS_14ConverterStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), in, &length, state)
	gopp.ErrPrint(err, rv)
	return rv
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

//  body block end
