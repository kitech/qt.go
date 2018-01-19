//  header block begin
// /usr/include/qt/QtCore/qstringmatcher.h
// #include <qstringmatcher.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
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
type QStringMatcher struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qstringmatcher.h:53
// index:0
// void QStringMatcher()
func NewQStringMatcher() *QStringMatcher {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStringMatcherC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QStringMatcher{cthis}
}

// /usr/include/qt/QtCore/qstringmatcher.h:54
// index:1
// void QStringMatcher(const class QString &, Qt::CaseSensitivity)
func NewQStringMatcher_1(pattern unsafe.Pointer, cs int) *QStringMatcher {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStringMatcherC2ERK7QStringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, cthis, pattern, &cs)
	gopp.ErrPrint(err, rv)
	return &QStringMatcher{cthis}
}

// /usr/include/qt/QtCore/qstringmatcher.h:56
// index:2
// void QStringMatcher(const class QChar *, int, Qt::CaseSensitivity)
func NewQStringMatcher_2(uc unsafe.Pointer, len int, cs int) *QStringMatcher {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStringMatcherC2EPK5QChariN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, cthis, uc, &len, &cs)
	gopp.ErrPrint(err, rv)
	return &QStringMatcher{cthis}
}

// /usr/include/qt/QtCore/qstringmatcher.h:59
// index:0
// void ~QStringMatcher()
func DeleteQStringMatcher(*QStringMatcher) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStringMatcherD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringmatcher.h:63
// index:0
// void setPattern(const class QString &)
func (this *QStringMatcher) SetPattern(pattern unsafe.Pointer) {
	// 0: (, const QString & pattern), (pattern)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStringMatcher10setPatternERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, pattern)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringmatcher.h:64
// index:0
// void setCaseSensitivity(Qt::CaseSensitivity)
func (this *QStringMatcher) SetCaseSensitivity(cs int) {
	// 0: (, Qt::CaseSensitivity cs), (&cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStringMatcher18setCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringmatcher.h:66
// index:0
// int indexIn(const class QString &, int)
func (this *QStringMatcher) IndexIn(str unsafe.Pointer, from int) {
	// 0: (, const QString & str, int from), (str, &from)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStringMatcher7indexInERK7QStringi", ffiqt.FFI_TYPE_VOID, this.cthis, str, &from)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringmatcher.h:67
// index:1
// int indexIn(const class QChar *, int, int)
func (this *QStringMatcher) IndexIn_1(str unsafe.Pointer, length int, from int) {
	// 1: (, const QChar * str, int length, int from), (str, &length, &from)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStringMatcher7indexInEPK5QCharii", ffiqt.FFI_TYPE_VOID, this.cthis, str, &length, &from)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringmatcher.h:68
// index:0
// QString pattern()
func (this *QStringMatcher) Pattern() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStringMatcher7patternEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringmatcher.h:69
// index:0
// inline
// Qt::CaseSensitivity caseSensitivity()
func (this *QStringMatcher) CaseSensitivity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStringMatcher15caseSensitivityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
