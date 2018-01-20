//  header block begin
// /usr/include/qt/QtCore/qbytearraymatcher.h
// #include <qbytearraymatcher.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QStaticByteArrayMatcherBase struct {
	*qtrt.CObject
}

func (this *QStaticByteArrayMatcherBase) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:93
// index:0
// inline
// void QStaticByteArrayMatcherBase(const char *, uint)
func NewQStaticByteArrayMatcherBase(pattern unsafe.Pointer, n uint) *QStaticByteArrayMatcherBase {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QStaticByteArrayMatcherBaseC2EPKcj", ffiqt.FFI_TYPE_VOID, cthis, pattern, &n)
	gopp.ErrPrint(err, rv)
	gothis := NewQStaticByteArrayMatcherBaseFromPointer(cthis)
	return gothis
}
func NewQStaticByteArrayMatcherBaseFromPointer(cthis unsafe.Pointer) *QStaticByteArrayMatcherBase {
	return &QStaticByteArrayMatcherBase{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qbytearraymatcher.h:98
// index:0
// int indexOfIn(const char *, uint, const char *, int, int)
func (this *QStaticByteArrayMatcherBase) IndexOfIn(needle unsafe.Pointer, nlen uint, haystack unsafe.Pointer, hlen int, from int) {
	// 0: (, needle const char *, nlen uint, haystack const char *, hlen int, from int), (needle, &nlen, haystack, &hlen, &from)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QStaticByteArrayMatcherBase9indexOfInEPKcjS1_ii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), needle, &nlen, haystack, &hlen, &from)
	gopp.ErrPrint(err, rv)
}

//  body block end
