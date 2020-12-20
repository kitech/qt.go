package qtcore

// /usr/include/qt/QtCore/qrect.h
// #include <qrect.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QRect struct {
	*qtrt.CObject
}
type QRect_ITF interface {
	QRect_PTR() *QRect
}

func (ptr *QRect) QRect_PTR() *QRect { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
func QRectFromptr(cthis Voidptr) *QRect {
	return &QRect{&qtrt.CObject{cthis}}
}
func (*QRect) Fromptr(cthis Voidptr) *QRect {
	return QRectFromptr(cthis)
}

func DeleteQRect(this *QRect) {
	rv, err := qtrt.Qtcc1(0, "_ZN5QRectD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10041() {
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
