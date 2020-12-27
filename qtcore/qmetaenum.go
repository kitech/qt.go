package qtcore

// /usr/include/qt/QtCore/qmetaobject.h
// #include <qmetaobject.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*
 */
// size 16
type QMetaEnum struct {
	*qtrt.CObject
}
type QMetaEnum_ITF interface {
	QMetaEnum_PTR() *QMetaEnum
}

func (ptr *QMetaEnum) QMetaEnum_PTR() *QMetaEnum { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QMetaEnumFromptr(cthis Voidptr) *QMetaEnum {
	return &QMetaEnum{&qtrt.CObject{cthis}}
}
func (*QMetaEnum) Fromptr(cthis Voidptr) *QMetaEnum {
	return QMetaEnumFromptr(cthis)
}

func DeleteQMetaEnum(this *QMetaEnum) {
	rv, err := qtrt.Qtcc3(859374498, "_ZN9QMetaEnumD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10049() {
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
