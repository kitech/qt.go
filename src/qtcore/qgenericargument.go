//  header block begin
// /usr/include/qt/QtCore/qobjectdefs.h
// #include <qobjectdefs.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QGenericArgument struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qobjectdefs.h:297
// index:0
// inline
// void QGenericArgument(const char *, const void *)
func NewQGenericArgument(aName unsafe.Pointer, aData unsafe.Pointer) *QGenericArgument {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QGenericArgumentC2EPKcPKv", ffiqt.FFI_TYPE_VOID, cthis, aName, aData)
	gopp.ErrPrint(err, rv)
	return &QGenericArgument{cthis}
}

// /usr/include/qt/QtCore/qobjectdefs.h:299
// index:0
// inline
// void * data()
func (this *QGenericArgument) Data() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QGenericArgument4dataEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobjectdefs.h:300
// index:0
// inline
// const char * name()
func (this *QGenericArgument) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QGenericArgument4nameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
