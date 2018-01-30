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
func (this *QChar) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQCharFromPointer(cthis unsafe.Pointer) *QChar {
	return &QChar{&qtrt.CObject{cthis}}
}
func (*QChar) NewFromPointer(cthis unsafe.Pointer) *QChar {
	return NewQCharFromPointer(cthis)
}

// /usr/include/qt/QtCore/qchar.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QChar()
func NewQChar() *QChar {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:82
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(ushort)
func NewQChar_1(rc uint16) *QChar {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Et", ffiqt.FFI_TYPE_POINTER, rc)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:83
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(uchar, uchar)
func NewQChar_2(c byte, r byte) *QChar {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Ehh", ffiqt.FFI_TYPE_POINTER, c, r)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:84
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(short)
func NewQChar_3(rc int16) *QChar {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Es", ffiqt.FFI_TYPE_POINTER, rc)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:85
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(uint)
func NewQChar_4(rc uint) *QChar {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Ej", ffiqt.FFI_TYPE_POINTER, rc)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:86
// index:5
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(int)
func NewQChar_5(rc int) *QChar {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Ei", ffiqt.FFI_TYPE_POINTER, rc)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:87
// index:6
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(enum QChar::SpecialCharacter)
func NewQChar_6(s int) *QChar {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2ENS_16SpecialCharacterE", ffiqt.FFI_TYPE_POINTER, s)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:90
// index:7
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(char16_t)
func NewQChar_7(ch int16) *QChar {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2EDs", ffiqt.FFI_TYPE_POINTER, ch)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:98
// index:8
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(char)
func NewQChar_8(c byte) *QChar {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Ec", ffiqt.FFI_TYPE_POINTER, c)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:100
// index:9
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(uchar)
func NewQChar_9(c byte) *QChar {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QCharC2Eh", ffiqt.FFI_TYPE_POINTER, c)
	gopp.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:394
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Category category()
func (this *QChar) Category() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar8categoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:504
// index:1
// Public static Visibility=Default Availability=Available
// [4] QChar::Category category(uint)
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
// Public inline Visibility=Default Availability=Available
// [4] QChar::Direction direction()
func (this *QChar) Direction() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar9directionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:505
// index:1
// Public static Visibility=Default Availability=Available
// [4] QChar::Direction direction(uint)
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
// Public inline Visibility=Default Availability=Available
// [4] QChar::JoiningType joiningType()
func (this *QChar) JoiningType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar11joiningTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:506
// index:1
// Public static Visibility=Default Availability=Available
// [4] QChar::JoiningType joiningType(uint)
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
// Public inline Visibility=Default Availability=Available
// [4] QChar::Joining joining()
func (this *QChar) Joining() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7joiningEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:508
// index:1
// Public static Visibility=Default Availability=Available
// [4] QChar::Joining joining(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] unsigned char combiningClass()
func (this *QChar) CombiningClass() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar14combiningClassEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("byte", rv).(byte) // 1111
}

// /usr/include/qt/QtCore/qchar.h:510
// index:1
// Public static Visibility=Default Availability=Available
// [1] unsigned char combiningClass(uint)
func (this *QChar) CombiningClass_1(ucs4 uint) byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar14combiningClassEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("byte", rv).(byte) // 1111
}
func QChar_CombiningClass_1(ucs4 uint) byte {
	var nilthis *QChar
	rv := nilthis.CombiningClass_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:413
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar mirroredChar()
func (this *QChar) MirroredChar() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar12mirroredCharEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:512
// index:1
// Public static Visibility=Default Availability=Available
// [4] uint mirroredChar(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool hasMirrored()
func (this *QChar) HasMirrored() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar11hasMirroredEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:513
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool hasMirrored(uint)
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
// Public Visibility=Default Availability=Available
// [8] QString decomposition()
func (this *QChar) Decomposition() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar13decompositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:515
// index:1
// Public static Visibility=Default Availability=Available
// [8] QString decomposition(uint)
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
// Public inline Visibility=Default Availability=Available
// [4] QChar::Decomposition decompositionTag()
func (this *QChar) DecompositionTag() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar16decompositionTagEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:516
// index:1
// Public static Visibility=Default Availability=Available
// [4] QChar::Decomposition decompositionTag(uint)
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
// Public inline Visibility=Default Availability=Available
// [4] int digitValue()
func (this *QChar) DigitValue() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar10digitValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qchar.h:518
// index:1
// Public static Visibility=Default Availability=Available
// [4] int digitValue(uint)
func (this *QChar) DigitValue_1(ucs4 uint) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar10digitValueEj", ffiqt.FFI_TYPE_POINTER, ucs4)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QChar_DigitValue_1(ucs4 uint) int {
	var nilthis *QChar
	rv := nilthis.DigitValue_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:420
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar toLower()
func (this *QChar) ToLower() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7toLowerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:519
// index:1
// Public static Visibility=Default Availability=Available
// [4] uint toLower(uint)
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
// Public inline Visibility=Default Availability=Available
// [2] QChar toUpper()
func (this *QChar) ToUpper() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7toUpperEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:520
// index:1
// Public static Visibility=Default Availability=Available
// [4] uint toUpper(uint)
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
// Public inline Visibility=Default Availability=Available
// [2] QChar toTitleCase()
func (this *QChar) ToTitleCase() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar11toTitleCaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:521
// index:1
// Public static Visibility=Default Availability=Available
// [4] uint toTitleCase(uint)
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
// Public inline Visibility=Default Availability=Available
// [2] QChar toCaseFolded()
func (this *QChar) ToCaseFolded() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar12toCaseFoldedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:522
// index:1
// Public static Visibility=Default Availability=Available
// [4] uint toCaseFolded(uint)
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
// Public inline Visibility=Default Availability=Available
// [4] QChar::Script script()
func (this *QChar) Script() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar6scriptEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:524
// index:1
// Public static Visibility=Default Availability=Available
// [4] QChar::Script script(uint)
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
// Public inline Visibility=Default Availability=Available
// [4] QChar::UnicodeVersion unicodeVersion()
func (this *QChar) UnicodeVersion() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar14unicodeVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:526
// index:1
// Public static Visibility=Default Availability=Available
// [4] QChar::UnicodeVersion unicodeVersion(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] char toLatin1()
func (this *QChar) ToLatin1() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar8toLatin1Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("byte", rv).(byte) // 1111
}

// /usr/include/qt/QtCore/qchar.h:433
// index:0
// Public inline Visibility=Default Availability=Available
// [2] ushort unicode()
func (this *QChar) Unicode() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7unicodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qchar.h:434
// index:1
// Public inline Visibility=Default Availability=Available
// [2] ushort & unicode()
func (this *QChar) Unicode_1() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7unicodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv)
}

// /usr/include/qt/QtCore/qchar.h:440
// index:0
// Public static inline Visibility=Default Availability=Available
// [2] QChar fromLatin1(char)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QChar) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:444
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isPrint()
func (this *QChar) IsPrint() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isPrintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:530
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool isPrint(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isSpace()
func (this *QChar) IsSpace() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isSpaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:531
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isSpace(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isMark()
func (this *QChar) IsMark() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar6isMarkEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:537
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool isMark(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isPunct()
func (this *QChar) IsPunct() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isPunctEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:538
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool isPunct(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isSymbol()
func (this *QChar) IsSymbol() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar8isSymbolEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:539
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool isSymbol(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isLetter()
func (this *QChar) IsLetter() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar8isLetterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:540
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isLetter(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isNumber()
func (this *QChar) IsNumber() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar8isNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:545
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isNumber(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isLetterOrNumber()
func (this *QChar) IsLetterOrNumber() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar16isLetterOrNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:547
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isLetterOrNumber(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isDigit()
func (this *QChar) IsDigit() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isDigitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:553
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isDigit(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isLower()
func (this *QChar) IsLower() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isLowerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:555
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isLower(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isUpper()
func (this *QChar) IsUpper() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar7isUpperEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:557
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isUpper(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isTitleCase()
func (this *QChar) IsTitleCase() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar11isTitleCaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:559
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isTitleCase(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isNonCharacter()
func (this *QChar) IsNonCharacter() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar14isNonCharacterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:467
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isNonCharacter(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isHighSurrogate()
func (this *QChar) IsHighSurrogate() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar15isHighSurrogateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:471
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isHighSurrogate(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isLowSurrogate()
func (this *QChar) IsLowSurrogate() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar14isLowSurrogateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:475
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isLowSurrogate(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] bool isSurrogate()
func (this *QChar) IsSurrogate() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar11isSurrogateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:479
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isSurrogate(uint)
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
// Public inline Visibility=Default Availability=Available
// [1] uchar cell()
func (this *QChar) Cell() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar4cellEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtCore/qchar.h:463
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar row()
func (this *QChar) Row() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QChar3rowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtCore/qchar.h:464
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCell(uchar)
func (this *QChar) SetCell(acell byte) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar7setCellEh", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), acell)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:465
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRow(uchar)
func (this *QChar) SetRow(arow byte) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar6setRowEh", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:483
// index:0
// Public static inline Visibility=Default Availability=Available
// [1] bool requiresSurrogates(uint)
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
// Public static inline Visibility=Default Availability=Available
// [4] uint surrogateToUcs4(ushort, ushort)
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
// Public static inline Visibility=Default Availability=Available
// [4] uint surrogateToUcs4(QChar, QChar)
func (this *QChar) SurrogateToUcs4_1(high *QChar /*123*/, low *QChar /*123*/) uint {
	var convArg0 = high.GetCthis()
	var convArg1 = low.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QChar15surrogateToUcs4ES_S_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
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
// Public static inline Visibility=Default Availability=Available
// [2] ushort highSurrogate(uint)
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
// Public static inline Visibility=Default Availability=Available
// [2] ushort lowSurrogate(uint)
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
// Public static Visibility=Default Availability=Available
// [4] QChar::UnicodeVersion currentUnicodeVersion()
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

type QChar__SpecialCharacter = int

const QChar__Null QChar__SpecialCharacter = 0
const QChar__Tabulation QChar__SpecialCharacter = 9
const QChar__LineFeed QChar__SpecialCharacter = 10
const QChar__CarriageReturn QChar__SpecialCharacter = 13
const QChar__Space QChar__SpecialCharacter = 32
const QChar__Nbsp QChar__SpecialCharacter = 160
const QChar__SoftHyphen QChar__SpecialCharacter = 173
const QChar__ReplacementCharacter QChar__SpecialCharacter = 65533
const QChar__ObjectReplacementCharacter QChar__SpecialCharacter = 65532
const QChar__ByteOrderMark QChar__SpecialCharacter = 65279
const QChar__ByteOrderSwapped QChar__SpecialCharacter = 65534
const QChar__ParagraphSeparator QChar__SpecialCharacter = 8233
const QChar__LineSeparator QChar__SpecialCharacter = 8232
const QChar__LastValidCodePoint QChar__SpecialCharacter = 1114111

type QChar__Category = int

const QChar__Mark_NonSpacing QChar__Category = 0
const QChar__Mark_SpacingCombining QChar__Category = 1
const QChar__Mark_Enclosing QChar__Category = 2
const QChar__Number_DecimalDigit QChar__Category = 3
const QChar__Number_Letter QChar__Category = 4
const QChar__Number_Other QChar__Category = 5
const QChar__Separator_Space QChar__Category = 6
const QChar__Separator_Line QChar__Category = 7
const QChar__Separator_Paragraph QChar__Category = 8
const QChar__Other_Control QChar__Category = 9
const QChar__Other_Format QChar__Category = 10
const QChar__Other_Surrogate QChar__Category = 11
const QChar__Other_PrivateUse QChar__Category = 12
const QChar__Other_NotAssigned QChar__Category = 13
const QChar__Letter_Uppercase QChar__Category = 14
const QChar__Letter_Lowercase QChar__Category = 15
const QChar__Letter_Titlecase QChar__Category = 16
const QChar__Letter_Modifier QChar__Category = 17
const QChar__Letter_Other QChar__Category = 18
const QChar__Punctuation_Connector QChar__Category = 19
const QChar__Punctuation_Dash QChar__Category = 20
const QChar__Punctuation_Open QChar__Category = 21
const QChar__Punctuation_Close QChar__Category = 22
const QChar__Punctuation_InitialQuote QChar__Category = 23
const QChar__Punctuation_FinalQuote QChar__Category = 24
const QChar__Punctuation_Other QChar__Category = 25
const QChar__Symbol_Math QChar__Category = 26
const QChar__Symbol_Currency QChar__Category = 27
const QChar__Symbol_Modifier QChar__Category = 28
const QChar__Symbol_Other QChar__Category = 29

type QChar__Script = int

const QChar__Script_Unknown QChar__Script = 0
const QChar__Script_Inherited QChar__Script = 1
const QChar__Script_Common QChar__Script = 2
const QChar__Script_Latin QChar__Script = 3
const QChar__Script_Greek QChar__Script = 4
const QChar__Script_Cyrillic QChar__Script = 5
const QChar__Script_Armenian QChar__Script = 6
const QChar__Script_Hebrew QChar__Script = 7
const QChar__Script_Arabic QChar__Script = 8
const QChar__Script_Syriac QChar__Script = 9
const QChar__Script_Thaana QChar__Script = 10
const QChar__Script_Devanagari QChar__Script = 11
const QChar__Script_Bengali QChar__Script = 12
const QChar__Script_Gurmukhi QChar__Script = 13
const QChar__Script_Gujarati QChar__Script = 14
const QChar__Script_Oriya QChar__Script = 15
const QChar__Script_Tamil QChar__Script = 16
const QChar__Script_Telugu QChar__Script = 17
const QChar__Script_Kannada QChar__Script = 18
const QChar__Script_Malayalam QChar__Script = 19
const QChar__Script_Sinhala QChar__Script = 20
const QChar__Script_Thai QChar__Script = 21
const QChar__Script_Lao QChar__Script = 22
const QChar__Script_Tibetan QChar__Script = 23
const QChar__Script_Myanmar QChar__Script = 24
const QChar__Script_Georgian QChar__Script = 25
const QChar__Script_Hangul QChar__Script = 26
const QChar__Script_Ethiopic QChar__Script = 27
const QChar__Script_Cherokee QChar__Script = 28
const QChar__Script_CanadianAboriginal QChar__Script = 29
const QChar__Script_Ogham QChar__Script = 30
const QChar__Script_Runic QChar__Script = 31
const QChar__Script_Khmer QChar__Script = 32
const QChar__Script_Mongolian QChar__Script = 33
const QChar__Script_Hiragana QChar__Script = 34
const QChar__Script_Katakana QChar__Script = 35
const QChar__Script_Bopomofo QChar__Script = 36
const QChar__Script_Han QChar__Script = 37
const QChar__Script_Yi QChar__Script = 38
const QChar__Script_OldItalic QChar__Script = 39
const QChar__Script_Gothic QChar__Script = 40
const QChar__Script_Deseret QChar__Script = 41
const QChar__Script_Tagalog QChar__Script = 42
const QChar__Script_Hanunoo QChar__Script = 43
const QChar__Script_Buhid QChar__Script = 44
const QChar__Script_Tagbanwa QChar__Script = 45
const QChar__Script_Coptic QChar__Script = 46
const QChar__Script_Limbu QChar__Script = 47
const QChar__Script_TaiLe QChar__Script = 48
const QChar__Script_LinearB QChar__Script = 49
const QChar__Script_Ugaritic QChar__Script = 50
const QChar__Script_Shavian QChar__Script = 51
const QChar__Script_Osmanya QChar__Script = 52
const QChar__Script_Cypriot QChar__Script = 53
const QChar__Script_Braille QChar__Script = 54
const QChar__Script_Buginese QChar__Script = 55
const QChar__Script_NewTaiLue QChar__Script = 56
const QChar__Script_Glagolitic QChar__Script = 57
const QChar__Script_Tifinagh QChar__Script = 58
const QChar__Script_SylotiNagri QChar__Script = 59
const QChar__Script_OldPersian QChar__Script = 60
const QChar__Script_Kharoshthi QChar__Script = 61
const QChar__Script_Balinese QChar__Script = 62
const QChar__Script_Cuneiform QChar__Script = 63
const QChar__Script_Phoenician QChar__Script = 64
const QChar__Script_PhagsPa QChar__Script = 65
const QChar__Script_Nko QChar__Script = 66
const QChar__Script_Sundanese QChar__Script = 67
const QChar__Script_Lepcha QChar__Script = 68
const QChar__Script_OlChiki QChar__Script = 69
const QChar__Script_Vai QChar__Script = 70
const QChar__Script_Saurashtra QChar__Script = 71
const QChar__Script_KayahLi QChar__Script = 72
const QChar__Script_Rejang QChar__Script = 73
const QChar__Script_Lycian QChar__Script = 74
const QChar__Script_Carian QChar__Script = 75
const QChar__Script_Lydian QChar__Script = 76
const QChar__Script_Cham QChar__Script = 77
const QChar__Script_TaiTham QChar__Script = 78
const QChar__Script_TaiViet QChar__Script = 79
const QChar__Script_Avestan QChar__Script = 80
const QChar__Script_EgyptianHieroglyphs QChar__Script = 81
const QChar__Script_Samaritan QChar__Script = 82
const QChar__Script_Lisu QChar__Script = 83
const QChar__Script_Bamum QChar__Script = 84
const QChar__Script_Javanese QChar__Script = 85
const QChar__Script_MeeteiMayek QChar__Script = 86
const QChar__Script_ImperialAramaic QChar__Script = 87
const QChar__Script_OldSouthArabian QChar__Script = 88
const QChar__Script_InscriptionalParthian QChar__Script = 89
const QChar__Script_InscriptionalPahlavi QChar__Script = 90
const QChar__Script_OldTurkic QChar__Script = 91
const QChar__Script_Kaithi QChar__Script = 92
const QChar__Script_Batak QChar__Script = 93
const QChar__Script_Brahmi QChar__Script = 94
const QChar__Script_Mandaic QChar__Script = 95
const QChar__Script_Chakma QChar__Script = 96
const QChar__Script_MeroiticCursive QChar__Script = 97
const QChar__Script_MeroiticHieroglyphs QChar__Script = 98
const QChar__Script_Miao QChar__Script = 99
const QChar__Script_Sharada QChar__Script = 100
const QChar__Script_SoraSompeng QChar__Script = 101
const QChar__Script_Takri QChar__Script = 102
const QChar__Script_CaucasianAlbanian QChar__Script = 103
const QChar__Script_BassaVah QChar__Script = 104
const QChar__Script_Duployan QChar__Script = 105
const QChar__Script_Elbasan QChar__Script = 106
const QChar__Script_Grantha QChar__Script = 107
const QChar__Script_PahawhHmong QChar__Script = 108
const QChar__Script_Khojki QChar__Script = 109
const QChar__Script_LinearA QChar__Script = 110
const QChar__Script_Mahajani QChar__Script = 111
const QChar__Script_Manichaean QChar__Script = 112
const QChar__Script_MendeKikakui QChar__Script = 113
const QChar__Script_Modi QChar__Script = 114
const QChar__Script_Mro QChar__Script = 115
const QChar__Script_OldNorthArabian QChar__Script = 116
const QChar__Script_Nabataean QChar__Script = 117
const QChar__Script_Palmyrene QChar__Script = 118
const QChar__Script_PauCinHau QChar__Script = 119
const QChar__Script_OldPermic QChar__Script = 120
const QChar__Script_PsalterPahlavi QChar__Script = 121
const QChar__Script_Siddham QChar__Script = 122
const QChar__Script_Khudawadi QChar__Script = 123
const QChar__Script_Tirhuta QChar__Script = 124
const QChar__Script_WarangCiti QChar__Script = 125
const QChar__Script_Ahom QChar__Script = 126
const QChar__Script_AnatolianHieroglyphs QChar__Script = 127
const QChar__Script_Hatran QChar__Script = 128
const QChar__Script_Multani QChar__Script = 129
const QChar__Script_OldHungarian QChar__Script = 130
const QChar__Script_SignWriting QChar__Script = 131
const QChar__ScriptCount QChar__Script = 132

type QChar__Direction = int

const QChar__DirL QChar__Direction = 0
const QChar__DirR QChar__Direction = 1
const QChar__DirEN QChar__Direction = 2
const QChar__DirES QChar__Direction = 3
const QChar__DirET QChar__Direction = 4
const QChar__DirAN QChar__Direction = 5
const QChar__DirCS QChar__Direction = 6
const QChar__DirB QChar__Direction = 7
const QChar__DirS QChar__Direction = 8
const QChar__DirWS QChar__Direction = 9
const QChar__DirON QChar__Direction = 10
const QChar__DirLRE QChar__Direction = 11
const QChar__DirLRO QChar__Direction = 12
const QChar__DirAL QChar__Direction = 13
const QChar__DirRLE QChar__Direction = 14
const QChar__DirRLO QChar__Direction = 15
const QChar__DirPDF QChar__Direction = 16
const QChar__DirNSM QChar__Direction = 17
const QChar__DirBN QChar__Direction = 18
const QChar__DirLRI QChar__Direction = 19
const QChar__DirRLI QChar__Direction = 20
const QChar__DirFSI QChar__Direction = 21
const QChar__DirPDI QChar__Direction = 22

type QChar__Decomposition = int

const QChar__NoDecomposition QChar__Decomposition = 0
const QChar__Canonical QChar__Decomposition = 1
const QChar__Font QChar__Decomposition = 2
const QChar__NoBreak QChar__Decomposition = 3
const QChar__Initial QChar__Decomposition = 4
const QChar__Medial QChar__Decomposition = 5
const QChar__Final QChar__Decomposition = 6
const QChar__Isolated QChar__Decomposition = 7
const QChar__Circle QChar__Decomposition = 8
const QChar__Super QChar__Decomposition = 9
const QChar__Sub QChar__Decomposition = 10
const QChar__Vertical QChar__Decomposition = 11
const QChar__Wide QChar__Decomposition = 12
const QChar__Narrow QChar__Decomposition = 13
const QChar__Small QChar__Decomposition = 14
const QChar__Square QChar__Decomposition = 15
const QChar__Compat QChar__Decomposition = 16
const QChar__Fraction QChar__Decomposition = 17

type QChar__JoiningType = int

const QChar__Joining_None QChar__JoiningType = 0
const QChar__Joining_Causing QChar__JoiningType = 1
const QChar__Joining_Dual QChar__JoiningType = 2
const QChar__Joining_Right QChar__JoiningType = 3
const QChar__Joining_Left QChar__JoiningType = 4
const QChar__Joining_Transparent QChar__JoiningType = 5

type QChar__Joining = int

const QChar__OtherJoining QChar__Joining = 0
const QChar__Dual QChar__Joining = 1
const QChar__Right QChar__Joining = 2
const QChar__Center QChar__Joining = 3

type QChar__CombiningClass = int

const QChar__Combining_BelowLeftAttached QChar__CombiningClass = 200
const QChar__Combining_BelowAttached QChar__CombiningClass = 202
const QChar__Combining_BelowRightAttached QChar__CombiningClass = 204
const QChar__Combining_LeftAttached QChar__CombiningClass = 208
const QChar__Combining_RightAttached QChar__CombiningClass = 210
const QChar__Combining_AboveLeftAttached QChar__CombiningClass = 212
const QChar__Combining_AboveAttached QChar__CombiningClass = 214
const QChar__Combining_AboveRightAttached QChar__CombiningClass = 216
const QChar__Combining_BelowLeft QChar__CombiningClass = 218
const QChar__Combining_Below QChar__CombiningClass = 220
const QChar__Combining_BelowRight QChar__CombiningClass = 222
const QChar__Combining_Left QChar__CombiningClass = 224
const QChar__Combining_Right QChar__CombiningClass = 226
const QChar__Combining_AboveLeft QChar__CombiningClass = 228
const QChar__Combining_Above QChar__CombiningClass = 230
const QChar__Combining_AboveRight QChar__CombiningClass = 232
const QChar__Combining_DoubleBelow QChar__CombiningClass = 233
const QChar__Combining_DoubleAbove QChar__CombiningClass = 234
const QChar__Combining_IotaSubscript QChar__CombiningClass = 240

type QChar__UnicodeVersion = int

const QChar__Unicode_Unassigned QChar__UnicodeVersion = 0
const QChar__Unicode_1_1 QChar__UnicodeVersion = 1
const QChar__Unicode_2_0 QChar__UnicodeVersion = 2
const QChar__Unicode_2_1_2 QChar__UnicodeVersion = 3
const QChar__Unicode_3_0 QChar__UnicodeVersion = 4
const QChar__Unicode_3_1 QChar__UnicodeVersion = 5
const QChar__Unicode_3_2 QChar__UnicodeVersion = 6
const QChar__Unicode_4_0 QChar__UnicodeVersion = 7
const QChar__Unicode_4_1 QChar__UnicodeVersion = 8
const QChar__Unicode_5_0 QChar__UnicodeVersion = 9
const QChar__Unicode_5_1 QChar__UnicodeVersion = 10
const QChar__Unicode_5_2 QChar__UnicodeVersion = 11
const QChar__Unicode_6_0 QChar__UnicodeVersion = 12
const QChar__Unicode_6_1 QChar__UnicodeVersion = 13
const QChar__Unicode_6_2 QChar__UnicodeVersion = 14
const QChar__Unicode_6_3 QChar__UnicodeVersion = 15
const QChar__Unicode_7_0 QChar__UnicodeVersion = 16
const QChar__Unicode_8_0 QChar__UnicodeVersion = 17

//  body block end
