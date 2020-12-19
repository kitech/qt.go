package qtcore

// /usr/include/qt/QtCore/qpoint.h
// #include <qpoint.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
// size 8
type QPoint struct {
	*qtrt.CObject
}
type QPoint_ITF interface {
	QPoint_PTR() *QPoint
}

func (ptr *QPoint) QPoint_PTR() *QPoint { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
func QPointFromptr(cthis unsafe.Pointer) *QPoint {
	return &QPoint{&qtrt.CObject{cthis}}
}
func (*QPoint) Fromptr(cthis unsafe.Pointer) *QPoint {
	return QPointFromptr(cthis)
}

func DeleteQPoint(this *QPoint) {
	rv, err := qtrt.Qtcc1(0, "_ZN6QPointD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10029() {
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
