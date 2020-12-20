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
// extern C begin: 0
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
// size 32
type QMetaProperty struct {
	*qtrt.CObject
}
type QMetaProperty_ITF interface {
	QMetaProperty_PTR() *QMetaProperty
}

func (ptr *QMetaProperty) QMetaProperty_PTR() *QMetaProperty { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
func QMetaPropertyFromptr(cthis Voidptr) *QMetaProperty {
	return &QMetaProperty{&qtrt.CObject{cthis}}
}
func (*QMetaProperty) Fromptr(cthis Voidptr) *QMetaProperty {
	return QMetaPropertyFromptr(cthis)
}

func DeleteQMetaProperty(this *QMetaProperty) {
	rv, err := qtrt.Qtcc1(0, "_ZN13QMetaPropertyD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10035() {
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
