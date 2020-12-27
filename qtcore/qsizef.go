package qtcore

// /usr/include/qt/QtCore/qsize.h
// #include <qsize.h>
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
type QSizeF struct {
	*qtrt.CObject
}
type QSizeF_ITF interface {
	QSizeF_PTR() *QSizeF
}

func (ptr *QSizeF) QSizeF_PTR() *QSizeF { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QSizeFFromptr(cthis Voidptr) *QSizeF {
	return &QSizeF{&qtrt.CObject{cthis}}
}
func (*QSizeF) Fromptr(cthis Voidptr) *QSizeF {
	return QSizeFFromptr(cthis)
}

func DeleteQSizeF(this *QSizeF) {
	rv, err := qtrt.Qtcc3(2561848487, "_ZN6QSizeFD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10057() {
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
