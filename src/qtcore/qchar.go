//  header block begin
// /usr/include/qt/QtCore/qchar.h
// #include <qchar.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QChar struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qchar.h:81
// index:0
// inline
// void QChar()
func NewQChar() *QChar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QChar{cthis}
}

// /usr/include/qt/QtCore/qchar.h:82
// index:1
// inline
// void QChar(ushort)
func NewQChar_1(rc uint16) *QChar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Et", ffiqt.FFI_TYPE_VOID, cthis, &rc)
	gopp.ErrPrint(err, rv)
	return &QChar{cthis}
}

// /usr/include/qt/QtCore/qchar.h:83
// index:2
// inline
// void QChar(uchar, uchar)
func NewQChar_2(c byte, r byte) *QChar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Ehh", ffiqt.FFI_TYPE_VOID, cthis, &c, &r)
	gopp.ErrPrint(err, rv)
	return &QChar{cthis}
}

// /usr/include/qt/QtCore/qchar.h:84
// index:3
// inline
// void QChar(short)
func NewQChar_3(rc int16) *QChar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Es", ffiqt.FFI_TYPE_VOID, cthis, &rc)
	gopp.ErrPrint(err, rv)
	return &QChar{cthis}
}

// /usr/include/qt/QtCore/qchar.h:85
// index:4
// inline
// void QChar(uint)
func NewQChar_4(rc uint) *QChar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Ej", ffiqt.FFI_TYPE_VOID, cthis, &rc)
	gopp.ErrPrint(err, rv)
	return &QChar{cthis}
}

// /usr/include/qt/QtCore/qchar.h:86
// index:5
// inline
// void QChar(int)
func NewQChar_5(rc int) *QChar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &rc)
	gopp.ErrPrint(err, rv)
	return &QChar{cthis}
}

// /usr/include/qt/QtCore/qchar.h:87
// index:6
// inline
// void QChar(enum QChar::SpecialCharacter)
func NewQChar_6(s int) *QChar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2ENS_16SpecialCharacterE", ffiqt.FFI_TYPE_VOID, cthis, &s)
	gopp.ErrPrint(err, rv)
	return &QChar{cthis}
}

// /usr/include/qt/QtCore/qchar.h:88
// index:7
// inline
// void QChar(struct QLatin1Char)
func NewQChar_7(ch unsafe.Pointer) *QChar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2E11QLatin1Char", ffiqt.FFI_TYPE_VOID, cthis, ch)
	gopp.ErrPrint(err, rv)
	return &QChar{cthis}
}

// /usr/include/qt/QtCore/qchar.h:90
// index:8
// inline
// void QChar(char16_t)
func NewQChar_8(ch int16) *QChar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2EDs", ffiqt.FFI_TYPE_VOID, cthis, &ch)
	gopp.ErrPrint(err, rv)
	return &QChar{cthis}
}

// /usr/include/qt/QtCore/qchar.h:98
// index:9
// inline
// void QChar(char)
func NewQChar_9(c byte) *QChar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Ec", ffiqt.FFI_TYPE_VOID, cthis, &c)
	gopp.ErrPrint(err, rv)
	return &QChar{cthis}
}

// /usr/include/qt/QtCore/qchar.h:100
// index:10
// inline
// void QChar(uchar)
func NewQChar_10(c byte) *QChar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Eh", ffiqt.FFI_TYPE_VOID, cthis, &c)
	gopp.ErrPrint(err, rv)
	return &QChar{cthis}
}

// /usr/include/qt/QtCore/qchar.h:394
// index:0
// inline
// QChar::Category category()
func (this *QChar) Category() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar8categoryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:504
// index:1
// static
// QChar::Category category(uint)
func (this *QChar) Category_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar8categoryEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_Category_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.Category_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:395
// index:0
// inline
// QChar::Direction direction()
func (this *QChar) Direction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar9directionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:505
// index:1
// static
// QChar::Direction direction(uint)
func (this *QChar) Direction_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar9directionEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_Direction_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.Direction_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:396
// index:0
// inline
// QChar::JoiningType joiningType()
func (this *QChar) JoiningType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar11joiningTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:506
// index:1
// static
// QChar::JoiningType joiningType(uint)
func (this *QChar) JoiningType_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar11joiningTypeEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_JoiningType_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.JoiningType_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:398
// index:0
// inline
// QChar::Joining joining()
func (this *QChar) Joining() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7joiningEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:508
// index:1
// static
// QChar::Joining joining(uint)
func (this *QChar) Joining_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7joiningEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_Joining_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.Joining_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:411
// index:0
// inline
// unsigned char combiningClass()
func (this *QChar) CombiningClass() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar14combiningClassEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:510
// index:1
// static
// unsigned char combiningClass(uint)
func (this *QChar) CombiningClass_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar14combiningClassEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_CombiningClass_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.CombiningClass_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:413
// index:0
// inline
// QChar mirroredChar()
func (this *QChar) MirroredChar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar12mirroredCharEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:512
// index:1
// static
// uint mirroredChar(uint)
func (this *QChar) MirroredChar_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar12mirroredCharEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_MirroredChar_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.MirroredChar_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:414
// index:0
// inline
// bool hasMirrored()
func (this *QChar) HasMirrored() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar11hasMirroredEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:513
// index:1
// static
// bool hasMirrored(uint)
func (this *QChar) HasMirrored_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar11hasMirroredEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_HasMirrored_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.HasMirrored_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:416
// index:0
// QString decomposition()
func (this *QChar) Decomposition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar13decompositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:515
// index:1
// static
// QString decomposition(uint)
func (this *QChar) Decomposition_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar13decompositionEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_Decomposition_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.Decomposition_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:417
// index:0
// inline
// QChar::Decomposition decompositionTag()
func (this *QChar) DecompositionTag() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar16decompositionTagEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:516
// index:1
// static
// QChar::Decomposition decompositionTag(uint)
func (this *QChar) DecompositionTag_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar16decompositionTagEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_DecompositionTag_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.DecompositionTag_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:419
// index:0
// inline
// int digitValue()
func (this *QChar) DigitValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar10digitValueEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:518
// index:1
// static
// int digitValue(uint)
func (this *QChar) DigitValue_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar10digitValueEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_DigitValue_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.DigitValue_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:420
// index:0
// inline
// QChar toLower()
func (this *QChar) ToLower() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7toLowerEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:519
// index:1
// static
// uint toLower(uint)
func (this *QChar) ToLower_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7toLowerEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_ToLower_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.ToLower_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:421
// index:0
// inline
// QChar toUpper()
func (this *QChar) ToUpper() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7toUpperEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:520
// index:1
// static
// uint toUpper(uint)
func (this *QChar) ToUpper_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7toUpperEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_ToUpper_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.ToUpper_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:422
// index:0
// inline
// QChar toTitleCase()
func (this *QChar) ToTitleCase() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar11toTitleCaseEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:521
// index:1
// static
// uint toTitleCase(uint)
func (this *QChar) ToTitleCase_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar11toTitleCaseEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_ToTitleCase_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.ToTitleCase_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:423
// index:0
// inline
// QChar toCaseFolded()
func (this *QChar) ToCaseFolded() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar12toCaseFoldedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:522
// index:1
// static
// uint toCaseFolded(uint)
func (this *QChar) ToCaseFolded_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar12toCaseFoldedEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_ToCaseFolded_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.ToCaseFolded_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:425
// index:0
// inline
// QChar::Script script()
func (this *QChar) Script() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar6scriptEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:524
// index:1
// static
// QChar::Script script(uint)
func (this *QChar) Script_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar6scriptEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_Script_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.Script_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:427
// index:0
// inline
// QChar::UnicodeVersion unicodeVersion()
func (this *QChar) UnicodeVersion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar14unicodeVersionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:526
// index:1
// static
// QChar::UnicodeVersion unicodeVersion(uint)
func (this *QChar) UnicodeVersion_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar14unicodeVersionEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_UnicodeVersion_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.UnicodeVersion_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:432
// index:0
// inline
// char toLatin1()
func (this *QChar) ToLatin1() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar8toLatin1Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:433
// index:0
// inline
// ushort unicode()
func (this *QChar) Unicode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7unicodeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:434
// index:1
// inline
// ushort & unicode()
func (this *QChar) Unicode_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7unicodeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:440
// index:0
// static inline
// QChar fromLatin1(char)
func (this *QChar) FromLatin1(c byte) {
	// 0: (char c), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar10fromLatin1Ec", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_FromLatin1(c byte) {
	// 0: (char c), (c)
	var nilthis *QChar
	nilthis.FromLatin1(c)
}

// /usr/include/qt/QtCore/qchar.h:442
// index:0
// inline
// bool isNull()
func (this *QChar) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar6isNullEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:444
// index:0
// inline
// bool isPrint()
func (this *QChar) IsPrint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isPrintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:530
// index:1
// static
// bool isPrint(uint)
func (this *QChar) IsPrint_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7isPrintEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsPrint_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsPrint_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:445
// index:0
// inline
// bool isSpace()
func (this *QChar) IsSpace() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isSpaceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:531
// index:1
// static inline
// bool isSpace(uint)
func (this *QChar) IsSpace_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7isSpaceEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsSpace_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsSpace_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:446
// index:0
// inline
// bool isMark()
func (this *QChar) IsMark() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar6isMarkEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:537
// index:1
// static
// bool isMark(uint)
func (this *QChar) IsMark_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar6isMarkEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsMark_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsMark_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:447
// index:0
// inline
// bool isPunct()
func (this *QChar) IsPunct() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isPunctEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:538
// index:1
// static
// bool isPunct(uint)
func (this *QChar) IsPunct_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7isPunctEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsPunct_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsPunct_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:448
// index:0
// inline
// bool isSymbol()
func (this *QChar) IsSymbol() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar8isSymbolEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:539
// index:1
// static
// bool isSymbol(uint)
func (this *QChar) IsSymbol_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar8isSymbolEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsSymbol_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsSymbol_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:449
// index:0
// inline
// bool isLetter()
func (this *QChar) IsLetter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar8isLetterEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:540
// index:1
// static inline
// bool isLetter(uint)
func (this *QChar) IsLetter_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar8isLetterEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsLetter_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsLetter_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:450
// index:0
// inline
// bool isNumber()
func (this *QChar) IsNumber() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar8isNumberEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:545
// index:1
// static inline
// bool isNumber(uint)
func (this *QChar) IsNumber_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar8isNumberEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsNumber_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsNumber_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:451
// index:0
// inline
// bool isLetterOrNumber()
func (this *QChar) IsLetterOrNumber() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar16isLetterOrNumberEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:547
// index:1
// static inline
// bool isLetterOrNumber(uint)
func (this *QChar) IsLetterOrNumber_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar16isLetterOrNumberEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsLetterOrNumber_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsLetterOrNumber_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:452
// index:0
// inline
// bool isDigit()
func (this *QChar) IsDigit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isDigitEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:553
// index:1
// static inline
// bool isDigit(uint)
func (this *QChar) IsDigit_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7isDigitEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsDigit_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsDigit_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:453
// index:0
// inline
// bool isLower()
func (this *QChar) IsLower() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isLowerEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:555
// index:1
// static inline
// bool isLower(uint)
func (this *QChar) IsLower_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7isLowerEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsLower_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsLower_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:454
// index:0
// inline
// bool isUpper()
func (this *QChar) IsUpper() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isUpperEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:557
// index:1
// static inline
// bool isUpper(uint)
func (this *QChar) IsUpper_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7isUpperEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsUpper_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsUpper_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:455
// index:0
// inline
// bool isTitleCase()
func (this *QChar) IsTitleCase() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar11isTitleCaseEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:559
// index:1
// static inline
// bool isTitleCase(uint)
func (this *QChar) IsTitleCase_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar11isTitleCaseEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsTitleCase_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsTitleCase_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:457
// index:0
// inline
// bool isNonCharacter()
func (this *QChar) IsNonCharacter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar14isNonCharacterEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:467
// index:1
// static inline
// bool isNonCharacter(uint)
func (this *QChar) IsNonCharacter_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar14isNonCharacterEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsNonCharacter_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsNonCharacter_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:458
// index:0
// inline
// bool isHighSurrogate()
func (this *QChar) IsHighSurrogate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar15isHighSurrogateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:471
// index:1
// static inline
// bool isHighSurrogate(uint)
func (this *QChar) IsHighSurrogate_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar15isHighSurrogateEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsHighSurrogate_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsHighSurrogate_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:459
// index:0
// inline
// bool isLowSurrogate()
func (this *QChar) IsLowSurrogate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar14isLowSurrogateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:475
// index:1
// static inline
// bool isLowSurrogate(uint)
func (this *QChar) IsLowSurrogate_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar14isLowSurrogateEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsLowSurrogate_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsLowSurrogate_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:460
// index:0
// inline
// bool isSurrogate()
func (this *QChar) IsSurrogate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar11isSurrogateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:479
// index:1
// static inline
// bool isSurrogate(uint)
func (this *QChar) IsSurrogate_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar11isSurrogateEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_IsSurrogate_1(ucs4 uint) {
	// 1: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.IsSurrogate_1(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:462
// index:0
// inline
// uchar cell()
func (this *QChar) Cell() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar4cellEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:463
// index:0
// inline
// uchar row()
func (this *QChar) Row() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar3rowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:464
// index:0
// inline
// void setCell(uchar)
func (this *QChar) SetCell(acell byte) {
	// 0: (, uchar acell), (&acell)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7setCellEh", ffiqt.FFI_TYPE_VOID, this.cthis, &acell)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:465
// index:0
// inline
// void setRow(uchar)
func (this *QChar) SetRow(arow byte) {
	// 0: (, uchar arow), (&arow)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar6setRowEh", ffiqt.FFI_TYPE_VOID, this.cthis, &arow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:483
// index:0
// static inline
// bool requiresSurrogates(uint)
func (this *QChar) RequiresSurrogates(ucs4 uint) {
	// 0: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar18requiresSurrogatesEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_RequiresSurrogates(ucs4 uint) {
	// 0: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.RequiresSurrogates(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:487
// index:0
// static inline
// uint surrogateToUcs4(ushort, ushort)
func (this *QChar) SurrogateToUcs4(high uint16, low uint16) {
	// 0: (ushort high, ushort low), (high, low)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar15surrogateToUcs4Ett", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_SurrogateToUcs4(high uint16, low uint16) {
	// 0: (ushort high, ushort low), (high, low)
	var nilthis *QChar
	nilthis.SurrogateToUcs4(high, low)
}

// /usr/include/qt/QtCore/qchar.h:491
// index:1
// static inline
// uint surrogateToUcs4(class QChar, class QChar)
func (this *QChar) SurrogateToUcs4_1(high unsafe.Pointer, low unsafe.Pointer) {
	// 1: (QChar high, QChar low), (high, low)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar15surrogateToUcs4ES_S_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_SurrogateToUcs4_1(high unsafe.Pointer, low unsafe.Pointer) {
	// 1: (QChar high, QChar low), (high, low)
	var nilthis *QChar
	nilthis.SurrogateToUcs4_1(high, low)
}

// /usr/include/qt/QtCore/qchar.h:495
// index:0
// static inline
// ushort highSurrogate(uint)
func (this *QChar) HighSurrogate(ucs4 uint) {
	// 0: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar13highSurrogateEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_HighSurrogate(ucs4 uint) {
	// 0: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.HighSurrogate(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:499
// index:0
// static inline
// ushort lowSurrogate(uint)
func (this *QChar) LowSurrogate(ucs4 uint) {
	// 0: (uint ucs4), (ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar12lowSurrogateEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_LowSurrogate(ucs4 uint) {
	// 0: (uint ucs4), (ucs4)
	var nilthis *QChar
	nilthis.LowSurrogate(ucs4)
}

// /usr/include/qt/QtCore/qchar.h:528
// index:0
// static
// QChar::UnicodeVersion currentUnicodeVersion()
func (this *QChar) CurrentUnicodeVersion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar21currentUnicodeVersionEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QChar_CurrentUnicodeVersion() {
	// 0: (), ()
	var nilthis *QChar
	nilthis.CurrentUnicodeVersion()
}

//  body block end
