package qtcore

// /usr/include/qt/QtCore/qstring.h
// #include <qstring.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 225
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin

type QCharRef struct {
	*qtrt.CObject
}

func (this *QCharRef) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCharRef) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCharRefFromPointer(cthis unsafe.Pointer) *QCharRef {
	return &QCharRef{&qtrt.CObject{cthis}}
}
func (*QCharRef) NewFromPointer(cthis unsafe.Pointer) *QCharRef {
	return NewQCharRefFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstring.h:1054
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QCharRef) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1055
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isPrint()
func (this *QCharRef) IsPrint() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7isPrintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1056
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isPunct()
func (this *QCharRef) IsPunct() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7isPunctEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1057
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSpace()
func (this *QCharRef) IsSpace() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7isSpaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1058
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isMark()
func (this *QCharRef) IsMark() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef6isMarkEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1059
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isLetter()
func (this *QCharRef) IsLetter() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef8isLetterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1060
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNumber()
func (this *QCharRef) IsNumber() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef8isNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1061
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isLetterOrNumber()
func (this *QCharRef) IsLetterOrNumber() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRef16isLetterOrNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1062
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDigit()
func (this *QCharRef) IsDigit() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7isDigitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1063
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isLower()
func (this *QCharRef) IsLower() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7isLowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1064
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUpper()
func (this *QCharRef) IsUpper() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7isUpperEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1065
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTitleCase()
func (this *QCharRef) IsTitleCase() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef11isTitleCaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1067
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int digitValue()
func (this *QCharRef) DigitValue() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef10digitValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1068
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar toLower()
func (this *QCharRef) ToLower() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7toLowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1069
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar toUpper()
func (this *QCharRef) ToUpper() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7toUpperEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1070
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar toTitleCase()
func (this *QCharRef) ToTitleCase() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef11toTitleCaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1072
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Category category()
func (this *QCharRef) Category() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef8categoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1073
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Direction direction()
func (this *QCharRef) Direction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef9directionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1074
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::JoiningType joiningType()
func (this *QCharRef) JoiningType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef11joiningTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1076
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Joining joining()
func (this *QCharRef) Joining() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7joiningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1089
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool hasMirrored()
func (this *QCharRef) HasMirrored() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef11hasMirroredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1090
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar mirroredChar()
func (this *QCharRef) MirroredChar() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef12mirroredCharEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1091
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString decomposition()
func (this *QCharRef) Decomposition() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef13decompositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:1092
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Decomposition decompositionTag()
func (this *QCharRef) DecompositionTag() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef16decompositionTagEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1093
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar combiningClass()
func (this *QCharRef) CombiningClass() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef14combiningClassEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1095
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Script script()
func (this *QCharRef) Script() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef6scriptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1097
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::UnicodeVersion unicodeVersion()
func (this *QCharRef) UnicodeVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef14unicodeVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1099
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar cell()
func (this *QCharRef) Cell() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef4cellEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1100
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar row()
func (this *QCharRef) Row() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef3rowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1101
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCell(uchar)
func (this *QCharRef) SetCell(cell byte) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRef7setCellEh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cell)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1102
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRow(uchar)
func (this *QCharRef) SetRow(row byte) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRef6setRowEh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1107
// index:0
// Public inline Visibility=Default Availability=Available
// [1] char toLatin1()
func (this *QCharRef) ToLatin1() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef8toLatin1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("byte", rv).(byte) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1108
// index:0
// Public inline Visibility=Default Availability=Available
// [2] ushort unicode()
func (this *QCharRef) Unicode() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7unicodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1109
// index:1
// Public inline Visibility=Default Availability=Available
// [2] ushort & unicode()
func (this *QCharRef) Unicode_1() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRef7unicodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return uint16(rv)
}

func DeleteQCharRef(this *QCharRef) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRefD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
