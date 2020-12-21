package qtcore

// /usr/include/qt/QtCore/qobjectdefs.h
// #include <qobjectdefs.h>
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
// size 16
type QGenericArgument struct {
	*qtrt.CObject
}
type QGenericArgument_ITF interface {
	QGenericArgument_PTR() *QGenericArgument
}

func (ptr *QGenericArgument) QGenericArgument_PTR() *QGenericArgument { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QGenericArgumentFromptr(cthis Voidptr) *QGenericArgument {
	return &QGenericArgument{&qtrt.CObject{cthis}}
}
func (*QGenericArgument) Fromptr(cthis Voidptr) *QGenericArgument {
	return QGenericArgumentFromptr(cthis)
}

func DeleteQGenericArgument(this *QGenericArgument) {
	rv, err := qtrt.Qtcc1(0, "_ZN16QGenericArgumentD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10001() {
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
