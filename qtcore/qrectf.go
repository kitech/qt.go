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
type QRectF struct {
	*qtrt.CObject
}
type QRectF_ITF interface {
	QRectF_PTR() *QRectF
}

func (ptr *QRectF) QRectF_PTR() *QRectF { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QRectFFromptr(cthis Voidptr) *QRectF {
	return &QRectF{&qtrt.CObject{cthis}}
}
func (*QRectF) Fromptr(cthis Voidptr) *QRectF {
	return QRectFFromptr(cthis)
}

func DeleteQRectF(this *QRectF) {
	rv, err := qtrt.Qtcc3(3802840456, "_ZN6QRectFD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10053() {
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
