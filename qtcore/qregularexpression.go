package qtcore

// /usr/include/qt/QtCore/qregularexpression.h
// #include <qregularexpression.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 68
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

type QRegularExpression struct {
	*qtrt.CObject
}

func (this *QRegularExpression) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRegularExpression) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRegularExpressionFromPointer(cthis unsafe.Pointer) *QRegularExpression {
	return &QRegularExpression{&qtrt.CObject{cthis}}
}
func (*QRegularExpression) NewFromPointer(cthis unsafe.Pointer) *QRegularExpression {
	return NewQRegularExpressionFromPointer(cthis)
}

// /usr/include/qt/QtCore/qregularexpression.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] QRegularExpression::PatternOptions patternOptions()
func (this *QRegularExpression) PatternOptions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression14patternOptionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRegularExpression()
func NewQRegularExpression() *QRegularExpression {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpressionC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegularExpression)
	return gothis
}

// /usr/include/qt/QtCore/qregularexpression.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QRegularExpression()
func DeleteQRegularExpression(this *QRegularExpression) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpressionD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qregularexpression.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QRegularExpression &)
func (this *QRegularExpression) Swap(other *QRegularExpression) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpression4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QString pattern()
func (this *QRegularExpression) Pattern() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression7patternEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPattern(const QString &)
func (this *QRegularExpression) SetPattern(pattern *QString) {
	var convArg0 = pattern.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpression10setPatternERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QRegularExpression) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int patternErrorOffset()
func (this *QRegularExpression) PatternErrorOffset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression18patternErrorOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QRegularExpression) ErrorString() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int captureCount()
func (this *QRegularExpression) CaptureCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression12captureCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void optimize()
func (this *QRegularExpression) Optimize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression8optimizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:143
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString escape(const QString &)
func (this *QRegularExpression) Escape(str *QString) *QString /*123*/ {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpression6escapeERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}
func QRegularExpression_Escape(str *QString) *QString /*123*/ {
	var nilthis *QRegularExpression
	rv := nilthis.Escape(str)
	return rv
}

type QRegularExpression__PatternOption = int

const QRegularExpression__NoPatternOption QRegularExpression__PatternOption = 0
const QRegularExpression__CaseInsensitiveOption QRegularExpression__PatternOption = 1
const QRegularExpression__DotMatchesEverythingOption QRegularExpression__PatternOption = 2
const QRegularExpression__MultilineOption QRegularExpression__PatternOption = 4
const QRegularExpression__ExtendedPatternSyntaxOption QRegularExpression__PatternOption = 8
const QRegularExpression__InvertedGreedinessOption QRegularExpression__PatternOption = 16
const QRegularExpression__DontCaptureOption QRegularExpression__PatternOption = 32
const QRegularExpression__UseUnicodePropertiesOption QRegularExpression__PatternOption = 64
const QRegularExpression__OptimizeOnFirstUsageOption QRegularExpression__PatternOption = 128
const QRegularExpression__DontAutomaticallyOptimizeOption QRegularExpression__PatternOption = 256

type QRegularExpression__MatchType = int

const QRegularExpression__NormalMatch QRegularExpression__MatchType = 0
const QRegularExpression__PartialPreferCompleteMatch QRegularExpression__MatchType = 1
const QRegularExpression__PartialPreferFirstMatch QRegularExpression__MatchType = 2
const QRegularExpression__NoMatch QRegularExpression__MatchType = 3

type QRegularExpression__MatchOption = int

const QRegularExpression__NoMatchOption QRegularExpression__MatchOption = 0
const QRegularExpression__AnchoredMatchOption QRegularExpression__MatchOption = 1
const QRegularExpression__DontCheckSubjectStringMatchOption QRegularExpression__MatchOption = 2

//  body block end
