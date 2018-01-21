package qtcore

// /usr/include/qt/QtCore/qchar.h
// #include <qchar.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
	*qtrt.CObject
}

func (this *QChar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQCharFromPointer(cthis unsafe.Pointer) *QChar {
	return &QChar{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qchar.h:81
// index:0
// Public inline
// void QChar()
func NewQChar() *QChar {
	cthis := qtrt.Calloc(1, 256) // 2
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:82
// index:1
// Public inline
// void QChar(ushort)
func NewQChar_1(rc uint16) *QChar {
	cthis := qtrt.Calloc(1, 256) // 2
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Et", ffiqt.FFI_TYPE_VOID, cthis, &rc)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:83
// index:2
// Public inline
// void QChar(uchar, uchar)
func NewQChar_2(c byte, r byte) *QChar {
	cthis := qtrt.Calloc(1, 256) // 2
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Ehh", ffiqt.FFI_TYPE_VOID, cthis, &c, &r)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:84
// index:3
// Public inline
// void QChar(short)
func NewQChar_3(rc int16) *QChar {
	cthis := qtrt.Calloc(1, 256) // 2
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Es", ffiqt.FFI_TYPE_VOID, cthis, &rc)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:85
// index:4
// Public inline
// void QChar(uint)
func NewQChar_4(rc uint) *QChar {
	cthis := qtrt.Calloc(1, 256) // 2
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Ej", ffiqt.FFI_TYPE_VOID, cthis, &rc)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:86
// index:5
// Public inline
// void QChar(int)
func NewQChar_5(rc int) *QChar {
	cthis := qtrt.Calloc(1, 256) // 2
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &rc)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:87
// index:6
// Public inline
// void QChar(enum QChar::SpecialCharacter)
func NewQChar_6(s int) *QChar {
	cthis := qtrt.Calloc(1, 256) // 2
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2ENS_16SpecialCharacterE", ffiqt.FFI_TYPE_VOID, cthis, &s)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:90
// index:7
// Public inline
// void QChar(char16_t)
func NewQChar_7(ch int16) *QChar {
	cthis := qtrt.Calloc(1, 256) // 2
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2EDs", ffiqt.FFI_TYPE_VOID, cthis, &ch)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:98
// index:8
// Public inline
// void QChar(char)
func NewQChar_8(c byte) *QChar {
	cthis := qtrt.Calloc(1, 256) // 2
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Ec", ffiqt.FFI_TYPE_VOID, cthis, &c)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:100
// index:9
// Public inline
// void QChar(uchar)
func NewQChar_9(c byte) *QChar {
	cthis := qtrt.Calloc(1, 256) // 2
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Eh", ffiqt.FFI_TYPE_VOID, cthis, &c)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:394
// index:0
// Public inline
// QChar::Category category()
func (this *QChar) Category() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar8categoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:504
// index:1
// Public static
// QChar::Category category(uint)
func (this *QChar) Category_1(ucs4 uint) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar8categoryEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QChar_Category_1(ucs4 uint) int {
	var nilthis *QChar
	rv := nilthis.Category_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:395
// index:0
// Public inline
// QChar::Direction direction()
func (this *QChar) Direction() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar9directionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:505
// index:1
// Public static
// QChar::Direction direction(uint)
func (this *QChar) Direction_1(ucs4 uint) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar9directionEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QChar_Direction_1(ucs4 uint) int {
	var nilthis *QChar
	rv := nilthis.Direction_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:396
// index:0
// Public inline
// QChar::JoiningType joiningType()
func (this *QChar) JoiningType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar11joiningTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:506
// index:1
// Public static
// QChar::JoiningType joiningType(uint)
func (this *QChar) JoiningType_1(ucs4 uint) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar11joiningTypeEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QChar_JoiningType_1(ucs4 uint) int {
	var nilthis *QChar
	rv := nilthis.JoiningType_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:398
// index:0
// Public inline
// QChar::Joining joining()
func (this *QChar) Joining() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7joiningEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:508
// index:1
// Public static
// QChar::Joining joining(uint)
func (this *QChar) Joining_1(ucs4 uint) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7joiningEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QChar_Joining_1(ucs4 uint) int {
	var nilthis *QChar
	rv := nilthis.Joining_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:411
// index:0
// Public inline
// unsigned char combiningClass()
func (this *QChar) CombiningClass() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar14combiningClassEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 111
}

// /usr/include/qt/QtCore/qchar.h:510
// index:1
// Public static
// unsigned char combiningClass(uint)
func (this *QChar) CombiningClass_1(ucs4 uint) byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar14combiningClassEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return byte(rv) // 111
}
func QChar_CombiningClass_1(ucs4 uint) byte {
	var nilthis *QChar
	rv := nilthis.CombiningClass_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:413
// index:0
// Public inline
// QChar mirroredChar()
func (this *QChar) MirroredChar() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar12mirroredCharEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:512
// index:1
// Public static
// uint mirroredChar(uint)
func (this *QChar) MirroredChar_1(ucs4 uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar12mirroredCharEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}
func QChar_MirroredChar_1(ucs4 uint) uint {
	var nilthis *QChar
	rv := nilthis.MirroredChar_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:414
// index:0
// Public inline
// bool hasMirrored()
func (this *QChar) HasMirrored() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar11hasMirroredEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:513
// index:1
// Public static
// bool hasMirrored(uint)
func (this *QChar) HasMirrored_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar11hasMirroredEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_HasMirrored_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.HasMirrored_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:416
// index:0
// Public
// QString decomposition()
func (this *QChar) Decomposition() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar13decompositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:515
// index:1
// Public static
// QString decomposition(uint)
func (this *QChar) Decomposition_1(ucs4 uint) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar13decompositionEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QChar_Decomposition_1(ucs4 uint) *QString /*123*/ {
	var nilthis *QChar
	rv := nilthis.Decomposition_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:417
// index:0
// Public inline
// QChar::Decomposition decompositionTag()
func (this *QChar) DecompositionTag() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar16decompositionTagEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:516
// index:1
// Public static
// QChar::Decomposition decompositionTag(uint)
func (this *QChar) DecompositionTag_1(ucs4 uint) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar16decompositionTagEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QChar_DecompositionTag_1(ucs4 uint) int {
	var nilthis *QChar
	rv := nilthis.DecompositionTag_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:419
// index:0
// Public inline
// int digitValue()
func (this *QChar) DigitValue() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar10digitValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qchar.h:518
// index:1
// Public static
// int digitValue(uint)
func (this *QChar) DigitValue_1(ucs4 uint) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar10digitValueEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QChar_DigitValue_1(ucs4 uint) int {
	var nilthis *QChar
	rv := nilthis.DigitValue_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:420
// index:0
// Public inline
// QChar toLower()
func (this *QChar) ToLower() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7toLowerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:519
// index:1
// Public static
// uint toLower(uint)
func (this *QChar) ToLower_1(ucs4 uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7toLowerEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}
func QChar_ToLower_1(ucs4 uint) uint {
	var nilthis *QChar
	rv := nilthis.ToLower_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:421
// index:0
// Public inline
// QChar toUpper()
func (this *QChar) ToUpper() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7toUpperEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:520
// index:1
// Public static
// uint toUpper(uint)
func (this *QChar) ToUpper_1(ucs4 uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7toUpperEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}
func QChar_ToUpper_1(ucs4 uint) uint {
	var nilthis *QChar
	rv := nilthis.ToUpper_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:422
// index:0
// Public inline
// QChar toTitleCase()
func (this *QChar) ToTitleCase() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar11toTitleCaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:521
// index:1
// Public static
// uint toTitleCase(uint)
func (this *QChar) ToTitleCase_1(ucs4 uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar11toTitleCaseEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}
func QChar_ToTitleCase_1(ucs4 uint) uint {
	var nilthis *QChar
	rv := nilthis.ToTitleCase_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:423
// index:0
// Public inline
// QChar toCaseFolded()
func (this *QChar) ToCaseFolded() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar12toCaseFoldedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:522
// index:1
// Public static
// uint toCaseFolded(uint)
func (this *QChar) ToCaseFolded_1(ucs4 uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar12toCaseFoldedEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}
func QChar_ToCaseFolded_1(ucs4 uint) uint {
	var nilthis *QChar
	rv := nilthis.ToCaseFolded_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:425
// index:0
// Public inline
// QChar::Script script()
func (this *QChar) Script() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar6scriptEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:524
// index:1
// Public static
// QChar::Script script(uint)
func (this *QChar) Script_1(ucs4 uint) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar6scriptEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QChar_Script_1(ucs4 uint) int {
	var nilthis *QChar
	rv := nilthis.Script_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:427
// index:0
// Public inline
// QChar::UnicodeVersion unicodeVersion()
func (this *QChar) UnicodeVersion() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar14unicodeVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:526
// index:1
// Public static
// QChar::UnicodeVersion unicodeVersion(uint)
func (this *QChar) UnicodeVersion_1(ucs4 uint) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar14unicodeVersionEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QChar_UnicodeVersion_1(ucs4 uint) int {
	var nilthis *QChar
	rv := nilthis.UnicodeVersion_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:432
// index:0
// Public inline
// char toLatin1()
func (this *QChar) ToLatin1() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar8toLatin1Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 111
}

// /usr/include/qt/QtCore/qchar.h:433
// index:0
// Public inline
// ushort unicode()
func (this *QChar) Unicode() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7unicodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qchar.h:434
// index:1
// Public inline
// ushort & unicode()
func (this *QChar) Unicode_1() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7unicodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv)
}

// /usr/include/qt/QtCore/qchar.h:440
// index:0
// Public static inline
// QChar fromLatin1(char)
func (this *QChar) FromLatin1(c byte) *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar10fromLatin1Ec", ffiqt.FFI_TYPE_POINTER, c)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QChar_FromLatin1(c byte) *QChar /*123*/ {
	var nilthis *QChar
	rv := nilthis.FromLatin1(c)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:442
// index:0
// Public inline
// bool isNull()
func (this *QChar) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:444
// index:0
// Public inline
// bool isPrint()
func (this *QChar) IsPrint() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isPrintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:530
// index:1
// Public static
// bool isPrint(uint)
func (this *QChar) IsPrint_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7isPrintEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsPrint_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsPrint_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:445
// index:0
// Public inline
// bool isSpace()
func (this *QChar) IsSpace() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isSpaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:531
// index:1
// Public static inline
// bool isSpace(uint)
func (this *QChar) IsSpace_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7isSpaceEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsSpace_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsSpace_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:446
// index:0
// Public inline
// bool isMark()
func (this *QChar) IsMark() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar6isMarkEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:537
// index:1
// Public static
// bool isMark(uint)
func (this *QChar) IsMark_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar6isMarkEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsMark_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsMark_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:447
// index:0
// Public inline
// bool isPunct()
func (this *QChar) IsPunct() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isPunctEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:538
// index:1
// Public static
// bool isPunct(uint)
func (this *QChar) IsPunct_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7isPunctEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsPunct_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsPunct_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:448
// index:0
// Public inline
// bool isSymbol()
func (this *QChar) IsSymbol() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar8isSymbolEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:539
// index:1
// Public static
// bool isSymbol(uint)
func (this *QChar) IsSymbol_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar8isSymbolEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsSymbol_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsSymbol_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:449
// index:0
// Public inline
// bool isLetter()
func (this *QChar) IsLetter() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar8isLetterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:540
// index:1
// Public static inline
// bool isLetter(uint)
func (this *QChar) IsLetter_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar8isLetterEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsLetter_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsLetter_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:450
// index:0
// Public inline
// bool isNumber()
func (this *QChar) IsNumber() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar8isNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:545
// index:1
// Public static inline
// bool isNumber(uint)
func (this *QChar) IsNumber_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar8isNumberEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsNumber_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsNumber_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:451
// index:0
// Public inline
// bool isLetterOrNumber()
func (this *QChar) IsLetterOrNumber() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar16isLetterOrNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:547
// index:1
// Public static inline
// bool isLetterOrNumber(uint)
func (this *QChar) IsLetterOrNumber_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar16isLetterOrNumberEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsLetterOrNumber_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsLetterOrNumber_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:452
// index:0
// Public inline
// bool isDigit()
func (this *QChar) IsDigit() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isDigitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:553
// index:1
// Public static inline
// bool isDigit(uint)
func (this *QChar) IsDigit_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7isDigitEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsDigit_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsDigit_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:453
// index:0
// Public inline
// bool isLower()
func (this *QChar) IsLower() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isLowerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:555
// index:1
// Public static inline
// bool isLower(uint)
func (this *QChar) IsLower_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7isLowerEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsLower_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsLower_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:454
// index:0
// Public inline
// bool isUpper()
func (this *QChar) IsUpper() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isUpperEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:557
// index:1
// Public static inline
// bool isUpper(uint)
func (this *QChar) IsUpper_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7isUpperEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsUpper_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsUpper_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:455
// index:0
// Public inline
// bool isTitleCase()
func (this *QChar) IsTitleCase() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar11isTitleCaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:559
// index:1
// Public static inline
// bool isTitleCase(uint)
func (this *QChar) IsTitleCase_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar11isTitleCaseEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsTitleCase_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsTitleCase_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:457
// index:0
// Public inline
// bool isNonCharacter()
func (this *QChar) IsNonCharacter() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar14isNonCharacterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:467
// index:1
// Public static inline
// bool isNonCharacter(uint)
func (this *QChar) IsNonCharacter_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar14isNonCharacterEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsNonCharacter_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsNonCharacter_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:458
// index:0
// Public inline
// bool isHighSurrogate()
func (this *QChar) IsHighSurrogate() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar15isHighSurrogateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:471
// index:1
// Public static inline
// bool isHighSurrogate(uint)
func (this *QChar) IsHighSurrogate_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar15isHighSurrogateEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsHighSurrogate_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsHighSurrogate_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:459
// index:0
// Public inline
// bool isLowSurrogate()
func (this *QChar) IsLowSurrogate() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar14isLowSurrogateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:475
// index:1
// Public static inline
// bool isLowSurrogate(uint)
func (this *QChar) IsLowSurrogate_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar14isLowSurrogateEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsLowSurrogate_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsLowSurrogate_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:460
// index:0
// Public inline
// bool isSurrogate()
func (this *QChar) IsSurrogate() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar11isSurrogateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:479
// index:1
// Public static inline
// bool isSurrogate(uint)
func (this *QChar) IsSurrogate_1(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar11isSurrogateEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_IsSurrogate_1(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.IsSurrogate_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:462
// index:0
// Public inline
// uchar cell()
func (this *QChar) Cell() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar4cellEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtCore/qchar.h:463
// index:0
// Public inline
// uchar row()
func (this *QChar) Row() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar3rowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtCore/qchar.h:464
// index:0
// Public inline
// void setCell(uchar)
func (this *QChar) SetCell(acell byte) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7setCellEh", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &acell)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:465
// index:0
// Public inline
// void setRow(uchar)
func (this *QChar) SetRow(arow byte) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar6setRowEh", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:483
// index:0
// Public static inline
// bool requiresSurrogates(uint)
func (this *QChar) RequiresSurrogates(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar18requiresSurrogatesEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QChar_RequiresSurrogates(ucs4 uint) bool {
	var nilthis *QChar
	rv := nilthis.RequiresSurrogates(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:487
// index:0
// Public static inline
// uint surrogateToUcs4(ushort, ushort)
func (this *QChar) SurrogateToUcs4(high uint16, low uint16) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar15surrogateToUcs4Ett", ffiqt.FFI_TYPE_POINTER, high, low)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}
func QChar_SurrogateToUcs4(high uint16, low uint16) uint {
	var nilthis *QChar
	rv := nilthis.SurrogateToUcs4(high, low)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:491
// index:1
// Public static inline
// uint surrogateToUcs4(class QChar, class QChar)
func (this *QChar) SurrogateToUcs4_1(high *QChar /*123*/, low *QChar /*123*/) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar15surrogateToUcs4ES_S_", ffiqt.FFI_TYPE_POINTER, high, low)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}
func QChar_SurrogateToUcs4_1(high *QChar /*123*/, low *QChar /*123*/) uint {
	var nilthis *QChar
	rv := nilthis.SurrogateToUcs4_1(high, low)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:495
// index:0
// Public static inline
// ushort highSurrogate(uint)
func (this *QChar) HighSurrogate(ucs4 uint) uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar13highSurrogateEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint16(rv) // 222
}
func QChar_HighSurrogate(ucs4 uint) uint16 {
	var nilthis *QChar
	rv := nilthis.HighSurrogate(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:499
// index:0
// Public static inline
// ushort lowSurrogate(uint)
func (this *QChar) LowSurrogate(ucs4 uint) uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar12lowSurrogateEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint16(rv) // 222
}
func QChar_LowSurrogate(ucs4 uint) uint16 {
	var nilthis *QChar
	rv := nilthis.LowSurrogate(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:528
// index:0
// Public static
// QChar::UnicodeVersion currentUnicodeVersion()
func (this *QChar) CurrentUnicodeVersion() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar21currentUnicodeVersionEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QChar_CurrentUnicodeVersion() int {
	var nilthis *QChar
	rv := nilthis.CurrentUnicodeVersion()
	return rv
}

//  body block end
