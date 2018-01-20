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

// /usr/include/qt/QtCore/qtextcodec.h:61
// index:0
// static
// QTextCodec * codecForName(const class QByteArray &)
func (this *QTextCodec) CodecForName(name unsafe.Pointer) {
	// 0: (name const QByteArray &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec12codecForNameERK10QByteArray", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTextCodec_CodecForName(name unsafe.Pointer) {
	// 0: (name const QByteArray &), (name)
	var nilthis *QTextCodec
	nilthis.CodecForName(name)
}

// /usr/include/qt/QtCore/qtextcodec.h:62
// index:1
// static inline
// QTextCodec * codecForName(const char *)
func (this *QTextCodec) CodecForName_1(name unsafe.Pointer) {
	// 1: (name const char *), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec12codecForNameEPKc", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTextCodec_CodecForName_1(name unsafe.Pointer) {
	// 1: (name const char *), (name)
	var nilthis *QTextCodec
	nilthis.CodecForName_1(name)
}

// /usr/include/qt/QtCore/qtextcodec.h:63
// index:0
// static
// QTextCodec * codecForMib(int)
func (this *QTextCodec) CodecForMib(mib int) {
	// 0: (mib int), (mib)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec11codecForMibEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTextCodec_CodecForMib(mib int) {
	// 0: (mib int), (mib)
	var nilthis *QTextCodec
	nilthis.CodecForMib(mib)
}

// /usr/include/qt/QtCore/qtextcodec.h:65
// index:0
// static
// QList<QByteArray> availableCodecs()
func (this *QTextCodec) AvailableCodecs() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec15availableCodecsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTextCodec_AvailableCodecs() {
	// 0: (), ()
	var nilthis *QTextCodec
	nilthis.AvailableCodecs()
}

// /usr/include/qt/QtCore/qtextcodec.h:66
// index:0
// static
// QList<int> availableMibs()
func (this *QTextCodec) AvailableMibs() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec13availableMibsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTextCodec_AvailableMibs() {
	// 0: (), ()
	var nilthis *QTextCodec
	nilthis.AvailableMibs()
}

// /usr/include/qt/QtCore/qtextcodec.h:68
// index:0
// static
// QTextCodec * codecForLocale()
func (this *QTextCodec) CodecForLocale() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec14codecForLocaleEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTextCodec_CodecForLocale() {
	// 0: (), ()
	var nilthis *QTextCodec
	nilthis.CodecForLocale()
}

// /usr/include/qt/QtCore/qtextcodec.h:69
// index:0
// static
// void setCodecForLocale(class QTextCodec *)
func (this *QTextCodec) SetCodecForLocale(c unsafe.Pointer) {
	// 0: (c QTextCodec *), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec17setCodecForLocaleEPS_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTextCodec_SetCodecForLocale(c unsafe.Pointer) {
	// 0: (c QTextCodec *), (c)
	var nilthis *QTextCodec
	nilthis.SetCodecForLocale(c)
}

// /usr/include/qt/QtCore/qtextcodec.h:75
// index:0
// static
// QTextCodec * codecForHtml(const class QByteArray &)
func (this *QTextCodec) CodecForHtml(ba unsafe.Pointer) {
	// 0: (ba const QByteArray &), (ba)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec12codecForHtmlERK10QByteArray", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTextCodec_CodecForHtml(ba unsafe.Pointer) {
	// 0: (ba const QByteArray &), (ba)
	var nilthis *QTextCodec
	nilthis.CodecForHtml(ba)
}

// /usr/include/qt/QtCore/qtextcodec.h:76
// index:1
// static
// QTextCodec * codecForHtml(const class QByteArray &, class QTextCodec *)
func (this *QTextCodec) CodecForHtml_1(ba unsafe.Pointer, defaultCodec unsafe.Pointer) {
	// 1: (ba const QByteArray &, defaultCodec QTextCodec *), (ba, defaultCodec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec12codecForHtmlERK10QByteArrayPS_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTextCodec_CodecForHtml_1(ba unsafe.Pointer, defaultCodec unsafe.Pointer) {
	// 1: (ba const QByteArray &, defaultCodec QTextCodec *), (ba, defaultCodec)
	var nilthis *QTextCodec
	nilthis.CodecForHtml_1(ba, defaultCodec)
}

// /usr/include/qt/QtCore/qtextcodec.h:78
// index:0
// static
// QTextCodec * codecForUtfText(const class QByteArray &)
func (this *QTextCodec) CodecForUtfText(ba unsafe.Pointer) {
	// 0: (ba const QByteArray &), (ba)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec15codecForUtfTextERK10QByteArray", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTextCodec_CodecForUtfText(ba unsafe.Pointer) {
	// 0: (ba const QByteArray &), (ba)
	var nilthis *QTextCodec
	nilthis.CodecForUtfText(ba)
}

// /usr/include/qt/QtCore/qtextcodec.h:79
// index:1
// static
// QTextCodec * codecForUtfText(const class QByteArray &, class QTextCodec *)
func (this *QTextCodec) CodecForUtfText_1(ba unsafe.Pointer, defaultCodec unsafe.Pointer) {
	// 1: (ba const QByteArray &, defaultCodec QTextCodec *), (ba, defaultCodec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodec15codecForUtfTextERK10QByteArrayPS_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTextCodec_CodecForUtfText_1(ba unsafe.Pointer, defaultCodec unsafe.Pointer) {
	// 1: (ba const QByteArray &, defaultCodec QTextCodec *), (ba, defaultCodec)
	var nilthis *QTextCodec
	nilthis.CodecForUtfText_1(ba, defaultCodec)
}

// /usr/include/qt/QtCore/qtextcodec.h:81
// index:0
// bool canEncode(class QChar)
func (this *QTextCodec) CanEncode(arg0 unsafe.Pointer) {
	// 0: (, QChar arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9canEncodeE5QChar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:83
// index:1
// bool canEncode(const class QString &)
func (this *QTextCodec) CanEncode_1(arg0 unsafe.Pointer) {
	// 1: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9canEncodeERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:85
// index:2
// bool canEncode(class QStringView)
func (this *QTextCodec) CanEncode_2(arg0 unsafe.Pointer) {
	// 2: (, QStringView arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9canEncodeE11QStringView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:87
// index:0
// QString toUnicode(const class QByteArray &)
func (this *QTextCodec) ToUnicode(arg0 unsafe.Pointer) {
	// 0: (, const QByteArray & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9toUnicodeERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:88
// index:1
// QString toUnicode(const char *)
func (this *QTextCodec) ToUnicode_1(chars unsafe.Pointer) {
	// 1: (, chars const char *), (chars)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9toUnicodeEPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), chars)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:114
// index:2
// inline
// QString toUnicode(const char *, int, struct QTextCodec::ConverterState *)
func (this *QTextCodec) ToUnicode_2(in unsafe.Pointer, length int, state unsafe.Pointer) {
	// 2: (, in const char *, length int, state QTextCodec::ConverterState *), (in, &length, state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec9toUnicodeEPKciPNS_14ConverterStateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), in, &length, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:90
// index:0
// QByteArray fromUnicode(const class QString &)
func (this *QTextCodec) FromUnicode(uc unsafe.Pointer) {
	// 0: (, uc const QString &), (uc)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec11fromUnicodeERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), uc)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:92
// index:1
// QByteArray fromUnicode(class QStringView)
func (this *QTextCodec) FromUnicode_1(uc unsafe.Pointer) {
	// 1: (, uc QStringView), (uc)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec11fromUnicodeE11QStringView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), uc)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:116
// index:2
// inline
// QByteArray fromUnicode(const class QChar *, int, struct QTextCodec::ConverterState *)
func (this *QTextCodec) FromUnicode_2(in unsafe.Pointer, length int, state unsafe.Pointer) {
	// 2: (, in const QChar *, length int, state QTextCodec::ConverterState *), (in, &length, state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec11fromUnicodeEPK5QChariPNS_14ConverterStateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), in, &length, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:119
// index:0
// QTextDecoder * makeDecoder(QTextCodec::ConversionFlags)
func (this *QTextCodec) MakeDecoder(flags int) {
	// 0: (, QFlags<QTextCodec::ConversionFlag> flags), (flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec11makeDecoderE6QFlagsINS_14ConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:120
// index:0
// QTextEncoder * makeEncoder(QTextCodec::ConversionFlags)
func (this *QTextCodec) MakeEncoder(flags int) {
	// 0: (, QFlags<QTextCodec::ConversionFlag> flags), (flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec11makeEncoderE6QFlagsINS_14ConversionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:122
// index:0
// pure virtual
// QByteArray name()
func (this *QTextCodec) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec4nameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:123
// index:0
// virtual
// QList<QByteArray> aliases()
func (this *QTextCodec) Aliases() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec7aliasesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:124
// index:0
// pure virtual
// int mibEnum()
func (this *QTextCodec) MibEnum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec7mibEnumEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:127
// index:0
// pure virtual
// QString convertToUnicode(const char *, int, struct QTextCodec::ConverterState *)
func (this *QTextCodec) ConvertToUnicode(in unsafe.Pointer, length int, state unsafe.Pointer) {
	// 0: (, in const char *, length int, state QTextCodec::ConverterState *), (in, &length, state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec16convertToUnicodeEPKciPNS_14ConverterStateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), in, &length, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:128
// index:0
// pure virtual
// QByteArray convertFromUnicode(const class QChar *, int, struct QTextCodec::ConverterState *)
func (this *QTextCodec) ConvertFromUnicode(in unsafe.Pointer, length int, state unsafe.Pointer) {
	// 0: (, in const QChar *, length int, state QTextCodec::ConverterState *), (in, &length, state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextCodec18convertFromUnicodeEPK5QChariPNS_14ConverterStateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), in, &length, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:130
// index:0
// void QTextCodec()
func NewQTextCodec() *QTextCodec {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodecC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextCodecFromPointer(cthis)
	return gothis
}
func NewQTextCodecFromPointer(cthis unsafe.Pointer) *QTextCodec {
	return &QTextCodec{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qtextcodec.h:131
// index:0
// virtual
// void ~QTextCodec()
func DeleteQTextCodec(*QTextCodec) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextCodecD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
