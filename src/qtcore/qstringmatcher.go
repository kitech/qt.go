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
	*qtrt.CObject
}

func (this *QStringMatcher) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQStringMatcherFromPointer(cthis unsafe.Pointer) *QStringMatcher {
	return &QStringMatcher{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qstringmatcher.h:53
// index:0
// Public
// void QStringMatcher()
func NewQStringMatcher() *QStringMatcher {
	cthis := qtrt.Calloc(1, 256) // 1048
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStringMatcherC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringMatcherFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstringmatcher.h:54
// index:1
// Public
// void QStringMatcher(const class QString &, Qt::CaseSensitivity)
func NewQStringMatcher_1(pattern *QString, cs int) *QStringMatcher {
	cthis := qtrt.Calloc(1, 256) // 1048
	var convArg0 = pattern.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStringMatcherC2ERK7QStringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &cs)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringMatcherFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstringmatcher.h:56
// index:2
// Public
// void QStringMatcher(const class QChar *, int, Qt::CaseSensitivity)
func NewQStringMatcher_2(uc unsafe.Pointer, len int, cs int) *QStringMatcher {
	cthis := qtrt.Calloc(1, 256) // 1048
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStringMatcherC2EPK5QChariN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, cthis, uc, &len, &cs)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringMatcherFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstringmatcher.h:59
// index:0
// Public
// void ~QStringMatcher()
func DeleteQStringMatcher(*QStringMatcher) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStringMatcherD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringmatcher.h:63
// index:0
// Public
// void setPattern(const class QString &)
func (this *QStringMatcher) SetPattern(pattern *QString) {
	var convArg0 = pattern.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStringMatcher10setPatternERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringmatcher.h:64
// index:0
// Public
// void setCaseSensitivity(Qt::CaseSensitivity)
func (this *QStringMatcher) SetCaseSensitivity(cs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStringMatcher18setCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringmatcher.h:66
// index:0
// Public
// int indexIn(const class QString &, int)
func (this *QStringMatcher) IndexIn(str *QString, from int) interface{} {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStringMatcher7indexInERK7QStringi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringmatcher.h:67
// index:1
// Public
// int indexIn(const class QChar *, int, int)
func (this *QStringMatcher) IndexIn_1(str unsafe.Pointer, length int, from int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStringMatcher7indexInEPK5QCharii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), str, &length, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringmatcher.h:68
// index:0
// Public
// QString pattern()
func (this *QStringMatcher) Pattern() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStringMatcher7patternEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringmatcher.h:69
// index:0
// Public inline
// Qt::CaseSensitivity caseSensitivity()
func (this *QStringMatcher) CaseSensitivity() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStringMatcher15caseSensitivityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
