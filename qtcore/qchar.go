package qtcore

// /usr/include/qt/QtCore/qchar.h
// #include <qchar.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QChar struct {
	*qtrt.CObject
}
type QChar_ITF interface {
	QChar_PTR() *QChar
}

func (ptr *QChar) QChar_PTR() *QChar { return ptr }

func (this *QChar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QChar) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
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

/*
Constructs a null QChar ('\0').

See also isNull().
*/
func NewQChar() *QChar {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QCharC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQChar)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:82
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(ushort)

/*
Constructs a null QChar ('\0').

See also isNull().
*/
func NewQChar_1(rc uint16) *QChar {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QCharC2Et", qtrt.FFI_TYPE_POINTER, rc)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQChar)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:83
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(uchar, uchar)

/*
Constructs a null QChar ('\0').

See also isNull().
*/
func NewQChar_2(c byte, r byte) *QChar {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QCharC2Ehh", qtrt.FFI_TYPE_POINTER, c, r)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQChar)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:84
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(short)

/*
Constructs a null QChar ('\0').

See also isNull().
*/
func NewQChar_3(rc int16) *QChar {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QCharC2Es", qtrt.FFI_TYPE_POINTER, rc)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQChar)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:85
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(uint)

/*
Constructs a null QChar ('\0').

See also isNull().
*/
func NewQChar_4(rc uint) *QChar {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QCharC2Ej", qtrt.FFI_TYPE_POINTER, rc)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQChar)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:86
// index:5
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(int)

/*
Constructs a null QChar ('\0').

See also isNull().
*/
func NewQChar_5(rc int) *QChar {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QCharC2Ei", qtrt.FFI_TYPE_POINTER, rc)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQChar)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:87
// index:6
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(enum QChar::SpecialCharacter)

/*
Constructs a null QChar ('\0').

See also isNull().
*/
func NewQChar_6(s int) *QChar {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QCharC2ENS_16SpecialCharacterE", qtrt.FFI_TYPE_POINTER, s)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQChar)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:90
// index:7
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(char16_t)

/*
Constructs a null QChar ('\0').

See also isNull().
*/
func NewQChar_7(ch int16) *QChar {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QCharC2EDs", qtrt.FFI_TYPE_POINTER, ch)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQChar)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:98
// index:8
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(char)

/*
Constructs a null QChar ('\0').

See also isNull().
*/
func NewQChar_8(c byte) *QChar {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QCharC2Ec", qtrt.FFI_TYPE_POINTER, c)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQChar)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:100
// index:9
// Public inline Visibility=Default Availability=Available
// [-2] void QChar(uchar)

/*
Constructs a null QChar ('\0').

See also isNull().
*/
func NewQChar_9(c byte) *QChar {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QCharC2Eh", qtrt.FFI_TYPE_POINTER, c)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCharFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQChar)
	return gothis
}

// /usr/include/qt/QtCore/qchar.h:394
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Category category() const

/*
Returns the character's category.
*/
func (this *QChar) Category() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar8categoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:504
// index:1
// Public static Visibility=Default Availability=Available
// [4] QChar::Category category(uint)

/*
Returns the character's category.
*/
func (this *QChar) Category_1(ucs4 uint) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar8categoryEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [4] QChar::Direction direction() const

/*
Returns the character's direction.
*/
func (this *QChar) Direction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar9directionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:505
// index:1
// Public static Visibility=Default Availability=Available
// [4] QChar::Direction direction(uint)

/*
Returns the character's direction.
*/
func (this *QChar) Direction_1(ucs4 uint) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar9directionEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [4] QChar::JoiningType joiningType() const

/*
Returns information about the joining type attributes of the character (needed for certain languages such as Arabic or Syriac).

This function was introduced in  Qt 5.3.
*/
func (this *QChar) JoiningType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar11joiningTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:506
// index:1
// Public static Visibility=Default Availability=Available
// [4] QChar::JoiningType joiningType(uint)

/*
Returns information about the joining type attributes of the character (needed for certain languages such as Arabic or Syriac).

This function was introduced in  Qt 5.3.
*/
func (this *QChar) JoiningType_1(ucs4 uint) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar11joiningTypeEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [4] QChar::Joining joining() const

/*

 */
func (this *QChar) Joining() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar7joiningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:508
// index:1
// Public static Visibility=Default Availability=Available
// [4] QChar::Joining joining(uint)

/*

 */
func (this *QChar) Joining_1(ucs4 uint) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar7joiningEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] unsigned char combiningClass() const

/*
Returns the combining class for the character as defined in the Unicode standard. This is mainly useful as a positioning hint for marks attached to a base character.

The Qt text rendering engine uses this information to correctly position non-spacing marks around a base character.
*/
func (this *QChar) CombiningClass() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar14combiningClassEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("byte", rv).(byte) // 1111
}

// /usr/include/qt/QtCore/qchar.h:510
// index:1
// Public static Visibility=Default Availability=Available
// [1] unsigned char combiningClass(uint)

/*
Returns the combining class for the character as defined in the Unicode standard. This is mainly useful as a positioning hint for marks attached to a base character.

The Qt text rendering engine uses this information to correctly position non-spacing marks around a base character.
*/
func (this *QChar) CombiningClass_1(ucs4 uint) byte {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar14combiningClassEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [2] QChar mirroredChar() const

/*
Returns the mirrored character if this character is a mirrored character; otherwise returns the character itself.

See also hasMirrored().
*/
func (this *QChar) MirroredChar() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar12mirroredCharEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:512
// index:1
// Public static Visibility=Default Availability=Available
// [4] uint mirroredChar(uint)

/*
Returns the mirrored character if this character is a mirrored character; otherwise returns the character itself.

See also hasMirrored().
*/
func (this *QChar) MirroredChar_1(ucs4 uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar12mirroredCharEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
	// unsigned int // 222
}
func QChar_MirroredChar_1(ucs4 uint) uint {
	var nilthis *QChar
	rv := nilthis.MirroredChar_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:414
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool hasMirrored() const

/*
Returns true if the character should be reversed if the text direction is reversed; otherwise returns false.

A bit faster equivalent of (ch.mirroredChar() != ch).

See also mirroredChar().
*/
func (this *QChar) HasMirrored() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar11hasMirroredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:513
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool hasMirrored(uint)

/*
Returns true if the character should be reversed if the text direction is reversed; otherwise returns false.

A bit faster equivalent of (ch.mirroredChar() != ch).

See also mirroredChar().
*/
func (this *QChar) HasMirrored_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar11hasMirroredEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [8] QString decomposition() const

/*
Decomposes a character into it's constituent parts. Returns an empty string if no decomposition exists.
*/
func (this *QChar) Decomposition() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar13decompositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qchar.h:515
// index:1
// Public static Visibility=Default Availability=Available
// [8] QString decomposition(uint)

/*
Decomposes a character into it's constituent parts. Returns an empty string if no decomposition exists.
*/
func (this *QChar) Decomposition_1(ucs4 uint) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar13decompositionEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QChar_Decomposition_1(ucs4 uint) string {
	var nilthis *QChar
	rv := nilthis.Decomposition_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:417
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Decomposition decompositionTag() const

/*
Returns the tag defining the composition of the character. Returns QChar::NoDecomposition if no decomposition exists.
*/
func (this *QChar) DecompositionTag() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar16decompositionTagEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:516
// index:1
// Public static Visibility=Default Availability=Available
// [4] QChar::Decomposition decompositionTag(uint)

/*
Returns the tag defining the composition of the character. Returns QChar::NoDecomposition if no decomposition exists.
*/
func (this *QChar) DecompositionTag_1(ucs4 uint) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar16decompositionTagEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [4] int digitValue() const

/*
Returns the numeric value of the digit, or -1 if the character is not a digit.
*/
func (this *QChar) DigitValue() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar10digitValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qchar.h:518
// index:1
// Public static Visibility=Default Availability=Available
// [4] int digitValue(uint)

/*
Returns the numeric value of the digit, or -1 if the character is not a digit.
*/
func (this *QChar) DigitValue_1(ucs4 uint) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar10digitValueEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [2] QChar toLower() const

/*
Returns the lowercase equivalent if the character is uppercase or titlecase; otherwise returns the character itself.
*/
func (this *QChar) ToLower() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar7toLowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:519
// index:1
// Public static Visibility=Default Availability=Available
// [4] uint toLower(uint)

/*
Returns the lowercase equivalent if the character is uppercase or titlecase; otherwise returns the character itself.
*/
func (this *QChar) ToLower_1(ucs4 uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar7toLowerEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
	// unsigned int // 222
}
func QChar_ToLower_1(ucs4 uint) uint {
	var nilthis *QChar
	rv := nilthis.ToLower_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:421
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar toUpper() const

/*
Returns the uppercase equivalent if the character is lowercase or titlecase; otherwise returns the character itself.
*/
func (this *QChar) ToUpper() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar7toUpperEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:520
// index:1
// Public static Visibility=Default Availability=Available
// [4] uint toUpper(uint)

/*
Returns the uppercase equivalent if the character is lowercase or titlecase; otherwise returns the character itself.
*/
func (this *QChar) ToUpper_1(ucs4 uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar7toUpperEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
	// unsigned int // 222
}
func QChar_ToUpper_1(ucs4 uint) uint {
	var nilthis *QChar
	rv := nilthis.ToUpper_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:422
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar toTitleCase() const

/*
Returns the title case equivalent if the character is lowercase or uppercase; otherwise returns the character itself.
*/
func (this *QChar) ToTitleCase() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar11toTitleCaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:521
// index:1
// Public static Visibility=Default Availability=Available
// [4] uint toTitleCase(uint)

/*
Returns the title case equivalent if the character is lowercase or uppercase; otherwise returns the character itself.
*/
func (this *QChar) ToTitleCase_1(ucs4 uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar11toTitleCaseEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
	// unsigned int // 222
}
func QChar_ToTitleCase_1(ucs4 uint) uint {
	var nilthis *QChar
	rv := nilthis.ToTitleCase_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:423
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar toCaseFolded() const

/*
Returns the case folded equivalent of the character. For most Unicode characters this is the same as toLower().
*/
func (this *QChar) ToCaseFolded() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar12toCaseFoldedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qchar.h:522
// index:1
// Public static Visibility=Default Availability=Available
// [4] uint toCaseFolded(uint)

/*
Returns the case folded equivalent of the character. For most Unicode characters this is the same as toLower().
*/
func (this *QChar) ToCaseFolded_1(ucs4 uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar12toCaseFoldedEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
	// unsigned int // 222
}
func QChar_ToCaseFolded_1(ucs4 uint) uint {
	var nilthis *QChar
	rv := nilthis.ToCaseFolded_1(ucs4)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:425
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Script script() const

/*
Returns the Unicode script property value for this character.

This function was introduced in  Qt 5.1.
*/
func (this *QChar) Script() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar6scriptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:524
// index:1
// Public static Visibility=Default Availability=Available
// [4] QChar::Script script(uint)

/*
Returns the Unicode script property value for this character.

This function was introduced in  Qt 5.1.
*/
func (this *QChar) Script_1(ucs4 uint) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar6scriptEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [4] QChar::UnicodeVersion unicodeVersion() const

/*
Returns the Unicode version that introduced this character.
*/
func (this *QChar) UnicodeVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar14unicodeVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qchar.h:526
// index:1
// Public static Visibility=Default Availability=Available
// [4] QChar::UnicodeVersion unicodeVersion(uint)

/*
Returns the Unicode version that introduced this character.
*/
func (this *QChar) UnicodeVersion_1(ucs4 uint) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar14unicodeVersionEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] char toLatin1() const

/*
Returns the Latin-1 character equivalent to the QChar, or 0. This is mainly useful for non-internationalized software.

Note: It is not possible to distinguish a non-Latin-1 character from a Latin-1 0 (NUL) character. Prefer to use unicode(), which does not have this ambiguity.

See also unicode().
*/
func (this *QChar) ToLatin1() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar8toLatin1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("byte", rv).(byte) // 1111
}

// /usr/include/qt/QtCore/qchar.h:433
// index:0
// Public inline Visibility=Default Availability=Available
// [2] ushort unicode() const

/*
Returns the numeric Unicode value of the QChar.
*/
func (this *QChar) Unicode() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar7unicodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
	// unsigned short // 222
}

// /usr/include/qt/QtCore/qchar.h:434
// index:1
// Public inline Visibility=Default Availability=Available
// [2] ushort & unicode()

/*
Returns the numeric Unicode value of the QChar.
*/
func (this *QChar) Unicode_1() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar7unicodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv)
}

// /usr/include/qt/QtCore/qchar.h:440
// index:0
// Public static inline Visibility=Default Availability=Available
// [2] QChar fromLatin1(char)

/*
Converts the Latin-1 character c to its equivalent QChar. This is mainly useful for non-internationalized software.

An alternative is to use QLatin1Char.

See also toLatin1() and unicode().
*/
func (this *QChar) FromLatin1(c byte) *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar10fromLatin1Ec", qtrt.FFI_TYPE_POINTER, c)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
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
// [1] bool isNull() const

/*
Returns true if the character is the Unicode character 0x0000 ('\0'); otherwise returns false.
*/
func (this *QChar) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:444
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isPrint() const

/*
Returns true if the character is a printable character; otherwise returns false. This is any character not of category Other_*.

Note that this gives no indication of whether the character is available in a particular font.
*/
func (this *QChar) IsPrint() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar7isPrintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:530
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool isPrint(uint)

/*
Returns true if the character is a printable character; otherwise returns false. This is any character not of category Other_*.

Note that this gives no indication of whether the character is available in a particular font.
*/
func (this *QChar) IsPrint_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar7isPrintEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] bool isSpace() const

/*
Returns true if the character is a separator character (Separator_* categories or certain code points from Other_Control category); otherwise returns false.
*/
func (this *QChar) IsSpace() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar7isSpaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:531
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isSpace(uint)

/*
Returns true if the character is a separator character (Separator_* categories or certain code points from Other_Control category); otherwise returns false.
*/
func (this *QChar) IsSpace_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar7isSpaceEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] bool isMark() const

/*
Returns true if the character is a mark (Mark_* categories); otherwise returns false.

See QChar::Category for more information regarding marks.
*/
func (this *QChar) IsMark() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar6isMarkEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:537
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool isMark(uint)

/*
Returns true if the character is a mark (Mark_* categories); otherwise returns false.

See QChar::Category for more information regarding marks.
*/
func (this *QChar) IsMark_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar6isMarkEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] bool isPunct() const

/*
Returns true if the character is a punctuation mark (Punctuation_* categories); otherwise returns false.
*/
func (this *QChar) IsPunct() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar7isPunctEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:538
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool isPunct(uint)

/*
Returns true if the character is a punctuation mark (Punctuation_* categories); otherwise returns false.
*/
func (this *QChar) IsPunct_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar7isPunctEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] bool isSymbol() const

/*
Returns true if the character is a symbol (Symbol_* categories); otherwise returns false.
*/
func (this *QChar) IsSymbol() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar8isSymbolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:539
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool isSymbol(uint)

/*
Returns true if the character is a symbol (Symbol_* categories); otherwise returns false.
*/
func (this *QChar) IsSymbol_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar8isSymbolEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] bool isLetter() const

/*
Returns true if the character is a letter (Letter_* categories); otherwise returns false.
*/
func (this *QChar) IsLetter() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar8isLetterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:540
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isLetter(uint)

/*
Returns true if the character is a letter (Letter_* categories); otherwise returns false.
*/
func (this *QChar) IsLetter_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar8isLetterEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] bool isNumber() const

/*
Returns true if the character is a number (Number_* categories, not just 0-9); otherwise returns false.

See also isDigit().
*/
func (this *QChar) IsNumber() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar8isNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:545
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isNumber(uint)

/*
Returns true if the character is a number (Number_* categories, not just 0-9); otherwise returns false.

See also isDigit().
*/
func (this *QChar) IsNumber_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar8isNumberEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] bool isLetterOrNumber() const

/*
Returns true if the character is a letter or number (Letter_* or Number_* categories); otherwise returns false.
*/
func (this *QChar) IsLetterOrNumber() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar16isLetterOrNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:547
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isLetterOrNumber(uint)

/*
Returns true if the character is a letter or number (Letter_* or Number_* categories); otherwise returns false.
*/
func (this *QChar) IsLetterOrNumber_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar16isLetterOrNumberEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] bool isDigit() const

/*
Returns true if the character is a decimal digit (Number_DecimalDigit); otherwise returns false.

See also isNumber().
*/
func (this *QChar) IsDigit() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar7isDigitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:553
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isDigit(uint)

/*
Returns true if the character is a decimal digit (Number_DecimalDigit); otherwise returns false.

See also isNumber().
*/
func (this *QChar) IsDigit_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar7isDigitEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] bool isLower() const

/*
Returns true if the character is a lowercase letter, for example category() is Letter_Lowercase.

See also isUpper(), toLower(), and toUpper().
*/
func (this *QChar) IsLower() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar7isLowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:555
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isLower(uint)

/*
Returns true if the character is a lowercase letter, for example category() is Letter_Lowercase.

See also isUpper(), toLower(), and toUpper().
*/
func (this *QChar) IsLower_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar7isLowerEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] bool isUpper() const

/*
Returns true if the character is an uppercase letter, for example category() is Letter_Uppercase.

See also isLower(), toUpper(), and toLower().
*/
func (this *QChar) IsUpper() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar7isUpperEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:557
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isUpper(uint)

/*
Returns true if the character is an uppercase letter, for example category() is Letter_Uppercase.

See also isLower(), toUpper(), and toLower().
*/
func (this *QChar) IsUpper_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar7isUpperEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] bool isTitleCase() const

/*
Returns true if the character is a titlecase letter, for example category() is Letter_Titlecase.

See also isLower(), toUpper(), toLower(), and toTitleCase().
*/
func (this *QChar) IsTitleCase() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar11isTitleCaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:559
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isTitleCase(uint)

/*
Returns true if the character is a titlecase letter, for example category() is Letter_Titlecase.

See also isLower(), toUpper(), toLower(), and toTitleCase().
*/
func (this *QChar) IsTitleCase_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar11isTitleCaseEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] bool isNonCharacter() const

/*
Returns true if the QChar is a non-character; false otherwise.

Unicode has a certain number of code points that are classified as "non-characters:" that is, they can be used for internal purposes in applications but cannot be used for text interchange. Those are the last two entries each Unicode Plane ([0xfffe..0xffff], [0x1fffe..0x1ffff], etc.) as well as the entries in range [0xfdd0..0xfdef].

This function was introduced in  Qt 5.0.
*/
func (this *QChar) IsNonCharacter() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar14isNonCharacterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:467
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isNonCharacter(uint)

/*
Returns true if the QChar is a non-character; false otherwise.

Unicode has a certain number of code points that are classified as "non-characters:" that is, they can be used for internal purposes in applications but cannot be used for text interchange. Those are the last two entries each Unicode Plane ([0xfffe..0xffff], [0x1fffe..0x1ffff], etc.) as well as the entries in range [0xfdd0..0xfdef].

This function was introduced in  Qt 5.0.
*/
func (this *QChar) IsNonCharacter_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar14isNonCharacterEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] bool isHighSurrogate() const

/*
Returns true if the QChar is the high part of a UTF16 surrogate (for example if its code point is in range [0xd800..0xdbff]); false otherwise.
*/
func (this *QChar) IsHighSurrogate() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar15isHighSurrogateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:471
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isHighSurrogate(uint)

/*
Returns true if the QChar is the high part of a UTF16 surrogate (for example if its code point is in range [0xd800..0xdbff]); false otherwise.
*/
func (this *QChar) IsHighSurrogate_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar15isHighSurrogateEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] bool isLowSurrogate() const

/*
Returns true if the QChar is the low part of a UTF16 surrogate (for example if its code point is in range [0xdc00..0xdfff]); false otherwise.
*/
func (this *QChar) IsLowSurrogate() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar14isLowSurrogateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:475
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isLowSurrogate(uint)

/*
Returns true if the QChar is the low part of a UTF16 surrogate (for example if its code point is in range [0xdc00..0xdfff]); false otherwise.
*/
func (this *QChar) IsLowSurrogate_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar14isLowSurrogateEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] bool isSurrogate() const

/*
Returns true if the QChar contains a code point that is in either the high or the low part of the UTF-16 surrogate range (for example if its code point is in range [0xd800..0xdfff]); false otherwise.

This function was introduced in  Qt 5.0.
*/
func (this *QChar) IsSurrogate() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar11isSurrogateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qchar.h:479
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool isSurrogate(uint)

/*
Returns true if the QChar contains a code point that is in either the high or the low part of the UTF-16 surrogate range (for example if its code point is in range [0xd800..0xdfff]); false otherwise.

This function was introduced in  Qt 5.0.
*/
func (this *QChar) IsSurrogate_1(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar11isSurrogateEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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
// [1] uchar cell() const

/*
Returns the cell (least significant byte) of the Unicode character.

See also row().
*/
func (this *QChar) Cell() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar4cellEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
	// unsigned char // 222
}

// /usr/include/qt/QtCore/qchar.h:463
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar row() const

/*
Returns the row (most significant byte) of the Unicode character.

See also cell().
*/
func (this *QChar) Row() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QChar3rowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
	// unsigned char // 222
}

// /usr/include/qt/QtCore/qchar.h:464
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCell(uchar)

/*

 */
func (this *QChar) SetCell(acell byte) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar7setCellEh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), acell)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:465
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRow(uchar)

/*

 */
func (this *QChar) SetRow(arow byte) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar6setRowEh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arow)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qchar.h:483
// index:0
// Public static inline Visibility=Default Availability=Available
// [1] bool requiresSurrogates(uint)

/*
Returns true if the UCS-4-encoded character specified by ucs4 can be split into the high and low parts of a UTF16 surrogate (for example if its code point is greater than or equals to 0x10000); false otherwise.
*/
func (this *QChar) RequiresSurrogates(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar18requiresSurrogatesEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
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

/*
Converts a UTF16 surrogate pair with the given high and low values to it's UCS-4-encoded code point.
*/
func (this *QChar) SurrogateToUcs4(high uint16, low uint16) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar15surrogateToUcs4Ett", qtrt.FFI_TYPE_POINTER, high, low)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
	// unsigned int // 222
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

/*
Converts a UTF16 surrogate pair with the given high and low values to it's UCS-4-encoded code point.
*/
func (this *QChar) SurrogateToUcs4_1(high QChar_ITF /*123*/, low QChar_ITF /*123*/) uint {
	var convArg0 unsafe.Pointer
	if high != nil && high.QChar_PTR() != nil {
		convArg0 = high.QChar_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if low != nil && low.QChar_PTR() != nil {
		convArg1 = low.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar15surrogateToUcs4ES_S_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
	// unsigned int // 222
}
func QChar_SurrogateToUcs4_1(high QChar_ITF /*123*/, low QChar_ITF /*123*/) uint {
	var nilthis *QChar
	rv := nilthis.SurrogateToUcs4_1(high, low)
	return rv
}

// /usr/include/qt/QtCore/qchar.h:495
// index:0
// Public static inline Visibility=Default Availability=Available
// [2] ushort highSurrogate(uint)

/*
Returns the high surrogate part of a UCS-4-encoded code point. The returned result is undefined if ucs4 is smaller than 0x10000.
*/
func (this *QChar) HighSurrogate(ucs4 uint) uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar13highSurrogateEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
	// unsigned short // 222
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

/*
Returns the low surrogate part of a UCS-4-encoded code point. The returned result is undefined if ucs4 is smaller than 0x10000.
*/
func (this *QChar) LowSurrogate(ucs4 uint) uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar12lowSurrogateEj", qtrt.FFI_TYPE_POINTER, ucs4)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
	// unsigned short // 222
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

/*
Returns the most recent supported Unicode version.
*/
func (this *QChar) CurrentUnicodeVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QChar21currentUnicodeVersionEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QChar_CurrentUnicodeVersion() int {
	var nilthis *QChar
	rv := nilthis.CurrentUnicodeVersion()
	return rv
}

func DeleteQChar(this *QChar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QCharD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*

 */
type QChar__SpecialCharacter = int

//
const QChar__Null QChar__SpecialCharacter = 0

//
const QChar__Tabulation QChar__SpecialCharacter = 9

//
const QChar__LineFeed QChar__SpecialCharacter = 10

//
const QChar__CarriageReturn QChar__SpecialCharacter = 13

//
const QChar__Space QChar__SpecialCharacter = 32

//
const QChar__Nbsp QChar__SpecialCharacter = 160

//
const QChar__SoftHyphen QChar__SpecialCharacter = 173

// xfffdThe character shown when a font has no glyph for a certain codepoint. A special question mark character is often used. Codecs use this codepoint when input data cannot be represented in Unicode.
const QChar__ReplacementCharacter QChar__SpecialCharacter = 65533

// xfffcUsed to represent an object such as an image when such objects cannot be presented.
const QChar__ObjectReplacementCharacter QChar__SpecialCharacter = 65532

// xfeff
const QChar__ByteOrderMark QChar__SpecialCharacter = 65279

// xfffe
const QChar__ByteOrderSwapped QChar__SpecialCharacter = 65534

//
const QChar__ParagraphSeparator QChar__SpecialCharacter = 8233

//
const QChar__LineSeparator QChar__SpecialCharacter = 8232

//
const QChar__LastValidCodePoint QChar__SpecialCharacter = 1114111

/*
This enum maps the Unicode character categories.

The following characters are normative in Unicode:



The following categories are informative in Unicode:



See also category().

*/
type QChar__Category = int

// Unicode class name Mn
const QChar__Mark_NonSpacing QChar__Category = 0

// Unicode class name Mc
const QChar__Mark_SpacingCombining QChar__Category = 1

// Unicode class name Me
const QChar__Mark_Enclosing QChar__Category = 2

// Unicode class name Nd
const QChar__Number_DecimalDigit QChar__Category = 3

// Unicode class name Nl
const QChar__Number_Letter QChar__Category = 4

// Unicode class name No
const QChar__Number_Other QChar__Category = 5

// Unicode class name Zs
const QChar__Separator_Space QChar__Category = 6

// Unicode class name Zl
const QChar__Separator_Line QChar__Category = 7

// Unicode class name Zp
const QChar__Separator_Paragraph QChar__Category = 8

// Unicode class name Cc
const QChar__Other_Control QChar__Category = 9

//
const QChar__Other_Format QChar__Category = 10

//
const QChar__Other_Surrogate QChar__Category = 11

//
const QChar__Other_PrivateUse QChar__Category = 12

//
const QChar__Other_NotAssigned QChar__Category = 13

//
const QChar__Letter_Uppercase QChar__Category = 14

//
const QChar__Letter_Lowercase QChar__Category = 15

//
const QChar__Letter_Titlecase QChar__Category = 16

//
const QChar__Letter_Modifier QChar__Category = 17

//
const QChar__Letter_Other QChar__Category = 18

//
const QChar__Punctuation_Connector QChar__Category = 19

//
const QChar__Punctuation_Dash QChar__Category = 20

//
const QChar__Punctuation_Open QChar__Category = 21

//
const QChar__Punctuation_Close QChar__Category = 22

//
const QChar__Punctuation_InitialQuote QChar__Category = 23

//
const QChar__Punctuation_FinalQuote QChar__Category = 24

//
const QChar__Punctuation_Other QChar__Category = 25

//
const QChar__Symbol_Math QChar__Category = 26

//
const QChar__Symbol_Currency QChar__Category = 27

//
const QChar__Symbol_Modifier QChar__Category = 28

//
const QChar__Symbol_Other QChar__Category = 29

/*
This enum type defines the Unicode script property values.

For details about the Unicode script property values see Unicode Standard Annex #24.

In order to conform to C/C++ naming conventions "Script_" is prepended to the codes used in the Unicode Standard.



This enum was introduced or modified in  Qt 5.1.

See also script().

*/
type QChar__Script = int

// For unassigned, private-use, noncharacter, and surrogate code points.
const QChar__Script_Unknown QChar__Script = 0

// For characters that may be used with multiple scripts and that inherit their script from the preceding characters. These include nonspacing marks, enclosing marks, and zero width joiner/non-joiner characters.
const QChar__Script_Inherited QChar__Script = 1

// For characters that may be used with multiple scripts and that do not inherit their script from the preceding characters.
const QChar__Script_Common QChar__Script = 2

//
const QChar__Script_Latin QChar__Script = 3

//
const QChar__Script_Greek QChar__Script = 4

//
const QChar__Script_Cyrillic QChar__Script = 5

//
const QChar__Script_Armenian QChar__Script = 6

//
const QChar__Script_Hebrew QChar__Script = 7

//
const QChar__Script_Arabic QChar__Script = 8

//
const QChar__Script_Syriac QChar__Script = 9

//
const QChar__Script_Thaana QChar__Script = 10

//
const QChar__Script_Devanagari QChar__Script = 11

//
const QChar__Script_Bengali QChar__Script = 12

//
const QChar__Script_Gurmukhi QChar__Script = 13

//
const QChar__Script_Gujarati QChar__Script = 14

//
const QChar__Script_Oriya QChar__Script = 15

//
const QChar__Script_Tamil QChar__Script = 16

//
const QChar__Script_Telugu QChar__Script = 17

//
const QChar__Script_Kannada QChar__Script = 18

//
const QChar__Script_Malayalam QChar__Script = 19

//
const QChar__Script_Sinhala QChar__Script = 20

//
const QChar__Script_Thai QChar__Script = 21

//
const QChar__Script_Lao QChar__Script = 22

//
const QChar__Script_Tibetan QChar__Script = 23

//
const QChar__Script_Myanmar QChar__Script = 24

//
const QChar__Script_Georgian QChar__Script = 25

//
const QChar__Script_Hangul QChar__Script = 26

//
const QChar__Script_Ethiopic QChar__Script = 27

//
const QChar__Script_Cherokee QChar__Script = 28

//
const QChar__Script_CanadianAboriginal QChar__Script = 29

//
const QChar__Script_Ogham QChar__Script = 30

//
const QChar__Script_Runic QChar__Script = 31

//
const QChar__Script_Khmer QChar__Script = 32

//
const QChar__Script_Mongolian QChar__Script = 33

//
const QChar__Script_Hiragana QChar__Script = 34

//
const QChar__Script_Katakana QChar__Script = 35

//
const QChar__Script_Bopomofo QChar__Script = 36

//
const QChar__Script_Han QChar__Script = 37

//
const QChar__Script_Yi QChar__Script = 38

//
const QChar__Script_OldItalic QChar__Script = 39

//
const QChar__Script_Gothic QChar__Script = 40

//
const QChar__Script_Deseret QChar__Script = 41

//
const QChar__Script_Tagalog QChar__Script = 42

//
const QChar__Script_Hanunoo QChar__Script = 43

//
const QChar__Script_Buhid QChar__Script = 44

//
const QChar__Script_Tagbanwa QChar__Script = 45

//
const QChar__Script_Coptic QChar__Script = 46

//
const QChar__Script_Limbu QChar__Script = 47

//
const QChar__Script_TaiLe QChar__Script = 48

//
const QChar__Script_LinearB QChar__Script = 49

//
const QChar__Script_Ugaritic QChar__Script = 50

//
const QChar__Script_Shavian QChar__Script = 51

//
const QChar__Script_Osmanya QChar__Script = 52

//
const QChar__Script_Cypriot QChar__Script = 53

//
const QChar__Script_Braille QChar__Script = 54

//
const QChar__Script_Buginese QChar__Script = 55

//
const QChar__Script_NewTaiLue QChar__Script = 56

//
const QChar__Script_Glagolitic QChar__Script = 57

//
const QChar__Script_Tifinagh QChar__Script = 58

//
const QChar__Script_SylotiNagri QChar__Script = 59

//
const QChar__Script_OldPersian QChar__Script = 60

//
const QChar__Script_Kharoshthi QChar__Script = 61

//
const QChar__Script_Balinese QChar__Script = 62

//
const QChar__Script_Cuneiform QChar__Script = 63

//
const QChar__Script_Phoenician QChar__Script = 64

//
const QChar__Script_PhagsPa QChar__Script = 65

//
const QChar__Script_Nko QChar__Script = 66

//
const QChar__Script_Sundanese QChar__Script = 67

//
const QChar__Script_Lepcha QChar__Script = 68

//
const QChar__Script_OlChiki QChar__Script = 69

//
const QChar__Script_Vai QChar__Script = 70

//
const QChar__Script_Saurashtra QChar__Script = 71

//
const QChar__Script_KayahLi QChar__Script = 72

//
const QChar__Script_Rejang QChar__Script = 73

//
const QChar__Script_Lycian QChar__Script = 74

//
const QChar__Script_Carian QChar__Script = 75

//
const QChar__Script_Lydian QChar__Script = 76

//
const QChar__Script_Cham QChar__Script = 77

//
const QChar__Script_TaiTham QChar__Script = 78

//
const QChar__Script_TaiViet QChar__Script = 79

//
const QChar__Script_Avestan QChar__Script = 80

//
const QChar__Script_EgyptianHieroglyphs QChar__Script = 81

//
const QChar__Script_Samaritan QChar__Script = 82

//
const QChar__Script_Lisu QChar__Script = 83

//
const QChar__Script_Bamum QChar__Script = 84

//
const QChar__Script_Javanese QChar__Script = 85

//
const QChar__Script_MeeteiMayek QChar__Script = 86

//
const QChar__Script_ImperialAramaic QChar__Script = 87

//
const QChar__Script_OldSouthArabian QChar__Script = 88

//
const QChar__Script_InscriptionalParthian QChar__Script = 89

//
const QChar__Script_InscriptionalPahlavi QChar__Script = 90

//
const QChar__Script_OldTurkic QChar__Script = 91

//
const QChar__Script_Kaithi QChar__Script = 92

//
const QChar__Script_Batak QChar__Script = 93

//
const QChar__Script_Brahmi QChar__Script = 94

//
const QChar__Script_Mandaic QChar__Script = 95

//
const QChar__Script_Chakma QChar__Script = 96

//
const QChar__Script_MeroiticCursive QChar__Script = 97

//
const QChar__Script_MeroiticHieroglyphs QChar__Script = 98

//
const QChar__Script_Miao QChar__Script = 99

//
const QChar__Script_Sharada QChar__Script = 100

//
const QChar__Script_SoraSompeng QChar__Script = 101

//
const QChar__Script_Takri QChar__Script = 102

//
const QChar__Script_CaucasianAlbanian QChar__Script = 103

//
const QChar__Script_BassaVah QChar__Script = 104

//
const QChar__Script_Duployan QChar__Script = 105

//
const QChar__Script_Elbasan QChar__Script = 106

//
const QChar__Script_Grantha QChar__Script = 107

//
const QChar__Script_PahawhHmong QChar__Script = 108

//
const QChar__Script_Khojki QChar__Script = 109

//
const QChar__Script_LinearA QChar__Script = 110

//
const QChar__Script_Mahajani QChar__Script = 111

//
const QChar__Script_Manichaean QChar__Script = 112

//
const QChar__Script_MendeKikakui QChar__Script = 113

//
const QChar__Script_Modi QChar__Script = 114

//
const QChar__Script_Mro QChar__Script = 115

//
const QChar__Script_OldNorthArabian QChar__Script = 116

//
const QChar__Script_Nabataean QChar__Script = 117

//
const QChar__Script_Palmyrene QChar__Script = 118

//
const QChar__Script_PauCinHau QChar__Script = 119

//
const QChar__Script_OldPermic QChar__Script = 120

//
const QChar__Script_PsalterPahlavi QChar__Script = 121

//
const QChar__Script_Siddham QChar__Script = 122

//
const QChar__Script_Khudawadi QChar__Script = 123

//
const QChar__Script_Tirhuta QChar__Script = 124

//
const QChar__Script_WarangCiti QChar__Script = 125

//
const QChar__Script_Ahom QChar__Script = 126

//
const QChar__Script_AnatolianHieroglyphs QChar__Script = 127

//
const QChar__Script_Hatran QChar__Script = 128

//
const QChar__Script_Multani QChar__Script = 129

//
const QChar__Script_OldHungarian QChar__Script = 130

//
const QChar__Script_SignWriting QChar__Script = 131

//
const QChar__ScriptCount QChar__Script = 132

/*
This enum type defines the Unicode direction attributes. See the Unicode Standard for a description of the values.

In order to conform to C/C++ naming conventions "Dir" is prepended to the codes used in the Unicode Standard.



See also direction().

*/
type QChar__Direction = int

//
const QChar__DirL QChar__Direction = 0

//
const QChar__DirR QChar__Direction = 1

//
const QChar__DirEN QChar__Direction = 2

//
const QChar__DirES QChar__Direction = 3

//
const QChar__DirET QChar__Direction = 4

//
const QChar__DirAN QChar__Direction = 5

//
const QChar__DirCS QChar__Direction = 6

//
const QChar__DirB QChar__Direction = 7

//
const QChar__DirS QChar__Direction = 8

//
const QChar__DirWS QChar__Direction = 9

//
const QChar__DirON QChar__Direction = 10

//
const QChar__DirLRE QChar__Direction = 11

//
const QChar__DirLRO QChar__Direction = 12

//
const QChar__DirAL QChar__Direction = 13

//
const QChar__DirRLE QChar__Direction = 14

//
const QChar__DirRLO QChar__Direction = 15

//
const QChar__DirPDF QChar__Direction = 16

//
const QChar__DirNSM QChar__Direction = 17

//
const QChar__DirBN QChar__Direction = 18

//
const QChar__DirLRI QChar__Direction = 19

//
const QChar__DirRLI QChar__Direction = 20

//
const QChar__DirFSI QChar__Direction = 21

//
const QChar__DirPDI QChar__Direction = 22

/*
This enum type defines the Unicode decomposition attributes. See the Unicode Standard for a description of the values.

ConstantValue
QChar::NoDecomposition0
QChar::Canonical1
QChar::Circle8
QChar::Final6
QChar::Font2
QChar::Initial4
QChar::Isolated7
QChar::Medial5
QChar::NoBreak3
QChar::Super9


See also decomposition().

*/
type QChar__Decomposition = int

//
const QChar__NoDecomposition QChar__Decomposition = 0

//
const QChar__Canonical QChar__Decomposition = 1

//
const QChar__Font QChar__Decomposition = 2

//
const QChar__NoBreak QChar__Decomposition = 3

//
const QChar__Initial QChar__Decomposition = 4

//
const QChar__Medial QChar__Decomposition = 5

//
const QChar__Final QChar__Decomposition = 6

//
const QChar__Isolated QChar__Decomposition = 7

//
const QChar__Circle QChar__Decomposition = 8

//
const QChar__Super QChar__Decomposition = 9

// 0
const QChar__Sub QChar__Decomposition = 10

// 1
const QChar__Vertical QChar__Decomposition = 11

// 2
const QChar__Wide QChar__Decomposition = 12

// 3
const QChar__Narrow QChar__Decomposition = 13

// 4
const QChar__Small QChar__Decomposition = 14

// 5
const QChar__Square QChar__Decomposition = 15

// 6
const QChar__Compat QChar__Decomposition = 16

// 7
const QChar__Fraction QChar__Decomposition = 17

/*
since 5.3

This enum type defines the Unicode joining type attributes. See the Unicode Standard for a description of the values.

In order to conform to C/C++ naming conventions "Joining_" is prepended to the codes used in the Unicode Standard.

ConstantValue
QChar::Joining_None0
QChar::Joining_Causing1
QChar::Joining_Dual2
QChar::Joining_Right3
QChar::Joining_Left4
QChar::Joining_Transparent5


See also joiningType().

*/
type QChar__JoiningType = int

//
const QChar__Joining_None QChar__JoiningType = 0

//
const QChar__Joining_Causing QChar__JoiningType = 1

//
const QChar__Joining_Dual QChar__JoiningType = 2

//
const QChar__Joining_Right QChar__JoiningType = 3

//
const QChar__Joining_Left QChar__JoiningType = 4

//
const QChar__Joining_Transparent QChar__JoiningType = 5

/*


 */
type QChar__Joining = int

//
const QChar__OtherJoining QChar__Joining = 0

//
const QChar__Dual QChar__Joining = 1

//
const QChar__Right QChar__Joining = 2

//
const QChar__Center QChar__Joining = 3

/*


 */
type QChar__CombiningClass = int

//
const QChar__Combining_BelowLeftAttached QChar__CombiningClass = 200

//
const QChar__Combining_BelowAttached QChar__CombiningClass = 202

//
const QChar__Combining_BelowRightAttached QChar__CombiningClass = 204

//
const QChar__Combining_LeftAttached QChar__CombiningClass = 208

//
const QChar__Combining_RightAttached QChar__CombiningClass = 210

//
const QChar__Combining_AboveLeftAttached QChar__CombiningClass = 212

//
const QChar__Combining_AboveAttached QChar__CombiningClass = 214

//
const QChar__Combining_AboveRightAttached QChar__CombiningClass = 216

//
const QChar__Combining_BelowLeft QChar__CombiningClass = 218

//
const QChar__Combining_Below QChar__CombiningClass = 220

//
const QChar__Combining_BelowRight QChar__CombiningClass = 222

//
const QChar__Combining_Left QChar__CombiningClass = 224

//
const QChar__Combining_Right QChar__CombiningClass = 226

//
const QChar__Combining_AboveLeft QChar__CombiningClass = 228

//
const QChar__Combining_Above QChar__CombiningClass = 230

//
const QChar__Combining_AboveRight QChar__CombiningClass = 232

//
const QChar__Combining_DoubleBelow QChar__CombiningClass = 233

//
const QChar__Combining_DoubleAbove QChar__CombiningClass = 234

//
const QChar__Combining_IotaSubscript QChar__CombiningClass = 240

/*
Specifies which version of the Unicode standard introduced a certain character.



See also unicodeVersion() and currentUnicodeVersion().

*/
type QChar__UnicodeVersion = int

//
const QChar__Unicode_Unassigned QChar__UnicodeVersion = 0

//
const QChar__Unicode_1_1 QChar__UnicodeVersion = 1

//
const QChar__Unicode_2_0 QChar__UnicodeVersion = 2

//
const QChar__Unicode_2_1_2 QChar__UnicodeVersion = 3

//
const QChar__Unicode_3_0 QChar__UnicodeVersion = 4

//
const QChar__Unicode_3_1 QChar__UnicodeVersion = 5

//
const QChar__Unicode_3_2 QChar__UnicodeVersion = 6

//
const QChar__Unicode_4_0 QChar__UnicodeVersion = 7

//
const QChar__Unicode_4_1 QChar__UnicodeVersion = 8

//
const QChar__Unicode_5_0 QChar__UnicodeVersion = 9

//
const QChar__Unicode_5_1 QChar__UnicodeVersion = 10

//
const QChar__Unicode_5_2 QChar__UnicodeVersion = 11

//
const QChar__Unicode_6_0 QChar__UnicodeVersion = 12

//
const QChar__Unicode_6_1 QChar__UnicodeVersion = 13

//
const QChar__Unicode_6_2 QChar__UnicodeVersion = 14

//
const QChar__Unicode_6_3 QChar__UnicodeVersion = 15

//
const QChar__Unicode_7_0 QChar__UnicodeVersion = 16

//
const QChar__Unicode_8_0 QChar__UnicodeVersion = 17

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
