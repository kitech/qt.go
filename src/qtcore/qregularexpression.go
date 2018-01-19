//  header block begin
// /usr/include/qt/QtCore/qregularexpression.h
// #include <qregularexpression.h>
// #include <QtCore>
package qtcore

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
type QRegularExpression struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qregularexpression.h:81
// index:0
// QRegularExpression::PatternOptions patternOptions()
func (this *QRegularExpression) PatternOptions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression14patternOptionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:84
// index:0
// void QRegularExpression()
func NewQRegularExpression() *QRegularExpression {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpressionC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QRegularExpression{cthis}
}

// /usr/include/qt/QtCore/qregularexpression.h:87
// index:0
// void ~QRegularExpression()
func DeleteQRegularExpression(*QRegularExpression) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpressionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:95
// index:0
// inline
// void swap(class QRegularExpression &)
func (this *QRegularExpression) Swap(other unsafe.Pointer) {
	// 0: (, QRegularExpression & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpression4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:97
// index:0
// QString pattern()
func (this *QRegularExpression) Pattern() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression7patternEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:98
// index:0
// void setPattern(const class QString &)
func (this *QRegularExpression) SetPattern(pattern unsafe.Pointer) {
	// 0: (, const QString & pattern), (pattern)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpression10setPatternERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, pattern)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:100
// index:0
// bool isValid()
func (this *QRegularExpression) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:101
// index:0
// int patternErrorOffset()
func (this *QRegularExpression) PatternErrorOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression18patternErrorOffsetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:102
// index:0
// QString errorString()
func (this *QRegularExpression) ErrorString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression11errorStringEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:104
// index:0
// int captureCount()
func (this *QRegularExpression) CaptureCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression12captureCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:105
// index:0
// QStringList namedCaptureGroups()
func (this *QRegularExpression) NamedCaptureGroups() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression18namedCaptureGroupsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:141
// index:0
// void optimize()
func (this *QRegularExpression) Optimize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression8optimizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:143
// index:0
// static
// QString escape(const class QString &)
func (this *QRegularExpression) Escape(str unsafe.Pointer) {
	// 0: (const QString & str), (str)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpression6escapeERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRegularExpression_Escape(str unsafe.Pointer) {
	// 0: (const QString & str), (str)
	var nilthis *QRegularExpression
	nilthis.Escape(str)
}

//  body block end
