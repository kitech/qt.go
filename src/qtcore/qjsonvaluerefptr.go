//  header block begin
// /usr/include/qt/QtCore/qjsonvalue.h
// #include <qjsonvalue.h>
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
type QJsonValueRefPtr struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qjsonvalue.h:237
// index:0
// inline
// void QJsonValueRefPtr(class QJsonArray *, int)
func NewQJsonValueRefPtr(array unsafe.Pointer, idx int) *QJsonValueRefPtr {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QJsonValueRefPtrC2EP10QJsonArrayi", ffiqt.FFI_TYPE_VOID, cthis, array, &idx)
	gopp.ErrPrint(err, rv)
	return &QJsonValueRefPtr{cthis}
}

// /usr/include/qt/QtCore/qjsonvalue.h:239
// index:1
// inline
// void QJsonValueRefPtr(class QJsonObject *, int)
func NewQJsonValueRefPtr_1(object unsafe.Pointer, idx int) *QJsonValueRefPtr {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QJsonValueRefPtrC2EP11QJsonObjecti", ffiqt.FFI_TYPE_VOID, cthis, object, &idx)
	gopp.ErrPrint(err, rv)
	return &QJsonValueRefPtr{cthis}
}

//  body block end
