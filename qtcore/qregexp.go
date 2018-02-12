package qtcore

// /usr/include/qt/QtCore/qregexp.h
// #include <qregexp.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QRegExp struct {
	*qtrt.CObject
}

func (this *QRegExp) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRegExp) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRegExpFromPointer(cthis unsafe.Pointer) *QRegExp {
	return &QRegExp{&qtrt.CObject{cthis}}
}
func (*QRegExp) NewFromPointer(cthis unsafe.Pointer) *QRegExp {
	return NewQRegExpFromPointer(cthis)
}

// /usr/include/qt/QtCore/qregexp.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRegExp()
func NewQRegExp() *QRegExp {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExpC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegExpFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegExp)
	return gothis
}

// /usr/include/qt/QtCore/qregexp.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRegExp(const QString &, Qt::CaseSensitivity, enum QRegExp::PatternSyntax)
func NewQRegExp_1(pattern string, cs int, syntax int) *QRegExp {
	var tmpArg0 = NewQString_5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExpC2ERK7QStringN2Qt15CaseSensitivityENS_13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, convArg0, cs, syntax)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegExpFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegExp)
	return gothis
}

// /usr/include/qt/QtCore/qregexp.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QRegExp()
func DeleteQRegExp(this *QRegExp) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExpD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qregexp.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QRegExp &)
func (this *QRegExp) Swap(other *QRegExp) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QRegExp) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregexp.h:85
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QRegExp) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregexp.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QString pattern()
func (this *QRegExp) Pattern() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp7patternEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregexp.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPattern(const QString &)
func (this *QRegExp) SetPattern(pattern string) {
	var tmpArg0 = NewQString_5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp10setPatternERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CaseSensitivity caseSensitivity()
func (this *QRegExp) CaseSensitivity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp15caseSensitivityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qregexp.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCaseSensitivity(Qt::CaseSensitivity)
func (this *QRegExp) SetCaseSensitivity(cs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp18setCaseSensitivityEN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] QRegExp::PatternSyntax patternSyntax()
func (this *QRegExp) PatternSyntax() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp13patternSyntaxEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qregexp.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPatternSyntax(enum QRegExp::PatternSyntax)
func (this *QRegExp) SetPatternSyntax(syntax int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp16setPatternSyntaxENS_13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), syntax)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:93
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMinimal()
func (this *QRegExp) IsMinimal() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp9isMinimalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregexp.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimal(_Bool)
func (this *QRegExp) SetMinimal(minimal bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp10setMinimalEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minimal)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool exactMatch(const QString &)
func (this *QRegExp) ExactMatch(str string) bool {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp10exactMatchERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregexp.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexIn(const QString &, int, enum QRegExp::CaretMode)
func (this *QRegExp) IndexIn(str string, offset int, caretMode int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp7indexInERK7QStringiNS_9CaretModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, caretMode)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastIndexIn(const QString &, int, enum QRegExp::CaretMode)
func (this *QRegExp) LastIndexIn(str string, offset int, caretMode int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp11lastIndexInERK7QStringiNS_9CaretModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, offset, caretMode)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int matchedLength()
func (this *QRegExp) MatchedLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp13matchedLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] int captureCount()
func (this *QRegExp) CaptureCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp12captureCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList capturedTexts()
func (this *QRegExp) CapturedTexts() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp13capturedTextsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qregexp.h:104
// index:1
// Public Visibility=Default Availability=Available
// [8] QStringList capturedTexts()
func (this *QRegExp) CapturedTexts_1() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp13capturedTextsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qregexp.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QString cap(int)
func (this *QRegExp) Cap(nth int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp3capEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregexp.h:106
// index:1
// Public Visibility=Default Availability=Available
// [8] QString cap(int)
func (this *QRegExp) Cap_1(nth int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp3capEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregexp.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] int pos(int)
func (this *QRegExp) Pos(nth int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp3posEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:108
// index:1
// Public Visibility=Default Availability=Available
// [4] int pos(int)
func (this *QRegExp) Pos_1(nth int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp3posEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregexp.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QRegExp) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QRegExp11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregexp.h:110
// index:1
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QRegExp) ErrorString_1() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregexp.h:113
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString escape(const QString &)
func (this *QRegExp) Escape(str string) string {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QRegExp6escapeERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QRegExp_Escape(str string) string {
	var nilthis *QRegExp
	rv := nilthis.Escape(str)
	return rv
}

type QRegExp__PatternSyntax = int

const QRegExp__RegExp QRegExp__PatternSyntax = 0
const QRegExp__Wildcard QRegExp__PatternSyntax = 1
const QRegExp__FixedString QRegExp__PatternSyntax = 2
const QRegExp__RegExp2 QRegExp__PatternSyntax = 3
const QRegExp__WildcardUnix QRegExp__PatternSyntax = 4
const QRegExp__W3CXmlSchema11 QRegExp__PatternSyntax = 5

type QRegExp__CaretMode = int

const QRegExp__CaretAtZero QRegExp__CaretMode = 0
const QRegExp__CaretAtOffset QRegExp__CaretMode = 1
const QRegExp__CaretWontMatch QRegExp__CaretMode = 2

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
		qtrt.KeepMe()
	}
}

//  keep block end
