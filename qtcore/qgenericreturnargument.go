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
type QGenericReturnArgument struct {
	*QGenericArgument
}
type QGenericReturnArgument_ITF interface {
	QGenericArgument_ITF
	QGenericReturnArgument_PTR() *QGenericReturnArgument
}

func (ptr *QGenericReturnArgument) QGenericReturnArgument_PTR() *QGenericReturnArgument { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QGenericReturnArgumentFromptr(cthis unsafe.Pointer) *QGenericReturnArgument {
	bcthis0 := QGenericArgumentFromptr(cthis)
	return &QGenericReturnArgument{bcthis0}
}
func (*QGenericReturnArgument) Fromptr(cthis unsafe.Pointer) *QGenericReturnArgument {
	return QGenericReturnArgumentFromptr(cthis)
}

func DeleteQGenericReturnArgument(this *QGenericReturnArgument) {
	rv, err := qtrt.Qtcc1(0, "_ZN22QGenericReturnArgumentD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10003() {
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
