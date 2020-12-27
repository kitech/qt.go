package qtcore

// /usr/include/qt/QtCore/qline.h
// #include <qline.h>
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
type QLineF struct {
	*qtrt.CObject
}
type QLineF_ITF interface {
	QLineF_PTR() *QLineF
}

func (ptr *QLineF) QLineF_PTR() *QLineF { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QLineFFromptr(cthis Voidptr) *QLineF {
	return &QLineF{&qtrt.CObject{cthis}}
}
func (*QLineF) Fromptr(cthis Voidptr) *QLineF {
	return QLineFFromptr(cthis)
}

func DeleteQLineF(this *QLineF) {
	rv, err := qtrt.Qtcc3(2019197402, "_ZN6QLineFD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QLineF__IntersectType = int

//
const QLineF__NoIntersection QLineF__IntersectType = 0

//
const QLineF__BoundedIntersection QLineF__IntersectType = 1

//
const QLineF__UnboundedIntersection QLineF__IntersectType = 2

func (this *QLineF) IntersectTypeItemName(val int) string {
	switch val {
	case QLineF__NoIntersection: // 0
		return "NoIntersection"
	case QLineF__BoundedIntersection: // 1
		return "BoundedIntersection"
	case QLineF__UnboundedIntersection: // 2
		return "UnboundedIntersection"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QLineF_IntersectTypeItemName(val int) string {
	var nilthis *QLineF
	return nilthis.IntersectTypeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10045() {
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
