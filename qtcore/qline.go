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
// extern C begin: 6
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
type QLine struct {
	*qtrt.CObject
}
type QLine_ITF interface {
	QLine_PTR() *QLine
}

func (ptr *QLine) QLine_PTR() *QLine { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QLineFromptr(cthis Voidptr) *QLine {
	return &QLine{&qtrt.CObject{cthis}}
}
func (*QLine) Fromptr(cthis Voidptr) *QLine {
	return QLineFromptr(cthis)
}

func DeleteQLine(this *QLine) {
	rv, err := qtrt.Qtcc3(2630803797, "_ZN5QLineD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10043() {
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
