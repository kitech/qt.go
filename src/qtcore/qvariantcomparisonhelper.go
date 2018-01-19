//  header block begin
// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 100
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
type QVariantComparisonHelper struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qvariant.h:560
// index:0
// inline
// void QVariantComparisonHelper(const class QVariant &)
func NewQVariantComparisonHelper(var_ unsafe.Pointer) *QVariantComparisonHelper {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QVariantComparisonHelperC2ERK8QVariant", ffiqt.FFI_TYPE_VOID, cthis, var_)
	gopp.ErrPrint(err, rv)
	return &QVariantComparisonHelper{cthis}
}

//  body block end
