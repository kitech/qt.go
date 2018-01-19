//  header block begin
// /usr/include/qt/QtCore/qobjectdefs.h
// #include <qobjectdefs.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QGenericReturnArgument struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qobjectdefs.h:310
// index:0
// inline
// void QGenericReturnArgument(const char *, void *)
func NewQGenericReturnArgument(aName unsafe.Pointer, aData unsafe.Pointer) *QGenericReturnArgument {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGenericReturnArgumentC2EPKcPv", ffiqt.FFI_TYPE_VOID, cthis, aName, aData)
	gopp.ErrPrint(err, rv)
	return &QGenericReturnArgument{cthis}
}

//  body block end
