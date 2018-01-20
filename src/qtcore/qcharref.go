//  header block begin
// /usr/include/qt/QtCore/qstring.h
// #include <qstring.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 226
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
type QCharRef struct {
	*qtrt.CObject
}

func (this *QCharRef) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQCharRefFromPointer(cthis unsafe.Pointer) *QCharRef {
	return &QCharRef{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qstring.h:1054
// index:0
// Public inline
// bool isNull()
func (this *QCharRef) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1055
// index:0
// Public inline
// bool isPrint()
func (this *QCharRef) IsPrint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7isPrintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1056
// index:0
// Public inline
// bool isPunct()
func (this *QCharRef) IsPunct() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7isPunctEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1057
// index:0
// Public inline
// bool isSpace()
func (this *QCharRef) IsSpace() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7isSpaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1058
// index:0
// Public inline
// bool isMark()
func (this *QCharRef) IsMark() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef6isMarkEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1059
// index:0
// Public inline
// bool isLetter()
func (this *QCharRef) IsLetter() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef8isLetterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1060
// index:0
// Public inline
// bool isNumber()
func (this *QCharRef) IsNumber() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef8isNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1061
// index:0
// Public inline
// bool isLetterOrNumber()
func (this *QCharRef) IsLetterOrNumber() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QCharRef16isLetterOrNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1062
// index:0
// Public inline
// bool isDigit()
func (this *QCharRef) IsDigit() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7isDigitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1063
// index:0
// Public inline
// bool isLower()
func (this *QCharRef) IsLower() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7isLowerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1064
// index:0
// Public inline
// bool isUpper()
func (this *QCharRef) IsUpper() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7isUpperEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1065
// index:0
// Public inline
// bool isTitleCase()
func (this *QCharRef) IsTitleCase() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef11isTitleCaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1067
// index:0
// Public inline
// int digitValue()
func (this *QCharRef) DigitValue() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef10digitValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1068
// index:0
// Public inline
// QChar toLower()
func (this *QCharRef) ToLower() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7toLowerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1069
// index:0
// Public inline
// QChar toUpper()
func (this *QCharRef) ToUpper() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7toUpperEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1070
// index:0
// Public inline
// QChar toTitleCase()
func (this *QCharRef) ToTitleCase() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef11toTitleCaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1072
// index:0
// Public inline
// QChar::Category category()
func (this *QCharRef) Category() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef8categoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1073
// index:0
// Public inline
// QChar::Direction direction()
func (this *QCharRef) Direction() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef9directionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1074
// index:0
// Public inline
// QChar::JoiningType joiningType()
func (this *QCharRef) JoiningType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef11joiningTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1076
// index:0
// Public inline
// QChar::Joining joining()
func (this *QCharRef) Joining() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7joiningEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1089
// index:0
// Public inline
// bool hasMirrored()
func (this *QCharRef) HasMirrored() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef11hasMirroredEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1090
// index:0
// Public inline
// QChar mirroredChar()
func (this *QCharRef) MirroredChar() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef12mirroredCharEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1091
// index:0
// Public inline
// QString decomposition()
func (this *QCharRef) Decomposition() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef13decompositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1092
// index:0
// Public inline
// QChar::Decomposition decompositionTag()
func (this *QCharRef) DecompositionTag() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef16decompositionTagEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1093
// index:0
// Public inline
// uchar combiningClass()
func (this *QCharRef) CombiningClass() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef14combiningClassEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1095
// index:0
// Public inline
// QChar::Script script()
func (this *QCharRef) Script() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef6scriptEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1097
// index:0
// Public inline
// QChar::UnicodeVersion unicodeVersion()
func (this *QCharRef) UnicodeVersion() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef14unicodeVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1099
// index:0
// Public inline
// uchar cell()
func (this *QCharRef) Cell() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef4cellEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1100
// index:0
// Public inline
// uchar row()
func (this *QCharRef) Row() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef3rowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1101
// index:0
// Public inline
// void setCell(uchar)
func (this *QCharRef) SetCell(cell byte) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QCharRef7setCellEh", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &cell)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1102
// index:0
// Public inline
// void setRow(uchar)
func (this *QCharRef) SetRow(row byte) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QCharRef6setRowEh", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1107
// index:0
// Public inline
// char toLatin1()
func (this *QCharRef) ToLatin1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef8toLatin1Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1108
// index:0
// Public inline
// ushort unicode()
func (this *QCharRef) Unicode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7unicodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1109
// index:1
// Public inline
// ushort & unicode()
func (this *QCharRef) Unicode_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QCharRef7unicodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
