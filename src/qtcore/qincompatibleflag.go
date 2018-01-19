//  header block begin
// /usr/include/qt/QtCore/qflags.h
// #include <qflags.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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
type QIncompatibleFlag struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qflags.h:80
// index:0
// inline
// void QIncompatibleFlag(int)
func NewQIncompatibleFlag(i int) *QIncompatibleFlag {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QIncompatibleFlagC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &i)
	gopp.ErrPrint(err, rv)
	return &QIncompatibleFlag{cthis}
}

//  body block end
