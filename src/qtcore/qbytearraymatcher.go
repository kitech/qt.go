//  header block begin
// /usr/include/qt/QtCore/qbytearraymatcher.h
// #include <qbytearraymatcher.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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
type QByteArrayMatcher struct {
	*qtrt.CObject
}

func (this *QByteArrayMatcher) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:53
// index:0
// void QByteArrayMatcher()
func NewQByteArrayMatcher() *QByteArrayMatcher {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QByteArrayMatcherC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQByteArrayMatcherFromPointer(cthis)
	return gothis
}
func NewQByteArrayMatcherFromPointer(cthis unsafe.Pointer) *QByteArrayMatcher {
	return &QByteArrayMatcher{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:54
// index:1
// void QByteArrayMatcher(const class QByteArray &)
func NewQByteArrayMatcher_1(pattern unsafe.Pointer) *QByteArrayMatcher {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QByteArrayMatcherC2ERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, pattern)
	gopp.ErrPrint(err, rv)
	gothis := NewQByteArrayMatcherFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:55
// index:2
// void QByteArrayMatcher(const char *, int)
func NewQByteArrayMatcher_2(pattern unsafe.Pointer, length int) *QByteArrayMatcher {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QByteArrayMatcherC2EPKci", ffiqt.FFI_TYPE_VOID, cthis, pattern, &length)
	gopp.ErrPrint(err, rv)
	gothis := NewQByteArrayMatcherFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:57
// index:0
// void ~QByteArrayMatcher()
func DeleteQByteArrayMatcher(*QByteArrayMatcher) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QByteArrayMatcherD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:61
// index:0
// void setPattern(const class QByteArray &)
func (this *QByteArrayMatcher) SetPattern(pattern unsafe.Pointer) {
	// 0: (, pattern const QByteArray &), (pattern)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QByteArrayMatcher10setPatternERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pattern)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:63
// index:0
// int indexIn(const class QByteArray &, int)
func (this *QByteArrayMatcher) IndexIn(ba unsafe.Pointer, from int) {
	// 0: (, ba const QByteArray &, from int), (ba, &from)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QByteArrayMatcher7indexInERK10QByteArrayi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ba, &from)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:64
// index:1
// int indexIn(const char *, int, int)
func (this *QByteArrayMatcher) IndexIn_1(str unsafe.Pointer, len int, from int) {
	// 1: (, str const char *, len int, from int), (str, &len, &from)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QByteArrayMatcher7indexInEPKcii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), str, &len, &from)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:65
// index:0
// inline
// QByteArray pattern()
func (this *QByteArrayMatcher) Pattern() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QByteArrayMatcher7patternEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
